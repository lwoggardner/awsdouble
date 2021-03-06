// Code generated by go doublegen; DO NOT EDIT.
// This file was generated at 2020-02-14T22:12:09+11:00

// Package costexplorerdouble provides a TestDouble implementation of costexploreriface.CostExplorerAPI
package costexplorerdouble

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/costexplorer/costexploreriface"
	"github.com/lwoggardner/awsdouble"
	"github.com/lwoggardner/godouble/godouble"
)

// CostExplorerDouble is TestDouble for costexploreriface.CostExplorerAPI
type CostExplorerDouble struct {
	costexploreriface.CostExplorerAPI
	*awsdouble.AWSTestDouble
}

// Constructor for CostExplorerDouble
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
func NewDouble(t godouble.T, configurators ...func(*awsdouble.AWSTestDouble)) *CostExplorerDouble {
	result := &CostExplorerDouble{}

	configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
		d.SetDefaultCall(result.defaultMethodCall)
		d.SetDefaultReturnValues(result.defaultReturnValues)
	}}, configurators...)

	result.AWSTestDouble = awsdouble.NewDouble(t, (*costexploreriface.CostExplorerAPI)(nil), configurators...)
	return result
}

func (d *CostExplorerDouble) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
	return d.DefaultReturnValues(m)
}

func (d *CostExplorerDouble) defaultMethodCall(m godouble.Method) godouble.MethodCall {
	switch m.Reflect().Name {

	case "CreateCostCategoryDefinitionWithContext":
		return m.Fake(d.fakeCreateCostCategoryDefinitionWithContext)

	case "DeleteCostCategoryDefinitionWithContext":
		return m.Fake(d.fakeDeleteCostCategoryDefinitionWithContext)

	case "DescribeCostCategoryDefinitionWithContext":
		return m.Fake(d.fakeDescribeCostCategoryDefinitionWithContext)

	case "GetCostAndUsageWithContext":
		return m.Fake(d.fakeGetCostAndUsageWithContext)

	case "GetCostAndUsageWithResourcesWithContext":
		return m.Fake(d.fakeGetCostAndUsageWithResourcesWithContext)

	case "GetCostForecastWithContext":
		return m.Fake(d.fakeGetCostForecastWithContext)

	case "GetDimensionValuesWithContext":
		return m.Fake(d.fakeGetDimensionValuesWithContext)

	case "GetReservationCoverageWithContext":
		return m.Fake(d.fakeGetReservationCoverageWithContext)

	case "GetReservationPurchaseRecommendationWithContext":
		return m.Fake(d.fakeGetReservationPurchaseRecommendationWithContext)

	case "GetReservationUtilizationWithContext":
		return m.Fake(d.fakeGetReservationUtilizationWithContext)

	case "GetRightsizingRecommendationWithContext":
		return m.Fake(d.fakeGetRightsizingRecommendationWithContext)

	case "GetSavingsPlansCoveragePages":
		return m.Fake(d.fakeGetSavingsPlansCoveragePages)

	case "GetSavingsPlansCoveragePagesWithContext":
		return m.Fake(d.fakeGetSavingsPlansCoveragePagesWithContext)

	case "GetSavingsPlansCoverageWithContext":
		return m.Fake(d.fakeGetSavingsPlansCoverageWithContext)

	case "GetSavingsPlansPurchaseRecommendationWithContext":
		return m.Fake(d.fakeGetSavingsPlansPurchaseRecommendationWithContext)

	case "GetSavingsPlansUtilizationDetailsPages":
		return m.Fake(d.fakeGetSavingsPlansUtilizationDetailsPages)

	case "GetSavingsPlansUtilizationDetailsPagesWithContext":
		return m.Fake(d.fakeGetSavingsPlansUtilizationDetailsPagesWithContext)

	case "GetSavingsPlansUtilizationDetailsWithContext":
		return m.Fake(d.fakeGetSavingsPlansUtilizationDetailsWithContext)

	case "GetSavingsPlansUtilizationWithContext":
		return m.Fake(d.fakeGetSavingsPlansUtilizationWithContext)

	case "GetTagsWithContext":
		return m.Fake(d.fakeGetTagsWithContext)

	case "GetUsageForecastWithContext":
		return m.Fake(d.fakeGetUsageForecastWithContext)

	case "ListCostCategoryDefinitionsWithContext":
		return m.Fake(d.fakeListCostCategoryDefinitionsWithContext)

	case "UpdateCostCategoryDefinitionWithContext":
		return m.Fake(d.fakeUpdateCostCategoryDefinitionWithContext)

	default:
		return nil
	}
}

