// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:12:58+11:00

// Package globalacceleratordouble provides a TestDouble implementation of globalacceleratoriface.GlobalAcceleratorAPI
package globalacceleratordouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/globalaccelerator"
	"github.com/aws/aws-sdk-go/service/globalaccelerator/globalacceleratoriface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// GlobalAcceleratorDouble is TestDouble for globalacceleratoriface.GlobalAcceleratorAPI
type GlobalAcceleratorDouble struct {
	globalacceleratoriface.GlobalAcceleratorAPI
	*awsdouble.AWSTestDouble
}

// Constructor for GlobalAcceleratorDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *GlobalAcceleratorDouble {
	result := &GlobalAcceleratorDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*globalacceleratoriface.GlobalAcceleratorAPI)(nil), configurators...)
	return result
}

func (d *GlobalAcceleratorDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *GlobalAcceleratorDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "CreateAcceleratorWithContext":
		return m.Fake(d.fakeCreateAcceleratorWithContext)

	case "CreateEndpointGroupWithContext":
		return m.Fake(d.fakeCreateEndpointGroupWithContext)

	case "CreateListenerWithContext":
		return m.Fake(d.fakeCreateListenerWithContext)

	case "DeleteAcceleratorWithContext":
		return m.Fake(d.fakeDeleteAcceleratorWithContext)

	case "DeleteEndpointGroupWithContext":
		return m.Fake(d.fakeDeleteEndpointGroupWithContext)

	case "DeleteListenerWithContext":
		return m.Fake(d.fakeDeleteListenerWithContext)

	case "DescribeAcceleratorAttributesWithContext":
		return m.Fake(d.fakeDescribeAcceleratorAttributesWithContext)

	case "DescribeAcceleratorWithContext":
		return m.Fake(d.fakeDescribeAcceleratorWithContext)

	case "DescribeEndpointGroupWithContext":
		return m.Fake(d.fakeDescribeEndpointGroupWithContext)

	case "DescribeListenerWithContext":
		return m.Fake(d.fakeDescribeListenerWithContext)

	case "ListAcceleratorsWithContext":
		return m.Fake(d.fakeListAcceleratorsWithContext)

	case "ListEndpointGroupsWithContext":
		return m.Fake(d.fakeListEndpointGroupsWithContext)

	case "ListListenersWithContext":
		return m.Fake(d.fakeListListenersWithContext)

	case "UpdateAcceleratorAttributesWithContext":
		return m.Fake(d.fakeUpdateAcceleratorAttributesWithContext)

	case "UpdateAcceleratorWithContext":
		return m.Fake(d.fakeUpdateAcceleratorWithContext)

	case "UpdateEndpointGroupWithContext":
		return m.Fake(d.fakeUpdateEndpointGroupWithContext)

	case "UpdateListenerWithContext":
		return m.Fake(d.fakeUpdateListenerWithContext)

	default:
		return nil
	}
}

func (d *GlobalAcceleratorDouble) CreateAccelerator(i0 *globalaccelerator.CreateAcceleratorInput) (r0 *globalaccelerator.CreateAcceleratorOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateAccelerator", i0)
	r0, _ = returns[0].(*globalaccelerator.CreateAcceleratorOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) CreateAcceleratorRequest(i0 *globalaccelerator.CreateAcceleratorInput) (r0 *request.Request, r1 *globalaccelerator.CreateAcceleratorOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateAcceleratorRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*globalaccelerator.CreateAcceleratorOutput)
	return
}

func (d *GlobalAcceleratorDouble) CreateAcceleratorWithContext(i0 context.Context, i1 *globalaccelerator.CreateAcceleratorInput, i2 ...request.Option) (r0 *globalaccelerator.CreateAcceleratorOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateAcceleratorWithContext", i0, i1, i2)
	r0, _ = returns[0].(*globalaccelerator.CreateAcceleratorOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) fakeCreateAcceleratorWithContext(ctx context.Context, in *globalaccelerator.CreateAcceleratorInput, _ ...request.Option) (*globalaccelerator.CreateAcceleratorOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateAccelerator cancelled", ctx.Err())
	default:
		return d.CreateAccelerator(in)
	}
}

