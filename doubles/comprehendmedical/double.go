// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:12:02+11:00

// Package comprehendmedicaldouble provides a TestDouble implementation of comprehendmedicaliface.ComprehendMedicalAPI
package comprehendmedicaldouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/comprehendmedical"
	"github.com/aws/aws-sdk-go/service/comprehendmedical/comprehendmedicaliface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// ComprehendMedicalDouble is TestDouble for comprehendmedicaliface.ComprehendMedicalAPI
type ComprehendMedicalDouble struct {
	comprehendmedicaliface.ComprehendMedicalAPI
	*awsdouble.AWSTestDouble
}

// Constructor for ComprehendMedicalDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *ComprehendMedicalDouble {
	result := &ComprehendMedicalDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*comprehendmedicaliface.ComprehendMedicalAPI)(nil), configurators...)
	return result
}

func (d *ComprehendMedicalDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *ComprehendMedicalDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "DescribeEntitiesDetectionV2JobWithContext":
		return m.Fake(d.fakeDescribeEntitiesDetectionV2JobWithContext)

	case "DescribePHIDetectionJobWithContext":
		return m.Fake(d.fakeDescribePHIDetectionJobWithContext)

	case "DetectEntitiesV2WithContext":
		return m.Fake(d.fakeDetectEntitiesV2WithContext)

	case "DetectEntitiesWithContext":
		return m.Fake(d.fakeDetectEntitiesWithContext)

	case "DetectPHIWithContext":
		return m.Fake(d.fakeDetectPHIWithContext)

	case "InferICD10CMWithContext":
		return m.Fake(d.fakeInferICD10CMWithContext)

	case "InferRxNormWithContext":
		return m.Fake(d.fakeInferRxNormWithContext)

	case "ListEntitiesDetectionV2JobsWithContext":
		return m.Fake(d.fakeListEntitiesDetectionV2JobsWithContext)

	case "ListPHIDetectionJobsWithContext":
		return m.Fake(d.fakeListPHIDetectionJobsWithContext)

	case "StartEntitiesDetectionV2JobWithContext":
		return m.Fake(d.fakeStartEntitiesDetectionV2JobWithContext)

	case "StartPHIDetectionJobWithContext":
		return m.Fake(d.fakeStartPHIDetectionJobWithContext)

	case "StopEntitiesDetectionV2JobWithContext":
		return m.Fake(d.fakeStopEntitiesDetectionV2JobWithContext)

	case "StopPHIDetectionJobWithContext":
		return m.Fake(d.fakeStopPHIDetectionJobWithContext)

	default:
		return nil
	}
}

func (d *ComprehendMedicalDouble) DescribeEntitiesDetectionV2Job(i0 *comprehendmedical.DescribeEntitiesDetectionV2JobInput) (r0 *comprehendmedical.DescribeEntitiesDetectionV2JobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeEntitiesDetectionV2Job", i0)
	r0, _ = returns[0].(*comprehendmedical.DescribeEntitiesDetectionV2JobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) DescribeEntitiesDetectionV2JobRequest(i0 *comprehendmedical.DescribeEntitiesDetectionV2JobInput) (r0 *request.Request, r1 *comprehendmedical.DescribeEntitiesDetectionV2JobOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeEntitiesDetectionV2JobRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*comprehendmedical.DescribeEntitiesDetectionV2JobOutput)
	return
}

func (d *ComprehendMedicalDouble) DescribeEntitiesDetectionV2JobWithContext(i0 context.Context, i1 *comprehendmedical.DescribeEntitiesDetectionV2JobInput, i2 ...request.Option) (r0 *comprehendmedical.DescribeEntitiesDetectionV2JobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeEntitiesDetectionV2JobWithContext", i0, i1, i2)
	r0, _ = returns[0].(*comprehendmedical.DescribeEntitiesDetectionV2JobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) fakeDescribeEntitiesDetectionV2JobWithContext(ctx context.Context, in *comprehendmedical.DescribeEntitiesDetectionV2JobInput, _ ...request.Option) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeEntitiesDetectionV2Job cancelled", ctx.Err())
	default:
		return d.DescribeEntitiesDetectionV2Job(in)
	}
}

