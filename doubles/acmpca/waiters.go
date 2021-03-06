// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package acmpcadouble

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/lwoggardner/awsdouble"
)

var waiters = map[string]*awsdouble.Waiter{

	"AuditReportCreated": &awsdouble.Waiter{
		OperationName: "DescribeCertificateAuthorityAuditReport",
		Waiter: request.Waiter{
			Name:        "WaitUntilAuditReportCreated",
			MaxAttempts: 60,
			Delay:       request.ConstantWaiterDelay(3 * time.Second),
			Acceptors: []request.WaiterAcceptor{
				{
					State:   request.SuccessWaiterState,
					Matcher: request.PathWaiterMatch, Argument: "AuditReportStatus",
					Expected: "SUCCESS",
				},
				{
					State:   request.FailureWaiterState,
					Matcher: request.PathWaiterMatch, Argument: "AuditReportStatus",
					Expected: "FAILED",
				},
			},
		},
	},

	"CertificateAuthorityCSRCreated": &awsdouble.Waiter{
		OperationName: "GetCertificateAuthorityCsr",
		Waiter: request.Waiter{
			Name:        "WaitUntilCertificateAuthorityCSRCreated",
			MaxAttempts: 60,
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
					Expected: "RequestInProgressException",
				},
			},
		},
	},

	"CertificateIssued": &awsdouble.Waiter{
		OperationName: "GetCertificate",
		Waiter: request.Waiter{
			Name:        "WaitUntilCertificateIssued",
			MaxAttempts: 60,
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
					Expected: "RequestInProgressException",
				},
			},
		},
	},
}
