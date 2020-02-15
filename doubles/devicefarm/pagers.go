// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package devicefarmdouble

import "github.com/aws/aws-sdk-go/aws/request"

var paginators = map[string]*request.Paginator{

	"GetOfferingStatus": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},

	"ListArtifacts": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},

	"ListDevicePools": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},

	"ListDevices": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},

	"ListJobs": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},

	"ListOfferingTransactions": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},

	"ListOfferings": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},

	"ListProjects": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},

	"ListRuns": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},

	"ListSamples": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},

	"ListSuites": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},

	"ListTestGridProjects": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResult",
		TruncationToken: "",
	},

	"ListTestGridSessionActions": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResult",
		TruncationToken: "",
	},

	"ListTestGridSessionArtifacts": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResult",
		TruncationToken: "",
	},

	"ListTestGridSessions": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "maxResult",
		TruncationToken: "",
	},

	"ListTests": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},

	"ListUniqueProblems": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},

	"ListUploads": &request.Paginator{
		InputTokens:     []string{"nextToken"},
		OutputTokens:    []string{"nextToken"},
		LimitToken:      "",
		TruncationToken: "",
	},
}
