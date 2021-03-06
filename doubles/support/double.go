// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:15:20+11:00

// Package supportdouble provides a TestDouble implementation of supportiface.SupportAPI
package supportdouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/support"
	"github.com/aws/aws-sdk-go/service/support/supportiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// SupportDouble is TestDouble for supportiface.SupportAPI
type SupportDouble struct {
	supportiface.SupportAPI
	*awsdouble.AWSTestDouble
}

// Constructor for SupportDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *SupportDouble {
	result := &SupportDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*supportiface.SupportAPI)(nil), configurators...)
	return result
}

func (d *SupportDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *SupportDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "AddAttachmentsToSetWithContext":
		return m.Fake(d.fakeAddAttachmentsToSetWithContext)

	case "AddCommunicationToCaseWithContext":
		return m.Fake(d.fakeAddCommunicationToCaseWithContext)

	case "CreateCaseWithContext":
		return m.Fake(d.fakeCreateCaseWithContext)

	case "DescribeAttachmentWithContext":
		return m.Fake(d.fakeDescribeAttachmentWithContext)

	case "DescribeCasesPages":
		return m.Fake(d.fakeDescribeCasesPages)

	case "DescribeCasesPagesWithContext":
		return m.Fake(d.fakeDescribeCasesPagesWithContext)

	case "DescribeCasesWithContext":
		return m.Fake(d.fakeDescribeCasesWithContext)

	case "DescribeCommunicationsPages":
		return m.Fake(d.fakeDescribeCommunicationsPages)

	case "DescribeCommunicationsPagesWithContext":
		return m.Fake(d.fakeDescribeCommunicationsPagesWithContext)

	case "DescribeCommunicationsWithContext":
		return m.Fake(d.fakeDescribeCommunicationsWithContext)

	case "DescribeServicesWithContext":
		return m.Fake(d.fakeDescribeServicesWithContext)

	case "DescribeSeverityLevelsWithContext":
		return m.Fake(d.fakeDescribeSeverityLevelsWithContext)

	case "DescribeTrustedAdvisorCheckRefreshStatusesWithContext":
		return m.Fake(d.fakeDescribeTrustedAdvisorCheckRefreshStatusesWithContext)

	case "DescribeTrustedAdvisorCheckResultWithContext":
		return m.Fake(d.fakeDescribeTrustedAdvisorCheckResultWithContext)

	case "DescribeTrustedAdvisorCheckSummariesWithContext":
		return m.Fake(d.fakeDescribeTrustedAdvisorCheckSummariesWithContext)

	case "DescribeTrustedAdvisorChecksWithContext":
		return m.Fake(d.fakeDescribeTrustedAdvisorChecksWithContext)

	case "RefreshTrustedAdvisorCheckWithContext":
		return m.Fake(d.fakeRefreshTrustedAdvisorCheckWithContext)

	case "ResolveCaseWithContext":
		return m.Fake(d.fakeResolveCaseWithContext)

	default:
		return nil
	}
}

func (d *SupportDouble) AddAttachmentsToSet(i0 *support.AddAttachmentsToSetInput) (r0 *support.AddAttachmentsToSetOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddAttachmentsToSet", i0)
	r0, _ = returns[0].(*support.AddAttachmentsToSetOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) AddAttachmentsToSetRequest(i0 *support.AddAttachmentsToSetInput) (r0 *request.Request, r1 *support.AddAttachmentsToSetOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddAttachmentsToSetRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*support.AddAttachmentsToSetOutput)
	return
}

func (d *SupportDouble) AddAttachmentsToSetWithContext(i0 context.Context, i1 *support.AddAttachmentsToSetInput, i2 ...request.Option) (r0 *support.AddAttachmentsToSetOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddAttachmentsToSetWithContext", i0, i1, i2)
	r0, _ = returns[0].(*support.AddAttachmentsToSetOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) fakeAddAttachmentsToSetWithContext(ctx context.Context, in *support.AddAttachmentsToSetInput, _ ...request.Option) (*support.AddAttachmentsToSetOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "AddAttachmentsToSet cancelled", ctx.Err())
	default:
		return d.AddAttachmentsToSet(in)
	}
}

