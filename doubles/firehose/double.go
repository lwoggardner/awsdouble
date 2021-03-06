// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:12:48+11:00

// Package firehosedouble provides a TestDouble implementation of firehoseiface.FirehoseAPI
package firehosedouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/firehose/firehoseiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// FirehoseDouble is TestDouble for firehoseiface.FirehoseAPI
type FirehoseDouble struct {
	firehoseiface.FirehoseAPI
	*awsdouble.AWSTestDouble
}

// Constructor for FirehoseDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *FirehoseDouble {
	result := &FirehoseDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*firehoseiface.FirehoseAPI)(nil), configurators...)
	return result
}

func (d *FirehoseDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *FirehoseDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "CreateDeliveryStreamWithContext":
		return m.Fake(d.fakeCreateDeliveryStreamWithContext)

	case "DeleteDeliveryStreamWithContext":
		return m.Fake(d.fakeDeleteDeliveryStreamWithContext)

	case "DescribeDeliveryStreamWithContext":
		return m.Fake(d.fakeDescribeDeliveryStreamWithContext)

	case "ListDeliveryStreamsWithContext":
		return m.Fake(d.fakeListDeliveryStreamsWithContext)

	case "ListTagsForDeliveryStreamWithContext":
		return m.Fake(d.fakeListTagsForDeliveryStreamWithContext)

	case "PutRecordBatchWithContext":
		return m.Fake(d.fakePutRecordBatchWithContext)

	case "PutRecordWithContext":
		return m.Fake(d.fakePutRecordWithContext)

	case "StartDeliveryStreamEncryptionWithContext":
		return m.Fake(d.fakeStartDeliveryStreamEncryptionWithContext)

	case "StopDeliveryStreamEncryptionWithContext":
		return m.Fake(d.fakeStopDeliveryStreamEncryptionWithContext)

	case "TagDeliveryStreamWithContext":
		return m.Fake(d.fakeTagDeliveryStreamWithContext)

	case "UntagDeliveryStreamWithContext":
		return m.Fake(d.fakeUntagDeliveryStreamWithContext)

	case "UpdateDestinationWithContext":
		return m.Fake(d.fakeUpdateDestinationWithContext)

	default:
		return nil
	}
}

func (d *FirehoseDouble) CreateDeliveryStream(i0 *firehose.CreateDeliveryStreamInput) (r0 *firehose.CreateDeliveryStreamOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateDeliveryStream", i0)
	r0, _ = returns[0].(*firehose.CreateDeliveryStreamOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) CreateDeliveryStreamRequest(i0 *firehose.CreateDeliveryStreamInput) (r0 *request.Request, r1 *firehose.CreateDeliveryStreamOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateDeliveryStreamRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*firehose.CreateDeliveryStreamOutput)
	return
}

func (d *FirehoseDouble) CreateDeliveryStreamWithContext(i0 context.Context, i1 *firehose.CreateDeliveryStreamInput, i2 ...request.Option) (r0 *firehose.CreateDeliveryStreamOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateDeliveryStreamWithContext", i0, i1, i2)
	r0, _ = returns[0].(*firehose.CreateDeliveryStreamOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) fakeCreateDeliveryStreamWithContext(ctx context.Context, in *firehose.CreateDeliveryStreamInput, _ ...request.Option) (*firehose.CreateDeliveryStreamOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateDeliveryStream cancelled", ctx.Err())
	default:
		return d.CreateDeliveryStream(in)
	}
}

func (d *FirehoseDouble) DeleteDeliveryStream(i0 *firehose.DeleteDeliveryStreamInput) (r0 *firehose.DeleteDeliveryStreamOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteDeliveryStream", i0)
	r0, _ = returns[0].(*firehose.DeleteDeliveryStreamOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) DeleteDeliveryStreamRequest(i0 *firehose.DeleteDeliveryStreamInput) (r0 *request.Request, r1 *firehose.DeleteDeliveryStreamOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteDeliveryStreamRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*firehose.DeleteDeliveryStreamOutput)
	return
}