func (d *CostExplorerDouble) CreateCostCategoryDefinition(i0 *costexplorer.CreateCostCategoryDefinitionInput) (r0 *costexplorer.CreateCostCategoryDefinitionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateCostCategoryDefinition", i0)
	r0, _ = returns[0].(*costexplorer.CreateCostCategoryDefinitionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) CreateCostCategoryDefinitionRequest(i0 *costexplorer.CreateCostCategoryDefinitionInput) (r0 *request.Request, r1 *costexplorer.CreateCostCategoryDefinitionOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateCostCategoryDefinitionRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*costexplorer.CreateCostCategoryDefinitionOutput)
	return
}

func (d *CostExplorerDouble) CreateCostCategoryDefinitionWithContext(i0 context.Context, i1 *costexplorer.CreateCostCategoryDefinitionInput, i2 ...request.Option) (r0 *costexplorer.CreateCostCategoryDefinitionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("CreateCostCategoryDefinitionWithContext", i0, i1, i2)
	r0, _ = returns[0].(*costexplorer.CreateCostCategoryDefinitionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) fakeCreateCostCategoryDefinitionWithContext(ctx context.Context, in *costexplorer.CreateCostCategoryDefinitionInput, _ ...request.Option) (*costexplorer.CreateCostCategoryDefinitionOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "CreateCostCategoryDefinition cancelled", ctx.Err())
	default:
		return d.CreateCostCategoryDefinition(in)
	}
}

func (d *CostExplorerDouble) DeleteCostCategoryDefinition(i0 *costexplorer.DeleteCostCategoryDefinitionInput) (r0 *costexplorer.DeleteCostCategoryDefinitionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteCostCategoryDefinition", i0)
	r0, _ = returns[0].(*costexplorer.DeleteCostCategoryDefinitionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) DeleteCostCategoryDefinitionRequest(i0 *costexplorer.DeleteCostCategoryDefinitionInput) (r0 *request.Request, r1 *costexplorer.DeleteCostCategoryDefinitionOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteCostCategoryDefinitionRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*costexplorer.DeleteCostCategoryDefinitionOutput)
	return
}

func (d *CostExplorerDouble) DeleteCostCategoryDefinitionWithContext(i0 context.Context, i1 *costexplorer.DeleteCostCategoryDefinitionInput, i2 ...request.Option) (r0 *costexplorer.DeleteCostCategoryDefinitionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DeleteCostCategoryDefinitionWithContext", i0, i1, i2)
	r0, _ = returns[0].(*costexplorer.DeleteCostCategoryDefinitionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) fakeDeleteCostCategoryDefinitionWithContext(ctx context.Context, in *costexplorer.DeleteCostCategoryDefinitionInput, _ ...request.Option) (*costexplorer.DeleteCostCategoryDefinitionOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DeleteCostCategoryDefinition cancelled", ctx.Err())
	default:
		return d.DeleteCostCategoryDefinition(in)
	}
}

func (d *CostExplorerDouble) DescribeCostCategoryDefinition(i0 *costexplorer.DescribeCostCategoryDefinitionInput) (r0 *costexplorer.DescribeCostCategoryDefinitionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeCostCategoryDefinition", i0)
	r0, _ = returns[0].(*costexplorer.DescribeCostCategoryDefinitionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) DescribeCostCategoryDefinitionRequest(i0 *costexplorer.DescribeCostCategoryDefinitionInput) (r0 *request.Request, r1 *costexplorer.DescribeCostCategoryDefinitionOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeCostCategoryDefinitionRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*costexplorer.DescribeCostCategoryDefinitionOutput)
	return
}

