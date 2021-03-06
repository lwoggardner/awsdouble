// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:14:38+11:00

// Package resourcegroupsdouble provides a TestDouble implementation of resourcegroupsiface.ResourceGroupsAPI
package resourcegroupsdouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/resourcegroups"
	"github.com/aws/aws-sdk-go/service/resourcegroups/resourcegroupsiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// ResourceGroupsDouble is TestDouble for resourcegroupsiface.ResourceGroupsAPI
type ResourceGroupsDouble struct {
	resourcegroupsiface.ResourceGroupsAPI
	*awsdouble.AWSTestDouble
}

// Constructor for ResourceGroupsDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *ResourceGroupsDouble {
	result := &ResourceGroupsDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*resourcegroupsiface.ResourceGroupsAPI)(nil), configurators...)
	return result
}

func (d *ResourceGroupsDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *ResourceGroupsDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "CreateGroupWithContext":
		return m.Fake(d.fakeCreateGroupWithContext)

	case "DeleteGroupWithContext":
		return m.Fake(d.fakeDeleteGroupWithContext)

	case "GetGroupQueryWithContext":
		return m.Fake(d.fakeGetGroupQueryWithContext)

	case "GetGroupWithContext":
		return m.Fake(d.fakeGetGroupWithContext)

	case "GetTagsWithContext":
		return m.Fake(d.fakeGetTagsWithContext)

	case "ListGroupResourcesPages":
		return m.Fake(d.fakeListGroupResourcesPages)

	case "ListGroupResourcesPagesWithContext":
		return m.Fake(d.fakeListGroupResourcesPagesWithContext)

	case "ListGroupResourcesWithContext":
		return m.Fake(d.fakeListGroupResourcesWithContext)

	case "ListGroupsPages":
		return m.Fake(d.fakeListGroupsPages)

	case "ListGroupsPagesWithContext":
		return m.Fake(d.fakeListGroupsPagesWithContext)

	case "ListGroupsWithContext":
		return m.Fake(d.fakeListGroupsWithContext)

	case "SearchResourcesPages":
		return m.Fake(d.fakeSearchResourcesPages)

	case "SearchResourcesPagesWithContext":
		return m.Fake(d.fakeSearchResourcesPagesWithContext)

	case "SearchResourcesWithContext":
		return m.Fake(d.fakeSearchResourcesWithContext)

	case "TagWithContext":
		return m.Fake(d.fakeTagWithContext)

	case "UntagWithContext":
		return m.Fake(d.fakeUntagWithContext)

	case "UpdateGroupQueryWithContext":
		return m.Fake(d.fakeUpdateGroupQueryWithContext)

	case "UpdateGroupWithContext":
		return m.Fake(d.fakeUpdateGroupWithContext)

	default:
		return nil
	}
}

func (d *ResourceGroupsDouble) CreateGroup(i0 *resourcegroups.CreateGroupInput) (r0 *resourcegroups.CreateGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateGroup", i0)
	r0, _ = returns[0].(*resourcegroups.CreateGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) CreateGroupRequest(i0 *resourcegroups.CreateGroupInput) (r0 *request.Request, r1 *resourcegroups.CreateGroupOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateGroupRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroups.CreateGroupOutput)
	return
}

func (d *ResourceGroupsDouble) CreateGroupWithContext(i0 context.Context, i1 *resourcegroups.CreateGroupInput, i2 ...request.Option) (r0 *resourcegroups.CreateGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateGroupWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroups.CreateGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) fakeCreateGroupWithContext(ctx context.Context, in *resourcegroups.CreateGroupInput, _ ...request.Option) (*resourcegroups.CreateGroupOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateGroup cancelled", ctx.Err())
	default:
		return d.CreateGroup(in)
	}
}

func (d *ResourceGroupsDouble) DeleteGroup(i0 *resourcegroups.DeleteGroupInput) (r0 *resourcegroups.DeleteGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteGroup", i0)
	r0, _ = returns[0].(*resourcegroups.DeleteGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) DeleteGroupRequest(i0 *resourcegroups.DeleteGroupInput) (r0 *request.Request, r1 *resourcegroups.DeleteGroupOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteGroupRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroups.DeleteGroupOutput)
	return
}

