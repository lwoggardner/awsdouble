// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package elbdouble

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/lwoggardner/awsdouble"
)

var waiters = map[string]*awsdouble.Waiter{

	"AnyInstanceInService": &awsdouble.Waiter{
		OperationName: "DescribeInstanceHealth",
		Waiter: request.Waiter{
			Name:        "WaitUntilAnyInstanceInService",
			MaxAttempts: 40,
			Delay:       request.ConstantWaiterDelay(15 * time.Second),
			Acceptors: []request.WaiterAcceptor{
				{
					State:   request.SuccessWaiterState,
					Matcher: request.PathAnyWaiterMatch, Argument: "InstanceStates[].State",
					Expected: "InService",
				},
			},
		},
	},

	"InstanceDeregistered": &awsdouble.Waiter{
		OperationName: "DescribeInstanceHealth",
		Waiter: request.Waiter{
			Name:        "WaitUntilInstanceDeregistered",
			MaxAttempts: 40,
			Delay:       request.ConstantWaiterDelay(15 * time.Second),
			Acceptors: []request.WaiterAcceptor{
				{
					State:   request.SuccessWaiterState,
					Matcher: request.PathAllWaiterMatch, Argument: "InstanceStates[].State",
					Expected: "OutOfService",
				},
				{
					State:    request.SuccessWaiterState,
					Matcher:  request.ErrorWaiterMatch,
					Expected: "InvalidInstance",
				},
			},
		},
	},

	"InstanceInService": &awsdouble.Waiter{
		OperationName: "DescribeInstanceHealth",
		Waiter: request.Waiter{
			Name:        "WaitUntilInstanceInService",
			MaxAttempts: 40,
			Delay:       request.ConstantWaiterDelay(15 * time.Second),
			Acceptors: []request.WaiterAcceptor{
				{
					State:   request.SuccessWaiterState,
					Matcher: request.PathAllWaiterMatch, Argument: "InstanceStates[].State",
					Expected: "InService",
				},
				{
					State:    request.RetryWaiterState,
					Matcher:  request.ErrorWaiterMatch,
					Expected: "InvalidInstance",
				},
			},
		},
	},
}
