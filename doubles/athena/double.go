// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:11:20+11:00

// Package athenadouble provides a TestDouble implementation of athenaiface.AthenaAPI
package athenadouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/athena/athenaiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// AthenaDouble is TestDouble for athenaiface.AthenaAPI
type AthenaDouble struct {
	athenaiface.AthenaAPI
	*awsdouble.AWSTestDouble
}

// Constructor for AthenaDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *AthenaDouble {
	result := &AthenaDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*athenaiface.AthenaAPI)(nil), configurators...)
	return result
}

func (d *AthenaDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *AthenaDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "BatchGetNamedQueryWithContext":
		return m.Fake(d.fakeBatchGetNamedQueryWithContext)

	case "BatchGetQueryExecutionWithContext":
		return m.Fake(d.fakeBatchGetQueryExecutionWithContext)

	case "CreateNamedQueryWithContext":
		return m.Fake(d.fakeCreateNamedQueryWithContext)

	case "CreateWorkGroupWithContext":
		return m.Fake(d.fakeCreateWorkGroupWithContext)

	case "DeleteNamedQueryWithContext":
		return m.Fake(d.fakeDeleteNamedQueryWithContext)

	case "DeleteWorkGroupWithContext":
		return m.Fake(d.fakeDeleteWorkGroupWithContext)

	case "GetNamedQueryWithContext":
		return m.Fake(d.fakeGetNamedQueryWithContext)

	case "GetQueryExecutionWithContext":
		return m.Fake(d.fakeGetQueryExecutionWithContext)

	case "GetQueryResultsPages":
		return m.Fake(d.fakeGetQueryResultsPages)

	case "GetQueryResultsPagesWithContext":
		return m.Fake(d.fakeGetQueryResultsPagesWithContext)

	case "GetQueryResultsWithContext":
		return m.Fake(d.fakeGetQueryResultsWithContext)

	case "GetWorkGroupWithContext":
		return m.Fake(d.fakeGetWorkGroupWithContext)

	case "ListNamedQueriesPages":
		return m.Fake(d.fakeListNamedQueriesPages)

	case "ListNamedQueriesPagesWithContext":
		return m.Fake(d.fakeListNamedQueriesPagesWithContext)

	case "ListNamedQueriesWithContext":
		return m.Fake(d.fakeListNamedQueriesWithContext)

	case "ListQueryExecutionsPages":
		return m.Fake(d.fakeListQueryExecutionsPages)

	case "ListQueryExecutionsPagesWithContext":
		return m.Fake(d.fakeListQueryExecutionsPagesWithContext)

	case "ListQueryExecutionsWithContext":
		return m.Fake(d.fakeListQueryExecutionsWithContext)

	case "ListTagsForResourceWithContext":
		return m.Fake(d.fakeListTagsForResourceWithContext)

	case "ListWorkGroupsPages":
		return m.Fake(d.fakeListWorkGroupsPages)

	case "ListWorkGroupsPagesWithContext":
		return m.Fake(d.fakeListWorkGroupsPagesWithContext)

	case "ListWorkGroupsWithContext":
		return m.Fake(d.fakeListWorkGroupsWithContext)

	case "StartQueryExecutionWithContext":
		return m.Fake(d.fakeStartQueryExecutionWithContext)

	case "StopQueryExecutionWithContext":
		return m.Fake(d.fakeStopQueryExecutionWithContext)

	case "TagResourceWithContext":
		return m.Fake(d.fakeTagResourceWithContext)

	case "UntagResourceWithContext":
		return m.Fake(d.fakeUntagResourceWithContext)

	case "UpdateWorkGroupWithContext":
		return m.Fake(d.fakeUpdateWorkGroupWithContext)

	default:
		return nil
	}
}

