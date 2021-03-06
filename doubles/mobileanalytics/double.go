// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:14:05+11:00

// Package mobileanalyticsdouble provides a TestDouble implementation of mobileanalyticsiface.MobileAnalyticsAPI
package mobileanalyticsdouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/mobileanalytics"
	"github.com/aws/aws-sdk-go/service/mobileanalytics/mobileanalyticsiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// MobileAnalyticsDouble is TestDouble for mobileanalyticsiface.MobileAnalyticsAPI
type MobileAnalyticsDouble struct {
	mobileanalyticsiface.MobileAnalyticsAPI
	*awsdouble.AWSTestDouble
}

// Constructor for MobileAnalyticsDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *MobileAnalyticsDouble {
	result := &MobileAnalyticsDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*mobileanalyticsiface.MobileAnalyticsAPI)(nil), configurators...)
	return result
}

func (d *MobileAnalyticsDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *MobileAnalyticsDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "PutEventsWithContext":
		return m.Fake(d.fakePutEventsWithContext)

	default:
		return nil
	}
}

func (d *MobileAnalyticsDouble) PutEvents(i0 *mobileanalytics.PutEventsInput) (r0 *mobileanalytics.PutEventsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutEvents", i0)
	r0, _ = returns[0].(*mobileanalytics.PutEventsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MobileAnalyticsDouble) PutEventsRequest(i0 *mobileanalytics.PutEventsInput) (r0 *request.Request, r1 *mobileanalytics.PutEventsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutEventsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mobileanalytics.PutEventsOutput)
	return
}

func (d *MobileAnalyticsDouble) PutEventsWithContext(i0 context.Context, i1 *mobileanalytics.PutEventsInput, i2 ...request.Option) (r0 *mobileanalytics.PutEventsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutEventsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mobileanalytics.PutEventsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MobileAnalyticsDouble) fakePutEventsWithContext(ctx context.Context, in *mobileanalytics.PutEventsInput, _ ...request.Option) (*mobileanalytics.PutEventsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "PutEvents cancelled", ctx.Err())
	default:
		return d.PutEvents(in)
	}
}
