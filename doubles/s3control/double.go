// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:14:47+11:00

// Package s3controldouble provides a TestDouble implementation of s3controliface.S3ControlAPI
package s3controldouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3control"
	"github.com/aws/aws-sdk-go/service/s3control/s3controliface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// S3ControlDouble is TestDouble for s3controliface.S3ControlAPI
type S3ControlDouble struct {
	s3controliface.S3ControlAPI
	*awsdouble.AWSTestDouble
}

// Constructor for S3ControlDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *S3ControlDouble {
	result := &S3ControlDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*s3controliface.S3ControlAPI)(nil), configurators...)
	return result
}

func (d *S3ControlDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *S3ControlDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "CreateAccessPointWithContext":
		return m.Fake(d.fakeCreateAccessPointWithContext)

	case "CreateJobWithContext":
		return m.Fake(d.fakeCreateJobWithContext)

	case "DeleteAccessPointPolicyWithContext":
		return m.Fake(d.fakeDeleteAccessPointPolicyWithContext)

	case "DeleteAccessPointWithContext":
		return m.Fake(d.fakeDeleteAccessPointWithContext)

	case "DeletePublicAccessBlockWithContext":
		return m.Fake(d.fakeDeletePublicAccessBlockWithContext)

	case "DescribeJobWithContext":
		return m.Fake(d.fakeDescribeJobWithContext)

	case "GetAccessPointPolicyStatusWithContext":
		return m.Fake(d.fakeGetAccessPointPolicyStatusWithContext)

	case "GetAccessPointPolicyWithContext":
		return m.Fake(d.fakeGetAccessPointPolicyWithContext)

	case "GetAccessPointWithContext":
		return m.Fake(d.fakeGetAccessPointWithContext)

	case "GetPublicAccessBlockWithContext":
		return m.Fake(d.fakeGetPublicAccessBlockWithContext)

	case "ListAccessPointsPages":
		return m.Fake(d.fakeListAccessPointsPages)

	case "ListAccessPointsPagesWithContext":
		return m.Fake(d.fakeListAccessPointsPagesWithContext)

	case "ListAccessPointsWithContext":
		return m.Fake(d.fakeListAccessPointsWithContext)

	case "ListJobsPages":
		return m.Fake(d.fakeListJobsPages)

	case "ListJobsPagesWithContext":
		return m.Fake(d.fakeListJobsPagesWithContext)

	case "ListJobsWithContext":
		return m.Fake(d.fakeListJobsWithContext)

	case "PutAccessPointPolicyWithContext":
		return m.Fake(d.fakePutAccessPointPolicyWithContext)

	case "PutPublicAccessBlockWithContext":
		return m.Fake(d.fakePutPublicAccessBlockWithContext)

	case "UpdateJobPriorityWithContext":
		return m.Fake(d.fakeUpdateJobPriorityWithContext)

	case "UpdateJobStatusWithContext":
		return m.Fake(d.fakeUpdateJobStatusWithContext)

	default:
		return nil
	}
}

func (d *S3ControlDouble) CreateAccessPoint(i0 *s3control.CreateAccessPointInput) (r0 *s3control.CreateAccessPointOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateAccessPoint", i0)
	r0, _ = returns[0].(*s3control.CreateAccessPointOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) CreateAccessPointRequest(i0 *s3control.CreateAccessPointInput) (r0 *request.Request, r1 *s3control.CreateAccessPointOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateAccessPointRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*s3control.CreateAccessPointOutput)
	return
}

func (d *S3ControlDouble) CreateAccessPointWithContext(i0 context.Context, i1 *s3control.CreateAccessPointInput, i2 ...request.Option) (r0 *s3control.CreateAccessPointOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateAccessPointWithContext", i0, i1, i2)
	r0, _ = returns[0].(*s3control.CreateAccessPointOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) fakeCreateAccessPointWithContext(ctx context.Context, in *s3control.CreateAccessPointInput, _ ...request.Option) (*s3control.CreateAccessPointOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateAccessPoint cancelled", ctx.Err())
	default:
		return d.CreateAccessPoint(in)
	}
}

func (d *S3ControlDouble) CreateJob(i0 *s3control.CreateJobInput) (r0 *s3control.CreateJobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateJob", i0)
	r0, _ = returns[0].(*s3control.CreateJobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) CreateJobRequest(i0 *s3control.CreateJobInput) (r0 *request.Request, r1 *s3control.CreateJobOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateJobRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*s3control.CreateJobOutput)
	return
}

