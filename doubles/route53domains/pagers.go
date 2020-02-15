// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package route53domainsdouble

import "github.com/aws/aws-sdk-go/aws/request"

var paginators = map[string]*request.Paginator{

	"ListDomains": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"NextPageMarker"},
		LimitToken:      "MaxItems",
		TruncationToken: "",
	},

	"ListOperations": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"NextPageMarker"},
		LimitToken:      "MaxItems",
		TruncationToken: "",
	},
}