func (d *SupportDouble) AddCommunicationToCase(i0 *support.AddCommunicationToCaseInput) (r0 *support.AddCommunicationToCaseOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddCommunicationToCase", i0)
	r0, _ = returns[0].(*support.AddCommunicationToCaseOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) AddCommunicationToCaseRequest(i0 *support.AddCommunicationToCaseInput) (r0 *request.Request, r1 *support.AddCommunicationToCaseOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddCommunicationToCaseRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*support.AddCommunicationToCaseOutput)
	return
}

func (d *SupportDouble) AddCommunicationToCaseWithContext(i0 context.Context, i1 *support.AddCommunicationToCaseInput, i2 ...request.Option) (r0 *support.AddCommunicationToCaseOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddCommunicationToCaseWithContext", i0, i1, i2)
	r0, _ = returns[0].(*support.AddCommunicationToCaseOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) fakeAddCommunicationToCaseWithContext(ctx context.Context, in *support.AddCommunicationToCaseInput, _ ...request.Option) (*support.AddCommunicationToCaseOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "AddCommunicationToCase cancelled", ctx.Err())
	default:
		return d.AddCommunicationToCase(in)
	}
}

func (d *SupportDouble) CreateCase(i0 *support.CreateCaseInput) (r0 *support.CreateCaseOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateCase", i0)
	r0, _ = returns[0].(*support.CreateCaseOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) CreateCaseRequest(i0 *support.CreateCaseInput) (r0 *request.Request, r1 *support.CreateCaseOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateCaseRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*support.CreateCaseOutput)
	return
}

func (d *SupportDouble) CreateCaseWithContext(i0 context.Context, i1 *support.CreateCaseInput, i2 ...request.Option) (r0 *support.CreateCaseOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateCaseWithContext", i0, i1, i2)
	r0, _ = returns[0].(*support.CreateCaseOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) fakeCreateCaseWithContext(ctx context.Context, in *support.CreateCaseInput, _ ...request.Option) (*support.CreateCaseOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateCase cancelled", ctx.Err())
	default:
		return d.CreateCase(in)
	}
}

func (d *SupportDouble) DescribeAttachment(i0 *support.DescribeAttachmentInput) (r0 *support.DescribeAttachmentOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeAttachment", i0)
	r0, _ = returns[0].(*support.DescribeAttachmentOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) DescribeAttachmentRequest(i0 *support.DescribeAttachmentInput) (r0 *request.Request, r1 *support.DescribeAttachmentOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeAttachmentRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*support.DescribeAttachmentOutput)
	return
}

func (d *SupportDouble) DescribeAttachmentWithContext(i0 context.Context, i1 *support.DescribeAttachmentInput, i2 ...request.Option) (r0 *support.DescribeAttachmentOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeAttachmentWithContext", i0, i1, i2)
	r0, _ = returns[0].(*support.DescribeAttachmentOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) fakeDescribeAttachmentWithContext(ctx context.Context, in *support.DescribeAttachmentInput, _ ...request.Option) (*support.DescribeAttachmentOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeAttachment cancelled", ctx.Err())
	default:
		return d.DescribeAttachment(in)
	}
}

