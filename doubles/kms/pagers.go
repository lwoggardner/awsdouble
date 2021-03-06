// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package kmsdouble

import "github.com/aws/aws-sdk-go/aws/request"

var paginators = map[string]*request.Paginator{

	"ListAliases": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"NextMarker"},
		LimitToken:      "Limit",
		TruncationToken: "Truncated",
	},

	"ListGrants": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"NextMarker"},
		LimitToken:      "Limit",
		TruncationToken: "Truncated",
	},

	"ListKeyPolicies": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"NextMarker"},
		LimitToken:      "Limit",
		TruncationToken: "Truncated",
	},

	"ListKeys": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"NextMarker"},
		LimitToken:      "Limit",
		TruncationToken: "Truncated",
	},
}
