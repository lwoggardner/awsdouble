// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:15:06+11:00

// Package signerdouble provides a TestDouble implementation of signeriface.SignerAPI
package signerdouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/signer"
	"github.com/aws/aws-sdk-go/service/signer/signeriface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// SignerDouble is TestDouble for signeriface.SignerAPI
type SignerDouble struct {
	signeriface.SignerAPI
	*awsdouble.AWSTestDouble
}

// Constructor for SignerDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *SignerDouble {
	result := &SignerDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*signeriface.SignerAPI)(nil), configurators...)
	return result
}

func (d *SignerDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *SignerDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "CancelSigningProfileWithContext":
		return m.Fake(d.fakeCancelSigningProfileWithContext)

	case "DescribeSigningJobWithContext":
		return m.Fake(d.fakeDescribeSigningJobWithContext)

	case "GetSigningPlatformWithContext":
		return m.Fake(d.fakeGetSigningPlatformWithContext)

	case "GetSigningProfileWithContext":
		return m.Fake(d.fakeGetSigningProfileWithContext)

	case "ListSigningJobsPages":
		return m.Fake(d.fakeListSigningJobsPages)

	case "ListSigningJobsPagesWithContext":
		return m.Fake(d.fakeListSigningJobsPagesWithContext)

	case "ListSigningJobsWithContext":
		return m.Fake(d.fakeListSigningJobsWithContext)

	case "ListSigningPlatformsPages":
		return m.Fake(d.fakeListSigningPlatformsPages)

	case "ListSigningPlatformsPagesWithContext":
		return m.Fake(d.fakeListSigningPlatformsPagesWithContext)

	case "ListSigningPlatformsWithContext":
		return m.Fake(d.fakeListSigningPlatformsWithContext)

	case "ListSigningProfilesPages":
		return m.Fake(d.fakeListSigningProfilesPages)

	case "ListSigningProfilesPagesWithContext":
		return m.Fake(d.fakeListSigningProfilesPagesWithContext)

	case "ListSigningProfilesWithContext":
		return m.Fake(d.fakeListSigningProfilesWithContext)

	case "ListTagsForResourceWithContext":
		return m.Fake(d.fakeListTagsForResourceWithContext)

	case "PutSigningProfileWithContext":
		return m.Fake(d.fakePutSigningProfileWithContext)

	case "StartSigningJobWithContext":
		return m.Fake(d.fakeStartSigningJobWithContext)

	case "TagResourceWithContext":
		return m.Fake(d.fakeTagResourceWithContext)

	case "UntagResourceWithContext":
		return m.Fake(d.fakeUntagResourceWithContext)

	case "WaitUntilSuccessfulSigningJob":
		return m.Fake(d.fakeWaitUntilSuccessfulSigningJob)

	case "WaitUntilSuccessfulSigningJobWithContext":
		return m.Fake(d.fakeWaitUntilSuccessfulSigningJobWithContext)

	default:
		return nil
	}
}

func (d *SignerDouble) CancelSigningProfile(i0 *signer.CancelSigningProfileInput) (r0 *signer.CancelSigningProfileOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CancelSigningProfile", i0)
	r0, _ = returns[0].(*signer.CancelSigningProfileOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) CancelSigningProfileRequest(i0 *signer.CancelSigningProfileInput) (r0 *request.Request, r1 *signer.CancelSigningProfileOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CancelSigningProfileRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*signer.CancelSigningProfileOutput)
	return
}

func (d *SignerDouble) CancelSigningProfileWithContext(i0 context.Context, i1 *signer.CancelSigningProfileInput, i2 ...request.Option) (r0 *signer.CancelSigningProfileOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CancelSigningProfileWithContext", i0, i1, i2)
	r0, _ = returns[0].(*signer.CancelSigningProfileOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) fakeCancelSigningProfileWithContext(ctx context.Context, in *signer.CancelSigningProfileInput, _ ...request.Option) (*signer.CancelSigningProfileOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CancelSigningProfile cancelled", ctx.Err())
	default:
		return d.CancelSigningProfile(in)
	}
}

func (d *SignerDouble) DescribeSigningJob(i0 *signer.DescribeSigningJobInput) (r0 *signer.DescribeSigningJobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeSigningJob", i0)
	r0, _ = returns[0].(*signer.DescribeSigningJobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) DescribeSigningJobRequest(i0 *signer.DescribeSigningJobInput) (r0 *request.Request, r1 *signer.DescribeSigningJobOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeSigningJobRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*signer.DescribeSigningJobOutput)
	return
}

