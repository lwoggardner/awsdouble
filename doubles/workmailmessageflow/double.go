// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:15:36+11:00

// Package workmailmessageflowdouble provides a TestDouble implementation of workmailmessageflowiface.WorkMailMessageFlowAPI
package workmailmessageflowdouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/workmailmessageflow"
	"github.com/aws/aws-sdk-go/service/workmailmessageflow/workmailmessageflowiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// WorkMailMessageFlowDouble is TestDouble for workmailmessageflowiface.WorkMailMessageFlowAPI
type WorkMailMessageFlowDouble struct {
	workmailmessageflowiface.WorkMailMessageFlowAPI
	*awsdouble.AWSTestDouble
}

// Constructor for WorkMailMessageFlowDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *WorkMailMessageFlowDouble {
	result := &WorkMailMessageFlowDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*workmailmessageflowiface.WorkMailMessageFlowAPI)(nil), configurators...)
	return result
}

func (d *WorkMailMessageFlowDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *WorkMailMessageFlowDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "GetRawMessageContentWithContext":
		return m.Fake(d.fakeGetRawMessageContentWithContext)

	default:
		return nil
	}
}

func (d *WorkMailMessageFlowDouble) GetRawMessageContent(i0 *workmailmessageflow.GetRawMessageContentInput) (r0 *workmailmessageflow.GetRawMessageContentOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetRawMessageContent", i0)
	r0, _ = returns[0].(*workmailmessageflow.GetRawMessageContentOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *WorkMailMessageFlowDouble) GetRawMessageContentRequest(i0 *workmailmessageflow.GetRawMessageContentInput) (r0 *request.Request, r1 *workmailmessageflow.GetRawMessageContentOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetRawMessageContentRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*workmailmessageflow.GetRawMessageContentOutput)
	return
}

func (d *WorkMailMessageFlowDouble) GetRawMessageContentWithContext(i0 context.Context, i1 *workmailmessageflow.GetRawMessageContentInput, i2 ...request.Option) (r0 *workmailmessageflow.GetRawMessageContentOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetRawMessageContentWithContext", i0, i1, i2)
	r0, _ = returns[0].(*workmailmessageflow.GetRawMessageContentOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *WorkMailMessageFlowDouble) fakeGetRawMessageContentWithContext(ctx context.Context, in *workmailmessageflow.GetRawMessageContentInput, _ ...request.Option) (*workmailmessageflow.GetRawMessageContentOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetRawMessageContent cancelled", ctx.Err())
	default:
		return d.GetRawMessageContent(in)
	}
}
