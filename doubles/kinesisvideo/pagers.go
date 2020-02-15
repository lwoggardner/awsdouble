// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package kinesisvideodouble

import "github.com/aws/aws-sdk-go/aws/request"

var paginators = map[string]*request.Paginator{

	"ListSignalingChannels": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "MaxResults",
		TruncationToken: "",
	},

	"ListStreams": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "MaxResults",
		TruncationToken: "",
	},
}
