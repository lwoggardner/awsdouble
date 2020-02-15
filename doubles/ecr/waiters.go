// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package ecrdouble

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/lwoggardner/awsdouble"
)

var waiters = map[string]*awsdouble.Waiter{

	"ImageScanComplete": &awsdouble.Waiter{
		OperationName: "DescribeImageScanFindings",
		Waiter: request.Waiter{
			Name:        "WaitUntilImageScanComplete",
			MaxAttempts: 60,
			Delay:       request.ConstantWaiterDelay(5 * time.Second),
			Acceptors: []request.WaiterAcceptor{
				{
					State:   request.SuccessWaiterState,
					Matcher: request.PathWaiterMatch, Argument: "imageScanStatus.status",
					Expected: "COMPLETE",
				},
				{
					State:   request.FailureWaiterState,
					Matcher: request.PathWaiterMatch, Argument: "imageScanStatus.status",
					Expected: "FAILED",
				},
			},
		},
	},

	"LifecyclePolicyPreviewComplete": &awsdouble.Waiter{
		OperationName: "GetLifecyclePolicyPreview",
		Waiter: request.Waiter{
			Name:        "WaitUntilLifecyclePolicyPreviewComplete",
			MaxAttempts: 20,
			Delay:       request.ConstantWaiterDelay(5 * time.Second),
			Acceptors: []request.WaiterAcceptor{
				{
					State:   request.SuccessWaiterState,
					Matcher: request.PathWaiterMatch, Argument: "status",
					Expected: "COMPLETE",
				},
				{
					State:   request.FailureWaiterState,
					Matcher: request.PathWaiterMatch, Argument: "status",
					Expected: "FAILED",
				},
			},
		},
	},
}