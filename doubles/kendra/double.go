// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:13:24+11:00

// Package kendradouble provides a TestDouble implementation of kendraiface.KendraAPI
package kendradouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kendra"
	"github.com/aws/aws-sdk-go/service/kendra/kendraiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// KendraDouble is TestDouble for kendraiface.KendraAPI
type KendraDouble struct {
	kendraiface.KendraAPI
	*awsdouble.AWSTestDouble
}

// Constructor for KendraDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *KendraDouble {
	result := &KendraDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*kendraiface.KendraAPI)(nil), configurators...)
	return result
}

func (d *KendraDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *KendraDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "BatchDeleteDocumentWithContext":
		return m.Fake(d.fakeBatchDeleteDocumentWithContext)

	case "BatchPutDocumentWithContext":
		return m.Fake(d.fakeBatchPutDocumentWithContext)

	case "CreateDataSourceWithContext":
		return m.Fake(d.fakeCreateDataSourceWithContext)

	case "CreateFaqWithContext":
		return m.Fake(d.fakeCreateFaqWithContext)

	case "CreateIndexWithContext":
		return m.Fake(d.fakeCreateIndexWithContext)

	case "DeleteFaqWithContext":
		return m.Fake(d.fakeDeleteFaqWithContext)

	case "DeleteIndexWithContext":
		return m.Fake(d.fakeDeleteIndexWithContext)

	case "DescribeDataSourceWithContext":
		return m.Fake(d.fakeDescribeDataSourceWithContext)

	case "DescribeFaqWithContext":
		return m.Fake(d.fakeDescribeFaqWithContext)

	case "DescribeIndexWithContext":
		return m.Fake(d.fakeDescribeIndexWithContext)

	case "ListDataSourceSyncJobsPages":
		return m.Fake(d.fakeListDataSourceSyncJobsPages)

	case "ListDataSourceSyncJobsPagesWithContext":
		return m.Fake(d.fakeListDataSourceSyncJobsPagesWithContext)

	case "ListDataSourceSyncJobsWithContext":
		return m.Fake(d.fakeListDataSourceSyncJobsWithContext)

	case "ListDataSourcesPages":
		return m.Fake(d.fakeListDataSourcesPages)

	case "ListDataSourcesPagesWithContext":
		return m.Fake(d.fakeListDataSourcesPagesWithContext)

	case "ListDataSourcesWithContext":
		return m.Fake(d.fakeListDataSourcesWithContext)

	case "ListFaqsWithContext":
		return m.Fake(d.fakeListFaqsWithContext)

	case "ListIndicesPages":
		return m.Fake(d.fakeListIndicesPages)

	case "ListIndicesPagesWithContext":
		return m.Fake(d.fakeListIndicesPagesWithContext)

	case "ListIndicesWithContext":
		return m.Fake(d.fakeListIndicesWithContext)

	case "QueryWithContext":
		return m.Fake(d.fakeQueryWithContext)

	case "StartDataSourceSyncJobWithContext":
		return m.Fake(d.fakeStartDataSourceSyncJobWithContext)

	case "StopDataSourceSyncJobWithContext":
		return m.Fake(d.fakeStopDataSourceSyncJobWithContext)

	case "SubmitFeedbackWithContext":
		return m.Fake(d.fakeSubmitFeedbackWithContext)

	case "UpdateDataSourceWithContext":
		return m.Fake(d.fakeUpdateDataSourceWithContext)

	case "UpdateIndexWithContext":
		return m.Fake(d.fakeUpdateIndexWithContext)

	default:
		return nil
	}
}

func (d *KendraDouble) BatchDeleteDocument(i0 *kendra.BatchDeleteDocumentInput) (r0 *kendra.BatchDeleteDocumentOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchDeleteDocument", i0)
	r0, _ = returns[0].(*kendra.BatchDeleteDocumentOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) BatchDeleteDocumentRequest(i0 *kendra.BatchDeleteDocumentInput) (r0 *request.Request, r1 *kendra.BatchDeleteDocumentOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchDeleteDocumentRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.BatchDeleteDocumentOutput)
	return
}

