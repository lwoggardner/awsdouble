// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:13:35+11:00

// Package lakeformationdouble provides a TestDouble implementation of lakeformationiface.LakeFormationAPI
package lakeformationdouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/lakeformation"
	"github.com/aws/aws-sdk-go/service/lakeformation/lakeformationiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// LakeFormationDouble is TestDouble for lakeformationiface.LakeFormationAPI
type LakeFormationDouble struct {
	lakeformationiface.LakeFormationAPI
	*awsdouble.AWSTestDouble
}

// Constructor for LakeFormationDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *LakeFormationDouble {
	result := &LakeFormationDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*lakeformationiface.LakeFormationAPI)(nil), configurators...)
	return result
}

func (d *LakeFormationDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *LakeFormationDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "BatchGrantPermissionsWithContext":
		return m.Fake(d.fakeBatchGrantPermissionsWithContext)

	case "BatchRevokePermissionsWithContext":
		return m.Fake(d.fakeBatchRevokePermissionsWithContext)

	case "DeregisterResourceWithContext":
		return m.Fake(d.fakeDeregisterResourceWithContext)

	case "DescribeResourceWithContext":
		return m.Fake(d.fakeDescribeResourceWithContext)

	case "GetDataLakeSettingsWithContext":
		return m.Fake(d.fakeGetDataLakeSettingsWithContext)

	case "GetEffectivePermissionsForPathPages":
		return m.Fake(d.fakeGetEffectivePermissionsForPathPages)

	case "GetEffectivePermissionsForPathPagesWithContext":
		return m.Fake(d.fakeGetEffectivePermissionsForPathPagesWithContext)

	case "GetEffectivePermissionsForPathWithContext":
		return m.Fake(d.fakeGetEffectivePermissionsForPathWithContext)

	case "GrantPermissionsWithContext":
		return m.Fake(d.fakeGrantPermissionsWithContext)

	case "ListPermissionsPages":
		return m.Fake(d.fakeListPermissionsPages)

	case "ListPermissionsPagesWithContext":
		return m.Fake(d.fakeListPermissionsPagesWithContext)

	case "ListPermissionsWithContext":
		return m.Fake(d.fakeListPermissionsWithContext)

	case "ListResourcesPages":
		return m.Fake(d.fakeListResourcesPages)

	case "ListResourcesPagesWithContext":
		return m.Fake(d.fakeListResourcesPagesWithContext)

	case "ListResourcesWithContext":
		return m.Fake(d.fakeListResourcesWithContext)

	case "PutDataLakeSettingsWithContext":
		return m.Fake(d.fakePutDataLakeSettingsWithContext)

	case "RegisterResourceWithContext":
		return m.Fake(d.fakeRegisterResourceWithContext)

	case "RevokePermissionsWithContext":
		return m.Fake(d.fakeRevokePermissionsWithContext)

	case "UpdateResourceWithContext":
		return m.Fake(d.fakeUpdateResourceWithContext)

	default:
		return nil
	}
}

func (d *LakeFormationDouble) BatchGrantPermissions(i0 *lakeformation.BatchGrantPermissionsInput) (r0 *lakeformation.BatchGrantPermissionsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchGrantPermissions", i0)
	r0, _ = returns[0].(*lakeformation.BatchGrantPermissionsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) BatchGrantPermissionsRequest(i0 *lakeformation.BatchGrantPermissionsInput) (r0 *request.Request, r1 *lakeformation.BatchGrantPermissionsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchGrantPermissionsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*lakeformation.BatchGrantPermissionsOutput)
	return
}

func (d *LakeFormationDouble) BatchGrantPermissionsWithContext(i0 context.Context, i1 *lakeformation.BatchGrantPermissionsInput, i2 ...request.Option) (r0 *lakeformation.BatchGrantPermissionsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchGrantPermissionsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*lakeformation.BatchGrantPermissionsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) fakeBatchGrantPermissionsWithContext(ctx context.Context, in *lakeformation.BatchGrantPermissionsInput, _ ...request.Option) (*lakeformation.BatchGrantPermissionsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "BatchGrantPermissions cancelled", ctx.Err())
	default:
		return d.BatchGrantPermissions(in)
	}
}

