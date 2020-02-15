// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:12:22+11:00

// Package dlmdouble provides a TestDouble implementation of dlmiface.DLMAPI
package dlmdouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dlm"
	"github.com/aws/aws-sdk-go/service/dlm/dlmiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// DLMDouble is TestDouble for dlmiface.DLMAPI
type DLMDouble struct {
	dlmiface.DLMAPI
	*awsdouble.AWSTestDouble
}

// Constructor for DLMDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *DLMDouble {
	result := &DLMDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*dlmiface.DLMAPI)(nil), configurators...)
	return result
}

func (d *DLMDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *DLMDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "CreateLifecyclePolicyWithContext":
		return m.Fake(d.fakeCreateLifecyclePolicyWithContext)

	case "DeleteLifecyclePolicyWithContext":
		return m.Fake(d.fakeDeleteLifecyclePolicyWithContext)

	case "GetLifecyclePoliciesWithContext":
		return m.Fake(d.fakeGetLifecyclePoliciesWithContext)

	case "GetLifecyclePolicyWithContext":
		return m.Fake(d.fakeGetLifecyclePolicyWithContext)

	case "ListTagsForResourceWithContext":
		return m.Fake(d.fakeListTagsForResourceWithContext)

	case "TagResourceWithContext":
		return m.Fake(d.fakeTagResourceWithContext)

	case "UntagResourceWithContext":
		return m.Fake(d.fakeUntagResourceWithContext)

	case "UpdateLifecyclePolicyWithContext":
		return m.Fake(d.fakeUpdateLifecyclePolicyWithContext)

	default:
		return nil
	}
}

func (d *DLMDouble) CreateLifecyclePolicy(i0 *dlm.CreateLifecyclePolicyInput) (r0 *dlm.CreateLifecyclePolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateLifecyclePolicy", i0)
	r0, _ = returns[0].(*dlm.CreateLifecyclePolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *DLMDouble) CreateLifecyclePolicyRequest(i0 *dlm.CreateLifecyclePolicyInput) (r0 *request.Request, r1 *dlm.CreateLifecyclePolicyOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateLifecyclePolicyRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*dlm.CreateLifecyclePolicyOutput)
	return
}

func (d *DLMDouble) CreateLifecyclePolicyWithContext(i0 context.Context, i1 *dlm.CreateLifecyclePolicyInput, i2 ...request.Option) (r0 *dlm.CreateLifecyclePolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateLifecyclePolicyWithContext", i0, i1, i2)
	r0, _ = returns[0].(*dlm.CreateLifecyclePolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *DLMDouble) fakeCreateLifecyclePolicyWithContext(ctx context.Context, in *dlm.CreateLifecyclePolicyInput, _ ...request.Option) (*dlm.CreateLifecyclePolicyOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateLifecyclePolicy cancelled", ctx.Err())
	default:
		return d.CreateLifecyclePolicy(in)
	}
}

func (d *DLMDouble) DeleteLifecyclePolicy(i0 *dlm.DeleteLifecyclePolicyInput) (r0 *dlm.DeleteLifecyclePolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteLifecyclePolicy", i0)
	r0, _ = returns[0].(*dlm.DeleteLifecyclePolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *DLMDouble) DeleteLifecyclePolicyRequest(i0 *dlm.DeleteLifecyclePolicyInput) (r0 *request.Request, r1 *dlm.DeleteLifecyclePolicyOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteLifecyclePolicyRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*dlm.DeleteLifecyclePolicyOutput)
	return
}

func (d *DLMDouble) DeleteLifecyclePolicyWithContext(i0 context.Context, i1 *dlm.DeleteLifecyclePolicyInput, i2 ...request.Option) (r0 *dlm.DeleteLifecyclePolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteLifecyclePolicyWithContext", i0, i1, i2)
	r0, _ = returns[0].(*dlm.DeleteLifecyclePolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *DLMDouble) fakeDeleteLifecyclePolicyWithContext(ctx context.Context, in *dlm.DeleteLifecyclePolicyInput, _ ...request.Option) (*dlm.DeleteLifecyclePolicyOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteLifecyclePolicy cancelled", ctx.Err())
	default:
		return d.DeleteLifecyclePolicy(in)
	}
}

func (d *DLMDouble) GetLifecyclePolicies(i0 *dlm.GetLifecyclePoliciesInput) (r0 *dlm.GetLifecyclePoliciesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetLifecyclePolicies", i0)
	r0, _ = returns[0].(*dlm.GetLifecyclePoliciesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *DLMDouble) GetLifecyclePoliciesRequest(i0 *dlm.GetLifecyclePoliciesInput) (r0 *request.Request, r1 *dlm.GetLifecyclePoliciesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetLifecyclePoliciesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*dlm.GetLifecyclePoliciesOutput)
	return
}

func (d *DLMDouble) GetLifecyclePoliciesWithContext(i0 context.Context, i1 *dlm.GetLifecyclePoliciesInput, i2 ...request.Option) (r0 *dlm.GetLifecyclePoliciesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetLifecyclePoliciesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*dlm.GetLifecyclePoliciesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *DLMDouble) fakeGetLifecyclePoliciesWithContext(ctx context.Context, in *dlm.GetLifecyclePoliciesInput, _ ...request.Option) (*dlm.GetLifecyclePoliciesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetLifecyclePolicies cancelled", ctx.Err())
	default:
		return d.GetLifecyclePolicies(in)
	}
}