func (d *KendraDouble) BatchDeleteDocumentWithContext(i0 context.Context, i1 *kendra.BatchDeleteDocumentInput, i2 ...request.Option) (r0 *kendra.BatchDeleteDocumentOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchDeleteDocumentWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.BatchDeleteDocumentOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeBatchDeleteDocumentWithContext(ctx context.Context, in *kendra.BatchDeleteDocumentInput, _ ...request.Option) (*kendra.BatchDeleteDocumentOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "BatchDeleteDocument cancelled", ctx.Err())
	default:
		return d.BatchDeleteDocument(in)
	}
}

func (d *KendraDouble) BatchPutDocument(i0 *kendra.BatchPutDocumentInput) (r0 *kendra.BatchPutDocumentOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchPutDocument", i0)
	r0, _ = returns[0].(*kendra.BatchPutDocumentOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) BatchPutDocumentRequest(i0 *kendra.BatchPutDocumentInput) (r0 *request.Request, r1 *kendra.BatchPutDocumentOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchPutDocumentRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.BatchPutDocumentOutput)
	return
}

func (d *KendraDouble) BatchPutDocumentWithContext(i0 context.Context, i1 *kendra.BatchPutDocumentInput, i2 ...request.Option) (r0 *kendra.BatchPutDocumentOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("BatchPutDocumentWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.BatchPutDocumentOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeBatchPutDocumentWithContext(ctx context.Context, in *kendra.BatchPutDocumentInput, _ ...request.Option) (*kendra.BatchPutDocumentOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "BatchPutDocument cancelled", ctx.Err())
	default:
		return d.BatchPutDocument(in)
	}
}

func (d *KendraDouble) CreateDataSource(i0 *kendra.CreateDataSourceInput) (r0 *kendra.CreateDataSourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateDataSource", i0)
	r0, _ = returns[0].(*kendra.CreateDataSourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) CreateDataSourceRequest(i0 *kendra.CreateDataSourceInput) (r0 *request.Request, r1 *kendra.CreateDataSourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateDataSourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.CreateDataSourceOutput)
	return
}

func (d *KendraDouble) CreateDataSourceWithContext(i0 context.Context, i1 *kendra.CreateDataSourceInput, i2 ...request.Option) (r0 *kendra.CreateDataSourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateDataSourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.CreateDataSourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeCreateDataSourceWithContext(ctx context.Context, in *kendra.CreateDataSourceInput, _ ...request.Option) (*kendra.CreateDataSourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateDataSource cancelled", ctx.Err())
	default:
		return d.CreateDataSource(in)
	}
}

func (d *KendraDouble) CreateFaq(i0 *kendra.CreateFaqInput) (r0 *kendra.CreateFaqOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateFaq", i0)
	r0, _ = returns[0].(*kendra.CreateFaqOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) CreateFaqRequest(i0 *kendra.CreateFaqInput) (r0 *request.Request, r1 *kendra.CreateFaqOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateFaqRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.CreateFaqOutput)
	return
}

func (d *KendraDouble) CreateFaqWithContext(i0 context.Context, i1 *kendra.CreateFaqInput, i2 ...request.Option) (r0 *kendra.CreateFaqOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateFaqWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.CreateFaqOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeCreateFaqWithContext(ctx context.Context, in *kendra.CreateFaqInput, _ ...request.Option) (*kendra.CreateFaqOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateFaq cancelled", ctx.Err())
	default:
		return d.CreateFaq(in)
	}
}

func (d *KendraDouble) CreateIndex(i0 *kendra.CreateIndexInput) (r0 *kendra.CreateIndexOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateIndex", i0)
	r0, _ = returns[0].(*kendra.CreateIndexOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) CreateIndexRequest(i0 *kendra.CreateIndexInput) (r0 *request.Request, r1 *kendra.CreateIndexOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateIndexRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.CreateIndexOutput)
	return
}

func (d *KendraDouble) CreateIndexWithContext(i0 context.Context, i1 *kendra.CreateIndexInput, i2 ...request.Option) (r0 *kendra.CreateIndexOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateIndexWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.CreateIndexOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeCreateIndexWithContext(ctx context.Context, in *kendra.CreateIndexInput, _ ...request.Option) (*kendra.CreateIndexOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateIndex cancelled", ctx.Err())
	default:
		return d.CreateIndex(in)
	}
}

