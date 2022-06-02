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
 * IBM OpenAPI SDK Code Generator Version: 3.48.1-52130155-20220425-145431
 */

// Package cdtoolchainv2 : Operations and models for the CdToolchainV2 service
package cdtoolchainv2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	common "github.com/IBM/continuous-delivery-go-sdk/common"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
)

// CdToolchainV2 : This swagger document describes the options and endpoints of the early Toolchain API.<br><br> All
// calls require an <strong>Authorization</strong> HTTP header to be set. <br><br> Following are the accepted
// authentication mechanisms and required credentials for each: <ul><li><b>Bearer:</b> an IBM Cloud IAM token
// (authorized for all endpoints)</li><li><b>Basic:</b> 'target_credentials' obtained from Section 3.1.4 or 3.2.5 in <a
// href="https://w3-connections.ibm.com/wikis/home?lang=en-us#!/wiki/W4e7425c664ea_4859_93fb_660b3ab8388b/page/TIAM%20Service%20and%20Broker%20Authentication">TIAM
// Service Vault wiki</a> using 'service_credentials'/'toolchain_credentials' or a fabric (id, secret) pair
// respectively. Refer to the 'Implementation Notes' of each endpoint below to determine how authorization is done based
// on the provided credentials.<br><br>Note: Requests can include the gzip,deflate encoding header.<br>Note: Resources
// can only have <b>one</b> owning organization_guid or resource_group_id, not both. If the organization_guid is
// provided, it will be returned in the response.</li></ul> <br><br> An optional 'transcation-id' header can be
// provided, otherwise, one will be generated and included in the response header.
//
// API Version: 2.0.0
type CdToolchainV2 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://otc-api.us-south.devops.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "cd_toolchain"

// CdToolchainV2Options : Service options
type CdToolchainV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewCdToolchainV2UsingExternalConfig : constructs an instance of CdToolchainV2 with passed in options and external configuration.
func NewCdToolchainV2UsingExternalConfig(options *CdToolchainV2Options) (cdToolchain *CdToolchainV2, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	cdToolchain, err = NewCdToolchainV2(options)
	if err != nil {
		return
	}

	err = cdToolchain.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = cdToolchain.Service.SetServiceURL(options.URL)
	}
	return
}

