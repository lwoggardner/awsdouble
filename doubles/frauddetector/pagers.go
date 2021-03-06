// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package frauddetectordouble

import "github.com/aws/aws-sdk-go/aws/request"

var paginators = map[string]*request.Paginator{

	"DescribeModelVersions": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},

	"GetDetectors": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},

	"GetExternalModels": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},

	"GetModels": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},

	"GetOutcomes": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},

	"GetRules": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},

	"GetVariables": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},
}
