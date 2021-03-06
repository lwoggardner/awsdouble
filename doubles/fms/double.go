// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:12:49+11:00

// Package fmsdouble provides a TestDouble implementation of fmsiface.FMSAPI
package fmsdouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/fms"
	"github.com/aws/aws-sdk-go/service/fms/fmsiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// FMSDouble is TestDouble for fmsiface.FMSAPI
type FMSDouble struct {
	fmsiface.FMSAPI
	*awsdouble.AWSTestDouble
}

// Constructor for FMSDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *FMSDouble {
	result := &FMSDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*fmsiface.FMSAPI)(nil), configurators...)
	return result
}

func (d *FMSDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *FMSDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "AssociateAdminAccountWithContext":
		return m.Fake(d.fakeAssociateAdminAccountWithContext)

	case "DeleteNotificationChannelWithContext":
		return m.Fake(d.fakeDeleteNotificationChannelWithContext)

	case "DeletePolicyWithContext":
		return m.Fake(d.fakeDeletePolicyWithContext)

	case "DisassociateAdminAccountWithContext":
		return m.Fake(d.fakeDisassociateAdminAccountWithContext)

	case "GetAdminAccountWithContext":
		return m.Fake(d.fakeGetAdminAccountWithContext)

	case "GetComplianceDetailWithContext":
		return m.Fake(d.fakeGetComplianceDetailWithContext)

	case "GetNotificationChannelWithContext":
		return m.Fake(d.fakeGetNotificationChannelWithContext)

	case "GetPolicyWithContext":
		return m.Fake(d.fakeGetPolicyWithContext)

	case "GetProtectionStatusWithContext":
		return m.Fake(d.fakeGetProtectionStatusWithContext)

	case "ListComplianceStatusPages":
		return m.Fake(d.fakeListComplianceStatusPages)

	case "ListComplianceStatusPagesWithContext":
		return m.Fake(d.fakeListComplianceStatusPagesWithContext)

	case "ListComplianceStatusWithContext":
		return m.Fake(d.fakeListComplianceStatusWithContext)

	case "ListMemberAccountsPages":
		return m.Fake(d.fakeListMemberAccountsPages)

	case "ListMemberAccountsPagesWithContext":
		return m.Fake(d.fakeListMemberAccountsPagesWithContext)

	case "ListMemberAccountsWithContext":
		return m.Fake(d.fakeListMemberAccountsWithContext)

	case "ListPoliciesPages":
		return m.Fake(d.fakeListPoliciesPages)

	case "ListPoliciesPagesWithContext":
		return m.Fake(d.fakeListPoliciesPagesWithContext)

	case "ListPoliciesWithContext":
		return m.Fake(d.fakeListPoliciesWithContext)

	case "ListTagsForResourceWithContext":
		return m.Fake(d.fakeListTagsForResourceWithContext)

	case "PutNotificationChannelWithContext":
		return m.Fake(d.fakePutNotificationChannelWithContext)

	case "PutPolicyWithContext":
		return m.Fake(d.fakePutPolicyWithContext)

	case "TagResourceWithContext":
		return m.Fake(d.fakeTagResourceWithContext)

	case "UntagResourceWithContext":
		return m.Fake(d.fakeUntagResourceWithContext)

	default:
		return nil
	}
}

func (d *FMSDouble) AssociateAdminAccount(i0 *fms.AssociateAdminAccountInput) (r0 *fms.AssociateAdminAccountOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AssociateAdminAccount", i0)
	r0, _ = returns[0].(*fms.AssociateAdminAccountOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) AssociateAdminAccountRequest(i0 *fms.AssociateAdminAccountInput) (r0 *request.Request, r1 *fms.AssociateAdminAccountOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AssociateAdminAccountRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*fms.AssociateAdminAccountOutput)
	return
}

func (d *FMSDouble) AssociateAdminAccountWithContext(i0 context.Context, i1 *fms.AssociateAdminAccountInput, i2 ...request.Option) (r0 *fms.AssociateAdminAccountOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AssociateAdminAccountWithContext", i0, i1, i2)
	r0, _ = returns[0].(*fms.AssociateAdminAccountOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) fakeAssociateAdminAccountWithContext(ctx context.Context, in *fms.AssociateAdminAccountInput, _ ...request.Option) (*fms.AssociateAdminAccountOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "AssociateAdminAccount cancelled", ctx.Err())
	default:
		return d.AssociateAdminAccount(in)
	}
}

