// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:14:06+11:00

// Package mqdouble provides a TestDouble implementation of mqiface.MQAPI
package mqdouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/mq"
	"github.com/aws/aws-sdk-go/service/mq/mqiface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// MQDouble is TestDouble for mqiface.MQAPI
type MQDouble struct {
	mqiface.MQAPI
	*awsdouble.AWSTestDouble
}

// Constructor for MQDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *MQDouble {
	result := &MQDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*mqiface.MQAPI)(nil), configurators...)
	return result
}

func (d *MQDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *MQDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "CreateBrokerWithContext":
		return m.Fake(d.fakeCreateBrokerWithContext)

	case "CreateConfigurationWithContext":
		return m.Fake(d.fakeCreateConfigurationWithContext)

	case "CreateTagsWithContext":
		return m.Fake(d.fakeCreateTagsWithContext)

	case "CreateUserWithContext":
		return m.Fake(d.fakeCreateUserWithContext)

	case "DeleteBrokerWithContext":
		return m.Fake(d.fakeDeleteBrokerWithContext)

	case "DeleteTagsWithContext":
		return m.Fake(d.fakeDeleteTagsWithContext)

	case "DeleteUserWithContext":
		return m.Fake(d.fakeDeleteUserWithContext)

	case "DescribeBrokerEngineTypesWithContext":
		return m.Fake(d.fakeDescribeBrokerEngineTypesWithContext)

	case "DescribeBrokerInstanceOptionsWithContext":
		return m.Fake(d.fakeDescribeBrokerInstanceOptionsWithContext)

	case "DescribeBrokerWithContext":
		return m.Fake(d.fakeDescribeBrokerWithContext)

	case "DescribeConfigurationRevisionWithContext":
		return m.Fake(d.fakeDescribeConfigurationRevisionWithContext)

	case "DescribeConfigurationWithContext":
		return m.Fake(d.fakeDescribeConfigurationWithContext)

	case "DescribeUserWithContext":
		return m.Fake(d.fakeDescribeUserWithContext)

	case "ListBrokersWithContext":
		return m.Fake(d.fakeListBrokersWithContext)

	case "ListConfigurationRevisionsWithContext":
		return m.Fake(d.fakeListConfigurationRevisionsWithContext)

	case "ListConfigurationsWithContext":
		return m.Fake(d.fakeListConfigurationsWithContext)

	case "ListTagsWithContext":
		return m.Fake(d.fakeListTagsWithContext)

	case "ListUsersWithContext":
		return m.Fake(d.fakeListUsersWithContext)

	case "RebootBrokerWithContext":
		return m.Fake(d.fakeRebootBrokerWithContext)

	case "UpdateBrokerWithContext":
		return m.Fake(d.fakeUpdateBrokerWithContext)

	case "UpdateConfigurationWithContext":
		return m.Fake(d.fakeUpdateConfigurationWithContext)

	case "UpdateUserWithContext":
		return m.Fake(d.fakeUpdateUserWithContext)

	default:
		return nil
	}
}

func (d *MQDouble) CreateBroker(i0 *mq.CreateBrokerRequest) (r0 *mq.CreateBrokerResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateBroker", i0)
	r0, _ = returns[0].(*mq.CreateBrokerResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) CreateBrokerRequest(i0 *mq.CreateBrokerRequest) (r0 *request.Request, r1 *mq.CreateBrokerResponse) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateBrokerRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.CreateBrokerResponse)
	return
}

func (d *MQDouble) CreateBrokerWithContext(i0 context.Context, i1 *mq.CreateBrokerRequest, i2 ...request.Option) (r0 *mq.CreateBrokerResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateBrokerWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.CreateBrokerResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeCreateBrokerWithContext(ctx context.Context, in *mq.CreateBrokerRequest, _ ...request.Option) (*mq.CreateBrokerResponse, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateBroker cancelled", ctx.Err())
	default:
		return d.CreateBroker(in)
	}
}