func (d *LakeFormationDouble) BatchRevokePermissions(i0 *lakeformation.BatchRevokePermissionsInput) (r0 *lakeformation.BatchRevokePermissionsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchRevokePermissions", i0)
	r0, _ = returns[0].(*lakeformation.BatchRevokePermissionsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) BatchRevokePermissionsRequest(i0 *lakeformation.BatchRevokePermissionsInput) (r0 *request.Request, r1 *lakeformation.BatchRevokePermissionsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchRevokePermissionsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*lakeformation.BatchRevokePermissionsOutput)
	return
}

func (d *LakeFormationDouble) BatchRevokePermissionsWithContext(i0 context.Context, i1 *lakeformation.BatchRevokePermissionsInput, i2 ...request.Option) (r0 *lakeformation.BatchRevokePermissionsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchRevokePermissionsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*lakeformation.BatchRevokePermissionsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) fakeBatchRevokePermissionsWithContext(ctx context.Context, in *lakeformation.BatchRevokePermissionsInput, _ ...request.Option) (*lakeformation.BatchRevokePermissionsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "BatchRevokePermissions cancelled", ctx.Err())
	default:
		return d.BatchRevokePermissions(in)
	}
}

func (d *LakeFormationDouble) DeregisterResource(i0 *lakeformation.DeregisterResourceInput) (r0 *lakeformation.DeregisterResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeregisterResource", i0)
	r0, _ = returns[0].(*lakeformation.DeregisterResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) DeregisterResourceRequest(i0 *lakeformation.DeregisterResourceInput) (r0 *request.Request, r1 *lakeformation.DeregisterResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeregisterResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*lakeformation.DeregisterResourceOutput)
	return
}

func (d *LakeFormationDouble) DeregisterResourceWithContext(i0 context.Context, i1 *lakeformation.DeregisterResourceInput, i2 ...request.Option) (r0 *lakeformation.DeregisterResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeregisterResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*lakeformation.DeregisterResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) fakeDeregisterResourceWithContext(ctx context.Context, in *lakeformation.DeregisterResourceInput, _ ...request.Option) (*lakeformation.DeregisterResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeregisterResource cancelled", ctx.Err())
	default:
		return d.DeregisterResource(in)
	}
}

func (d *LakeFormationDouble) DescribeResource(i0 *lakeformation.DescribeResourceInput) (r0 *lakeformation.DescribeResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeResource", i0)
	r0, _ = returns[0].(*lakeformation.DescribeResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) DescribeResourceRequest(i0 *lakeformation.DescribeResourceInput) (r0 *request.Request, r1 *lakeformation.DescribeResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*lakeformation.DescribeResourceOutput)
	return
}

func (d *LakeFormationDouble) DescribeResourceWithContext(i0 context.Context, i1 *lakeformation.DescribeResourceInput, i2 ...request.Option) (r0 *lakeformation.DescribeResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*lakeformation.DescribeResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) fakeDescribeResourceWithContext(ctx context.Context, in *lakeformation.DescribeResourceInput, _ ...request.Option) (*lakeformation.DescribeResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeResource cancelled", ctx.Err())
	default:
		return d.DescribeResource(in)
	}
}

func (d *LakeFormationDouble) GetDataLakeSettings(i0 *lakeformation.GetDataLakeSettingsInput) (r0 *lakeformation.GetDataLakeSettingsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetDataLakeSettings", i0)
	r0, _ = returns[0].(*lakeformation.GetDataLakeSettingsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) GetDataLakeSettingsRequest(i0 *lakeformation.GetDataLakeSettingsInput) (r0 *request.Request, r1 *lakeformation.GetDataLakeSettingsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetDataLakeSettingsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*lakeformation.GetDataLakeSettingsOutput)
	return
}

