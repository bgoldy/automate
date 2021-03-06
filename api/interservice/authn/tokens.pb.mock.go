// Code generated by protoc-gen-grpc-mock. DO NOT EDIT.
// source: api/interservice/authn/tokens.proto

package authn

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// verify that the mock satisfies the TokensMgmtServer interface (at compile time)
var _ TokensMgmtServer = &TokensMgmtServerMock{}

// NewTokensMgmtServerMock gives you a fresh instance of TokensMgmtServerMock.
func NewTokensMgmtServerMock() *TokensMgmtServerMock {
	return &TokensMgmtServerMock{validateRequests: true}
}

// NewTokensMgmtServerMockWithoutValidation gives you a fresh instance of
// TokensMgmtServerMock which does not attempt to validate requests before passing
// them to their respective '*Func'.
func NewTokensMgmtServerMockWithoutValidation() *TokensMgmtServerMock {
	return &TokensMgmtServerMock{}
}

// TokensMgmtServerMock is the mock-what-you-want struct that stubs all not-overridden
// methods with "not implemented" returns
type TokensMgmtServerMock struct {
	validateRequests         bool
	GetTokensFunc            func(context.Context, *GetTokensReq) (*Tokens, error)
	CreateTokenFunc          func(context.Context, *CreateTokenReq) (*Token, error)
	CreateTokenWithValueFunc func(context.Context, *CreateTokenWithValueReq) (*Token, error)
	UpdateTokenFunc          func(context.Context, *UpdateTokenReq) (*Token, error)
	GetTokenFunc             func(context.Context, *GetTokenReq) (*Token, error)
	DeleteTokenFunc          func(context.Context, *DeleteTokenReq) (*DeleteTokenResp, error)
}

func (m *TokensMgmtServerMock) GetTokens(ctx context.Context, req *GetTokensReq) (*Tokens, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetTokensFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetTokens' not implemented")
}

func (m *TokensMgmtServerMock) CreateToken(ctx context.Context, req *CreateTokenReq) (*Token, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.CreateTokenFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'CreateToken' not implemented")
}

func (m *TokensMgmtServerMock) CreateTokenWithValue(ctx context.Context, req *CreateTokenWithValueReq) (*Token, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.CreateTokenWithValueFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'CreateTokenWithValue' not implemented")
}

func (m *TokensMgmtServerMock) UpdateToken(ctx context.Context, req *UpdateTokenReq) (*Token, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.UpdateTokenFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'UpdateToken' not implemented")
}

func (m *TokensMgmtServerMock) GetToken(ctx context.Context, req *GetTokenReq) (*Token, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetTokenFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetToken' not implemented")
}

func (m *TokensMgmtServerMock) DeleteToken(ctx context.Context, req *DeleteTokenReq) (*DeleteTokenResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.DeleteTokenFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'DeleteToken' not implemented")
}

// Reset resets all overridden functions
func (m *TokensMgmtServerMock) Reset() {
	m.GetTokensFunc = nil
	m.CreateTokenFunc = nil
	m.CreateTokenWithValueFunc = nil
	m.UpdateTokenFunc = nil
	m.GetTokenFunc = nil
	m.DeleteTokenFunc = nil
}