func (d *ResourceGroupsDouble) DeleteGroupWithContext(i0 context.Context, i1 *resourcegroups.DeleteGroupInput, i2 ...request.Option) (r0 *resourcegroups.DeleteGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteGroupWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroups.DeleteGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) fakeDeleteGroupWithContext(ctx context.Context, in *resourcegroups.DeleteGroupInput, _ ...request.Option) (*resourcegroups.DeleteGroupOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteGroup cancelled", ctx.Err())
	default:
		return d.DeleteGroup(in)
	}
}

func (d *ResourceGroupsDouble) GetGroup(i0 *resourcegroups.GetGroupInput) (r0 *resourcegroups.GetGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetGroup", i0)
	r0, _ = returns[0].(*resourcegroups.GetGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) GetGroupQuery(i0 *resourcegroups.GetGroupQueryInput) (r0 *resourcegroups.GetGroupQueryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetGroupQuery", i0)
	r0, _ = returns[0].(*resourcegroups.GetGroupQueryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) GetGroupQueryRequest(i0 *resourcegroups.GetGroupQueryInput) (r0 *request.Request, r1 *resourcegroups.GetGroupQueryOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetGroupQueryRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroups.GetGroupQueryOutput)
	return
}

func (d *ResourceGroupsDouble) GetGroupQueryWithContext(i0 context.Context, i1 *resourcegroups.GetGroupQueryInput, i2 ...request.Option) (r0 *resourcegroups.GetGroupQueryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetGroupQueryWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroups.GetGroupQueryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) fakeGetGroupQueryWithContext(ctx context.Context, in *resourcegroups.GetGroupQueryInput, _ ...request.Option) (*resourcegroups.GetGroupQueryOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetGroupQuery cancelled", ctx.Err())
	default:
		return d.GetGroupQuery(in)
	}
}

func (d *ResourceGroupsDouble) GetGroupRequest(i0 *resourcegroups.GetGroupInput) (r0 *request.Request, r1 *resourcegroups.GetGroupOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetGroupRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroups.GetGroupOutput)
	return
}

func (d *ResourceGroupsDouble) GetGroupWithContext(i0 context.Context, i1 *resourcegroups.GetGroupInput, i2 ...request.Option) (r0 *resourcegroups.GetGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetGroupWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroups.GetGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) fakeGetGroupWithContext(ctx context.Context, in *resourcegroups.GetGroupInput, _ ...request.Option) (*resourcegroups.GetGroupOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetGroup cancelled", ctx.Err())
	default:
		return d.GetGroup(in)
	}
}

func (d *ResourceGroupsDouble) GetTags(i0 *resourcegroups.GetTagsInput) (r0 *resourcegroups.GetTagsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetTags", i0)
	r0, _ = returns[0].(*resourcegroups.GetTagsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) GetTagsRequest(i0 *resourcegroups.GetTagsInput) (r0 *request.Request, r1 *resourcegroups.GetTagsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetTagsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroups.GetTagsOutput)
	return
}

func (d *ResourceGroupsDouble) GetTagsWithContext(i0 context.Context, i1 *resourcegroups.GetTagsInput, i2 ...request.Option) (r0 *resourcegroups.GetTagsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetTagsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroups.GetTagsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) fakeGetTagsWithContext(ctx context.Context, in *resourcegroups.GetTagsInput, _ ...request.Option) (*resourcegroups.GetTagsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetTags cancelled", ctx.Err())
	default:
		return d.GetTags(in)
	}
}