func (d *CostExplorerDouble) DescribeCostCategoryDefinitionWithContext(i0 context.Context, i1 *costexplorer.DescribeCostCategoryDefinitionInput, i2 ...request.Option) (r0 *costexplorer.DescribeCostCategoryDefinitionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("DescribeCostCategoryDefinitionWithContext", i0, i1, i2)
	r0, _ = returns[0].(*costexplorer.DescribeCostCategoryDefinitionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) fakeDescribeCostCategoryDefinitionWithContext(ctx context.Context, in *costexplorer.DescribeCostCategoryDefinitionInput, _ ...request.Option) (*costexplorer.DescribeCostCategoryDefinitionOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "DescribeCostCategoryDefinition cancelled", ctx.Err())
	default:
		return d.DescribeCostCategoryDefinition(in)
	}
}

func (d *CostExplorerDouble) GetCostAndUsage(i0 *costexplorer.GetCostAndUsageInput) (r0 *costexplorer.GetCostAndUsageOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetCostAndUsage", i0)
	r0, _ = returns[0].(*costexplorer.GetCostAndUsageOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) GetCostAndUsageRequest(i0 *costexplorer.GetCostAndUsageInput) (r0 *request.Request, r1 *costexplorer.GetCostAndUsageOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetCostAndUsageRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*costexplorer.GetCostAndUsageOutput)
	return
}

func (d *CostExplorerDouble) GetCostAndUsageWithContext(i0 context.Context, i1 *costexplorer.GetCostAndUsageInput, i2 ...request.Option) (r0 *costexplorer.GetCostAndUsageOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetCostAndUsageWithContext", i0, i1, i2)
	r0, _ = returns[0].(*costexplorer.GetCostAndUsageOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) fakeGetCostAndUsageWithContext(ctx context.Context, in *costexplorer.GetCostAndUsageInput, _ ...request.Option) (*costexplorer.GetCostAndUsageOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetCostAndUsage cancelled", ctx.Err())
	default:
		return d.GetCostAndUsage(in)
	}
}

func (d *CostExplorerDouble) GetCostAndUsageWithResources(i0 *costexplorer.GetCostAndUsageWithResourcesInput) (r0 *costexplorer.GetCostAndUsageWithResourcesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetCostAndUsageWithResources", i0)
	r0, _ = returns[0].(*costexplorer.GetCostAndUsageWithResourcesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) GetCostAndUsageWithResourcesRequest(i0 *costexplorer.GetCostAndUsageWithResourcesInput) (r0 *request.Request, r1 *costexplorer.GetCostAndUsageWithResourcesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetCostAndUsageWithResourcesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*costexplorer.GetCostAndUsageWithResourcesOutput)
	return
}

func (d *CostExplorerDouble) GetCostAndUsageWithResourcesWithContext(i0 context.Context, i1 *costexplorer.GetCostAndUsageWithResourcesInput, i2 ...request.Option) (r0 *costexplorer.GetCostAndUsageWithResourcesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetCostAndUsageWithResourcesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*costexplorer.GetCostAndUsageWithResourcesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) fakeGetCostAndUsageWithResourcesWithContext(ctx context.Context, in *costexplorer.GetCostAndUsageWithResourcesInput, _ ...request.Option) (*costexplorer.GetCostAndUsageWithResourcesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetCostAndUsageWithResources cancelled", ctx.Err())
	default:
		return d.GetCostAndUsageWithResources(in)
	}
}

func (d *CostExplorerDouble) GetCostForecast(i0 *costexplorer.GetCostForecastInput) (r0 *costexplorer.GetCostForecastOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetCostForecast", i0)
	r0, _ = returns[0].(*costexplorer.GetCostForecastOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) GetCostForecastRequest(i0 *costexplorer.GetCostForecastInput) (r0 *request.Request, r1 *costexplorer.GetCostForecastOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetCostForecastRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*costexplorer.GetCostForecastOutput)
	return
}

func (d *CostExplorerDouble) GetCostForecastWithContext(i0 context.Context, i1 *costexplorer.GetCostForecastInput, i2 ...request.Option) (r0 *costexplorer.GetCostForecastOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetCostForecastWithContext", i0, i1, i2)
	r0, _ = returns[0].(*costexplorer.GetCostForecastOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) fakeGetCostForecastWithContext(ctx context.Context, in *costexplorer.GetCostForecastInput, _ ...request.Option) (*costexplorer.GetCostForecastOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetCostForecast cancelled", ctx.Err())
	default:
		return d.GetCostForecast(in)
	}
}

