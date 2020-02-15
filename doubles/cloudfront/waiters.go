// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package cloudfrontdouble

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/lwoggardner/awsdouble"
)

var waiters = map[string]*awsdouble.Waiter{

	"DistributionDeployed": &awsdouble.Waiter{
		OperationName: "GetDistribution",
		Waiter: request.Waiter{
			Name:        "WaitUntilDistributionDeployed",
			MaxAttempts: 35,
			Delay:       request.ConstantWaiterDelay(60 * time.Second),
			Acceptors: []request.WaiterAcceptor{
				{
					State:   request.SuccessWaiterState,
					Matcher: request.PathWaiterMatch, Argument: "Distribution.Status",
					Expected: "Deployed",
				},
			},
		},
	},

	"InvalidationCompleted": &awsdouble.Waiter{
		OperationName: "GetInvalidation",
		Waiter: request.Waiter{
			Name:        "WaitUntilInvalidationCompleted",
			MaxAttempts: 30,
			Delay:       request.ConstantWaiterDelay(20 * time.Second),
			Acceptors: []request.WaiterAcceptor{
				{
					State:   request.SuccessWaiterState,
					Matcher: request.PathWaiterMatch, Argument: "Invalidation.Status",
					Expected: "Completed",
				},
			},
		},
	},

	"StreamingDistributionDeployed": &awsdouble.Waiter{
		OperationName: "GetStreamingDistribution",
		Waiter: request.Waiter{
			Name:        "WaitUntilStreamingDistributionDeployed",
			MaxAttempts: 25,
			Delay:       request.ConstantWaiterDelay(60 * time.Second),
			Acceptors: []request.WaiterAcceptor{
				{
					State:   request.SuccessWaiterState,
					Matcher: request.PathWaiterMatch, Argument: "StreamingDistribution.Status",
					Expected: "Deployed",
				},
			},
		},
	},
}