func (d *MQDouble) CreateConfiguration(i0 *mq.CreateConfigurationRequest) (r0 *mq.CreateConfigurationResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateConfiguration", i0)
	r0, _ = returns[0].(*mq.CreateConfigurationResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) CreateConfigurationRequest(i0 *mq.CreateConfigurationRequest) (r0 *request.Request, r1 *mq.CreateConfigurationResponse) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateConfigurationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.CreateConfigurationResponse)
	return
}

func (d *MQDouble) CreateConfigurationWithContext(i0 context.Context, i1 *mq.CreateConfigurationRequest, i2 ...request.Option) (r0 *mq.CreateConfigurationResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateConfigurationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.CreateConfigurationResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeCreateConfigurationWithContext(ctx context.Context, in *mq.CreateConfigurationRequest, _ ...request.Option) (*mq.CreateConfigurationResponse, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateConfiguration cancelled", ctx.Err())
	default:
		return d.CreateConfiguration(in)
	}
}

func (d *MQDouble) CreateTags(i0 *mq.CreateTagsInput) (r0 *mq.CreateTagsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateTags", i0)
	r0, _ = returns[0].(*mq.CreateTagsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) CreateTagsRequest(i0 *mq.CreateTagsInput) (r0 *request.Request, r1 *mq.CreateTagsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateTagsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.CreateTagsOutput)
	return
}

func (d *MQDouble) CreateTagsWithContext(i0 context.Context, i1 *mq.CreateTagsInput, i2 ...request.Option) (r0 *mq.CreateTagsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateTagsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.CreateTagsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeCreateTagsWithContext(ctx context.Context, in *mq.CreateTagsInput, _ ...request.Option) (*mq.CreateTagsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateTags cancelled", ctx.Err())
	default:
		return d.CreateTags(in)
	}
}

func (d *MQDouble) CreateUser(i0 *mq.CreateUserRequest) (r0 *mq.CreateUserOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateUser", i0)
	r0, _ = returns[0].(*mq.CreateUserOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) CreateUserRequest(i0 *mq.CreateUserRequest) (r0 *request.Request, r1 *mq.CreateUserOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateUserRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.CreateUserOutput)
	return
}

func (d *MQDouble) CreateUserWithContext(i0 context.Context, i1 *mq.CreateUserRequest, i2 ...request.Option) (r0 *mq.CreateUserOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateUserWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.CreateUserOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeCreateUserWithContext(ctx context.Context, in *mq.CreateUserRequest, _ ...request.Option) (*mq.CreateUserOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateUser cancelled", ctx.Err())
	default:
		return d.CreateUser(in)
	}
}

func (d *MQDouble) DeleteBroker(i0 *mq.DeleteBrokerInput) (r0 *mq.DeleteBrokerResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteBroker", i0)
	r0, _ = returns[0].(*mq.DeleteBrokerResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) DeleteBrokerRequest(i0 *mq.DeleteBrokerInput) (r0 *request.Request, r1 *mq.DeleteBrokerResponse) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteBrokerRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.DeleteBrokerResponse)
	return
}

func (d *MQDouble) DeleteBrokerWithContext(i0 context.Context, i1 *mq.DeleteBrokerInput, i2 ...request.Option) (r0 *mq.DeleteBrokerResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteBrokerWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.DeleteBrokerResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeDeleteBrokerWithContext(ctx context.Context, in *mq.DeleteBrokerInput, _ ...request.Option) (*mq.DeleteBrokerResponse, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteBroker cancelled", ctx.Err())
	default:
		return d.DeleteBroker(in)
	}
}

func (d *MQDouble) DeleteTags(i0 *mq.DeleteTagsInput) (r0 *mq.DeleteTagsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteTags", i0)
	r0, _ = returns[0].(*mq.DeleteTagsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) DeleteTagsRequest(i0 *mq.DeleteTagsInput) (r0 *request.Request, r1 *mq.DeleteTagsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteTagsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.DeleteTagsOutput)
	return
}

