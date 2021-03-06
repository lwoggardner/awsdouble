// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:13:26+11:00

// Package kinesisanalyticsdouble provides a TestDouble implementation of kinesisanalyticsiface.KinesisAnalyticsAPI
package kinesisanalyticsdouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics/kinesisanalyticsiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// KinesisAnalyticsDouble is TestDouble for kinesisanalyticsiface.KinesisAnalyticsAPI
type KinesisAnalyticsDouble struct {
	kinesisanalyticsiface.KinesisAnalyticsAPI
	*awsdouble.AWSTestDouble
}

// Constructor for KinesisAnalyticsDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *KinesisAnalyticsDouble {
	result := &KinesisAnalyticsDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*kinesisanalyticsiface.KinesisAnalyticsAPI)(nil), configurators...)
	return result
}

func (d *KinesisAnalyticsDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *KinesisAnalyticsDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "AddApplicationCloudWatchLoggingOptionWithContext":
		return m.Fake(d.fakeAddApplicationCloudWatchLoggingOptionWithContext)

	case "AddApplicationInputProcessingConfigurationWithContext":
		return m.Fake(d.fakeAddApplicationInputProcessingConfigurationWithContext)

	case "AddApplicationInputWithContext":
		return m.Fake(d.fakeAddApplicationInputWithContext)

	case "AddApplicationOutputWithContext":
		return m.Fake(d.fakeAddApplicationOutputWithContext)

	case "AddApplicationReferenceDataSourceWithContext":
		return m.Fake(d.fakeAddApplicationReferenceDataSourceWithContext)

	case "CreateApplicationWithContext":
		return m.Fake(d.fakeCreateApplicationWithContext)

	case "DeleteApplicationCloudWatchLoggingOptionWithContext":
		return m.Fake(d.fakeDeleteApplicationCloudWatchLoggingOptionWithContext)

	case "DeleteApplicationInputProcessingConfigurationWithContext":
		return m.Fake(d.fakeDeleteApplicationInputProcessingConfigurationWithContext)

	case "DeleteApplicationOutputWithContext":
		return m.Fake(d.fakeDeleteApplicationOutputWithContext)

	case "DeleteApplicationReferenceDataSourceWithContext":
		return m.Fake(d.fakeDeleteApplicationReferenceDataSourceWithContext)

	case "DeleteApplicationWithContext":
		return m.Fake(d.fakeDeleteApplicationWithContext)

	case "DescribeApplicationWithContext":
		return m.Fake(d.fakeDescribeApplicationWithContext)

	case "DiscoverInputSchemaWithContext":
		return m.Fake(d.fakeDiscoverInputSchemaWithContext)

	case "ListApplicationsWithContext":
		return m.Fake(d.fakeListApplicationsWithContext)

	case "ListTagsForResourceWithContext":
		return m.Fake(d.fakeListTagsForResourceWithContext)

	case "StartApplicationWithContext":
		return m.Fake(d.fakeStartApplicationWithContext)

	case "StopApplicationWithContext":
		return m.Fake(d.fakeStopApplicationWithContext)

	case "TagResourceWithContext":
		return m.Fake(d.fakeTagResourceWithContext)

	case "UntagResourceWithContext":
		return m.Fake(d.fakeUntagResourceWithContext)

	case "UpdateApplicationWithContext":
		return m.Fake(d.fakeUpdateApplicationWithContext)

	default:
		return nil
	}
}

func (d *KinesisAnalyticsDouble) AddApplicationCloudWatchLoggingOption(i0 *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) (r0 *kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddApplicationCloudWatchLoggingOption", i0)
	r0, _ = returns[0].(*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) AddApplicationCloudWatchLoggingOptionRequest(i0 *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) (r0 *request.Request, r1 *kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddApplicationCloudWatchLoggingOptionRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput)
	return
}

func (d *KinesisAnalyticsDouble) AddApplicationCloudWatchLoggingOptionWithContext(i0 context.Context, i1 *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput, i2 ...request.Option) (r0 *kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddApplicationCloudWatchLoggingOptionWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeAddApplicationCloudWatchLoggingOptionWithContext(ctx context.Context, in *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput, _ ...request.Option) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "AddApplicationCloudWatchLoggingOption cancelled", ctx.Err())
	default:
		return d.AddApplicationCloudWatchLoggingOption(in)
	}
}

