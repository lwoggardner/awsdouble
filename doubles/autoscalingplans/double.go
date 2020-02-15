// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:11:23+11:00

// Package autoscalingplansdouble provides a TestDouble implementation of autoscalingplansiface.AutoScalingPlansAPI
package autoscalingplansdouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/autoscalingplans"
	"github.com/aws/aws-sdk-go/service/autoscalingplans/autoscalingplansiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// AutoScalingPlansDouble is TestDouble for autoscalingplansiface.AutoScalingPlansAPI
type AutoScalingPlansDouble struct {
	autoscalingplansiface.AutoScalingPlansAPI
	*awsdouble.AWSTestDouble
}

// Constructor for AutoScalingPlansDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *AutoScalingPlansDouble {
	result := &AutoScalingPlansDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*autoscalingplansiface.AutoScalingPlansAPI)(nil), configurators...)
	return result
}

func (d *AutoScalingPlansDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *AutoScalingPlansDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "CreateScalingPlanWithContext":
		return m.Fake(d.fakeCreateScalingPlanWithContext)

	case "DeleteScalingPlanWithContext":
		return m.Fake(d.fakeDeleteScalingPlanWithContext)

	case "DescribeScalingPlanResourcesWithContext":
		return m.Fake(d.fakeDescribeScalingPlanResourcesWithContext)

	case "DescribeScalingPlansWithContext":
		return m.Fake(d.fakeDescribeScalingPlansWithContext)

	case "GetScalingPlanResourceForecastDataWithContext":
		return m.Fake(d.fakeGetScalingPlanResourceForecastDataWithContext)

	case "UpdateScalingPlanWithContext":
		return m.Fake(d.fakeUpdateScalingPlanWithContext)

	default:
		return nil
	}
}

func (d *AutoScalingPlansDouble) CreateScalingPlan(i0 *autoscalingplans.CreateScalingPlanInput) (r0 *autoscalingplans.CreateScalingPlanOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateScalingPlan", i0)
	r0, _ = returns[0].(*autoscalingplans.CreateScalingPlanOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AutoScalingPlansDouble) CreateScalingPlanRequest(i0 *autoscalingplans.CreateScalingPlanInput) (r0 *request.Request, r1 *autoscalingplans.CreateScalingPlanOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateScalingPlanRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*autoscalingplans.CreateScalingPlanOutput)
	return
}

func (d *AutoScalingPlansDouble) CreateScalingPlanWithContext(i0 context.Context, i1 *autoscalingplans.CreateScalingPlanInput, i2 ...request.Option) (r0 *autoscalingplans.CreateScalingPlanOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateScalingPlanWithContext", i0, i1, i2)
	r0, _ = returns[0].(*autoscalingplans.CreateScalingPlanOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AutoScalingPlansDouble) fakeCreateScalingPlanWithContext(ctx context.Context, in *autoscalingplans.CreateScalingPlanInput, _ ...request.Option) (*autoscalingplans.CreateScalingPlanOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateScalingPlan cancelled", ctx.Err())
	default:
		return d.CreateScalingPlan(in)
	}
}

func (d *AutoScalingPlansDouble) DeleteScalingPlan(i0 *autoscalingplans.DeleteScalingPlanInput) (r0 *autoscalingplans.DeleteScalingPlanOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteScalingPlan", i0)
	r0, _ = returns[0].(*autoscalingplans.DeleteScalingPlanOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AutoScalingPlansDouble) DeleteScalingPlanRequest(i0 *autoscalingplans.DeleteScalingPlanInput) (r0 *request.Request, r1 *autoscalingplans.DeleteScalingPlanOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteScalingPlanRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*autoscalingplans.DeleteScalingPlanOutput)
	return
}

func (d *AutoScalingPlansDouble) DeleteScalingPlanWithContext(i0 context.Context, i1 *autoscalingplans.DeleteScalingPlanInput, i2 ...request.Option) (r0 *autoscalingplans.DeleteScalingPlanOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteScalingPlanWithContext", i0, i1, i2)
	r0, _ = returns[0].(*autoscalingplans.DeleteScalingPlanOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AutoScalingPlansDouble) fakeDeleteScalingPlanWithContext(ctx context.Context, in *autoscalingplans.DeleteScalingPlanInput, _ ...request.Option) (*autoscalingplans.DeleteScalingPlanOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteScalingPlan cancelled", ctx.Err())
	default:
		return d.DeleteScalingPlan(in)
	}
}

func (d *AutoScalingPlansDouble) DescribeScalingPlanResources(i0 *autoscalingplans.DescribeScalingPlanResourcesInput) (r0 *autoscalingplans.DescribeScalingPlanResourcesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeScalingPlanResources", i0)
	r0, _ = returns[0].(*autoscalingplans.DescribeScalingPlanResourcesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AutoScalingPlansDouble) DescribeScalingPlanResourcesRequest(i0 *autoscalingplans.DescribeScalingPlanResourcesInput) (r0 *request.Request, r1 *autoscalingplans.DescribeScalingPlanResourcesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeScalingPlanResourcesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*autoscalingplans.DescribeScalingPlanResourcesOutput)
	return
}

