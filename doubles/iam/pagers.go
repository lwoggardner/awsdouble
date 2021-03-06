// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package iamdouble

import "github.com/aws/aws-sdk-go/aws/request"

var paginators = map[string]*request.Paginator{

	"GetAccountAuthorizationDetails": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"GetGroup": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListAccessKeys": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListAccountAliases": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListAttachedGroupPolicies": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListAttachedRolePolicies": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListAttachedUserPolicies": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListEntitiesForPolicy": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListGroupPolicies": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListGroups": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListGroupsForUser": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListInstanceProfiles": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListInstanceProfilesForRole": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListMFADevices": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListPolicies": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListPolicyVersions": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListRolePolicies": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListRoles": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListSSHPublicKeys": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListServerCertificates": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListSigningCertificates": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListUserPolicies": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListUsers": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"ListVirtualMFADevices": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"SimulateCustomPolicy": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},

	"SimulatePrincipalPolicy": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "MaxItems",
		TruncationToken: "IsTruncated",
	},
}
