// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package inspectordouble

import "github.com/aws/aws-sdk-go/aws/request"

var paginators = map[string]*request.Paginator{

	"GetExclusionsPreview": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},

	"ListAssessmentRunAgents": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},

	"ListAssessmentRuns": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},

	"ListAssessmentTargets": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},

	"ListAssessmentTemplates": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},

	"ListEventSubscriptions": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},

	"ListExclusions": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},

	"ListFindings": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},

	"ListRulesPackages": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},

	"PreviewAgents": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResults",
		TruncationToken: "",
	},
}
