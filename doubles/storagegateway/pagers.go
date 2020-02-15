// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package storagegatewaydouble

import "github.com/aws/aws-sdk-go/aws/request"

var paginators = map[string]*request.Paginator{

	"DescribeTapeArchives": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "Limit",
		TruncationToken: "",
	},

	"DescribeTapeRecoveryPoints": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "Limit",
		TruncationToken: "",
	},

	"DescribeTapes": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "Limit",
		TruncationToken: "",
	},

	"DescribeVTLDevices": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "Limit",
		TruncationToken: "",
	},

	"ListFileShares": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"NextMarker"},
		LimitToken:      "Limit",
		TruncationToken: "",
	},

	"ListGateways": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "Limit",
		TruncationToken: "",
	},

	"ListTagsForResource": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "Limit",
		TruncationToken: "",
	},

	"ListTapes": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "Limit",
		TruncationToken: "",
	},

	"ListVolumes": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "Limit",
		TruncationToken: "",
	},
}