func (d *GlobalAcceleratorDouble) CreateEndpointGroup(i0 *globalaccelerator.CreateEndpointGroupInput) (r0 *globalaccelerator.CreateEndpointGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateEndpointGroup", i0)
	r0, _ = returns[0].(*globalaccelerator.CreateEndpointGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) CreateEndpointGroupRequest(i0 *globalaccelerator.CreateEndpointGroupInput) (r0 *request.Request, r1 *globalaccelerator.CreateEndpointGroupOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateEndpointGroupRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*globalaccelerator.CreateEndpointGroupOutput)
	return
}

func (d *GlobalAcceleratorDouble) CreateEndpointGroupWithContext(i0 context.Context, i1 *globalaccelerator.CreateEndpointGroupInput, i2 ...request.Option) (r0 *globalaccelerator.CreateEndpointGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateEndpointGroupWithContext", i0, i1, i2)
	r0, _ = returns[0].(*globalaccelerator.CreateEndpointGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) fakeCreateEndpointGroupWithContext(ctx context.Context, in *globalaccelerator.CreateEndpointGroupInput, _ ...request.Option) (*globalaccelerator.CreateEndpointGroupOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateEndpointGroup cancelled", ctx.Err())
	default:
		return d.CreateEndpointGroup(in)
	}
}

func (d *GlobalAcceleratorDouble) CreateListener(i0 *globalaccelerator.CreateListenerInput) (r0 *globalaccelerator.CreateListenerOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateListener", i0)
	r0, _ = returns[0].(*globalaccelerator.CreateListenerOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) CreateListenerRequest(i0 *globalaccelerator.CreateListenerInput) (r0 *request.Request, r1 *globalaccelerator.CreateListenerOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateListenerRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*globalaccelerator.CreateListenerOutput)
	return
}

func (d *GlobalAcceleratorDouble) CreateListenerWithContext(i0 context.Context, i1 *globalaccelerator.CreateListenerInput, i2 ...request.Option) (r0 *globalaccelerator.CreateListenerOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateListenerWithContext", i0, i1, i2)
	r0, _ = returns[0].(*globalaccelerator.CreateListenerOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) fakeCreateListenerWithContext(ctx context.Context, in *globalaccelerator.CreateListenerInput, _ ...request.Option) (*globalaccelerator.CreateListenerOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateListener cancelled", ctx.Err())
	default:
		return d.CreateListener(in)
	}
}

func (d *GlobalAcceleratorDouble) DeleteAccelerator(i0 *globalaccelerator.DeleteAcceleratorInput) (r0 *globalaccelerator.DeleteAcceleratorOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteAccelerator", i0)
	r0, _ = returns[0].(*globalaccelerator.DeleteAcceleratorOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) DeleteAcceleratorRequest(i0 *globalaccelerator.DeleteAcceleratorInput) (r0 *request.Request, r1 *globalaccelerator.DeleteAcceleratorOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteAcceleratorRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*globalaccelerator.DeleteAcceleratorOutput)
	return
}

func (d *GlobalAcceleratorDouble) DeleteAcceleratorWithContext(i0 context.Context, i1 *globalaccelerator.DeleteAcceleratorInput, i2 ...request.Option) (r0 *globalaccelerator.DeleteAcceleratorOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteAcceleratorWithContext", i0, i1, i2)
	r0, _ = returns[0].(*globalaccelerator.DeleteAcceleratorOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) fakeDeleteAcceleratorWithContext(ctx context.Context, in *globalaccelerator.DeleteAcceleratorInput, _ ...request.Option) (*globalaccelerator.DeleteAcceleratorOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteAccelerator cancelled", ctx.Err())
	default:
		return d.DeleteAccelerator(in)
	}
}