func (d *AutoScalingPlansDouble) DescribeScalingPlanResourcesWithContext(i0 context.Context, i1 *autoscalingplans.DescribeScalingPlanResourcesInput, i2 ...request.Option) (r0 *autoscalingplans.DescribeScalingPlanResourcesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeScalingPlanResourcesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*autoscalingplans.DescribeScalingPlanResourcesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AutoScalingPlansDouble) fakeDescribeScalingPlanResourcesWithContext(ctx context.Context, in *autoscalingplans.DescribeScalingPlanResourcesInput, _ ...request.Option) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeScalingPlanResources cancelled", ctx.Err())
	default:
		return d.DescribeScalingPlanResources(in)
	}
}

func (d *AutoScalingPlansDouble) DescribeScalingPlans(i0 *autoscalingplans.DescribeScalingPlansInput) (r0 *autoscalingplans.DescribeScalingPlansOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeScalingPlans", i0)
	r0, _ = returns[0].(*autoscalingplans.DescribeScalingPlansOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AutoScalingPlansDouble) DescribeScalingPlansRequest(i0 *autoscalingplans.DescribeScalingPlansInput) (r0 *request.Request, r1 *autoscalingplans.DescribeScalingPlansOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeScalingPlansRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*autoscalingplans.DescribeScalingPlansOutput)
	return
}

func (d *AutoScalingPlansDouble) DescribeScalingPlansWithContext(i0 context.Context, i1 *autoscalingplans.DescribeScalingPlansInput, i2 ...request.Option) (r0 *autoscalingplans.DescribeScalingPlansOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeScalingPlansWithContext", i0, i1, i2)
	r0, _ = returns[0].(*autoscalingplans.DescribeScalingPlansOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AutoScalingPlansDouble) fakeDescribeScalingPlansWithContext(ctx context.Context, in *autoscalingplans.DescribeScalingPlansInput, _ ...request.Option) (*autoscalingplans.DescribeScalingPlansOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeScalingPlans cancelled", ctx.Err())
	default:
		return d.DescribeScalingPlans(in)
	}
}

func (d *AutoScalingPlansDouble) GetScalingPlanResourceForecastData(i0 *autoscalingplans.GetScalingPlanResourceForecastDataInput) (r0 *autoscalingplans.GetScalingPlanResourceForecastDataOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetScalingPlanResourceForecastData", i0)
	r0, _ = returns[0].(*autoscalingplans.GetScalingPlanResourceForecastDataOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AutoScalingPlansDouble) GetScalingPlanResourceForecastDataRequest(i0 *autoscalingplans.GetScalingPlanResourceForecastDataInput) (r0 *request.Request, r1 *autoscalingplans.GetScalingPlanResourceForecastDataOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetScalingPlanResourceForecastDataRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*autoscalingplans.GetScalingPlanResourceForecastDataOutput)
	return
}

func (d *AutoScalingPlansDouble) GetScalingPlanResourceForecastDataWithContext(i0 context.Context, i1 *autoscalingplans.GetScalingPlanResourceForecastDataInput, i2 ...request.Option) (r0 *autoscalingplans.GetScalingPlanResourceForecastDataOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetScalingPlanResourceForecastDataWithContext", i0, i1, i2)
	r0, _ = returns[0].(*autoscalingplans.GetScalingPlanResourceForecastDataOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AutoScalingPlansDouble) fakeGetScalingPlanResourceForecastDataWithContext(ctx context.Context, in *autoscalingplans.GetScalingPlanResourceForecastDataInput, _ ...request.Option) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetScalingPlanResourceForecastData cancelled", ctx.Err())
	default:
		return d.GetScalingPlanResourceForecastData(in)
	}
}

func (d *AutoScalingPlansDouble) UpdateScalingPlan(i0 *autoscalingplans.UpdateScalingPlanInput) (r0 *autoscalingplans.UpdateScalingPlanOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateScalingPlan", i0)
	r0, _ = returns[0].(*autoscalingplans.UpdateScalingPlanOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AutoScalingPlansDouble) UpdateScalingPlanRequest(i0 *autoscalingplans.UpdateScalingPlanInput) (r0 *request.Request, r1 *autoscalingplans.UpdateScalingPlanOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateScalingPlanRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*autoscalingplans.UpdateScalingPlanOutput)
	return
}

func (d *AutoScalingPlansDouble) UpdateScalingPlanWithContext(i0 context.Context, i1 *autoscalingplans.UpdateScalingPlanInput, i2 ...request.Option) (r0 *autoscalingplans.UpdateScalingPlanOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateScalingPlanWithContext", i0, i1, i2)
	r0, _ = returns[0].(*autoscalingplans.UpdateScalingPlanOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AutoScalingPlansDouble) fakeUpdateScalingPlanWithContext(ctx context.Context, in *autoscalingplans.UpdateScalingPlanInput, _ ...request.Option) (*autoscalingplans.UpdateScalingPlanOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateScalingPlan cancelled", ctx.Err())
	default:
		return d.UpdateScalingPlan(in)
	}
}
