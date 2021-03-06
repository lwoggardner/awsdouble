// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:14:39+11:00

// Package resourcegroupstaggingapidouble provides a TestDouble implementation of resourcegroupstaggingapiiface.ResourceGroupsTaggingAPIAPI
package resourcegroupstaggingapidouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi/resourcegroupstaggingapiiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// ResourceGroupsTaggingAPIDouble is TestDouble for resourcegroupstaggingapiiface.ResourceGroupsTaggingAPIAPI
type ResourceGroupsTaggingAPIDouble struct {
	resourcegroupstaggingapiiface.ResourceGroupsTaggingAPIAPI
	*awsdouble.AWSTestDouble
}

// Constructor for ResourceGroupsTaggingAPIDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *ResourceGroupsTaggingAPIDouble {
	result := &ResourceGroupsTaggingAPIDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*resourcegroupstaggingapiiface.ResourceGroupsTaggingAPIAPI)(nil), configurators...)
	return result
}

func (d *ResourceGroupsTaggingAPIDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *ResourceGroupsTaggingAPIDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "DescribeReportCreationWithContext":
		return m.Fake(d.fakeDescribeReportCreationWithContext)

	case "GetComplianceSummaryPages":
		return m.Fake(d.fakeGetComplianceSummaryPages)

	case "GetComplianceSummaryPagesWithContext":
		return m.Fake(d.fakeGetComplianceSummaryPagesWithContext)

	case "GetComplianceSummaryWithContext":
		return m.Fake(d.fakeGetComplianceSummaryWithContext)

	case "GetResourcesPages":
		return m.Fake(d.fakeGetResourcesPages)

	case "GetResourcesPagesWithContext":
		return m.Fake(d.fakeGetResourcesPagesWithContext)

	case "GetResourcesWithContext":
		return m.Fake(d.fakeGetResourcesWithContext)

	case "GetTagKeysPages":
		return m.Fake(d.fakeGetTagKeysPages)

	case "GetTagKeysPagesWithContext":
		return m.Fake(d.fakeGetTagKeysPagesWithContext)

	case "GetTagKeysWithContext":
		return m.Fake(d.fakeGetTagKeysWithContext)

	case "GetTagValuesPages":
		return m.Fake(d.fakeGetTagValuesPages)

	case "GetTagValuesPagesWithContext":
		return m.Fake(d.fakeGetTagValuesPagesWithContext)

	case "GetTagValuesWithContext":
		return m.Fake(d.fakeGetTagValuesWithContext)

	case "StartReportCreationWithContext":
		return m.Fake(d.fakeStartReportCreationWithContext)

	case "TagResourcesWithContext":
		return m.Fake(d.fakeTagResourcesWithContext)

	case "UntagResourcesWithContext":
		return m.Fake(d.fakeUntagResourcesWithContext)

	default:
		return nil
	}
}

func (d *ResourceGroupsTaggingAPIDouble) DescribeReportCreation(i0 *resourcegroupstaggingapi.DescribeReportCreationInput) (r0 *resourcegroupstaggingapi.DescribeReportCreationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeReportCreation", i0)
	r0, _ = returns[0].(*resourcegroupstaggingapi.DescribeReportCreationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) DescribeReportCreationRequest(i0 *resourcegroupstaggingapi.DescribeReportCreationInput) (r0 *request.Request, r1 *resourcegroupstaggingapi.DescribeReportCreationOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeReportCreationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroupstaggingapi.DescribeReportCreationOutput)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) DescribeReportCreationWithContext(i0 context.Context, i1 *resourcegroupstaggingapi.DescribeReportCreationInput, i2 ...request.Option) (r0 *resourcegroupstaggingapi.DescribeReportCreationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeReportCreationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroupstaggingapi.DescribeReportCreationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) fakeDescribeReportCreationWithContext(ctx context.Context, in *resourcegroupstaggingapi.DescribeReportCreationInput, _ ...request.Option) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeReportCreation cancelled", ctx.Err())
	default:
		return d.DescribeReportCreation(in)
	}
}