func (d *LakeFormationDouble) GetDataLakeSettingsWithContext(i0 context.Context, i1 *lakeformation.GetDataLakeSettingsInput, i2 ...request.Option) (r0 *lakeformation.GetDataLakeSettingsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetDataLakeSettingsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*lakeformation.GetDataLakeSettingsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) fakeGetDataLakeSettingsWithContext(ctx context.Context, in *lakeformation.GetDataLakeSettingsInput, _ ...request.Option) (*lakeformation.GetDataLakeSettingsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetDataLakeSettings cancelled", ctx.Err())
	default:
		return d.GetDataLakeSettings(in)
	}
}

func (d *LakeFormationDouble) GetEffectivePermissionsForPath(i0 *lakeformation.GetEffectivePermissionsForPathInput) (r0 *lakeformation.GetEffectivePermissionsForPathOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetEffectivePermissionsForPath", i0)
	r0, _ = returns[0].(*lakeformation.GetEffectivePermissionsForPathOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) GetEffectivePermissionsForPathPages(i0 *lakeformation.GetEffectivePermissionsForPathInput, i1 func(*lakeformation.GetEffectivePermissionsForPathOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetEffectivePermissionsForPathPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *LakeFormationDouble) fakeGetEffectivePermissionsForPathPages(in *lakeformation.GetEffectivePermissionsForPathInput, pager func(*lakeformation.GetEffectivePermissionsForPathOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("GetEffectivePermissionsForPath", paginators, in, pager)
}

func (d *LakeFormationDouble) GetEffectivePermissionsForPathPagesWithContext(i0 context.Context, i1 *lakeformation.GetEffectivePermissionsForPathInput, i2 func(*lakeformation.GetEffectivePermissionsForPathOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetEffectivePermissionsForPathPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *LakeFormationDouble) fakeGetEffectivePermissionsForPathPagesWithContext(ctx context.Context, in *lakeformation.GetEffectivePermissionsForPathInput, pager func(*lakeformation.GetEffectivePermissionsForPathOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("GetEffectivePermissionsForPath", paginators, ctx, in, pager, options...)
}

func (d *LakeFormationDouble) GetEffectivePermissionsForPathRequest(i0 *lakeformation.GetEffectivePermissionsForPathInput) (r0 *request.Request, r1 *lakeformation.GetEffectivePermissionsForPathOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetEffectivePermissionsForPathRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*lakeformation.GetEffectivePermissionsForPathOutput)
	return
}

func (d *LakeFormationDouble) GetEffectivePermissionsForPathWithContext(i0 context.Context, i1 *lakeformation.GetEffectivePermissionsForPathInput, i2 ...request.Option) (r0 *lakeformation.GetEffectivePermissionsForPathOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetEffectivePermissionsForPathWithContext", i0, i1, i2)
	r0, _ = returns[0].(*lakeformation.GetEffectivePermissionsForPathOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) fakeGetEffectivePermissionsForPathWithContext(ctx context.Context, in *lakeformation.GetEffectivePermissionsForPathInput, _ ...request.Option) (*lakeformation.GetEffectivePermissionsForPathOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetEffectivePermissionsForPath cancelled", ctx.Err())
	default:
		return d.GetEffectivePermissionsForPath(in)
	}
}

func (d *LakeFormationDouble) GrantPermissions(i0 *lakeformation.GrantPermissionsInput) (r0 *lakeformation.GrantPermissionsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GrantPermissions", i0)
	r0, _ = returns[0].(*lakeformation.GrantPermissionsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) GrantPermissionsRequest(i0 *lakeformation.GrantPermissionsInput) (r0 *request.Request, r1 *lakeformation.GrantPermissionsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GrantPermissionsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*lakeformation.GrantPermissionsOutput)
	return
}

func (d *LakeFormationDouble) GrantPermissionsWithContext(i0 context.Context, i1 *lakeformation.GrantPermissionsInput, i2 ...request.Option) (r0 *lakeformation.GrantPermissionsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GrantPermissionsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*lakeformation.GrantPermissionsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) fakeGrantPermissionsWithContext(ctx context.Context, in *lakeformation.GrantPermissionsInput, _ ...request.Option) (*lakeformation.GrantPermissionsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GrantPermissions cancelled", ctx.Err())
	default:
		return d.GrantPermissions(in)
	}
}

func (d *LakeFormationDouble) ListPermissions(i0 *lakeformation.ListPermissionsInput) (r0 *lakeformation.ListPermissionsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListPermissions", i0)
	r0, _ = returns[0].(*lakeformation.ListPermissionsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) ListPermissionsPages(i0 *lakeformation.ListPermissionsInput, i1 func(*lakeformation.ListPermissionsOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListPermissionsPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *LakeFormationDouble) fakeListPermissionsPages(in *lakeformation.ListPermissionsInput, pager func(*lakeformation.ListPermissionsOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListPermissions", paginators, in, pager)
}

func (d *LakeFormationDouble) ListPermissionsPagesWithContext(i0 context.Context, i1 *lakeformation.ListPermissionsInput, i2 func(*lakeformation.ListPermissionsOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListPermissionsPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *LakeFormationDouble) fakeListPermissionsPagesWithContext(ctx context.Context, in *lakeformation.ListPermissionsInput, pager func(*lakeformation.ListPermissionsOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListPermissions", paginators, ctx, in, pager, options...)
}

func (d *LakeFormationDouble) ListPermissionsRequest(i0 *lakeformation.ListPermissionsInput) (r0 *request.Request, r1 *lakeformation.ListPermissionsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListPermissionsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*lakeformation.ListPermissionsOutput)
	return
}

func (d *LakeFormationDouble) ListPermissionsWithContext(i0 context.Context, i1 *lakeformation.ListPermissionsInput, i2 ...request.Option) (r0 *lakeformation.ListPermissionsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListPermissionsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*lakeformation.ListPermissionsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) fakeListPermissionsWithContext(ctx context.Context, in *lakeformation.ListPermissionsInput, _ ...request.Option) (*lakeformation.ListPermissionsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListPermissions cancelled", ctx.Err())
	default:
		return d.ListPermissions(in)
	}
}

func (d *LakeFormationDouble) ListResources(i0 *lakeformation.ListResourcesInput) (r0 *lakeformation.ListResourcesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListResources", i0)
	r0, _ = returns[0].(*lakeformation.ListResourcesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) ListResourcesPages(i0 *lakeformation.ListResourcesInput, i1 func(*lakeformation.ListResourcesOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListResourcesPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *LakeFormationDouble) fakeListResourcesPages(in *lakeformation.ListResourcesInput, pager func(*lakeformation.ListResourcesOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListResources", paginators, in, pager)
}

func (d *LakeFormationDouble) ListResourcesPagesWithContext(i0 context.Context, i1 *lakeformation.ListResourcesInput, i2 func(*lakeformation.ListResourcesOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListResourcesPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *LakeFormationDouble) fakeListResourcesPagesWithContext(ctx context.Context, in *lakeformation.ListResourcesInput, pager func(*lakeformation.ListResourcesOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListResources", paginators, ctx, in, pager, options...)
}

func (d *LakeFormationDouble) ListResourcesRequest(i0 *lakeformation.ListResourcesInput) (r0 *request.Request, r1 *lakeformation.ListResourcesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListResourcesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*lakeformation.ListResourcesOutput)
	return
}

func (d *LakeFormationDouble) ListResourcesWithContext(i0 context.Context, i1 *lakeformation.ListResourcesInput, i2 ...request.Option) (r0 *lakeformation.ListResourcesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListResourcesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*lakeformation.ListResourcesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) fakeListResourcesWithContext(ctx context.Context, in *lakeformation.ListResourcesInput, _ ...request.Option) (*lakeformation.ListResourcesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListResources cancelled", ctx.Err())
	default:
		return d.ListResources(in)
	}
}

func (d *LakeFormationDouble) PutDataLakeSettings(i0 *lakeformation.PutDataLakeSettingsInput) (r0 *lakeformation.PutDataLakeSettingsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutDataLakeSettings", i0)
	r0, _ = returns[0].(*lakeformation.PutDataLakeSettingsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) PutDataLakeSettingsRequest(i0 *lakeformation.PutDataLakeSettingsInput) (r0 *request.Request, r1 *lakeformation.PutDataLakeSettingsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutDataLakeSettingsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*lakeformation.PutDataLakeSettingsOutput)
	return
}

func (d *LakeFormationDouble) PutDataLakeSettingsWithContext(i0 context.Context, i1 *lakeformation.PutDataLakeSettingsInput, i2 ...request.Option) (r0 *lakeformation.PutDataLakeSettingsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("PutDataLakeSettingsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*lakeformation.PutDataLakeSettingsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) fakePutDataLakeSettingsWithContext(ctx context.Context, in *lakeformation.PutDataLakeSettingsInput, _ ...request.Option) (*lakeformation.PutDataLakeSettingsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "PutDataLakeSettings cancelled", ctx.Err())
	default:
		return d.PutDataLakeSettings(in)
	}
}

func (d *LakeFormationDouble) RegisterResource(i0 *lakeformation.RegisterResourceInput) (r0 *lakeformation.RegisterResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("RegisterResource", i0)
	r0, _ = returns[0].(*lakeformation.RegisterResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) RegisterResourceRequest(i0 *lakeformation.RegisterResourceInput) (r0 *request.Request, r1 *lakeformation.RegisterResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("RegisterResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*lakeformation.RegisterResourceOutput)
	return
}

func (d *LakeFormationDouble) RegisterResourceWithContext(i0 context.Context, i1 *lakeformation.RegisterResourceInput, i2 ...request.Option) (r0 *lakeformation.RegisterResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("RegisterResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*lakeformation.RegisterResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) fakeRegisterResourceWithContext(ctx context.Context, in *lakeformation.RegisterResourceInput, _ ...request.Option) (*lakeformation.RegisterResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "RegisterResource cancelled", ctx.Err())
	default:
		return d.RegisterResource(in)
	}
}

func (d *LakeFormationDouble) RevokePermissions(i0 *lakeformation.RevokePermissionsInput) (r0 *lakeformation.RevokePermissionsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("RevokePermissions", i0)
	r0, _ = returns[0].(*lakeformation.RevokePermissionsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) RevokePermissionsRequest(i0 *lakeformation.RevokePermissionsInput) (r0 *request.Request, r1 *lakeformation.RevokePermissionsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("RevokePermissionsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*lakeformation.RevokePermissionsOutput)
	return
}

func (d *LakeFormationDouble) RevokePermissionsWithContext(i0 context.Context, i1 *lakeformation.RevokePermissionsInput, i2 ...request.Option) (r0 *lakeformation.RevokePermissionsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("RevokePermissionsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*lakeformation.RevokePermissionsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) fakeRevokePermissionsWithContext(ctx context.Context, in *lakeformation.RevokePermissionsInput, _ ...request.Option) (*lakeformation.RevokePermissionsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "RevokePermissions cancelled", ctx.Err())
	default:
		return d.RevokePermissions(in)
	}
}

func (d *LakeFormationDouble) UpdateResource(i0 *lakeformation.UpdateResourceInput) (r0 *lakeformation.UpdateResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateResource", i0)
	r0, _ = returns[0].(*lakeformation.UpdateResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) UpdateResourceRequest(i0 *lakeformation.UpdateResourceInput) (r0 *request.Request, r1 *lakeformation.UpdateResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*lakeformation.UpdateResourceOutput)
	return
}

func (d *LakeFormationDouble) UpdateResourceWithContext(i0 context.Context, i1 *lakeformation.UpdateResourceInput, i2 ...request.Option) (r0 *lakeformation.UpdateResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*lakeformation.UpdateResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *LakeFormationDouble) fakeUpdateResourceWithContext(ctx context.Context, in *lakeformation.UpdateResourceInput, _ ...request.Option) (*lakeformation.UpdateResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateResource cancelled", ctx.Err())
	default:
		return d.UpdateResource(in)
	}
}