func (d *ComprehendMedicalDouble) DescribePHIDetectionJob(i0 *comprehendmedical.DescribePHIDetectionJobInput) (r0 *comprehendmedical.DescribePHIDetectionJobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribePHIDetectionJob", i0)
	r0, _ = returns[0].(*comprehendmedical.DescribePHIDetectionJobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) DescribePHIDetectionJobRequest(i0 *comprehendmedical.DescribePHIDetectionJobInput) (r0 *request.Request, r1 *comprehendmedical.DescribePHIDetectionJobOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribePHIDetectionJobRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*comprehendmedical.DescribePHIDetectionJobOutput)
	return
}

func (d *ComprehendMedicalDouble) DescribePHIDetectionJobWithContext(i0 context.Context, i1 *comprehendmedical.DescribePHIDetectionJobInput, i2 ...request.Option) (r0 *comprehendmedical.DescribePHIDetectionJobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribePHIDetectionJobWithContext", i0, i1, i2)
	r0, _ = returns[0].(*comprehendmedical.DescribePHIDetectionJobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) fakeDescribePHIDetectionJobWithContext(ctx context.Context, in *comprehendmedical.DescribePHIDetectionJobInput, _ ...request.Option) (*comprehendmedical.DescribePHIDetectionJobOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribePHIDetectionJob cancelled", ctx.Err())
	default:
		return d.DescribePHIDetectionJob(in)
	}
}

func (d *ComprehendMedicalDouble) DetectEntities(i0 *comprehendmedical.DetectEntitiesInput) (r0 *comprehendmedical.DetectEntitiesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DetectEntities", i0)
	r0, _ = returns[0].(*comprehendmedical.DetectEntitiesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) DetectEntitiesRequest(i0 *comprehendmedical.DetectEntitiesInput) (r0 *request.Request, r1 *comprehendmedical.DetectEntitiesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DetectEntitiesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*comprehendmedical.DetectEntitiesOutput)
	return
}

func (d *ComprehendMedicalDouble) DetectEntitiesV2(i0 *comprehendmedical.DetectEntitiesV2Input) (r0 *comprehendmedical.DetectEntitiesV2Output, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DetectEntitiesV2", i0)
	r0, _ = returns[0].(*comprehendmedical.DetectEntitiesV2Output)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) DetectEntitiesV2Request(i0 *comprehendmedical.DetectEntitiesV2Input) (r0 *request.Request, r1 *comprehendmedical.DetectEntitiesV2Output) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DetectEntitiesV2Request", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*comprehendmedical.DetectEntitiesV2Output)
	return
}

func (d *ComprehendMedicalDouble) DetectEntitiesV2WithContext(i0 context.Context, i1 *comprehendmedical.DetectEntitiesV2Input, i2 ...request.Option) (r0 *comprehendmedical.DetectEntitiesV2Output, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DetectEntitiesV2WithContext", i0, i1, i2)
	r0, _ = returns[0].(*comprehendmedical.DetectEntitiesV2Output)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) fakeDetectEntitiesV2WithContext(ctx context.Context, in *comprehendmedical.DetectEntitiesV2Input, _ ...request.Option) (*comprehendmedical.DetectEntitiesV2Output, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DetectEntitiesV2 cancelled", ctx.Err())
	default:
		return d.DetectEntitiesV2(in)
	}
}

func (d *ComprehendMedicalDouble) DetectEntitiesWithContext(i0 context.Context, i1 *comprehendmedical.DetectEntitiesInput, i2 ...request.Option) (r0 *comprehendmedical.DetectEntitiesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DetectEntitiesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*comprehendmedical.DetectEntitiesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) fakeDetectEntitiesWithContext(ctx context.Context, in *comprehendmedical.DetectEntitiesInput, _ ...request.Option) (*comprehendmedical.DetectEntitiesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DetectEntities cancelled", ctx.Err())
	default:
		return d.DetectEntities(in)
	}
}

