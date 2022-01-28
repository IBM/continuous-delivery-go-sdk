/**
 * (C) Copyright IBM Corp. 2022.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 99-SNAPSHOT-f17fbac6-20220127-113737
 */

// Package exampleservicev1 : Operations and models for the ExampleServiceV1 service
package exampleservicev1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.ibm.com/CloudEngineering/go-sdk-template/common"
)

// ExampleServiceV1 : The IBM SDK Squad Example Service. The Example service serves as an example of a service.
//
// API Version: 1.0.0
// See: https://{invalid host}.cloud.ibm.com/apidocs/example-service
type ExampleServiceV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "http://cloud.ibm.com/mysdk/v1"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "example_service"

// ExampleServiceV1Options : Service options
type ExampleServiceV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewExampleServiceV1UsingExternalConfig : constructs an instance of ExampleServiceV1 with passed in options and external configuration.
func NewExampleServiceV1UsingExternalConfig(options *ExampleServiceV1Options) (exampleService *ExampleServiceV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	exampleService, err = NewExampleServiceV1(options)
	if err != nil {
		return
	}

	err = exampleService.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = exampleService.Service.SetServiceURL(options.URL)
	}
	return
}

// NewExampleServiceV1 : constructs an instance of ExampleServiceV1 with passed in options.
func NewExampleServiceV1(options *ExampleServiceV1Options) (service *ExampleServiceV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &ExampleServiceV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "exampleService" suitable for processing requests.
func (exampleService *ExampleServiceV1) Clone() *ExampleServiceV1 {
	if core.IsNil(exampleService) {
		return nil
	}
	clone := *exampleService
	clone.Service = exampleService.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (exampleService *ExampleServiceV1) SetServiceURL(url string) error {
	return exampleService.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (exampleService *ExampleServiceV1) GetServiceURL() string {
	return exampleService.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (exampleService *ExampleServiceV1) SetDefaultHeaders(headers http.Header) {
	exampleService.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (exampleService *ExampleServiceV1) SetEnableGzipCompression(enableGzip bool) {
	exampleService.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (exampleService *ExampleServiceV1) GetEnableGzipCompression() bool {
	return exampleService.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (exampleService *ExampleServiceV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	exampleService.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (exampleService *ExampleServiceV1) DisableRetries() {
	exampleService.Service.DisableRetries()
}

// ListResources : List all resources
func (exampleService *ExampleServiceV1) ListResources(listResourcesOptions *ListResourcesOptions) (result *Resources, response *core.DetailedResponse, err error) {
	return exampleService.ListResourcesWithContext(context.Background(), listResourcesOptions)
}

// ListResourcesWithContext is an alternate form of the ListResources method which supports a Context parameter
func (exampleService *ExampleServiceV1) ListResourcesWithContext(ctx context.Context, listResourcesOptions *ListResourcesOptions) (result *Resources, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listResourcesOptions, "listResourcesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = exampleService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(exampleService.Service.Options.URL, `/resources`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listResourcesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("example_service", "V1", "ListResources")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listResourcesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listResourcesOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = exampleService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateResource : Create a resource
func (exampleService *ExampleServiceV1) CreateResource(createResourceOptions *CreateResourceOptions) (result *Resource, response *core.DetailedResponse, err error) {
	return exampleService.CreateResourceWithContext(context.Background(), createResourceOptions)
}

// CreateResourceWithContext is an alternate form of the CreateResource method which supports a Context parameter
func (exampleService *ExampleServiceV1) CreateResourceWithContext(ctx context.Context, createResourceOptions *CreateResourceOptions) (result *Resource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createResourceOptions, "createResourceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createResourceOptions, "createResourceOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = exampleService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(exampleService.Service.Options.URL, `/resources`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createResourceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("example_service", "V1", "CreateResource")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createResourceOptions.Name != nil {
		body["name"] = createResourceOptions.Name
	}
	if createResourceOptions.Tag != nil {
		body["tag"] = createResourceOptions.Tag
	}
	if createResourceOptions.ResourceID != nil {
		body["resource_id"] = createResourceOptions.ResourceID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = exampleService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetResource : Info for a specific resource
func (exampleService *ExampleServiceV1) GetResource(getResourceOptions *GetResourceOptions) (result *Resource, response *core.DetailedResponse, err error) {
	return exampleService.GetResourceWithContext(context.Background(), getResourceOptions)
}

// GetResourceWithContext is an alternate form of the GetResource method which supports a Context parameter
func (exampleService *ExampleServiceV1) GetResourceWithContext(ctx context.Context, getResourceOptions *GetResourceOptions) (result *Resource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getResourceOptions, "getResourceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getResourceOptions, "getResourceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"resource_id": *getResourceOptions.ResourceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = exampleService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(exampleService.Service.Options.URL, `/resources/{resource_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getResourceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("example_service", "V1", "GetResource")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = exampleService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetResourceEncoded : Info for a specific resource
func (exampleService *ExampleServiceV1) GetResourceEncoded(getResourceEncodedOptions *GetResourceEncodedOptions) (result *Resource, response *core.DetailedResponse, err error) {
	return exampleService.GetResourceEncodedWithContext(context.Background(), getResourceEncodedOptions)
}

// GetResourceEncodedWithContext is an alternate form of the GetResourceEncoded method which supports a Context parameter
func (exampleService *ExampleServiceV1) GetResourceEncodedWithContext(ctx context.Context, getResourceEncodedOptions *GetResourceEncodedOptions) (result *Resource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getResourceEncodedOptions, "getResourceEncodedOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getResourceEncodedOptions, "getResourceEncodedOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"url_encoded_resource_id": *getResourceEncodedOptions.UrlEncodedResourceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = exampleService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(exampleService.Service.Options.URL, `/resources/encoded/{url_encoded_resource_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getResourceEncodedOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("example_service", "V1", "GetResourceEncoded")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = exampleService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateResourceOptions : The CreateResource options.
type CreateResourceOptions struct {
	// The name of the resource.
	Name *string `json:"name" validate:"required"`

	// A tag value for the resource.
	Tag *string `json:"tag" validate:"required"`

	// The id of the resource. If not specified, it will be assigned by the server.
	ResourceID *string `json:"resource_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateResourceOptions : Instantiate CreateResourceOptions
func (*ExampleServiceV1) NewCreateResourceOptions(name string, tag string) *CreateResourceOptions {
	return &CreateResourceOptions{
		Name: core.StringPtr(name),
		Tag: core.StringPtr(tag),
	}
}

// SetName : Allow user to set Name
func (_options *CreateResourceOptions) SetName(name string) *CreateResourceOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetTag : Allow user to set Tag
func (_options *CreateResourceOptions) SetTag(tag string) *CreateResourceOptions {
	_options.Tag = core.StringPtr(tag)
	return _options
}

// SetResourceID : Allow user to set ResourceID
func (_options *CreateResourceOptions) SetResourceID(resourceID string) *CreateResourceOptions {
	_options.ResourceID = core.StringPtr(resourceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateResourceOptions) SetHeaders(param map[string]string) *CreateResourceOptions {
	options.Headers = param
	return options
}

// GetResourceEncodedOptions : The GetResourceEncoded options.
type GetResourceEncodedOptions struct {
	// The id of the resource to retrieve.
	UrlEncodedResourceID *string `json:"url_encoded_resource_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetResourceEncodedOptions : Instantiate GetResourceEncodedOptions
func (*ExampleServiceV1) NewGetResourceEncodedOptions(urlEncodedResourceID string) *GetResourceEncodedOptions {
	return &GetResourceEncodedOptions{
		UrlEncodedResourceID: core.StringPtr(urlEncodedResourceID),
	}
}

// SetUrlEncodedResourceID : Allow user to set UrlEncodedResourceID
func (_options *GetResourceEncodedOptions) SetUrlEncodedResourceID(urlEncodedResourceID string) *GetResourceEncodedOptions {
	_options.UrlEncodedResourceID = core.StringPtr(urlEncodedResourceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetResourceEncodedOptions) SetHeaders(param map[string]string) *GetResourceEncodedOptions {
	options.Headers = param
	return options
}

// GetResourceOptions : The GetResource options.
type GetResourceOptions struct {
	// The id of the resource to retrieve.
	ResourceID *string `json:"resource_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetResourceOptions : Instantiate GetResourceOptions
func (*ExampleServiceV1) NewGetResourceOptions(resourceID string) *GetResourceOptions {
	return &GetResourceOptions{
		ResourceID: core.StringPtr(resourceID),
	}
}

// SetResourceID : Allow user to set ResourceID
func (_options *GetResourceOptions) SetResourceID(resourceID string) *GetResourceOptions {
	_options.ResourceID = core.StringPtr(resourceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetResourceOptions) SetHeaders(param map[string]string) *GetResourceOptions {
	options.Headers = param
	return options
}

// ListResourcesOptions : The ListResources options.
type ListResourcesOptions struct {
	// How many items to return at one time (max 100).
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListResourcesOptions : Instantiate ListResourcesOptions
func (*ExampleServiceV1) NewListResourcesOptions() *ListResourcesOptions {
	return &ListResourcesOptions{}
}

// SetLimit : Allow user to set Limit
func (_options *ListResourcesOptions) SetLimit(limit int64) *ListResourcesOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListResourcesOptions) SetHeaders(param map[string]string) *ListResourcesOptions {
	options.Headers = param
	return options
}

// Resource : A resource.
type Resource struct {
	// The id of the resource. If not specified, it will be assigned by the server.
	ResourceID *string `json:"resource_id,omitempty"`

	// The name of the resource.
	Name *string `json:"name" validate:"required"`

	// A tag value for the resource.
	Tag *string `json:"tag" validate:"required"`

	// This is a read only string.
	ReadOnly *string `json:"read_only,omitempty"`
}

// NewResource : Instantiate Resource (Generic Model Constructor)
func (*ExampleServiceV1) NewResource(name string, tag string) (_model *Resource, err error) {
	_model = &Resource{
		Name: core.StringPtr(name),
		Tag: core.StringPtr(tag),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalResource unmarshals an instance of Resource from the specified map of raw messages.
func UnmarshalResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Resource)
	err = core.UnmarshalPrimitive(m, "resource_id", &obj.ResourceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tag", &obj.Tag)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "read_only", &obj.ReadOnly)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Resources : List of resources.
type Resources struct {
	// Offset value for this portion of the resource list.
	Offset *int64 `json:"offset" validate:"required"`

	// Limit value specified or defaulted in the list_resources request.
	Limit *int64 `json:"limit" validate:"required"`

	// A list of resources.
	Resources []Resource `json:"resources" validate:"required"`
}

// UnmarshalResources unmarshals an instance of Resources from the specified map of raw messages.
func UnmarshalResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Resources)
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalResource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
