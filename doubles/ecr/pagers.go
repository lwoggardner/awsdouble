// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package ecrdouble

import "github.com/aws/aws-sdk-go/aws/request"

var paginators = map[string]*request.Paginator{

	"DescribeImageScanFindings": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},

	"DescribeImages": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},

	"DescribeRepositories": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},

	"GetLifecyclePolicyPreview": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},

	"ListImages": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},
}
