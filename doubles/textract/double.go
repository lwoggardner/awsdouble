// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:15:22+11:00

// Package textractdouble provides a TestDouble implementation of textractiface.TextractAPI
package textractdouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/textract"
	"github.com/aws/aws-sdk-go/service/textract/textractiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// TextractDouble is TestDouble for textractiface.TextractAPI
type TextractDouble struct {
	textractiface.TextractAPI
	*awsdouble.AWSTestDouble
}

// Constructor for TextractDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *TextractDouble {
	result := &TextractDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*textractiface.TextractAPI)(nil), configurators...)
	return result
}

func (d *TextractDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *TextractDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "AnalyzeDocumentWithContext":
		return m.Fake(d.fakeAnalyzeDocumentWithContext)

	case "DetectDocumentTextWithContext":
		return m.Fake(d.fakeDetectDocumentTextWithContext)

	case "GetDocumentAnalysisWithContext":
		return m.Fake(d.fakeGetDocumentAnalysisWithContext)

	case "GetDocumentTextDetectionWithContext":
		return m.Fake(d.fakeGetDocumentTextDetectionWithContext)

	case "StartDocumentAnalysisWithContext":
		return m.Fake(d.fakeStartDocumentAnalysisWithContext)

	case "StartDocumentTextDetectionWithContext":
		return m.Fake(d.fakeStartDocumentTextDetectionWithContext)

	default:
		return nil
	}
}

func (d *TextractDouble) AnalyzeDocument(i0 *textract.AnalyzeDocumentInput) (r0 *textract.AnalyzeDocumentOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AnalyzeDocument", i0)
	r0, _ = returns[0].(*textract.AnalyzeDocumentOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *TextractDouble) AnalyzeDocumentRequest(i0 *textract.AnalyzeDocumentInput) (r0 *request.Request, r1 *textract.AnalyzeDocumentOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AnalyzeDocumentRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*textract.AnalyzeDocumentOutput)
	return
}

func (d *TextractDouble) AnalyzeDocumentWithContext(i0 context.Context, i1 *textract.AnalyzeDocumentInput, i2 ...request.Option) (r0 *textract.AnalyzeDocumentOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AnalyzeDocumentWithContext", i0, i1, i2)
	r0, _ = returns[0].(*textract.AnalyzeDocumentOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *TextractDouble) fakeAnalyzeDocumentWithContext(ctx context.Context, in *textract.AnalyzeDocumentInput, _ ...request.Option) (*textract.AnalyzeDocumentOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "AnalyzeDocument cancelled", ctx.Err())
	default:
		return d.AnalyzeDocument(in)
	}
}

func (d *TextractDouble) DetectDocumentText(i0 *textract.DetectDocumentTextInput) (r0 *textract.DetectDocumentTextOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DetectDocumentText", i0)
	r0, _ = returns[0].(*textract.DetectDocumentTextOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *TextractDouble) DetectDocumentTextRequest(i0 *textract.DetectDocumentTextInput) (r0 *request.Request, r1 *textract.DetectDocumentTextOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DetectDocumentTextRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*textract.DetectDocumentTextOutput)
	return
}

func (d *TextractDouble) DetectDocumentTextWithContext(i0 context.Context, i1 *textract.DetectDocumentTextInput, i2 ...request.Option) (r0 *textract.DetectDocumentTextOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DetectDocumentTextWithContext", i0, i1, i2)
	r0, _ = returns[0].(*textract.DetectDocumentTextOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *TextractDouble) fakeDetectDocumentTextWithContext(ctx context.Context, in *textract.DetectDocumentTextInput, _ ...request.Option) (*textract.DetectDocumentTextOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DetectDocumentText cancelled", ctx.Err())
	default:
		return d.DetectDocumentText(in)
	}
}

func (d *TextractDouble) GetDocumentAnalysis(i0 *textract.GetDocumentAnalysisInput) (r0 *textract.GetDocumentAnalysisOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetDocumentAnalysis", i0)
	r0, _ = returns[0].(*textract.GetDocumentAnalysisOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *TextractDouble) GetDocumentAnalysisRequest(i0 *textract.GetDocumentAnalysisInput) (r0 *request.Request, r1 *textract.GetDocumentAnalysisOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetDocumentAnalysisRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*textract.GetDocumentAnalysisOutput)
	return
}