func (d *KinesisAnalyticsDouble) AddApplicationInput(i0 *kinesisanalytics.AddApplicationInputInput) (r0 *kinesisanalytics.AddApplicationInputOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddApplicationInput", i0)
	r0, _ = returns[0].(*kinesisanalytics.AddApplicationInputOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) AddApplicationInputProcessingConfiguration(i0 *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) (r0 *kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddApplicationInputProcessingConfiguration", i0)
	r0, _ = returns[0].(*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) AddApplicationInputProcessingConfigurationRequest(i0 *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) (r0 *request.Request, r1 *kinesisanalytics.AddApplicationInputProcessingConfigurationOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddApplicationInputProcessingConfigurationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput)
	return
}

func (d *KinesisAnalyticsDouble) AddApplicationInputProcessingConfigurationWithContext(i0 context.Context, i1 *kinesisanalytics.AddApplicationInputProcessingConfigurationInput, i2 ...request.Option) (r0 *kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddApplicationInputProcessingConfigurationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeAddApplicationInputProcessingConfigurationWithContext(ctx context.Context, in *kinesisanalytics.AddApplicationInputProcessingConfigurationInput, _ ...request.Option) (*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "AddApplicationInputProcessingConfiguration cancelled", ctx.Err())
	default:
		return d.AddApplicationInputProcessingConfiguration(in)
	}
}

func (d *KinesisAnalyticsDouble) AddApplicationInputRequest(i0 *kinesisanalytics.AddApplicationInputInput) (r0 *request.Request, r1 *kinesisanalytics.AddApplicationInputOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddApplicationInputRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.AddApplicationInputOutput)
	return
}

func (d *KinesisAnalyticsDouble) AddApplicationInputWithContext(i0 context.Context, i1 *kinesisanalytics.AddApplicationInputInput, i2 ...request.Option) (r0 *kinesisanalytics.AddApplicationInputOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddApplicationInputWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.AddApplicationInputOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeAddApplicationInputWithContext(ctx context.Context, in *kinesisanalytics.AddApplicationInputInput, _ ...request.Option) (*kinesisanalytics.AddApplicationInputOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "AddApplicationInput cancelled", ctx.Err())
	default:
		return d.AddApplicationInput(in)
	}
}

func (d *KinesisAnalyticsDouble) AddApplicationOutput(i0 *kinesisanalytics.AddApplicationOutputInput) (r0 *kinesisanalytics.AddApplicationOutputOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddApplicationOutput", i0)
	r0, _ = returns[0].(*kinesisanalytics.AddApplicationOutputOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) AddApplicationOutputRequest(i0 *kinesisanalytics.AddApplicationOutputInput) (r0 *request.Request, r1 *kinesisanalytics.AddApplicationOutputOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddApplicationOutputRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.AddApplicationOutputOutput)
	return
}

func (d *KinesisAnalyticsDouble) AddApplicationOutputWithContext(i0 context.Context, i1 *kinesisanalytics.AddApplicationOutputInput, i2 ...request.Option) (r0 *kinesisanalytics.AddApplicationOutputOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddApplicationOutputWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.AddApplicationOutputOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeAddApplicationOutputWithContext(ctx context.Context, in *kinesisanalytics.AddApplicationOutputInput, _ ...request.Option) (*kinesisanalytics.AddApplicationOutputOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "AddApplicationOutput cancelled", ctx.Err())
	default:
		return d.AddApplicationOutput(in)
	}
}

func (d *KinesisAnalyticsDouble) AddApplicationReferenceDataSource(i0 *kinesisanalytics.AddApplicationReferenceDataSourceInput) (r0 *kinesisanalytics.AddApplicationReferenceDataSourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddApplicationReferenceDataSource", i0)
	r0, _ = returns[0].(*kinesisanalytics.AddApplicationReferenceDataSourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) AddApplicationReferenceDataSourceRequest(i0 *kinesisanalytics.AddApplicationReferenceDataSourceInput) (r0 *request.Request, r1 *kinesisanalytics.AddApplicationReferenceDataSourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddApplicationReferenceDataSourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.AddApplicationReferenceDataSourceOutput)
	return
}

