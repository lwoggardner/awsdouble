// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:14:26+11:00

// Package pricingdouble provides a TestDouble implementation of pricingiface.PricingAPI
package pricingdouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/pricing"
	"github.com/aws/aws-sdk-go/service/pricing/pricingiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// PricingDouble is TestDouble for pricingiface.PricingAPI
type PricingDouble struct {
	pricingiface.PricingAPI
	*awsdouble.AWSTestDouble
}

// Constructor for PricingDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *PricingDouble {
	result := &PricingDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*pricingiface.PricingAPI)(nil), configurators...)
	return result
}

func (d *PricingDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *PricingDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "DescribeServicesPages":
		return m.Fake(d.fakeDescribeServicesPages)

	case "DescribeServicesPagesWithContext":
		return m.Fake(d.fakeDescribeServicesPagesWithContext)

	case "DescribeServicesWithContext":
		return m.Fake(d.fakeDescribeServicesWithContext)

	case "GetAttributeValuesPages":
		return m.Fake(d.fakeGetAttributeValuesPages)

	case "GetAttributeValuesPagesWithContext":
		return m.Fake(d.fakeGetAttributeValuesPagesWithContext)

	case "GetAttributeValuesWithContext":
		return m.Fake(d.fakeGetAttributeValuesWithContext)

	case "GetProductsPages":
		return m.Fake(d.fakeGetProductsPages)

	case "GetProductsPagesWithContext":
		return m.Fake(d.fakeGetProductsPagesWithContext)

	case "GetProductsWithContext":
		return m.Fake(d.fakeGetProductsWithContext)

	default:
		return nil
	}
}

func (d *PricingDouble) DescribeServices(i0 *pricing.DescribeServicesInput) (r0 *pricing.DescribeServicesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeServices", i0)
	r0, _ = returns[0].(*pricing.DescribeServicesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *PricingDouble) DescribeServicesPages(i0 *pricing.DescribeServicesInput, i1 func(*pricing.DescribeServicesOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeServicesPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *PricingDouble) fakeDescribeServicesPages(in *pricing.DescribeServicesInput, pager func(*pricing.DescribeServicesOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("DescribeServices", paginators, in, pager)
}

func (d *PricingDouble) DescribeServicesPagesWithContext(i0 context.Context, i1 *pricing.DescribeServicesInput, i2 func(*pricing.DescribeServicesOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeServicesPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *PricingDouble) fakeDescribeServicesPagesWithContext(ctx context.Context, in *pricing.DescribeServicesInput, pager func(*pricing.DescribeServicesOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("DescribeServices", paginators, ctx, in, pager, options...)
}

func (d *PricingDouble) DescribeServicesRequest(i0 *pricing.DescribeServicesInput) (r0 *request.Request, r1 *pricing.DescribeServicesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeServicesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*pricing.DescribeServicesOutput)
	return
}

func (d *PricingDouble) DescribeServicesWithContext(i0 context.Context, i1 *pricing.DescribeServicesInput, i2 ...request.Option) (r0 *pricing.DescribeServicesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeServicesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*pricing.DescribeServicesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *PricingDouble) fakeDescribeServicesWithContext(ctx context.Context, in *pricing.DescribeServicesInput, _ ...request.Option) (*pricing.DescribeServicesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeServices cancelled", ctx.Err())
	default:
		return d.DescribeServices(in)
	}
}

func (d *PricingDouble) GetAttributeValues(i0 *pricing.GetAttributeValuesInput) (r0 *pricing.GetAttributeValuesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetAttributeValues", i0)
	r0, _ = returns[0].(*pricing.GetAttributeValuesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *PricingDouble) GetAttributeValuesPages(i0 *pricing.GetAttributeValuesInput, i1 func(*pricing.GetAttributeValuesOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetAttributeValuesPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *PricingDouble) fakeGetAttributeValuesPages(in *pricing.GetAttributeValuesInput, pager func(*pricing.GetAttributeValuesOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("GetAttributeValues", paginators, in, pager)
}

func (d *PricingDouble) GetAttributeValuesPagesWithContext(i0 context.Context, i1 *pricing.GetAttributeValuesInput, i2 func(*pricing.GetAttributeValuesOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetAttributeValuesPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *PricingDouble) fakeGetAttributeValuesPagesWithContext(ctx context.Context, in *pricing.GetAttributeValuesInput, pager func(*pricing.GetAttributeValuesOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("GetAttributeValues", paginators, ctx, in, pager, options...)
}

func (d *PricingDouble) GetAttributeValuesRequest(i0 *pricing.GetAttributeValuesInput) (r0 *request.Request, r1 *pricing.GetAttributeValuesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetAttributeValuesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*pricing.GetAttributeValuesOutput)
	return
}

func (d *PricingDouble) GetAttributeValuesWithContext(i0 context.Context, i1 *pricing.GetAttributeValuesInput, i2 ...request.Option) (r0 *pricing.GetAttributeValuesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetAttributeValuesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*pricing.GetAttributeValuesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *PricingDouble) fakeGetAttributeValuesWithContext(ctx context.Context, in *pricing.GetAttributeValuesInput, _ ...request.Option) (*pricing.GetAttributeValuesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetAttributeValues cancelled", ctx.Err())
	default:
		return d.GetAttributeValues(in)
	}
}

func (d *PricingDouble) GetProducts(i0 *pricing.GetProductsInput) (r0 *pricing.GetProductsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetProducts", i0)
	r0, _ = returns[0].(*pricing.GetProductsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *PricingDouble) GetProductsPages(i0 *pricing.GetProductsInput, i1 func(*pricing.GetProductsOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetProductsPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *PricingDouble) fakeGetProductsPages(in *pricing.GetProductsInput, pager func(*pricing.GetProductsOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("GetProducts", paginators, in, pager)
}

func (d *PricingDouble) GetProductsPagesWithContext(i0 context.Context, i1 *pricing.GetProductsInput, i2 func(*pricing.GetProductsOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetProductsPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *PricingDouble) fakeGetProductsPagesWithContext(ctx context.Context, in *pricing.GetProductsInput, pager func(*pricing.GetProductsOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("GetProducts", paginators, ctx, in, pager, options...)
}

func (d *PricingDouble) GetProductsRequest(i0 *pricing.GetProductsInput) (r0 *request.Request, r1 *pricing.GetProductsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetProductsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*pricing.GetProductsOutput)
	return
}

func (d *PricingDouble) GetProductsWithContext(i0 context.Context, i1 *pricing.GetProductsInput, i2 ...request.Option) (r0 *pricing.GetProductsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetProductsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*pricing.GetProductsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *PricingDouble) fakeGetProductsWithContext(ctx context.Context, in *pricing.GetProductsInput, _ ...request.Option) (*pricing.GetProductsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetProducts cancelled", ctx.Err())
	default:
		return d.GetProducts(in)
	}
}