func (d *GlobalAcceleratorDouble) DeleteEndpointGroup(i0 *globalaccelerator.DeleteEndpointGroupInput) (r0 *globalaccelerator.DeleteEndpointGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteEndpointGroup", i0)
	r0, _ = returns[0].(*globalaccelerator.DeleteEndpointGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) DeleteEndpointGroupRequest(i0 *globalaccelerator.DeleteEndpointGroupInput) (r0 *request.Request, r1 *globalaccelerator.DeleteEndpointGroupOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteEndpointGroupRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*globalaccelerator.DeleteEndpointGroupOutput)
	return
}

func (d *GlobalAcceleratorDouble) DeleteEndpointGroupWithContext(i0 context.Context, i1 *globalaccelerator.DeleteEndpointGroupInput, i2 ...request.Option) (r0 *globalaccelerator.DeleteEndpointGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteEndpointGroupWithContext", i0, i1, i2)
	r0, _ = returns[0].(*globalaccelerator.DeleteEndpointGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) fakeDeleteEndpointGroupWithContext(ctx context.Context, in *globalaccelerator.DeleteEndpointGroupInput, _ ...request.Option) (*globalaccelerator.DeleteEndpointGroupOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteEndpointGroup cancelled", ctx.Err())
	default:
		return d.DeleteEndpointGroup(in)
	}
}

func (d *GlobalAcceleratorDouble) DeleteListener(i0 *globalaccelerator.DeleteListenerInput) (r0 *globalaccelerator.DeleteListenerOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteListener", i0)
	r0, _ = returns[0].(*globalaccelerator.DeleteListenerOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) DeleteListenerRequest(i0 *globalaccelerator.DeleteListenerInput) (r0 *request.Request, r1 *globalaccelerator.DeleteListenerOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteListenerRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*globalaccelerator.DeleteListenerOutput)
	return
}

func (d *GlobalAcceleratorDouble) DeleteListenerWithContext(i0 context.Context, i1 *globalaccelerator.DeleteListenerInput, i2 ...request.Option) (r0 *globalaccelerator.DeleteListenerOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteListenerWithContext", i0, i1, i2)
	r0, _ = returns[0].(*globalaccelerator.DeleteListenerOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) fakeDeleteListenerWithContext(ctx context.Context, in *globalaccelerator.DeleteListenerInput, _ ...request.Option) (*globalaccelerator.DeleteListenerOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteListener cancelled", ctx.Err())
	default:
		return d.DeleteListener(in)
	}
}

func (d *GlobalAcceleratorDouble) DescribeAccelerator(i0 *globalaccelerator.DescribeAcceleratorInput) (r0 *globalaccelerator.DescribeAcceleratorOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeAccelerator", i0)
	r0, _ = returns[0].(*globalaccelerator.DescribeAcceleratorOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) DescribeAcceleratorAttributes(i0 *globalaccelerator.DescribeAcceleratorAttributesInput) (r0 *globalaccelerator.DescribeAcceleratorAttributesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeAcceleratorAttributes", i0)
	r0, _ = returns[0].(*globalaccelerator.DescribeAcceleratorAttributesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) DescribeAcceleratorAttributesRequest(i0 *globalaccelerator.DescribeAcceleratorAttributesInput) (r0 *request.Request, r1 *globalaccelerator.DescribeAcceleratorAttributesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeAcceleratorAttributesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*globalaccelerator.DescribeAcceleratorAttributesOutput)
	return
}

func (d *GlobalAcceleratorDouble) DescribeAcceleratorAttributesWithContext(i0 context.Context, i1 *globalaccelerator.DescribeAcceleratorAttributesInput, i2 ...request.Option) (r0 *globalaccelerator.DescribeAcceleratorAttributesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeAcceleratorAttributesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*globalaccelerator.DescribeAcceleratorAttributesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) fakeDescribeAcceleratorAttributesWithContext(ctx context.Context, in *globalaccelerator.DescribeAcceleratorAttributesInput, _ ...request.Option) (*globalaccelerator.DescribeAcceleratorAttributesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeAcceleratorAttributes cancelled", ctx.Err())
	default:
		return d.DescribeAcceleratorAttributes(in)
	}
}