func (d *FMSDouble) DeleteNotificationChannel(i0 *fms.DeleteNotificationChannelInput) (r0 *fms.DeleteNotificationChannelOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteNotificationChannel", i0)
	r0, _ = returns[0].(*fms.DeleteNotificationChannelOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) DeleteNotificationChannelRequest(i0 *fms.DeleteNotificationChannelInput) (r0 *request.Request, r1 *fms.DeleteNotificationChannelOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteNotificationChannelRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*fms.DeleteNotificationChannelOutput)
	return
}

func (d *FMSDouble) DeleteNotificationChannelWithContext(i0 context.Context, i1 *fms.DeleteNotificationChannelInput, i2 ...request.Option) (r0 *fms.DeleteNotificationChannelOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteNotificationChannelWithContext", i0, i1, i2)
	r0, _ = returns[0].(*fms.DeleteNotificationChannelOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) fakeDeleteNotificationChannelWithContext(ctx context.Context, in *fms.DeleteNotificationChannelInput, _ ...request.Option) (*fms.DeleteNotificationChannelOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteNotificationChannel cancelled", ctx.Err())
	default:
		return d.DeleteNotificationChannel(in)
	}
}

func (d *FMSDouble) DeletePolicy(i0 *fms.DeletePolicyInput) (r0 *fms.DeletePolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeletePolicy", i0)
	r0, _ = returns[0].(*fms.DeletePolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) DeletePolicyRequest(i0 *fms.DeletePolicyInput) (r0 *request.Request, r1 *fms.DeletePolicyOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeletePolicyRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*fms.DeletePolicyOutput)
	return
}

func (d *FMSDouble) DeletePolicyWithContext(i0 context.Context, i1 *fms.DeletePolicyInput, i2 ...request.Option) (r0 *fms.DeletePolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeletePolicyWithContext", i0, i1, i2)
	r0, _ = returns[0].(*fms.DeletePolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) fakeDeletePolicyWithContext(ctx context.Context, in *fms.DeletePolicyInput, _ ...request.Option) (*fms.DeletePolicyOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeletePolicy cancelled", ctx.Err())
	default:
		return d.DeletePolicy(in)
	}
}

func (d *FMSDouble) DisassociateAdminAccount(i0 *fms.DisassociateAdminAccountInput) (r0 *fms.DisassociateAdminAccountOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DisassociateAdminAccount", i0)
	r0, _ = returns[0].(*fms.DisassociateAdminAccountOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) DisassociateAdminAccountRequest(i0 *fms.DisassociateAdminAccountInput) (r0 *request.Request, r1 *fms.DisassociateAdminAccountOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DisassociateAdminAccountRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*fms.DisassociateAdminAccountOutput)
	return
}

func (d *FMSDouble) DisassociateAdminAccountWithContext(i0 context.Context, i1 *fms.DisassociateAdminAccountInput, i2 ...request.Option) (r0 *fms.DisassociateAdminAccountOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DisassociateAdminAccountWithContext", i0, i1, i2)
	r0, _ = returns[0].(*fms.DisassociateAdminAccountOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) fakeDisassociateAdminAccountWithContext(ctx context.Context, in *fms.DisassociateAdminAccountInput, _ ...request.Option) (*fms.DisassociateAdminAccountOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DisassociateAdminAccount cancelled", ctx.Err())
	default:
		return d.DisassociateAdminAccount(in)
	}
}

func (d *FMSDouble) GetAdminAccount(i0 *fms.GetAdminAccountInput) (r0 *fms.GetAdminAccountOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetAdminAccount", i0)
	r0, _ = returns[0].(*fms.GetAdminAccountOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) GetAdminAccountRequest(i0 *fms.GetAdminAccountInput) (r0 *request.Request, r1 *fms.GetAdminAccountOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetAdminAccountRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*fms.GetAdminAccountOutput)
	return
}

func (d *FMSDouble) GetAdminAccountWithContext(i0 context.Context, i1 *fms.GetAdminAccountInput, i2 ...request.Option) (r0 *fms.GetAdminAccountOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetAdminAccountWithContext", i0, i1, i2)
	r0, _ = returns[0].(*fms.GetAdminAccountOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) fakeGetAdminAccountWithContext(ctx context.Context, in *fms.GetAdminAccountInput, _ ...request.Option) (*fms.GetAdminAccountOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetAdminAccount cancelled", ctx.Err())
	default:
		return d.GetAdminAccount(in)
	}
}

func (d *FMSDouble) GetComplianceDetail(i0 *fms.GetComplianceDetailInput) (r0 *fms.GetComplianceDetailOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetComplianceDetail", i0)
	r0, _ = returns[0].(*fms.GetComplianceDetailOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) GetComplianceDetailRequest(i0 *fms.GetComplianceDetailInput) (r0 *request.Request, r1 *fms.GetComplianceDetailOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetComplianceDetailRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*fms.GetComplianceDetailOutput)
	return
}

