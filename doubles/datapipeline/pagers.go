// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package datapipelinedouble

import "github.com/aws/aws-sdk-go/aws/request"

var paginators = map[string]*request.Paginator{

	"DescribeObjects": &request.Paginator{
		InputTokens:     []string{"marker"},
		OutputTokens:    []string{"marker"},
		LimitToken:      "",
		TruncationToken: "hasMoreResults",
	},

	"ListPipelines": &request.Paginator{
		InputTokens:     []string{"marker"},
		OutputTokens:    []string{"marker"},
		LimitToken:      "",
		TruncationToken: "hasMoreResults",
	},

	"QueryObjects": &request.Paginator{
		InputTokens:     []string{"marker"},
		OutputTokens:    []string{"marker"},
		LimitToken:      "limit",
		TruncationToken: "hasMoreResults",
	},
}