// NewCdToolchainV2 : constructs an instance of CdToolchainV2 with passed in options.
func NewCdToolchainV2(options *CdToolchainV2Options) (service *CdToolchainV2, err error) {
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

	service = &CdToolchainV2{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	var endpoints = map[string]string{
		"us-south": "https://otc-api.us-south.devops.cloud.ibm.com", // The otc-api enpoint in the us-south region
		"us-east": "https://otc-api.us-east.devops.cloud.ibm.com", // The otc-api enpoint in the us-east region
		"eu-de": "https://otc-api.eu-de.devops.cloud.ibm.com", // The otc-api enpoint in the eu-de region
		"eu-gb": "https://otc-api.eu-gb.devops.cloud.ibm.com", // The otc-api enpoint in the eu-gb region
		"jp-osa": "https://otc-api.jp-osa.devops.cloud.ibm.com", // The otc-api enpoint in the jp-osa region
		"jp-tok": "https://otc-api.jp-tok.devops.cloud.ibm.com", // The otc-api enpoint in the jp-tok region
		"au-syd": "https://otc-api.au-syd.devops.cloud.ibm.com", // The otc-api enpoint in the au-syd region
		"ca-tor": "https://otc-api.ca-tor.devops.cloud.ibm.com", // The otc-api enpoint in the ca-tor region
		"br-sao": "https://otc-api.br-sao.devops.cloud.ibm.com", // The otc-api enpoint in the br-sao region
		"mon01": "https://otc-api.mon01.devops.cloud.ibm.com", // The otc-api enpoint in the mon01 region
		"eu-fr2": "https://otc-api.eu-fr2.devops.cloud.ibm.com", // The otc-api enpoint in the eu-fr2 region
	}

	if url, ok := endpoints[region]; ok {
		return url, nil
	}
	return "", fmt.Errorf("service URL for region '%s' not found", region)
}

// Clone makes a copy of "cdToolchain" suitable for processing requests.
func (cdToolchain *CdToolchainV2) Clone() *CdToolchainV2 {
	if core.IsNil(cdToolchain) {
		return nil
	}
	clone := *cdToolchain
	clone.Service = cdToolchain.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (cdToolchain *CdToolchainV2) SetServiceURL(url string) error {
	return cdToolchain.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (cdToolchain *CdToolchainV2) GetServiceURL() string {
	return cdToolchain.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (cdToolchain *CdToolchainV2) SetDefaultHeaders(headers http.Header) {
	cdToolchain.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (cdToolchain *CdToolchainV2) SetEnableGzipCompression(enableGzip bool) {
	cdToolchain.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (cdToolchain *CdToolchainV2) GetEnableGzipCompression() bool {
	return cdToolchain.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (cdToolchain *CdToolchainV2) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	cdToolchain.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (cdToolchain *CdToolchainV2) DisableRetries() {
	cdToolchain.Service.DisableRetries()
}

// ListToolchains : Returns a list of toolchains
// Returns a list of toolchains that the caller is authorized to access and that meet the provided query parameters.
// <br><br><b>Basic Authorization:</b>'target_credentials' obtained using 'toolchain_credentials'.
func (cdToolchain *CdToolchainV2) ListToolchains(listToolchainsOptions *ListToolchainsOptions) (result *GetToolchainsResponse, response *core.DetailedResponse, err error) {
	return cdToolchain.ListToolchainsWithContext(context.Background(), listToolchainsOptions)
}

// ListToolchainsWithContext is an alternate form of the ListToolchains method which supports a Context parameter
func (cdToolchain *CdToolchainV2) ListToolchainsWithContext(ctx context.Context, listToolchainsOptions *ListToolchainsOptions) (result *GetToolchainsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listToolchainsOptions, "listToolchainsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listToolchainsOptions, "listToolchainsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cdToolchain.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cdToolchain.Service.Options.URL, `/api/v2/toolchains`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listToolchainsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cd_toolchain", "V2", "ListToolchains")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("resource_group_id", fmt.Sprint(*listToolchainsOptions.ResourceGroupID))
	if listToolchainsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listToolchainsOptions.Limit))
	}
	if listToolchainsOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listToolchainsOptions.Offset))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cdToolchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetToolchainsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateToolchain : Create a toolchain
// Creates a new toolchain based off of provided parameters in the POST body. <br><br><b>Basic Authorization:</b>
// Unauthorized.
func (cdToolchain *CdToolchainV2) CreateToolchain(createToolchainOptions *CreateToolchainOptions) (result *PostToolchainResponse, response *core.DetailedResponse, err error) {
	return cdToolchain.CreateToolchainWithContext(context.Background(), createToolchainOptions)
}

// CreateToolchainWithContext is an alternate form of the CreateToolchain method which supports a Context parameter
func (cdToolchain *CdToolchainV2) CreateToolchainWithContext(ctx context.Context, createToolchainOptions *CreateToolchainOptions) (result *PostToolchainResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createToolchainOptions, "createToolchainOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createToolchainOptions, "createToolchainOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cdToolchain.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cdToolchain.Service.Options.URL, `/api/v2/toolchains`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createToolchainOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cd_toolchain", "V2", "CreateToolchain")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createToolchainOptions.Name != nil {
		body["name"] = createToolchainOptions.Name
	}
	if createToolchainOptions.ResourceGroupID != nil {
		body["resource_group_id"] = createToolchainOptions.ResourceGroupID
	}
	if createToolchainOptions.Description != nil {
		body["description"] = createToolchainOptions.Description
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
	response, err = cdToolchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPostToolchainResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetToolchainByID : Fetch a toolchain
// Returns data for a single toolchain identified by id. <br><br><b>Basic Authorization:</b> 'target_credentials'
// obtained using 'toolchain_credentials' scoped to this toolchain.
func (cdToolchain *CdToolchainV2) GetToolchainByID(getToolchainByIDOptions *GetToolchainByIDOptions) (result *GetToolchainByIDResponse, response *core.DetailedResponse, err error) {
	return cdToolchain.GetToolchainByIDWithContext(context.Background(), getToolchainByIDOptions)
}

// GetToolchainByIDWithContext is an alternate form of the GetToolchainByID method which supports a Context parameter
func (cdToolchain *CdToolchainV2) GetToolchainByIDWithContext(ctx context.Context, getToolchainByIDOptions *GetToolchainByIDOptions) (result *GetToolchainByIDResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getToolchainByIDOptions, "getToolchainByIDOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getToolchainByIDOptions, "getToolchainByIDOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"toolchain_id": *getToolchainByIDOptions.ToolchainID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cdToolchain.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cdToolchain.Service.Options.URL, `/api/v2/toolchains/{toolchain_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getToolchainByIDOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cd_toolchain", "V2", "GetToolchainByID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cdToolchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetToolchainByIDResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteToolchain : Delete a toolchain
// Delete the toolchain with the specified ID. <br><br><b>Basic Authorization:</b> Unauthorized.
func (cdToolchain *CdToolchainV2) DeleteToolchain(deleteToolchainOptions *DeleteToolchainOptions) (response *core.DetailedResponse, err error) {
	return cdToolchain.DeleteToolchainWithContext(context.Background(), deleteToolchainOptions)
}

// DeleteToolchainWithContext is an alternate form of the DeleteToolchain method which supports a Context parameter
func (cdToolchain *CdToolchainV2) DeleteToolchainWithContext(ctx context.Context, deleteToolchainOptions *DeleteToolchainOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteToolchainOptions, "deleteToolchainOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteToolchainOptions, "deleteToolchainOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"toolchain_id": *deleteToolchainOptions.ToolchainID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cdToolchain.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cdToolchain.Service.Options.URL, `/api/v2/toolchains/{toolchain_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteToolchainOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cd_toolchain", "V2", "DeleteToolchain")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = cdToolchain.Service.Request(request, nil)

	return
}

// UpdateToolchain : Update a toolchain
// Update the toolchain with the specified ID. <br><br><b>Basic Authorization:</b> Unauthorized.
func (cdToolchain *CdToolchainV2) UpdateToolchain(updateToolchainOptions *UpdateToolchainOptions) (response *core.DetailedResponse, err error) {
	return cdToolchain.UpdateToolchainWithContext(context.Background(), updateToolchainOptions)
}

// UpdateToolchainWithContext is an alternate form of the UpdateToolchain method which supports a Context parameter
func (cdToolchain *CdToolchainV2) UpdateToolchainWithContext(ctx context.Context, updateToolchainOptions *UpdateToolchainOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateToolchainOptions, "updateToolchainOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateToolchainOptions, "updateToolchainOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"toolchain_id": *updateToolchainOptions.ToolchainID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cdToolchain.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cdToolchain.Service.Options.URL, `/api/v2/toolchains/{toolchain_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateToolchainOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cd_toolchain", "V2", "UpdateToolchain")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateToolchainOptions.Name != nil {
		body["name"] = updateToolchainOptions.Name
	}
	if updateToolchainOptions.Description != nil {
		body["description"] = updateToolchainOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = cdToolchain.Service.Request(request, nil)

	return
}

// ListIntegrations : Returns a list of tool integrations bound to toolchain
// Returns a list of tool integrations bound to toolchain that the caller is authorized to access and that meet the
// provided query parameters. <br><br><b>Basic Authorization:</b>'target_credentials' obtained using
// 'toolchain_credentials'.
func (cdToolchain *CdToolchainV2) ListIntegrations(listIntegrationsOptions *ListIntegrationsOptions) (result *GetIntegrationsResponse, response *core.DetailedResponse, err error) {
	return cdToolchain.ListIntegrationsWithContext(context.Background(), listIntegrationsOptions)
}

// ListIntegrationsWithContext is an alternate form of the ListIntegrations method which supports a Context parameter
func (cdToolchain *CdToolchainV2) ListIntegrationsWithContext(ctx context.Context, listIntegrationsOptions *ListIntegrationsOptions) (result *GetIntegrationsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listIntegrationsOptions, "listIntegrationsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listIntegrationsOptions, "listIntegrationsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"toolchain_id": *listIntegrationsOptions.ToolchainID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cdToolchain.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cdToolchain.Service.Options.URL, `/api/v2/toolchains/{toolchain_id}/integrations`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listIntegrationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cd_toolchain", "V2", "ListIntegrations")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listIntegrationsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listIntegrationsOptions.Limit))
	}
	if listIntegrationsOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listIntegrationsOptions.Offset))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cdToolchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetIntegrationsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateIntegration : Create a tool integration
// Provisions a new tool integration based off of provided parameters in the POST body and binds it to the specified
// toolchain. <br><br><b>Basic Authorization:</b> Unauthorized.
func (cdToolchain *CdToolchainV2) CreateIntegration(createIntegrationOptions *CreateIntegrationOptions) (result *PostIntegrationResponse, response *core.DetailedResponse, err error) {
	return cdToolchain.CreateIntegrationWithContext(context.Background(), createIntegrationOptions)
}

// CreateIntegrationWithContext is an alternate form of the CreateIntegration method which supports a Context parameter
func (cdToolchain *CdToolchainV2) CreateIntegrationWithContext(ctx context.Context, createIntegrationOptions *CreateIntegrationOptions) (result *PostIntegrationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createIntegrationOptions, "createIntegrationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createIntegrationOptions, "createIntegrationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"toolchain_id": *createIntegrationOptions.ToolchainID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cdToolchain.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cdToolchain.Service.Options.URL, `/api/v2/toolchains/{toolchain_id}/integrations`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createIntegrationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cd_toolchain", "V2", "CreateIntegration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createIntegrationOptions.ToolID != nil {
		body["tool_id"] = createIntegrationOptions.ToolID
	}
	if createIntegrationOptions.Name != nil {
		body["name"] = createIntegrationOptions.Name
	}
	if createIntegrationOptions.Parameters != nil {
		body["parameters"] = createIntegrationOptions.Parameters
	}
	if createIntegrationOptions.ParametersReferences != nil {
		body["parameters_references"] = createIntegrationOptions.ParametersReferences
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
	response, err = cdToolchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPostIntegrationResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetIntegrationByID : Fetch a tool integration
// Returns a tool integration that is bound to the provided toolchain. <br><br><b>Basic
// Authorization:</b>'target_credentials' obtained using:<br> - 'service_credentials' scoped to this service instance.
// <br> - 'toolchain_credentials' scoped to the toolchain (if any) that this service instance is bound to. <br> - fabric
// (id, secret) pair.
func (cdToolchain *CdToolchainV2) GetIntegrationByID(getIntegrationByIDOptions *GetIntegrationByIDOptions) (result *GetIntegrationByIDResponse, response *core.DetailedResponse, err error) {
	return cdToolchain.GetIntegrationByIDWithContext(context.Background(), getIntegrationByIDOptions)
}

// GetIntegrationByIDWithContext is an alternate form of the GetIntegrationByID method which supports a Context parameter
func (cdToolchain *CdToolchainV2) GetIntegrationByIDWithContext(ctx context.Context, getIntegrationByIDOptions *GetIntegrationByIDOptions) (result *GetIntegrationByIDResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getIntegrationByIDOptions, "getIntegrationByIDOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getIntegrationByIDOptions, "getIntegrationByIDOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"toolchain_id": *getIntegrationByIDOptions.ToolchainID,
		"integration_id": *getIntegrationByIDOptions.IntegrationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cdToolchain.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cdToolchain.Service.Options.URL, `/api/v2/toolchains/{toolchain_id}/integrations/{integration_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getIntegrationByIDOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cd_toolchain", "V2", "GetIntegrationByID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cdToolchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetIntegrationByIDResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteIntegration : Delete a tool integration
// Delete the tool integration with the specified ID. <br><br><b>Basic Authorization:</b> Unauthorized.
func (cdToolchain *CdToolchainV2) DeleteIntegration(deleteIntegrationOptions *DeleteIntegrationOptions) (response *core.DetailedResponse, err error) {
	return cdToolchain.DeleteIntegrationWithContext(context.Background(), deleteIntegrationOptions)
}

// DeleteIntegrationWithContext is an alternate form of the DeleteIntegration method which supports a Context parameter
func (cdToolchain *CdToolchainV2) DeleteIntegrationWithContext(ctx context.Context, deleteIntegrationOptions *DeleteIntegrationOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteIntegrationOptions, "deleteIntegrationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteIntegrationOptions, "deleteIntegrationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"toolchain_id": *deleteIntegrationOptions.ToolchainID,
		"integration_id": *deleteIntegrationOptions.IntegrationID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cdToolchain.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cdToolchain.Service.Options.URL, `/api/v2/toolchains/{toolchain_id}/integrations/{integration_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteIntegrationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cd_toolchain", "V2", "DeleteIntegration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = cdToolchain.Service.Request(request, nil)

	return
}

// UpdateIntegration : Update a tool integration
// Update the tool integration with the specified ID. <br><br><b>Basic Authorization:</b> Unauthorized.
func (cdToolchain *CdToolchainV2) UpdateIntegration(updateIntegrationOptions *UpdateIntegrationOptions) (response *core.DetailedResponse, err error) {
	return cdToolchain.UpdateIntegrationWithContext(context.Background(), updateIntegrationOptions)
}

// UpdateIntegrationWithContext is an alternate form of the UpdateIntegration method which supports a Context parameter
func (cdToolchain *CdToolchainV2) UpdateIntegrationWithContext(ctx context.Context, updateIntegrationOptions *UpdateIntegrationOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateIntegrationOptions, "updateIntegrationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateIntegrationOptions, "updateIntegrationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"toolchain_id": *updateIntegrationOptions.ToolchainID,
		"integration_id": *updateIntegrationOptions.IntegrationID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cdToolchain.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cdToolchain.Service.Options.URL, `/api/v2/toolchains/{toolchain_id}/integrations/{integration_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateIntegrationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cd_toolchain", "V2", "UpdateIntegration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateIntegrationOptions.Name != nil {
		body["name"] = updateIntegrationOptions.Name
	}
	if updateIntegrationOptions.ToolID != nil {
		body["tool_id"] = updateIntegrationOptions.ToolID
	}
	if updateIntegrationOptions.Parameters != nil {
		body["parameters"] = updateIntegrationOptions.Parameters
	}
	if updateIntegrationOptions.ParametersReferences != nil {
		body["parameters_references"] = updateIntegrationOptions.ParametersReferences
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = cdToolchain.Service.Request(request, nil)

	return
}

// CreateIntegrationOptions : The CreateIntegration options.
type CreateIntegrationOptions struct {
	// ID of the toolchain to bind integration to.
	ToolchainID *string `json:"toolchain_id" validate:"required,ne="`

	// The unique short name of the integration that should be provisioned.
	ToolID *string `json:"tool_id" validate:"required"`

	// Name of tool integration.
	Name *string `json:"name,omitempty"`

	// Parameters to be used to create the integration.
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// Decoded values used on provision in the broker that reference fields in the parameters.
	ParametersReferences map[string]interface{} `json:"parameters_references,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateIntegrationOptions : Instantiate CreateIntegrationOptions
func (*CdToolchainV2) NewCreateIntegrationOptions(toolchainID string, toolID string) *CreateIntegrationOptions {
	return &CreateIntegrationOptions{
		ToolchainID: core.StringPtr(toolchainID),
		ToolID: core.StringPtr(toolID),
	}
}

// SetToolchainID : Allow user to set ToolchainID
func (_options *CreateIntegrationOptions) SetToolchainID(toolchainID string) *CreateIntegrationOptions {
	_options.ToolchainID = core.StringPtr(toolchainID)
	return _options
}

// SetToolID : Allow user to set ToolID
func (_options *CreateIntegrationOptions) SetToolID(toolID string) *CreateIntegrationOptions {
	_options.ToolID = core.StringPtr(toolID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateIntegrationOptions) SetName(name string) *CreateIntegrationOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetParameters : Allow user to set Parameters
func (_options *CreateIntegrationOptions) SetParameters(parameters map[string]interface{}) *CreateIntegrationOptions {
	_options.Parameters = parameters
	return _options
}

// SetParametersReferences : Allow user to set ParametersReferences
func (_options *CreateIntegrationOptions) SetParametersReferences(parametersReferences map[string]interface{}) *CreateIntegrationOptions {
	_options.ParametersReferences = parametersReferences
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateIntegrationOptions) SetHeaders(param map[string]string) *CreateIntegrationOptions {
	options.Headers = param
	return options
}

// CreateToolchainOptions : The CreateToolchain options.
type CreateToolchainOptions struct {
	// Toolchain name.
	Name *string `json:"name" validate:"required"`

	// Resource group where toolchain will be created.
	ResourceGroupID *string `json:"resource_group_id" validate:"required"`

	// Describes the toolchain.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateToolchainOptions : Instantiate CreateToolchainOptions
func (*CdToolchainV2) NewCreateToolchainOptions(name string, resourceGroupID string) *CreateToolchainOptions {
	return &CreateToolchainOptions{
		Name: core.StringPtr(name),
		ResourceGroupID: core.StringPtr(resourceGroupID),
	}
}

// SetName : Allow user to set Name
func (_options *CreateToolchainOptions) SetName(name string) *CreateToolchainOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetResourceGroupID : Allow user to set ResourceGroupID
func (_options *CreateToolchainOptions) SetResourceGroupID(resourceGroupID string) *CreateToolchainOptions {
	_options.ResourceGroupID = core.StringPtr(resourceGroupID)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateToolchainOptions) SetDescription(description string) *CreateToolchainOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateToolchainOptions) SetHeaders(param map[string]string) *CreateToolchainOptions {
	options.Headers = param
	return options
}

// DeleteIntegrationOptions : The DeleteIntegration options.
type DeleteIntegrationOptions struct {
	// ID of the toolchain.
	ToolchainID *string `json:"toolchain_id" validate:"required,ne="`

	// ID of the tool integration bound to the toolchain.
	IntegrationID *string `json:"integration_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteIntegrationOptions : Instantiate DeleteIntegrationOptions
func (*CdToolchainV2) NewDeleteIntegrationOptions(toolchainID string, integrationID string) *DeleteIntegrationOptions {
	return &DeleteIntegrationOptions{
		ToolchainID: core.StringPtr(toolchainID),
		IntegrationID: core.StringPtr(integrationID),
	}
}

// SetToolchainID : Allow user to set ToolchainID
func (_options *DeleteIntegrationOptions) SetToolchainID(toolchainID string) *DeleteIntegrationOptions {
	_options.ToolchainID = core.StringPtr(toolchainID)
	return _options
}

// SetIntegrationID : Allow user to set IntegrationID
func (_options *DeleteIntegrationOptions) SetIntegrationID(integrationID string) *DeleteIntegrationOptions {
	_options.IntegrationID = core.StringPtr(integrationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteIntegrationOptions) SetHeaders(param map[string]string) *DeleteIntegrationOptions {
	options.Headers = param
	return options
}

// DeleteToolchainOptions : The DeleteToolchain options.
type DeleteToolchainOptions struct {
	// ID of the toolchain.
	ToolchainID *string `json:"toolchain_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteToolchainOptions : Instantiate DeleteToolchainOptions
func (*CdToolchainV2) NewDeleteToolchainOptions(toolchainID string) *DeleteToolchainOptions {
	return &DeleteToolchainOptions{
		ToolchainID: core.StringPtr(toolchainID),
	}
}

// SetToolchainID : Allow user to set ToolchainID
func (_options *DeleteToolchainOptions) SetToolchainID(toolchainID string) *DeleteToolchainOptions {
	_options.ToolchainID = core.StringPtr(toolchainID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteToolchainOptions) SetHeaders(param map[string]string) *DeleteToolchainOptions {
	options.Headers = param
	return options
}

// GetIntegrationByIDOptions : The GetIntegrationByID options.
type GetIntegrationByIDOptions struct {
	// ID of the toolchain.
	ToolchainID *string `json:"toolchain_id" validate:"required,ne="`

	// ID of the tool integration bound to the toolchain.
	IntegrationID *string `json:"integration_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetIntegrationByIDOptions : Instantiate GetIntegrationByIDOptions
func (*CdToolchainV2) NewGetIntegrationByIDOptions(toolchainID string, integrationID string) *GetIntegrationByIDOptions {
	return &GetIntegrationByIDOptions{
		ToolchainID: core.StringPtr(toolchainID),
		IntegrationID: core.StringPtr(integrationID),
	}
}

// SetToolchainID : Allow user to set ToolchainID
func (_options *GetIntegrationByIDOptions) SetToolchainID(toolchainID string) *GetIntegrationByIDOptions {
	_options.ToolchainID = core.StringPtr(toolchainID)
	return _options
}

// SetIntegrationID : Allow user to set IntegrationID
func (_options *GetIntegrationByIDOptions) SetIntegrationID(integrationID string) *GetIntegrationByIDOptions {
	_options.IntegrationID = core.StringPtr(integrationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetIntegrationByIDOptions) SetHeaders(param map[string]string) *GetIntegrationByIDOptions {
	options.Headers = param
	return options
}

// GetToolchainByIDOptions : The GetToolchainByID options.
type GetToolchainByIDOptions struct {
	// ID of the toolchain.
	ToolchainID *string `json:"toolchain_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetToolchainByIDOptions : Instantiate GetToolchainByIDOptions
func (*CdToolchainV2) NewGetToolchainByIDOptions(toolchainID string) *GetToolchainByIDOptions {
	return &GetToolchainByIDOptions{
		ToolchainID: core.StringPtr(toolchainID),
	}
}

// SetToolchainID : Allow user to set ToolchainID
func (_options *GetToolchainByIDOptions) SetToolchainID(toolchainID string) *GetToolchainByIDOptions {
	_options.ToolchainID = core.StringPtr(toolchainID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetToolchainByIDOptions) SetHeaders(param map[string]string) *GetToolchainByIDOptions {
	options.Headers = param
	return options
}

// ListIntegrationsOptions : The ListIntegrations options.
type ListIntegrationsOptions struct {
	// ID of the toolchain that integrations are bound to.
	ToolchainID *string `json:"toolchain_id" validate:"required,ne="`

	// Limit the number of results. Valid value 0 < limit <= 200.
	Limit *int64 `json:"limit,omitempty"`

	// Offset the number of results from the beginning of the list. Valid value 0 <= offset < 200.
	Offset *int64 `json:"offset,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListIntegrationsOptions : Instantiate ListIntegrationsOptions
func (*CdToolchainV2) NewListIntegrationsOptions(toolchainID string) *ListIntegrationsOptions {
	return &ListIntegrationsOptions{
		ToolchainID: core.StringPtr(toolchainID),
	}
}

// SetToolchainID : Allow user to set ToolchainID
func (_options *ListIntegrationsOptions) SetToolchainID(toolchainID string) *ListIntegrationsOptions {
	_options.ToolchainID = core.StringPtr(toolchainID)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListIntegrationsOptions) SetLimit(limit int64) *ListIntegrationsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *ListIntegrationsOptions) SetOffset(offset int64) *ListIntegrationsOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListIntegrationsOptions) SetHeaders(param map[string]string) *ListIntegrationsOptions {
	options.Headers = param
	return options
}

// ListToolchainsOptions : The ListToolchains options.
type ListToolchainsOptions struct {
	// The resource group ID where the toolchains exist.
	ResourceGroupID *string `json:"resource_group_id" validate:"required"`

	// Limit the number of results. Valid value 0 < limit <= 200.
	Limit *int64 `json:"limit,omitempty"`

	// Offset the number of results from the beginning of the list. Valid value 0 <= offset < 200.
	Offset *int64 `json:"offset,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListToolchainsOptions : Instantiate ListToolchainsOptions
func (*CdToolchainV2) NewListToolchainsOptions(resourceGroupID string) *ListToolchainsOptions {
	return &ListToolchainsOptions{
		ResourceGroupID: core.StringPtr(resourceGroupID),
	}
}

// SetResourceGroupID : Allow user to set ResourceGroupID
func (_options *ListToolchainsOptions) SetResourceGroupID(resourceGroupID string) *ListToolchainsOptions {
	_options.ResourceGroupID = core.StringPtr(resourceGroupID)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListToolchainsOptions) SetLimit(limit int64) *ListToolchainsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *ListToolchainsOptions) SetOffset(offset int64) *ListToolchainsOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListToolchainsOptions) SetHeaders(param map[string]string) *ListToolchainsOptions {
	options.Headers = param
	return options
}

// ToolIntegrationReferent : Information on URIs to access this resource through the UI or API.
type ToolIntegrationReferent struct {
	// URI representing the this resource through the UI.
	UIHref *string `json:"ui_href,omitempty"`

	// URI representing the this resource through an API.
	APIHref *string `json:"api_href,omitempty"`
}

// UnmarshalToolIntegrationReferent unmarshals an instance of ToolIntegrationReferent from the specified map of raw messages.
func UnmarshalToolIntegrationReferent(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ToolIntegrationReferent)
	err = core.UnmarshalPrimitive(m, "ui_href", &obj.UIHref)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "api_href", &obj.APIHref)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateIntegrationOptions : The UpdateIntegration options.
type UpdateIntegrationOptions struct {
	// ID of the toolchain.
	ToolchainID *string `json:"toolchain_id" validate:"required,ne="`

	// ID of the tool integration bound to the toolchain.
	IntegrationID *string `json:"integration_id" validate:"required,ne="`

	// Name of tool integration.
	Name *string `json:"name,omitempty"`

	// The unique short name of the integration that should be provisioned or updated.
	ToolID *string `json:"tool_id,omitempty"`

	// Parameters to be used to create the integration.
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// Decoded values used on provision in the broker that reference fields in the parameters.
	ParametersReferences map[string]interface{} `json:"parameters_references,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateIntegrationOptions : Instantiate UpdateIntegrationOptions
func (*CdToolchainV2) NewUpdateIntegrationOptions(toolchainID string, integrationID string) *UpdateIntegrationOptions {
	return &UpdateIntegrationOptions{
		ToolchainID: core.StringPtr(toolchainID),
		IntegrationID: core.StringPtr(integrationID),
	}
}

// SetToolchainID : Allow user to set ToolchainID
func (_options *UpdateIntegrationOptions) SetToolchainID(toolchainID string) *UpdateIntegrationOptions {
	_options.ToolchainID = core.StringPtr(toolchainID)
	return _options
}

// SetIntegrationID : Allow user to set IntegrationID
func (_options *UpdateIntegrationOptions) SetIntegrationID(integrationID string) *UpdateIntegrationOptions {
	_options.IntegrationID = core.StringPtr(integrationID)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateIntegrationOptions) SetName(name string) *UpdateIntegrationOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetToolID : Allow user to set ToolID
func (_options *UpdateIntegrationOptions) SetToolID(toolID string) *UpdateIntegrationOptions {
	_options.ToolID = core.StringPtr(toolID)
	return _options
}

// SetParameters : Allow user to set Parameters
func (_options *UpdateIntegrationOptions) SetParameters(parameters map[string]interface{}) *UpdateIntegrationOptions {
	_options.Parameters = parameters
	return _options
}

// SetParametersReferences : Allow user to set ParametersReferences
func (_options *UpdateIntegrationOptions) SetParametersReferences(parametersReferences map[string]interface{}) *UpdateIntegrationOptions {
	_options.ParametersReferences = parametersReferences
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateIntegrationOptions) SetHeaders(param map[string]string) *UpdateIntegrationOptions {
	options.Headers = param
	return options
}

// UpdateToolchainOptions : The UpdateToolchain options.
type UpdateToolchainOptions struct {
	// ID of the toolchain.
	ToolchainID *string `json:"toolchain_id" validate:"required,ne="`

	// The name of the toolchain.
	Name *string `json:"name,omitempty"`

	// An optional description.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateToolchainOptions : Instantiate UpdateToolchainOptions
func (*CdToolchainV2) NewUpdateToolchainOptions(toolchainID string) *UpdateToolchainOptions {
	return &UpdateToolchainOptions{
		ToolchainID: core.StringPtr(toolchainID),
	}
}

// SetToolchainID : Allow user to set ToolchainID
func (_options *UpdateToolchainOptions) SetToolchainID(toolchainID string) *UpdateToolchainOptions {
	_options.ToolchainID = core.StringPtr(toolchainID)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateToolchainOptions) SetName(name string) *UpdateToolchainOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateToolchainOptions) SetDescription(description string) *UpdateToolchainOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateToolchainOptions) SetHeaders(param map[string]string) *UpdateToolchainOptions {
	options.Headers = param
	return options
}

// GetIntegrationByIDResponse : Response structure for GET tool integration.
type GetIntegrationByIDResponse struct {
	// Tool integration ID.
	ID *string `json:"id" validate:"required"`

	// Resource group where tool integration can be found.
	ResourceGroupID *string `json:"resource_group_id" validate:"required"`

	// Tool integration CRN.
	CRN *string `json:"crn" validate:"required"`

	// The unique name of the provisioned integration.
	ToolID *string `json:"tool_id" validate:"required"`

	// ID of toolchain which the integration is bound to.
	ToolchainID *string `json:"toolchain_id" validate:"required"`

	// CRN of toolchain which the integration is bound to.
	ToolchainCRN *string `json:"toolchain_crn" validate:"required"`

	// URI representing the tool integration.
	Href *string `json:"href" validate:"required"`

	// Information on URIs to access this resource through the UI or API.
	Referent *ToolIntegrationReferent `json:"referent" validate:"required"`

	// Tool integration name.
	Name *string `json:"name,omitempty"`

	// Latest tool integration update timestamp.
	UpdatedAt *strfmt.DateTime `json:"updated_at" validate:"required"`

	// Parameters to be used to create the integration.
	Parameters map[string]interface{} `json:"parameters" validate:"required"`

	// Current configuration state of the tool integration.
	State *string `json:"state" validate:"required"`
}

// Constants associated with the GetIntegrationByIDResponse.State property.
// Current configuration state of the tool integration.
const (
	GetIntegrationByIDResponseStateConfiguredConst = "configured"
	GetIntegrationByIDResponseStateConfiguringConst = "configuring"
	GetIntegrationByIDResponseStateMisconfiguredConst = "misconfigured"
	GetIntegrationByIDResponseStateUnconfiguredConst = "unconfigured"
)

// UnmarshalGetIntegrationByIDResponse unmarshals an instance of GetIntegrationByIDResponse from the specified map of raw messages.
func UnmarshalGetIntegrationByIDResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetIntegrationByIDResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_id", &obj.ResourceGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.CRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tool_id", &obj.ToolID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "toolchain_id", &obj.ToolchainID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "toolchain_crn", &obj.ToolchainCRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "referent", &obj.Referent, UnmarshalToolIntegrationReferent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameters", &obj.Parameters)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetIntegrationsResponse : Response structure for GET tool integrations.
type GetIntegrationsResponse struct {
	// Maximum number of tool integrations returned from collection.
	Limit *int64 `json:"limit" validate:"required"`

	// Offset applied to tool integrations collection.
	Offset *int64 `json:"offset" validate:"required"`

	// Total number of tool integrations found in collection.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// Information about retrieving first tool integration results from the collection.
	First *GetIntegrationsResponseFirst `json:"first" validate:"required"`

	// Information about retrieving previous tool integration results from the collection.
	Previous *GetIntegrationsResponsePrevious `json:"previous,omitempty"`

	// Information about retrieving next tool integration results from the collection.
	Next *GetIntegrationsResponseNext `json:"next,omitempty"`

	// Information about retrieving last tool integration results from the collection.
	Last *GetIntegrationsResponseLast `json:"last" validate:"required"`

	// Tool integration results returned from the collection.
	Integrations []ToolIntegration `json:"integrations" validate:"required"`
}

// UnmarshalGetIntegrationsResponse unmarshals an instance of GetIntegrationsResponse from the specified map of raw messages.
func UnmarshalGetIntegrationsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetIntegrationsResponse)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalGetIntegrationsResponseFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "previous", &obj.Previous, UnmarshalGetIntegrationsResponsePrevious)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalGetIntegrationsResponseNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalGetIntegrationsResponseLast)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "integrations", &obj.Integrations, UnmarshalToolIntegration)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *GetIntegrationsResponse) GetNextOffset() (*int64, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	offset, err := core.GetQueryParam(resp.Next.Href, "offset")
	if err != nil || offset == nil {
		return nil, err
	}
	var offsetValue int64
	offsetValue, err = strconv.ParseInt(*offset, 10, 64)
	if err != nil {
		return nil, err
	}
	return core.Int64Ptr(offsetValue), nil
}

// GetIntegrationsResponseFirst : Information about retrieving first tool integration results from the collection.
type GetIntegrationsResponseFirst struct {
	// URI that can be used to get first results from the collection.
	Href *string `json:"href,omitempty"`
}

// UnmarshalGetIntegrationsResponseFirst unmarshals an instance of GetIntegrationsResponseFirst from the specified map of raw messages.
func UnmarshalGetIntegrationsResponseFirst(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetIntegrationsResponseFirst)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetIntegrationsResponseLast : Information about retrieving last tool integration results from the collection.
type GetIntegrationsResponseLast struct {
	// URI that can be used to get last results from the collection.
	Href *string `json:"href,omitempty"`
}

// UnmarshalGetIntegrationsResponseLast unmarshals an instance of GetIntegrationsResponseLast from the specified map of raw messages.
func UnmarshalGetIntegrationsResponseLast(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetIntegrationsResponseLast)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetIntegrationsResponseNext : Information about retrieving next tool integration results from the collection.
type GetIntegrationsResponseNext struct {
	// URI that can be used to get next results from the collection.
	Href *string `json:"href,omitempty"`
}

// UnmarshalGetIntegrationsResponseNext unmarshals an instance of GetIntegrationsResponseNext from the specified map of raw messages.
func UnmarshalGetIntegrationsResponseNext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetIntegrationsResponseNext)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetIntegrationsResponsePrevious : Information about retrieving previous tool integration results from the collection.
type GetIntegrationsResponsePrevious struct {
	// URI that can be used to get previous results from the collection.
	Href *string `json:"href,omitempty"`
}

// UnmarshalGetIntegrationsResponsePrevious unmarshals an instance of GetIntegrationsResponsePrevious from the specified map of raw messages.
func UnmarshalGetIntegrationsResponsePrevious(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetIntegrationsResponsePrevious)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetToolchainByIDResponse : Response structure for GET toolchains.
type GetToolchainByIDResponse struct {
	// Toolchain ID.
	ID *string `json:"id" validate:"required"`

	// Toolchain name.
	Name *string `json:"name" validate:"required"`

	// Toolchain description.
	Description *string `json:"description" validate:"required"`

	// Account ID where toolchain can be found.
	AccountID *string `json:"account_id" validate:"required"`

	// Toolchain region.
	Location *string `json:"location" validate:"required"`

	// Resource group where toolchain can be found.
	ResourceGroupID *string `json:"resource_group_id" validate:"required"`

	// Toolchain CRN.
	CRN *string `json:"crn" validate:"required"`

	// URI that can be used to retrieve toolchain.
	Href *string `json:"href" validate:"required"`

	// Toolchain creation timestamp.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// Latest toolchain update timestamp.
	UpdatedAt *strfmt.DateTime `json:"updated_at" validate:"required"`

	// Identity that created the toolchain.
	CreatedBy *string `json:"created_by" validate:"required"`

	// Tags associated with the toolchain.
	Tags []string `json:"tags" validate:"required"`
}

// UnmarshalGetToolchainByIDResponse unmarshals an instance of GetToolchainByIDResponse from the specified map of raw messages.
func UnmarshalGetToolchainByIDResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetToolchainByIDResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_id", &obj.ResourceGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.CRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetToolchainsResponse : Response structure for GET toolchains.
type GetToolchainsResponse struct {
	// Maximum number of toolchains returned from collection.
	Limit *int64 `json:"limit" validate:"required"`

	// Offset applied to toolchains collection.
	Offset *int64 `json:"offset" validate:"required"`

	// Total number of toolchains found in collection.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// Information about retrieving first toolchain results from the collection.
	First *GetToolchainsResponseFirst `json:"first" validate:"required"`

	// Information about retrieving previous toolchain results from the collection.
	Previous *GetToolchainsResponsePrevious `json:"previous,omitempty"`

	// Information about retrieving next toolchain results from the collection.
	Next *GetToolchainsResponseNext `json:"next,omitempty"`

	// Information about retrieving last toolchain results from the collection.
	Last *GetToolchainsResponseLast `json:"last" validate:"required"`

	// Toolchain results returned from the collection.
	Toolchains []Toolchain `json:"toolchains" validate:"required"`
}

// UnmarshalGetToolchainsResponse unmarshals an instance of GetToolchainsResponse from the specified map of raw messages.
func UnmarshalGetToolchainsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetToolchainsResponse)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalGetToolchainsResponseFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "previous", &obj.Previous, UnmarshalGetToolchainsResponsePrevious)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalGetToolchainsResponseNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalGetToolchainsResponseLast)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "toolchains", &obj.Toolchains, UnmarshalToolchain)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *GetToolchainsResponse) GetNextOffset() (*int64, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	offset, err := core.GetQueryParam(resp.Next.Href, "offset")
	if err != nil || offset == nil {
		return nil, err
	}
	var offsetValue int64
	offsetValue, err = strconv.ParseInt(*offset, 10, 64)
	if err != nil {
		return nil, err
	}
	return core.Int64Ptr(offsetValue), nil
}

// GetToolchainsResponseFirst : Information about retrieving first toolchain results from the collection.
type GetToolchainsResponseFirst struct {
	// URI that can be used to get first results from the collection.
	Href *string `json:"href,omitempty"`
}

// UnmarshalGetToolchainsResponseFirst unmarshals an instance of GetToolchainsResponseFirst from the specified map of raw messages.
func UnmarshalGetToolchainsResponseFirst(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetToolchainsResponseFirst)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetToolchainsResponseLast : Information about retrieving last toolchain results from the collection.
type GetToolchainsResponseLast struct {
	// URI that can be used to get last results from the collection.
	Href *string `json:"href,omitempty"`
}

// UnmarshalGetToolchainsResponseLast unmarshals an instance of GetToolchainsResponseLast from the specified map of raw messages.
func UnmarshalGetToolchainsResponseLast(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetToolchainsResponseLast)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetToolchainsResponseNext : Information about retrieving next toolchain results from the collection.
type GetToolchainsResponseNext struct {
	// URI that can be used to get next results from the collection.
	Href *string `json:"href,omitempty"`
}

// UnmarshalGetToolchainsResponseNext unmarshals an instance of GetToolchainsResponseNext from the specified map of raw messages.
func UnmarshalGetToolchainsResponseNext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetToolchainsResponseNext)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetToolchainsResponsePrevious : Information about retrieving previous toolchain results from the collection.
type GetToolchainsResponsePrevious struct {
	// URI that can be used to get previous results from the collection.
	Href *string `json:"href,omitempty"`
}

