// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package dynamodbdouble

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/lwoggardner/awsdouble"
)

var waiters = map[string]*awsdouble.Waiter{

	"TableExists": &awsdouble.Waiter{
		OperationName: "DescribeTable",
		Waiter: request.Waiter{
			Name:        "WaitUntilTableExists",
			MaxAttempts: 25,
			Delay:       request.ConstantWaiterDelay(20 * time.Second),
			Acceptors: []request.WaiterAcceptor{
				{
					State:   request.SuccessWaiterState,
					Matcher: request.PathWaiterMatch, Argument: "Table.TableStatus",
					Expected: "ACTIVE",
				},
				{
					State:    request.RetryWaiterState,
					Matcher:  request.ErrorWaiterMatch,
					Expected: "ResourceNotFoundException",
				},
			},
		},
	},

	"TableNotExists": &awsdouble.Waiter{
		OperationName: "DescribeTable",
		Waiter: request.Waiter{
			Name:        "WaitUntilTableNotExists",
			MaxAttempts: 25,
			Delay:       request.ConstantWaiterDelay(20 * time.Second),
			Acceptors: []request.WaiterAcceptor{
				{
					State:    request.SuccessWaiterState,
					Matcher:  request.ErrorWaiterMatch,
					Expected: "ResourceNotFoundException",
				},
			},
		},
	},
}