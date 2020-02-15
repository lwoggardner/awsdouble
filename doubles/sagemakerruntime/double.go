// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:14:50+11:00

// Package sagemakerruntimedouble provides a TestDouble implementation of sagemakerruntimeiface.SageMakerRuntimeAPI
package sagemakerruntimedouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sagemakerruntime"
	"github.com/aws/aws-sdk-go/service/sagemakerruntime/sagemakerruntimeiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// SageMakerRuntimeDouble is TestDouble for sagemakerruntimeiface.SageMakerRuntimeAPI
type SageMakerRuntimeDouble struct {
	sagemakerruntimeiface.SageMakerRuntimeAPI
	*awsdouble.AWSTestDouble
}

// Constructor for SageMakerRuntimeDouble
// Default configuration will ensure
//   * API operations not explicitly stubbed will return an empty output struct pointer, and nil error.
//     To simulate long polling, "Poll" operations will return these values after a random delay of up to 100ms.
//   * WithContext methods implement a 'Fake' method that returns a cancellation error if the context is cancelled
//      before the method is called.
//   * Pages and PagesWithContext methods similarly implement a 'Fake' method that paginates over the underlying method.
//
//    This allows tests to only stub the simple api methods and be generally unconcerned whether the SUT is using the
//      Context or Pagination forms of the API.
//
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *SageMakerRuntimeDouble {
	result := &SageMakerRuntimeDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*sagemakerruntimeiface.SageMakerRuntimeAPI)(nil), configurators...)
	return result
}

func (d *SageMakerRuntimeDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *SageMakerRuntimeDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "InvokeEndpointWithContext":
		return m.Fake(d.fakeInvokeEndpointWithContext)

	default:
		return nil
	}
}

func (d *SageMakerRuntimeDouble) InvokeEndpoint(i0 *sagemakerruntime.InvokeEndpointInput) (r0 *sagemakerruntime.InvokeEndpointOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("InvokeEndpoint", i0)
	r0, _ = returns[0].(*sagemakerruntime.InvokeEndpointOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SageMakerRuntimeDouble) InvokeEndpointRequest(i0 *sagemakerruntime.InvokeEndpointInput) (r0 *request.Request, r1 *sagemakerruntime.InvokeEndpointOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("InvokeEndpointRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*sagemakerruntime.InvokeEndpointOutput)
	return
}

func (d *SageMakerRuntimeDouble) InvokeEndpointWithContext(i0 context.Context, i1 *sagemakerruntime.InvokeEndpointInput, i2 ...request.Option) (r0 *sagemakerruntime.InvokeEndpointOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("InvokeEndpointWithContext", i0, i1, i2)
	r0, _ = returns[0].(*sagemakerruntime.InvokeEndpointOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SageMakerRuntimeDouble) fakeInvokeEndpointWithContext(ctx context.Context, in *sagemakerruntime.InvokeEndpointInput, _ ...request.Option) (*sagemakerruntime.InvokeEndpointOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "InvokeEndpoint cancelled", ctx.Err())
	default:
		return d.InvokeEndpoint(in)
	}
}
