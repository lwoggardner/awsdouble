// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:13:57+11:00

// Package mediastoredouble provides a TestDouble implementation of mediastoreiface.MediaStoreAPI
package mediastoredouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/mediastore"
	"github.com/aws/aws-sdk-go/service/mediastore/mediastoreiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// MediaStoreDouble is TestDouble for mediastoreiface.MediaStoreAPI
type MediaStoreDouble struct {
	mediastoreiface.MediaStoreAPI
	*awsdouble.AWSTestDouble
}

// Constructor for MediaStoreDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *MediaStoreDouble {
	result := &MediaStoreDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*mediastoreiface.MediaStoreAPI)(nil), configurators...)
	return result
}

func (d *MediaStoreDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *MediaStoreDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "CreateContainerWithContext":
		return m.Fake(d.fakeCreateContainerWithContext)

	case "DeleteContainerPolicyWithContext":
		return m.Fake(d.fakeDeleteContainerPolicyWithContext)

	case "DeleteContainerWithContext":
		return m.Fake(d.fakeDeleteContainerWithContext)

	case "DeleteCorsPolicyWithContext":
		return m.Fake(d.fakeDeleteCorsPolicyWithContext)

	case "DeleteLifecyclePolicyWithContext":
		return m.Fake(d.fakeDeleteLifecyclePolicyWithContext)

	case "DescribeContainerWithContext":
		return m.Fake(d.fakeDescribeContainerWithContext)

	case "GetContainerPolicyWithContext":
		return m.Fake(d.fakeGetContainerPolicyWithContext)

	case "GetCorsPolicyWithContext":
		return m.Fake(d.fakeGetCorsPolicyWithContext)

	case "GetLifecyclePolicyWithContext":
		return m.Fake(d.fakeGetLifecyclePolicyWithContext)

	case "ListContainersPages":
		return m.Fake(d.fakeListContainersPages)

	case "ListContainersPagesWithContext":
		return m.Fake(d.fakeListContainersPagesWithContext)

	case "ListContainersWithContext":
		return m.Fake(d.fakeListContainersWithContext)

	case "ListTagsForResourceWithContext":
		return m.Fake(d.fakeListTagsForResourceWithContext)

	case "PutContainerPolicyWithContext":
		return m.Fake(d.fakePutContainerPolicyWithContext)

	case "PutCorsPolicyWithContext":
		return m.Fake(d.fakePutCorsPolicyWithContext)

	case "PutLifecyclePolicyWithContext":
		return m.Fake(d.fakePutLifecyclePolicyWithContext)

	case "StartAccessLoggingWithContext":
		return m.Fake(d.fakeStartAccessLoggingWithContext)

	case "StopAccessLoggingWithContext":
		return m.Fake(d.fakeStopAccessLoggingWithContext)

	case "TagResourceWithContext":
		return m.Fake(d.fakeTagResourceWithContext)

	case "UntagResourceWithContext":
		return m.Fake(d.fakeUntagResourceWithContext)

	default:
		return nil
	}
}

func (d *MediaStoreDouble) CreateContainer(i0 *mediastore.CreateContainerInput) (r0 *mediastore.CreateContainerOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateContainer", i0)
	r0, _ = returns[0].(*mediastore.CreateContainerOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) CreateContainerRequest(i0 *mediastore.CreateContainerInput) (r0 *request.Request, r1 *mediastore.CreateContainerOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateContainerRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mediastore.CreateContainerOutput)
	return
}

func (d *MediaStoreDouble) CreateContainerWithContext(i0 context.Context, i1 *mediastore.CreateContainerInput, i2 ...request.Option) (r0 *mediastore.CreateContainerOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateContainerWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mediastore.CreateContainerOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) fakeCreateContainerWithContext(ctx context.Context, in *mediastore.CreateContainerInput, _ ...request.Option) (*mediastore.CreateContainerOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateContainer cancelled", ctx.Err())
	default:
		return d.CreateContainer(in)
	}
}

