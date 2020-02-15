// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package codedeploydouble

import "github.com/aws/aws-sdk-go/aws/request"

var paginators = map[string]*request.Paginator{

	"ListApplicationRevisions": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},

	"ListApplications": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},

	"ListDeploymentConfigs": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},

	"ListDeploymentGroups": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},

	"ListDeploymentInstances": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},

	"ListDeployments": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},
}