func (d *AthenaDouble) BatchGetNamedQuery(i0 *athena.BatchGetNamedQueryInput) (r0 *athena.BatchGetNamedQueryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchGetNamedQuery", i0)
	r0, _ = returns[0].(*athena.BatchGetNamedQueryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) BatchGetNamedQueryRequest(i0 *athena.BatchGetNamedQueryInput) (r0 *request.Request, r1 *athena.BatchGetNamedQueryOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchGetNamedQueryRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*athena.BatchGetNamedQueryOutput)
	return
}

func (d *AthenaDouble) BatchGetNamedQueryWithContext(i0 context.Context, i1 *athena.BatchGetNamedQueryInput, i2 ...request.Option) (r0 *athena.BatchGetNamedQueryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchGetNamedQueryWithContext", i0, i1, i2)
	r0, _ = returns[0].(*athena.BatchGetNamedQueryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) fakeBatchGetNamedQueryWithContext(ctx context.Context, in *athena.BatchGetNamedQueryInput, _ ...request.Option) (*athena.BatchGetNamedQueryOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "BatchGetNamedQuery cancelled", ctx.Err())
	default:
		return d.BatchGetNamedQuery(in)
	}
}

func (d *AthenaDouble) BatchGetQueryExecution(i0 *athena.BatchGetQueryExecutionInput) (r0 *athena.BatchGetQueryExecutionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchGetQueryExecution", i0)
	r0, _ = returns[0].(*athena.BatchGetQueryExecutionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) BatchGetQueryExecutionRequest(i0 *athena.BatchGetQueryExecutionInput) (r0 *request.Request, r1 *athena.BatchGetQueryExecutionOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchGetQueryExecutionRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*athena.BatchGetQueryExecutionOutput)
	return
}

func (d *AthenaDouble) BatchGetQueryExecutionWithContext(i0 context.Context, i1 *athena.BatchGetQueryExecutionInput, i2 ...request.Option) (r0 *athena.BatchGetQueryExecutionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchGetQueryExecutionWithContext", i0, i1, i2)
	r0, _ = returns[0].(*athena.BatchGetQueryExecutionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) fakeBatchGetQueryExecutionWithContext(ctx context.Context, in *athena.BatchGetQueryExecutionInput, _ ...request.Option) (*athena.BatchGetQueryExecutionOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "BatchGetQueryExecution cancelled", ctx.Err())
	default:
		return d.BatchGetQueryExecution(in)
	}
}

func (d *AthenaDouble) CreateNamedQuery(i0 *athena.CreateNamedQueryInput) (r0 *athena.CreateNamedQueryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateNamedQuery", i0)
	r0, _ = returns[0].(*athena.CreateNamedQueryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) CreateNamedQueryRequest(i0 *athena.CreateNamedQueryInput) (r0 *request.Request, r1 *athena.CreateNamedQueryOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateNamedQueryRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*athena.CreateNamedQueryOutput)
	return
}

func (d *AthenaDouble) CreateNamedQueryWithContext(i0 context.Context, i1 *athena.CreateNamedQueryInput, i2 ...request.Option) (r0 *athena.CreateNamedQueryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateNamedQueryWithContext", i0, i1, i2)
	r0, _ = returns[0].(*athena.CreateNamedQueryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) fakeCreateNamedQueryWithContext(ctx context.Context, in *athena.CreateNamedQueryInput, _ ...request.Option) (*athena.CreateNamedQueryOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateNamedQuery cancelled", ctx.Err())
	default:
		return d.CreateNamedQuery(in)
	}
}

func (d *AthenaDouble) CreateWorkGroup(i0 *athena.CreateWorkGroupInput) (r0 *athena.CreateWorkGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateWorkGroup", i0)
	r0, _ = returns[0].(*athena.CreateWorkGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) CreateWorkGroupRequest(i0 *athena.CreateWorkGroupInput) (r0 *request.Request, r1 *athena.CreateWorkGroupOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateWorkGroupRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*athena.CreateWorkGroupOutput)
	return
}

