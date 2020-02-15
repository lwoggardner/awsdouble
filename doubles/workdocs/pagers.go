// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package workdocsdouble

import "github.com/aws/aws-sdk-go/aws/request"

var paginators = map[string]*request.Paginator{

	"DescribeDocumentVersions": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "Limit",
		TruncationToken: "",
	},

	"DescribeFolderContents": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "Limit",
		TruncationToken: "",
	},

	"DescribeUsers": &request.Paginator{
		InputTokens:     []string{"Marker"},
		OutputTokens:    []string{"Marker"},
		LimitToken:      "Limit",
		TruncationToken: "",
	},
}