func (d *CostExplorerDouble) GetDimensionValues(i0 *costexplorer.GetDimensionValuesInput) (r0 *costexplorer.GetDimensionValuesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetDimensionValues", i0)
	r0, _ = returns[0].(*costexplorer.GetDimensionValuesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) GetDimensionValuesRequest(i0 *costexplorer.GetDimensionValuesInput) (r0 *request.Request, r1 *costexplorer.GetDimensionValuesOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetDimensionValuesRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*costexplorer.GetDimensionValuesOutput)
	return
}

func (d *CostExplorerDouble) GetDimensionValuesWithContext(i0 context.Context, i1 *costexplorer.GetDimensionValuesInput, i2 ...request.Option) (r0 *costexplorer.GetDimensionValuesOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetDimensionValuesWithContext", i0, i1, i2)
	r0, _ = returns[0].(*costexplorer.GetDimensionValuesOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) fakeGetDimensionValuesWithContext(ctx context.Context, in *costexplorer.GetDimensionValuesInput, _ ...request.Option) (*costexplorer.GetDimensionValuesOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetDimensionValues cancelled", ctx.Err())
	default:
		return d.GetDimensionValues(in)
	}
}

func (d *CostExplorerDouble) GetReservationCoverage(i0 *costexplorer.GetReservationCoverageInput) (r0 *costexplorer.GetReservationCoverageOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetReservationCoverage", i0)
	r0, _ = returns[0].(*costexplorer.GetReservationCoverageOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) GetReservationCoverageRequest(i0 *costexplorer.GetReservationCoverageInput) (r0 *request.Request, r1 *costexplorer.GetReservationCoverageOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetReservationCoverageRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*costexplorer.GetReservationCoverageOutput)
	return
}

func (d *CostExplorerDouble) GetReservationCoverageWithContext(i0 context.Context, i1 *costexplorer.GetReservationCoverageInput, i2 ...request.Option) (r0 *costexplorer.GetReservationCoverageOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetReservationCoverageWithContext", i0, i1, i2)
	r0, _ = returns[0].(*costexplorer.GetReservationCoverageOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) fakeGetReservationCoverageWithContext(ctx context.Context, in *costexplorer.GetReservationCoverageInput, _ ...request.Option) (*costexplorer.GetReservationCoverageOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetReservationCoverage cancelled", ctx.Err())
	default:
		return d.GetReservationCoverage(in)
	}
}

func (d *CostExplorerDouble) GetReservationPurchaseRecommendation(i0 *costexplorer.GetReservationPurchaseRecommendationInput) (r0 *costexplorer.GetReservationPurchaseRecommendationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetReservationPurchaseRecommendation", i0)
	r0, _ = returns[0].(*costexplorer.GetReservationPurchaseRecommendationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) GetReservationPurchaseRecommendationRequest(i0 *costexplorer.GetReservationPurchaseRecommendationInput) (r0 *request.Request, r1 *costexplorer.GetReservationPurchaseRecommendationOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetReservationPurchaseRecommendationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*costexplorer.GetReservationPurchaseRecommendationOutput)
	return
}

func (d *CostExplorerDouble) GetReservationPurchaseRecommendationWithContext(i0 context.Context, i1 *costexplorer.GetReservationPurchaseRecommendationInput, i2 ...request.Option) (r0 *costexplorer.GetReservationPurchaseRecommendationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetReservationPurchaseRecommendationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*costexplorer.GetReservationPurchaseRecommendationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) fakeGetReservationPurchaseRecommendationWithContext(ctx context.Context, in *costexplorer.GetReservationPurchaseRecommendationInput, _ ...request.Option) (*costexplorer.GetReservationPurchaseRecommendationOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetReservationPurchaseRecommendation cancelled", ctx.Err())
	default:
		return d.GetReservationPurchaseRecommendation(in)
	}
}

