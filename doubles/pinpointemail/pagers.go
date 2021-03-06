// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package pinpointemaildouble

import "github.com/aws/aws-sdk-go/aws/request"

var paginators = map[string]*request.Paginator{

	"GetDedicatedIps": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "PageSize",
		TruncationToken: "",
	},

	"ListConfigurationSets": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "PageSize",
		TruncationToken: "",
	},

	"ListDedicatedIpPools": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "PageSize",
		TruncationToken: "",
	},

	"ListDeliverabilityTestReports": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "PageSize",
		TruncationToken: "",
	},

	"ListDomainDeliverabilityCampaigns": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "PageSize",
		TruncationToken: "",
	},

	"ListEmailIdentities": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "PageSize",
		TruncationToken: "",
	},
}