func (d *DLMDouble) GetLifecyclePolicy(i0 *dlm.GetLifecyclePolicyInput) (r0 *dlm.GetLifecyclePolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetLifecyclePolicy", i0)
	r0, _ = returns[0].(*dlm.GetLifecyclePolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *DLMDouble) GetLifecyclePolicyRequest(i0 *dlm.GetLifecyclePolicyInput) (r0 *request.Request, r1 *dlm.GetLifecyclePolicyOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetLifecyclePolicyRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*dlm.GetLifecyclePolicyOutput)
	return
}

func (d *DLMDouble) GetLifecyclePolicyWithContext(i0 context.Context, i1 *dlm.GetLifecyclePolicyInput, i2 ...request.Option) (r0 *dlm.GetLifecyclePolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetLifecyclePolicyWithContext", i0, i1, i2)
	r0, _ = returns[0].(*dlm.GetLifecyclePolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *DLMDouble) fakeGetLifecyclePolicyWithContext(ctx context.Context, in *dlm.GetLifecyclePolicyInput, _ ...request.Option) (*dlm.GetLifecyclePolicyOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetLifecyclePolicy cancelled", ctx.Err())
	default:
		return d.GetLifecyclePolicy(in)
	}
}

func (d *DLMDouble) ListTagsForResource(i0 *dlm.ListTagsForResourceInput) (r0 *dlm.ListTagsForResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForResource", i0)
	r0, _ = returns[0].(*dlm.ListTagsForResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *DLMDouble) ListTagsForResourceRequest(i0 *dlm.ListTagsForResourceInput) (r0 *request.Request, r1 *dlm.ListTagsForResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*dlm.ListTagsForResourceOutput)
	return
}

func (d *DLMDouble) ListTagsForResourceWithContext(i0 context.Context, i1 *dlm.ListTagsForResourceInput, i2 ...request.Option) (r0 *dlm.ListTagsForResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*dlm.ListTagsForResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *DLMDouble) fakeListTagsForResourceWithContext(ctx context.Context, in *dlm.ListTagsForResourceInput, _ ...request.Option) (*dlm.ListTagsForResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListTagsForResource cancelled", ctx.Err())
	default:
		return d.ListTagsForResource(in)
	}
}

func (d *DLMDouble) TagResource(i0 *dlm.TagResourceInput) (r0 *dlm.TagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResource", i0)
	r0, _ = returns[0].(*dlm.TagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *DLMDouble) TagResourceRequest(i0 *dlm.TagResourceInput) (r0 *request.Request, r1 *dlm.TagResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*dlm.TagResourceOutput)
	return
}

func (d *DLMDouble) TagResourceWithContext(i0 context.Context, i1 *dlm.TagResourceInput, i2 ...request.Option) (r0 *dlm.TagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*dlm.TagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *DLMDouble) fakeTagResourceWithContext(ctx context.Context, in *dlm.TagResourceInput, _ ...request.Option) (*dlm.TagResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "TagResource cancelled", ctx.Err())
	default:
		return d.TagResource(in)
	}
}

func (d *DLMDouble) UntagResource(i0 *dlm.UntagResourceInput) (r0 *dlm.UntagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResource", i0)
	r0, _ = returns[0].(*dlm.UntagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *DLMDouble) UntagResourceRequest(i0 *dlm.UntagResourceInput) (r0 *request.Request, r1 *dlm.UntagResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*dlm.UntagResourceOutput)
	return
}

func (d *DLMDouble) UntagResourceWithContext(i0 context.Context, i1 *dlm.UntagResourceInput, i2 ...request.Option) (r0 *dlm.UntagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*dlm.UntagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *DLMDouble) fakeUntagResourceWithContext(ctx context.Context, in *dlm.UntagResourceInput, _ ...request.Option) (*dlm.UntagResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UntagResource cancelled", ctx.Err())
	default:
		return d.UntagResource(in)
	}
}

func (d *DLMDouble) UpdateLifecyclePolicy(i0 *dlm.UpdateLifecyclePolicyInput) (r0 *dlm.UpdateLifecyclePolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateLifecyclePolicy", i0)
	r0, _ = returns[0].(*dlm.UpdateLifecyclePolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *DLMDouble) UpdateLifecyclePolicyRequest(i0 *dlm.UpdateLifecyclePolicyInput) (r0 *request.Request, r1 *dlm.UpdateLifecyclePolicyOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateLifecyclePolicyRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*dlm.UpdateLifecyclePolicyOutput)
	return
}

func (d *DLMDouble) UpdateLifecyclePolicyWithContext(i0 context.Context, i1 *dlm.UpdateLifecyclePolicyInput, i2 ...request.Option) (r0 *dlm.UpdateLifecyclePolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateLifecyclePolicyWithContext", i0, i1, i2)
	r0, _ = returns[0].(*dlm.UpdateLifecyclePolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *DLMDouble) fakeUpdateLifecyclePolicyWithContext(ctx context.Context, in *dlm.UpdateLifecyclePolicyInput, _ ...request.Option) (*dlm.UpdateLifecyclePolicyOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateLifecyclePolicy cancelled", ctx.Err())
	default:
		return d.UpdateLifecyclePolicy(in)
	}
}