func (d *CostExplorerDouble) GetReservationUtilization(i0 *costexplorer.GetReservationUtilizationInput) (r0 *costexplorer.GetReservationUtilizationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetReservationUtilization", i0)
	r0, _ = returns[0].(*costexplorer.GetReservationUtilizationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) GetReservationUtilizationRequest(i0 *costexplorer.GetReservationUtilizationInput) (r0 *request.Request, r1 *costexplorer.GetReservationUtilizationOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetReservationUtilizationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*costexplorer.GetReservationUtilizationOutput)
	return
}

func (d *CostExplorerDouble) GetReservationUtilizationWithContext(i0 context.Context, i1 *costexplorer.GetReservationUtilizationInput, i2 ...request.Option) (r0 *costexplorer.GetReservationUtilizationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetReservationUtilizationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*costexplorer.GetReservationUtilizationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) fakeGetReservationUtilizationWithContext(ctx context.Context, in *costexplorer.GetReservationUtilizationInput, _ ...request.Option) (*costexplorer.GetReservationUtilizationOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetReservationUtilization cancelled", ctx.Err())
	default:
		return d.GetReservationUtilization(in)
	}
}

func (d *CostExplorerDouble) GetRightsizingRecommendation(i0 *costexplorer.GetRightsizingRecommendationInput) (r0 *costexplorer.GetRightsizingRecommendationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetRightsizingRecommendation", i0)
	r0, _ = returns[0].(*costexplorer.GetRightsizingRecommendationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) GetRightsizingRecommendationRequest(i0 *costexplorer.GetRightsizingRecommendationInput) (r0 *request.Request, r1 *costexplorer.GetRightsizingRecommendationOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetRightsizingRecommendationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*costexplorer.GetRightsizingRecommendationOutput)
	return
}

func (d *CostExplorerDouble) GetRightsizingRecommendationWithContext(i0 context.Context, i1 *costexplorer.GetRightsizingRecommendationInput, i2 ...request.Option) (r0 *costexplorer.GetRightsizingRecommendationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetRightsizingRecommendationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*costexplorer.GetRightsizingRecommendationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) fakeGetRightsizingRecommendationWithContext(ctx context.Context, in *costexplorer.GetRightsizingRecommendationInput, _ ...request.Option) (*costexplorer.GetRightsizingRecommendationOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetRightsizingRecommendation cancelled", ctx.Err())
	default:
		return d.GetRightsizingRecommendation(in)
	}
}

