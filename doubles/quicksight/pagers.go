// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package quicksightdouble

import "github.com/aws/aws-sdk-go/aws/request"

var paginators = map[string]*request.Paginator{

	"ListDashboardVersions": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "MaxResults",
		TruncationToken: "",
	},

	"ListDashboards": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "MaxResults",
		TruncationToken: "",
	},

	"ListDataSets": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "MaxResults",
		TruncationToken: "",
	},

	"ListDataSources": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "MaxResults",
		TruncationToken: "",
	},

	"ListIngestions": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "MaxResults",
		TruncationToken: "",
	},

	"ListTemplateAliases": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "MaxResults",
		TruncationToken: "",
	},

	"ListTemplateVersions": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "MaxResults",
		TruncationToken: "",
	},

	"ListTemplates": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "MaxResults",
		TruncationToken: "",
	},
}