func (d *KinesisAnalyticsDouble) AddApplicationReferenceDataSourceWithContext(i0 context.Context, i1 *kinesisanalytics.AddApplicationReferenceDataSourceInput, i2 ...request.Option) (r0 *kinesisanalytics.AddApplicationReferenceDataSourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("AddApplicationReferenceDataSourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.AddApplicationReferenceDataSourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeAddApplicationReferenceDataSourceWithContext(ctx context.Context, in *kinesisanalytics.AddApplicationReferenceDataSourceInput, _ ...request.Option) (*kinesisanalytics.AddApplicationReferenceDataSourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "AddApplicationReferenceDataSource cancelled", ctx.Err())
	default:
		return d.AddApplicationReferenceDataSource(in)
	}
}

func (d *KinesisAnalyticsDouble) CreateApplication(i0 *kinesisanalytics.CreateApplicationInput) (r0 *kinesisanalytics.CreateApplicationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateApplication", i0)
	r0, _ = returns[0].(*kinesisanalytics.CreateApplicationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) CreateApplicationRequest(i0 *kinesisanalytics.CreateApplicationInput) (r0 *request.Request, r1 *kinesisanalytics.CreateApplicationOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateApplicationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.CreateApplicationOutput)
	return
}

func (d *KinesisAnalyticsDouble) CreateApplicationWithContext(i0 context.Context, i1 *kinesisanalytics.CreateApplicationInput, i2 ...request.Option) (r0 *kinesisanalytics.CreateApplicationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateApplicationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.CreateApplicationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeCreateApplicationWithContext(ctx context.Context, in *kinesisanalytics.CreateApplicationInput, _ ...request.Option) (*kinesisanalytics.CreateApplicationOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateApplication cancelled", ctx.Err())
	default:
		return d.CreateApplication(in)
	}
}