func (d *FirehoseDouble) DeleteDeliveryStreamWithContext(i0 context.Context, i1 *firehose.DeleteDeliveryStreamInput, i2 ...request.Option) (r0 *firehose.DeleteDeliveryStreamOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteDeliveryStreamWithContext", i0, i1, i2)
	r0, _ = returns[0].(*firehose.DeleteDeliveryStreamOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) fakeDeleteDeliveryStreamWithContext(ctx context.Context, in *firehose.DeleteDeliveryStreamInput, _ ...request.Option) (*firehose.DeleteDeliveryStreamOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteDeliveryStream cancelled", ctx.Err())
	default:
		return d.DeleteDeliveryStream(in)
	}
}

func (d *FirehoseDouble) DescribeDeliveryStream(i0 *firehose.DescribeDeliveryStreamInput) (r0 *firehose.DescribeDeliveryStreamOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeDeliveryStream", i0)
	r0, _ = returns[0].(*firehose.DescribeDeliveryStreamOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) DescribeDeliveryStreamRequest(i0 *firehose.DescribeDeliveryStreamInput) (r0 *request.Request, r1 *firehose.DescribeDeliveryStreamOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeDeliveryStreamRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*firehose.DescribeDeliveryStreamOutput)
	return
}

func (d *FirehoseDouble) DescribeDeliveryStreamWithContext(i0 context.Context, i1 *firehose.DescribeDeliveryStreamInput, i2 ...request.Option) (r0 *firehose.DescribeDeliveryStreamOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeDeliveryStreamWithContext", i0, i1, i2)
	r0, _ = returns[0].(*firehose.DescribeDeliveryStreamOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) fakeDescribeDeliveryStreamWithContext(ctx context.Context, in *firehose.DescribeDeliveryStreamInput, _ ...request.Option) (*firehose.DescribeDeliveryStreamOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeDeliveryStream cancelled", ctx.Err())
	default:
		return d.DescribeDeliveryStream(in)
	}
}

func (d *FirehoseDouble) ListDeliveryStreams(i0 *firehose.ListDeliveryStreamsInput) (r0 *firehose.ListDeliveryStreamsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListDeliveryStreams", i0)
	r0, _ = returns[0].(*firehose.ListDeliveryStreamsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) ListDeliveryStreamsRequest(i0 *firehose.ListDeliveryStreamsInput) (r0 *request.Request, r1 *firehose.ListDeliveryStreamsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListDeliveryStreamsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*firehose.ListDeliveryStreamsOutput)
	return
}

func (d *FirehoseDouble) ListDeliveryStreamsWithContext(i0 context.Context, i1 *firehose.ListDeliveryStreamsInput, i2 ...request.Option) (r0 *firehose.ListDeliveryStreamsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListDeliveryStreamsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*firehose.ListDeliveryStreamsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) fakeListDeliveryStreamsWithContext(ctx context.Context, in *firehose.ListDeliveryStreamsInput, _ ...request.Option) (*firehose.ListDeliveryStreamsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListDeliveryStreams cancelled", ctx.Err())
	default:
		return d.ListDeliveryStreams(in)
	}
}

func (d *FirehoseDouble) ListTagsForDeliveryStream(i0 *firehose.ListTagsForDeliveryStreamInput) (r0 *firehose.ListTagsForDeliveryStreamOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForDeliveryStream", i0)
	r0, _ = returns[0].(*firehose.ListTagsForDeliveryStreamOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) ListTagsForDeliveryStreamRequest(i0 *firehose.ListTagsForDeliveryStreamInput) (r0 *request.Request, r1 *firehose.ListTagsForDeliveryStreamOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForDeliveryStreamRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*firehose.ListTagsForDeliveryStreamOutput)
	return
}

func (d *FirehoseDouble) ListTagsForDeliveryStreamWithContext(i0 context.Context, i1 *firehose.ListTagsForDeliveryStreamInput, i2 ...request.Option) (r0 *firehose.ListTagsForDeliveryStreamOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForDeliveryStreamWithContext", i0, i1, i2)
	r0, _ = returns[0].(*firehose.ListTagsForDeliveryStreamOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) fakeListTagsForDeliveryStreamWithContext(ctx context.Context, in *firehose.ListTagsForDeliveryStreamInput, _ ...request.Option) (*firehose.ListTagsForDeliveryStreamOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListTagsForDeliveryStream cancelled", ctx.Err())
	default:
		return d.ListTagsForDeliveryStream(in)
	}
}

