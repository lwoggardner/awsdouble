// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package opsworkscmdouble

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/lwoggardner/awsdouble"
)

var waiters = map[string]*awsdouble.Waiter{

	"NodeAssociated": &awsdouble.Waiter{
		OperationName: "DescribeNodeAssociationStatus",
		Waiter: request.Waiter{
			Name:        "WaitUntilNodeAssociated",
			MaxAttempts: 15,
			Delay:       request.ConstantWaiterDelay(15 * time.Second),
			Acceptors: []request.WaiterAcceptor{
				{
					State:   request.SuccessWaiterState,
					Matcher: request.PathWaiterMatch, Argument: "NodeAssociationStatus",
					Expected: "SUCCESS",
				},
				{
					State:   request.FailureWaiterState,
					Matcher: request.PathWaiterMatch, Argument: "NodeAssociationStatus",
					Expected: "FAILED",
				},
			},
		},
	},
}