func (d *GlobalAcceleratorDouble) DescribeAcceleratorRequest(i0 *globalaccelerator.DescribeAcceleratorInput) (r0 *request.Request, r1 *globalaccelerator.DescribeAcceleratorOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeAcceleratorRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*globalaccelerator.DescribeAcceleratorOutput)
	return
}

func (d *GlobalAcceleratorDouble) DescribeAcceleratorWithContext(i0 context.Context, i1 *globalaccelerator.DescribeAcceleratorInput, i2 ...request.Option) (r0 *globalaccelerator.DescribeAcceleratorOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeAcceleratorWithContext", i0, i1, i2)
	r0, _ = returns[0].(*globalaccelerator.DescribeAcceleratorOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) fakeDescribeAcceleratorWithContext(ctx context.Context, in *globalaccelerator.DescribeAcceleratorInput, _ ...request.Option) (*globalaccelerator.DescribeAcceleratorOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeAccelerator cancelled", ctx.Err())
	default:
		return d.DescribeAccelerator(in)
	}
}

func (d *GlobalAcceleratorDouble) DescribeEndpointGroup(i0 *globalaccelerator.DescribeEndpointGroupInput) (r0 *globalaccelerator.DescribeEndpointGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeEndpointGroup", i0)
	r0, _ = returns[0].(*globalaccelerator.DescribeEndpointGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) DescribeEndpointGroupRequest(i0 *globalaccelerator.DescribeEndpointGroupInput) (r0 *request.Request, r1 *globalaccelerator.DescribeEndpointGroupOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeEndpointGroupRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*globalaccelerator.DescribeEndpointGroupOutput)
	return
}

func (d *GlobalAcceleratorDouble) DescribeEndpointGroupWithContext(i0 context.Context, i1 *globalaccelerator.DescribeEndpointGroupInput, i2 ...request.Option) (r0 *globalaccelerator.DescribeEndpointGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeEndpointGroupWithContext", i0, i1, i2)
	r0, _ = returns[0].(*globalaccelerator.DescribeEndpointGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) fakeDescribeEndpointGroupWithContext(ctx context.Context, in *globalaccelerator.DescribeEndpointGroupInput, _ ...request.Option) (*globalaccelerator.DescribeEndpointGroupOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeEndpointGroup cancelled", ctx.Err())
	default:
		return d.DescribeEndpointGroup(in)
	}
}

func (d *GlobalAcceleratorDouble) DescribeListener(i0 *globalaccelerator.DescribeListenerInput) (r0 *globalaccelerator.DescribeListenerOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeListener", i0)
	r0, _ = returns[0].(*globalaccelerator.DescribeListenerOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) DescribeListenerRequest(i0 *globalaccelerator.DescribeListenerInput) (r0 *request.Request, r1 *globalaccelerator.DescribeListenerOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeListenerRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*globalaccelerator.DescribeListenerOutput)
	return
}

func (d *GlobalAcceleratorDouble) DescribeListenerWithContext(i0 context.Context, i1 *globalaccelerator.DescribeListenerInput, i2 ...request.Option) (r0 *globalaccelerator.DescribeListenerOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeListenerWithContext", i0, i1, i2)
	r0, _ = returns[0].(*globalaccelerator.DescribeListenerOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) fakeDescribeListenerWithContext(ctx context.Context, in *globalaccelerator.DescribeListenerInput, _ ...request.Option) (*globalaccelerator.DescribeListenerOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeListener cancelled", ctx.Err())
	default:
		return d.DescribeListener(in)
	}
}

func (d *GlobalAcceleratorDouble) ListAccelerators(i0 *globalaccelerator.ListAcceleratorsInput) (r0 *globalaccelerator.ListAcceleratorsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListAccelerators", i0)
	r0, _ = returns[0].(*globalaccelerator.ListAcceleratorsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) ListAcceleratorsRequest(i0 *globalaccelerator.ListAcceleratorsInput) (r0 *request.Request, r1 *globalaccelerator.ListAcceleratorsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListAcceleratorsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*globalaccelerator.ListAcceleratorsOutput)
	return
}