func (d *FMSDouble) GetComplianceDetailWithContext(i0 context.Context, i1 *fms.GetComplianceDetailInput, i2 ...request.Option) (r0 *fms.GetComplianceDetailOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetComplianceDetailWithContext", i0, i1, i2)
	r0, _ = returns[0].(*fms.GetComplianceDetailOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) fakeGetComplianceDetailWithContext(ctx context.Context, in *fms.GetComplianceDetailInput, _ ...request.Option) (*fms.GetComplianceDetailOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetComplianceDetail cancelled", ctx.Err())
	default:
		return d.GetComplianceDetail(in)
	}
}

func (d *FMSDouble) GetNotificationChannel(i0 *fms.GetNotificationChannelInput) (r0 *fms.GetNotificationChannelOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetNotificationChannel", i0)
	r0, _ = returns[0].(*fms.GetNotificationChannelOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) GetNotificationChannelRequest(i0 *fms.GetNotificationChannelInput) (r0 *request.Request, r1 *fms.GetNotificationChannelOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetNotificationChannelRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*fms.GetNotificationChannelOutput)
	return
}

func (d *FMSDouble) GetNotificationChannelWithContext(i0 context.Context, i1 *fms.GetNotificationChannelInput, i2 ...request.Option) (r0 *fms.GetNotificationChannelOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetNotificationChannelWithContext", i0, i1, i2)
	r0, _ = returns[0].(*fms.GetNotificationChannelOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) fakeGetNotificationChannelWithContext(ctx context.Context, in *fms.GetNotificationChannelInput, _ ...request.Option) (*fms.GetNotificationChannelOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetNotificationChannel cancelled", ctx.Err())
	default:
		return d.GetNotificationChannel(in)
	}
}

func (d *FMSDouble) GetPolicy(i0 *fms.GetPolicyInput) (r0 *fms.GetPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetPolicy", i0)
	r0, _ = returns[0].(*fms.GetPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) GetPolicyRequest(i0 *fms.GetPolicyInput) (r0 *request.Request, r1 *fms.GetPolicyOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetPolicyRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*fms.GetPolicyOutput)
	return
}

func (d *FMSDouble) GetPolicyWithContext(i0 context.Context, i1 *fms.GetPolicyInput, i2 ...request.Option) (r0 *fms.GetPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetPolicyWithContext", i0, i1, i2)
	r0, _ = returns[0].(*fms.GetPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) fakeGetPolicyWithContext(ctx context.Context, in *fms.GetPolicyInput, _ ...request.Option) (*fms.GetPolicyOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetPolicy cancelled", ctx.Err())
	default:
		return d.GetPolicy(in)
	}
}

func (d *FMSDouble) GetProtectionStatus(i0 *fms.GetProtectionStatusInput) (r0 *fms.GetProtectionStatusOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetProtectionStatus", i0)
	r0, _ = returns[0].(*fms.GetProtectionStatusOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) GetProtectionStatusRequest(i0 *fms.GetProtectionStatusInput) (r0 *request.Request, r1 *fms.GetProtectionStatusOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetProtectionStatusRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*fms.GetProtectionStatusOutput)
	return
}

func (d *FMSDouble) GetProtectionStatusWithContext(i0 context.Context, i1 *fms.GetProtectionStatusInput, i2 ...request.Option) (r0 *fms.GetProtectionStatusOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetProtectionStatusWithContext", i0, i1, i2)
	r0, _ = returns[0].(*fms.GetProtectionStatusOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) fakeGetProtectionStatusWithContext(ctx context.Context, in *fms.GetProtectionStatusInput, _ ...request.Option) (*fms.GetProtectionStatusOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetProtectionStatus cancelled", ctx.Err())
	default:
		return d.GetProtectionStatus(in)
	}
}