func (d *ResourceGroupsTaggingAPIDouble) GetComplianceSummary(i0 *resourcegroupstaggingapi.GetComplianceSummaryInput) (r0 *resourcegroupstaggingapi.GetComplianceSummaryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetComplianceSummary", i0)
	r0, _ = returns[0].(*resourcegroupstaggingapi.GetComplianceSummaryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) GetComplianceSummaryPages(i0 *resourcegroupstaggingapi.GetComplianceSummaryInput, i1 func(*resourcegroupstaggingapi.GetComplianceSummaryOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetComplianceSummaryPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) fakeGetComplianceSummaryPages(in *resourcegroupstaggingapi.GetComplianceSummaryInput, pager func(*resourcegroupstaggingapi.GetComplianceSummaryOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("GetComplianceSummary", paginators, in, pager)
}

func (d *ResourceGroupsTaggingAPIDouble) GetComplianceSummaryPagesWithContext(i0 context.Context, i1 *resourcegroupstaggingapi.GetComplianceSummaryInput, i2 func(*resourcegroupstaggingapi.GetComplianceSummaryOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetComplianceSummaryPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) fakeGetComplianceSummaryPagesWithContext(ctx context.Context, in *resourcegroupstaggingapi.GetComplianceSummaryInput, pager func(*resourcegroupstaggingapi.GetComplianceSummaryOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("GetComplianceSummary", paginators, ctx, in, pager, options...)
}

func (d *ResourceGroupsTaggingAPIDouble) GetComplianceSummaryRequest(i0 *resourcegroupstaggingapi.GetComplianceSummaryInput) (r0 *request.Request, r1 *resourcegroupstaggingapi.GetComplianceSummaryOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetComplianceSummaryRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroupstaggingapi.GetComplianceSummaryOutput)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) GetComplianceSummaryWithContext(i0 context.Context, i1 *resourcegroupstaggingapi.GetComplianceSummaryInput, i2 ...request.Option) (r0 *resourcegroupstaggingapi.GetComplianceSummaryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetComplianceSummaryWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroupstaggingapi.GetComplianceSummaryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) fakeGetComplianceSummaryWithContext(ctx context.Context, in *resourcegroupstaggingapi.GetComplianceSummaryInput, _ ...request.Option) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetComplianceSummary cancelled", ctx.Err())
	default:
		return d.GetComplianceSummary(in)
	}
}

func (d *ResourceGroupsTaggingAPIDouble) GetResources(i0 *resourcegroupstaggingapi.GetResourcesInput) (r0 *resourcegroupstaggingapi.GetResourcesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetResources", i0)
	r0, _ = returns[0].(*resourcegroupstaggingapi.GetResourcesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) GetResourcesPages(i0 *resourcegroupstaggingapi.GetResourcesInput, i1 func(*resourcegroupstaggingapi.GetResourcesOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetResourcesPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) fakeGetResourcesPages(in *resourcegroupstaggingapi.GetResourcesInput, pager func(*resourcegroupstaggingapi.GetResourcesOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("GetResources", paginators, in, pager)
}

func (d *ResourceGroupsTaggingAPIDouble) GetResourcesPagesWithContext(i0 context.Context, i1 *resourcegroupstaggingapi.GetResourcesInput, i2 func(*resourcegroupstaggingapi.GetResourcesOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetResourcesPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) fakeGetResourcesPagesWithContext(ctx context.Context, in *resourcegroupstaggingapi.GetResourcesInput, pager func(*resourcegroupstaggingapi.GetResourcesOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("GetResources", paginators, ctx, in, pager, options...)
}

func (d *ResourceGroupsTaggingAPIDouble) GetResourcesRequest(i0 *resourcegroupstaggingapi.GetResourcesInput) (r0 *request.Request, r1 *resourcegroupstaggingapi.GetResourcesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetResourcesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroupstaggingapi.GetResourcesOutput)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) GetResourcesWithContext(i0 context.Context, i1 *resourcegroupstaggingapi.GetResourcesInput, i2 ...request.Option) (r0 *resourcegroupstaggingapi.GetResourcesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetResourcesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroupstaggingapi.GetResourcesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) fakeGetResourcesWithContext(ctx context.Context, in *resourcegroupstaggingapi.GetResourcesInput, _ ...request.Option) (*resourcegroupstaggingapi.GetResourcesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetResources cancelled", ctx.Err())
	default:
		return d.GetResources(in)
	}
}

func (d *ResourceGroupsTaggingAPIDouble) GetTagKeys(i0 *resourcegroupstaggingapi.GetTagKeysInput) (r0 *resourcegroupstaggingapi.GetTagKeysOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetTagKeys", i0)
	r0, _ = returns[0].(*resourcegroupstaggingapi.GetTagKeysOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) GetTagKeysPages(i0 *resourcegroupstaggingapi.GetTagKeysInput, i1 func(*resourcegroupstaggingapi.GetTagKeysOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetTagKeysPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) fakeGetTagKeysPages(in *resourcegroupstaggingapi.GetTagKeysInput, pager func(*resourcegroupstaggingapi.GetTagKeysOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("GetTagKeys", paginators, in, pager)
}

func (d *ResourceGroupsTaggingAPIDouble) GetTagKeysPagesWithContext(i0 context.Context, i1 *resourcegroupstaggingapi.GetTagKeysInput, i2 func(*resourcegroupstaggingapi.GetTagKeysOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetTagKeysPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) fakeGetTagKeysPagesWithContext(ctx context.Context, in *resourcegroupstaggingapi.GetTagKeysInput, pager func(*resourcegroupstaggingapi.GetTagKeysOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("GetTagKeys", paginators, ctx, in, pager, options...)
}

func (d *ResourceGroupsTaggingAPIDouble) GetTagKeysRequest(i0 *resourcegroupstaggingapi.GetTagKeysInput) (r0 *request.Request, r1 *resourcegroupstaggingapi.GetTagKeysOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetTagKeysRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroupstaggingapi.GetTagKeysOutput)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) GetTagKeysWithContext(i0 context.Context, i1 *resourcegroupstaggingapi.GetTagKeysInput, i2 ...request.Option) (r0 *resourcegroupstaggingapi.GetTagKeysOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetTagKeysWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroupstaggingapi.GetTagKeysOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) fakeGetTagKeysWithContext(ctx context.Context, in *resourcegroupstaggingapi.GetTagKeysInput, _ ...request.Option) (*resourcegroupstaggingapi.GetTagKeysOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetTagKeys cancelled", ctx.Err())
	default:
		return d.GetTagKeys(in)
	}
}

func (d *ResourceGroupsTaggingAPIDouble) GetTagValues(i0 *resourcegroupstaggingapi.GetTagValuesInput) (r0 *resourcegroupstaggingapi.GetTagValuesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetTagValues", i0)
	r0, _ = returns[0].(*resourcegroupstaggingapi.GetTagValuesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) GetTagValuesPages(i0 *resourcegroupstaggingapi.GetTagValuesInput, i1 func(*resourcegroupstaggingapi.GetTagValuesOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetTagValuesPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) fakeGetTagValuesPages(in *resourcegroupstaggingapi.GetTagValuesInput, pager func(*resourcegroupstaggingapi.GetTagValuesOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("GetTagValues", paginators, in, pager)
}

func (d *ResourceGroupsTaggingAPIDouble) GetTagValuesPagesWithContext(i0 context.Context, i1 *resourcegroupstaggingapi.GetTagValuesInput, i2 func(*resourcegroupstaggingapi.GetTagValuesOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetTagValuesPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) fakeGetTagValuesPagesWithContext(ctx context.Context, in *resourcegroupstaggingapi.GetTagValuesInput, pager func(*resourcegroupstaggingapi.GetTagValuesOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("GetTagValues", paginators, ctx, in, pager, options...)
}

func (d *ResourceGroupsTaggingAPIDouble) GetTagValuesRequest(i0 *resourcegroupstaggingapi.GetTagValuesInput) (r0 *request.Request, r1 *resourcegroupstaggingapi.GetTagValuesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetTagValuesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroupstaggingapi.GetTagValuesOutput)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) GetTagValuesWithContext(i0 context.Context, i1 *resourcegroupstaggingapi.GetTagValuesInput, i2 ...request.Option) (r0 *resourcegroupstaggingapi.GetTagValuesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetTagValuesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroupstaggingapi.GetTagValuesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) fakeGetTagValuesWithContext(ctx context.Context, in *resourcegroupstaggingapi.GetTagValuesInput, _ ...request.Option) (*resourcegroupstaggingapi.GetTagValuesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetTagValues cancelled", ctx.Err())
	default:
		return d.GetTagValues(in)
	}
}

func (d *ResourceGroupsTaggingAPIDouble) StartReportCreation(i0 *resourcegroupstaggingapi.StartReportCreationInput) (r0 *resourcegroupstaggingapi.StartReportCreationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartReportCreation", i0)
	r0, _ = returns[0].(*resourcegroupstaggingapi.StartReportCreationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) StartReportCreationRequest(i0 *resourcegroupstaggingapi.StartReportCreationInput) (r0 *request.Request, r1 *resourcegroupstaggingapi.StartReportCreationOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartReportCreationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroupstaggingapi.StartReportCreationOutput)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) StartReportCreationWithContext(i0 context.Context, i1 *resourcegroupstaggingapi.StartReportCreationInput, i2 ...request.Option) (r0 *resourcegroupstaggingapi.StartReportCreationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartReportCreationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroupstaggingapi.StartReportCreationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) fakeStartReportCreationWithContext(ctx context.Context, in *resourcegroupstaggingapi.StartReportCreationInput, _ ...request.Option) (*resourcegroupstaggingapi.StartReportCreationOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StartReportCreation cancelled", ctx.Err())
	default:
		return d.StartReportCreation(in)
	}
}

func (d *ResourceGroupsTaggingAPIDouble) TagResources(i0 *resourcegroupstaggingapi.TagResourcesInput) (r0 *resourcegroupstaggingapi.TagResourcesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResources", i0)
	r0, _ = returns[0].(*resourcegroupstaggingapi.TagResourcesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) TagResourcesRequest(i0 *resourcegroupstaggingapi.TagResourcesInput) (r0 *request.Request, r1 *resourcegroupstaggingapi.TagResourcesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResourcesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroupstaggingapi.TagResourcesOutput)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) TagResourcesWithContext(i0 context.Context, i1 *resourcegroupstaggingapi.TagResourcesInput, i2 ...request.Option) (r0 *resourcegroupstaggingapi.TagResourcesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResourcesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroupstaggingapi.TagResourcesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) fakeTagResourcesWithContext(ctx context.Context, in *resourcegroupstaggingapi.TagResourcesInput, _ ...request.Option) (*resourcegroupstaggingapi.TagResourcesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "TagResources cancelled", ctx.Err())
	default:
		return d.TagResources(in)
	}
}

func (d *ResourceGroupsTaggingAPIDouble) UntagResources(i0 *resourcegroupstaggingapi.UntagResourcesInput) (r0 *resourcegroupstaggingapi.UntagResourcesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResources", i0)
	r0, _ = returns[0].(*resourcegroupstaggingapi.UntagResourcesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) UntagResourcesRequest(i0 *resourcegroupstaggingapi.UntagResourcesInput) (r0 *request.Request, r1 *resourcegroupstaggingapi.UntagResourcesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResourcesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroupstaggingapi.UntagResourcesOutput)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) UntagResourcesWithContext(i0 context.Context, i1 *resourcegroupstaggingapi.UntagResourcesInput, i2 ...request.Option) (r0 *resourcegroupstaggingapi.UntagResourcesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResourcesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroupstaggingapi.UntagResourcesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsTaggingAPIDouble) fakeUntagResourcesWithContext(ctx context.Context, in *resourcegroupstaggingapi.UntagResourcesInput, _ ...request.Option) (*resourcegroupstaggingapi.UntagResourcesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UntagResources cancelled", ctx.Err())
	default:
		return d.UntagResources(in)
	}
}