func (d *MQDouble) DeleteTagsWithContext(i0 context.Context, i1 *mq.DeleteTagsInput, i2 ...request.Option) (r0 *mq.DeleteTagsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteTagsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.DeleteTagsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeDeleteTagsWithContext(ctx context.Context, in *mq.DeleteTagsInput, _ ...request.Option) (*mq.DeleteTagsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteTags cancelled", ctx.Err())
	default:
		return d.DeleteTags(in)
	}
}

func (d *MQDouble) DeleteUser(i0 *mq.DeleteUserInput) (r0 *mq.DeleteUserOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteUser", i0)
	r0, _ = returns[0].(*mq.DeleteUserOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) DeleteUserRequest(i0 *mq.DeleteUserInput) (r0 *request.Request, r1 *mq.DeleteUserOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteUserRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.DeleteUserOutput)
	return
}

func (d *MQDouble) DeleteUserWithContext(i0 context.Context, i1 *mq.DeleteUserInput, i2 ...request.Option) (r0 *mq.DeleteUserOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteUserWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.DeleteUserOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeDeleteUserWithContext(ctx context.Context, in *mq.DeleteUserInput, _ ...request.Option) (*mq.DeleteUserOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteUser cancelled", ctx.Err())
	default:
		return d.DeleteUser(in)
	}
}

func (d *MQDouble) DescribeBroker(i0 *mq.DescribeBrokerInput) (r0 *mq.DescribeBrokerResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeBroker", i0)
	r0, _ = returns[0].(*mq.DescribeBrokerResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) DescribeBrokerEngineTypes(i0 *mq.DescribeBrokerEngineTypesInput) (r0 *mq.DescribeBrokerEngineTypesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeBrokerEngineTypes", i0)
	r0, _ = returns[0].(*mq.DescribeBrokerEngineTypesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) DescribeBrokerEngineTypesRequest(i0 *mq.DescribeBrokerEngineTypesInput) (r0 *request.Request, r1 *mq.DescribeBrokerEngineTypesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeBrokerEngineTypesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.DescribeBrokerEngineTypesOutput)
	return
}

func (d *MQDouble) DescribeBrokerEngineTypesWithContext(i0 context.Context, i1 *mq.DescribeBrokerEngineTypesInput, i2 ...request.Option) (r0 *mq.DescribeBrokerEngineTypesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeBrokerEngineTypesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.DescribeBrokerEngineTypesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeDescribeBrokerEngineTypesWithContext(ctx context.Context, in *mq.DescribeBrokerEngineTypesInput, _ ...request.Option) (*mq.DescribeBrokerEngineTypesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeBrokerEngineTypes cancelled", ctx.Err())
	default:
		return d.DescribeBrokerEngineTypes(in)
	}
}

func (d *MQDouble) DescribeBrokerInstanceOptions(i0 *mq.DescribeBrokerInstanceOptionsInput) (r0 *mq.DescribeBrokerInstanceOptionsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeBrokerInstanceOptions", i0)
	r0, _ = returns[0].(*mq.DescribeBrokerInstanceOptionsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) DescribeBrokerInstanceOptionsRequest(i0 *mq.DescribeBrokerInstanceOptionsInput) (r0 *request.Request, r1 *mq.DescribeBrokerInstanceOptionsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeBrokerInstanceOptionsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.DescribeBrokerInstanceOptionsOutput)
	return
}

func (d *MQDouble) DescribeBrokerInstanceOptionsWithContext(i0 context.Context, i1 *mq.DescribeBrokerInstanceOptionsInput, i2 ...request.Option) (r0 *mq.DescribeBrokerInstanceOptionsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeBrokerInstanceOptionsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.DescribeBrokerInstanceOptionsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeDescribeBrokerInstanceOptionsWithContext(ctx context.Context, in *mq.DescribeBrokerInstanceOptionsInput, _ ...request.Option) (*mq.DescribeBrokerInstanceOptionsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeBrokerInstanceOptions cancelled", ctx.Err())
	default:
		return d.DescribeBrokerInstanceOptions(in)
	}
}

func (d *MQDouble) DescribeBrokerRequest(i0 *mq.DescribeBrokerInput) (r0 *request.Request, r1 *mq.DescribeBrokerResponse) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeBrokerRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.DescribeBrokerResponse)
	return
}