func (d *CostExplorerDouble) GetSavingsPlansCoverage(i0 *costexplorer.GetSavingsPlansCoverageInput) (r0 *costexplorer.GetSavingsPlansCoverageOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSavingsPlansCoverage", i0)
	r0, _ = returns[0].(*costexplorer.GetSavingsPlansCoverageOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) GetSavingsPlansCoveragePages(i0 *costexplorer.GetSavingsPlansCoverageInput, i1 func(*costexplorer.GetSavingsPlansCoverageOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSavingsPlansCoveragePages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *CostExplorerDouble) fakeGetSavingsPlansCoveragePages(in *costexplorer.GetSavingsPlansCoverageInput, pager func(*costexplorer.GetSavingsPlansCoverageOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("GetSavingsPlansCoverage", paginators, in, pager)
}

func (d *CostExplorerDouble) GetSavingsPlansCoveragePagesWithContext(i0 context.Context, i1 *costexplorer.GetSavingsPlansCoverageInput, i2 func(*costexplorer.GetSavingsPlansCoverageOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSavingsPlansCoveragePagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *CostExplorerDouble) fakeGetSavingsPlansCoveragePagesWithContext(ctx context.Context, in *costexplorer.GetSavingsPlansCoverageInput, pager func(*costexplorer.GetSavingsPlansCoverageOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("GetSavingsPlansCoverage", paginators, ctx, in, pager, options...)
}

func (d *CostExplorerDouble) GetSavingsPlansCoverageRequest(i0 *costexplorer.GetSavingsPlansCoverageInput) (r0 *request.Request, r1 *costexplorer.GetSavingsPlansCoverageOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSavingsPlansCoverageRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*costexplorer.GetSavingsPlansCoverageOutput)
	return
}

func (d *CostExplorerDouble) GetSavingsPlansCoverageWithContext(i0 context.Context, i1 *costexplorer.GetSavingsPlansCoverageInput, i2 ...request.Option) (r0 *costexplorer.GetSavingsPlansCoverageOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSavingsPlansCoverageWithContext", i0, i1, i2)
	r0, _ = returns[0].(*costexplorer.GetSavingsPlansCoverageOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) fakeGetSavingsPlansCoverageWithContext(ctx context.Context, in *costexplorer.GetSavingsPlansCoverageInput, _ ...request.Option) (*costexplorer.GetSavingsPlansCoverageOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetSavingsPlansCoverage cancelled", ctx.Err())
	default:
		return d.GetSavingsPlansCoverage(in)
	}
}

func (d *CostExplorerDouble) GetSavingsPlansPurchaseRecommendation(i0 *costexplorer.GetSavingsPlansPurchaseRecommendationInput) (r0 *costexplorer.GetSavingsPlansPurchaseRecommendationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSavingsPlansPurchaseRecommendation", i0)
	r0, _ = returns[0].(*costexplorer.GetSavingsPlansPurchaseRecommendationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) GetSavingsPlansPurchaseRecommendationRequest(i0 *costexplorer.GetSavingsPlansPurchaseRecommendationInput) (r0 *request.Request, r1 *costexplorer.GetSavingsPlansPurchaseRecommendationOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSavingsPlansPurchaseRecommendationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*costexplorer.GetSavingsPlansPurchaseRecommendationOutput)
	return
}

func (d *CostExplorerDouble) GetSavingsPlansPurchaseRecommendationWithContext(i0 context.Context, i1 *costexplorer.GetSavingsPlansPurchaseRecommendationInput, i2 ...request.Option) (r0 *costexplorer.GetSavingsPlansPurchaseRecommendationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSavingsPlansPurchaseRecommendationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*costexplorer.GetSavingsPlansPurchaseRecommendationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) fakeGetSavingsPlansPurchaseRecommendationWithContext(ctx context.Context, in *costexplorer.GetSavingsPlansPurchaseRecommendationInput, _ ...request.Option) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetSavingsPlansPurchaseRecommendation cancelled", ctx.Err())
	default:
		return d.GetSavingsPlansPurchaseRecommendation(in)
	}
}

func (d *CostExplorerDouble) GetSavingsPlansUtilization(i0 *costexplorer.GetSavingsPlansUtilizationInput) (r0 *costexplorer.GetSavingsPlansUtilizationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSavingsPlansUtilization", i0)
	r0, _ = returns[0].(*costexplorer.GetSavingsPlansUtilizationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) GetSavingsPlansUtilizationDetails(i0 *costexplorer.GetSavingsPlansUtilizationDetailsInput) (r0 *costexplorer.GetSavingsPlansUtilizationDetailsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSavingsPlansUtilizationDetails", i0)
	r0, _ = returns[0].(*costexplorer.GetSavingsPlansUtilizationDetailsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) GetSavingsPlansUtilizationDetailsPages(i0 *costexplorer.GetSavingsPlansUtilizationDetailsInput, i1 func(*costexplorer.GetSavingsPlansUtilizationDetailsOutput, bool) bool) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSavingsPlansUtilizationDetailsPages", i0, i1)
	r0, _ = returns[0].(error)
	return
}

func (d *CostExplorerDouble) fakeGetSavingsPlansUtilizationDetailsPages(in *costexplorer.GetSavingsPlansUtilizationDetailsInput, pager func(*costexplorer.GetSavingsPlansUtilizationDetailsOutput, bool) (shouldContinue bool)) error {
	return d.Paginate("GetSavingsPlansUtilizationDetails", paginators, in, pager)
}

func (d *CostExplorerDouble) GetSavingsPlansUtilizationDetailsPagesWithContext(i0 context.Context, i1 *costexplorer.GetSavingsPlansUtilizationDetailsInput, i2 func(*costexplorer.GetSavingsPlansUtilizationDetailsOutput, bool) bool, i3 ...request.Option) (r0 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSavingsPlansUtilizationDetailsPagesWithContext", i0, i1, i2, i3)
	r0, _ = returns[0].(error)
	return
}

func (d *CostExplorerDouble) fakeGetSavingsPlansUtilizationDetailsPagesWithContext(ctx context.Context, in *costexplorer.GetSavingsPlansUtilizationDetailsInput, pager func(*costexplorer.GetSavingsPlansUtilizationDetailsOutput, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("GetSavingsPlansUtilizationDetails", paginators, ctx, in, pager, options...)
}

func (d *CostExplorerDouble) GetSavingsPlansUtilizationDetailsRequest(i0 *costexplorer.GetSavingsPlansUtilizationDetailsInput) (r0 *request.Request, r1 *costexplorer.GetSavingsPlansUtilizationDetailsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSavingsPlansUtilizationDetailsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*costexplorer.GetSavingsPlansUtilizationDetailsOutput)
	return
}

func (d *CostExplorerDouble) GetSavingsPlansUtilizationDetailsWithContext(i0 context.Context, i1 *costexplorer.GetSavingsPlansUtilizationDetailsInput, i2 ...request.Option) (r0 *costexplorer.GetSavingsPlansUtilizationDetailsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSavingsPlansUtilizationDetailsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*costexplorer.GetSavingsPlansUtilizationDetailsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) fakeGetSavingsPlansUtilizationDetailsWithContext(ctx context.Context, in *costexplorer.GetSavingsPlansUtilizationDetailsInput, _ ...request.Option) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetSavingsPlansUtilizationDetails cancelled", ctx.Err())
	default:
		return d.GetSavingsPlansUtilizationDetails(in)
	}
}

func (d *CostExplorerDouble) GetSavingsPlansUtilizationRequest(i0 *costexplorer.GetSavingsPlansUtilizationInput) (r0 *request.Request, r1 *costexplorer.GetSavingsPlansUtilizationOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSavingsPlansUtilizationRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*costexplorer.GetSavingsPlansUtilizationOutput)
	return
}