func (d *KendraDouble) DeleteFaq(i0 *kendra.DeleteFaqInput) (r0 *kendra.DeleteFaqOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteFaq", i0)
	r0, _ = returns[0].(*kendra.DeleteFaqOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) DeleteFaqRequest(i0 *kendra.DeleteFaqInput) (r0 *request.Request, r1 *kendra.DeleteFaqOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteFaqRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.DeleteFaqOutput)
	return
}

func (d *KendraDouble) DeleteFaqWithContext(i0 context.Context, i1 *kendra.DeleteFaqInput, i2 ...request.Option) (r0 *kendra.DeleteFaqOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteFaqWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.DeleteFaqOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeDeleteFaqWithContext(ctx context.Context, in *kendra.DeleteFaqInput, _ ...request.Option) (*kendra.DeleteFaqOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteFaq cancelled", ctx.Err())
	default:
		return d.DeleteFaq(in)
	}
}

func (d *KendraDouble) DeleteIndex(i0 *kendra.DeleteIndexInput) (r0 *kendra.DeleteIndexOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteIndex", i0)
	r0, _ = returns[0].(*kendra.DeleteIndexOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) DeleteIndexRequest(i0 *kendra.DeleteIndexInput) (r0 *request.Request, r1 *kendra.DeleteIndexOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteIndexRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.DeleteIndexOutput)
	return
}

func (d *KendraDouble) DeleteIndexWithContext(i0 context.Context, i1 *kendra.DeleteIndexInput, i2 ...request.Option) (r0 *kendra.DeleteIndexOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteIndexWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.DeleteIndexOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeDeleteIndexWithContext(ctx context.Context, in *kendra.DeleteIndexInput, _ ...request.Option) (*kendra.DeleteIndexOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteIndex cancelled", ctx.Err())
	default:
		return d.DeleteIndex(in)
	}
}

func (d *KendraDouble) DescribeDataSource(i0 *kendra.DescribeDataSourceInput) (r0 *kendra.DescribeDataSourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeDataSource", i0)
	r0, _ = returns[0].(*kendra.DescribeDataSourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) DescribeDataSourceRequest(i0 *kendra.DescribeDataSourceInput) (r0 *request.Request, r1 *kendra.DescribeDataSourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeDataSourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.DescribeDataSourceOutput)
	return
}

func (d *KendraDouble) DescribeDataSourceWithContext(i0 context.Context, i1 *kendra.DescribeDataSourceInput, i2 ...request.Option) (r0 *kendra.DescribeDataSourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeDataSourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.DescribeDataSourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeDescribeDataSourceWithContext(ctx context.Context, in *kendra.DescribeDataSourceInput, _ ...request.Option) (*kendra.DescribeDataSourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeDataSource cancelled", ctx.Err())
	default:
		return d.DescribeDataSource(in)
	}
}

func (d *KendraDouble) DescribeFaq(i0 *kendra.DescribeFaqInput) (r0 *kendra.DescribeFaqOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeFaq", i0)
	r0, _ = returns[0].(*kendra.DescribeFaqOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) DescribeFaqRequest(i0 *kendra.DescribeFaqInput) (r0 *request.Request, r1 *kendra.DescribeFaqOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeFaqRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.DescribeFaqOutput)
	return
}

func (d *KendraDouble) DescribeFaqWithContext(i0 context.Context, i1 *kendra.DescribeFaqInput, i2 ...request.Option) (r0 *kendra.DescribeFaqOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeFaqWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.DescribeFaqOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeDescribeFaqWithContext(ctx context.Context, in *kendra.DescribeFaqInput, _ ...request.Option) (*kendra.DescribeFaqOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeFaq cancelled", ctx.Err())
	default:
		return d.DescribeFaq(in)
	}
}

func (d *KendraDouble) DescribeIndex(i0 *kendra.DescribeIndexInput) (r0 *kendra.DescribeIndexOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeIndex", i0)
	r0, _ = returns[0].(*kendra.DescribeIndexOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) DescribeIndexRequest(i0 *kendra.DescribeIndexInput) (r0 *request.Request, r1 *kendra.DescribeIndexOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeIndexRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.DescribeIndexOutput)
	return
}

