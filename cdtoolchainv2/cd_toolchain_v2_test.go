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

package cdtoolchainv2_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/continuous-delivery-go-sdk/cdtoolchainv2"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`CdToolchainV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cdToolchainService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cdToolchainService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
				URL: "https://cdtoolchainv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cdToolchainService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CD_TOOLCHAIN_URL": "https://cdtoolchainv2/api",
				"CD_TOOLCHAIN_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2UsingExternalConfig(&cdtoolchainv2.CdToolchainV2Options{
				})
				Expect(cdToolchainService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := cdToolchainService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != cdToolchainService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(cdToolchainService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(cdToolchainService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2UsingExternalConfig(&cdtoolchainv2.CdToolchainV2Options{
					URL: "https://testService/api",
				})
				Expect(cdToolchainService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := cdToolchainService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != cdToolchainService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(cdToolchainService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(cdToolchainService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2UsingExternalConfig(&cdtoolchainv2.CdToolchainV2Options{
				})
				err := cdToolchainService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := cdToolchainService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != cdToolchainService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(cdToolchainService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(cdToolchainService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CD_TOOLCHAIN_URL": "https://cdtoolchainv2/api",
				"CD_TOOLCHAIN_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2UsingExternalConfig(&cdtoolchainv2.CdToolchainV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cdToolchainService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CD_TOOLCHAIN_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2UsingExternalConfig(&cdtoolchainv2.CdToolchainV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cdToolchainService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = cdtoolchainv2.GetServiceURLForRegion("us-south")
			Expect(url).To(Equal("https://otc-api.us-south.devops.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("us-east")
			Expect(url).To(Equal("https://otc-api.us-east.devops.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("eu-de")
			Expect(url).To(Equal("https://otc-api.eu-de.devops.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("eu-gb")
			Expect(url).To(Equal("https://otc-api.eu-gb.devops.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("jp-osa")
			Expect(url).To(Equal("https://otc-api.jp-osa.devops.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("jp-tok")
			Expect(url).To(Equal("https://otc-api.jp-tok.devops.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("au-syd")
			Expect(url).To(Equal("https://otc-api.au-syd.devops.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("ca-tor")
			Expect(url).To(Equal("https://otc-api.ca-tor.devops.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("br-sao")
			Expect(url).To(Equal("https://otc-api.br-sao.devops.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("mon01")
			Expect(url).To(Equal("https://otc-api.mon01.devops.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("eu-fr2")
			Expect(url).To(Equal("https://otc-api.eu-fr2.devops.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListToolchains(listToolchainsOptions *ListToolchainsOptions) - Operation response error`, func() {
		listToolchainsPath := "/api/v2/toolchains"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listToolchainsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListToolchains with error: Operation response processing error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the ListToolchainsOptions model
				listToolchainsOptionsModel := new(cdtoolchainv2.ListToolchainsOptions)
				listToolchainsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listToolchainsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listToolchainsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listToolchainsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdToolchainService.ListToolchains(listToolchainsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdToolchainService.EnableRetries(0, 0)
				result, response, operationErr = cdToolchainService.ListToolchains(listToolchainsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListToolchains(listToolchainsOptions *ListToolchainsOptions)`, func() {
		listToolchainsPath := "/api/v2/toolchains"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listToolchainsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 3, "offset": 6, "total_count": 12, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "toolchains": [{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "location": "Location", "resource_group_id": "ResourceGroupID", "crn": "CRN", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "tags": ["Tags"]}]}`)
				}))
			})
			It(`Invoke ListToolchains successfully with retries`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())
				cdToolchainService.EnableRetries(0, 0)

				// Construct an instance of the ListToolchainsOptions model
				listToolchainsOptionsModel := new(cdtoolchainv2.ListToolchainsOptions)
				listToolchainsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listToolchainsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listToolchainsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listToolchainsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdToolchainService.ListToolchainsWithContext(ctx, listToolchainsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdToolchainService.DisableRetries()
				result, response, operationErr := cdToolchainService.ListToolchains(listToolchainsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdToolchainService.ListToolchainsWithContext(ctx, listToolchainsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listToolchainsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 3, "offset": 6, "total_count": 12, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "toolchains": [{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "location": "Location", "resource_group_id": "ResourceGroupID", "crn": "CRN", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "tags": ["Tags"]}]}`)
				}))
			})
			It(`Invoke ListToolchains successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdToolchainService.ListToolchains(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListToolchainsOptions model
				listToolchainsOptionsModel := new(cdtoolchainv2.ListToolchainsOptions)
				listToolchainsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listToolchainsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listToolchainsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listToolchainsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdToolchainService.ListToolchains(listToolchainsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListToolchains with error: Operation validation and request error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the ListToolchainsOptions model
				listToolchainsOptionsModel := new(cdtoolchainv2.ListToolchainsOptions)
				listToolchainsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listToolchainsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listToolchainsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listToolchainsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdToolchainService.ListToolchains(listToolchainsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListToolchainsOptions model with no property values
				listToolchainsOptionsModelNew := new(cdtoolchainv2.ListToolchainsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdToolchainService.ListToolchains(listToolchainsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListToolchains successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the ListToolchainsOptions model
				listToolchainsOptionsModel := new(cdtoolchainv2.ListToolchainsOptions)
				listToolchainsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listToolchainsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listToolchainsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listToolchainsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdToolchainService.ListToolchains(listToolchainsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	Context(`Test pagination helper method on response`, func() {
		It(`Invoke GetNextOffset successfully`, func() {
			responseObject := new(cdtoolchainv2.GetToolchainsResponse)
			nextObject := new(cdtoolchainv2.GetToolchainsResponseNext)
			nextObject.Href = core.StringPtr("ibm.com?offset=135")
			responseObject.Next = nextObject

			value, err := responseObject.GetNextOffset()
			Expect(err).To(BeNil())
			Expect(value).To(Equal(core.Int64Ptr(int64(135))))
		})
		It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
			responseObject := new(cdtoolchainv2.GetToolchainsResponse)

			value, err := responseObject.GetNextOffset()
			Expect(err).To(BeNil())
			Expect(value).To(BeNil())
		})
		It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
			responseObject := new(cdtoolchainv2.GetToolchainsResponse)
			nextObject := new(cdtoolchainv2.GetToolchainsResponseNext)
			nextObject.Href = core.StringPtr("ibm.com")
			responseObject.Next = nextObject

			value, err := responseObject.GetNextOffset()
			Expect(err).To(BeNil())
			Expect(value).To(BeNil())
		})
		It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
			responseObject := new(cdtoolchainv2.GetToolchainsResponse)
			nextObject := new(cdtoolchainv2.GetToolchainsResponseNext)
			nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
			responseObject.Next = nextObject

			value, err := responseObject.GetNextOffset()
			Expect(err).NotTo(BeNil())
			Expect(value).To(BeNil())
		})
	})
	})
	Describe(`CreateToolchain(createToolchainOptions *CreateToolchainOptions) - Operation response error`, func() {
		createToolchainPath := "/api/v2/toolchains"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createToolchainPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateToolchain with error: Operation response processing error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the CreateToolchainOptions model
				createToolchainOptionsModel := new(cdtoolchainv2.CreateToolchainOptions)
				createToolchainOptionsModel.Name = core.StringPtr("TestToolchainV2")
				createToolchainOptionsModel.ResourceGroupID = core.StringPtr("6a9a01f2cff54a7f966f803d92877123")
				createToolchainOptionsModel.Description = core.StringPtr("A sample toolchain to test the API")
				createToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdToolchainService.CreateToolchain(createToolchainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdToolchainService.EnableRetries(0, 0)
				result, response, operationErr = cdToolchainService.CreateToolchain(createToolchainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateToolchain(createToolchainOptions *CreateToolchainOptions)`, func() {
		createToolchainPath := "/api/v2/toolchains"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createToolchainPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "location": "Location", "resource_group_id": "ResourceGroupID", "crn": "CRN", "href": "Href", "created_by": "CreatedBy", "tags": ["Tags"]}`)
				}))
			})
			It(`Invoke CreateToolchain successfully with retries`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())
				cdToolchainService.EnableRetries(0, 0)

				// Construct an instance of the CreateToolchainOptions model
				createToolchainOptionsModel := new(cdtoolchainv2.CreateToolchainOptions)
				createToolchainOptionsModel.Name = core.StringPtr("TestToolchainV2")
				createToolchainOptionsModel.ResourceGroupID = core.StringPtr("6a9a01f2cff54a7f966f803d92877123")
				createToolchainOptionsModel.Description = core.StringPtr("A sample toolchain to test the API")
				createToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdToolchainService.CreateToolchainWithContext(ctx, createToolchainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdToolchainService.DisableRetries()
				result, response, operationErr := cdToolchainService.CreateToolchain(createToolchainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdToolchainService.CreateToolchainWithContext(ctx, createToolchainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createToolchainPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "location": "Location", "resource_group_id": "ResourceGroupID", "crn": "CRN", "href": "Href", "created_by": "CreatedBy", "tags": ["Tags"]}`)
				}))
			})
			It(`Invoke CreateToolchain successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdToolchainService.CreateToolchain(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateToolchainOptions model
				createToolchainOptionsModel := new(cdtoolchainv2.CreateToolchainOptions)
				createToolchainOptionsModel.Name = core.StringPtr("TestToolchainV2")
				createToolchainOptionsModel.ResourceGroupID = core.StringPtr("6a9a01f2cff54a7f966f803d92877123")
				createToolchainOptionsModel.Description = core.StringPtr("A sample toolchain to test the API")
				createToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdToolchainService.CreateToolchain(createToolchainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateToolchain with error: Operation validation and request error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the CreateToolchainOptions model
				createToolchainOptionsModel := new(cdtoolchainv2.CreateToolchainOptions)
				createToolchainOptionsModel.Name = core.StringPtr("TestToolchainV2")
				createToolchainOptionsModel.ResourceGroupID = core.StringPtr("6a9a01f2cff54a7f966f803d92877123")
				createToolchainOptionsModel.Description = core.StringPtr("A sample toolchain to test the API")
				createToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdToolchainService.CreateToolchain(createToolchainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateToolchainOptions model with no property values
				createToolchainOptionsModelNew := new(cdtoolchainv2.CreateToolchainOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdToolchainService.CreateToolchain(createToolchainOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateToolchain successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the CreateToolchainOptions model
				createToolchainOptionsModel := new(cdtoolchainv2.CreateToolchainOptions)
				createToolchainOptionsModel.Name = core.StringPtr("TestToolchainV2")
				createToolchainOptionsModel.ResourceGroupID = core.StringPtr("6a9a01f2cff54a7f966f803d92877123")
				createToolchainOptionsModel.Description = core.StringPtr("A sample toolchain to test the API")
				createToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdToolchainService.CreateToolchain(createToolchainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetToolchainByID(getToolchainByIDOptions *GetToolchainByIDOptions) - Operation response error`, func() {
		getToolchainByIDPath := "/api/v2/toolchains/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getToolchainByIDPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetToolchainByID with error: Operation response processing error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the GetToolchainByIDOptions model
				getToolchainByIDOptionsModel := new(cdtoolchainv2.GetToolchainByIDOptions)
				getToolchainByIDOptionsModel.ToolchainID = core.StringPtr("testString")
				getToolchainByIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdToolchainService.GetToolchainByID(getToolchainByIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdToolchainService.EnableRetries(0, 0)
				result, response, operationErr = cdToolchainService.GetToolchainByID(getToolchainByIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetToolchainByID(getToolchainByIDOptions *GetToolchainByIDOptions)`, func() {
		getToolchainByIDPath := "/api/v2/toolchains/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getToolchainByIDPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "location": "Location", "resource_group_id": "ResourceGroupID", "crn": "CRN", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "tags": ["Tags"]}`)
				}))
			})
			It(`Invoke GetToolchainByID successfully with retries`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())
				cdToolchainService.EnableRetries(0, 0)

				// Construct an instance of the GetToolchainByIDOptions model
				getToolchainByIDOptionsModel := new(cdtoolchainv2.GetToolchainByIDOptions)
				getToolchainByIDOptionsModel.ToolchainID = core.StringPtr("testString")
				getToolchainByIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdToolchainService.GetToolchainByIDWithContext(ctx, getToolchainByIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdToolchainService.DisableRetries()
				result, response, operationErr := cdToolchainService.GetToolchainByID(getToolchainByIDOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdToolchainService.GetToolchainByIDWithContext(ctx, getToolchainByIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getToolchainByIDPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "location": "Location", "resource_group_id": "ResourceGroupID", "crn": "CRN", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "tags": ["Tags"]}`)
				}))
			})
			It(`Invoke GetToolchainByID successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdToolchainService.GetToolchainByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetToolchainByIDOptions model
				getToolchainByIDOptionsModel := new(cdtoolchainv2.GetToolchainByIDOptions)
				getToolchainByIDOptionsModel.ToolchainID = core.StringPtr("testString")
				getToolchainByIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdToolchainService.GetToolchainByID(getToolchainByIDOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetToolchainByID with error: Operation validation and request error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the GetToolchainByIDOptions model
				getToolchainByIDOptionsModel := new(cdtoolchainv2.GetToolchainByIDOptions)
				getToolchainByIDOptionsModel.ToolchainID = core.StringPtr("testString")
				getToolchainByIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdToolchainService.GetToolchainByID(getToolchainByIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetToolchainByIDOptions model with no property values
				getToolchainByIDOptionsModelNew := new(cdtoolchainv2.GetToolchainByIDOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdToolchainService.GetToolchainByID(getToolchainByIDOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetToolchainByID successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the GetToolchainByIDOptions model
				getToolchainByIDOptionsModel := new(cdtoolchainv2.GetToolchainByIDOptions)
				getToolchainByIDOptionsModel.ToolchainID = core.StringPtr("testString")
				getToolchainByIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdToolchainService.GetToolchainByID(getToolchainByIDOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteToolchain(deleteToolchainOptions *DeleteToolchainOptions)`, func() {
		deleteToolchainPath := "/api/v2/toolchains/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteToolchainPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteToolchain successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := cdToolchainService.DeleteToolchain(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteToolchainOptions model
				deleteToolchainOptionsModel := new(cdtoolchainv2.DeleteToolchainOptions)
				deleteToolchainOptionsModel.ToolchainID = core.StringPtr("testString")
				deleteToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = cdToolchainService.DeleteToolchain(deleteToolchainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteToolchain with error: Operation validation and request error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the DeleteToolchainOptions model
				deleteToolchainOptionsModel := new(cdtoolchainv2.DeleteToolchainOptions)
				deleteToolchainOptionsModel.ToolchainID = core.StringPtr("testString")
				deleteToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := cdToolchainService.DeleteToolchain(deleteToolchainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteToolchainOptions model with no property values
				deleteToolchainOptionsModelNew := new(cdtoolchainv2.DeleteToolchainOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = cdToolchainService.DeleteToolchain(deleteToolchainOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateToolchain(updateToolchainOptions *UpdateToolchainOptions)`, func() {
		updateToolchainPath := "/api/v2/toolchains/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateToolchainPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(204)
				}))
			})
			It(`Invoke UpdateToolchain successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := cdToolchainService.UpdateToolchain(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateToolchainOptions model
				updateToolchainOptionsModel := new(cdtoolchainv2.UpdateToolchainOptions)
				updateToolchainOptionsModel.ToolchainID = core.StringPtr("testString")
				updateToolchainOptionsModel.Name = core.StringPtr("newToolchainName")
				updateToolchainOptionsModel.Description = core.StringPtr("New toolchain description")
				updateToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = cdToolchainService.UpdateToolchain(updateToolchainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateToolchain with error: Operation validation and request error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the UpdateToolchainOptions model
				updateToolchainOptionsModel := new(cdtoolchainv2.UpdateToolchainOptions)
				updateToolchainOptionsModel.ToolchainID = core.StringPtr("testString")
				updateToolchainOptionsModel.Name = core.StringPtr("newToolchainName")
				updateToolchainOptionsModel.Description = core.StringPtr("New toolchain description")
				updateToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := cdToolchainService.UpdateToolchain(updateToolchainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateToolchainOptions model with no property values
				updateToolchainOptionsModelNew := new(cdtoolchainv2.UpdateToolchainOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = cdToolchainService.UpdateToolchain(updateToolchainOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListIntegrations(listIntegrationsOptions *ListIntegrationsOptions) - Operation response error`, func() {
		listIntegrationsPath := "/api/v2/toolchains/testString/integrations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listIntegrationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListIntegrations with error: Operation response processing error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the ListIntegrationsOptions model
				listIntegrationsOptionsModel := new(cdtoolchainv2.ListIntegrationsOptions)
				listIntegrationsOptionsModel.ToolchainID = core.StringPtr("testString")
				listIntegrationsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listIntegrationsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listIntegrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdToolchainService.ListIntegrations(listIntegrationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdToolchainService.EnableRetries(0, 0)
				result, response, operationErr = cdToolchainService.ListIntegrations(listIntegrationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListIntegrations(listIntegrationsOptions *ListIntegrationsOptions)`, func() {
		listIntegrationsPath := "/api/v2/toolchains/testString/integrations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listIntegrationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "integrations": [{"id": "ID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "tool_id": "ToolID", "toolchain_id": "ToolchainID", "toolchain_crn": "ToolchainCRN", "href": "Href", "referent": {"ui_href": "UIHref", "api_href": "APIHref"}, "name": "Name", "updated_at": "2019-01-01T12:00:00.000Z", "parameters": {"mapKey": "anyValue"}, "state": "configured"}]}`)
				}))
			})
			It(`Invoke ListIntegrations successfully with retries`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())
				cdToolchainService.EnableRetries(0, 0)

				// Construct an instance of the ListIntegrationsOptions model
				listIntegrationsOptionsModel := new(cdtoolchainv2.ListIntegrationsOptions)
				listIntegrationsOptionsModel.ToolchainID = core.StringPtr("testString")
				listIntegrationsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listIntegrationsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listIntegrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdToolchainService.ListIntegrationsWithContext(ctx, listIntegrationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdToolchainService.DisableRetries()
				result, response, operationErr := cdToolchainService.ListIntegrations(listIntegrationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdToolchainService.ListIntegrationsWithContext(ctx, listIntegrationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listIntegrationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "integrations": [{"id": "ID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "tool_id": "ToolID", "toolchain_id": "ToolchainID", "toolchain_crn": "ToolchainCRN", "href": "Href", "referent": {"ui_href": "UIHref", "api_href": "APIHref"}, "name": "Name", "updated_at": "2019-01-01T12:00:00.000Z", "parameters": {"mapKey": "anyValue"}, "state": "configured"}]}`)
				}))
			})
			It(`Invoke ListIntegrations successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdToolchainService.ListIntegrations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListIntegrationsOptions model
				listIntegrationsOptionsModel := new(cdtoolchainv2.ListIntegrationsOptions)
				listIntegrationsOptionsModel.ToolchainID = core.StringPtr("testString")
				listIntegrationsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listIntegrationsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listIntegrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdToolchainService.ListIntegrations(listIntegrationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListIntegrations with error: Operation validation and request error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the ListIntegrationsOptions model
				listIntegrationsOptionsModel := new(cdtoolchainv2.ListIntegrationsOptions)
				listIntegrationsOptionsModel.ToolchainID = core.StringPtr("testString")
				listIntegrationsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listIntegrationsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listIntegrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdToolchainService.ListIntegrations(listIntegrationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListIntegrationsOptions model with no property values
				listIntegrationsOptionsModelNew := new(cdtoolchainv2.ListIntegrationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdToolchainService.ListIntegrations(listIntegrationsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListIntegrations successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the ListIntegrationsOptions model
				listIntegrationsOptionsModel := new(cdtoolchainv2.ListIntegrationsOptions)
				listIntegrationsOptionsModel.ToolchainID = core.StringPtr("testString")
				listIntegrationsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listIntegrationsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listIntegrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdToolchainService.ListIntegrations(listIntegrationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	Context(`Test pagination helper method on response`, func() {
		It(`Invoke GetNextOffset successfully`, func() {
			responseObject := new(cdtoolchainv2.GetIntegrationsResponse)
			nextObject := new(cdtoolchainv2.GetIntegrationsResponseNext)
			nextObject.Href = core.StringPtr("ibm.com?offset=135")
			responseObject.Next = nextObject

			value, err := responseObject.GetNextOffset()
			Expect(err).To(BeNil())
			Expect(value).To(Equal(core.Int64Ptr(int64(135))))
		})
		It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
			responseObject := new(cdtoolchainv2.GetIntegrationsResponse)

			value, err := responseObject.GetNextOffset()
			Expect(err).To(BeNil())
			Expect(value).To(BeNil())
		})
		It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
			responseObject := new(cdtoolchainv2.GetIntegrationsResponse)
			nextObject := new(cdtoolchainv2.GetIntegrationsResponseNext)
			nextObject.Href = core.StringPtr("ibm.com")
			responseObject.Next = nextObject

			value, err := responseObject.GetNextOffset()
			Expect(err).To(BeNil())
			Expect(value).To(BeNil())
		})
		It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
			responseObject := new(cdtoolchainv2.GetIntegrationsResponse)
			nextObject := new(cdtoolchainv2.GetIntegrationsResponseNext)
			nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
			responseObject.Next = nextObject

			value, err := responseObject.GetNextOffset()
			Expect(err).NotTo(BeNil())
			Expect(value).To(BeNil())
		})
	})
	})
	Describe(`CreateIntegration(createIntegrationOptions *CreateIntegrationOptions) - Operation response error`, func() {
		createIntegrationPath := "/api/v2/toolchains/testString/integrations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createIntegrationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateIntegration with error: Operation response processing error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the CreateIntegrationOptions model
				createIntegrationOptionsModel := new(cdtoolchainv2.CreateIntegrationOptions)
				createIntegrationOptionsModel.ToolchainID = core.StringPtr("testString")
				createIntegrationOptionsModel.ToolID = core.StringPtr("todolist")
				createIntegrationOptionsModel.Name = core.StringPtr("testString")
				createIntegrationOptionsModel.Parameters = make(map[string]interface{})
				createIntegrationOptionsModel.ParametersReferences = make(map[string]interface{})
				createIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdToolchainService.CreateIntegration(createIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdToolchainService.EnableRetries(0, 0)
				result, response, operationErr = cdToolchainService.CreateIntegration(createIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateIntegration(createIntegrationOptions *CreateIntegrationOptions)`, func() {
		createIntegrationPath := "/api/v2/toolchains/testString/integrations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createIntegrationPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "tool_id": "ToolID", "toolchain_id": "ToolchainID", "toolchain_crn": "ToolchainCRN", "href": "Href", "referent": {"ui_href": "UIHref", "api_href": "APIHref"}, "name": "MyToolIntegration", "parameters": {"mapKey": "anyValue"}, "state": "configured"}`)
				}))
			})
			It(`Invoke CreateIntegration successfully with retries`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())
				cdToolchainService.EnableRetries(0, 0)

				// Construct an instance of the CreateIntegrationOptions model
				createIntegrationOptionsModel := new(cdtoolchainv2.CreateIntegrationOptions)
				createIntegrationOptionsModel.ToolchainID = core.StringPtr("testString")
				createIntegrationOptionsModel.ToolID = core.StringPtr("todolist")
				createIntegrationOptionsModel.Name = core.StringPtr("testString")
				createIntegrationOptionsModel.Parameters = make(map[string]interface{})
				createIntegrationOptionsModel.ParametersReferences = make(map[string]interface{})
				createIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdToolchainService.CreateIntegrationWithContext(ctx, createIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdToolchainService.DisableRetries()
				result, response, operationErr := cdToolchainService.CreateIntegration(createIntegrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdToolchainService.CreateIntegrationWithContext(ctx, createIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createIntegrationPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "tool_id": "ToolID", "toolchain_id": "ToolchainID", "toolchain_crn": "ToolchainCRN", "href": "Href", "referent": {"ui_href": "UIHref", "api_href": "APIHref"}, "name": "MyToolIntegration", "parameters": {"mapKey": "anyValue"}, "state": "configured"}`)
				}))
			})
			It(`Invoke CreateIntegration successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdToolchainService.CreateIntegration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateIntegrationOptions model
				createIntegrationOptionsModel := new(cdtoolchainv2.CreateIntegrationOptions)
				createIntegrationOptionsModel.ToolchainID = core.StringPtr("testString")
				createIntegrationOptionsModel.ToolID = core.StringPtr("todolist")
				createIntegrationOptionsModel.Name = core.StringPtr("testString")
				createIntegrationOptionsModel.Parameters = make(map[string]interface{})
				createIntegrationOptionsModel.ParametersReferences = make(map[string]interface{})
				createIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdToolchainService.CreateIntegration(createIntegrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateIntegration with error: Operation validation and request error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the CreateIntegrationOptions model
				createIntegrationOptionsModel := new(cdtoolchainv2.CreateIntegrationOptions)
				createIntegrationOptionsModel.ToolchainID = core.StringPtr("testString")
				createIntegrationOptionsModel.ToolID = core.StringPtr("todolist")
				createIntegrationOptionsModel.Name = core.StringPtr("testString")
				createIntegrationOptionsModel.Parameters = make(map[string]interface{})
				createIntegrationOptionsModel.ParametersReferences = make(map[string]interface{})
				createIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdToolchainService.CreateIntegration(createIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateIntegrationOptions model with no property values
				createIntegrationOptionsModelNew := new(cdtoolchainv2.CreateIntegrationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdToolchainService.CreateIntegration(createIntegrationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateIntegration successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the CreateIntegrationOptions model
				createIntegrationOptionsModel := new(cdtoolchainv2.CreateIntegrationOptions)
				createIntegrationOptionsModel.ToolchainID = core.StringPtr("testString")
				createIntegrationOptionsModel.ToolID = core.StringPtr("todolist")
				createIntegrationOptionsModel.Name = core.StringPtr("testString")
				createIntegrationOptionsModel.Parameters = make(map[string]interface{})
				createIntegrationOptionsModel.ParametersReferences = make(map[string]interface{})
				createIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdToolchainService.CreateIntegration(createIntegrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetIntegrationByID(getIntegrationByIDOptions *GetIntegrationByIDOptions) - Operation response error`, func() {
		getIntegrationByIDPath := "/api/v2/toolchains/testString/integrations/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getIntegrationByIDPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetIntegrationByID with error: Operation response processing error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the GetIntegrationByIDOptions model
				getIntegrationByIDOptionsModel := new(cdtoolchainv2.GetIntegrationByIDOptions)
				getIntegrationByIDOptionsModel.ToolchainID = core.StringPtr("testString")
				getIntegrationByIDOptionsModel.IntegrationID = core.StringPtr("testString")
				getIntegrationByIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdToolchainService.GetIntegrationByID(getIntegrationByIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdToolchainService.EnableRetries(0, 0)
				result, response, operationErr = cdToolchainService.GetIntegrationByID(getIntegrationByIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetIntegrationByID(getIntegrationByIDOptions *GetIntegrationByIDOptions)`, func() {
		getIntegrationByIDPath := "/api/v2/toolchains/testString/integrations/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getIntegrationByIDPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "tool_id": "ToolID", "toolchain_id": "ToolchainID", "toolchain_crn": "ToolchainCRN", "href": "Href", "referent": {"ui_href": "UIHref", "api_href": "APIHref"}, "name": "Name", "updated_at": "2019-01-01T12:00:00.000Z", "parameters": {"mapKey": "anyValue"}, "state": "configured"}`)
				}))
			})
			It(`Invoke GetIntegrationByID successfully with retries`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())
				cdToolchainService.EnableRetries(0, 0)

				// Construct an instance of the GetIntegrationByIDOptions model
				getIntegrationByIDOptionsModel := new(cdtoolchainv2.GetIntegrationByIDOptions)
				getIntegrationByIDOptionsModel.ToolchainID = core.StringPtr("testString")
				getIntegrationByIDOptionsModel.IntegrationID = core.StringPtr("testString")
				getIntegrationByIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdToolchainService.GetIntegrationByIDWithContext(ctx, getIntegrationByIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdToolchainService.DisableRetries()
				result, response, operationErr := cdToolchainService.GetIntegrationByID(getIntegrationByIDOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdToolchainService.GetIntegrationByIDWithContext(ctx, getIntegrationByIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getIntegrationByIDPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "tool_id": "ToolID", "toolchain_id": "ToolchainID", "toolchain_crn": "ToolchainCRN", "href": "Href", "referent": {"ui_href": "UIHref", "api_href": "APIHref"}, "name": "Name", "updated_at": "2019-01-01T12:00:00.000Z", "parameters": {"mapKey": "anyValue"}, "state": "configured"}`)
				}))
			})
			It(`Invoke GetIntegrationByID successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdToolchainService.GetIntegrationByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetIntegrationByIDOptions model
				getIntegrationByIDOptionsModel := new(cdtoolchainv2.GetIntegrationByIDOptions)
				getIntegrationByIDOptionsModel.ToolchainID = core.StringPtr("testString")
				getIntegrationByIDOptionsModel.IntegrationID = core.StringPtr("testString")
				getIntegrationByIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdToolchainService.GetIntegrationByID(getIntegrationByIDOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetIntegrationByID with error: Operation validation and request error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the GetIntegrationByIDOptions model
				getIntegrationByIDOptionsModel := new(cdtoolchainv2.GetIntegrationByIDOptions)
				getIntegrationByIDOptionsModel.ToolchainID = core.StringPtr("testString")
				getIntegrationByIDOptionsModel.IntegrationID = core.StringPtr("testString")
				getIntegrationByIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdToolchainService.GetIntegrationByID(getIntegrationByIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetIntegrationByIDOptions model with no property values
				getIntegrationByIDOptionsModelNew := new(cdtoolchainv2.GetIntegrationByIDOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdToolchainService.GetIntegrationByID(getIntegrationByIDOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetIntegrationByID successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the GetIntegrationByIDOptions model
				getIntegrationByIDOptionsModel := new(cdtoolchainv2.GetIntegrationByIDOptions)
				getIntegrationByIDOptionsModel.ToolchainID = core.StringPtr("testString")
				getIntegrationByIDOptionsModel.IntegrationID = core.StringPtr("testString")
				getIntegrationByIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdToolchainService.GetIntegrationByID(getIntegrationByIDOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteIntegration(deleteIntegrationOptions *DeleteIntegrationOptions)`, func() {
		deleteIntegrationPath := "/api/v2/toolchains/testString/integrations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteIntegrationPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteIntegration successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := cdToolchainService.DeleteIntegration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteIntegrationOptions model
				deleteIntegrationOptionsModel := new(cdtoolchainv2.DeleteIntegrationOptions)
				deleteIntegrationOptionsModel.ToolchainID = core.StringPtr("testString")
				deleteIntegrationOptionsModel.IntegrationID = core.StringPtr("testString")
				deleteIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = cdToolchainService.DeleteIntegration(deleteIntegrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteIntegration with error: Operation validation and request error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the DeleteIntegrationOptions model
				deleteIntegrationOptionsModel := new(cdtoolchainv2.DeleteIntegrationOptions)
				deleteIntegrationOptionsModel.ToolchainID = core.StringPtr("testString")
				deleteIntegrationOptionsModel.IntegrationID = core.StringPtr("testString")
				deleteIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := cdToolchainService.DeleteIntegration(deleteIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteIntegrationOptions model with no property values
				deleteIntegrationOptionsModelNew := new(cdtoolchainv2.DeleteIntegrationOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = cdToolchainService.DeleteIntegration(deleteIntegrationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateIntegration(updateIntegrationOptions *UpdateIntegrationOptions)`, func() {
		updateIntegrationPath := "/api/v2/toolchains/testString/integrations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateIntegrationPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(204)
				}))
			})
			It(`Invoke UpdateIntegration successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := cdToolchainService.UpdateIntegration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateIntegrationOptions model
				updateIntegrationOptionsModel := new(cdtoolchainv2.UpdateIntegrationOptions)
				updateIntegrationOptionsModel.ToolchainID = core.StringPtr("testString")
				updateIntegrationOptionsModel.IntegrationID = core.StringPtr("testString")
				updateIntegrationOptionsModel.Name = core.StringPtr("MyToolIntegration")
				updateIntegrationOptionsModel.ToolID = core.StringPtr("todolist")
				updateIntegrationOptionsModel.Parameters = make(map[string]interface{})
				updateIntegrationOptionsModel.ParametersReferences = make(map[string]interface{})
				updateIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = cdToolchainService.UpdateIntegration(updateIntegrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateIntegration with error: Operation validation and request error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the UpdateIntegrationOptions model
				updateIntegrationOptionsModel := new(cdtoolchainv2.UpdateIntegrationOptions)
				updateIntegrationOptionsModel.ToolchainID = core.StringPtr("testString")
				updateIntegrationOptionsModel.IntegrationID = core.StringPtr("testString")
				updateIntegrationOptionsModel.Name = core.StringPtr("MyToolIntegration")
				updateIntegrationOptionsModel.ToolID = core.StringPtr("todolist")
				updateIntegrationOptionsModel.Parameters = make(map[string]interface{})
				updateIntegrationOptionsModel.ParametersReferences = make(map[string]interface{})
				updateIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := cdToolchainService.UpdateIntegration(updateIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateIntegrationOptions model with no property values
				updateIntegrationOptionsModelNew := new(cdtoolchainv2.UpdateIntegrationOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = cdToolchainService.UpdateIntegration(updateIntegrationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			cdToolchainService, _ := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
				URL:           "http://cdtoolchainv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateIntegrationOptions successfully`, func() {
				// Construct an instance of the CreateIntegrationOptions model
				toolchainID := "testString"
				createIntegrationOptionsToolID := "todolist"
				createIntegrationOptionsModel := cdToolchainService.NewCreateIntegrationOptions(toolchainID, createIntegrationOptionsToolID)
				createIntegrationOptionsModel.SetToolchainID("testString")
				createIntegrationOptionsModel.SetToolID("todolist")
				createIntegrationOptionsModel.SetName("testString")
				createIntegrationOptionsModel.SetParameters(make(map[string]interface{}))
				createIntegrationOptionsModel.SetParametersReferences(make(map[string]interface{}))
				createIntegrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createIntegrationOptionsModel).ToNot(BeNil())
				Expect(createIntegrationOptionsModel.ToolchainID).To(Equal(core.StringPtr("testString")))
				Expect(createIntegrationOptionsModel.ToolID).To(Equal(core.StringPtr("todolist")))
				Expect(createIntegrationOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createIntegrationOptionsModel.Parameters).To(Equal(make(map[string]interface{})))
				Expect(createIntegrationOptionsModel.ParametersReferences).To(Equal(make(map[string]interface{})))
				Expect(createIntegrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateToolchainOptions successfully`, func() {
				// Construct an instance of the CreateToolchainOptions model
				createToolchainOptionsName := "TestToolchainV2"
				createToolchainOptionsResourceGroupID := "6a9a01f2cff54a7f966f803d92877123"
				createToolchainOptionsModel := cdToolchainService.NewCreateToolchainOptions(createToolchainOptionsName, createToolchainOptionsResourceGroupID)
				createToolchainOptionsModel.SetName("TestToolchainV2")
				createToolchainOptionsModel.SetResourceGroupID("6a9a01f2cff54a7f966f803d92877123")
				createToolchainOptionsModel.SetDescription("A sample toolchain to test the API")
				createToolchainOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createToolchainOptionsModel).ToNot(BeNil())
				Expect(createToolchainOptionsModel.Name).To(Equal(core.StringPtr("TestToolchainV2")))
				Expect(createToolchainOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("6a9a01f2cff54a7f966f803d92877123")))
				Expect(createToolchainOptionsModel.Description).To(Equal(core.StringPtr("A sample toolchain to test the API")))
				Expect(createToolchainOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteIntegrationOptions successfully`, func() {
				// Construct an instance of the DeleteIntegrationOptions model
				toolchainID := "testString"
				integrationID := "testString"
				deleteIntegrationOptionsModel := cdToolchainService.NewDeleteIntegrationOptions(toolchainID, integrationID)
				deleteIntegrationOptionsModel.SetToolchainID("testString")
				deleteIntegrationOptionsModel.SetIntegrationID("testString")
				deleteIntegrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteIntegrationOptionsModel).ToNot(BeNil())
				Expect(deleteIntegrationOptionsModel.ToolchainID).To(Equal(core.StringPtr("testString")))
				Expect(deleteIntegrationOptionsModel.IntegrationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteIntegrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteToolchainOptions successfully`, func() {
				// Construct an instance of the DeleteToolchainOptions model
				toolchainID := "testString"
				deleteToolchainOptionsModel := cdToolchainService.NewDeleteToolchainOptions(toolchainID)
				deleteToolchainOptionsModel.SetToolchainID("testString")
				deleteToolchainOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteToolchainOptionsModel).ToNot(BeNil())
				Expect(deleteToolchainOptionsModel.ToolchainID).To(Equal(core.StringPtr("testString")))
				Expect(deleteToolchainOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetIntegrationByIDOptions successfully`, func() {
				// Construct an instance of the GetIntegrationByIDOptions model
				toolchainID := "testString"
				integrationID := "testString"
				getIntegrationByIDOptionsModel := cdToolchainService.NewGetIntegrationByIDOptions(toolchainID, integrationID)
				getIntegrationByIDOptionsModel.SetToolchainID("testString")
				getIntegrationByIDOptionsModel.SetIntegrationID("testString")
				getIntegrationByIDOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getIntegrationByIDOptionsModel).ToNot(BeNil())
				Expect(getIntegrationByIDOptionsModel.ToolchainID).To(Equal(core.StringPtr("testString")))
				Expect(getIntegrationByIDOptionsModel.IntegrationID).To(Equal(core.StringPtr("testString")))
				Expect(getIntegrationByIDOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetToolchainByIDOptions successfully`, func() {
				// Construct an instance of the GetToolchainByIDOptions model
				toolchainID := "testString"
				getToolchainByIDOptionsModel := cdToolchainService.NewGetToolchainByIDOptions(toolchainID)
				getToolchainByIDOptionsModel.SetToolchainID("testString")
				getToolchainByIDOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getToolchainByIDOptionsModel).ToNot(BeNil())
				Expect(getToolchainByIDOptionsModel.ToolchainID).To(Equal(core.StringPtr("testString")))
				Expect(getToolchainByIDOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListIntegrationsOptions successfully`, func() {
				// Construct an instance of the ListIntegrationsOptions model
				toolchainID := "testString"
				listIntegrationsOptionsModel := cdToolchainService.NewListIntegrationsOptions(toolchainID)
				listIntegrationsOptionsModel.SetToolchainID("testString")
				listIntegrationsOptionsModel.SetLimit(int64(1))
				listIntegrationsOptionsModel.SetOffset(int64(0))
				listIntegrationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listIntegrationsOptionsModel).ToNot(BeNil())
				Expect(listIntegrationsOptionsModel.ToolchainID).To(Equal(core.StringPtr("testString")))
				Expect(listIntegrationsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listIntegrationsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listIntegrationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListToolchainsOptions successfully`, func() {
				// Construct an instance of the ListToolchainsOptions model
				resourceGroupID := "testString"
				listToolchainsOptionsModel := cdToolchainService.NewListToolchainsOptions(resourceGroupID)
				listToolchainsOptionsModel.SetResourceGroupID("testString")
				listToolchainsOptionsModel.SetLimit(int64(1))
				listToolchainsOptionsModel.SetOffset(int64(0))
				listToolchainsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listToolchainsOptionsModel).ToNot(BeNil())
				Expect(listToolchainsOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listToolchainsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listToolchainsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listToolchainsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateIntegrationOptions successfully`, func() {
				// Construct an instance of the UpdateIntegrationOptions model
				toolchainID := "testString"
				integrationID := "testString"
				updateIntegrationOptionsModel := cdToolchainService.NewUpdateIntegrationOptions(toolchainID, integrationID)
				updateIntegrationOptionsModel.SetToolchainID("testString")
				updateIntegrationOptionsModel.SetIntegrationID("testString")
				updateIntegrationOptionsModel.SetName("MyToolIntegration")
				updateIntegrationOptionsModel.SetToolID("todolist")
				updateIntegrationOptionsModel.SetParameters(make(map[string]interface{}))
				updateIntegrationOptionsModel.SetParametersReferences(make(map[string]interface{}))
				updateIntegrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateIntegrationOptionsModel).ToNot(BeNil())
				Expect(updateIntegrationOptionsModel.ToolchainID).To(Equal(core.StringPtr("testString")))
				Expect(updateIntegrationOptionsModel.IntegrationID).To(Equal(core.StringPtr("testString")))
				Expect(updateIntegrationOptionsModel.Name).To(Equal(core.StringPtr("MyToolIntegration")))
				Expect(updateIntegrationOptionsModel.ToolID).To(Equal(core.StringPtr("todolist")))
				Expect(updateIntegrationOptionsModel.Parameters).To(Equal(make(map[string]interface{})))
				Expect(updateIntegrationOptionsModel.ParametersReferences).To(Equal(make(map[string]interface{})))
				Expect(updateIntegrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateToolchainOptions successfully`, func() {
				// Construct an instance of the UpdateToolchainOptions model
				toolchainID := "testString"
				updateToolchainOptionsModel := cdToolchainService.NewUpdateToolchainOptions(toolchainID)
				updateToolchainOptionsModel.SetToolchainID("testString")
				updateToolchainOptionsModel.SetName("newToolchainName")
				updateToolchainOptionsModel.SetDescription("New toolchain description")
				updateToolchainOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateToolchainOptionsModel).ToNot(BeNil())
				Expect(updateToolchainOptionsModel.ToolchainID).To(Equal(core.StringPtr("testString")))
				Expect(updateToolchainOptionsModel.Name).To(Equal(core.StringPtr("newToolchainName")))
				Expect(updateToolchainOptionsModel.Description).To(Equal(core.StringPtr("New toolchain description")))
				Expect(updateToolchainOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