func (d *FirehoseDouble) PutRecord(i0 *firehose.PutRecordInput) (r0 *firehose.PutRecordOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutRecord", i0)
	r0, _ = returns[0].(*firehose.PutRecordOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) PutRecordBatch(i0 *firehose.PutRecordBatchInput) (r0 *firehose.PutRecordBatchOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutRecordBatch", i0)
	r0, _ = returns[0].(*firehose.PutRecordBatchOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) PutRecordBatchRequest(i0 *firehose.PutRecordBatchInput) (r0 *request.Request, r1 *firehose.PutRecordBatchOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutRecordBatchRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*firehose.PutRecordBatchOutput)
	return
}

func (d *FirehoseDouble) PutRecordBatchWithContext(i0 context.Context, i1 *firehose.PutRecordBatchInput, i2 ...request.Option) (r0 *firehose.PutRecordBatchOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutRecordBatchWithContext", i0, i1, i2)
	r0, _ = returns[0].(*firehose.PutRecordBatchOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) fakePutRecordBatchWithContext(ctx context.Context, in *firehose.PutRecordBatchInput, _ ...request.Option) (*firehose.PutRecordBatchOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "PutRecordBatch cancelled", ctx.Err())
	default:
		return d.PutRecordBatch(in)
	}
}

func (d *FirehoseDouble) PutRecordRequest(i0 *firehose.PutRecordInput) (r0 *request.Request, r1 *firehose.PutRecordOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutRecordRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*firehose.PutRecordOutput)
	return
}

func (d *FirehoseDouble) PutRecordWithContext(i0 context.Context, i1 *firehose.PutRecordInput, i2 ...request.Option) (r0 *firehose.PutRecordOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutRecordWithContext", i0, i1, i2)
	r0, _ = returns[0].(*firehose.PutRecordOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) fakePutRecordWithContext(ctx context.Context, in *firehose.PutRecordInput, _ ...request.Option) (*firehose.PutRecordOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "PutRecord cancelled", ctx.Err())
	default:
		return d.PutRecord(in)
	}
}

func (d *FirehoseDouble) StartDeliveryStreamEncryption(i0 *firehose.StartDeliveryStreamEncryptionInput) (r0 *firehose.StartDeliveryStreamEncryptionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartDeliveryStreamEncryption", i0)
	r0, _ = returns[0].(*firehose.StartDeliveryStreamEncryptionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) StartDeliveryStreamEncryptionRequest(i0 *firehose.StartDeliveryStreamEncryptionInput) (r0 *request.Request, r1 *firehose.StartDeliveryStreamEncryptionOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartDeliveryStreamEncryptionRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*firehose.StartDeliveryStreamEncryptionOutput)
	return
}

func (d *FirehoseDouble) StartDeliveryStreamEncryptionWithContext(i0 context.Context, i1 *firehose.StartDeliveryStreamEncryptionInput, i2 ...request.Option) (r0 *firehose.StartDeliveryStreamEncryptionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartDeliveryStreamEncryptionWithContext", i0, i1, i2)
	r0, _ = returns[0].(*firehose.StartDeliveryStreamEncryptionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) fakeStartDeliveryStreamEncryptionWithContext(ctx context.Context, in *firehose.StartDeliveryStreamEncryptionInput, _ ...request.Option) (*firehose.StartDeliveryStreamEncryptionOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StartDeliveryStreamEncryption cancelled", ctx.Err())
	default:
		return d.StartDeliveryStreamEncryption(in)
	}
}

func (d *FirehoseDouble) StopDeliveryStreamEncryption(i0 *firehose.StopDeliveryStreamEncryptionInput) (r0 *firehose.StopDeliveryStreamEncryptionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopDeliveryStreamEncryption", i0)
	r0, _ = returns[0].(*firehose.StopDeliveryStreamEncryptionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) StopDeliveryStreamEncryptionRequest(i0 *firehose.StopDeliveryStreamEncryptionInput) (r0 *request.Request, r1 *firehose.StopDeliveryStreamEncryptionOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopDeliveryStreamEncryptionRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*firehose.StopDeliveryStreamEncryptionOutput)
	return
}