func (d *GlobalAcceleratorDouble) ListAcceleratorsWithContext(i0 context.Context, i1 *globalaccelerator.ListAcceleratorsInput, i2 ...request.Option) (r0 *globalaccelerator.ListAcceleratorsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListAcceleratorsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*globalaccelerator.ListAcceleratorsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) fakeListAcceleratorsWithContext(ctx context.Context, in *globalaccelerator.ListAcceleratorsInput, _ ...request.Option) (*globalaccelerator.ListAcceleratorsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListAccelerators cancelled", ctx.Err())
	default:
		return d.ListAccelerators(in)
	}
}

func (d *GlobalAcceleratorDouble) ListEndpointGroups(i0 *globalaccelerator.ListEndpointGroupsInput) (r0 *globalaccelerator.ListEndpointGroupsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListEndpointGroups", i0)
	r0, _ = returns[0].(*globalaccelerator.ListEndpointGroupsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) ListEndpointGroupsRequest(i0 *globalaccelerator.ListEndpointGroupsInput) (r0 *request.Request, r1 *globalaccelerator.ListEndpointGroupsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListEndpointGroupsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*globalaccelerator.ListEndpointGroupsOutput)
	return
}

func (d *GlobalAcceleratorDouble) ListEndpointGroupsWithContext(i0 context.Context, i1 *globalaccelerator.ListEndpointGroupsInput, i2 ...request.Option) (r0 *globalaccelerator.ListEndpointGroupsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListEndpointGroupsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*globalaccelerator.ListEndpointGroupsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) fakeListEndpointGroupsWithContext(ctx context.Context, in *globalaccelerator.ListEndpointGroupsInput, _ ...request.Option) (*globalaccelerator.ListEndpointGroupsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListEndpointGroups cancelled", ctx.Err())
	default:
		return d.ListEndpointGroups(in)
	}
}

func (d *GlobalAcceleratorDouble) ListListeners(i0 *globalaccelerator.ListListenersInput) (r0 *globalaccelerator.ListListenersOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListListeners", i0)
	r0, _ = returns[0].(*globalaccelerator.ListListenersOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) ListListenersRequest(i0 *globalaccelerator.ListListenersInput) (r0 *request.Request, r1 *globalaccelerator.ListListenersOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListListenersRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*globalaccelerator.ListListenersOutput)
	return
}

func (d *GlobalAcceleratorDouble) ListListenersWithContext(i0 context.Context, i1 *globalaccelerator.ListListenersInput, i2 ...request.Option) (r0 *globalaccelerator.ListListenersOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListListenersWithContext", i0, i1, i2)
	r0, _ = returns[0].(*globalaccelerator.ListListenersOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) fakeListListenersWithContext(ctx context.Context, in *globalaccelerator.ListListenersInput, _ ...request.Option) (*globalaccelerator.ListListenersOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListListeners cancelled", ctx.Err())
	default:
		return d.ListListeners(in)
	}
}

func (d *GlobalAcceleratorDouble) UpdateAccelerator(i0 *globalaccelerator.UpdateAcceleratorInput) (r0 *globalaccelerator.UpdateAcceleratorOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateAccelerator", i0)
	r0, _ = returns[0].(*globalaccelerator.UpdateAcceleratorOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) UpdateAcceleratorAttributes(i0 *globalaccelerator.UpdateAcceleratorAttributesInput) (r0 *globalaccelerator.UpdateAcceleratorAttributesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateAcceleratorAttributes", i0)
	r0, _ = returns[0].(*globalaccelerator.UpdateAcceleratorAttributesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) UpdateAcceleratorAttributesRequest(i0 *globalaccelerator.UpdateAcceleratorAttributesInput) (r0 *request.Request, r1 *globalaccelerator.UpdateAcceleratorAttributesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateAcceleratorAttributesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*globalaccelerator.UpdateAcceleratorAttributesOutput)
	return
}

func (d *GlobalAcceleratorDouble) UpdateAcceleratorAttributesWithContext(i0 context.Context, i1 *globalaccelerator.UpdateAcceleratorAttributesInput, i2 ...request.Option) (r0 *globalaccelerator.UpdateAcceleratorAttributesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateAcceleratorAttributesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*globalaccelerator.UpdateAcceleratorAttributesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) fakeUpdateAcceleratorAttributesWithContext(ctx context.Context, in *globalaccelerator.UpdateAcceleratorAttributesInput, _ ...request.Option) (*globalaccelerator.UpdateAcceleratorAttributesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateAcceleratorAttributes cancelled", ctx.Err())
	default:
		return d.UpdateAcceleratorAttributes(in)
	}
}