func (d *MQDouble) DescribeBrokerWithContext(i0 context.Context, i1 *mq.DescribeBrokerInput, i2 ...request.Option) (r0 *mq.DescribeBrokerResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeBrokerWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.DescribeBrokerResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeDescribeBrokerWithContext(ctx context.Context, in *mq.DescribeBrokerInput, _ ...request.Option) (*mq.DescribeBrokerResponse, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeBroker cancelled", ctx.Err())
	default:
		return d.DescribeBroker(in)
	}
}

func (d *MQDouble) DescribeConfiguration(i0 *mq.DescribeConfigurationInput) (r0 *mq.DescribeConfigurationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeConfiguration", i0)
	r0, _ = returns[0].(*mq.DescribeConfigurationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) DescribeConfigurationRequest(i0 *mq.DescribeConfigurationInput) (r0 *request.Request, r1 *mq.DescribeConfigurationOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeConfigurationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.DescribeConfigurationOutput)
	return
}

func (d *MQDouble) DescribeConfigurationRevision(i0 *mq.DescribeConfigurationRevisionInput) (r0 *mq.DescribeConfigurationRevisionResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeConfigurationRevision", i0)
	r0, _ = returns[0].(*mq.DescribeConfigurationRevisionResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) DescribeConfigurationRevisionRequest(i0 *mq.DescribeConfigurationRevisionInput) (r0 *request.Request, r1 *mq.DescribeConfigurationRevisionResponse) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeConfigurationRevisionRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.DescribeConfigurationRevisionResponse)
	return
}

func (d *MQDouble) DescribeConfigurationRevisionWithContext(i0 context.Context, i1 *mq.DescribeConfigurationRevisionInput, i2 ...request.Option) (r0 *mq.DescribeConfigurationRevisionResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeConfigurationRevisionWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.DescribeConfigurationRevisionResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeDescribeConfigurationRevisionWithContext(ctx context.Context, in *mq.DescribeConfigurationRevisionInput, _ ...request.Option) (*mq.DescribeConfigurationRevisionResponse, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeConfigurationRevision cancelled", ctx.Err())
	default:
		return d.DescribeConfigurationRevision(in)
	}
}

func (d *MQDouble) DescribeConfigurationWithContext(i0 context.Context, i1 *mq.DescribeConfigurationInput, i2 ...request.Option) (r0 *mq.DescribeConfigurationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeConfigurationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.DescribeConfigurationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeDescribeConfigurationWithContext(ctx context.Context, in *mq.DescribeConfigurationInput, _ ...request.Option) (*mq.DescribeConfigurationOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeConfiguration cancelled", ctx.Err())
	default:
		return d.DescribeConfiguration(in)
	}
}

func (d *MQDouble) DescribeUser(i0 *mq.DescribeUserInput) (r0 *mq.DescribeUserResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeUser", i0)
	r0, _ = returns[0].(*mq.DescribeUserResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) DescribeUserRequest(i0 *mq.DescribeUserInput) (r0 *request.Request, r1 *mq.DescribeUserResponse) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeUserRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.DescribeUserResponse)
	return
}

func (d *MQDouble) DescribeUserWithContext(i0 context.Context, i1 *mq.DescribeUserInput, i2 ...request.Option) (r0 *mq.DescribeUserResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeUserWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.DescribeUserResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeDescribeUserWithContext(ctx context.Context, in *mq.DescribeUserInput, _ ...request.Option) (*mq.DescribeUserResponse, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeUser cancelled", ctx.Err())
	default:
		return d.DescribeUser(in)
	}
}

func (d *MQDouble) ListBrokers(i0 *mq.ListBrokersInput) (r0 *mq.ListBrokersResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListBrokers", i0)
	r0, _ = returns[0].(*mq.ListBrokersResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) ListBrokersRequest(i0 *mq.ListBrokersInput) (r0 *request.Request, r1 *mq.ListBrokersResponse) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListBrokersRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.ListBrokersResponse)
	return
}

