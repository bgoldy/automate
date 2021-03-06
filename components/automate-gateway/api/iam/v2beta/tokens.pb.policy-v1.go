// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2beta/tokens.proto

package v2beta

import (
	policy "github.com/chef/automate/components/automate-gateway/api/authz/policy"
	request "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/request"
)

func init() {
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Tokens/CreateToken", "auth:tokens", "create", "POST", "/iam/v2beta/tokens", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateTokenReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "value":
					return m.Value
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Tokens/GetToken", "auth:tokens:{id}", "get", "GET", "/iam/v2beta/tokens/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetTokenReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Tokens/UpdateToken", "auth:tokens:{id}", "update", "PUT", "/iam/v2beta/tokens/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateTokenReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Tokens/DeleteToken", "auth:tokens:{id}", "delete", "DELETE", "/iam/v2beta/tokens/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteTokenReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Tokens/ListTokens", "auth:tokens", "read", "GET", "/iam/v2beta/tokens", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}