// UnmarshalGetToolchainsResponsePrevious unmarshals an instance of GetToolchainsResponsePrevious from the specified map of raw messages.
func UnmarshalGetToolchainsResponsePrevious(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetToolchainsResponsePrevious)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PostIntegrationResponse : Response structure for POST tool integration.
type PostIntegrationResponse struct {
	// ID of created tool integration.
	ID *string `json:"id" validate:"required"`

	// Resource group where tool integration was created.
	ResourceGroupID *string `json:"resource_group_id" validate:"required"`

	// CRN of created tool integration.
	CRN *string `json:"crn" validate:"required"`

	// The unique name of the provisioned integration.
	ToolID *string `json:"tool_id" validate:"required"`

	// ID of toolchain which the created integration was bound to.
	ToolchainID *string `json:"toolchain_id" validate:"required"`

	// CRN of toolchain which the created integration was bound to.
	ToolchainCRN *string `json:"toolchain_crn" validate:"required"`

	// URI representing the created tool integration.
	Href *string `json:"href" validate:"required"`

	// Information on URIs to access this resource through the UI or API.
	Referent *Referent `json:"referent" validate:"required"`

	// Name of tool integration.
	Name *string `json:"name,omitempty"`

	// Parameters to be used to create the integration.
	Parameters map[string]interface{} `json:"parameters" validate:"required"`

	// Current configuration state of the tool integration.
	State *string `json:"state" validate:"required"`
}