func (d *ComprehendMedicalDouble) DetectPHI(i0 *comprehendmedical.DetectPHIInput) (r0 *comprehendmedical.DetectPHIOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DetectPHI", i0)
	r0, _ = returns[0].(*comprehendmedical.DetectPHIOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) DetectPHIRequest(i0 *comprehendmedical.DetectPHIInput) (r0 *request.Request, r1 *comprehendmedical.DetectPHIOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DetectPHIRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*comprehendmedical.DetectPHIOutput)
	return
}

func (d *ComprehendMedicalDouble) DetectPHIWithContext(i0 context.Context, i1 *comprehendmedical.DetectPHIInput, i2 ...request.Option) (r0 *comprehendmedical.DetectPHIOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DetectPHIWithContext", i0, i1, i2)
	r0, _ = returns[0].(*comprehendmedical.DetectPHIOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) fakeDetectPHIWithContext(ctx context.Context, in *comprehendmedical.DetectPHIInput, _ ...request.Option) (*comprehendmedical.DetectPHIOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DetectPHI cancelled", ctx.Err())
	default:
		return d.DetectPHI(in)
	}
}

func (d *ComprehendMedicalDouble) InferICD10CM(i0 *comprehendmedical.InferICD10CMInput) (r0 *comprehendmedical.InferICD10CMOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("InferICD10CM", i0)
	r0, _ = returns[0].(*comprehendmedical.InferICD10CMOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) InferICD10CMRequest(i0 *comprehendmedical.InferICD10CMInput) (r0 *request.Request, r1 *comprehendmedical.InferICD10CMOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("InferICD10CMRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*comprehendmedical.InferICD10CMOutput)
	return
}

func (d *ComprehendMedicalDouble) InferICD10CMWithContext(i0 context.Context, i1 *comprehendmedical.InferICD10CMInput, i2 ...request.Option) (r0 *comprehendmedical.InferICD10CMOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("InferICD10CMWithContext", i0, i1, i2)
	r0, _ = returns[0].(*comprehendmedical.InferICD10CMOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) fakeInferICD10CMWithContext(ctx context.Context, in *comprehendmedical.InferICD10CMInput, _ ...request.Option) (*comprehendmedical.InferICD10CMOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "InferICD10CM cancelled", ctx.Err())
	default:
		return d.InferICD10CM(in)
	}
}

func (d *ComprehendMedicalDouble) InferRxNorm(i0 *comprehendmedical.InferRxNormInput) (r0 *comprehendmedical.InferRxNormOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("InferRxNorm", i0)
	r0, _ = returns[0].(*comprehendmedical.InferRxNormOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) InferRxNormRequest(i0 *comprehendmedical.InferRxNormInput) (r0 *request.Request, r1 *comprehendmedical.InferRxNormOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("InferRxNormRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*comprehendmedical.InferRxNormOutput)
	return
}

func (d *ComprehendMedicalDouble) InferRxNormWithContext(i0 context.Context, i1 *comprehendmedical.InferRxNormInput, i2 ...request.Option) (r0 *comprehendmedical.InferRxNormOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("InferRxNormWithContext", i0, i1, i2)
	r0, _ = returns[0].(*comprehendmedical.InferRxNormOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) fakeInferRxNormWithContext(ctx context.Context, in *comprehendmedical.InferRxNormInput, _ ...request.Option) (*comprehendmedical.InferRxNormOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "InferRxNorm cancelled", ctx.Err())
	default:
		return d.InferRxNorm(in)
	}
}

func (d *ComprehendMedicalDouble) ListEntitiesDetectionV2Jobs(i0 *comprehendmedical.ListEntitiesDetectionV2JobsInput) (r0 *comprehendmedical.ListEntitiesDetectionV2JobsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListEntitiesDetectionV2Jobs", i0)
	r0, _ = returns[0].(*comprehendmedical.ListEntitiesDetectionV2JobsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) ListEntitiesDetectionV2JobsRequest(i0 *comprehendmedical.ListEntitiesDetectionV2JobsInput) (r0 *request.Request, r1 *comprehendmedical.ListEntitiesDetectionV2JobsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListEntitiesDetectionV2JobsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*comprehendmedical.ListEntitiesDetectionV2JobsOutput)
	return
}

func (d *ComprehendMedicalDouble) ListEntitiesDetectionV2JobsWithContext(i0 context.Context, i1 *comprehendmedical.ListEntitiesDetectionV2JobsInput, i2 ...request.Option) (r0 *comprehendmedical.ListEntitiesDetectionV2JobsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListEntitiesDetectionV2JobsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*comprehendmedical.ListEntitiesDetectionV2JobsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) fakeListEntitiesDetectionV2JobsWithContext(ctx context.Context, in *comprehendmedical.ListEntitiesDetectionV2JobsInput, _ ...request.Option) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListEntitiesDetectionV2Jobs cancelled", ctx.Err())
	default:
		return d.ListEntitiesDetectionV2Jobs(in)
	}
}