func (d *KinesisAnalyticsDouble) DeleteApplication(i0 *kinesisanalytics.DeleteApplicationInput) (r0 *kinesisanalytics.DeleteApplicationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteApplication", i0)
	r0, _ = returns[0].(*kinesisanalytics.DeleteApplicationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) DeleteApplicationCloudWatchLoggingOption(i0 *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) (r0 *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteApplicationCloudWatchLoggingOption", i0)
	r0, _ = returns[0].(*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) DeleteApplicationCloudWatchLoggingOptionRequest(i0 *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) (r0 *request.Request, r1 *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteApplicationCloudWatchLoggingOptionRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput)
	return
}

func (d *KinesisAnalyticsDouble) DeleteApplicationCloudWatchLoggingOptionWithContext(i0 context.Context, i1 *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput, i2 ...request.Option) (r0 *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteApplicationCloudWatchLoggingOptionWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeDeleteApplicationCloudWatchLoggingOptionWithContext(ctx context.Context, in *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput, _ ...request.Option) (*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteApplicationCloudWatchLoggingOption cancelled", ctx.Err())
	default:
		return d.DeleteApplicationCloudWatchLoggingOption(in)
	}
}

func (d *KinesisAnalyticsDouble) DeleteApplicationInputProcessingConfiguration(i0 *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) (r0 *kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteApplicationInputProcessingConfiguration", i0)
	r0, _ = returns[0].(*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) DeleteApplicationInputProcessingConfigurationRequest(i0 *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) (r0 *request.Request, r1 *kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteApplicationInputProcessingConfigurationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput)
	return
}

func (d *KinesisAnalyticsDouble) DeleteApplicationInputProcessingConfigurationWithContext(i0 context.Context, i1 *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput, i2 ...request.Option) (r0 *kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteApplicationInputProcessingConfigurationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeDeleteApplicationInputProcessingConfigurationWithContext(ctx context.Context, in *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput, _ ...request.Option) (*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteApplicationInputProcessingConfiguration cancelled", ctx.Err())
	default:
		return d.DeleteApplicationInputProcessingConfiguration(in)
	}
}

func (d *KinesisAnalyticsDouble) DeleteApplicationOutput(i0 *kinesisanalytics.DeleteApplicationOutputInput) (r0 *kinesisanalytics.DeleteApplicationOutputOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteApplicationOutput", i0)
	r0, _ = returns[0].(*kinesisanalytics.DeleteApplicationOutputOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) DeleteApplicationOutputRequest(i0 *kinesisanalytics.DeleteApplicationOutputInput) (r0 *request.Request, r1 *kinesisanalytics.DeleteApplicationOutputOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteApplicationOutputRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.DeleteApplicationOutputOutput)
	return
}

func (d *KinesisAnalyticsDouble) DeleteApplicationOutputWithContext(i0 context.Context, i1 *kinesisanalytics.DeleteApplicationOutputInput, i2 ...request.Option) (r0 *kinesisanalytics.DeleteApplicationOutputOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteApplicationOutputWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.DeleteApplicationOutputOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeDeleteApplicationOutputWithContext(ctx context.Context, in *kinesisanalytics.DeleteApplicationOutputInput, _ ...request.Option) (*kinesisanalytics.DeleteApplicationOutputOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteApplicationOutput cancelled", ctx.Err())
	default:
		return d.DeleteApplicationOutput(in)
	}
}

func (d *KinesisAnalyticsDouble) DeleteApplicationReferenceDataSource(i0 *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) (r0 *kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteApplicationReferenceDataSource", i0)
	r0, _ = returns[0].(*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) DeleteApplicationReferenceDataSourceRequest(i0 *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) (r0 *request.Request, r1 *kinesisanalytics.DeleteApplicationReferenceDataSourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteApplicationReferenceDataSourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput)
	return
}

func (d *KinesisAnalyticsDouble) DeleteApplicationReferenceDataSourceWithContext(i0 context.Context, i1 *kinesisanalytics.DeleteApplicationReferenceDataSourceInput, i2 ...request.Option) (r0 *kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteApplicationReferenceDataSourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeDeleteApplicationReferenceDataSourceWithContext(ctx context.Context, in *kinesisanalytics.DeleteApplicationReferenceDataSourceInput, _ ...request.Option) (*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteApplicationReferenceDataSource cancelled", ctx.Err())
	default:
		return d.DeleteApplicationReferenceDataSource(in)
	}
}

func (d *KinesisAnalyticsDouble) DeleteApplicationRequest(i0 *kinesisanalytics.DeleteApplicationInput) (r0 *request.Request, r1 *kinesisanalytics.DeleteApplicationOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteApplicationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.DeleteApplicationOutput)
	return
}

func (d *KinesisAnalyticsDouble) DeleteApplicationWithContext(i0 context.Context, i1 *kinesisanalytics.DeleteApplicationInput, i2 ...request.Option) (r0 *kinesisanalytics.DeleteApplicationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteApplicationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.DeleteApplicationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeDeleteApplicationWithContext(ctx context.Context, in *kinesisanalytics.DeleteApplicationInput, _ ...request.Option) (*kinesisanalytics.DeleteApplicationOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteApplication cancelled", ctx.Err())
	default:
		return d.DeleteApplication(in)
	}
}

func (d *KinesisAnalyticsDouble) DescribeApplication(i0 *kinesisanalytics.DescribeApplicationInput) (r0 *kinesisanalytics.DescribeApplicationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeApplication", i0)
	r0, _ = returns[0].(*kinesisanalytics.DescribeApplicationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) DescribeApplicationRequest(i0 *kinesisanalytics.DescribeApplicationInput) (r0 *request.Request, r1 *kinesisanalytics.DescribeApplicationOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeApplicationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.DescribeApplicationOutput)
	return
}

func (d *KinesisAnalyticsDouble) DescribeApplicationWithContext(i0 context.Context, i1 *kinesisanalytics.DescribeApplicationInput, i2 ...request.Option) (r0 *kinesisanalytics.DescribeApplicationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeApplicationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.DescribeApplicationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeDescribeApplicationWithContext(ctx context.Context, in *kinesisanalytics.DescribeApplicationInput, _ ...request.Option) (*kinesisanalytics.DescribeApplicationOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeApplication cancelled", ctx.Err())
	default:
		return d.DescribeApplication(in)
	}
}