func (d *KendraDouble) DescribeIndexWithContext(i0 context.Context, i1 *kendra.DescribeIndexInput, i2 ...request.Option) (r0 *kendra.DescribeIndexOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeIndexWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.DescribeIndexOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeDescribeIndexWithContext(ctx context.Context, in *kendra.DescribeIndexInput, _ ...request.Option) (*kendra.DescribeIndexOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeIndex cancelled", ctx.Err())
	default:
		return d.DescribeIndex(in)
	}
}

func (d *KendraDouble) ListDataSourceSyncJobs(i0 *kendra.ListDataSourceSyncJobsInput) (r0 *kendra.ListDataSourceSyncJobsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListDataSourceSyncJobs", i0)
	r0, _ = returns[0].(*kendra.ListDataSourceSyncJobsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) ListDataSourceSyncJobsPages(i0 *kendra.ListDataSourceSyncJobsInput, i1 func(*kendra.ListDataSourceSyncJobsOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListDataSourceSyncJobsPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *KendraDouble) fakeListDataSourceSyncJobsPages(in *kendra.ListDataSourceSyncJobsInput, pager func(*kendra.ListDataSourceSyncJobsOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListDataSourceSyncJobs", paginators, in, pager)
}

func (d *KendraDouble) ListDataSourceSyncJobsPagesWithContext(i0 context.Context, i1 *kendra.ListDataSourceSyncJobsInput, i2 func(*kendra.ListDataSourceSyncJobsOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListDataSourceSyncJobsPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *KendraDouble) fakeListDataSourceSyncJobsPagesWithContext(ctx context.Context, in *kendra.ListDataSourceSyncJobsInput, pager func(*kendra.ListDataSourceSyncJobsOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListDataSourceSyncJobs", paginators, ctx, in, pager, options...)
}

func (d *KendraDouble) ListDataSourceSyncJobsRequest(i0 *kendra.ListDataSourceSyncJobsInput) (r0 *request.Request, r1 *kendra.ListDataSourceSyncJobsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListDataSourceSyncJobsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.ListDataSourceSyncJobsOutput)
	return
}

func (d *KendraDouble) ListDataSourceSyncJobsWithContext(i0 context.Context, i1 *kendra.ListDataSourceSyncJobsInput, i2 ...request.Option) (r0 *kendra.ListDataSourceSyncJobsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListDataSourceSyncJobsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.ListDataSourceSyncJobsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeListDataSourceSyncJobsWithContext(ctx context.Context, in *kendra.ListDataSourceSyncJobsInput, _ ...request.Option) (*kendra.ListDataSourceSyncJobsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListDataSourceSyncJobs cancelled", ctx.Err())
	default:
		return d.ListDataSourceSyncJobs(in)
	}
}

func (d *KendraDouble) ListDataSources(i0 *kendra.ListDataSourcesInput) (r0 *kendra.ListDataSourcesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListDataSources", i0)
	r0, _ = returns[0].(*kendra.ListDataSourcesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) ListDataSourcesPages(i0 *kendra.ListDataSourcesInput, i1 func(*kendra.ListDataSourcesOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListDataSourcesPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *KendraDouble) fakeListDataSourcesPages(in *kendra.ListDataSourcesInput, pager func(*kendra.ListDataSourcesOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListDataSources", paginators, in, pager)
}

func (d *KendraDouble) ListDataSourcesPagesWithContext(i0 context.Context, i1 *kendra.ListDataSourcesInput, i2 func(*kendra.ListDataSourcesOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListDataSourcesPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *KendraDouble) fakeListDataSourcesPagesWithContext(ctx context.Context, in *kendra.ListDataSourcesInput, pager func(*kendra.ListDataSourcesOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListDataSources", paginators, ctx, in, pager, options...)
}

func (d *KendraDouble) ListDataSourcesRequest(i0 *kendra.ListDataSourcesInput) (r0 *request.Request, r1 *kendra.ListDataSourcesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListDataSourcesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.ListDataSourcesOutput)
	return
}

func (d *KendraDouble) ListDataSourcesWithContext(i0 context.Context, i1 *kendra.ListDataSourcesInput, i2 ...request.Option) (r0 *kendra.ListDataSourcesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListDataSourcesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.ListDataSourcesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeListDataSourcesWithContext(ctx context.Context, in *kendra.ListDataSourcesInput, _ ...request.Option) (*kendra.ListDataSourcesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListDataSources cancelled", ctx.Err())
	default:
		return d.ListDataSources(in)
	}
}

func (d *KendraDouble) ListFaqs(i0 *kendra.ListFaqsInput) (r0 *kendra.ListFaqsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListFaqs", i0)
	r0, _ = returns[0].(*kendra.ListFaqsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) ListFaqsRequest(i0 *kendra.ListFaqsInput) (r0 *request.Request, r1 *kendra.ListFaqsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListFaqsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.ListFaqsOutput)
	return
}

func (d *KendraDouble) ListFaqsWithContext(i0 context.Context, i1 *kendra.ListFaqsInput, i2 ...request.Option) (r0 *kendra.ListFaqsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListFaqsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.ListFaqsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeListFaqsWithContext(ctx context.Context, in *kendra.ListFaqsInput, _ ...request.Option) (*kendra.ListFaqsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListFaqs cancelled", ctx.Err())
	default:
		return d.ListFaqs(in)
	}
}

func (d *KendraDouble) ListIndices(i0 *kendra.ListIndicesInput) (r0 *kendra.ListIndicesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListIndices", i0)
	r0, _ = returns[0].(*kendra.ListIndicesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) ListIndicesPages(i0 *kendra.ListIndicesInput, i1 func(*kendra.ListIndicesOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListIndicesPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *KendraDouble) fakeListIndicesPages(in *kendra.ListIndicesInput, pager func(*kendra.ListIndicesOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("ListIndices", paginators, in, pager)
}

func (d *KendraDouble) ListIndicesPagesWithContext(i0 context.Context, i1 *kendra.ListIndicesInput, i2 func(*kendra.ListIndicesOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListIndicesPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *KendraDouble) fakeListIndicesPagesWithContext(ctx context.Context, in *kendra.ListIndicesInput, pager func(*kendra.ListIndicesOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("ListIndices", paginators, ctx, in, pager, options...)
}

func (d *KendraDouble) ListIndicesRequest(i0 *kendra.ListIndicesInput) (r0 *request.Request, r1 *kendra.ListIndicesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListIndicesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.ListIndicesOutput)
	return
}

func (d *KendraDouble) ListIndicesWithContext(i0 context.Context, i1 *kendra.ListIndicesInput, i2 ...request.Option) (r0 *kendra.ListIndicesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListIndicesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.ListIndicesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeListIndicesWithContext(ctx context.Context, in *kendra.ListIndicesInput, _ ...request.Option) (*kendra.ListIndicesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListIndices cancelled", ctx.Err())
	default:
		return d.ListIndices(in)
	}
}

func (d *KendraDouble) Query(i0 *kendra.QueryInput) (r0 *kendra.QueryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("Query", i0)
	r0, _ = returns[0].(*kendra.QueryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) QueryRequest(i0 *kendra.QueryInput) (r0 *request.Request, r1 *kendra.QueryOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("QueryRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.QueryOutput)
	return
}

func (d *KendraDouble) QueryWithContext(i0 context.Context, i1 *kendra.QueryInput, i2 ...request.Option) (r0 *kendra.QueryOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("QueryWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.QueryOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeQueryWithContext(ctx context.Context, in *kendra.QueryInput, _ ...request.Option) (*kendra.QueryOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "Query cancelled", ctx.Err())
	default:
		return d.Query(in)
	}
}

func (d *KendraDouble) StartDataSourceSyncJob(i0 *kendra.StartDataSourceSyncJobInput) (r0 *kendra.StartDataSourceSyncJobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartDataSourceSyncJob", i0)
	r0, _ = returns[0].(*kendra.StartDataSourceSyncJobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) StartDataSourceSyncJobRequest(i0 *kendra.StartDataSourceSyncJobInput) (r0 *request.Request, r1 *kendra.StartDataSourceSyncJobOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartDataSourceSyncJobRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.StartDataSourceSyncJobOutput)
	return
}

func (d *KendraDouble) StartDataSourceSyncJobWithContext(i0 context.Context, i1 *kendra.StartDataSourceSyncJobInput, i2 ...request.Option) (r0 *kendra.StartDataSourceSyncJobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartDataSourceSyncJobWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.StartDataSourceSyncJobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeStartDataSourceSyncJobWithContext(ctx context.Context, in *kendra.StartDataSourceSyncJobInput, _ ...request.Option) (*kendra.StartDataSourceSyncJobOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StartDataSourceSyncJob cancelled", ctx.Err())
	default:
		return d.StartDataSourceSyncJob(in)
	}
}

func (d *KendraDouble) StopDataSourceSyncJob(i0 *kendra.StopDataSourceSyncJobInput) (r0 *kendra.StopDataSourceSyncJobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopDataSourceSyncJob", i0)
	r0, _ = returns[0].(*kendra.StopDataSourceSyncJobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) StopDataSourceSyncJobRequest(i0 *kendra.StopDataSourceSyncJobInput) (r0 *request.Request, r1 *kendra.StopDataSourceSyncJobOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopDataSourceSyncJobRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.StopDataSourceSyncJobOutput)
	return
}

func (d *KendraDouble) StopDataSourceSyncJobWithContext(i0 context.Context, i1 *kendra.StopDataSourceSyncJobInput, i2 ...request.Option) (r0 *kendra.StopDataSourceSyncJobOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopDataSourceSyncJobWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.StopDataSourceSyncJobOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeStopDataSourceSyncJobWithContext(ctx context.Context, in *kendra.StopDataSourceSyncJobInput, _ ...request.Option) (*kendra.StopDataSourceSyncJobOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StopDataSourceSyncJob cancelled", ctx.Err())
	default:
		return d.StopDataSourceSyncJob(in)
	}
}

func (d *KendraDouble) SubmitFeedback(i0 *kendra.SubmitFeedbackInput) (r0 *kendra.SubmitFeedbackOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("SubmitFeedback", i0)
	r0, _ = returns[0].(*kendra.SubmitFeedbackOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) SubmitFeedbackRequest(i0 *kendra.SubmitFeedbackInput) (r0 *request.Request, r1 *kendra.SubmitFeedbackOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("SubmitFeedbackRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.SubmitFeedbackOutput)
	return
}

func (d *KendraDouble) SubmitFeedbackWithContext(i0 context.Context, i1 *kendra.SubmitFeedbackInput, i2 ...request.Option) (r0 *kendra.SubmitFeedbackOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("SubmitFeedbackWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.SubmitFeedbackOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeSubmitFeedbackWithContext(ctx context.Context, in *kendra.SubmitFeedbackInput, _ ...request.Option) (*kendra.SubmitFeedbackOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "SubmitFeedback cancelled", ctx.Err())
	default:
		return d.SubmitFeedback(in)
	}
}

func (d *KendraDouble) UpdateDataSource(i0 *kendra.UpdateDataSourceInput) (r0 *kendra.UpdateDataSourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateDataSource", i0)
	r0, _ = returns[0].(*kendra.UpdateDataSourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) UpdateDataSourceRequest(i0 *kendra.UpdateDataSourceInput) (r0 *request.Request, r1 *kendra.UpdateDataSourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateDataSourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.UpdateDataSourceOutput)
	return
}

func (d *KendraDouble) UpdateDataSourceWithContext(i0 context.Context, i1 *kendra.UpdateDataSourceInput, i2 ...request.Option) (r0 *kendra.UpdateDataSourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateDataSourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.UpdateDataSourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeUpdateDataSourceWithContext(ctx context.Context, in *kendra.UpdateDataSourceInput, _ ...request.Option) (*kendra.UpdateDataSourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateDataSource cancelled", ctx.Err())
	default:
		return d.UpdateDataSource(in)
	}
}

func (d *KendraDouble) UpdateIndex(i0 *kendra.UpdateIndexInput) (r0 *kendra.UpdateIndexOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateIndex", i0)
	r0, _ = returns[0].(*kendra.UpdateIndexOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) UpdateIndexRequest(i0 *kendra.UpdateIndexInput) (r0 *request.Request, r1 *kendra.UpdateIndexOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateIndexRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kendra.UpdateIndexOutput)
	return
}

func (d *KendraDouble) UpdateIndexWithContext(i0 context.Context, i1 *kendra.UpdateIndexInput, i2 ...request.Option) (r0 *kendra.UpdateIndexOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateIndexWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kendra.UpdateIndexOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KendraDouble) fakeUpdateIndexWithContext(ctx context.Context, in *kendra.UpdateIndexInput, _ ...request.Option) (*kendra.UpdateIndexOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateIndex cancelled", ctx.Err())
	default:
		return d.UpdateIndex(in)
	}
}
