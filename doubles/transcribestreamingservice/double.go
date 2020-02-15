// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:15:25+11:00

// Package transcribestreamingservicedouble provides a TestDouble implementation of transcribestreamingserviceiface.TranscribeStreamingServiceAPI
package transcribestreamingservicedouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/transcribestreamingservice"
	"github.com/aws/aws-sdk-go/service/transcribestreamingservice/transcribestreamingserviceiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// TranscribeStreamingServiceDouble is TestDouble for transcribestreamingserviceiface.TranscribeStreamingServiceAPI
type TranscribeStreamingServiceDouble struct {
	transcribestreamingserviceiface.TranscribeStreamingServiceAPI
	*awsdouble.AWSTestDouble
}

// Constructor for TranscribeStreamingServiceDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *TranscribeStreamingServiceDouble {
	result := &TranscribeStreamingServiceDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*transcribestreamingserviceiface.TranscribeStreamingServiceAPI)(nil), configurators...)
	return result
}

func (d *TranscribeStreamingServiceDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *TranscribeStreamingServiceDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "StartStreamTranscriptionWithContext":
		return m.Fake(d.fakeStartStreamTranscriptionWithContext)

	default:
		return nil
	}
}

func (d *TranscribeStreamingServiceDouble) StartStreamTranscription(i0 *transcribestreamingservice.StartStreamTranscriptionInput) (r0 *transcribestreamingservice.StartStreamTranscriptionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartStreamTranscription", i0)
	r0, _ = returns[0].(*transcribestreamingservice.StartStreamTranscriptionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *TranscribeStreamingServiceDouble) StartStreamTranscriptionRequest(i0 *transcribestreamingservice.StartStreamTranscriptionInput) (r0 *request.Request, r1 *transcribestreamingservice.StartStreamTranscriptionOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartStreamTranscriptionRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*transcribestreamingservice.StartStreamTranscriptionOutput)
	return
}

func (d *TranscribeStreamingServiceDouble) StartStreamTranscriptionWithContext(i0 context.Context, i1 *transcribestreamingservice.StartStreamTranscriptionInput, i2 ...request.Option) (r0 *transcribestreamingservice.StartStreamTranscriptionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartStreamTranscriptionWithContext", i0, i1, i2)
	r0, _ = returns[0].(*transcribestreamingservice.StartStreamTranscriptionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *TranscribeStreamingServiceDouble) fakeStartStreamTranscriptionWithContext(ctx context.Context, in *transcribestreamingservice.StartStreamTranscriptionInput, _ ...request.Option) (*transcribestreamingservice.StartStreamTranscriptionOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StartStreamTranscription cancelled", ctx.Err())
	default:
		return d.StartStreamTranscription(in)
	}
}
