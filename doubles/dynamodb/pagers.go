// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package dynamodbdouble

import "github.com/aws/aws-sdk-go/aws/request"

var paginators = map[string]*request.Paginator{

	"BatchGetItem": &request.Paginator{
		InputTokens:     []string{"RequestItems"},
		OutputTokens:    []string{"UnprocessedKeys"},
		LimitToken:      "",
		TruncationToken: "",
	},

	"ListContributorInsights": &request.Paginator{
		InputTokens:     []string{"NextToken"},
		OutputTokens:    []string{"NextToken"},
		LimitToken:      "MaxResults",
		TruncationToken: "",
	},

	"ListTables": &request.Paginator{
		InputTokens:     []string{"ExclusiveStartTableName"},
		OutputTokens:    []string{"LastEvaluatedTableName"},
		LimitToken:      "Limit",
		TruncationToken: "",
	},

	"Query": &request.Paginator{
		InputTokens:     []string{"ExclusiveStartKey"},
		OutputTokens:    []string{"LastEvaluatedKey"},
		LimitToken:      "Limit",
		TruncationToken: "",
	},

	"Scan": &request.Paginator{
		InputTokens:     []string{"ExclusiveStartKey"},
		OutputTokens:    []string{"LastEvaluatedKey"},
		LimitToken:      "Limit",
		TruncationToken: "",
	},
}
