// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package appstreamdouble

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/lwoggardner/awsdouble"
)

var waiters = map[string]*awsdouble.Waiter{

	"FleetStarted": &awsdouble.Waiter{
		OperationName: "DescribeFleets",
		Waiter: request.Waiter{
			Name:        "WaitUntilFleetStarted",
			MaxAttempts: 40,
			Delay:       request.ConstantWaiterDelay(30 * time.Second),
			Acceptors: []request.WaiterAcceptor{
				{
					State:   request.SuccessWaiterState,
					Matcher: request.PathAllWaiterMatch, Argument: "Fleets[].State",
					Expected: "ACTIVE",
				},
				{
					State:   request.FailureWaiterState,
					Matcher: request.PathAnyWaiterMatch, Argument: "Fleets[].State",
					Expected: "PENDING_DEACTIVATE",
				},
				{
					State:   request.FailureWaiterState,
					Matcher: request.PathAnyWaiterMatch, Argument: "Fleets[].State",
					Expected: "INACTIVE",
				},
			},
		},
	},

	"FleetStopped": &awsdouble.Waiter{
		OperationName: "DescribeFleets",
		Waiter: request.Waiter{
			Name:        "WaitUntilFleetStopped",
			MaxAttempts: 40,
			Delay:       request.ConstantWaiterDelay(30 * time.Second),
			Acceptors: []request.WaiterAcceptor{
				{
					State:   request.SuccessWaiterState,
					Matcher: request.PathAllWaiterMatch, Argument: "Fleets[].State",
					Expected: "INACTIVE",
				},
				{
					State:   request.FailureWaiterState,
					Matcher: request.PathAnyWaiterMatch, Argument: "Fleets[].State",
					Expected: "PENDING_ACTIVATE",
				},
				{
					State:   request.FailureWaiterState,
					Matcher: request.PathAnyWaiterMatch, Argument: "Fleets[].State",
					Expected: "ACTIVE",
				},
			},
		},
	},
}
