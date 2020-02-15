// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:13:50+11:00

// Package marketplacemeteringdouble provides a TestDouble implementation of marketplacemeteringiface.MarketplaceMeteringAPI
package marketplacemeteringdouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/marketplacemetering"
	"github.com/aws/aws-sdk-go/service/marketplacemetering/marketplacemeteringiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// MarketplaceMeteringDouble is TestDouble for marketplacemeteringiface.MarketplaceMeteringAPI
type MarketplaceMeteringDouble struct {
	marketplacemeteringiface.MarketplaceMeteringAPI
	*awsdouble.AWSTestDouble
}

// Constructor for MarketplaceMeteringDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *MarketplaceMeteringDouble {
	result := &MarketplaceMeteringDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*marketplacemeteringiface.MarketplaceMeteringAPI)(nil), configurators...)
	return result
}

func (d *MarketplaceMeteringDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *MarketplaceMeteringDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "BatchMeterUsageWithContext":
		return m.Fake(d.fakeBatchMeterUsageWithContext)

	case "MeterUsageWithContext":
		return m.Fake(d.fakeMeterUsageWithContext)

	case "RegisterUsageWithContext":
		return m.Fake(d.fakeRegisterUsageWithContext)

	case "ResolveCustomerWithContext":
		return m.Fake(d.fakeResolveCustomerWithContext)

	default:
		return nil
	}
}

func (d *MarketplaceMeteringDouble) BatchMeterUsage(i0 *marketplacemetering.BatchMeterUsageInput) (r0 *marketplacemetering.BatchMeterUsageOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchMeterUsage", i0)
	r0, _ = returns[0].(*marketplacemetering.BatchMeterUsageOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MarketplaceMeteringDouble) BatchMeterUsageRequest(i0 *marketplacemetering.BatchMeterUsageInput) (r0 *request.Request, r1 *marketplacemetering.BatchMeterUsageOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchMeterUsageRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*marketplacemetering.BatchMeterUsageOutput)
	return
}

func (d *MarketplaceMeteringDouble) BatchMeterUsageWithContext(i0 context.Context, i1 *marketplacemetering.BatchMeterUsageInput, i2 ...request.Option) (r0 *marketplacemetering.BatchMeterUsageOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchMeterUsageWithContext", i0, i1, i2)
	r0, _ = returns[0].(*marketplacemetering.BatchMeterUsageOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MarketplaceMeteringDouble) fakeBatchMeterUsageWithContext(ctx context.Context, in *marketplacemetering.BatchMeterUsageInput, _ ...request.Option) (*marketplacemetering.BatchMeterUsageOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "BatchMeterUsage cancelled", ctx.Err())
	default:
		return d.BatchMeterUsage(in)
	}
}

func (d *MarketplaceMeteringDouble) MeterUsage(i0 *marketplacemetering.MeterUsageInput) (r0 *marketplacemetering.MeterUsageOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("MeterUsage", i0)
	r0, _ = returns[0].(*marketplacemetering.MeterUsageOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MarketplaceMeteringDouble) MeterUsageRequest(i0 *marketplacemetering.MeterUsageInput) (r0 *request.Request, r1 *marketplacemetering.MeterUsageOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("MeterUsageRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*marketplacemetering.MeterUsageOutput)
	return
}

func (d *MarketplaceMeteringDouble) MeterUsageWithContext(i0 context.Context, i1 *marketplacemetering.MeterUsageInput, i2 ...request.Option) (r0 *marketplacemetering.MeterUsageOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("MeterUsageWithContext", i0, i1, i2)
	r0, _ = returns[0].(*marketplacemetering.MeterUsageOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MarketplaceMeteringDouble) fakeMeterUsageWithContext(ctx context.Context, in *marketplacemetering.MeterUsageInput, _ ...request.Option) (*marketplacemetering.MeterUsageOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "MeterUsage cancelled", ctx.Err())
	default:
		return d.MeterUsage(in)
	}
}

func (d *MarketplaceMeteringDouble) RegisterUsage(i0 *marketplacemetering.RegisterUsageInput) (r0 *marketplacemetering.RegisterUsageOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("RegisterUsage", i0)
	r0, _ = returns[0].(*marketplacemetering.RegisterUsageOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MarketplaceMeteringDouble) RegisterUsageRequest(i0 *marketplacemetering.RegisterUsageInput) (r0 *request.Request, r1 *marketplacemetering.RegisterUsageOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("RegisterUsageRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*marketplacemetering.RegisterUsageOutput)
	return
}

func (d *MarketplaceMeteringDouble) RegisterUsageWithContext(i0 context.Context, i1 *marketplacemetering.RegisterUsageInput, i2 ...request.Option) (r0 *marketplacemetering.RegisterUsageOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("RegisterUsageWithContext", i0, i1, i2)
	r0, _ = returns[0].(*marketplacemetering.RegisterUsageOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MarketplaceMeteringDouble) fakeRegisterUsageWithContext(ctx context.Context, in *marketplacemetering.RegisterUsageInput, _ ...request.Option) (*marketplacemetering.RegisterUsageOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "RegisterUsage cancelled", ctx.Err())
	default:
		return d.RegisterUsage(in)
	}
}

func (d *MarketplaceMeteringDouble) ResolveCustomer(i0 *marketplacemetering.ResolveCustomerInput) (r0 *marketplacemetering.ResolveCustomerOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ResolveCustomer", i0)
	r0, _ = returns[0].(*marketplacemetering.ResolveCustomerOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MarketplaceMeteringDouble) ResolveCustomerRequest(i0 *marketplacemetering.ResolveCustomerInput) (r0 *request.Request, r1 *marketplacemetering.ResolveCustomerOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ResolveCustomerRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*marketplacemetering.ResolveCustomerOutput)
	return
}

func (d *MarketplaceMeteringDouble) ResolveCustomerWithContext(i0 context.Context, i1 *marketplacemetering.ResolveCustomerInput, i2 ...request.Option) (r0 *marketplacemetering.ResolveCustomerOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ResolveCustomerWithContext", i0, i1, i2)
	r0, _ = returns[0].(*marketplacemetering.ResolveCustomerOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MarketplaceMeteringDouble) fakeResolveCustomerWithContext(ctx context.Context, in *marketplacemetering.ResolveCustomerInput, _ ...request.Option) (*marketplacemetering.ResolveCustomerOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ResolveCustomer cancelled", ctx.Err())
	default:
		return d.ResolveCustomer(in)
	}
}