func (d *SupportDouble) DescribeCases(i0 *support.DescribeCasesInput) (r0 *support.DescribeCasesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeCases", i0)
	r0, _ = returns[0].(*support.DescribeCasesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) DescribeCasesPages(i0 *support.DescribeCasesInput, i1 func(*support.DescribeCasesOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeCasesPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *SupportDouble) fakeDescribeCasesPages(in *support.DescribeCasesInput, pager func(*support.DescribeCasesOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("DescribeCases", paginators, in, pager)
}

func (d *SupportDouble) DescribeCasesPagesWithContext(i0 context.Context, i1 *support.DescribeCasesInput, i2 func(*support.DescribeCasesOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeCasesPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *SupportDouble) fakeDescribeCasesPagesWithContext(ctx context.Context, in *support.DescribeCasesInput, pager func(*support.DescribeCasesOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("DescribeCases", paginators, ctx, in, pager, options...)
}

func (d *SupportDouble) DescribeCasesRequest(i0 *support.DescribeCasesInput) (r0 *request.Request, r1 *support.DescribeCasesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeCasesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*support.DescribeCasesOutput)
	return
}

func (d *SupportDouble) DescribeCasesWithContext(i0 context.Context, i1 *support.DescribeCasesInput, i2 ...request.Option) (r0 *support.DescribeCasesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeCasesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*support.DescribeCasesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) fakeDescribeCasesWithContext(ctx context.Context, in *support.DescribeCasesInput, _ ...request.Option) (*support.DescribeCasesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeCases cancelled", ctx.Err())
	default:
		return d.DescribeCases(in)
	}
}

func (d *SupportDouble) DescribeCommunications(i0 *support.DescribeCommunicationsInput) (r0 *support.DescribeCommunicationsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeCommunications", i0)
	r0, _ = returns[0].(*support.DescribeCommunicationsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) DescribeCommunicationsPages(i0 *support.DescribeCommunicationsInput, i1 func(*support.DescribeCommunicationsOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeCommunicationsPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *SupportDouble) fakeDescribeCommunicationsPages(in *support.DescribeCommunicationsInput, pager func(*support.DescribeCommunicationsOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("DescribeCommunications", paginators, in, pager)
}

func (d *SupportDouble) DescribeCommunicationsPagesWithContext(i0 context.Context, i1 *support.DescribeCommunicationsInput, i2 func(*support.DescribeCommunicationsOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeCommunicationsPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *SupportDouble) fakeDescribeCommunicationsPagesWithContext(ctx context.Context, in *support.DescribeCommunicationsInput, pager func(*support.DescribeCommunicationsOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("DescribeCommunications", paginators, ctx, in, pager, options...)
}

func (d *SupportDouble) DescribeCommunicationsRequest(i0 *support.DescribeCommunicationsInput) (r0 *request.Request, r1 *support.DescribeCommunicationsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeCommunicationsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*support.DescribeCommunicationsOutput)
	return
}

func (d *SupportDouble) DescribeCommunicationsWithContext(i0 context.Context, i1 *support.DescribeCommunicationsInput, i2 ...request.Option) (r0 *support.DescribeCommunicationsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeCommunicationsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*support.DescribeCommunicationsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) fakeDescribeCommunicationsWithContext(ctx context.Context, in *support.DescribeCommunicationsInput, _ ...request.Option) (*support.DescribeCommunicationsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeCommunications cancelled", ctx.Err())
	default:
		return d.DescribeCommunications(in)
	}
}

func (d *SupportDouble) DescribeServices(i0 *support.DescribeServicesInput) (r0 *support.DescribeServicesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeServices", i0)
	r0, _ = returns[0].(*support.DescribeServicesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) DescribeServicesRequest(i0 *support.DescribeServicesInput) (r0 *request.Request, r1 *support.DescribeServicesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeServicesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*support.DescribeServicesOutput)
	return
}

func (d *SupportDouble) DescribeServicesWithContext(i0 context.Context, i1 *support.DescribeServicesInput, i2 ...request.Option) (r0 *support.DescribeServicesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeServicesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*support.DescribeServicesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) fakeDescribeServicesWithContext(ctx context.Context, in *support.DescribeServicesInput, _ ...request.Option) (*support.DescribeServicesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeServices cancelled", ctx.Err())
	default:
		return d.DescribeServices(in)
	}
}

func (d *SupportDouble) DescribeSeverityLevels(i0 *support.DescribeSeverityLevelsInput) (r0 *support.DescribeSeverityLevelsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeSeverityLevels", i0)
	r0, _ = returns[0].(*support.DescribeSeverityLevelsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) DescribeSeverityLevelsRequest(i0 *support.DescribeSeverityLevelsInput) (r0 *request.Request, r1 *support.DescribeSeverityLevelsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeSeverityLevelsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*support.DescribeSeverityLevelsOutput)
	return
}

func (d *SupportDouble) DescribeSeverityLevelsWithContext(i0 context.Context, i1 *support.DescribeSeverityLevelsInput, i2 ...request.Option) (r0 *support.DescribeSeverityLevelsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeSeverityLevelsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*support.DescribeSeverityLevelsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) fakeDescribeSeverityLevelsWithContext(ctx context.Context, in *support.DescribeSeverityLevelsInput, _ ...request.Option) (*support.DescribeSeverityLevelsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeSeverityLevels cancelled", ctx.Err())
	default:
		return d.DescribeSeverityLevels(in)
	}
}

func (d *SupportDouble) DescribeTrustedAdvisorCheckRefreshStatuses(i0 *support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (r0 *support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeTrustedAdvisorCheckRefreshStatuses", i0)
	r0, _ = returns[0].(*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) DescribeTrustedAdvisorCheckRefreshStatusesRequest(i0 *support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (r0 *request.Request, r1 *support.DescribeTrustedAdvisorCheckRefreshStatusesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeTrustedAdvisorCheckRefreshStatusesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput)
	return
}

func (d *SupportDouble) DescribeTrustedAdvisorCheckRefreshStatusesWithContext(i0 context.Context, i1 *support.DescribeTrustedAdvisorCheckRefreshStatusesInput, i2 ...request.Option) (r0 *support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeTrustedAdvisorCheckRefreshStatusesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) fakeDescribeTrustedAdvisorCheckRefreshStatusesWithContext(ctx context.Context, in *support.DescribeTrustedAdvisorCheckRefreshStatusesInput, _ ...request.Option) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeTrustedAdvisorCheckRefreshStatuses cancelled", ctx.Err())
	default:
		return d.DescribeTrustedAdvisorCheckRefreshStatuses(in)
	}
}

func (d *SupportDouble) DescribeTrustedAdvisorCheckResult(i0 *support.DescribeTrustedAdvisorCheckResultInput) (r0 *support.DescribeTrustedAdvisorCheckResultOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeTrustedAdvisorCheckResult", i0)
	r0, _ = returns[0].(*support.DescribeTrustedAdvisorCheckResultOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) DescribeTrustedAdvisorCheckResultRequest(i0 *support.DescribeTrustedAdvisorCheckResultInput) (r0 *request.Request, r1 *support.DescribeTrustedAdvisorCheckResultOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeTrustedAdvisorCheckResultRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*support.DescribeTrustedAdvisorCheckResultOutput)
	return
}

func (d *SupportDouble) DescribeTrustedAdvisorCheckResultWithContext(i0 context.Context, i1 *support.DescribeTrustedAdvisorCheckResultInput, i2 ...request.Option) (r0 *support.DescribeTrustedAdvisorCheckResultOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeTrustedAdvisorCheckResultWithContext", i0, i1, i2)
	r0, _ = returns[0].(*support.DescribeTrustedAdvisorCheckResultOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) fakeDescribeTrustedAdvisorCheckResultWithContext(ctx context.Context, in *support.DescribeTrustedAdvisorCheckResultInput, _ ...request.Option) (*support.DescribeTrustedAdvisorCheckResultOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeTrustedAdvisorCheckResult cancelled", ctx.Err())
	default:
		return d.DescribeTrustedAdvisorCheckResult(in)
	}
}

func (d *SupportDouble) DescribeTrustedAdvisorCheckSummaries(i0 *support.DescribeTrustedAdvisorCheckSummariesInput) (r0 *support.DescribeTrustedAdvisorCheckSummariesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeTrustedAdvisorCheckSummaries", i0)
	r0, _ = returns[0].(*support.DescribeTrustedAdvisorCheckSummariesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) DescribeTrustedAdvisorCheckSummariesRequest(i0 *support.DescribeTrustedAdvisorCheckSummariesInput) (r0 *request.Request, r1 *support.DescribeTrustedAdvisorCheckSummariesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeTrustedAdvisorCheckSummariesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*support.DescribeTrustedAdvisorCheckSummariesOutput)
	return
}

func (d *SupportDouble) DescribeTrustedAdvisorCheckSummariesWithContext(i0 context.Context, i1 *support.DescribeTrustedAdvisorCheckSummariesInput, i2 ...request.Option) (r0 *support.DescribeTrustedAdvisorCheckSummariesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeTrustedAdvisorCheckSummariesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*support.DescribeTrustedAdvisorCheckSummariesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) fakeDescribeTrustedAdvisorCheckSummariesWithContext(ctx context.Context, in *support.DescribeTrustedAdvisorCheckSummariesInput, _ ...request.Option) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeTrustedAdvisorCheckSummaries cancelled", ctx.Err())
	default:
		return d.DescribeTrustedAdvisorCheckSummaries(in)
	}
}

func (d *SupportDouble) DescribeTrustedAdvisorChecks(i0 *support.DescribeTrustedAdvisorChecksInput) (r0 *support.DescribeTrustedAdvisorChecksOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeTrustedAdvisorChecks", i0)
	r0, _ = returns[0].(*support.DescribeTrustedAdvisorChecksOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) DescribeTrustedAdvisorChecksRequest(i0 *support.DescribeTrustedAdvisorChecksInput) (r0 *request.Request, r1 *support.DescribeTrustedAdvisorChecksOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeTrustedAdvisorChecksRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*support.DescribeTrustedAdvisorChecksOutput)
	return
}

func (d *SupportDouble) DescribeTrustedAdvisorChecksWithContext(i0 context.Context, i1 *support.DescribeTrustedAdvisorChecksInput, i2 ...request.Option) (r0 *support.DescribeTrustedAdvisorChecksOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeTrustedAdvisorChecksWithContext", i0, i1, i2)
	r0, _ = returns[0].(*support.DescribeTrustedAdvisorChecksOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) fakeDescribeTrustedAdvisorChecksWithContext(ctx context.Context, in *support.DescribeTrustedAdvisorChecksInput, _ ...request.Option) (*support.DescribeTrustedAdvisorChecksOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeTrustedAdvisorChecks cancelled", ctx.Err())
	default:
		return d.DescribeTrustedAdvisorChecks(in)
	}
}

func (d *SupportDouble) RefreshTrustedAdvisorCheck(i0 *support.RefreshTrustedAdvisorCheckInput) (r0 *support.RefreshTrustedAdvisorCheckOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("RefreshTrustedAdvisorCheck", i0)
	r0, _ = returns[0].(*support.RefreshTrustedAdvisorCheckOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) RefreshTrustedAdvisorCheckRequest(i0 *support.RefreshTrustedAdvisorCheckInput) (r0 *request.Request, r1 *support.RefreshTrustedAdvisorCheckOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("RefreshTrustedAdvisorCheckRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*support.RefreshTrustedAdvisorCheckOutput)
	return
}

func (d *SupportDouble) RefreshTrustedAdvisorCheckWithContext(i0 context.Context, i1 *support.RefreshTrustedAdvisorCheckInput, i2 ...request.Option) (r0 *support.RefreshTrustedAdvisorCheckOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("RefreshTrustedAdvisorCheckWithContext", i0, i1, i2)
	r0, _ = returns[0].(*support.RefreshTrustedAdvisorCheckOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) fakeRefreshTrustedAdvisorCheckWithContext(ctx context.Context, in *support.RefreshTrustedAdvisorCheckInput, _ ...request.Option) (*support.RefreshTrustedAdvisorCheckOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "RefreshTrustedAdvisorCheck cancelled", ctx.Err())
	default:
		return d.RefreshTrustedAdvisorCheck(in)
	}
}

func (d *SupportDouble) ResolveCase(i0 *support.ResolveCaseInput) (r0 *support.ResolveCaseOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ResolveCase", i0)
	r0, _ = returns[0].(*support.ResolveCaseOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) ResolveCaseRequest(i0 *support.ResolveCaseInput) (r0 *request.Request, r1 *support.ResolveCaseOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ResolveCaseRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*support.ResolveCaseOutput)
	return
}

func (d *SupportDouble) ResolveCaseWithContext(i0 context.Context, i1 *support.ResolveCaseInput, i2 ...request.Option) (r0 *support.ResolveCaseOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ResolveCaseWithContext", i0, i1, i2)
	r0, _ = returns[0].(*support.ResolveCaseOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SupportDouble) fakeResolveCaseWithContext(ctx context.Context, in *support.ResolveCaseInput, _ ...request.Option) (*support.ResolveCaseOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ResolveCase cancelled", ctx.Err())
	default:
		return d.ResolveCase(in)
	}
}