func (d *SignerDouble) DescribeSigningJobWithContext(i0 context.Context, i1 *signer.DescribeSigningJobInput, i2 ...request.Option) (r0 *signer.DescribeSigningJobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeSigningJobWithContext", i0, i1, i2)
	r0, _ = returns[0].(*signer.DescribeSigningJobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) fakeDescribeSigningJobWithContext(ctx context.Context, in *signer.DescribeSigningJobInput, _ ...request.Option) (*signer.DescribeSigningJobOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeSigningJob cancelled", ctx.Err())
	default:
		return d.DescribeSigningJob(in)
	}
}

func (d *SignerDouble) GetSigningPlatform(i0 *signer.GetSigningPlatformInput) (r0 *signer.GetSigningPlatformOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSigningPlatform", i0)
	r0, _ = returns[0].(*signer.GetSigningPlatformOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) GetSigningPlatformRequest(i0 *signer.GetSigningPlatformInput) (r0 *request.Request, r1 *signer.GetSigningPlatformOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSigningPlatformRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*signer.GetSigningPlatformOutput)
	return
}

func (d *SignerDouble) GetSigningPlatformWithContext(i0 context.Context, i1 *signer.GetSigningPlatformInput, i2 ...request.Option) (r0 *signer.GetSigningPlatformOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSigningPlatformWithContext", i0, i1, i2)
	r0, _ = returns[0].(*signer.GetSigningPlatformOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) fakeGetSigningPlatformWithContext(ctx context.Context, in *signer.GetSigningPlatformInput, _ ...request.Option) (*signer.GetSigningPlatformOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetSigningPlatform cancelled", ctx.Err())
	default:
		return d.GetSigningPlatform(in)
	}
}

func (d *SignerDouble) GetSigningProfile(i0 *signer.GetSigningProfileInput) (r0 *signer.GetSigningProfileOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSigningProfile", i0)
	r0, _ = returns[0].(*signer.GetSigningProfileOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) GetSigningProfileRequest(i0 *signer.GetSigningProfileInput) (r0 *request.Request, r1 *signer.GetSigningProfileOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSigningProfileRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*signer.GetSigningProfileOutput)
	return
}

func (d *SignerDouble) GetSigningProfileWithContext(i0 context.Context, i1 *signer.GetSigningProfileInput, i2 ...request.Option) (r0 *signer.GetSigningProfileOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSigningProfileWithContext", i0, i1, i2)
	r0, _ = returns[0].(*signer.GetSigningProfileOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) fakeGetSigningProfileWithContext(ctx context.Context, in *signer.GetSigningProfileInput, _ ...request.Option) (*signer.GetSigningProfileOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetSigningProfile cancelled", ctx.Err())
	default:
		return d.GetSigningProfile(in)
	}
}

