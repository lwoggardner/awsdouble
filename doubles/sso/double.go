// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:15:15+11:00

// Package ssodouble provides a TestDouble implementation of ssoiface.SSOAPI
package ssodouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sso"
	"github.com/aws/aws-sdk-go/service/sso/ssoiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// SSODouble is TestDouble for ssoiface.SSOAPI
type SSODouble struct {
	ssoiface.SSOAPI
	*awsdouble.AWSTestDouble
}

// Constructor for SSODouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *SSODouble {
	result := &SSODouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*ssoiface.SSOAPI)(nil), configurators...)
	return result
}

func (d *SSODouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *SSODouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "GetRoleCredentialsWithContext":
		return m.Fake(d.fakeGetRoleCredentialsWithContext)

	case "ListAccountRolesPages":
		return m.Fake(d.fakeListAccountRolesPages)

	case "ListAccountRolesPagesWithContext":
		return m.Fake(d.fakeListAccountRolesPagesWithContext)

	case "ListAccountRolesWithContext":
		return m.Fake(d.fakeListAccountRolesWithContext)

	case "ListAccountsPages":
		return m.Fake(d.fakeListAccountsPages)

	case "ListAccountsPagesWithContext":
		return m.Fake(d.fakeListAccountsPagesWithContext)

	case "ListAccountsWithContext":
		return m.Fake(d.fakeListAccountsWithContext)

	case "LogoutWithContext":
		return m.Fake(d.fakeLogoutWithContext)

	default:
		return nil
	}
}

func (d *SSODouble) GetRoleCredentials(i0 *sso.GetRoleCredentialsInput) (r0 *sso.GetRoleCredentialsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetRoleCredentials", i0)
	r0, _ = returns[0].(*sso.GetRoleCredentialsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SSODouble) GetRoleCredentialsRequest(i0 *sso.GetRoleCredentialsInput) (r0 *request.Request, r1 *sso.GetRoleCredentialsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetRoleCredentialsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*sso.GetRoleCredentialsOutput)
	return
}

func (d *SSODouble) GetRoleCredentialsWithContext(i0 context.Context, i1 *sso.GetRoleCredentialsInput, i2 ...request.Option) (r0 *sso.GetRoleCredentialsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetRoleCredentialsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*sso.GetRoleCredentialsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SSODouble) fakeGetRoleCredentialsWithContext(ctx context.Context, in *sso.GetRoleCredentialsInput, _ ...request.Option) (*sso.GetRoleCredentialsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetRoleCredentials cancelled", ctx.Err())
	default:
		return d.GetRoleCredentials(in)
	}
}

func (d *SSODouble) ListAccountRoles(i0 *sso.ListAccountRolesInput) (r0 *sso.ListAccountRolesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListAccountRoles", i0)
	r0, _ = returns[0].(*sso.ListAccountRolesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SSODouble) ListAccountRolesPages(i0 *sso.ListAccountRolesInput, i1 func(*sso.ListAccountRolesOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListAccountRolesPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *SSODouble) fakeListAccountRolesPages(in *sso.ListAccountRolesInput, pager func(*sso.ListAccountRolesOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListAccountRoles", paginators, in, pager)
}

func (d *SSODouble) ListAccountRolesPagesWithContext(i0 context.Context, i1 *sso.ListAccountRolesInput, i2 func(*sso.ListAccountRolesOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListAccountRolesPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *SSODouble) fakeListAccountRolesPagesWithContext(ctx context.Context, in *sso.ListAccountRolesInput, pager func(*sso.ListAccountRolesOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListAccountRoles", paginators, ctx, in, pager, options...)
}

func (d *SSODouble) ListAccountRolesRequest(i0 *sso.ListAccountRolesInput) (r0 *request.Request, r1 *sso.ListAccountRolesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListAccountRolesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*sso.ListAccountRolesOutput)
	return
}

func (d *SSODouble) ListAccountRolesWithContext(i0 context.Context, i1 *sso.ListAccountRolesInput, i2 ...request.Option) (r0 *sso.ListAccountRolesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListAccountRolesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*sso.ListAccountRolesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SSODouble) fakeListAccountRolesWithContext(ctx context.Context, in *sso.ListAccountRolesInput, _ ...request.Option) (*sso.ListAccountRolesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListAccountRoles cancelled", ctx.Err())
	default:
		return d.ListAccountRoles(in)
	}
}

func (d *SSODouble) ListAccounts(i0 *sso.ListAccountsInput) (r0 *sso.ListAccountsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListAccounts", i0)
	r0, _ = returns[0].(*sso.ListAccountsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SSODouble) ListAccountsPages(i0 *sso.ListAccountsInput, i1 func(*sso.ListAccountsOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListAccountsPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *SSODouble) fakeListAccountsPages(in *sso.ListAccountsInput, pager func(*sso.ListAccountsOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListAccounts", paginators, in, pager)
}

func (d *SSODouble) ListAccountsPagesWithContext(i0 context.Context, i1 *sso.ListAccountsInput, i2 func(*sso.ListAccountsOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListAccountsPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *SSODouble) fakeListAccountsPagesWithContext(ctx context.Context, in *sso.ListAccountsInput, pager func(*sso.ListAccountsOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListAccounts", paginators, ctx, in, pager, options...)
}

func (d *SSODouble) ListAccountsRequest(i0 *sso.ListAccountsInput) (r0 *request.Request, r1 *sso.ListAccountsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListAccountsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*sso.ListAccountsOutput)
	return
}

func (d *SSODouble) ListAccountsWithContext(i0 context.Context, i1 *sso.ListAccountsInput, i2 ...request.Option) (r0 *sso.ListAccountsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListAccountsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*sso.ListAccountsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SSODouble) fakeListAccountsWithContext(ctx context.Context, in *sso.ListAccountsInput, _ ...request.Option) (*sso.ListAccountsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListAccounts cancelled", ctx.Err())
	default:
		return d.ListAccounts(in)
	}
}

func (d *SSODouble) Logout(i0 *sso.LogoutInput) (r0 *sso.LogoutOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("Logout", i0)
	r0, _ = returns[0].(*sso.LogoutOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SSODouble) LogoutRequest(i0 *sso.LogoutInput) (r0 *request.Request, r1 *sso.LogoutOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("LogoutRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*sso.LogoutOutput)
	return
}

func (d *SSODouble) LogoutWithContext(i0 context.Context, i1 *sso.LogoutInput, i2 ...request.Option) (r0 *sso.LogoutOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("LogoutWithContext", i0, i1, i2)
	r0, _ = returns[0].(*sso.LogoutOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SSODouble) fakeLogoutWithContext(ctx context.Context, in *sso.LogoutInput, _ ...request.Option) (*sso.LogoutOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "Logout cancelled", ctx.Err())
	default:
		return d.Logout(in)
	}
}