// Constants associated with the PostIntegrationResponse.State property.
// Current configuration state of the tool integration.
const (
	PostIntegrationResponseStateConfiguredConst = "configured"
	PostIntegrationResponseStateConfiguringConst = "configuring"
	PostIntegrationResponseStateMisconfiguredConst = "misconfigured"
	PostIntegrationResponseStateUnconfiguredConst = "unconfigured"
)

// UnmarshalPostIntegrationResponse unmarshals an instance of PostIntegrationResponse from the specified map of raw messages.
func UnmarshalPostIntegrationResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PostIntegrationResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_id", &obj.ResourceGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.CRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tool_id", &obj.ToolID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "toolchain_id", &obj.ToolchainID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "toolchain_crn", &obj.ToolchainCRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "referent", &obj.Referent, UnmarshalReferent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameters", &obj.Parameters)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PostToolchainResponse : POST toolchains response body.
type PostToolchainResponse struct {
	// ID of created toolchain.
	ID *string `json:"id" validate:"required"`

	// Name of created toolchain.
	Name *string `json:"name" validate:"required"`

	// Description of created toolchain.
	Description *string `json:"description" validate:"required"`

	// Account ID where toolchain was created.
	AccountID *string `json:"account_id" validate:"required"`

	// Region where toolchain is created.
	Location *string `json:"location" validate:"required"`

	// Resource group where toolchain is created.
	ResourceGroupID *string `json:"resource_group_id" validate:"required"`

	// CRN of created toolchain.
	CRN *string `json:"crn" validate:"required"`

	// URI representing the created toolchain.
	Href *string `json:"href" validate:"required"`

	// Identity that created the toolchain.
	CreatedBy *string `json:"created_by" validate:"required"`

	// Tags associated with the created toolchain.
	Tags []string `json:"tags" validate:"required"`
}