func (d *GlobalAcceleratorDouble) UpdateAcceleratorRequest(i0 *globalaccelerator.UpdateAcceleratorInput) (r0 *request.Request, r1 *globalaccelerator.UpdateAcceleratorOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateAcceleratorRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*globalaccelerator.UpdateAcceleratorOutput)
	return
}

func (d *GlobalAcceleratorDouble) UpdateAcceleratorWithContext(i0 context.Context, i1 *globalaccelerator.UpdateAcceleratorInput, i2 ...request.Option) (r0 *globalaccelerator.UpdateAcceleratorOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateAcceleratorWithContext", i0, i1, i2)
	r0, _ = returns[0].(*globalaccelerator.UpdateAcceleratorOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) fakeUpdateAcceleratorWithContext(ctx context.Context, in *globalaccelerator.UpdateAcceleratorInput, _ ...request.Option) (*globalaccelerator.UpdateAcceleratorOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateAccelerator cancelled", ctx.Err())
	default:
		return d.UpdateAccelerator(in)
	}
}

func (d *GlobalAcceleratorDouble) UpdateEndpointGroup(i0 *globalaccelerator.UpdateEndpointGroupInput) (r0 *globalaccelerator.UpdateEndpointGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateEndpointGroup", i0)
	r0, _ = returns[0].(*globalaccelerator.UpdateEndpointGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) UpdateEndpointGroupRequest(i0 *globalaccelerator.UpdateEndpointGroupInput) (r0 *request.Request, r1 *globalaccelerator.UpdateEndpointGroupOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateEndpointGroupRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*globalaccelerator.UpdateEndpointGroupOutput)
	return
}

func (d *GlobalAcceleratorDouble) UpdateEndpointGroupWithContext(i0 context.Context, i1 *globalaccelerator.UpdateEndpointGroupInput, i2 ...request.Option) (r0 *globalaccelerator.UpdateEndpointGroupOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateEndpointGroupWithContext", i0, i1, i2)
	r0, _ = returns[0].(*globalaccelerator.UpdateEndpointGroupOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) fakeUpdateEndpointGroupWithContext(ctx context.Context, in *globalaccelerator.UpdateEndpointGroupInput, _ ...request.Option) (*globalaccelerator.UpdateEndpointGroupOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateEndpointGroup cancelled", ctx.Err())
	default:
		return d.UpdateEndpointGroup(in)
	}
}

func (d *GlobalAcceleratorDouble) UpdateListener(i0 *globalaccelerator.UpdateListenerInput) (r0 *globalaccelerator.UpdateListenerOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateListener", i0)
	r0, _ = returns[0].(*globalaccelerator.UpdateListenerOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) UpdateListenerRequest(i0 *globalaccelerator.UpdateListenerInput) (r0 *request.Request, r1 *globalaccelerator.UpdateListenerOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateListenerRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*globalaccelerator.UpdateListenerOutput)
	return
}

func (d *GlobalAcceleratorDouble) UpdateListenerWithContext(i0 context.Context, i1 *globalaccelerator.UpdateListenerInput, i2 ...request.Option) (r0 *globalaccelerator.UpdateListenerOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateListenerWithContext", i0, i1, i2)
	r0, _ = returns[0].(*globalaccelerator.UpdateListenerOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *GlobalAcceleratorDouble) fakeUpdateListenerWithContext(ctx context.Context, in *globalaccelerator.UpdateListenerInput, _ ...request.Option) (*globalaccelerator.UpdateListenerOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateListener cancelled", ctx.Err())
	default:
		return d.UpdateListener(in)
	}
}