func (d *MediaStoreDouble) DeleteContainer(i0 *mediastore.DeleteContainerInput) (r0 *mediastore.DeleteContainerOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteContainer", i0)
	r0, _ = returns[0].(*mediastore.DeleteContainerOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) DeleteContainerPolicy(i0 *mediastore.DeleteContainerPolicyInput) (r0 *mediastore.DeleteContainerPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteContainerPolicy", i0)
	r0, _ = returns[0].(*mediastore.DeleteContainerPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) DeleteContainerPolicyRequest(i0 *mediastore.DeleteContainerPolicyInput) (r0 *request.Request, r1 *mediastore.DeleteContainerPolicyOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteContainerPolicyRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mediastore.DeleteContainerPolicyOutput)
	return
}

func (d *MediaStoreDouble) DeleteContainerPolicyWithContext(i0 context.Context, i1 *mediastore.DeleteContainerPolicyInput, i2 ...request.Option) (r0 *mediastore.DeleteContainerPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteContainerPolicyWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mediastore.DeleteContainerPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) fakeDeleteContainerPolicyWithContext(ctx context.Context, in *mediastore.DeleteContainerPolicyInput, _ ...request.Option) (*mediastore.DeleteContainerPolicyOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteContainerPolicy cancelled", ctx.Err())
	default:
		return d.DeleteContainerPolicy(in)
	}
}

func (d *MediaStoreDouble) DeleteContainerRequest(i0 *mediastore.DeleteContainerInput) (r0 *request.Request, r1 *mediastore.DeleteContainerOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteContainerRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mediastore.DeleteContainerOutput)
	return
}

func (d *MediaStoreDouble) DeleteContainerWithContext(i0 context.Context, i1 *mediastore.DeleteContainerInput, i2 ...request.Option) (r0 *mediastore.DeleteContainerOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteContainerWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mediastore.DeleteContainerOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) fakeDeleteContainerWithContext(ctx context.Context, in *mediastore.DeleteContainerInput, _ ...request.Option) (*mediastore.DeleteContainerOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteContainer cancelled", ctx.Err())
	default:
		return d.DeleteContainer(in)
	}
}

func (d *MediaStoreDouble) DeleteCorsPolicy(i0 *mediastore.DeleteCorsPolicyInput) (r0 *mediastore.DeleteCorsPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteCorsPolicy", i0)
	r0, _ = returns[0].(*mediastore.DeleteCorsPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) DeleteCorsPolicyRequest(i0 *mediastore.DeleteCorsPolicyInput) (r0 *request.Request, r1 *mediastore.DeleteCorsPolicyOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteCorsPolicyRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mediastore.DeleteCorsPolicyOutput)
	return
}

func (d *MediaStoreDouble) DeleteCorsPolicyWithContext(i0 context.Context, i1 *mediastore.DeleteCorsPolicyInput, i2 ...request.Option) (r0 *mediastore.DeleteCorsPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteCorsPolicyWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mediastore.DeleteCorsPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) fakeDeleteCorsPolicyWithContext(ctx context.Context, in *mediastore.DeleteCorsPolicyInput, _ ...request.Option) (*mediastore.DeleteCorsPolicyOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteCorsPolicy cancelled", ctx.Err())
	default:
		return d.DeleteCorsPolicy(in)
	}
}

func (d *MediaStoreDouble) DeleteLifecyclePolicy(i0 *mediastore.DeleteLifecyclePolicyInput) (r0 *mediastore.DeleteLifecyclePolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteLifecyclePolicy", i0)
	r0, _ = returns[0].(*mediastore.DeleteLifecyclePolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) DeleteLifecyclePolicyRequest(i0 *mediastore.DeleteLifecyclePolicyInput) (r0 *request.Request, r1 *mediastore.DeleteLifecyclePolicyOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteLifecyclePolicyRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mediastore.DeleteLifecyclePolicyOutput)
	return
}

func (d *MediaStoreDouble) DeleteLifecyclePolicyWithContext(i0 context.Context, i1 *mediastore.DeleteLifecyclePolicyInput, i2 ...request.Option) (r0 *mediastore.DeleteLifecyclePolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteLifecyclePolicyWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mediastore.DeleteLifecyclePolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) fakeDeleteLifecyclePolicyWithContext(ctx context.Context, in *mediastore.DeleteLifecyclePolicyInput, _ ...request.Option) (*mediastore.DeleteLifecyclePolicyOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteLifecyclePolicy cancelled", ctx.Err())
	default:
		return d.DeleteLifecyclePolicy(in)
	}
}

func (d *MediaStoreDouble) DescribeContainer(i0 *mediastore.DescribeContainerInput) (r0 *mediastore.DescribeContainerOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeContainer", i0)
	r0, _ = returns[0].(*mediastore.DescribeContainerOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) DescribeContainerRequest(i0 *mediastore.DescribeContainerInput) (r0 *request.Request, r1 *mediastore.DescribeContainerOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeContainerRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mediastore.DescribeContainerOutput)
	return
}

