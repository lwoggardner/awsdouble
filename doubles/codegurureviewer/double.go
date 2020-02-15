// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:11:50+11:00

// Package codegurureviewerdouble provides a TestDouble implementation of codegururevieweriface.CodeGuruReviewerAPI
package codegurureviewerdouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codegurureviewer"
	"github.com/aws/aws-sdk-go/service/codegurureviewer/codegururevieweriface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// CodeGuruReviewerDouble is TestDouble for codegururevieweriface.CodeGuruReviewerAPI
type CodeGuruReviewerDouble struct {
	codegururevieweriface.CodeGuruReviewerAPI
	*awsdouble.AWSTestDouble
}

// Constructor for CodeGuruReviewerDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *CodeGuruReviewerDouble {
	result := &CodeGuruReviewerDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*codegururevieweriface.CodeGuruReviewerAPI)(nil), configurators...)
	return result
}

func (d *CodeGuruReviewerDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *CodeGuruReviewerDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "AssociateRepositoryWithContext":
		return m.Fake(d.fakeAssociateRepositoryWithContext)

	case "DescribeRepositoryAssociationWithContext":
		return m.Fake(d.fakeDescribeRepositoryAssociationWithContext)

	case "DisassociateRepositoryWithContext":
		return m.Fake(d.fakeDisassociateRepositoryWithContext)

	case "ListRepositoryAssociationsPages":
		return m.Fake(d.fakeListRepositoryAssociationsPages)

	case "ListRepositoryAssociationsPagesWithContext":
		return m.Fake(d.fakeListRepositoryAssociationsPagesWithContext)

	case "ListRepositoryAssociationsWithContext":
		return m.Fake(d.fakeListRepositoryAssociationsWithContext)

	default:
		return nil
	}
}

func (d *CodeGuruReviewerDouble) AssociateRepository(i0 *codegurureviewer.AssociateRepositoryInput) (r0 *codegurureviewer.AssociateRepositoryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AssociateRepository", i0)
	r0, _ = returns[0].(*codegurureviewer.AssociateRepositoryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CodeGuruReviewerDouble) AssociateRepositoryRequest(i0 *codegurureviewer.AssociateRepositoryInput) (r0 *request.Request, r1 *codegurureviewer.AssociateRepositoryOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AssociateRepositoryRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*codegurureviewer.AssociateRepositoryOutput)
	return
}

func (d *CodeGuruReviewerDouble) AssociateRepositoryWithContext(i0 context.Context, i1 *codegurureviewer.AssociateRepositoryInput, i2 ...request.Option) (r0 *codegurureviewer.AssociateRepositoryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AssociateRepositoryWithContext", i0, i1, i2)
	r0, _ = returns[0].(*codegurureviewer.AssociateRepositoryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CodeGuruReviewerDouble) fakeAssociateRepositoryWithContext(ctx context.Context, in *codegurureviewer.AssociateRepositoryInput, _ ...request.Option) (*codegurureviewer.AssociateRepositoryOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "AssociateRepository cancelled", ctx.Err())
	default:
		return d.AssociateRepository(in)
	}
}

func (d *CodeGuruReviewerDouble) DescribeRepositoryAssociation(i0 *codegurureviewer.DescribeRepositoryAssociationInput) (r0 *codegurureviewer.DescribeRepositoryAssociationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeRepositoryAssociation", i0)
	r0, _ = returns[0].(*codegurureviewer.DescribeRepositoryAssociationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CodeGuruReviewerDouble) DescribeRepositoryAssociationRequest(i0 *codegurureviewer.DescribeRepositoryAssociationInput) (r0 *request.Request, r1 *codegurureviewer.DescribeRepositoryAssociationOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeRepositoryAssociationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*codegurureviewer.DescribeRepositoryAssociationOutput)
	return
}

func (d *CodeGuruReviewerDouble) DescribeRepositoryAssociationWithContext(i0 context.Context, i1 *codegurureviewer.DescribeRepositoryAssociationInput, i2 ...request.Option) (r0 *codegurureviewer.DescribeRepositoryAssociationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeRepositoryAssociationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*codegurureviewer.DescribeRepositoryAssociationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CodeGuruReviewerDouble) fakeDescribeRepositoryAssociationWithContext(ctx context.Context, in *codegurureviewer.DescribeRepositoryAssociationInput, _ ...request.Option) (*codegurureviewer.DescribeRepositoryAssociationOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeRepositoryAssociation cancelled", ctx.Err())
	default:
		return d.DescribeRepositoryAssociation(in)
	}
}