func (d *KinesisAnalyticsDouble) DiscoverInputSchema(i0 *kinesisanalytics.DiscoverInputSchemaInput) (r0 *kinesisanalytics.DiscoverInputSchemaOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DiscoverInputSchema", i0)
	r0, _ = returns[0].(*kinesisanalytics.DiscoverInputSchemaOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) DiscoverInputSchemaRequest(i0 *kinesisanalytics.DiscoverInputSchemaInput) (r0 *request.Request, r1 *kinesisanalytics.DiscoverInputSchemaOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DiscoverInputSchemaRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.DiscoverInputSchemaOutput)
	return
}

func (d *KinesisAnalyticsDouble) DiscoverInputSchemaWithContext(i0 context.Context, i1 *kinesisanalytics.DiscoverInputSchemaInput, i2 ...request.Option) (r0 *kinesisanalytics.DiscoverInputSchemaOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DiscoverInputSchemaWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.DiscoverInputSchemaOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeDiscoverInputSchemaWithContext(ctx context.Context, in *kinesisanalytics.DiscoverInputSchemaInput, _ ...request.Option) (*kinesisanalytics.DiscoverInputSchemaOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DiscoverInputSchema cancelled", ctx.Err())
	default:
		return d.DiscoverInputSchema(in)
	}
}

func (d *KinesisAnalyticsDouble) ListApplications(i0 *kinesisanalytics.ListApplicationsInput) (r0 *kinesisanalytics.ListApplicationsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListApplications", i0)
	r0, _ = returns[0].(*kinesisanalytics.ListApplicationsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) ListApplicationsRequest(i0 *kinesisanalytics.ListApplicationsInput) (r0 *request.Request, r1 *kinesisanalytics.ListApplicationsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListApplicationsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.ListApplicationsOutput)
	return
}

func (d *KinesisAnalyticsDouble) ListApplicationsWithContext(i0 context.Context, i1 *kinesisanalytics.ListApplicationsInput, i2 ...request.Option) (r0 *kinesisanalytics.ListApplicationsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListApplicationsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.ListApplicationsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeListApplicationsWithContext(ctx context.Context, in *kinesisanalytics.ListApplicationsInput, _ ...request.Option) (*kinesisanalytics.ListApplicationsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListApplications cancelled", ctx.Err())
	default:
		return d.ListApplications(in)
	}
}

func (d *KinesisAnalyticsDouble) ListTagsForResource(i0 *kinesisanalytics.ListTagsForResourceInput) (r0 *kinesisanalytics.ListTagsForResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForResource", i0)
	r0, _ = returns[0].(*kinesisanalytics.ListTagsForResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) ListTagsForResourceRequest(i0 *kinesisanalytics.ListTagsForResourceInput) (r0 *request.Request, r1 *kinesisanalytics.ListTagsForResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.ListTagsForResourceOutput)
	return
}

func (d *KinesisAnalyticsDouble) ListTagsForResourceWithContext(i0 context.Context, i1 *kinesisanalytics.ListTagsForResourceInput, i2 ...request.Option) (r0 *kinesisanalytics.ListTagsForResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsForResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.ListTagsForResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeListTagsForResourceWithContext(ctx context.Context, in *kinesisanalytics.ListTagsForResourceInput, _ ...request.Option) (*kinesisanalytics.ListTagsForResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListTagsForResource cancelled", ctx.Err())
	default:
		return d.ListTagsForResource(in)
	}
}

func (d *KinesisAnalyticsDouble) StartApplication(i0 *kinesisanalytics.StartApplicationInput) (r0 *kinesisanalytics.StartApplicationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartApplication", i0)
	r0, _ = returns[0].(*kinesisanalytics.StartApplicationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) StartApplicationRequest(i0 *kinesisanalytics.StartApplicationInput) (r0 *request.Request, r1 *kinesisanalytics.StartApplicationOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartApplicationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.StartApplicationOutput)
	return
}

func (d *KinesisAnalyticsDouble) StartApplicationWithContext(i0 context.Context, i1 *kinesisanalytics.StartApplicationInput, i2 ...request.Option) (r0 *kinesisanalytics.StartApplicationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StartApplicationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.StartApplicationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeStartApplicationWithContext(ctx context.Context, in *kinesisanalytics.StartApplicationInput, _ ...request.Option) (*kinesisanalytics.StartApplicationOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StartApplication cancelled", ctx.Err())
	default:
		return d.StartApplication(in)
	}
}