func (d *MediaStoreDouble) DescribeContainerWithContext(i0 context.Context, i1 *mediastore.DescribeContainerInput, i2 ...request.Option) (r0 *mediastore.DescribeContainerOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeContainerWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mediastore.DescribeContainerOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) fakeDescribeContainerWithContext(ctx context.Context, in *mediastore.DescribeContainerInput, _ ...request.Option) (*mediastore.DescribeContainerOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeContainer cancelled", ctx.Err())
	default:
		return d.DescribeContainer(in)
	}
}

func (d *MediaStoreDouble) GetContainerPolicy(i0 *mediastore.GetContainerPolicyInput) (r0 *mediastore.GetContainerPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetContainerPolicy", i0)
	r0, _ = returns[0].(*mediastore.GetContainerPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) GetContainerPolicyRequest(i0 *mediastore.GetContainerPolicyInput) (r0 *request.Request, r1 *mediastore.GetContainerPolicyOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetContainerPolicyRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mediastore.GetContainerPolicyOutput)
	return
}

func (d *MediaStoreDouble) GetContainerPolicyWithContext(i0 context.Context, i1 *mediastore.GetContainerPolicyInput, i2 ...request.Option) (r0 *mediastore.GetContainerPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetContainerPolicyWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mediastore.GetContainerPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) fakeGetContainerPolicyWithContext(ctx context.Context, in *mediastore.GetContainerPolicyInput, _ ...request.Option) (*mediastore.GetContainerPolicyOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetContainerPolicy cancelled", ctx.Err())
	default:
		return d.GetContainerPolicy(in)
	}
}

func (d *MediaStoreDouble) GetCorsPolicy(i0 *mediastore.GetCorsPolicyInput) (r0 *mediastore.GetCorsPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetCorsPolicy", i0)
	r0, _ = returns[0].(*mediastore.GetCorsPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) GetCorsPolicyRequest(i0 *mediastore.GetCorsPolicyInput) (r0 *request.Request, r1 *mediastore.GetCorsPolicyOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetCorsPolicyRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mediastore.GetCorsPolicyOutput)
	return
}

func (d *MediaStoreDouble) GetCorsPolicyWithContext(i0 context.Context, i1 *mediastore.GetCorsPolicyInput, i2 ...request.Option) (r0 *mediastore.GetCorsPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetCorsPolicyWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mediastore.GetCorsPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) fakeGetCorsPolicyWithContext(ctx context.Context, in *mediastore.GetCorsPolicyInput, _ ...request.Option) (*mediastore.GetCorsPolicyOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetCorsPolicy cancelled", ctx.Err())
	default:
		return d.GetCorsPolicy(in)
	}
}

func (d *MediaStoreDouble) GetLifecyclePolicy(i0 *mediastore.GetLifecyclePolicyInput) (r0 *mediastore.GetLifecyclePolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetLifecyclePolicy", i0)
	r0, _ = returns[0].(*mediastore.GetLifecyclePolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) GetLifecyclePolicyRequest(i0 *mediastore.GetLifecyclePolicyInput) (r0 *request.Request, r1 *mediastore.GetLifecyclePolicyOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetLifecyclePolicyRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mediastore.GetLifecyclePolicyOutput)
	return
}

func (d *MediaStoreDouble) GetLifecyclePolicyWithContext(i0 context.Context, i1 *mediastore.GetLifecyclePolicyInput, i2 ...request.Option) (r0 *mediastore.GetLifecyclePolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetLifecyclePolicyWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mediastore.GetLifecyclePolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) fakeGetLifecyclePolicyWithContext(ctx context.Context, in *mediastore.GetLifecyclePolicyInput, _ ...request.Option) (*mediastore.GetLifecyclePolicyOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetLifecyclePolicy cancelled", ctx.Err())
	default:
		return d.GetLifecyclePolicy(in)
	}
}