// UnmarshalPostToolchainResponse unmarshals an instance of PostToolchainResponse from the specified map of raw messages.
func UnmarshalPostToolchainResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PostToolchainResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_id", &obj.ResourceGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.CRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Referent : Information on URIs to access this resource through the UI or API.
type Referent struct {
	// URI representing the this resource through the UI.
	UIHref *string `json:"ui_href,omitempty"`

	// URI representing the this resource through an API.
	APIHref *string `json:"api_href,omitempty"`
}

// UnmarshalReferent unmarshals an instance of Referent from the specified map of raw messages.
func UnmarshalReferent(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Referent)
	err = core.UnmarshalPrimitive(m, "ui_href", &obj.UIHref)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "api_href", &obj.APIHref)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ToolIntegration : Model describing tool integration resource.
type ToolIntegration struct {
	// Tool integration ID.
	ID *string `json:"id" validate:"required"`

	// Resource group where tool integration can be found.
	ResourceGroupID *string `json:"resource_group_id" validate:"required"`

	// Tool integration CRN.
	CRN *string `json:"crn" validate:"required"`

	// The unique name of the provisioned integration.
	ToolID *string `json:"tool_id" validate:"required"`

	// ID of toolchain which the integration is bound to.
	ToolchainID *string `json:"toolchain_id" validate:"required"`

	// CRN of toolchain which the integration is bound to.
	ToolchainCRN *string `json:"toolchain_crn" validate:"required"`

	// URI representing the tool integration.
	Href *string `json:"href" validate:"required"`

	// Information on URIs to access this resource through the UI or API.
	Referent *ToolIntegrationReferent `json:"referent" validate:"required"`

	// Tool integration name.
	Name *string `json:"name,omitempty"`

	// Latest tool integration update timestamp.
	UpdatedAt *strfmt.DateTime `json:"updated_at" validate:"required"`

	// Parameters to be used to create the integration.
	Parameters map[string]interface{} `json:"parameters" validate:"required"`

	// Current configuration state of the tool integration.
	State *string `json:"state" validate:"required"`
}

