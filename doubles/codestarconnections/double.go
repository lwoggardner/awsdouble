// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:11:54+11:00

// Package codestarconnectionsdouble provides a TestDouble implementation of codestarconnectionsiface.CodeStarConnectionsAPI
package codestarconnectionsdouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codestarconnections"
	"github.com/aws/aws-sdk-go/service/codestarconnections/codestarconnectionsiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// CodeStarConnectionsDouble is TestDouble for codestarconnectionsiface.CodeStarConnectionsAPI
type CodeStarConnectionsDouble struct {
	codestarconnectionsiface.CodeStarConnectionsAPI
	*awsdouble.AWSTestDouble
}

// Constructor for CodeStarConnectionsDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *CodeStarConnectionsDouble {
	result := &CodeStarConnectionsDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*codestarconnectionsiface.CodeStarConnectionsAPI)(nil), configurators...)
	return result
}

func (d *CodeStarConnectionsDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *CodeStarConnectionsDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "CreateConnectionWithContext":
		return m.Fake(d.fakeCreateConnectionWithContext)

	case "DeleteConnectionWithContext":
		return m.Fake(d.fakeDeleteConnectionWithContext)

	case "GetConnectionWithContext":
		return m.Fake(d.fakeGetConnectionWithContext)

	case "ListConnectionsPages":
		return m.Fake(d.fakeListConnectionsPages)

	case "ListConnectionsPagesWithContext":
		return m.Fake(d.fakeListConnectionsPagesWithContext)

	case "ListConnectionsWithContext":
		return m.Fake(d.fakeListConnectionsWithContext)

	default:
		return nil
	}
}

func (d *CodeStarConnectionsDouble) CreateConnection(i0 *codestarconnections.CreateConnectionInput) (r0 *codestarconnections.CreateConnectionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateConnection", i0)
	r0, _ = returns[0].(*codestarconnections.CreateConnectionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CodeStarConnectionsDouble) CreateConnectionRequest(i0 *codestarconnections.CreateConnectionInput) (r0 *request.Request, r1 *codestarconnections.CreateConnectionOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateConnectionRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*codestarconnections.CreateConnectionOutput)
	return
}

func (d *CodeStarConnectionsDouble) CreateConnectionWithContext(i0 context.Context, i1 *codestarconnections.CreateConnectionInput, i2 ...request.Option) (r0 *codestarconnections.CreateConnectionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateConnectionWithContext", i0, i1, i2)
	r0, _ = returns[0].(*codestarconnections.CreateConnectionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CodeStarConnectionsDouble) fakeCreateConnectionWithContext(ctx context.Context, in *codestarconnections.CreateConnectionInput, _ ...request.Option) (*codestarconnections.CreateConnectionOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateConnection cancelled", ctx.Err())
	default:
		return d.CreateConnection(in)
	}
}

func (d *CodeStarConnectionsDouble) DeleteConnection(i0 *codestarconnections.DeleteConnectionInput) (r0 *codestarconnections.DeleteConnectionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteConnection", i0)
	r0, _ = returns[0].(*codestarconnections.DeleteConnectionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CodeStarConnectionsDouble) DeleteConnectionRequest(i0 *codestarconnections.DeleteConnectionInput) (r0 *request.Request, r1 *codestarconnections.DeleteConnectionOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteConnectionRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*codestarconnections.DeleteConnectionOutput)
	return
}

func (d *CodeStarConnectionsDouble) DeleteConnectionWithContext(i0 context.Context, i1 *codestarconnections.DeleteConnectionInput, i2 ...request.Option) (r0 *codestarconnections.DeleteConnectionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteConnectionWithContext", i0, i1, i2)
	r0, _ = returns[0].(*codestarconnections.DeleteConnectionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CodeStarConnectionsDouble) fakeDeleteConnectionWithContext(ctx context.Context, in *codestarconnections.DeleteConnectionInput, _ ...request.Option) (*codestarconnections.DeleteConnectionOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteConnection cancelled", ctx.Err())
	default:
		return d.DeleteConnection(in)
	}
}

func (d *CodeStarConnectionsDouble) GetConnection(i0 *codestarconnections.GetConnectionInput) (r0 *codestarconnections.GetConnectionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetConnection", i0)
	r0, _ = returns[0].(*codestarconnections.GetConnectionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CodeStarConnectionsDouble) GetConnectionRequest(i0 *codestarconnections.GetConnectionInput) (r0 *request.Request, r1 *codestarconnections.GetConnectionOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetConnectionRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*codestarconnections.GetConnectionOutput)
	return
}

func (d *CodeStarConnectionsDouble) GetConnectionWithContext(i0 context.Context, i1 *codestarconnections.GetConnectionInput, i2 ...request.Option) (r0 *codestarconnections.GetConnectionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetConnectionWithContext", i0, i1, i2)
	r0, _ = returns[0].(*codestarconnections.GetConnectionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CodeStarConnectionsDouble) fakeGetConnectionWithContext(ctx context.Context, in *codestarconnections.GetConnectionInput, _ ...request.Option) (*codestarconnections.GetConnectionOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetConnection cancelled", ctx.Err())
	default:
		return d.GetConnection(in)
	}
}

func (d *CodeStarConnectionsDouble) ListConnections(i0 *codestarconnections.ListConnectionsInput) (r0 *codestarconnections.ListConnectionsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListConnections", i0)
	r0, _ = returns[0].(*codestarconnections.ListConnectionsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CodeStarConnectionsDouble) ListConnectionsPages(i0 *codestarconnections.ListConnectionsInput, i1 func(*codestarconnections.ListConnectionsOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListConnectionsPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *CodeStarConnectionsDouble) fakeListConnectionsPages(in *codestarconnections.ListConnectionsInput, pager func(*codestarconnections.ListConnectionsOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListConnections", paginators, in, pager)
}

func (d *CodeStarConnectionsDouble) ListConnectionsPagesWithContext(i0 context.Context, i1 *codestarconnections.ListConnectionsInput, i2 func(*codestarconnections.ListConnectionsOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListConnectionsPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *CodeStarConnectionsDouble) fakeListConnectionsPagesWithContext(ctx context.Context, in *codestarconnections.ListConnectionsInput, pager func(*codestarconnections.ListConnectionsOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListConnections", paginators, ctx, in, pager, options...)
}

func (d *CodeStarConnectionsDouble) ListConnectionsRequest(i0 *codestarconnections.ListConnectionsInput) (r0 *request.Request, r1 *codestarconnections.ListConnectionsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListConnectionsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*codestarconnections.ListConnectionsOutput)
	return
}

func (d *CodeStarConnectionsDouble) ListConnectionsWithContext(i0 context.Context, i1 *codestarconnections.ListConnectionsInput, i2 ...request.Option) (r0 *codestarconnections.ListConnectionsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListConnectionsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*codestarconnections.ListConnectionsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CodeStarConnectionsDouble) fakeListConnectionsWithContext(ctx context.Context, in *codestarconnections.ListConnectionsInput, _ ...request.Option) (*codestarconnections.ListConnectionsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListConnections cancelled", ctx.Err())
	default:
		return d.ListConnections(in)
	}
}