func (d *FMSDouble) ListComplianceStatus(i0 *fms.ListComplianceStatusInput) (r0 *fms.ListComplianceStatusOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListComplianceStatus", i0)
	r0, _ = returns[0].(*fms.ListComplianceStatusOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) ListComplianceStatusPages(i0 *fms.ListComplianceStatusInput, i1 func(*fms.ListComplianceStatusOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListComplianceStatusPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *FMSDouble) fakeListComplianceStatusPages(in *fms.ListComplianceStatusInput, pager func(*fms.ListComplianceStatusOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListComplianceStatus", paginators, in, pager)
}

func (d *FMSDouble) ListComplianceStatusPagesWithContext(i0 context.Context, i1 *fms.ListComplianceStatusInput, i2 func(*fms.ListComplianceStatusOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListComplianceStatusPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *FMSDouble) fakeListComplianceStatusPagesWithContext(ctx context.Context, in *fms.ListComplianceStatusInput, pager func(*fms.ListComplianceStatusOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListComplianceStatus", paginators, ctx, in, pager, options...)
}

func (d *FMSDouble) ListComplianceStatusRequest(i0 *fms.ListComplianceStatusInput) (r0 *request.Request, r1 *fms.ListComplianceStatusOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListComplianceStatusRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*fms.ListComplianceStatusOutput)
	return
}

func (d *FMSDouble) ListComplianceStatusWithContext(i0 context.Context, i1 *fms.ListComplianceStatusInput, i2 ...request.Option) (r0 *fms.ListComplianceStatusOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListComplianceStatusWithContext", i0, i1, i2)
	r0, _ = returns[0].(*fms.ListComplianceStatusOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) fakeListComplianceStatusWithContext(ctx context.Context, in *fms.ListComplianceStatusInput, _ ...request.Option) (*fms.ListComplianceStatusOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListComplianceStatus cancelled", ctx.Err())
	default:
		return d.ListComplianceStatus(in)
	}
}

func (d *FMSDouble) ListMemberAccounts(i0 *fms.ListMemberAccountsInput) (r0 *fms.ListMemberAccountsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListMemberAccounts", i0)
	r0, _ = returns[0].(*fms.ListMemberAccountsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) ListMemberAccountsPages(i0 *fms.ListMemberAccountsInput, i1 func(*fms.ListMemberAccountsOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListMemberAccountsPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *FMSDouble) fakeListMemberAccountsPages(in *fms.ListMemberAccountsInput, pager func(*fms.ListMemberAccountsOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListMemberAccounts", paginators, in, pager)
}

func (d *FMSDouble) ListMemberAccountsPagesWithContext(i0 context.Context, i1 *fms.ListMemberAccountsInput, i2 func(*fms.ListMemberAccountsOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListMemberAccountsPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *FMSDouble) fakeListMemberAccountsPagesWithContext(ctx context.Context, in *fms.ListMemberAccountsInput, pager func(*fms.ListMemberAccountsOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListMemberAccounts", paginators, ctx, in, pager, options...)
}

func (d *FMSDouble) ListMemberAccountsRequest(i0 *fms.ListMemberAccountsInput) (r0 *request.Request, r1 *fms.ListMemberAccountsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListMemberAccountsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*fms.ListMemberAccountsOutput)
	return
}

func (d *FMSDouble) ListMemberAccountsWithContext(i0 context.Context, i1 *fms.ListMemberAccountsInput, i2 ...request.Option) (r0 *fms.ListMemberAccountsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListMemberAccountsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*fms.ListMemberAccountsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) fakeListMemberAccountsWithContext(ctx context.Context, in *fms.ListMemberAccountsInput, _ ...request.Option) (*fms.ListMemberAccountsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListMemberAccounts cancelled", ctx.Err())
	default:
		return d.ListMemberAccounts(in)
	}
}

func (d *FMSDouble) ListPolicies(i0 *fms.ListPoliciesInput) (r0 *fms.ListPoliciesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListPolicies", i0)
	r0, _ = returns[0].(*fms.ListPoliciesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) ListPoliciesPages(i0 *fms.ListPoliciesInput, i1 func(*fms.ListPoliciesOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListPoliciesPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *FMSDouble) fakeListPoliciesPages(in *fms.ListPoliciesInput, pager func(*fms.ListPoliciesOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListPolicies", paginators, in, pager)
}

func (d *FMSDouble) ListPoliciesPagesWithContext(i0 context.Context, i1 *fms.ListPoliciesInput, i2 func(*fms.ListPoliciesOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListPoliciesPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *FMSDouble) fakeListPoliciesPagesWithContext(ctx context.Context, in *fms.ListPoliciesInput, pager func(*fms.ListPoliciesOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListPolicies", paginators, ctx, in, pager, options...)
}

func (d *FMSDouble) ListPoliciesRequest(i0 *fms.ListPoliciesInput) (r0 *request.Request, r1 *fms.ListPoliciesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListPoliciesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*fms.ListPoliciesOutput)
	return
}

func (d *FMSDouble) ListPoliciesWithContext(i0 context.Context, i1 *fms.ListPoliciesInput, i2 ...request.Option) (r0 *fms.ListPoliciesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListPoliciesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*fms.ListPoliciesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) fakeListPoliciesWithContext(ctx context.Context, in *fms.ListPoliciesInput, _ ...request.Option) (*fms.ListPoliciesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListPolicies cancelled", ctx.Err())
	default:
		return d.ListPolicies(in)
	}
}

func (d *FMSDouble) ListTagsForResource(i0 *fms.ListTagsForResourceInput) (r0 *fms.ListTagsForResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForResource", i0)
	r0, _ = returns[0].(*fms.ListTagsForResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) ListTagsForResourceRequest(i0 *fms.ListTagsForResourceInput) (r0 *request.Request, r1 *fms.ListTagsForResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*fms.ListTagsForResourceOutput)
	return
}

func (d *FMSDouble) ListTagsForResourceWithContext(i0 context.Context, i1 *fms.ListTagsForResourceInput, i2 ...request.Option) (r0 *fms.ListTagsForResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*fms.ListTagsForResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) fakeListTagsForResourceWithContext(ctx context.Context, in *fms.ListTagsForResourceInput, _ ...request.Option) (*fms.ListTagsForResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListTagsForResource cancelled", ctx.Err())
	default:
		return d.ListTagsForResource(in)
	}
}

func (d *FMSDouble) PutNotificationChannel(i0 *fms.PutNotificationChannelInput) (r0 *fms.PutNotificationChannelOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutNotificationChannel", i0)
	r0, _ = returns[0].(*fms.PutNotificationChannelOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) PutNotificationChannelRequest(i0 *fms.PutNotificationChannelInput) (r0 *request.Request, r1 *fms.PutNotificationChannelOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutNotificationChannelRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*fms.PutNotificationChannelOutput)
	return
}

func (d *FMSDouble) PutNotificationChannelWithContext(i0 context.Context, i1 *fms.PutNotificationChannelInput, i2 ...request.Option) (r0 *fms.PutNotificationChannelOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutNotificationChannelWithContext", i0, i1, i2)
	r0, _ = returns[0].(*fms.PutNotificationChannelOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) fakePutNotificationChannelWithContext(ctx context.Context, in *fms.PutNotificationChannelInput, _ ...request.Option) (*fms.PutNotificationChannelOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "PutNotificationChannel cancelled", ctx.Err())
	default:
		return d.PutNotificationChannel(in)
	}
}

func (d *FMSDouble) PutPolicy(i0 *fms.PutPolicyInput) (r0 *fms.PutPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutPolicy", i0)
	r0, _ = returns[0].(*fms.PutPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) PutPolicyRequest(i0 *fms.PutPolicyInput) (r0 *request.Request, r1 *fms.PutPolicyOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutPolicyRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*fms.PutPolicyOutput)
	return
}

func (d *FMSDouble) PutPolicyWithContext(i0 context.Context, i1 *fms.PutPolicyInput, i2 ...request.Option) (r0 *fms.PutPolicyOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutPolicyWithContext", i0, i1, i2)
	r0, _ = returns[0].(*fms.PutPolicyOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) fakePutPolicyWithContext(ctx context.Context, in *fms.PutPolicyInput, _ ...request.Option) (*fms.PutPolicyOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "PutPolicy cancelled", ctx.Err())
	default:
		return d.PutPolicy(in)
	}
}

func (d *FMSDouble) TagResource(i0 *fms.TagResourceInput) (r0 *fms.TagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResource", i0)
	r0, _ = returns[0].(*fms.TagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) TagResourceRequest(i0 *fms.TagResourceInput) (r0 *request.Request, r1 *fms.TagResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*fms.TagResourceOutput)
	return
}

func (d *FMSDouble) TagResourceWithContext(i0 context.Context, i1 *fms.TagResourceInput, i2 ...request.Option) (r0 *fms.TagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*fms.TagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) fakeTagResourceWithContext(ctx context.Context, in *fms.TagResourceInput, _ ...request.Option) (*fms.TagResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "TagResource cancelled", ctx.Err())
	default:
		return d.TagResource(in)
	}
}

func (d *FMSDouble) UntagResource(i0 *fms.UntagResourceInput) (r0 *fms.UntagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResource", i0)
	r0, _ = returns[0].(*fms.UntagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) UntagResourceRequest(i0 *fms.UntagResourceInput) (r0 *request.Request, r1 *fms.UntagResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*fms.UntagResourceOutput)
	return
}

func (d *FMSDouble) UntagResourceWithContext(i0 context.Context, i1 *fms.UntagResourceInput, i2 ...request.Option) (r0 *fms.UntagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*fms.UntagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *FMSDouble) fakeUntagResourceWithContext(ctx context.Context, in *fms.UntagResourceInput, _ ...request.Option) (*fms.UntagResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UntagResource cancelled", ctx.Err())
	default:
		return d.UntagResource(in)
	}
}