func (d *CostExplorerDouble) GetSavingsPlansUtilizationWithContext(i0 context.Context, i1 *costexplorer.GetSavingsPlansUtilizationInput, i2 ...request.Option) (r0 *costexplorer.GetSavingsPlansUtilizationOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetSavingsPlansUtilizationWithContext", i0, i1, i2)
	r0, _ = returns[0].(*costexplorer.GetSavingsPlansUtilizationOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) fakeGetSavingsPlansUtilizationWithContext(ctx context.Context, in *costexplorer.GetSavingsPlansUtilizationInput, _ ...request.Option) (*costexplorer.GetSavingsPlansUtilizationOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetSavingsPlansUtilization cancelled", ctx.Err())
	default:
		return d.GetSavingsPlansUtilization(in)
	}
}

func (d *CostExplorerDouble) GetTags(i0 *costexplorer.GetTagsInput) (r0 *costexplorer.GetTagsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetTags", i0)
	r0, _ = returns[0].(*costexplorer.GetTagsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) GetTagsRequest(i0 *costexplorer.GetTagsInput) (r0 *request.Request, r1 *costexplorer.GetTagsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetTagsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*costexplorer.GetTagsOutput)
	return
}

func (d *CostExplorerDouble) GetTagsWithContext(i0 context.Context, i1 *costexplorer.GetTagsInput, i2 ...request.Option) (r0 *costexplorer.GetTagsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetTagsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*costexplorer.GetTagsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) fakeGetTagsWithContext(ctx context.Context, in *costexplorer.GetTagsInput, _ ...request.Option) (*costexplorer.GetTagsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetTags cancelled", ctx.Err())
	default:
		return d.GetTags(in)
	}
}

func (d *CostExplorerDouble) GetUsageForecast(i0 *costexplorer.GetUsageForecastInput) (r0 *costexplorer.GetUsageForecastOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetUsageForecast", i0)
	r0, _ = returns[0].(*costexplorer.GetUsageForecastOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) GetUsageForecastRequest(i0 *costexplorer.GetUsageForecastInput) (r0 *request.Request, r1 *costexplorer.GetUsageForecastOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetUsageForecastRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*costexplorer.GetUsageForecastOutput)
	return
}