func (d *AthenaDouble) CreateWorkGroupWithContext(i0 context.Context, i1 *athena.CreateWorkGroupInput, i2 ...request.Option) (r0 *athena.CreateWorkGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateWorkGroupWithContext", i0, i1, i2)
	r0, _ = returns[0].(*athena.CreateWorkGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) fakeCreateWorkGroupWithContext(ctx context.Context, in *athena.CreateWorkGroupInput, _ ...request.Option) (*athena.CreateWorkGroupOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateWorkGroup cancelled", ctx.Err())
	default:
		return d.CreateWorkGroup(in)
	}
}

func (d *AthenaDouble) DeleteNamedQuery(i0 *athena.DeleteNamedQueryInput) (r0 *athena.DeleteNamedQueryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteNamedQuery", i0)
	r0, _ = returns[0].(*athena.DeleteNamedQueryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) DeleteNamedQueryRequest(i0 *athena.DeleteNamedQueryInput) (r0 *request.Request, r1 *athena.DeleteNamedQueryOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteNamedQueryRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*athena.DeleteNamedQueryOutput)
	return
}

func (d *AthenaDouble) DeleteNamedQueryWithContext(i0 context.Context, i1 *athena.DeleteNamedQueryInput, i2 ...request.Option) (r0 *athena.DeleteNamedQueryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteNamedQueryWithContext", i0, i1, i2)
	r0, _ = returns[0].(*athena.DeleteNamedQueryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) fakeDeleteNamedQueryWithContext(ctx context.Context, in *athena.DeleteNamedQueryInput, _ ...request.Option) (*athena.DeleteNamedQueryOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteNamedQuery cancelled", ctx.Err())
	default:
		return d.DeleteNamedQuery(in)
	}
}

func (d *AthenaDouble) DeleteWorkGroup(i0 *athena.DeleteWorkGroupInput) (r0 *athena.DeleteWorkGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteWorkGroup", i0)
	r0, _ = returns[0].(*athena.DeleteWorkGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) DeleteWorkGroupRequest(i0 *athena.DeleteWorkGroupInput) (r0 *request.Request, r1 *athena.DeleteWorkGroupOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteWorkGroupRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*athena.DeleteWorkGroupOutput)
	return
}

func (d *AthenaDouble) DeleteWorkGroupWithContext(i0 context.Context, i1 *athena.DeleteWorkGroupInput, i2 ...request.Option) (r0 *athena.DeleteWorkGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteWorkGroupWithContext", i0, i1, i2)
	r0, _ = returns[0].(*athena.DeleteWorkGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) fakeDeleteWorkGroupWithContext(ctx context.Context, in *athena.DeleteWorkGroupInput, _ ...request.Option) (*athena.DeleteWorkGroupOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteWorkGroup cancelled", ctx.Err())
	default:
		return d.DeleteWorkGroup(in)
	}
}

func (d *AthenaDouble) GetNamedQuery(i0 *athena.GetNamedQueryInput) (r0 *athena.GetNamedQueryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetNamedQuery", i0)
	r0, _ = returns[0].(*athena.GetNamedQueryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) GetNamedQueryRequest(i0 *athena.GetNamedQueryInput) (r0 *request.Request, r1 *athena.GetNamedQueryOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetNamedQueryRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*athena.GetNamedQueryOutput)
	return
}

func (d *AthenaDouble) GetNamedQueryWithContext(i0 context.Context, i1 *athena.GetNamedQueryInput, i2 ...request.Option) (r0 *athena.GetNamedQueryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetNamedQueryWithContext", i0, i1, i2)
	r0, _ = returns[0].(*athena.GetNamedQueryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) fakeGetNamedQueryWithContext(ctx context.Context, in *athena.GetNamedQueryInput, _ ...request.Option) (*athena.GetNamedQueryOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetNamedQuery cancelled", ctx.Err())
	default:
		return d.GetNamedQuery(in)
	}
}

