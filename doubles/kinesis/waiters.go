// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package kinesisdouble

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/lwoggardner/awsdouble"
)

var waiters = map[string]*awsdouble.Waiter{

	"StreamExists": &awsdouble.Waiter{
		OperationName: "DescribeStream",
		Waiter: request.Waiter{
			Name:        "WaitUntilStreamExists",
			MaxAttempts: 18,
			Delay:       request.ConstantWaiterDelay(10 * time.Second),
			Acceptors: []request.WaiterAcceptor{
				{
					State:   request.SuccessWaiterState,
					Matcher: request.PathWaiterMatch, Argument: "StreamDescription.StreamStatus",
					Expected: "ACTIVE",
				},
			},
		},
	},

	"StreamNotExists": &awsdouble.Waiter{
		OperationName: "DescribeStream",
		Waiter: request.Waiter{
			Name:        "WaitUntilStreamNotExists",
			MaxAttempts: 18,
			Delay:       request.ConstantWaiterDelay(10 * time.Second),
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