func (d *S3ControlDouble) CreateJobWithContext(i0 context.Context, i1 *s3control.CreateJobInput, i2 ...request.Option) (r0 *s3control.CreateJobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateJobWithContext", i0, i1, i2)
	r0, _ = returns[0].(*s3control.CreateJobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) fakeCreateJobWithContext(ctx context.Context, in *s3control.CreateJobInput, _ ...request.Option) (*s3control.CreateJobOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateJob cancelled", ctx.Err())
	default:
		return d.CreateJob(in)
	}
}

func (d *S3ControlDouble) DeleteAccessPoint(i0 *s3control.DeleteAccessPointInput) (r0 *s3control.DeleteAccessPointOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteAccessPoint", i0)
	r0, _ = returns[0].(*s3control.DeleteAccessPointOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) DeleteAccessPointPolicy(i0 *s3control.DeleteAccessPointPolicyInput) (r0 *s3control.DeleteAccessPointPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteAccessPointPolicy", i0)
	r0, _ = returns[0].(*s3control.DeleteAccessPointPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) DeleteAccessPointPolicyRequest(i0 *s3control.DeleteAccessPointPolicyInput) (r0 *request.Request, r1 *s3control.DeleteAccessPointPolicyOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteAccessPointPolicyRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*s3control.DeleteAccessPointPolicyOutput)
	return
}

func (d *S3ControlDouble) DeleteAccessPointPolicyWithContext(i0 context.Context, i1 *s3control.DeleteAccessPointPolicyInput, i2 ...request.Option) (r0 *s3control.DeleteAccessPointPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteAccessPointPolicyWithContext", i0, i1, i2)
	r0, _ = returns[0].(*s3control.DeleteAccessPointPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) fakeDeleteAccessPointPolicyWithContext(ctx context.Context, in *s3control.DeleteAccessPointPolicyInput, _ ...request.Option) (*s3control.DeleteAccessPointPolicyOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteAccessPointPolicy cancelled", ctx.Err())
	default:
		return d.DeleteAccessPointPolicy(in)
	}
}

func (d *S3ControlDouble) DeleteAccessPointRequest(i0 *s3control.DeleteAccessPointInput) (r0 *request.Request, r1 *s3control.DeleteAccessPointOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteAccessPointRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*s3control.DeleteAccessPointOutput)
	return
}

func (d *S3ControlDouble) DeleteAccessPointWithContext(i0 context.Context, i1 *s3control.DeleteAccessPointInput, i2 ...request.Option) (r0 *s3control.DeleteAccessPointOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteAccessPointWithContext", i0, i1, i2)
	r0, _ = returns[0].(*s3control.DeleteAccessPointOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) fakeDeleteAccessPointWithContext(ctx context.Context, in *s3control.DeleteAccessPointInput, _ ...request.Option) (*s3control.DeleteAccessPointOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteAccessPoint cancelled", ctx.Err())
	default:
		return d.DeleteAccessPoint(in)
	}
}

func (d *S3ControlDouble) DeletePublicAccessBlock(i0 *s3control.DeletePublicAccessBlockInput) (r0 *s3control.DeletePublicAccessBlockOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeletePublicAccessBlock", i0)
	r0, _ = returns[0].(*s3control.DeletePublicAccessBlockOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) DeletePublicAccessBlockRequest(i0 *s3control.DeletePublicAccessBlockInput) (r0 *request.Request, r1 *s3control.DeletePublicAccessBlockOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeletePublicAccessBlockRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*s3control.DeletePublicAccessBlockOutput)
	return
}

func (d *S3ControlDouble) DeletePublicAccessBlockWithContext(i0 context.Context, i1 *s3control.DeletePublicAccessBlockInput, i2 ...request.Option) (r0 *s3control.DeletePublicAccessBlockOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeletePublicAccessBlockWithContext", i0, i1, i2)
	r0, _ = returns[0].(*s3control.DeletePublicAccessBlockOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) fakeDeletePublicAccessBlockWithContext(ctx context.Context, in *s3control.DeletePublicAccessBlockInput, _ ...request.Option) (*s3control.DeletePublicAccessBlockOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeletePublicAccessBlock cancelled", ctx.Err())
	default:
		return d.DeletePublicAccessBlock(in)
	}
}

func (d *S3ControlDouble) DescribeJob(i0 *s3control.DescribeJobInput) (r0 *s3control.DescribeJobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeJob", i0)
	r0, _ = returns[0].(*s3control.DescribeJobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) DescribeJobRequest(i0 *s3control.DescribeJobInput) (r0 *request.Request, r1 *s3control.DescribeJobOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeJobRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*s3control.DescribeJobOutput)
	return
}