func (d *FirehoseDouble) StopDeliveryStreamEncryptionWithContext(i0 context.Context, i1 *firehose.StopDeliveryStreamEncryptionInput, i2 ...request.Option) (r0 *firehose.StopDeliveryStreamEncryptionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopDeliveryStreamEncryptionWithContext", i0, i1, i2)
	r0, _ = returns[0].(*firehose.StopDeliveryStreamEncryptionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) fakeStopDeliveryStreamEncryptionWithContext(ctx context.Context, in *firehose.StopDeliveryStreamEncryptionInput, _ ...request.Option) (*firehose.StopDeliveryStreamEncryptionOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StopDeliveryStreamEncryption cancelled", ctx.Err())
	default:
		return d.StopDeliveryStreamEncryption(in)
	}
}

func (d *FirehoseDouble) TagDeliveryStream(i0 *firehose.TagDeliveryStreamInput) (r0 *firehose.TagDeliveryStreamOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagDeliveryStream", i0)
	r0, _ = returns[0].(*firehose.TagDeliveryStreamOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) TagDeliveryStreamRequest(i0 *firehose.TagDeliveryStreamInput) (r0 *request.Request, r1 *firehose.TagDeliveryStreamOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagDeliveryStreamRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*firehose.TagDeliveryStreamOutput)
	return
}

func (d *FirehoseDouble) TagDeliveryStreamWithContext(i0 context.Context, i1 *firehose.TagDeliveryStreamInput, i2 ...request.Option) (r0 *firehose.TagDeliveryStreamOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagDeliveryStreamWithContext", i0, i1, i2)
	r0, _ = returns[0].(*firehose.TagDeliveryStreamOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) fakeTagDeliveryStreamWithContext(ctx context.Context, in *firehose.TagDeliveryStreamInput, _ ...request.Option) (*firehose.TagDeliveryStreamOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "TagDeliveryStream cancelled", ctx.Err())
	default:
		return d.TagDeliveryStream(in)
	}
}

func (d *FirehoseDouble) UntagDeliveryStream(i0 *firehose.UntagDeliveryStreamInput) (r0 *firehose.UntagDeliveryStreamOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagDeliveryStream", i0)
	r0, _ = returns[0].(*firehose.UntagDeliveryStreamOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) UntagDeliveryStreamRequest(i0 *firehose.UntagDeliveryStreamInput) (r0 *request.Request, r1 *firehose.UntagDeliveryStreamOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagDeliveryStreamRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*firehose.UntagDeliveryStreamOutput)
	return
}

func (d *FirehoseDouble) UntagDeliveryStreamWithContext(i0 context.Context, i1 *firehose.UntagDeliveryStreamInput, i2 ...request.Option) (r0 *firehose.UntagDeliveryStreamOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagDeliveryStreamWithContext", i0, i1, i2)
	r0, _ = returns[0].(*firehose.UntagDeliveryStreamOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) fakeUntagDeliveryStreamWithContext(ctx context.Context, in *firehose.UntagDeliveryStreamInput, _ ...request.Option) (*firehose.UntagDeliveryStreamOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UntagDeliveryStream cancelled", ctx.Err())
	default:
		return d.UntagDeliveryStream(in)
	}
}

func (d *FirehoseDouble) UpdateDestination(i0 *firehose.UpdateDestinationInput) (r0 *firehose.UpdateDestinationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateDestination", i0)
	r0, _ = returns[0].(*firehose.UpdateDestinationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) UpdateDestinationRequest(i0 *firehose.UpdateDestinationInput) (r0 *request.Request, r1 *firehose.UpdateDestinationOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateDestinationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*firehose.UpdateDestinationOutput)
	return
}

func (d *FirehoseDouble) UpdateDestinationWithContext(i0 context.Context, i1 *firehose.UpdateDestinationInput, i2 ...request.Option) (r0 *firehose.UpdateDestinationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateDestinationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*firehose.UpdateDestinationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FirehoseDouble) fakeUpdateDestinationWithContext(ctx context.Context, in *firehose.UpdateDestinationInput, _ ...request.Option) (*firehose.UpdateDestinationOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateDestination cancelled", ctx.Err())
	default:
		return d.UpdateDestination(in)
	}
}