func (d *MQDouble) ListBrokersWithContext(i0 context.Context, i1 *mq.ListBrokersInput, i2 ...request.Option) (r0 *mq.ListBrokersResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListBrokersWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.ListBrokersResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeListBrokersWithContext(ctx context.Context, in *mq.ListBrokersInput, _ ...request.Option) (*mq.ListBrokersResponse, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListBrokers cancelled", ctx.Err())
	default:
		return d.ListBrokers(in)
	}
}

func (d *MQDouble) ListConfigurationRevisions(i0 *mq.ListConfigurationRevisionsInput) (r0 *mq.ListConfigurationRevisionsResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListConfigurationRevisions", i0)
	r0, _ = returns[0].(*mq.ListConfigurationRevisionsResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) ListConfigurationRevisionsRequest(i0 *mq.ListConfigurationRevisionsInput) (r0 *request.Request, r1 *mq.ListConfigurationRevisionsResponse) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListConfigurationRevisionsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.ListConfigurationRevisionsResponse)
	return
}

func (d *MQDouble) ListConfigurationRevisionsWithContext(i0 context.Context, i1 *mq.ListConfigurationRevisionsInput, i2 ...request.Option) (r0 *mq.ListConfigurationRevisionsResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListConfigurationRevisionsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.ListConfigurationRevisionsResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeListConfigurationRevisionsWithContext(ctx context.Context, in *mq.ListConfigurationRevisionsInput, _ ...request.Option) (*mq.ListConfigurationRevisionsResponse, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListConfigurationRevisions cancelled", ctx.Err())
	default:
		return d.ListConfigurationRevisions(in)
	}
}

func (d *MQDouble) ListConfigurations(i0 *mq.ListConfigurationsInput) (r0 *mq.ListConfigurationsResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListConfigurations", i0)
	r0, _ = returns[0].(*mq.ListConfigurationsResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) ListConfigurationsRequest(i0 *mq.ListConfigurationsInput) (r0 *request.Request, r1 *mq.ListConfigurationsResponse) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListConfigurationsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.ListConfigurationsResponse)
	return
}

func (d *MQDouble) ListConfigurationsWithContext(i0 context.Context, i1 *mq.ListConfigurationsInput, i2 ...request.Option) (r0 *mq.ListConfigurationsResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListConfigurationsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.ListConfigurationsResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeListConfigurationsWithContext(ctx context.Context, in *mq.ListConfigurationsInput, _ ...request.Option) (*mq.ListConfigurationsResponse, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListConfigurations cancelled", ctx.Err())
	default:
		return d.ListConfigurations(in)
	}
}

func (d *MQDouble) ListTags(i0 *mq.ListTagsInput) (r0 *mq.ListTagsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTags", i0)
	r0, _ = returns[0].(*mq.ListTagsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) ListTagsRequest(i0 *mq.ListTagsInput) (r0 *request.Request, r1 *mq.ListTagsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.ListTagsOutput)
	return
}

func (d *MQDouble) ListTagsWithContext(i0 context.Context, i1 *mq.ListTagsInput, i2 ...request.Option) (r0 *mq.ListTagsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListTagsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.ListTagsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeListTagsWithContext(ctx context.Context, in *mq.ListTagsInput, _ ...request.Option) (*mq.ListTagsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListTags cancelled", ctx.Err())
	default:
		return d.ListTags(in)
	}
}

func (d *MQDouble) ListUsers(i0 *mq.ListUsersInput) (r0 *mq.ListUsersResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListUsers", i0)
	r0, _ = returns[0].(*mq.ListUsersResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) ListUsersRequest(i0 *mq.ListUsersInput) (r0 *request.Request, r1 *mq.ListUsersResponse) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListUsersRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.ListUsersResponse)
	return
}

func (d *MQDouble) ListUsersWithContext(i0 context.Context, i1 *mq.ListUsersInput, i2 ...request.Option) (r0 *mq.ListUsersResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListUsersWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.ListUsersResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeListUsersWithContext(ctx context.Context, in *mq.ListUsersInput, _ ...request.Option) (*mq.ListUsersResponse, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListUsers cancelled", ctx.Err())
	default:
		return d.ListUsers(in)
	}
}