func (d *S3ControlDouble) DescribeJobWithContext(i0 context.Context, i1 *s3control.DescribeJobInput, i2 ...request.Option) (r0 *s3control.DescribeJobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeJobWithContext", i0, i1, i2)
	r0, _ = returns[0].(*s3control.DescribeJobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) fakeDescribeJobWithContext(ctx context.Context, in *s3control.DescribeJobInput, _ ...request.Option) (*s3control.DescribeJobOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeJob cancelled", ctx.Err())
	default:
		return d.DescribeJob(in)
	}
}

func (d *S3ControlDouble) GetAccessPoint(i0 *s3control.GetAccessPointInput) (r0 *s3control.GetAccessPointOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetAccessPoint", i0)
	r0, _ = returns[0].(*s3control.GetAccessPointOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) GetAccessPointPolicy(i0 *s3control.GetAccessPointPolicyInput) (r0 *s3control.GetAccessPointPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetAccessPointPolicy", i0)
	r0, _ = returns[0].(*s3control.GetAccessPointPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) GetAccessPointPolicyRequest(i0 *s3control.GetAccessPointPolicyInput) (r0 *request.Request, r1 *s3control.GetAccessPointPolicyOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetAccessPointPolicyRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*s3control.GetAccessPointPolicyOutput)
	return
}

func (d *S3ControlDouble) GetAccessPointPolicyStatus(i0 *s3control.GetAccessPointPolicyStatusInput) (r0 *s3control.GetAccessPointPolicyStatusOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetAccessPointPolicyStatus", i0)
	r0, _ = returns[0].(*s3control.GetAccessPointPolicyStatusOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) GetAccessPointPolicyStatusRequest(i0 *s3control.GetAccessPointPolicyStatusInput) (r0 *request.Request, r1 *s3control.GetAccessPointPolicyStatusOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetAccessPointPolicyStatusRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*s3control.GetAccessPointPolicyStatusOutput)
	return
}

func (d *S3ControlDouble) GetAccessPointPolicyStatusWithContext(i0 context.Context, i1 *s3control.GetAccessPointPolicyStatusInput, i2 ...request.Option) (r0 *s3control.GetAccessPointPolicyStatusOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetAccessPointPolicyStatusWithContext", i0, i1, i2)
	r0, _ = returns[0].(*s3control.GetAccessPointPolicyStatusOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) fakeGetAccessPointPolicyStatusWithContext(ctx context.Context, in *s3control.GetAccessPointPolicyStatusInput, _ ...request.Option) (*s3control.GetAccessPointPolicyStatusOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetAccessPointPolicyStatus cancelled", ctx.Err())
	default:
		return d.GetAccessPointPolicyStatus(in)
	}
}

func (d *S3ControlDouble) GetAccessPointPolicyWithContext(i0 context.Context, i1 *s3control.GetAccessPointPolicyInput, i2 ...request.Option) (r0 *s3control.GetAccessPointPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetAccessPointPolicyWithContext", i0, i1, i2)
	r0, _ = returns[0].(*s3control.GetAccessPointPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) fakeGetAccessPointPolicyWithContext(ctx context.Context, in *s3control.GetAccessPointPolicyInput, _ ...request.Option) (*s3control.GetAccessPointPolicyOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetAccessPointPolicy cancelled", ctx.Err())
	default:
		return d.GetAccessPointPolicy(in)
	}
}

func (d *S3ControlDouble) GetAccessPointRequest(i0 *s3control.GetAccessPointInput) (r0 *request.Request, r1 *s3control.GetAccessPointOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetAccessPointRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*s3control.GetAccessPointOutput)
	return
}

func (d *S3ControlDouble) GetAccessPointWithContext(i0 context.Context, i1 *s3control.GetAccessPointInput, i2 ...request.Option) (r0 *s3control.GetAccessPointOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetAccessPointWithContext", i0, i1, i2)
	r0, _ = returns[0].(*s3control.GetAccessPointOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) fakeGetAccessPointWithContext(ctx context.Context, in *s3control.GetAccessPointInput, _ ...request.Option) (*s3control.GetAccessPointOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetAccessPoint cancelled", ctx.Err())
	default:
		return d.GetAccessPoint(in)
	}
}

func (d *S3ControlDouble) GetPublicAccessBlock(i0 *s3control.GetPublicAccessBlockInput) (r0 *s3control.GetPublicAccessBlockOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetPublicAccessBlock", i0)
	r0, _ = returns[0].(*s3control.GetPublicAccessBlockOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) GetPublicAccessBlockRequest(i0 *s3control.GetPublicAccessBlockInput) (r0 *request.Request, r1 *s3control.GetPublicAccessBlockOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetPublicAccessBlockRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*s3control.GetPublicAccessBlockOutput)
	return
}

