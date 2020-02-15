// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package resourcegroupsdouble

import "github.com/aws/aws-sdk-go/aws/request"

var paginators = map[string]*request.Paginator{

	"ListGroupResources": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "MaxResults",
		TruncationToken: "",
	},

	"ListGroups": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "MaxResults",
		TruncationToken: "",
	},

	"SearchResources": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "MaxResults",
		TruncationToken: "",
	},
}