func (d *MediaStoreDouble) ListContainers(i0 *mediastore.ListContainersInput) (r0 *mediastore.ListContainersOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListContainers", i0)
	r0, _ = returns[0].(*mediastore.ListContainersOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) ListContainersPages(i0 *mediastore.ListContainersInput, i1 func(*mediastore.ListContainersOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListContainersPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *MediaStoreDouble) fakeListContainersPages(in *mediastore.ListContainersInput, pager func(*mediastore.ListContainersOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListContainers", paginators, in, pager)
}

func (d *MediaStoreDouble) ListContainersPagesWithContext(i0 context.Context, i1 *mediastore.ListContainersInput, i2 func(*mediastore.ListContainersOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListContainersPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *MediaStoreDouble) fakeListContainersPagesWithContext(ctx context.Context, in *mediastore.ListContainersInput, pager func(*mediastore.ListContainersOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListContainers", paginators, ctx, in, pager, options...)
}

func (d *MediaStoreDouble) ListContainersRequest(i0 *mediastore.ListContainersInput) (r0 *request.Request, r1 *mediastore.ListContainersOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListContainersRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mediastore.ListContainersOutput)
	return
}

func (d *MediaStoreDouble) ListContainersWithContext(i0 context.Context, i1 *mediastore.ListContainersInput, i2 ...request.Option) (r0 *mediastore.ListContainersOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListContainersWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mediastore.ListContainersOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) fakeListContainersWithContext(ctx context.Context, in *mediastore.ListContainersInput, _ ...request.Option) (*mediastore.ListContainersOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListContainers cancelled", ctx.Err())
	default:
		return d.ListContainers(in)
	}
}

func (d *MediaStoreDouble) ListTagsForResource(i0 *mediastore.ListTagsForResourceInput) (r0 *mediastore.ListTagsForResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForResource", i0)
	r0, _ = returns[0].(*mediastore.ListTagsForResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) ListTagsForResourceRequest(i0 *mediastore.ListTagsForResourceInput) (r0 *request.Request, r1 *mediastore.ListTagsForResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mediastore.ListTagsForResourceOutput)
	return
}

func (d *MediaStoreDouble) ListTagsForResourceWithContext(i0 context.Context, i1 *mediastore.ListTagsForResourceInput, i2 ...request.Option) (r0 *mediastore.ListTagsForResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mediastore.ListTagsForResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) fakeListTagsForResourceWithContext(ctx context.Context, in *mediastore.ListTagsForResourceInput, _ ...request.Option) (*mediastore.ListTagsForResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListTagsForResource cancelled", ctx.Err())
	default:
		return d.ListTagsForResource(in)
	}
}

func (d *MediaStoreDouble) PutContainerPolicy(i0 *mediastore.PutContainerPolicyInput) (r0 *mediastore.PutContainerPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutContainerPolicy", i0)
	r0, _ = returns[0].(*mediastore.PutContainerPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) PutContainerPolicyRequest(i0 *mediastore.PutContainerPolicyInput) (r0 *request.Request, r1 *mediastore.PutContainerPolicyOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutContainerPolicyRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mediastore.PutContainerPolicyOutput)
	return
}

func (d *MediaStoreDouble) PutContainerPolicyWithContext(i0 context.Context, i1 *mediastore.PutContainerPolicyInput, i2 ...request.Option) (r0 *mediastore.PutContainerPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutContainerPolicyWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mediastore.PutContainerPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) fakePutContainerPolicyWithContext(ctx context.Context, in *mediastore.PutContainerPolicyInput, _ ...request.Option) (*mediastore.PutContainerPolicyOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "PutContainerPolicy cancelled", ctx.Err())
	default:
		return d.PutContainerPolicy(in)
	}
}

func (d *MediaStoreDouble) PutCorsPolicy(i0 *mediastore.PutCorsPolicyInput) (r0 *mediastore.PutCorsPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutCorsPolicy", i0)
	r0, _ = returns[0].(*mediastore.PutCorsPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) PutCorsPolicyRequest(i0 *mediastore.PutCorsPolicyInput) (r0 *request.Request, r1 *mediastore.PutCorsPolicyOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutCorsPolicyRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mediastore.PutCorsPolicyOutput)
	return
}

func (d *MediaStoreDouble) PutCorsPolicyWithContext(i0 context.Context, i1 *mediastore.PutCorsPolicyInput, i2 ...request.Option) (r0 *mediastore.PutCorsPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutCorsPolicyWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mediastore.PutCorsPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) fakePutCorsPolicyWithContext(ctx context.Context, in *mediastore.PutCorsPolicyInput, _ ...request.Option) (*mediastore.PutCorsPolicyOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "PutCorsPolicy cancelled", ctx.Err())
	default:
		return d.PutCorsPolicy(in)
	}
}

func (d *MediaStoreDouble) PutLifecyclePolicy(i0 *mediastore.PutLifecyclePolicyInput) (r0 *mediastore.PutLifecyclePolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutLifecyclePolicy", i0)
	r0, _ = returns[0].(*mediastore.PutLifecyclePolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) PutLifecyclePolicyRequest(i0 *mediastore.PutLifecyclePolicyInput) (r0 *request.Request, r1 *mediastore.PutLifecyclePolicyOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutLifecyclePolicyRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mediastore.PutLifecyclePolicyOutput)
	return
}