func (d *CodeGuruReviewerDouble) DisassociateRepository(i0 *codegurureviewer.DisassociateRepositoryInput) (r0 *codegurureviewer.DisassociateRepositoryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DisassociateRepository", i0)
	r0, _ = returns[0].(*codegurureviewer.DisassociateRepositoryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CodeGuruReviewerDouble) DisassociateRepositoryRequest(i0 *codegurureviewer.DisassociateRepositoryInput) (r0 *request.Request, r1 *codegurureviewer.DisassociateRepositoryOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DisassociateRepositoryRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*codegurureviewer.DisassociateRepositoryOutput)
	return
}

func (d *CodeGuruReviewerDouble) DisassociateRepositoryWithContext(i0 context.Context, i1 *codegurureviewer.DisassociateRepositoryInput, i2 ...request.Option) (r0 *codegurureviewer.DisassociateRepositoryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DisassociateRepositoryWithContext", i0, i1, i2)
	r0, _ = returns[0].(*codegurureviewer.DisassociateRepositoryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CodeGuruReviewerDouble) fakeDisassociateRepositoryWithContext(ctx context.Context, in *codegurureviewer.DisassociateRepositoryInput, _ ...request.Option) (*codegurureviewer.DisassociateRepositoryOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DisassociateRepository cancelled", ctx.Err())
	default:
		return d.DisassociateRepository(in)
	}
}

func (d *CodeGuruReviewerDouble) ListRepositoryAssociations(i0 *codegurureviewer.ListRepositoryAssociationsInput) (r0 *codegurureviewer.ListRepositoryAssociationsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListRepositoryAssociations", i0)
	r0, _ = returns[0].(*codegurureviewer.ListRepositoryAssociationsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CodeGuruReviewerDouble) ListRepositoryAssociationsPages(i0 *codegurureviewer.ListRepositoryAssociationsInput, i1 func(*codegurureviewer.ListRepositoryAssociationsOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListRepositoryAssociationsPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *CodeGuruReviewerDouble) fakeListRepositoryAssociationsPages(in *codegurureviewer.ListRepositoryAssociationsInput, pager func(*codegurureviewer.ListRepositoryAssociationsOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListRepositoryAssociations", paginators, in, pager)
}

func (d *CodeGuruReviewerDouble) ListRepositoryAssociationsPagesWithContext(i0 context.Context, i1 *codegurureviewer.ListRepositoryAssociationsInput, i2 func(*codegurureviewer.ListRepositoryAssociationsOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListRepositoryAssociationsPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *CodeGuruReviewerDouble) fakeListRepositoryAssociationsPagesWithContext(ctx context.Context, in *codegurureviewer.ListRepositoryAssociationsInput, pager func(*codegurureviewer.ListRepositoryAssociationsOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListRepositoryAssociations", paginators, ctx, in, pager, options...)
}

func (d *CodeGuruReviewerDouble) ListRepositoryAssociationsRequest(i0 *codegurureviewer.ListRepositoryAssociationsInput) (r0 *request.Request, r1 *codegurureviewer.ListRepositoryAssociationsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListRepositoryAssociationsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*codegurureviewer.ListRepositoryAssociationsOutput)
	return
}

func (d *CodeGuruReviewerDouble) ListRepositoryAssociationsWithContext(i0 context.Context, i1 *codegurureviewer.ListRepositoryAssociationsInput, i2 ...request.Option) (r0 *codegurureviewer.ListRepositoryAssociationsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListRepositoryAssociationsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*codegurureviewer.ListRepositoryAssociationsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CodeGuruReviewerDouble) fakeListRepositoryAssociationsWithContext(ctx context.Context, in *codegurureviewer.ListRepositoryAssociationsInput, _ ...request.Option) (*codegurureviewer.ListRepositoryAssociationsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListRepositoryAssociations cancelled", ctx.Err())
	default:
		return d.ListRepositoryAssociations(in)
	}
}