func (d *TextractDouble) GetDocumentAnalysisWithContext(i0 context.Context, i1 *textract.GetDocumentAnalysisInput, i2 ...request.Option) (r0 *textract.GetDocumentAnalysisOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetDocumentAnalysisWithContext", i0, i1, i2)
	r0, _ = returns[0].(*textract.GetDocumentAnalysisOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *TextractDouble) fakeGetDocumentAnalysisWithContext(ctx context.Context, in *textract.GetDocumentAnalysisInput, _ ...request.Option) (*textract.GetDocumentAnalysisOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetDocumentAnalysis cancelled", ctx.Err())
	default:
		return d.GetDocumentAnalysis(in)
	}
}

func (d *TextractDouble) GetDocumentTextDetection(i0 *textract.GetDocumentTextDetectionInput) (r0 *textract.GetDocumentTextDetectionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetDocumentTextDetection", i0)
	r0, _ = returns[0].(*textract.GetDocumentTextDetectionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *TextractDouble) GetDocumentTextDetectionRequest(i0 *textract.GetDocumentTextDetectionInput) (r0 *request.Request, r1 *textract.GetDocumentTextDetectionOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetDocumentTextDetectionRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*textract.GetDocumentTextDetectionOutput)
	return
}

func (d *TextractDouble) GetDocumentTextDetectionWithContext(i0 context.Context, i1 *textract.GetDocumentTextDetectionInput, i2 ...request.Option) (r0 *textract.GetDocumentTextDetectionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetDocumentTextDetectionWithContext", i0, i1, i2)
	r0, _ = returns[0].(*textract.GetDocumentTextDetectionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *TextractDouble) fakeGetDocumentTextDetectionWithContext(ctx context.Context, in *textract.GetDocumentTextDetectionInput, _ ...request.Option) (*textract.GetDocumentTextDetectionOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetDocumentTextDetection cancelled", ctx.Err())
	default:
		return d.GetDocumentTextDetection(in)
	}
}

func (d *TextractDouble) StartDocumentAnalysis(i0 *textract.StartDocumentAnalysisInput) (r0 *textract.StartDocumentAnalysisOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartDocumentAnalysis", i0)
	r0, _ = returns[0].(*textract.StartDocumentAnalysisOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *TextractDouble) StartDocumentAnalysisRequest(i0 *textract.StartDocumentAnalysisInput) (r0 *request.Request, r1 *textract.StartDocumentAnalysisOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartDocumentAnalysisRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*textract.StartDocumentAnalysisOutput)
	return
}

func (d *TextractDouble) StartDocumentAnalysisWithContext(i0 context.Context, i1 *textract.StartDocumentAnalysisInput, i2 ...request.Option) (r0 *textract.StartDocumentAnalysisOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartDocumentAnalysisWithContext", i0, i1, i2)
	r0, _ = returns[0].(*textract.StartDocumentAnalysisOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *TextractDouble) fakeStartDocumentAnalysisWithContext(ctx context.Context, in *textract.StartDocumentAnalysisInput, _ ...request.Option) (*textract.StartDocumentAnalysisOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StartDocumentAnalysis cancelled", ctx.Err())
	default:
		return d.StartDocumentAnalysis(in)
	}
}

func (d *TextractDouble) StartDocumentTextDetection(i0 *textract.StartDocumentTextDetectionInput) (r0 *textract.StartDocumentTextDetectionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartDocumentTextDetection", i0)
	r0, _ = returns[0].(*textract.StartDocumentTextDetectionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *TextractDouble) StartDocumentTextDetectionRequest(i0 *textract.StartDocumentTextDetectionInput) (r0 *request.Request, r1 *textract.StartDocumentTextDetectionOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartDocumentTextDetectionRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*textract.StartDocumentTextDetectionOutput)
	return
}

func (d *TextractDouble) StartDocumentTextDetectionWithContext(i0 context.Context, i1 *textract.StartDocumentTextDetectionInput, i2 ...request.Option) (r0 *textract.StartDocumentTextDetectionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartDocumentTextDetectionWithContext", i0, i1, i2)
	r0, _ = returns[0].(*textract.StartDocumentTextDetectionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *TextractDouble) fakeStartDocumentTextDetectionWithContext(ctx context.Context, in *textract.StartDocumentTextDetectionInput, _ ...request.Option) (*textract.StartDocumentTextDetectionOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StartDocumentTextDetection cancelled", ctx.Err())
	default:
		return d.StartDocumentTextDetection(in)
	}
}