func (d *SignerDouble) ListSigningJobs(i0 *signer.ListSigningJobsInput) (r0 *signer.ListSigningJobsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListSigningJobs", i0)
	r0, _ = returns[0].(*signer.ListSigningJobsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) ListSigningJobsPages(i0 *signer.ListSigningJobsInput, i1 func(*signer.ListSigningJobsOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListSigningJobsPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *SignerDouble) fakeListSigningJobsPages(in *signer.ListSigningJobsInput, pager func(*signer.ListSigningJobsOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListSigningJobs", paginators, in, pager)
}

func (d *SignerDouble) ListSigningJobsPagesWithContext(i0 context.Context, i1 *signer.ListSigningJobsInput, i2 func(*signer.ListSigningJobsOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListSigningJobsPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *SignerDouble) fakeListSigningJobsPagesWithContext(ctx context.Context, in *signer.ListSigningJobsInput, pager func(*signer.ListSigningJobsOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListSigningJobs", paginators, ctx, in, pager, options...)
}

func (d *SignerDouble) ListSigningJobsRequest(i0 *signer.ListSigningJobsInput) (r0 *request.Request, r1 *signer.ListSigningJobsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListSigningJobsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*signer.ListSigningJobsOutput)
	return
}

func (d *SignerDouble) ListSigningJobsWithContext(i0 context.Context, i1 *signer.ListSigningJobsInput, i2 ...request.Option) (r0 *signer.ListSigningJobsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListSigningJobsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*signer.ListSigningJobsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) fakeListSigningJobsWithContext(ctx context.Context, in *signer.ListSigningJobsInput, _ ...request.Option) (*signer.ListSigningJobsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListSigningJobs cancelled", ctx.Err())
	default:
		return d.ListSigningJobs(in)
	}
}

func (d *SignerDouble) ListSigningPlatforms(i0 *signer.ListSigningPlatformsInput) (r0 *signer.ListSigningPlatformsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListSigningPlatforms", i0)
	r0, _ = returns[0].(*signer.ListSigningPlatformsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) ListSigningPlatformsPages(i0 *signer.ListSigningPlatformsInput, i1 func(*signer.ListSigningPlatformsOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListSigningPlatformsPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *SignerDouble) fakeListSigningPlatformsPages(in *signer.ListSigningPlatformsInput, pager func(*signer.ListSigningPlatformsOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListSigningPlatforms", paginators, in, pager)
}

func (d *SignerDouble) ListSigningPlatformsPagesWithContext(i0 context.Context, i1 *signer.ListSigningPlatformsInput, i2 func(*signer.ListSigningPlatformsOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListSigningPlatformsPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *SignerDouble) fakeListSigningPlatformsPagesWithContext(ctx context.Context, in *signer.ListSigningPlatformsInput, pager func(*signer.ListSigningPlatformsOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListSigningPlatforms", paginators, ctx, in, pager, options...)
}

func (d *SignerDouble) ListSigningPlatformsRequest(i0 *signer.ListSigningPlatformsInput) (r0 *request.Request, r1 *signer.ListSigningPlatformsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListSigningPlatformsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*signer.ListSigningPlatformsOutput)
	return
}

func (d *SignerDouble) ListSigningPlatformsWithContext(i0 context.Context, i1 *signer.ListSigningPlatformsInput, i2 ...request.Option) (r0 *signer.ListSigningPlatformsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListSigningPlatformsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*signer.ListSigningPlatformsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) fakeListSigningPlatformsWithContext(ctx context.Context, in *signer.ListSigningPlatformsInput, _ ...request.Option) (*signer.ListSigningPlatformsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListSigningPlatforms cancelled", ctx.Err())
	default:
		return d.ListSigningPlatforms(in)
	}
}

func (d *SignerDouble) ListSigningProfiles(i0 *signer.ListSigningProfilesInput) (r0 *signer.ListSigningProfilesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListSigningProfiles", i0)
	r0, _ = returns[0].(*signer.ListSigningProfilesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) ListSigningProfilesPages(i0 *signer.ListSigningProfilesInput, i1 func(*signer.ListSigningProfilesOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListSigningProfilesPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *SignerDouble) fakeListSigningProfilesPages(in *signer.ListSigningProfilesInput, pager func(*signer.ListSigningProfilesOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListSigningProfiles", paginators, in, pager)
}

func (d *SignerDouble) ListSigningProfilesPagesWithContext(i0 context.Context, i1 *signer.ListSigningProfilesInput, i2 func(*signer.ListSigningProfilesOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListSigningProfilesPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *SignerDouble) fakeListSigningProfilesPagesWithContext(ctx context.Context, in *signer.ListSigningProfilesInput, pager func(*signer.ListSigningProfilesOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListSigningProfiles", paginators, ctx, in, pager, options...)
}

func (d *SignerDouble) ListSigningProfilesRequest(i0 *signer.ListSigningProfilesInput) (r0 *request.Request, r1 *signer.ListSigningProfilesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListSigningProfilesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*signer.ListSigningProfilesOutput)
	return
}

func (d *SignerDouble) ListSigningProfilesWithContext(i0 context.Context, i1 *signer.ListSigningProfilesInput, i2 ...request.Option) (r0 *signer.ListSigningProfilesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListSigningProfilesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*signer.ListSigningProfilesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) fakeListSigningProfilesWithContext(ctx context.Context, in *signer.ListSigningProfilesInput, _ ...request.Option) (*signer.ListSigningProfilesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListSigningProfiles cancelled", ctx.Err())
	default:
		return d.ListSigningProfiles(in)
	}
}

func (d *SignerDouble) ListTagsForResource(i0 *signer.ListTagsForResourceInput) (r0 *signer.ListTagsForResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForResource", i0)
	r0, _ = returns[0].(*signer.ListTagsForResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) ListTagsForResourceRequest(i0 *signer.ListTagsForResourceInput) (r0 *request.Request, r1 *signer.ListTagsForResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*signer.ListTagsForResourceOutput)
	return
}

func (d *SignerDouble) ListTagsForResourceWithContext(i0 context.Context, i1 *signer.ListTagsForResourceInput, i2 ...request.Option) (r0 *signer.ListTagsForResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*signer.ListTagsForResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) fakeListTagsForResourceWithContext(ctx context.Context, in *signer.ListTagsForResourceInput, _ ...request.Option) (*signer.ListTagsForResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListTagsForResource cancelled", ctx.Err())
	default:
		return d.ListTagsForResource(in)
	}
}

func (d *SignerDouble) PutSigningProfile(i0 *signer.PutSigningProfileInput) (r0 *signer.PutSigningProfileOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutSigningProfile", i0)
	r0, _ = returns[0].(*signer.PutSigningProfileOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) PutSigningProfileRequest(i0 *signer.PutSigningProfileInput) (r0 *request.Request, r1 *signer.PutSigningProfileOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutSigningProfileRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*signer.PutSigningProfileOutput)
	return
}

func (d *SignerDouble) PutSigningProfileWithContext(i0 context.Context, i1 *signer.PutSigningProfileInput, i2 ...request.Option) (r0 *signer.PutSigningProfileOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutSigningProfileWithContext", i0, i1, i2)
	r0, _ = returns[0].(*signer.PutSigningProfileOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) fakePutSigningProfileWithContext(ctx context.Context, in *signer.PutSigningProfileInput, _ ...request.Option) (*signer.PutSigningProfileOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "PutSigningProfile cancelled", ctx.Err())
	default:
		return d.PutSigningProfile(in)
	}
}

func (d *SignerDouble) StartSigningJob(i0 *signer.StartSigningJobInput) (r0 *signer.StartSigningJobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartSigningJob", i0)
	r0, _ = returns[0].(*signer.StartSigningJobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) StartSigningJobRequest(i0 *signer.StartSigningJobInput) (r0 *request.Request, r1 *signer.StartSigningJobOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartSigningJobRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*signer.StartSigningJobOutput)
	return
}

func (d *SignerDouble) StartSigningJobWithContext(i0 context.Context, i1 *signer.StartSigningJobInput, i2 ...request.Option) (r0 *signer.StartSigningJobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartSigningJobWithContext", i0, i1, i2)
	r0, _ = returns[0].(*signer.StartSigningJobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) fakeStartSigningJobWithContext(ctx context.Context, in *signer.StartSigningJobInput, _ ...request.Option) (*signer.StartSigningJobOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StartSigningJob cancelled", ctx.Err())
	default:
		return d.StartSigningJob(in)
	}
}

func (d *SignerDouble) TagResource(i0 *signer.TagResourceInput) (r0 *signer.TagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResource", i0)
	r0, _ = returns[0].(*signer.TagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) TagResourceRequest(i0 *signer.TagResourceInput) (r0 *request.Request, r1 *signer.TagResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*signer.TagResourceOutput)
	return
}

func (d *SignerDouble) TagResourceWithContext(i0 context.Context, i1 *signer.TagResourceInput, i2 ...request.Option) (r0 *signer.TagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*signer.TagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) fakeTagResourceWithContext(ctx context.Context, in *signer.TagResourceInput, _ ...request.Option) (*signer.TagResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "TagResource cancelled", ctx.Err())
	default:
		return d.TagResource(in)
	}
}

func (d *SignerDouble) UntagResource(i0 *signer.UntagResourceInput) (r0 *signer.UntagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResource", i0)
	r0, _ = returns[0].(*signer.UntagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) UntagResourceRequest(i0 *signer.UntagResourceInput) (r0 *request.Request, r1 *signer.UntagResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*signer.UntagResourceOutput)
	return
}

func (d *SignerDouble) UntagResourceWithContext(i0 context.Context, i1 *signer.UntagResourceInput, i2 ...request.Option) (r0 *signer.UntagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*signer.UntagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SignerDouble) fakeUntagResourceWithContext(ctx context.Context, in *signer.UntagResourceInput, _ ...request.Option) (*signer.UntagResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UntagResource cancelled", ctx.Err())
	default:
		return d.UntagResource(in)
	}
}

func (d *SignerDouble) WaitUntilSuccessfulSigningJob(i0 *signer.DescribeSigningJobInput) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("WaitUntilSuccessfulSigningJob", i0)
	r0, _ = returns[0].(error)
	return
}

func (d *SignerDouble) fakeWaitUntilSuccessfulSigningJob(in *signer.DescribeSigningJobInput) error {
	return d.WaitUntil("SuccessfulSigningJob", waiters, in)
}

func (d *SignerDouble) WaitUntilSuccessfulSigningJobWithContext(i0 context.Context, i1 *signer.DescribeSigningJobInput, i2 ...request.WaiterOption) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("WaitUntilSuccessfulSigningJobWithContext", i0, i1, i2)
	r0, _ = returns[0].(error)
	return
}

func (d *SignerDouble) fakeWaitUntilSuccessfulSigningJobWithContext(ctx context.Context, in *signer.DescribeSigningJobInput, waitOption ...request.WaiterOption) error {
	return d.WaitWithContext("SuccessfulSigningJob", waiters, ctx, in, waitOption...)
}