func (d *CostExplorerDouble) GetUsageForecastWithContext(i0 context.Context, i1 *costexplorer.GetUsageForecastInput, i2 ...request.Option) (r0 *costexplorer.GetUsageForecastOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("GetUsageForecastWithContext", i0, i1, i2)
	r0, _ = returns[0].(*costexplorer.GetUsageForecastOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) fakeGetUsageForecastWithContext(ctx context.Context, in *costexplorer.GetUsageForecastInput, _ ...request.Option) (*costexplorer.GetUsageForecastOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "GetUsageForecast cancelled", ctx.Err())
	default:
		return d.GetUsageForecast(in)
	}
}

func (d *CostExplorerDouble) ListCostCategoryDefinitions(i0 *costexplorer.ListCostCategoryDefinitionsInput) (r0 *costexplorer.ListCostCategoryDefinitionsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListCostCategoryDefinitions", i0)
	r0, _ = returns[0].(*costexplorer.ListCostCategoryDefinitionsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) ListCostCategoryDefinitionsRequest(i0 *costexplorer.ListCostCategoryDefinitionsInput) (r0 *request.Request, r1 *costexplorer.ListCostCategoryDefinitionsOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListCostCategoryDefinitionsRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*costexplorer.ListCostCategoryDefinitionsOutput)
	return
}

func (d *CostExplorerDouble) ListCostCategoryDefinitionsWithContext(i0 context.Context, i1 *costexplorer.ListCostCategoryDefinitionsInput, i2 ...request.Option) (r0 *costexplorer.ListCostCategoryDefinitionsOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("ListCostCategoryDefinitionsWithContext", i0, i1, i2)
	r0, _ = returns[0].(*costexplorer.ListCostCategoryDefinitionsOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) fakeListCostCategoryDefinitionsWithContext(ctx context.Context, in *costexplorer.ListCostCategoryDefinitionsInput, _ ...request.Option) (*costexplorer.ListCostCategoryDefinitionsOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "ListCostCategoryDefinitions cancelled", ctx.Err())
	default:
		return d.ListCostCategoryDefinitions(in)
	}
}

func (d *CostExplorerDouble) UpdateCostCategoryDefinition(i0 *costexplorer.UpdateCostCategoryDefinitionInput) (r0 *costexplorer.UpdateCostCategoryDefinitionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateCostCategoryDefinition", i0)
	r0, _ = returns[0].(*costexplorer.UpdateCostCategoryDefinitionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) UpdateCostCategoryDefinitionRequest(i0 *costexplorer.UpdateCostCategoryDefinitionInput) (r0 *request.Request, r1 *costexplorer.UpdateCostCategoryDefinitionOutput) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateCostCategoryDefinitionRequest", i0)
	r0, _ = returns[0].(*request.Request)
	r1, _ = returns[1].(*costexplorer.UpdateCostCategoryDefinitionOutput)
	return
}

func (d *CostExplorerDouble) UpdateCostCategoryDefinitionWithContext(i0 context.Context, i1 *costexplorer.UpdateCostCategoryDefinitionInput, i2 ...request.Option) (r0 *costexplorer.UpdateCostCategoryDefinitionOutput, r1 error) {
	d.TestDouble.T().Helper()
	returns := d.TestDouble.Invoke("UpdateCostCategoryDefinitionWithContext", i0, i1, i2)
	r0, _ = returns[0].(*costexplorer.UpdateCostCategoryDefinitionOutput)
	r1, _ = returns[1].(error)
	return
}

func (d *CostExplorerDouble) fakeUpdateCostCategoryDefinitionWithContext(ctx context.Context, in *costexplorer.UpdateCostCategoryDefinitionInput, _ ...request.Option) (*costexplorer.UpdateCostCategoryDefinitionOutput, error) {

	select {
	case <-ctx.Done():
		return nil, awserr.New(request.CanceledErrorCode, "UpdateCostCategoryDefinition cancelled", ctx.Err())
	default:
		return d.UpdateCostCategoryDefinition(in)
	}
}
