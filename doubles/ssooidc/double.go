// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:15:16+11:00

// Package ssooidcdouble provides a TestDouble implementation of ssooidciface.SSOOIDCAPI
package ssooidcdouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ssooidc"
	"github.com/aws/aws-sdk-go/service/ssooidc/ssooidciface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// SSOOIDCDouble is TestDouble for ssooidciface.SSOOIDCAPI
type SSOOIDCDouble struct {
	ssooidciface.SSOOIDCAPI
	*awsdouble.AWSTestDouble
}

// Constructor for SSOOIDCDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *SSOOIDCDouble {
	result := &SSOOIDCDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*ssooidciface.SSOOIDCAPI)(nil), configurators...)
	return result
}

func (d *SSOOIDCDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *SSOOIDCDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "CreateTokenWithContext":
		return m.Fake(d.fakeCreateTokenWithContext)

	case "RegisterClientWithContext":
		return m.Fake(d.fakeRegisterClientWithContext)

	case "StartDeviceAuthorizationWithContext":
		return m.Fake(d.fakeStartDeviceAuthorizationWithContext)

	default:
		return nil
	}
}

func (d *SSOOIDCDouble) CreateToken(i0 *ssooidc.CreateTokenInput) (r0 *ssooidc.CreateTokenOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateToken", i0)
	r0, _ = returns[0].(*ssooidc.CreateTokenOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SSOOIDCDouble) CreateTokenRequest(i0 *ssooidc.CreateTokenInput) (r0 *request.Request, r1 *ssooidc.CreateTokenOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateTokenRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*ssooidc.CreateTokenOutput)
	return
}

func (d *SSOOIDCDouble) CreateTokenWithContext(i0 context.Context, i1 *ssooidc.CreateTokenInput, i2 ...request.Option) (r0 *ssooidc.CreateTokenOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateTokenWithContext", i0, i1, i2)
	r0, _ = returns[0].(*ssooidc.CreateTokenOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SSOOIDCDouble) fakeCreateTokenWithContext(ctx context.Context, in *ssooidc.CreateTokenInput, _ ...request.Option) (*ssooidc.CreateTokenOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateToken cancelled", ctx.Err())
	default:
		return d.CreateToken(in)
	}
}

func (d *SSOOIDCDouble) RegisterClient(i0 *ssooidc.RegisterClientInput) (r0 *ssooidc.RegisterClientOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("RegisterClient", i0)
	r0, _ = returns[0].(*ssooidc.RegisterClientOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SSOOIDCDouble) RegisterClientRequest(i0 *ssooidc.RegisterClientInput) (r0 *request.Request, r1 *ssooidc.RegisterClientOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("RegisterClientRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*ssooidc.RegisterClientOutput)
	return
}

func (d *SSOOIDCDouble) RegisterClientWithContext(i0 context.Context, i1 *ssooidc.RegisterClientInput, i2 ...request.Option) (r0 *ssooidc.RegisterClientOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("RegisterClientWithContext", i0, i1, i2)
	r0, _ = returns[0].(*ssooidc.RegisterClientOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SSOOIDCDouble) fakeRegisterClientWithContext(ctx context.Context, in *ssooidc.RegisterClientInput, _ ...request.Option) (*ssooidc.RegisterClientOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "RegisterClient cancelled", ctx.Err())
	default:
		return d.RegisterClient(in)
	}
}

func (d *SSOOIDCDouble) StartDeviceAuthorization(i0 *ssooidc.StartDeviceAuthorizationInput) (r0 *ssooidc.StartDeviceAuthorizationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartDeviceAuthorization", i0)
	r0, _ = returns[0].(*ssooidc.StartDeviceAuthorizationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SSOOIDCDouble) StartDeviceAuthorizationRequest(i0 *ssooidc.StartDeviceAuthorizationInput) (r0 *request.Request, r1 *ssooidc.StartDeviceAuthorizationOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartDeviceAuthorizationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*ssooidc.StartDeviceAuthorizationOutput)
	return
}

func (d *SSOOIDCDouble) StartDeviceAuthorizationWithContext(i0 context.Context, i1 *ssooidc.StartDeviceAuthorizationInput, i2 ...request.Option) (r0 *ssooidc.StartDeviceAuthorizationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartDeviceAuthorizationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*ssooidc.StartDeviceAuthorizationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *SSOOIDCDouble) fakeStartDeviceAuthorizationWithContext(ctx context.Context, in *ssooidc.StartDeviceAuthorizationInput, _ ...request.Option) (*ssooidc.StartDeviceAuthorizationOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StartDeviceAuthorization cancelled", ctx.Err())
	default:
		return d.StartDeviceAuthorization(in)
	}
}