func (d *MediaStoreDouble) PutLifecyclePolicyWithContext(i0 context.Context, i1 *mediastore.PutLifecyclePolicyInput, i2 ...request.Option) (r0 *mediastore.PutLifecyclePolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutLifecyclePolicyWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mediastore.PutLifecyclePolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) fakePutLifecyclePolicyWithContext(ctx context.Context, in *mediastore.PutLifecyclePolicyInput, _ ...request.Option) (*mediastore.PutLifecyclePolicyOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "PutLifecyclePolicy cancelled", ctx.Err())
	default:
		return d.PutLifecyclePolicy(in)
	}
}

func (d *MediaStoreDouble) StartAccessLogging(i0 *mediastore.StartAccessLoggingInput) (r0 *mediastore.StartAccessLoggingOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartAccessLogging", i0)
	r0, _ = returns[0].(*mediastore.StartAccessLoggingOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) StartAccessLoggingRequest(i0 *mediastore.StartAccessLoggingInput) (r0 *request.Request, r1 *mediastore.StartAccessLoggingOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartAccessLoggingRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mediastore.StartAccessLoggingOutput)
	return
}

func (d *MediaStoreDouble) StartAccessLoggingWithContext(i0 context.Context, i1 *mediastore.StartAccessLoggingInput, i2 ...request.Option) (r0 *mediastore.StartAccessLoggingOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartAccessLoggingWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mediastore.StartAccessLoggingOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) fakeStartAccessLoggingWithContext(ctx context.Context, in *mediastore.StartAccessLoggingInput, _ ...request.Option) (*mediastore.StartAccessLoggingOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StartAccessLogging cancelled", ctx.Err())
	default:
		return d.StartAccessLogging(in)
	}
}

func (d *MediaStoreDouble) StopAccessLogging(i0 *mediastore.StopAccessLoggingInput) (r0 *mediastore.StopAccessLoggingOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopAccessLogging", i0)
	r0, _ = returns[0].(*mediastore.StopAccessLoggingOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) StopAccessLoggingRequest(i0 *mediastore.StopAccessLoggingInput) (r0 *request.Request, r1 *mediastore.StopAccessLoggingOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopAccessLoggingRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mediastore.StopAccessLoggingOutput)
	return
}

func (d *MediaStoreDouble) StopAccessLoggingWithContext(i0 context.Context, i1 *mediastore.StopAccessLoggingInput, i2 ...request.Option) (r0 *mediastore.StopAccessLoggingOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopAccessLoggingWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mediastore.StopAccessLoggingOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) fakeStopAccessLoggingWithContext(ctx context.Context, in *mediastore.StopAccessLoggingInput, _ ...request.Option) (*mediastore.StopAccessLoggingOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StopAccessLogging cancelled", ctx.Err())
	default:
		return d.StopAccessLogging(in)
	}
}

func (d *MediaStoreDouble) TagResource(i0 *mediastore.TagResourceInput) (r0 *mediastore.TagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResource", i0)
	r0, _ = returns[0].(*mediastore.TagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) TagResourceRequest(i0 *mediastore.TagResourceInput) (r0 *request.Request, r1 *mediastore.TagResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mediastore.TagResourceOutput)
	return
}

func (d *MediaStoreDouble) TagResourceWithContext(i0 context.Context, i1 *mediastore.TagResourceInput, i2 ...request.Option) (r0 *mediastore.TagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mediastore.TagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) fakeTagResourceWithContext(ctx context.Context, in *mediastore.TagResourceInput, _ ...request.Option) (*mediastore.TagResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "TagResource cancelled", ctx.Err())
	default:
		return d.TagResource(in)
	}
}

func (d *MediaStoreDouble) UntagResource(i0 *mediastore.UntagResourceInput) (r0 *mediastore.UntagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResource", i0)
	r0, _ = returns[0].(*mediastore.UntagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) UntagResourceRequest(i0 *mediastore.UntagResourceInput) (r0 *request.Request, r1 *mediastore.UntagResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mediastore.UntagResourceOutput)
	return
}

func (d *MediaStoreDouble) UntagResourceWithContext(i0 context.Context, i1 *mediastore.UntagResourceInput, i2 ...request.Option) (r0 *mediastore.UntagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mediastore.UntagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MediaStoreDouble) fakeUntagResourceWithContext(ctx context.Context, in *mediastore.UntagResourceInput, _ ...request.Option) (*mediastore.UntagResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UntagResource cancelled", ctx.Err())
	default:
		return d.UntagResource(in)
	}
}