// Constants associated with the ToolIntegration.State property.
// Current configuration state of the tool integration.
const (
	ToolIntegrationStateConfiguredConst = "configured"
	ToolIntegrationStateConfiguringConst = "configuring"
	ToolIntegrationStateMisconfiguredConst = "misconfigured"
	ToolIntegrationStateUnconfiguredConst = "unconfigured"
)

// UnmarshalToolIntegration unmarshals an instance of ToolIntegration from the specified map of raw messages.
func UnmarshalToolIntegration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ToolIntegration)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_id", &obj.ResourceGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.CRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tool_id", &obj.ToolID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "toolchain_id", &obj.ToolchainID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "toolchain_crn", &obj.ToolchainCRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "referent", &obj.Referent, UnmarshalToolIntegrationReferent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameters", &obj.Parameters)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Toolchain : Model describing toolchain resource.
type Toolchain struct {
	// Toolchain ID.
	ID *string `json:"id" validate:"required"`

	// Toolchain name.
	Name *string `json:"name" validate:"required"`

	// Toolchain description.
	Description *string `json:"description" validate:"required"`

	// Account ID where toolchain can be found.
	AccountID *string `json:"account_id" validate:"required"`

	// Toolchain region.
	Location *string `json:"location" validate:"required"`

	// Resource group where toolchain can be found.
	ResourceGroupID *string `json:"resource_group_id" validate:"required"`

	// Toolchain CRN.
	CRN *string `json:"crn" validate:"required"`

	// URI that can be used to retrieve toolchain.
	Href *string `json:"href" validate:"required"`

	// Toolchain creation timestamp.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// Latest toolchain update timestamp.
	UpdatedAt *strfmt.DateTime `json:"updated_at" validate:"required"`

	// Identity that created the toolchain.
	CreatedBy *string `json:"created_by" validate:"required"`

	// Tags associated with the toolchain.
	Tags []string `json:"tags" validate:"required"`
}

// UnmarshalToolchain unmarshals an instance of Toolchain from the specified map of raw messages.
func UnmarshalToolchain(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Toolchain)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_id", &obj.ResourceGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.CRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