func (d *ResourceGroupsDouble) ListGroupResources(i0 *resourcegroups.ListGroupResourcesInput) (r0 *resourcegroups.ListGroupResourcesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListGroupResources", i0)
	r0, _ = returns[0].(*resourcegroups.ListGroupResourcesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) ListGroupResourcesPages(i0 *resourcegroups.ListGroupResourcesInput, i1 func(*resourcegroups.ListGroupResourcesOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListGroupResourcesPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *ResourceGroupsDouble) fakeListGroupResourcesPages(in *resourcegroups.ListGroupResourcesInput, pager func(*resourcegroups.ListGroupResourcesOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListGroupResources", paginators, in, pager)
}

func (d *ResourceGroupsDouble) ListGroupResourcesPagesWithContext(i0 context.Context, i1 *resourcegroups.ListGroupResourcesInput, i2 func(*resourcegroups.ListGroupResourcesOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListGroupResourcesPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *ResourceGroupsDouble) fakeListGroupResourcesPagesWithContext(ctx context.Context, in *resourcegroups.ListGroupResourcesInput, pager func(*resourcegroups.ListGroupResourcesOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListGroupResources", paginators, ctx, in, pager, options...)
}

func (d *ResourceGroupsDouble) ListGroupResourcesRequest(i0 *resourcegroups.ListGroupResourcesInput) (r0 *request.Request, r1 *resourcegroups.ListGroupResourcesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListGroupResourcesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroups.ListGroupResourcesOutput)
	return
}

func (d *ResourceGroupsDouble) ListGroupResourcesWithContext(i0 context.Context, i1 *resourcegroups.ListGroupResourcesInput, i2 ...request.Option) (r0 *resourcegroups.ListGroupResourcesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListGroupResourcesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroups.ListGroupResourcesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) fakeListGroupResourcesWithContext(ctx context.Context, in *resourcegroups.ListGroupResourcesInput, _ ...request.Option) (*resourcegroups.ListGroupResourcesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListGroupResources cancelled", ctx.Err())
	default:
		return d.ListGroupResources(in)
	}
}

func (d *ResourceGroupsDouble) ListGroups(i0 *resourcegroups.ListGroupsInput) (r0 *resourcegroups.ListGroupsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListGroups", i0)
	r0, _ = returns[0].(*resourcegroups.ListGroupsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) ListGroupsPages(i0 *resourcegroups.ListGroupsInput, i1 func(*resourcegroups.ListGroupsOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListGroupsPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *ResourceGroupsDouble) fakeListGroupsPages(in *resourcegroups.ListGroupsInput, pager func(*resourcegroups.ListGroupsOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListGroups", paginators, in, pager)
}

func (d *ResourceGroupsDouble) ListGroupsPagesWithContext(i0 context.Context, i1 *resourcegroups.ListGroupsInput, i2 func(*resourcegroups.ListGroupsOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListGroupsPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *ResourceGroupsDouble) fakeListGroupsPagesWithContext(ctx context.Context, in *resourcegroups.ListGroupsInput, pager func(*resourcegroups.ListGroupsOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListGroups", paginators, ctx, in, pager, options...)
}

func (d *ResourceGroupsDouble) ListGroupsRequest(i0 *resourcegroups.ListGroupsInput) (r0 *request.Request, r1 *resourcegroups.ListGroupsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListGroupsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroups.ListGroupsOutput)
	return
}

func (d *ResourceGroupsDouble) ListGroupsWithContext(i0 context.Context, i1 *resourcegroups.ListGroupsInput, i2 ...request.Option) (r0 *resourcegroups.ListGroupsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListGroupsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroups.ListGroupsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) fakeListGroupsWithContext(ctx context.Context, in *resourcegroups.ListGroupsInput, _ ...request.Option) (*resourcegroups.ListGroupsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListGroups cancelled", ctx.Err())
	default:
		return d.ListGroups(in)
	}
}

func (d *ResourceGroupsDouble) SearchResources(i0 *resourcegroups.SearchResourcesInput) (r0 *resourcegroups.SearchResourcesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("SearchResources", i0)
	r0, _ = returns[0].(*resourcegroups.SearchResourcesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) SearchResourcesPages(i0 *resourcegroups.SearchResourcesInput, i1 func(*resourcegroups.SearchResourcesOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("SearchResourcesPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *ResourceGroupsDouble) fakeSearchResourcesPages(in *resourcegroups.SearchResourcesInput, pager func(*resourcegroups.SearchResourcesOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("SearchResources", paginators, in, pager)
}

func (d *ResourceGroupsDouble) SearchResourcesPagesWithContext(i0 context.Context, i1 *resourcegroups.SearchResourcesInput, i2 func(*resourcegroups.SearchResourcesOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("SearchResourcesPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *ResourceGroupsDouble) fakeSearchResourcesPagesWithContext(ctx context.Context, in *resourcegroups.SearchResourcesInput, pager func(*resourcegroups.SearchResourcesOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("SearchResources", paginators, ctx, in, pager, options...)
}

func (d *ResourceGroupsDouble) SearchResourcesRequest(i0 *resourcegroups.SearchResourcesInput) (r0 *request.Request, r1 *resourcegroups.SearchResourcesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("SearchResourcesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroups.SearchResourcesOutput)
	return
}

func (d *ResourceGroupsDouble) SearchResourcesWithContext(i0 context.Context, i1 *resourcegroups.SearchResourcesInput, i2 ...request.Option) (r0 *resourcegroups.SearchResourcesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("SearchResourcesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroups.SearchResourcesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) fakeSearchResourcesWithContext(ctx context.Context, in *resourcegroups.SearchResourcesInput, _ ...request.Option) (*resourcegroups.SearchResourcesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "SearchResources cancelled", ctx.Err())
	default:
		return d.SearchResources(in)
	}
}

func (d *ResourceGroupsDouble) Tag(i0 *resourcegroups.TagInput) (r0 *resourcegroups.TagOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("Tag", i0)
	r0, _ = returns[0].(*resourcegroups.TagOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) TagRequest(i0 *resourcegroups.TagInput) (r0 *request.Request, r1 *resourcegroups.TagOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroups.TagOutput)
	return
}

func (d *ResourceGroupsDouble) TagWithContext(i0 context.Context, i1 *resourcegroups.TagInput, i2 ...request.Option) (r0 *resourcegroups.TagOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroups.TagOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) fakeTagWithContext(ctx context.Context, in *resourcegroups.TagInput, _ ...request.Option) (*resourcegroups.TagOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "Tag cancelled", ctx.Err())
	default:
		return d.Tag(in)
	}
}

func (d *ResourceGroupsDouble) Untag(i0 *resourcegroups.UntagInput) (r0 *resourcegroups.UntagOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("Untag", i0)
	r0, _ = returns[0].(*resourcegroups.UntagOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) UntagRequest(i0 *resourcegroups.UntagInput) (r0 *request.Request, r1 *resourcegroups.UntagOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroups.UntagOutput)
	return
}

func (d *ResourceGroupsDouble) UntagWithContext(i0 context.Context, i1 *resourcegroups.UntagInput, i2 ...request.Option) (r0 *resourcegroups.UntagOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroups.UntagOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) fakeUntagWithContext(ctx context.Context, in *resourcegroups.UntagInput, _ ...request.Option) (*resourcegroups.UntagOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "Untag cancelled", ctx.Err())
	default:
		return d.Untag(in)
	}
}

func (d *ResourceGroupsDouble) UpdateGroup(i0 *resourcegroups.UpdateGroupInput) (r0 *resourcegroups.UpdateGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateGroup", i0)
	r0, _ = returns[0].(*resourcegroups.UpdateGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) UpdateGroupQuery(i0 *resourcegroups.UpdateGroupQueryInput) (r0 *resourcegroups.UpdateGroupQueryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateGroupQuery", i0)
	r0, _ = returns[0].(*resourcegroups.UpdateGroupQueryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) UpdateGroupQueryRequest(i0 *resourcegroups.UpdateGroupQueryInput) (r0 *request.Request, r1 *resourcegroups.UpdateGroupQueryOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateGroupQueryRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroups.UpdateGroupQueryOutput)
	return
}

func (d *ResourceGroupsDouble) UpdateGroupQueryWithContext(i0 context.Context, i1 *resourcegroups.UpdateGroupQueryInput, i2 ...request.Option) (r0 *resourcegroups.UpdateGroupQueryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateGroupQueryWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroups.UpdateGroupQueryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) fakeUpdateGroupQueryWithContext(ctx context.Context, in *resourcegroups.UpdateGroupQueryInput, _ ...request.Option) (*resourcegroups.UpdateGroupQueryOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateGroupQuery cancelled", ctx.Err())
	default:
		return d.UpdateGroupQuery(in)
	}
}

func (d *ResourceGroupsDouble) UpdateGroupRequest(i0 *resourcegroups.UpdateGroupInput) (r0 *request.Request, r1 *resourcegroups.UpdateGroupOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateGroupRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*resourcegroups.UpdateGroupOutput)
	return
}

func (d *ResourceGroupsDouble) UpdateGroupWithContext(i0 context.Context, i1 *resourcegroups.UpdateGroupInput, i2 ...request.Option) (r0 *resourcegroups.UpdateGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateGroupWithContext", i0, i1, i2)
	r0, _ = returns[0].(*resourcegroups.UpdateGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ResourceGroupsDouble) fakeUpdateGroupWithContext(ctx context.Context, in *resourcegroups.UpdateGroupInput, _ ...request.Option) (*resourcegroups.UpdateGroupOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateGroup cancelled", ctx.Err())
	default:
		return d.UpdateGroup(in)
	}
}