func (d *MQDouble) RebootBroker(i0 *mq.RebootBrokerInput) (r0 *mq.RebootBrokerOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("RebootBroker", i0)
	r0, _ = returns[0].(*mq.RebootBrokerOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) RebootBrokerRequest(i0 *mq.RebootBrokerInput) (r0 *request.Request, r1 *mq.RebootBrokerOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("RebootBrokerRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.RebootBrokerOutput)
	return
}

func (d *MQDouble) RebootBrokerWithContext(i0 context.Context, i1 *mq.RebootBrokerInput, i2 ...request.Option) (r0 *mq.RebootBrokerOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("RebootBrokerWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.RebootBrokerOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeRebootBrokerWithContext(ctx context.Context, in *mq.RebootBrokerInput, _ ...request.Option) (*mq.RebootBrokerOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "RebootBroker cancelled", ctx.Err())
	default:
		return d.RebootBroker(in)
	}
}

func (d *MQDouble) UpdateBroker(i0 *mq.UpdateBrokerRequest) (r0 *mq.UpdateBrokerResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateBroker", i0)
	r0, _ = returns[0].(*mq.UpdateBrokerResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) UpdateBrokerRequest(i0 *mq.UpdateBrokerRequest) (r0 *request.Request, r1 *mq.UpdateBrokerResponse) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateBrokerRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.UpdateBrokerResponse)
	return
}

func (d *MQDouble) UpdateBrokerWithContext(i0 context.Context, i1 *mq.UpdateBrokerRequest, i2 ...request.Option) (r0 *mq.UpdateBrokerResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateBrokerWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.UpdateBrokerResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeUpdateBrokerWithContext(ctx context.Context, in *mq.UpdateBrokerRequest, _ ...request.Option) (*mq.UpdateBrokerResponse, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateBroker cancelled", ctx.Err())
	default:
		return d.UpdateBroker(in)
	}
}

func (d *MQDouble) UpdateConfiguration(i0 *mq.UpdateConfigurationRequest) (r0 *mq.UpdateConfigurationResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateConfiguration", i0)
	r0, _ = returns[0].(*mq.UpdateConfigurationResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) UpdateConfigurationRequest(i0 *mq.UpdateConfigurationRequest) (r0 *request.Request, r1 *mq.UpdateConfigurationResponse) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateConfigurationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.UpdateConfigurationResponse)
	return
}

func (d *MQDouble) UpdateConfigurationWithContext(i0 context.Context, i1 *mq.UpdateConfigurationRequest, i2 ...request.Option) (r0 *mq.UpdateConfigurationResponse, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateConfigurationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.UpdateConfigurationResponse)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeUpdateConfigurationWithContext(ctx context.Context, in *mq.UpdateConfigurationRequest, _ ...request.Option) (*mq.UpdateConfigurationResponse, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateConfiguration cancelled", ctx.Err())
	default:
		return d.UpdateConfiguration(in)
	}
}

func (d *MQDouble) UpdateUser(i0 *mq.UpdateUserRequest) (r0 *mq.UpdateUserOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateUser", i0)
	r0, _ = returns[0].(*mq.UpdateUserOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) UpdateUserRequest(i0 *mq.UpdateUserRequest) (r0 *request.Request, r1 *mq.UpdateUserOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateUserRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*mq.UpdateUserOutput)
	return
}

func (d *MQDouble) UpdateUserWithContext(i0 context.Context, i1 *mq.UpdateUserRequest, i2 ...request.Option) (r0 *mq.UpdateUserOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateUserWithContext", i0, i1, i2)
	r0, _ = returns[0].(*mq.UpdateUserOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *MQDouble) fakeUpdateUserWithContext(ctx context.Context, in *mq.UpdateUserRequest, _ ...request.Option) (*mq.UpdateUserOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateUser cancelled", ctx.Err())
	default:
		return d.UpdateUser(in)
	}
}