func (d *KinesisAnalyticsDouble) StopApplication(i0 *kinesisanalytics.StopApplicationInput) (r0 *kinesisanalytics.StopApplicationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopApplication", i0)
	r0, _ = returns[0].(*kinesisanalytics.StopApplicationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) StopApplicationRequest(i0 *kinesisanalytics.StopApplicationInput) (r0 *request.Request, r1 *kinesisanalytics.StopApplicationOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopApplicationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.StopApplicationOutput)
	return
}

func (d *KinesisAnalyticsDouble) StopApplicationWithContext(i0 context.Context, i1 *kinesisanalytics.StopApplicationInput, i2 ...request.Option) (r0 *kinesisanalytics.StopApplicationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("StopApplicationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.StopApplicationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeStopApplicationWithContext(ctx context.Context, in *kinesisanalytics.StopApplicationInput, _ ...request.Option) (*kinesisanalytics.StopApplicationOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "StopApplication cancelled", ctx.Err())
	default:
		return d.StopApplication(in)
	}
}

func (d *KinesisAnalyticsDouble) TagResource(i0 *kinesisanalytics.TagResourceInput) (r0 *kinesisanalytics.TagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResource", i0)
	r0, _ = returns[0].(*kinesisanalytics.TagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) TagResourceRequest(i0 *kinesisanalytics.TagResourceInput) (r0 *request.Request, r1 *kinesisanalytics.TagResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.TagResourceOutput)
	return
}

func (d *KinesisAnalyticsDouble) TagResourceWithContext(i0 context.Context, i1 *kinesisanalytics.TagResourceInput, i2 ...request.Option) (r0 *kinesisanalytics.TagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("TagResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.TagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeTagResourceWithContext(ctx context.Context, in *kinesisanalytics.TagResourceInput, _ ...request.Option) (*kinesisanalytics.TagResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "TagResource cancelled", ctx.Err())
	default:
		return d.TagResource(in)
	}
}

func (d *KinesisAnalyticsDouble) UntagResource(i0 *kinesisanalytics.UntagResourceInput) (r0 *kinesisanalytics.UntagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResource", i0)
	r0, _ = returns[0].(*kinesisanalytics.UntagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) UntagResourceRequest(i0 *kinesisanalytics.UntagResourceInput) (r0 *request.Request, r1 *kinesisanalytics.UntagResourceOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResourceRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.UntagResourceOutput)
	return
}

func (d *KinesisAnalyticsDouble) UntagResourceWithContext(i0 context.Context, i1 *kinesisanalytics.UntagResourceInput, i2 ...request.Option) (r0 *kinesisanalytics.UntagResourceOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UntagResourceWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.UntagResourceOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeUntagResourceWithContext(ctx context.Context, in *kinesisanalytics.UntagResourceInput, _ ...request.Option) (*kinesisanalytics.UntagResourceOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UntagResource cancelled", ctx.Err())
	default:
		return d.UntagResource(in)
	}
}

func (d *KinesisAnalyticsDouble) UpdateApplication(i0 *kinesisanalytics.UpdateApplicationInput) (r0 *kinesisanalytics.UpdateApplicationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateApplication", i0)
	r0, _ = returns[0].(*kinesisanalytics.UpdateApplicationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) UpdateApplicationRequest(i0 *kinesisanalytics.UpdateApplicationInput) (r0 *request.Request, r1 *kinesisanalytics.UpdateApplicationOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateApplicationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*kinesisanalytics.UpdateApplicationOutput)
	return
}

func (d *KinesisAnalyticsDouble) UpdateApplicationWithContext(i0 context.Context, i1 *kinesisanalytics.UpdateApplicationInput, i2 ...request.Option) (r0 *kinesisanalytics.UpdateApplicationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateApplicationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*kinesisanalytics.UpdateApplicationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *KinesisAnalyticsDouble) fakeUpdateApplicationWithContext(ctx context.Context, in *kinesisanalytics.UpdateApplicationInput, _ ...request.Option) (*kinesisanalytics.UpdateApplicationOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateApplication cancelled", ctx.Err())
	default:
		return d.UpdateApplication(in)
	}
}
