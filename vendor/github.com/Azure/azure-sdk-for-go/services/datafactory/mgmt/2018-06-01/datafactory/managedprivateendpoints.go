package datafactory

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ManagedPrivateEndpointsClient is the the Azure Data Factory V2 management API provides a RESTful set of web services
// that interact with Azure Data Factory V2 services.
type ManagedPrivateEndpointsClient struct {
	BaseClient
}

// NewManagedPrivateEndpointsClient creates an instance of the ManagedPrivateEndpointsClient client.
func NewManagedPrivateEndpointsClient(subscriptionID string) ManagedPrivateEndpointsClient {
	return NewManagedPrivateEndpointsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagedPrivateEndpointsClientWithBaseURI creates an instance of the ManagedPrivateEndpointsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewManagedPrivateEndpointsClientWithBaseURI(baseURI string, subscriptionID string) ManagedPrivateEndpointsClient {
	return ManagedPrivateEndpointsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a managed private endpoint.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// managedVirtualNetworkName - managed virtual network name
// managedPrivateEndpointName - managed private endpoint name
// managedPrivateEndpoint - managed private endpoint resource definition.
// ifMatch - eTag of the managed private endpoint entity. Should only be specified for update, for which it
// should match existing entity or can be * for unconditional update.
func (client ManagedPrivateEndpointsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string, managedPrivateEndpoint ManagedPrivateEndpointResource, ifMatch string) (result ManagedPrivateEndpointResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedPrivateEndpointsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: managedVirtualNetworkName,
			Constraints: []validation.Constraint{{Target: "managedVirtualNetworkName", Name: validation.MaxLength, Rule: 127, Chain: nil},
				{Target: "managedVirtualNetworkName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "managedVirtualNetworkName", Name: validation.Pattern, Rule: `^([_A-Za-z0-9]|([_A-Za-z0-9][-_A-Za-z0-9]{0,125}[_A-Za-z0-9]))$`, Chain: nil}}},
		{TargetValue: managedPrivateEndpointName,
			Constraints: []validation.Constraint{{Target: "managedPrivateEndpointName", Name: validation.MaxLength, Rule: 127, Chain: nil},
				{Target: "managedPrivateEndpointName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "managedPrivateEndpointName", Name: validation.Pattern, Rule: `^([_A-Za-z0-9]|([_A-Za-z0-9][-_A-Za-z0-9]{0,125}[_A-Za-z0-9]))$`, Chain: nil}}},
		{TargetValue: managedPrivateEndpoint,
			Constraints: []validation.Constraint{{Target: "managedPrivateEndpoint.Properties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.ManagedPrivateEndpointsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, factoryName, managedVirtualNetworkName, managedPrivateEndpointName, managedPrivateEndpoint, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.ManagedPrivateEndpointsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.ManagedPrivateEndpointsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.ManagedPrivateEndpointsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ManagedPrivateEndpointsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string, managedPrivateEndpoint ManagedPrivateEndpointResource, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":                autorest.Encode("path", factoryName),
		"managedPrivateEndpointName": autorest.Encode("path", managedPrivateEndpointName),
		"managedVirtualNetworkName":  autorest.Encode("path", managedVirtualNetworkName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/managedVirtualNetworks/{managedVirtualNetworkName}/managedPrivateEndpoints/{managedPrivateEndpointName}", pathParameters),
		autorest.WithJSON(managedPrivateEndpoint),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedPrivateEndpointsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ManagedPrivateEndpointsClient) CreateOrUpdateResponder(resp *http.Response) (result ManagedPrivateEndpointResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a managed private endpoint.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// managedVirtualNetworkName - managed virtual network name
// managedPrivateEndpointName - managed private endpoint name
func (client ManagedPrivateEndpointsClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedPrivateEndpointsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: managedVirtualNetworkName,
			Constraints: []validation.Constraint{{Target: "managedVirtualNetworkName", Name: validation.MaxLength, Rule: 127, Chain: nil},
				{Target: "managedVirtualNetworkName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "managedVirtualNetworkName", Name: validation.Pattern, Rule: `^([_A-Za-z0-9]|([_A-Za-z0-9][-_A-Za-z0-9]{0,125}[_A-Za-z0-9]))$`, Chain: nil}}},
		{TargetValue: managedPrivateEndpointName,
			Constraints: []validation.Constraint{{Target: "managedPrivateEndpointName", Name: validation.MaxLength, Rule: 127, Chain: nil},
				{Target: "managedPrivateEndpointName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "managedPrivateEndpointName", Name: validation.Pattern, Rule: `^([_A-Za-z0-9]|([_A-Za-z0-9][-_A-Za-z0-9]{0,125}[_A-Za-z0-9]))$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.ManagedPrivateEndpointsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, factoryName, managedVirtualNetworkName, managedPrivateEndpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.ManagedPrivateEndpointsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "datafactory.ManagedPrivateEndpointsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.ManagedPrivateEndpointsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ManagedPrivateEndpointsClient) DeletePreparer(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":                autorest.Encode("path", factoryName),
		"managedPrivateEndpointName": autorest.Encode("path", managedPrivateEndpointName),
		"managedVirtualNetworkName":  autorest.Encode("path", managedVirtualNetworkName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/managedVirtualNetworks/{managedVirtualNetworkName}/managedPrivateEndpoints/{managedPrivateEndpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedPrivateEndpointsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ManagedPrivateEndpointsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a managed private endpoint.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// managedVirtualNetworkName - managed virtual network name
// managedPrivateEndpointName - managed private endpoint name
// ifNoneMatch - eTag of the managed private endpoint entity. Should only be specified for get. If the ETag
// matches the existing entity tag, or if * was provided, then no content will be returned.
func (client ManagedPrivateEndpointsClient) Get(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string, ifNoneMatch string) (result ManagedPrivateEndpointResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedPrivateEndpointsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: managedVirtualNetworkName,
			Constraints: []validation.Constraint{{Target: "managedVirtualNetworkName", Name: validation.MaxLength, Rule: 127, Chain: nil},
				{Target: "managedVirtualNetworkName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "managedVirtualNetworkName", Name: validation.Pattern, Rule: `^([_A-Za-z0-9]|([_A-Za-z0-9][-_A-Za-z0-9]{0,125}[_A-Za-z0-9]))$`, Chain: nil}}},
		{TargetValue: managedPrivateEndpointName,
			Constraints: []validation.Constraint{{Target: "managedPrivateEndpointName", Name: validation.MaxLength, Rule: 127, Chain: nil},
				{Target: "managedPrivateEndpointName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "managedPrivateEndpointName", Name: validation.Pattern, Rule: `^([_A-Za-z0-9]|([_A-Za-z0-9][-_A-Za-z0-9]{0,125}[_A-Za-z0-9]))$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.ManagedPrivateEndpointsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, factoryName, managedVirtualNetworkName, managedPrivateEndpointName, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.ManagedPrivateEndpointsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.ManagedPrivateEndpointsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.ManagedPrivateEndpointsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ManagedPrivateEndpointsClient) GetPreparer(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":                autorest.Encode("path", factoryName),
		"managedPrivateEndpointName": autorest.Encode("path", managedPrivateEndpointName),
		"managedVirtualNetworkName":  autorest.Encode("path", managedVirtualNetworkName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/managedVirtualNetworks/{managedVirtualNetworkName}/managedPrivateEndpoints/{managedPrivateEndpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedPrivateEndpointsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ManagedPrivateEndpointsClient) GetResponder(resp *http.Response) (result ManagedPrivateEndpointResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByFactory lists managed private endpoints.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// managedVirtualNetworkName - managed virtual network name
func (client ManagedPrivateEndpointsClient) ListByFactory(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string) (result ManagedPrivateEndpointListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedPrivateEndpointsClient.ListByFactory")
		defer func() {
			sc := -1
			if result.mpelr.Response.Response != nil {
				sc = result.mpelr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: managedVirtualNetworkName,
			Constraints: []validation.Constraint{{Target: "managedVirtualNetworkName", Name: validation.MaxLength, Rule: 127, Chain: nil},
				{Target: "managedVirtualNetworkName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "managedVirtualNetworkName", Name: validation.Pattern, Rule: `^([_A-Za-z0-9]|([_A-Za-z0-9][-_A-Za-z0-9]{0,125}[_A-Za-z0-9]))$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.ManagedPrivateEndpointsClient", "ListByFactory", err.Error())
	}

	result.fn = client.listByFactoryNextResults
	req, err := client.ListByFactoryPreparer(ctx, resourceGroupName, factoryName, managedVirtualNetworkName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.ManagedPrivateEndpointsClient", "ListByFactory", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByFactorySender(req)
	if err != nil {
		result.mpelr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.ManagedPrivateEndpointsClient", "ListByFactory", resp, "Failure sending request")
		return
	}

	result.mpelr, err = client.ListByFactoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.ManagedPrivateEndpointsClient", "ListByFactory", resp, "Failure responding to request")
		return
	}
	if result.mpelr.hasNextLink() && result.mpelr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByFactoryPreparer prepares the ListByFactory request.
func (client ManagedPrivateEndpointsClient) ListByFactoryPreparer(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":               autorest.Encode("path", factoryName),
		"managedVirtualNetworkName": autorest.Encode("path", managedVirtualNetworkName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/managedVirtualNetworks/{managedVirtualNetworkName}/managedPrivateEndpoints", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByFactorySender sends the ListByFactory request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedPrivateEndpointsClient) ListByFactorySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByFactoryResponder handles the response to the ListByFactory request. The method always
// closes the http.Response Body.
func (client ManagedPrivateEndpointsClient) ListByFactoryResponder(resp *http.Response) (result ManagedPrivateEndpointListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByFactoryNextResults retrieves the next set of results, if any.
func (client ManagedPrivateEndpointsClient) listByFactoryNextResults(ctx context.Context, lastResults ManagedPrivateEndpointListResponse) (result ManagedPrivateEndpointListResponse, err error) {
	req, err := lastResults.managedPrivateEndpointListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datafactory.ManagedPrivateEndpointsClient", "listByFactoryNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByFactorySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datafactory.ManagedPrivateEndpointsClient", "listByFactoryNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByFactoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.ManagedPrivateEndpointsClient", "listByFactoryNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByFactoryComplete enumerates all values, automatically crossing page boundaries as required.
func (client ManagedPrivateEndpointsClient) ListByFactoryComplete(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string) (result ManagedPrivateEndpointListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedPrivateEndpointsClient.ListByFactory")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByFactory(ctx, resourceGroupName, factoryName, managedVirtualNetworkName)
	return
}