func (d *ComprehendMedicalDouble) ListPHIDetectionJobs(i0 *comprehendmedical.ListPHIDetectionJobsInput) (r0 *comprehendmedical.ListPHIDetectionJobsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListPHIDetectionJobs", i0)
	r0, _ = returns[0].(*comprehendmedical.ListPHIDetectionJobsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) ListPHIDetectionJobsRequest(i0 *comprehendmedical.ListPHIDetectionJobsInput) (r0 *request.Request, r1 *comprehendmedical.ListPHIDetectionJobsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListPHIDetectionJobsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*comprehendmedical.ListPHIDetectionJobsOutput)
	return
}

func (d *ComprehendMedicalDouble) ListPHIDetectionJobsWithContext(i0 context.Context, i1 *comprehendmedical.ListPHIDetectionJobsInput, i2 ...request.Option) (r0 *comprehendmedical.ListPHIDetectionJobsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListPHIDetectionJobsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*comprehendmedical.ListPHIDetectionJobsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) fakeListPHIDetectionJobsWithContext(ctx context.Context, in *comprehendmedical.ListPHIDetectionJobsInput, _ ...request.Option) (*comprehendmedical.ListPHIDetectionJobsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListPHIDetectionJobs cancelled", ctx.Err())
	default:
		return d.ListPHIDetectionJobs(in)
	}
}

func (d *ComprehendMedicalDouble) StartEntitiesDetectionV2Job(i0 *comprehendmedical.StartEntitiesDetectionV2JobInput) (r0 *comprehendmedical.StartEntitiesDetectionV2JobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartEntitiesDetectionV2Job", i0)
	r0, _ = returns[0].(*comprehendmedical.StartEntitiesDetectionV2JobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) StartEntitiesDetectionV2JobRequest(i0 *comprehendmedical.StartEntitiesDetectionV2JobInput) (r0 *request.Request, r1 *comprehendmedical.StartEntitiesDetectionV2JobOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartEntitiesDetectionV2JobRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*comprehendmedical.StartEntitiesDetectionV2JobOutput)
	return
}

func (d *ComprehendMedicalDouble) StartEntitiesDetectionV2JobWithContext(i0 context.Context, i1 *comprehendmedical.StartEntitiesDetectionV2JobInput, i2 ...request.Option) (r0 *comprehendmedical.StartEntitiesDetectionV2JobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartEntitiesDetectionV2JobWithContext", i0, i1, i2)
	r0, _ = returns[0].(*comprehendmedical.StartEntitiesDetectionV2JobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) fakeStartEntitiesDetectionV2JobWithContext(ctx context.Context, in *comprehendmedical.StartEntitiesDetectionV2JobInput, _ ...request.Option) (*comprehendmedical.StartEntitiesDetectionV2JobOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StartEntitiesDetectionV2Job cancelled", ctx.Err())
	default:
		return d.StartEntitiesDetectionV2Job(in)
	}
}

