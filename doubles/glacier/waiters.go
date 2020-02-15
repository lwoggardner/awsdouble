// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package glacierdouble

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/lwoggardner/awsdouble"
)

var waiters = map[string]*awsdouble.Waiter{

	"VaultExists": &awsdouble.Waiter{
		OperationName: "DescribeVault",
		Waiter: request.Waiter{
			Name:        "WaitUntilVaultExists",
			MaxAttempts: 15,
			Delay:       request.ConstantWaiterDelay(3 * time.Second),
			Acceptors: []request.WaiterAcceptor{
				{
					State:    request.SuccessWaiterState,
					Matcher:  request.StatusWaiterMatch,
					Expected: 200,
				},
				{
					State:    request.RetryWaiterState,
					Matcher:  request.ErrorWaiterMatch,
					Expected: "ResourceNotFoundException",
				},
			},
		},
	},

	"VaultNotExists": &awsdouble.Waiter{
		OperationName: "DescribeVault",
		Waiter: request.Waiter{
			Name:        "WaitUntilVaultNotExists",
			MaxAttempts: 15,
			Delay:       request.ConstantWaiterDelay(3 * time.Second),
			Acceptors: []request.WaiterAcceptor{
				{
					State:    request.RetryWaiterState,
					Matcher:  request.StatusWaiterMatch,
					Expected: 200,
				},
				{
					State:    request.SuccessWaiterState,
					Matcher:  request.ErrorWaiterMatch,
					Expected: "ResourceNotFoundException",
				},
			},
		},
	},
}