func (d *AthenaDouble) GetQueryExecution(i0 *athena.GetQueryExecutionInput) (r0 *athena.GetQueryExecutionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetQueryExecution", i0)
	r0, _ = returns[0].(*athena.GetQueryExecutionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) GetQueryExecutionRequest(i0 *athena.GetQueryExecutionInput) (r0 *request.Request, r1 *athena.GetQueryExecutionOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetQueryExecutionRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*athena.GetQueryExecutionOutput)
	return
}

func (d *AthenaDouble) GetQueryExecutionWithContext(i0 context.Context, i1 *athena.GetQueryExecutionInput, i2 ...request.Option) (r0 *athena.GetQueryExecutionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetQueryExecutionWithContext", i0, i1, i2)
	r0, _ = returns[0].(*athena.GetQueryExecutionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) fakeGetQueryExecutionWithContext(ctx context.Context, in *athena.GetQueryExecutionInput, _ ...request.Option) (*athena.GetQueryExecutionOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetQueryExecution cancelled", ctx.Err())
	default:
		return d.GetQueryExecution(in)
	}
}

func (d *AthenaDouble) GetQueryResults(i0 *athena.GetQueryResultsInput) (r0 *athena.GetQueryResultsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetQueryResults", i0)
	r0, _ = returns[0].(*athena.GetQueryResultsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) GetQueryResultsPages(i0 *athena.GetQueryResultsInput, i1 func(*athena.GetQueryResultsOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetQueryResultsPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *AthenaDouble) fakeGetQueryResultsPages(in *athena.GetQueryResultsInput, pager func(*athena.GetQueryResultsOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("GetQueryResults", paginators, in, pager)
}

func (d *AthenaDouble) GetQueryResultsPagesWithContext(i0 context.Context, i1 *athena.GetQueryResultsInput, i2 func(*athena.GetQueryResultsOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetQueryResultsPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *AthenaDouble) fakeGetQueryResultsPagesWithContext(ctx context.Context, in *athena.GetQueryResultsInput, pager func(*athena.GetQueryResultsOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("GetQueryResults", paginators, ctx, in, pager, options...)
}

func (d *AthenaDouble) GetQueryResultsRequest(i0 *athena.GetQueryResultsInput) (r0 *request.Request, r1 *athena.GetQueryResultsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetQueryResultsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*athena.GetQueryResultsOutput)
	return
}

func (d *AthenaDouble) GetQueryResultsWithContext(i0 context.Context, i1 *athena.GetQueryResultsInput, i2 ...request.Option) (r0 *athena.GetQueryResultsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetQueryResultsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*athena.GetQueryResultsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) fakeGetQueryResultsWithContext(ctx context.Context, in *athena.GetQueryResultsInput, _ ...request.Option) (*athena.GetQueryResultsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetQueryResults cancelled", ctx.Err())
	default:
		return d.GetQueryResults(in)
	}
}

func (d *AthenaDouble) GetWorkGroup(i0 *athena.GetWorkGroupInput) (r0 *athena.GetWorkGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetWorkGroup", i0)
	r0, _ = returns[0].(*athena.GetWorkGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) GetWorkGroupRequest(i0 *athena.GetWorkGroupInput) (r0 *request.Request, r1 *athena.GetWorkGroupOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetWorkGroupRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*athena.GetWorkGroupOutput)
	return
}

func (d *AthenaDouble) GetWorkGroupWithContext(i0 context.Context, i1 *athena.GetWorkGroupInput, i2 ...request.Option) (r0 *athena.GetWorkGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetWorkGroupWithContext", i0, i1, i2)
	r0, _ = returns[0].(*athena.GetWorkGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) fakeGetWorkGroupWithContext(ctx context.Context, in *athena.GetWorkGroupInput, _ ...request.Option) (*athena.GetWorkGroupOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetWorkGroup cancelled", ctx.Err())
	default:
		return d.GetWorkGroup(in)
	}
}

func (d *AthenaDouble) ListNamedQueries(i0 *athena.ListNamedQueriesInput) (r0 *athena.ListNamedQueriesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListNamedQueries", i0)
	r0, _ = returns[0].(*athena.ListNamedQueriesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) ListNamedQueriesPages(i0 *athena.ListNamedQueriesInput, i1 func(*athena.ListNamedQueriesOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListNamedQueriesPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *AthenaDouble) fakeListNamedQueriesPages(in *athena.ListNamedQueriesInput, pager func(*athena.ListNamedQueriesOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListNamedQueries", paginators, in, pager)
}

func (d *AthenaDouble) ListNamedQueriesPagesWithContext(i0 context.Context, i1 *athena.ListNamedQueriesInput, i2 func(*athena.ListNamedQueriesOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListNamedQueriesPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *AthenaDouble) fakeListNamedQueriesPagesWithContext(ctx context.Context, in *athena.ListNamedQueriesInput, pager func(*athena.ListNamedQueriesOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListNamedQueries", paginators, ctx, in, pager, options...)
}

func (d *AthenaDouble) ListNamedQueriesRequest(i0 *athena.ListNamedQueriesInput) (r0 *request.Request, r1 *athena.ListNamedQueriesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListNamedQueriesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*athena.ListNamedQueriesOutput)
	return
}

func (d *AthenaDouble) ListNamedQueriesWithContext(i0 context.Context, i1 *athena.ListNamedQueriesInput, i2 ...request.Option) (r0 *athena.ListNamedQueriesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListNamedQueriesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*athena.ListNamedQueriesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) fakeListNamedQueriesWithContext(ctx context.Context, in *athena.ListNamedQueriesInput, _ ...request.Option) (*athena.ListNamedQueriesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListNamedQueries cancelled", ctx.Err())
	default:
		return d.ListNamedQueries(in)
	}
}

func (d *AthenaDouble) ListQueryExecutions(i0 *athena.ListQueryExecutionsInput) (r0 *athena.ListQueryExecutionsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListQueryExecutions", i0)
	r0, _ = returns[0].(*athena.ListQueryExecutionsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) ListQueryExecutionsPages(i0 *athena.ListQueryExecutionsInput, i1 func(*athena.ListQueryExecutionsOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListQueryExecutionsPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *AthenaDouble) fakeListQueryExecutionsPages(in *athena.ListQueryExecutionsInput, pager func(*athena.ListQueryExecutionsOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListQueryExecutions", paginators, in, pager)
}

func (d *AthenaDouble) ListQueryExecutionsPagesWithContext(i0 context.Context, i1 *athena.ListQueryExecutionsInput, i2 func(*athena.ListQueryExecutionsOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListQueryExecutionsPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *AthenaDouble) fakeListQueryExecutionsPagesWithContext(ctx context.Context, in *athena.ListQueryExecutionsInput, pager func(*athena.ListQueryExecutionsOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListQueryExecutions", paginators, ctx, in, pager, options...)
}

func (d *AthenaDouble) ListQueryExecutionsRequest(i0 *athena.ListQueryExecutionsInput) (r0 *request.Request, r1 *athena.ListQueryExecutionsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListQueryExecutionsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*athena.ListQueryExecutionsOutput)
	return
}

func (d *AthenaDouble) ListQueryExecutionsWithContext(i0 context.Context, i1 *athena.ListQueryExecutionsInput, i2 ...request.Option) (r0 *athena.ListQueryExecutionsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListQueryExecutionsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*athena.ListQueryExecutionsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) fakeListQueryExecutionsWithContext(ctx context.Context, in *athena.ListQueryExecutionsInput, _ ...request.Option) (*athena.ListQueryExecutionsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListQueryExecutions cancelled", ctx.Err())
	default:
		return d.ListQueryExecutions(in)
	}
}

func (d *AthenaDouble) ListTagsForResource(i0 *athena.ListTagsForResourceInput) (r0 *athena.ListTagsForResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForResource", i0)
	r0, _ = returns[0].(*athena.ListTagsForResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) ListTagsForResourceRequest(i0 *athena.ListTagsForResourceInput) (r0 *request.Request, r1 *athena.ListTagsForResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*athena.ListTagsForResourceOutput)
	return
}

func (d *AthenaDouble) ListTagsForResourceWithContext(i0 context.Context, i1 *athena.ListTagsForResourceInput, i2 ...request.Option) (r0 *athena.ListTagsForResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*athena.ListTagsForResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) fakeListTagsForResourceWithContext(ctx context.Context, in *athena.ListTagsForResourceInput, _ ...request.Option) (*athena.ListTagsForResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListTagsForResource cancelled", ctx.Err())
	default:
		return d.ListTagsForResource(in)
	}
}

func (d *AthenaDouble) ListWorkGroups(i0 *athena.ListWorkGroupsInput) (r0 *athena.ListWorkGroupsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListWorkGroups", i0)
	r0, _ = returns[0].(*athena.ListWorkGroupsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) ListWorkGroupsPages(i0 *athena.ListWorkGroupsInput, i1 func(*athena.ListWorkGroupsOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListWorkGroupsPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *AthenaDouble) fakeListWorkGroupsPages(in *athena.ListWorkGroupsInput, pager func(*athena.ListWorkGroupsOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListWorkGroups", paginators, in, pager)
}

func (d *AthenaDouble) ListWorkGroupsPagesWithContext(i0 context.Context, i1 *athena.ListWorkGroupsInput, i2 func(*athena.ListWorkGroupsOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListWorkGroupsPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *AthenaDouble) fakeListWorkGroupsPagesWithContext(ctx context.Context, in *athena.ListWorkGroupsInput, pager func(*athena.ListWorkGroupsOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListWorkGroups", paginators, ctx, in, pager, options...)
}

func (d *AthenaDouble) ListWorkGroupsRequest(i0 *athena.ListWorkGroupsInput) (r0 *request.Request, r1 *athena.ListWorkGroupsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListWorkGroupsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*athena.ListWorkGroupsOutput)
	return
}

func (d *AthenaDouble) ListWorkGroupsWithContext(i0 context.Context, i1 *athena.ListWorkGroupsInput, i2 ...request.Option) (r0 *athena.ListWorkGroupsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListWorkGroupsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*athena.ListWorkGroupsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) fakeListWorkGroupsWithContext(ctx context.Context, in *athena.ListWorkGroupsInput, _ ...request.Option) (*athena.ListWorkGroupsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListWorkGroups cancelled", ctx.Err())
	default:
		return d.ListWorkGroups(in)
	}
}

func (d *AthenaDouble) StartQueryExecution(i0 *athena.StartQueryExecutionInput) (r0 *athena.StartQueryExecutionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartQueryExecution", i0)
	r0, _ = returns[0].(*athena.StartQueryExecutionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) StartQueryExecutionRequest(i0 *athena.StartQueryExecutionInput) (r0 *request.Request, r1 *athena.StartQueryExecutionOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartQueryExecutionRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*athena.StartQueryExecutionOutput)
	return
}

func (d *AthenaDouble) StartQueryExecutionWithContext(i0 context.Context, i1 *athena.StartQueryExecutionInput, i2 ...request.Option) (r0 *athena.StartQueryExecutionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartQueryExecutionWithContext", i0, i1, i2)
	r0, _ = returns[0].(*athena.StartQueryExecutionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) fakeStartQueryExecutionWithContext(ctx context.Context, in *athena.StartQueryExecutionInput, _ ...request.Option) (*athena.StartQueryExecutionOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StartQueryExecution cancelled", ctx.Err())
	default:
		return d.StartQueryExecution(in)
	}
}

func (d *AthenaDouble) StopQueryExecution(i0 *athena.StopQueryExecutionInput) (r0 *athena.StopQueryExecutionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopQueryExecution", i0)
	r0, _ = returns[0].(*athena.StopQueryExecutionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) StopQueryExecutionRequest(i0 *athena.StopQueryExecutionInput) (r0 *request.Request, r1 *athena.StopQueryExecutionOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopQueryExecutionRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*athena.StopQueryExecutionOutput)
	return
}

func (d *AthenaDouble) StopQueryExecutionWithContext(i0 context.Context, i1 *athena.StopQueryExecutionInput, i2 ...request.Option) (r0 *athena.StopQueryExecutionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopQueryExecutionWithContext", i0, i1, i2)
	r0, _ = returns[0].(*athena.StopQueryExecutionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) fakeStopQueryExecutionWithContext(ctx context.Context, in *athena.StopQueryExecutionInput, _ ...request.Option) (*athena.StopQueryExecutionOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StopQueryExecution cancelled", ctx.Err())
	default:
		return d.StopQueryExecution(in)
	}
}

func (d *AthenaDouble) TagResource(i0 *athena.TagResourceInput) (r0 *athena.TagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResource", i0)
	r0, _ = returns[0].(*athena.TagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) TagResourceRequest(i0 *athena.TagResourceInput) (r0 *request.Request, r1 *athena.TagResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*athena.TagResourceOutput)
	return
}

func (d *AthenaDouble) TagResourceWithContext(i0 context.Context, i1 *athena.TagResourceInput, i2 ...request.Option) (r0 *athena.TagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*athena.TagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) fakeTagResourceWithContext(ctx context.Context, in *athena.TagResourceInput, _ ...request.Option) (*athena.TagResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "TagResource cancelled", ctx.Err())
	default:
		return d.TagResource(in)
	}
}

func (d *AthenaDouble) UntagResource(i0 *athena.UntagResourceInput) (r0 *athena.UntagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResource", i0)
	r0, _ = returns[0].(*athena.UntagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) UntagResourceRequest(i0 *athena.UntagResourceInput) (r0 *request.Request, r1 *athena.UntagResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*athena.UntagResourceOutput)
	return
}

func (d *AthenaDouble) UntagResourceWithContext(i0 context.Context, i1 *athena.UntagResourceInput, i2 ...request.Option) (r0 *athena.UntagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*athena.UntagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) fakeUntagResourceWithContext(ctx context.Context, in *athena.UntagResourceInput, _ ...request.Option) (*athena.UntagResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UntagResource cancelled", ctx.Err())
	default:
		return d.UntagResource(in)
	}
}

func (d *AthenaDouble) UpdateWorkGroup(i0 *athena.UpdateWorkGroupInput) (r0 *athena.UpdateWorkGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateWorkGroup", i0)
	r0, _ = returns[0].(*athena.UpdateWorkGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) UpdateWorkGroupRequest(i0 *athena.UpdateWorkGroupInput) (r0 *request.Request, r1 *athena.UpdateWorkGroupOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateWorkGroupRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*athena.UpdateWorkGroupOutput)
	return
}

func (d *AthenaDouble) UpdateWorkGroupWithContext(i0 context.Context, i1 *athena.UpdateWorkGroupInput, i2 ...request.Option) (r0 *athena.UpdateWorkGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateWorkGroupWithContext", i0, i1, i2)
	r0, _ = returns[0].(*athena.UpdateWorkGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *AthenaDouble) fakeUpdateWorkGroupWithContext(ctx context.Context, in *athena.UpdateWorkGroupInput, _ ...request.Option) (*athena.UpdateWorkGroupOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateWorkGroup cancelled", ctx.Err())
	default:
		return d.UpdateWorkGroup(in)
	}
}
