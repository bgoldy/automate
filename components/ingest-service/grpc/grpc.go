package grpc

import (
	"context"
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/reflection"

	iam_v2 "github.com/chef/automate/api/interservice/authz/v2"
	dls "github.com/chef/automate/api/interservice/data_lifecycle"
	"github.com/chef/automate/api/interservice/es_sidecar"
	automate_event "github.com/chef/automate/api/interservice/event"
	"github.com/chef/automate/api/interservice/ingest"
	"github.com/chef/automate/components/ingest-service/backend/elastic"
	"github.com/chef/automate/components/ingest-service/config"
	"github.com/chef/automate/components/ingest-service/migration"
	"github.com/chef/automate/components/ingest-service/server"
	"github.com/chef/automate/components/ingest-service/serveropts"
)

// Spawn starts a gRPC Server listening on the provided host and port,
// it uses the backend URL to initialize the Pipelines
//
// You can start a server in a goroutine to work independently like:
// ```
// go server.Spawn("localhost", "2193", "elasticsearch:9200")
// ```
//
// Maybe even spawn multiple servers
func Spawn(opts *serveropts.Opts) error {
	// Initialize the backend client
	client, err := elastic.New(opts.ElasticSearchUrl)

	if err != nil {
		log.WithFields(log.Fields{
			"url":   opts.ElasticSearchUrl,
			"error": err.Error(),
		}).Error("could not connect to elasticsearch")
		return err
	}

	var (
		grpcServer  = opts.ConnFactory.NewServer()
		A1migration = migration.New(client)
		uri         = fmt.Sprintf("%s:%d", opts.Host, opts.Port)
	)

	log.WithFields(log.Fields{"uri": uri}).Info("Starting gRPC Server")
	conn, err := net.Listen("tcp", uri)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("TCP listen failed")
		return err
	}

	// Register Status Server
	ingestStatus := server.NewIngestStatus(client, A1migration)
	ingest.RegisterIngestStatusServer(grpcServer, ingestStatus)

	// Initialize elasticsearch indices or trigger a migration
	//
	// This initialization will happen inside a goroutine so that we can then
	// report the health of the system. We might have to do the same for the
	// migration tasks.
	migrationNeeded, err := A1migration.MigrationNeeded()
	if err != nil {
		return err
	}

	if migrationNeeded {
		A1migration.Start()
	} else {
		A1migration.MarkUnneeded()
		client.InitializeStore(context.Background())
	}

	// Authz Interface
	authzConn, err := opts.ConnFactory.Dial("authz-service", opts.AuthzAddress)
	if err != nil {
		// This should never happend
		log.WithFields(log.Fields{"error": err}).Fatal("Failed to create Authz connection")
		return err
	}
	defer authzConn.Close()

	authzProjectsClient := iam_v2.NewProjectsClient(authzConn)

	// event Interface
	eventConn, err := opts.ConnFactory.Dial("event-service", opts.EventAddress)
	if err != nil {
		// This should never happend
		log.WithFields(log.Fields{"error": err}).Fatal("Failed to create Event connection")
		return err
	}
	defer eventConn.Close()

	eventServiceClient := automate_event.NewEventServiceClient(eventConn)

	// ChefRuns
	chefIngest := server.NewChefIngestServer(client, authzProjectsClient)
	ingest.RegisterChefIngesterServer(grpcServer, chefIngest)

	// Pass the chef ingest server to give status about the pipelines
	ingestStatus.SetChefIngestServer(chefIngest)

	// JobScheduler
	// Create a job scheduler so we can triggers jobs like mark
	// nodes as missing on a specific threshold every minute
	jobScheduler := server.NewJobScheduler()
	defer jobScheduler.Close()

	configManager := config.NewManager()
	defer configManager.Close()

	// JobSchedulerServer
	jobSchedulerServer := server.NewJobSchedulerServer(client, jobScheduler, configManager)
	ingest.RegisterJobSchedulerServer(grpcServer, jobSchedulerServer)

	// EventHandler
	eventHandlerServer := server.NewAutomateEventHandlerServer(client, *chefIngest,
		authzProjectsClient, eventServiceClient)
	ingest.RegisterEventHandlerServer(grpcServer, eventHandlerServer)

	// Data Lifecycle Interface
	esSidecarConn, err := opts.ConnFactory.Dial("es-sidecar-service", opts.EsSidecarAddress)
	if err != nil {
		// This should never happend
		log.WithFields(log.Fields{"error": err}).Fatal("Failed to create ES Sidecar connection")
		return err
	}
	defer esSidecarConn.Close()

	purgePolicies := []server.PurgePolicy{}
	if opts.PurgeConvergeHistoryAfterDays >= 0 {
		purgePolicies = append(purgePolicies, server.PurgePolicy{
			IndexName:          "converge-history",
			PurgeOlderThanDays: opts.PurgeConvergeHistoryAfterDays,
		})
	}

	if opts.PurgeActionsAfterDays >= 0 {
		purgePolicies = append(purgePolicies, server.PurgePolicy{
			IndexName:          "actions",
			PurgeOlderThanDays: opts.PurgeActionsAfterDays,
		})
	}
	dataLifecycleServer := server.NewDataLifecycleManageableServer(es_sidecar.NewEsSidecarClient(esSidecarConn), purgePolicies)
	dls.RegisterDataLifecycleManageableServer(grpcServer, dataLifecycleServer)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	return grpcServer.Serve(conn)
}