func (d *ComprehendMedicalDouble) StartPHIDetectionJob(i0 *comprehendmedical.StartPHIDetectionJobInput) (r0 *comprehendmedical.StartPHIDetectionJobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartPHIDetectionJob", i0)
	r0, _ = returns[0].(*comprehendmedical.StartPHIDetectionJobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) StartPHIDetectionJobRequest(i0 *comprehendmedical.StartPHIDetectionJobInput) (r0 *request.Request, r1 *comprehendmedical.StartPHIDetectionJobOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartPHIDetectionJobRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*comprehendmedical.StartPHIDetectionJobOutput)
	return
}

func (d *ComprehendMedicalDouble) StartPHIDetectionJobWithContext(i0 context.Context, i1 *comprehendmedical.StartPHIDetectionJobInput, i2 ...request.Option) (r0 *comprehendmedical.StartPHIDetectionJobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartPHIDetectionJobWithContext", i0, i1, i2)
	r0, _ = returns[0].(*comprehendmedical.StartPHIDetectionJobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) fakeStartPHIDetectionJobWithContext(ctx context.Context, in *comprehendmedical.StartPHIDetectionJobInput, _ ...request.Option) (*comprehendmedical.StartPHIDetectionJobOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StartPHIDetectionJob cancelled", ctx.Err())
	default:
		return d.StartPHIDetectionJob(in)
	}
}

func (d *ComprehendMedicalDouble) StopEntitiesDetectionV2Job(i0 *comprehendmedical.StopEntitiesDetectionV2JobInput) (r0 *comprehendmedical.StopEntitiesDetectionV2JobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopEntitiesDetectionV2Job", i0)
	r0, _ = returns[0].(*comprehendmedical.StopEntitiesDetectionV2JobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) StopEntitiesDetectionV2JobRequest(i0 *comprehendmedical.StopEntitiesDetectionV2JobInput) (r0 *request.Request, r1 *comprehendmedical.StopEntitiesDetectionV2JobOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopEntitiesDetectionV2JobRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*comprehendmedical.StopEntitiesDetectionV2JobOutput)
	return
}

func (d *ComprehendMedicalDouble) StopEntitiesDetectionV2JobWithContext(i0 context.Context, i1 *comprehendmedical.StopEntitiesDetectionV2JobInput, i2 ...request.Option) (r0 *comprehendmedical.StopEntitiesDetectionV2JobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopEntitiesDetectionV2JobWithContext", i0, i1, i2)
	r0, _ = returns[0].(*comprehendmedical.StopEntitiesDetectionV2JobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) fakeStopEntitiesDetectionV2JobWithContext(ctx context.Context, in *comprehendmedical.StopEntitiesDetectionV2JobInput, _ ...request.Option) (*comprehendmedical.StopEntitiesDetectionV2JobOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StopEntitiesDetectionV2Job cancelled", ctx.Err())
	default:
		return d.StopEntitiesDetectionV2Job(in)
	}
}

func (d *ComprehendMedicalDouble) StopPHIDetectionJob(i0 *comprehendmedical.StopPHIDetectionJobInput) (r0 *comprehendmedical.StopPHIDetectionJobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopPHIDetectionJob", i0)
	r0, _ = returns[0].(*comprehendmedical.StopPHIDetectionJobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) StopPHIDetectionJobRequest(i0 *comprehendmedical.StopPHIDetectionJobInput) (r0 *request.Request, r1 *comprehendmedical.StopPHIDetectionJobOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopPHIDetectionJobRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*comprehendmedical.StopPHIDetectionJobOutput)
	return
}

func (d *ComprehendMedicalDouble) StopPHIDetectionJobWithContext(i0 context.Context, i1 *comprehendmedical.StopPHIDetectionJobInput, i2 ...request.Option) (r0 *comprehendmedical.StopPHIDetectionJobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopPHIDetectionJobWithContext", i0, i1, i2)
	r0, _ = returns[0].(*comprehendmedical.StopPHIDetectionJobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *ComprehendMedicalDouble) fakeStopPHIDetectionJobWithContext(ctx context.Context, in *comprehendmedical.StopPHIDetectionJobInput, _ ...request.Option) (*comprehendmedical.StopPHIDetectionJobOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StopPHIDetectionJob cancelled", ctx.Err())
	default:
		return d.StopPHIDetectionJob(in)
	}
}