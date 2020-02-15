// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:13:15+11:00

// Package iotdataplanedouble provides a TestDouble implementation of iotdataplaneiface.IoTDataPlaneAPI
package iotdataplanedouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"github.com/aws/aws-sdk-go/service/iotdataplane/iotdataplaneiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// IoTDataPlaneDouble is TestDouble for iotdataplaneiface.IoTDataPlaneAPI
type IoTDataPlaneDouble struct {
	iotdataplaneiface.IoTDataPlaneAPI
	*awsdouble.AWSTestDouble
}

// Constructor for IoTDataPlaneDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *IoTDataPlaneDouble {
	result := &IoTDataPlaneDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*iotdataplaneiface.IoTDataPlaneAPI)(nil), configurators...)
	return result
}

func (d *IoTDataPlaneDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *IoTDataPlaneDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "DeleteThingShadowWithContext":
		return m.Fake(d.fakeDeleteThingShadowWithContext)

	case "GetThingShadowWithContext":
		return m.Fake(d.fakeGetThingShadowWithContext)

	case "PublishWithContext":
		return m.Fake(d.fakePublishWithContext)

	case "UpdateThingShadowWithContext":
		return m.Fake(d.fakeUpdateThingShadowWithContext)

	default:
		return nil
	}
}

func (d *IoTDataPlaneDouble) DeleteThingShadow(i0 *iotdataplane.DeleteThingShadowInput) (r0 *iotdataplane.DeleteThingShadowOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteThingShadow", i0)
	r0, _ = returns[0].(*iotdataplane.DeleteThingShadowOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *IoTDataPlaneDouble) DeleteThingShadowRequest(i0 *iotdataplane.DeleteThingShadowInput) (r0 *request.Request, r1 *iotdataplane.DeleteThingShadowOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteThingShadowRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*iotdataplane.DeleteThingShadowOutput)
	return
}

func (d *IoTDataPlaneDouble) DeleteThingShadowWithContext(i0 context.Context, i1 *iotdataplane.DeleteThingShadowInput, i2 ...request.Option) (r0 *iotdataplane.DeleteThingShadowOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteThingShadowWithContext", i0, i1, i2)
	r0, _ = returns[0].(*iotdataplane.DeleteThingShadowOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *IoTDataPlaneDouble) fakeDeleteThingShadowWithContext(ctx context.Context, in *iotdataplane.DeleteThingShadowInput, _ ...request.Option) (*iotdataplane.DeleteThingShadowOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteThingShadow cancelled", ctx.Err())
	default:
		return d.DeleteThingShadow(in)
	}
}

func (d *IoTDataPlaneDouble) GetThingShadow(i0 *iotdataplane.GetThingShadowInput) (r0 *iotdataplane.GetThingShadowOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetThingShadow", i0)
	r0, _ = returns[0].(*iotdataplane.GetThingShadowOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *IoTDataPlaneDouble) GetThingShadowRequest(i0 *iotdataplane.GetThingShadowInput) (r0 *request.Request, r1 *iotdataplane.GetThingShadowOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetThingShadowRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*iotdataplane.GetThingShadowOutput)
	return
}

func (d *IoTDataPlaneDouble) GetThingShadowWithContext(i0 context.Context, i1 *iotdataplane.GetThingShadowInput, i2 ...request.Option) (r0 *iotdataplane.GetThingShadowOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetThingShadowWithContext", i0, i1, i2)
	r0, _ = returns[0].(*iotdataplane.GetThingShadowOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *IoTDataPlaneDouble) fakeGetThingShadowWithContext(ctx context.Context, in *iotdataplane.GetThingShadowInput, _ ...request.Option) (*iotdataplane.GetThingShadowOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetThingShadow cancelled", ctx.Err())
	default:
		return d.GetThingShadow(in)
	}
}

func (d *IoTDataPlaneDouble) Publish(i0 *iotdataplane.PublishInput) (r0 *iotdataplane.PublishOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("Publish", i0)
	r0, _ = returns[0].(*iotdataplane.PublishOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *IoTDataPlaneDouble) PublishRequest(i0 *iotdataplane.PublishInput) (r0 *request.Request, r1 *iotdataplane.PublishOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PublishRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*iotdataplane.PublishOutput)
	return
}

func (d *IoTDataPlaneDouble) PublishWithContext(i0 context.Context, i1 *iotdataplane.PublishInput, i2 ...request.Option) (r0 *iotdataplane.PublishOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PublishWithContext", i0, i1, i2)
	r0, _ = returns[0].(*iotdataplane.PublishOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *IoTDataPlaneDouble) fakePublishWithContext(ctx context.Context, in *iotdataplane.PublishInput, _ ...request.Option) (*iotdataplane.PublishOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "Publish cancelled", ctx.Err())
	default:
		return d.Publish(in)
	}
}

func (d *IoTDataPlaneDouble) UpdateThingShadow(i0 *iotdataplane.UpdateThingShadowInput) (r0 *iotdataplane.UpdateThingShadowOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateThingShadow", i0)
	r0, _ = returns[0].(*iotdataplane.UpdateThingShadowOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *IoTDataPlaneDouble) UpdateThingShadowRequest(i0 *iotdataplane.UpdateThingShadowInput) (r0 *request.Request, r1 *iotdataplane.UpdateThingShadowOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateThingShadowRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*iotdataplane.UpdateThingShadowOutput)
	return
}

func (d *IoTDataPlaneDouble) UpdateThingShadowWithContext(i0 context.Context, i1 *iotdataplane.UpdateThingShadowInput, i2 ...request.Option) (r0 *iotdataplane.UpdateThingShadowOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateThingShadowWithContext", i0, i1, i2)
	r0, _ = returns[0].(*iotdataplane.UpdateThingShadowOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *IoTDataPlaneDouble) fakeUpdateThingShadowWithContext(ctx context.Context, in *iotdataplane.UpdateThingShadowInput, _ ...request.Option) (*iotdataplane.UpdateThingShadowOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateThingShadow cancelled", ctx.Err())
	default:
		return d.UpdateThingShadow(in)
	}
}