func (d *S3ControlDouble) GetPublicAccessBlockWithContext(i0 context.Context, i1 *s3control.GetPublicAccessBlockInput, i2 ...request.Option) (r0 *s3control.GetPublicAccessBlockOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetPublicAccessBlockWithContext", i0, i1, i2)
	r0, _ = returns[0].(*s3control.GetPublicAccessBlockOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) fakeGetPublicAccessBlockWithContext(ctx context.Context, in *s3control.GetPublicAccessBlockInput, _ ...request.Option) (*s3control.GetPublicAccessBlockOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetPublicAccessBlock cancelled", ctx.Err())
	default:
		return d.GetPublicAccessBlock(in)
	}
}

func (d *S3ControlDouble) ListAccessPoints(i0 *s3control.ListAccessPointsInput) (r0 *s3control.ListAccessPointsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListAccessPoints", i0)
	r0, _ = returns[0].(*s3control.ListAccessPointsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) ListAccessPointsPages(i0 *s3control.ListAccessPointsInput, i1 func(*s3control.ListAccessPointsOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListAccessPointsPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *S3ControlDouble) fakeListAccessPointsPages(in *s3control.ListAccessPointsInput, pager func(*s3control.ListAccessPointsOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListAccessPoints", paginators, in, pager)
}

func (d *S3ControlDouble) ListAccessPointsPagesWithContext(i0 context.Context, i1 *s3control.ListAccessPointsInput, i2 func(*s3control.ListAccessPointsOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListAccessPointsPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *S3ControlDouble) fakeListAccessPointsPagesWithContext(ctx context.Context, in *s3control.ListAccessPointsInput, pager func(*s3control.ListAccessPointsOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListAccessPoints", paginators, ctx, in, pager, options...)
}

func (d *S3ControlDouble) ListAccessPointsRequest(i0 *s3control.ListAccessPointsInput) (r0 *request.Request, r1 *s3control.ListAccessPointsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListAccessPointsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*s3control.ListAccessPointsOutput)
	return
}

func (d *S3ControlDouble) ListAccessPointsWithContext(i0 context.Context, i1 *s3control.ListAccessPointsInput, i2 ...request.Option) (r0 *s3control.ListAccessPointsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListAccessPointsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*s3control.ListAccessPointsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) fakeListAccessPointsWithContext(ctx context.Context, in *s3control.ListAccessPointsInput, _ ...request.Option) (*s3control.ListAccessPointsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListAccessPoints cancelled", ctx.Err())
	default:
		return d.ListAccessPoints(in)
	}
}

func (d *S3ControlDouble) ListJobs(i0 *s3control.ListJobsInput) (r0 *s3control.ListJobsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListJobs", i0)
	r0, _ = returns[0].(*s3control.ListJobsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) ListJobsPages(i0 *s3control.ListJobsInput, i1 func(*s3control.ListJobsOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListJobsPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *S3ControlDouble) fakeListJobsPages(in *s3control.ListJobsInput, pager func(*s3control.ListJobsOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListJobs", paginators, in, pager)
}

func (d *S3ControlDouble) ListJobsPagesWithContext(i0 context.Context, i1 *s3control.ListJobsInput, i2 func(*s3control.ListJobsOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListJobsPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *S3ControlDouble) fakeListJobsPagesWithContext(ctx context.Context, in *s3control.ListJobsInput, pager func(*s3control.ListJobsOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListJobs", paginators, ctx, in, pager, options...)
}

func (d *S3ControlDouble) ListJobsRequest(i0 *s3control.ListJobsInput) (r0 *request.Request, r1 *s3control.ListJobsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListJobsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*s3control.ListJobsOutput)
	return
}

func (d *S3ControlDouble) ListJobsWithContext(i0 context.Context, i1 *s3control.ListJobsInput, i2 ...request.Option) (r0 *s3control.ListJobsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListJobsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*s3control.ListJobsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) fakeListJobsWithContext(ctx context.Context, in *s3control.ListJobsInput, _ ...request.Option) (*s3control.ListJobsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListJobs cancelled", ctx.Err())
	default:
		return d.ListJobs(in)
	}
}

func (d *S3ControlDouble) PutAccessPointPolicy(i0 *s3control.PutAccessPointPolicyInput) (r0 *s3control.PutAccessPointPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutAccessPointPolicy", i0)
	r0, _ = returns[0].(*s3control.PutAccessPointPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) PutAccessPointPolicyRequest(i0 *s3control.PutAccessPointPolicyInput) (r0 *request.Request, r1 *s3control.PutAccessPointPolicyOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutAccessPointPolicyRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*s3control.PutAccessPointPolicyOutput)
	return
}

func (d *S3ControlDouble) PutAccessPointPolicyWithContext(i0 context.Context, i1 *s3control.PutAccessPointPolicyInput, i2 ...request.Option) (r0 *s3control.PutAccessPointPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutAccessPointPolicyWithContext", i0, i1, i2)
	r0, _ = returns[0].(*s3control.PutAccessPointPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) fakePutAccessPointPolicyWithContext(ctx context.Context, in *s3control.PutAccessPointPolicyInput, _ ...request.Option) (*s3control.PutAccessPointPolicyOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "PutAccessPointPolicy cancelled", ctx.Err())
	default:
		return d.PutAccessPointPolicy(in)
	}
}

func (d *S3ControlDouble) PutPublicAccessBlock(i0 *s3control.PutPublicAccessBlockInput) (r0 *s3control.PutPublicAccessBlockOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutPublicAccessBlock", i0)
	r0, _ = returns[0].(*s3control.PutPublicAccessBlockOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) PutPublicAccessBlockRequest(i0 *s3control.PutPublicAccessBlockInput) (r0 *request.Request, r1 *s3control.PutPublicAccessBlockOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutPublicAccessBlockRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*s3control.PutPublicAccessBlockOutput)
	return
}

func (d *S3ControlDouble) PutPublicAccessBlockWithContext(i0 context.Context, i1 *s3control.PutPublicAccessBlockInput, i2 ...request.Option) (r0 *s3control.PutPublicAccessBlockOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutPublicAccessBlockWithContext", i0, i1, i2)
	r0, _ = returns[0].(*s3control.PutPublicAccessBlockOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) fakePutPublicAccessBlockWithContext(ctx context.Context, in *s3control.PutPublicAccessBlockInput, _ ...request.Option) (*s3control.PutPublicAccessBlockOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "PutPublicAccessBlock cancelled", ctx.Err())
	default:
		return d.PutPublicAccessBlock(in)
	}
}

func (d *S3ControlDouble) UpdateJobPriority(i0 *s3control.UpdateJobPriorityInput) (r0 *s3control.UpdateJobPriorityOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateJobPriority", i0)
	r0, _ = returns[0].(*s3control.UpdateJobPriorityOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) UpdateJobPriorityRequest(i0 *s3control.UpdateJobPriorityInput) (r0 *request.Request, r1 *s3control.UpdateJobPriorityOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateJobPriorityRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*s3control.UpdateJobPriorityOutput)
	return
}

func (d *S3ControlDouble) UpdateJobPriorityWithContext(i0 context.Context, i1 *s3control.UpdateJobPriorityInput, i2 ...request.Option) (r0 *s3control.UpdateJobPriorityOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateJobPriorityWithContext", i0, i1, i2)
	r0, _ = returns[0].(*s3control.UpdateJobPriorityOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) fakeUpdateJobPriorityWithContext(ctx context.Context, in *s3control.UpdateJobPriorityInput, _ ...request.Option) (*s3control.UpdateJobPriorityOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateJobPriority cancelled", ctx.Err())
	default:
		return d.UpdateJobPriority(in)
	}
}

func (d *S3ControlDouble) UpdateJobStatus(i0 *s3control.UpdateJobStatusInput) (r0 *s3control.UpdateJobStatusOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateJobStatus", i0)
	r0, _ = returns[0].(*s3control.UpdateJobStatusOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) UpdateJobStatusRequest(i0 *s3control.UpdateJobStatusInput) (r0 *request.Request, r1 *s3control.UpdateJobStatusOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateJobStatusRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*s3control.UpdateJobStatusOutput)
	return
}

func (d *S3ControlDouble) UpdateJobStatusWithContext(i0 context.Context, i1 *s3control.UpdateJobStatusInput, i2 ...request.Option) (r0 *s3control.UpdateJobStatusOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateJobStatusWithContext", i0, i1, i2)
	r0, _ = returns[0].(*s3control.UpdateJobStatusOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *S3ControlDouble) fakeUpdateJobStatusWithContext(ctx context.Context, in *s3control.UpdateJobStatusInput, _ ...request.Option) (*s3control.UpdateJobStatusOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateJobStatus cancelled", ctx.Err())
	default:
		return d.UpdateJobStatus(in)
	}
}