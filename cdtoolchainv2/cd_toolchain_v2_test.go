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
			Expect(url).To(Equal("https://api.us-south.devops.cloud.ibm.com/toolchain/v2"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("us-east")
			Expect(url).To(Equal("https://api.us-east.devops.cloud.ibm.com/toolchain/v2"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("eu-de")
			Expect(url).To(Equal("https://api.eu-de.devops.cloud.ibm.com/toolchain/v2"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("eu-gb")
			Expect(url).To(Equal("https://api.eu-gb.devops.cloud.ibm.com/toolchain/v2"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("jp-osa")
			Expect(url).To(Equal("https://api.jp-osa.devops.cloud.ibm.com/toolchain/v2"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("jp-tok")
			Expect(url).To(Equal("https://api.jp-tok.devops.cloud.ibm.com/toolchain/v2"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("au-syd")
			Expect(url).To(Equal("https://api.au-syd.devops.cloud.ibm.com/toolchain/v2"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("ca-tor")
			Expect(url).To(Equal("https://api.ca-tor.devops.cloud.ibm.com/toolchain/v2"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("br-sao")
			Expect(url).To(Equal("https://api.br-sao.devops.cloud.ibm.com/toolchain/v2"))
			Expect(err).To(BeNil())

			url, err = cdtoolchainv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListToolchains(listToolchainsOptions *ListToolchainsOptions) - Operation response error`, func() {
		listToolchainsPath := "/toolchains"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listToolchainsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
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
				listToolchainsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listToolchainsOptionsModel.Start = core.StringPtr("testString")
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
		listToolchainsPath := "/toolchains"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listToolchainsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "limit": 5, "first": {"href": "Href"}, "previous": {"start": "Start", "href": "Href"}, "next": {"start": "Start", "href": "Href"}, "last": {"start": "Start", "href": "Href"}, "toolchains": [{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "location": "Location", "resource_group_id": "ResourceGroupID", "crn": "CRN", "href": "Href", "ui_href": "UIHref", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "tags": ["Tags"]}]}`)
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
				listToolchainsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listToolchainsOptionsModel.Start = core.StringPtr("testString")
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
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "limit": 5, "first": {"href": "Href"}, "previous": {"start": "Start", "href": "Href"}, "next": {"start": "Start", "href": "Href"}, "last": {"start": "Start", "href": "Href"}, "toolchains": [{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "location": "Location", "resource_group_id": "ResourceGroupID", "crn": "CRN", "href": "Href", "ui_href": "UIHref", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "tags": ["Tags"]}]}`)
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
				listToolchainsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listToolchainsOptionsModel.Start = core.StringPtr("testString")
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
				listToolchainsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listToolchainsOptionsModel.Start = core.StringPtr("testString")
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
				listToolchainsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listToolchainsOptionsModel.Start = core.StringPtr("testString")
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
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(cdtoolchainv2.ToolchainCollection)
				nextObject := new(cdtoolchainv2.ToolchainCollectionNext)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(cdtoolchainv2.ToolchainCollection)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listToolchainsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"toolchains":[{"id":"ID","name":"Name","description":"Description","account_id":"AccountID","location":"Location","resource_group_id":"ResourceGroupID","crn":"CRN","href":"Href","ui_href":"UIHref","created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","tags":["Tags"]}],"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"toolchains":[{"id":"ID","name":"Name","description":"Description","account_id":"AccountID","location":"Location","resource_group_id":"ResourceGroupID","crn":"CRN","href":"Href","ui_href":"UIHref","created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","tags":["Tags"]}],"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ToolchainsPager.GetNext successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				listToolchainsOptionsModel := &cdtoolchainv2.ListToolchainsOptions{
					ResourceGroupID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := cdToolchainService.NewToolchainsPager(listToolchainsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []cdtoolchainv2.ToolchainModel
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ToolchainsPager.GetAll successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				listToolchainsOptionsModel := &cdtoolchainv2.ListToolchainsOptions{
					ResourceGroupID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := cdToolchainService.NewToolchainsPager(listToolchainsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateToolchain(createToolchainOptions *CreateToolchainOptions) - Operation response error`, func() {
		createToolchainPath := "/toolchains"
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
		createToolchainPath := "/toolchains"
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "location": "Location", "resource_group_id": "ResourceGroupID", "crn": "CRN", "href": "Href", "ui_href": "UIHref", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "tags": ["Tags"]}`)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "location": "Location", "resource_group_id": "ResourceGroupID", "crn": "CRN", "href": "Href", "ui_href": "UIHref", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "tags": ["Tags"]}`)
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
		getToolchainByIDPath := "/toolchains/testString"
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
		getToolchainByIDPath := "/toolchains/testString"
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "location": "Location", "resource_group_id": "ResourceGroupID", "crn": "CRN", "href": "Href", "ui_href": "UIHref", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "tags": ["Tags"]}`)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "location": "Location", "resource_group_id": "ResourceGroupID", "crn": "CRN", "href": "Href", "ui_href": "UIHref", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "tags": ["Tags"]}`)
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
		deleteToolchainPath := "/toolchains/testString"
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
	Describe(`UpdateToolchain(updateToolchainOptions *UpdateToolchainOptions) - Operation response error`, func() {
		updateToolchainPath := "/toolchains/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateToolchainPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateToolchain with error: Operation response processing error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the ToolchainPrototypePatch model
				toolchainPrototypePatchModel := new(cdtoolchainv2.ToolchainPrototypePatch)
				toolchainPrototypePatchModel.Name = core.StringPtr("newToolchainName")
				toolchainPrototypePatchModel.Description = core.StringPtr("New toolchain description")
				toolchainPrototypePatchModelAsPatch, asPatchErr := toolchainPrototypePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateToolchainOptions model
				updateToolchainOptionsModel := new(cdtoolchainv2.UpdateToolchainOptions)
				updateToolchainOptionsModel.ToolchainID = core.StringPtr("testString")
				updateToolchainOptionsModel.ToolchainPrototypePatch = toolchainPrototypePatchModelAsPatch
				updateToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdToolchainService.UpdateToolchain(updateToolchainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdToolchainService.EnableRetries(0, 0)
				result, response, operationErr = cdToolchainService.UpdateToolchain(updateToolchainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateToolchain(updateToolchainOptions *UpdateToolchainOptions)`, func() {
		updateToolchainPath := "/toolchains/testString"
		Context(`Using mock server endpoint with timeout`, func() {
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "location": "Location", "resource_group_id": "ResourceGroupID", "crn": "CRN", "href": "Href", "ui_href": "UIHref", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "tags": ["Tags"]}`)
				}))
			})
			It(`Invoke UpdateToolchain successfully with retries`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())
				cdToolchainService.EnableRetries(0, 0)

				// Construct an instance of the ToolchainPrototypePatch model
				toolchainPrototypePatchModel := new(cdtoolchainv2.ToolchainPrototypePatch)
				toolchainPrototypePatchModel.Name = core.StringPtr("newToolchainName")
				toolchainPrototypePatchModel.Description = core.StringPtr("New toolchain description")
				toolchainPrototypePatchModelAsPatch, asPatchErr := toolchainPrototypePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateToolchainOptions model
				updateToolchainOptionsModel := new(cdtoolchainv2.UpdateToolchainOptions)
				updateToolchainOptionsModel.ToolchainID = core.StringPtr("testString")
				updateToolchainOptionsModel.ToolchainPrototypePatch = toolchainPrototypePatchModelAsPatch
				updateToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdToolchainService.UpdateToolchainWithContext(ctx, updateToolchainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdToolchainService.DisableRetries()
				result, response, operationErr := cdToolchainService.UpdateToolchain(updateToolchainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdToolchainService.UpdateToolchainWithContext(ctx, updateToolchainOptionsModel)
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "location": "Location", "resource_group_id": "ResourceGroupID", "crn": "CRN", "href": "Href", "ui_href": "UIHref", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "tags": ["Tags"]}`)
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
				result, response, operationErr := cdToolchainService.UpdateToolchain(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ToolchainPrototypePatch model
				toolchainPrototypePatchModel := new(cdtoolchainv2.ToolchainPrototypePatch)
				toolchainPrototypePatchModel.Name = core.StringPtr("newToolchainName")
				toolchainPrototypePatchModel.Description = core.StringPtr("New toolchain description")
				toolchainPrototypePatchModelAsPatch, asPatchErr := toolchainPrototypePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateToolchainOptions model
				updateToolchainOptionsModel := new(cdtoolchainv2.UpdateToolchainOptions)
				updateToolchainOptionsModel.ToolchainID = core.StringPtr("testString")
				updateToolchainOptionsModel.ToolchainPrototypePatch = toolchainPrototypePatchModelAsPatch
				updateToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdToolchainService.UpdateToolchain(updateToolchainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateToolchain with error: Operation validation and request error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the ToolchainPrototypePatch model
				toolchainPrototypePatchModel := new(cdtoolchainv2.ToolchainPrototypePatch)
				toolchainPrototypePatchModel.Name = core.StringPtr("newToolchainName")
				toolchainPrototypePatchModel.Description = core.StringPtr("New toolchain description")
				toolchainPrototypePatchModelAsPatch, asPatchErr := toolchainPrototypePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateToolchainOptions model
				updateToolchainOptionsModel := new(cdtoolchainv2.UpdateToolchainOptions)
				updateToolchainOptionsModel.ToolchainID = core.StringPtr("testString")
				updateToolchainOptionsModel.ToolchainPrototypePatch = toolchainPrototypePatchModelAsPatch
				updateToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdToolchainService.UpdateToolchain(updateToolchainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateToolchainOptions model with no property values
				updateToolchainOptionsModelNew := new(cdtoolchainv2.UpdateToolchainOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdToolchainService.UpdateToolchain(updateToolchainOptionsModelNew)
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
			It(`Invoke UpdateToolchain successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the ToolchainPrototypePatch model
				toolchainPrototypePatchModel := new(cdtoolchainv2.ToolchainPrototypePatch)
				toolchainPrototypePatchModel.Name = core.StringPtr("newToolchainName")
				toolchainPrototypePatchModel.Description = core.StringPtr("New toolchain description")
				toolchainPrototypePatchModelAsPatch, asPatchErr := toolchainPrototypePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateToolchainOptions model
				updateToolchainOptionsModel := new(cdtoolchainv2.UpdateToolchainOptions)
				updateToolchainOptionsModel.ToolchainID = core.StringPtr("testString")
				updateToolchainOptionsModel.ToolchainPrototypePatch = toolchainPrototypePatchModelAsPatch
				updateToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdToolchainService.UpdateToolchain(updateToolchainOptionsModel)
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
	Describe(`ListTools(listToolsOptions *ListToolsOptions) - Operation response error`, func() {
		listToolsPath := "/toolchains/testString/tools"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listToolsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTools with error: Operation response processing error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the ListToolsOptions model
				listToolsOptionsModel := new(cdtoolchainv2.ListToolsOptions)
				listToolsOptionsModel.ToolchainID = core.StringPtr("testString")
				listToolsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listToolsOptionsModel.Start = core.StringPtr("testString")
				listToolsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdToolchainService.ListTools(listToolsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdToolchainService.EnableRetries(0, 0)
				result, response, operationErr = cdToolchainService.ListTools(listToolsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTools(listToolsOptions *ListToolsOptions)`, func() {
		listToolsPath := "/toolchains/testString/tools"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listToolsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "total_count": 10, "first": {"href": "Href"}, "previous": {"start": "Start", "href": "Href"}, "next": {"start": "Start", "href": "Href"}, "last": {"start": "Start", "href": "Href"}, "tools": [{"id": "ID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "tool_type_id": "ToolTypeID", "toolchain_id": "ToolchainID", "toolchain_crn": "ToolchainCRN", "href": "Href", "referent": {"ui_href": "UIHref", "api_href": "APIHref"}, "name": "Name", "updated_at": "2019-01-01T12:00:00.000Z", "parameters": {"anyKey": "anyValue"}, "state": "configured"}]}`)
				}))
			})
			It(`Invoke ListTools successfully with retries`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())
				cdToolchainService.EnableRetries(0, 0)

				// Construct an instance of the ListToolsOptions model
				listToolsOptionsModel := new(cdtoolchainv2.ListToolsOptions)
				listToolsOptionsModel.ToolchainID = core.StringPtr("testString")
				listToolsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listToolsOptionsModel.Start = core.StringPtr("testString")
				listToolsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdToolchainService.ListToolsWithContext(ctx, listToolsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdToolchainService.DisableRetries()
				result, response, operationErr := cdToolchainService.ListTools(listToolsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdToolchainService.ListToolsWithContext(ctx, listToolsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listToolsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "total_count": 10, "first": {"href": "Href"}, "previous": {"start": "Start", "href": "Href"}, "next": {"start": "Start", "href": "Href"}, "last": {"start": "Start", "href": "Href"}, "tools": [{"id": "ID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "tool_type_id": "ToolTypeID", "toolchain_id": "ToolchainID", "toolchain_crn": "ToolchainCRN", "href": "Href", "referent": {"ui_href": "UIHref", "api_href": "APIHref"}, "name": "Name", "updated_at": "2019-01-01T12:00:00.000Z", "parameters": {"anyKey": "anyValue"}, "state": "configured"}]}`)
				}))
			})
			It(`Invoke ListTools successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdToolchainService.ListTools(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListToolsOptions model
				listToolsOptionsModel := new(cdtoolchainv2.ListToolsOptions)
				listToolsOptionsModel.ToolchainID = core.StringPtr("testString")
				listToolsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listToolsOptionsModel.Start = core.StringPtr("testString")
				listToolsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdToolchainService.ListTools(listToolsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTools with error: Operation validation and request error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the ListToolsOptions model
				listToolsOptionsModel := new(cdtoolchainv2.ListToolsOptions)
				listToolsOptionsModel.ToolchainID = core.StringPtr("testString")
				listToolsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listToolsOptionsModel.Start = core.StringPtr("testString")
				listToolsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdToolchainService.ListTools(listToolsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListToolsOptions model with no property values
				listToolsOptionsModelNew := new(cdtoolchainv2.ListToolsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdToolchainService.ListTools(listToolsOptionsModelNew)
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
			It(`Invoke ListTools successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the ListToolsOptions model
				listToolsOptionsModel := new(cdtoolchainv2.ListToolsOptions)
				listToolsOptionsModel.ToolchainID = core.StringPtr("testString")
				listToolsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listToolsOptionsModel.Start = core.StringPtr("testString")
				listToolsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdToolchainService.ListTools(listToolsOptionsModel)
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
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(cdtoolchainv2.ToolchainToolCollection)
				nextObject := new(cdtoolchainv2.ToolchainToolCollectionNext)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(cdtoolchainv2.ToolchainToolCollection)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listToolsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"tools":[{"id":"ID","resource_group_id":"ResourceGroupID","crn":"CRN","tool_type_id":"ToolTypeID","toolchain_id":"ToolchainID","toolchain_crn":"ToolchainCRN","href":"Href","referent":{"ui_href":"UIHref","api_href":"APIHref"},"name":"Name","updated_at":"2019-01-01T12:00:00.000Z","parameters":{"anyKey":"anyValue"},"state":"configured"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"tools":[{"id":"ID","resource_group_id":"ResourceGroupID","crn":"CRN","tool_type_id":"ToolTypeID","toolchain_id":"ToolchainID","toolchain_crn":"ToolchainCRN","href":"Href","referent":{"ui_href":"UIHref","api_href":"APIHref"},"name":"Name","updated_at":"2019-01-01T12:00:00.000Z","parameters":{"anyKey":"anyValue"},"state":"configured"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ToolsPager.GetNext successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				listToolsOptionsModel := &cdtoolchainv2.ListToolsOptions{
					ToolchainID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := cdToolchainService.NewToolsPager(listToolsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []cdtoolchainv2.ToolModel
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ToolsPager.GetAll successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				listToolsOptionsModel := &cdtoolchainv2.ListToolsOptions{
					ToolchainID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := cdToolchainService.NewToolsPager(listToolsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateTool(createToolOptions *CreateToolOptions) - Operation response error`, func() {
		createToolPath := "/toolchains/testString/tools"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createToolPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTool with error: Operation response processing error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the CreateToolOptions model
				createToolOptionsModel := new(cdtoolchainv2.CreateToolOptions)
				createToolOptionsModel.ToolchainID = core.StringPtr("testString")
				createToolOptionsModel.ToolTypeID = core.StringPtr("draservicebroker")
				createToolOptionsModel.Name = core.StringPtr("testString")
				createToolOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				createToolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdToolchainService.CreateTool(createToolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdToolchainService.EnableRetries(0, 0)
				result, response, operationErr = cdToolchainService.CreateTool(createToolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTool(createToolOptions *CreateToolOptions)`, func() {
		createToolPath := "/toolchains/testString/tools"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createToolPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "tool_type_id": "ToolTypeID", "toolchain_id": "ToolchainID", "toolchain_crn": "ToolchainCRN", "href": "Href", "referent": {"ui_href": "UIHref", "api_href": "APIHref"}, "name": "Name", "updated_at": "2019-01-01T12:00:00.000Z", "parameters": {"anyKey": "anyValue"}, "state": "configured"}`)
				}))
			})
			It(`Invoke CreateTool successfully with retries`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())
				cdToolchainService.EnableRetries(0, 0)

				// Construct an instance of the CreateToolOptions model
				createToolOptionsModel := new(cdtoolchainv2.CreateToolOptions)
				createToolOptionsModel.ToolchainID = core.StringPtr("testString")
				createToolOptionsModel.ToolTypeID = core.StringPtr("draservicebroker")
				createToolOptionsModel.Name = core.StringPtr("testString")
				createToolOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				createToolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdToolchainService.CreateToolWithContext(ctx, createToolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdToolchainService.DisableRetries()
				result, response, operationErr := cdToolchainService.CreateTool(createToolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdToolchainService.CreateToolWithContext(ctx, createToolOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createToolPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "tool_type_id": "ToolTypeID", "toolchain_id": "ToolchainID", "toolchain_crn": "ToolchainCRN", "href": "Href", "referent": {"ui_href": "UIHref", "api_href": "APIHref"}, "name": "Name", "updated_at": "2019-01-01T12:00:00.000Z", "parameters": {"anyKey": "anyValue"}, "state": "configured"}`)
				}))
			})
			It(`Invoke CreateTool successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdToolchainService.CreateTool(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateToolOptions model
				createToolOptionsModel := new(cdtoolchainv2.CreateToolOptions)
				createToolOptionsModel.ToolchainID = core.StringPtr("testString")
				createToolOptionsModel.ToolTypeID = core.StringPtr("draservicebroker")
				createToolOptionsModel.Name = core.StringPtr("testString")
				createToolOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				createToolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdToolchainService.CreateTool(createToolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTool with error: Operation validation and request error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the CreateToolOptions model
				createToolOptionsModel := new(cdtoolchainv2.CreateToolOptions)
				createToolOptionsModel.ToolchainID = core.StringPtr("testString")
				createToolOptionsModel.ToolTypeID = core.StringPtr("draservicebroker")
				createToolOptionsModel.Name = core.StringPtr("testString")
				createToolOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				createToolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdToolchainService.CreateTool(createToolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateToolOptions model with no property values
				createToolOptionsModelNew := new(cdtoolchainv2.CreateToolOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdToolchainService.CreateTool(createToolOptionsModelNew)
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
			It(`Invoke CreateTool successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the CreateToolOptions model
				createToolOptionsModel := new(cdtoolchainv2.CreateToolOptions)
				createToolOptionsModel.ToolchainID = core.StringPtr("testString")
				createToolOptionsModel.ToolTypeID = core.StringPtr("draservicebroker")
				createToolOptionsModel.Name = core.StringPtr("testString")
				createToolOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				createToolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdToolchainService.CreateTool(createToolOptionsModel)
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
	Describe(`GetToolByID(getToolByIDOptions *GetToolByIDOptions) - Operation response error`, func() {
		getToolByIDPath := "/toolchains/testString/tools/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getToolByIDPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetToolByID with error: Operation response processing error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the GetToolByIDOptions model
				getToolByIDOptionsModel := new(cdtoolchainv2.GetToolByIDOptions)
				getToolByIDOptionsModel.ToolchainID = core.StringPtr("testString")
				getToolByIDOptionsModel.ToolID = core.StringPtr("testString")
				getToolByIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdToolchainService.GetToolByID(getToolByIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdToolchainService.EnableRetries(0, 0)
				result, response, operationErr = cdToolchainService.GetToolByID(getToolByIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetToolByID(getToolByIDOptions *GetToolByIDOptions)`, func() {
		getToolByIDPath := "/toolchains/testString/tools/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getToolByIDPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "tool_type_id": "ToolTypeID", "toolchain_id": "ToolchainID", "toolchain_crn": "ToolchainCRN", "href": "Href", "referent": {"ui_href": "UIHref", "api_href": "APIHref"}, "name": "Name", "updated_at": "2019-01-01T12:00:00.000Z", "parameters": {"anyKey": "anyValue"}, "state": "configured"}`)
				}))
			})
			It(`Invoke GetToolByID successfully with retries`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())
				cdToolchainService.EnableRetries(0, 0)

				// Construct an instance of the GetToolByIDOptions model
				getToolByIDOptionsModel := new(cdtoolchainv2.GetToolByIDOptions)
				getToolByIDOptionsModel.ToolchainID = core.StringPtr("testString")
				getToolByIDOptionsModel.ToolID = core.StringPtr("testString")
				getToolByIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdToolchainService.GetToolByIDWithContext(ctx, getToolByIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdToolchainService.DisableRetries()
				result, response, operationErr := cdToolchainService.GetToolByID(getToolByIDOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdToolchainService.GetToolByIDWithContext(ctx, getToolByIDOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getToolByIDPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "tool_type_id": "ToolTypeID", "toolchain_id": "ToolchainID", "toolchain_crn": "ToolchainCRN", "href": "Href", "referent": {"ui_href": "UIHref", "api_href": "APIHref"}, "name": "Name", "updated_at": "2019-01-01T12:00:00.000Z", "parameters": {"anyKey": "anyValue"}, "state": "configured"}`)
				}))
			})
			It(`Invoke GetToolByID successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdToolchainService.GetToolByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetToolByIDOptions model
				getToolByIDOptionsModel := new(cdtoolchainv2.GetToolByIDOptions)
				getToolByIDOptionsModel.ToolchainID = core.StringPtr("testString")
				getToolByIDOptionsModel.ToolID = core.StringPtr("testString")
				getToolByIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdToolchainService.GetToolByID(getToolByIDOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetToolByID with error: Operation validation and request error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the GetToolByIDOptions model
				getToolByIDOptionsModel := new(cdtoolchainv2.GetToolByIDOptions)
				getToolByIDOptionsModel.ToolchainID = core.StringPtr("testString")
				getToolByIDOptionsModel.ToolID = core.StringPtr("testString")
				getToolByIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdToolchainService.GetToolByID(getToolByIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetToolByIDOptions model with no property values
				getToolByIDOptionsModelNew := new(cdtoolchainv2.GetToolByIDOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdToolchainService.GetToolByID(getToolByIDOptionsModelNew)
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
			It(`Invoke GetToolByID successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the GetToolByIDOptions model
				getToolByIDOptionsModel := new(cdtoolchainv2.GetToolByIDOptions)
				getToolByIDOptionsModel.ToolchainID = core.StringPtr("testString")
				getToolByIDOptionsModel.ToolID = core.StringPtr("testString")
				getToolByIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdToolchainService.GetToolByID(getToolByIDOptionsModel)
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
	Describe(`DeleteTool(deleteToolOptions *DeleteToolOptions)`, func() {
		deleteToolPath := "/toolchains/testString/tools/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteToolPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTool successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := cdToolchainService.DeleteTool(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteToolOptions model
				deleteToolOptionsModel := new(cdtoolchainv2.DeleteToolOptions)
				deleteToolOptionsModel.ToolchainID = core.StringPtr("testString")
				deleteToolOptionsModel.ToolID = core.StringPtr("testString")
				deleteToolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = cdToolchainService.DeleteTool(deleteToolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTool with error: Operation validation and request error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the DeleteToolOptions model
				deleteToolOptionsModel := new(cdtoolchainv2.DeleteToolOptions)
				deleteToolOptionsModel.ToolchainID = core.StringPtr("testString")
				deleteToolOptionsModel.ToolID = core.StringPtr("testString")
				deleteToolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := cdToolchainService.DeleteTool(deleteToolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteToolOptions model with no property values
				deleteToolOptionsModelNew := new(cdtoolchainv2.DeleteToolOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = cdToolchainService.DeleteTool(deleteToolOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTool(updateToolOptions *UpdateToolOptions) - Operation response error`, func() {
		updateToolPath := "/toolchains/testString/tools/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateToolPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTool with error: Operation response processing error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the ToolchainToolPrototypePatch model
				toolchainToolPrototypePatchModel := new(cdtoolchainv2.ToolchainToolPrototypePatch)
				toolchainToolPrototypePatchModel.Name = core.StringPtr("MyTool")
				toolchainToolPrototypePatchModel.ToolTypeID = core.StringPtr("draservicebroker")
				toolchainToolPrototypePatchModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				toolchainToolPrototypePatchModelAsPatch, asPatchErr := toolchainToolPrototypePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateToolOptions model
				updateToolOptionsModel := new(cdtoolchainv2.UpdateToolOptions)
				updateToolOptionsModel.ToolchainID = core.StringPtr("testString")
				updateToolOptionsModel.ToolID = core.StringPtr("testString")
				updateToolOptionsModel.ToolchainToolPrototypePatch = toolchainToolPrototypePatchModelAsPatch
				updateToolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdToolchainService.UpdateTool(updateToolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdToolchainService.EnableRetries(0, 0)
				result, response, operationErr = cdToolchainService.UpdateTool(updateToolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTool(updateToolOptions *UpdateToolOptions)`, func() {
		updateToolPath := "/toolchains/testString/tools/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateToolPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "tool_type_id": "ToolTypeID", "toolchain_id": "ToolchainID", "toolchain_crn": "ToolchainCRN", "href": "Href", "referent": {"ui_href": "UIHref", "api_href": "APIHref"}, "name": "Name", "updated_at": "2019-01-01T12:00:00.000Z", "parameters": {"anyKey": "anyValue"}, "state": "configured"}`)
				}))
			})
			It(`Invoke UpdateTool successfully with retries`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())
				cdToolchainService.EnableRetries(0, 0)

				// Construct an instance of the ToolchainToolPrototypePatch model
				toolchainToolPrototypePatchModel := new(cdtoolchainv2.ToolchainToolPrototypePatch)
				toolchainToolPrototypePatchModel.Name = core.StringPtr("MyTool")
				toolchainToolPrototypePatchModel.ToolTypeID = core.StringPtr("draservicebroker")
				toolchainToolPrototypePatchModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				toolchainToolPrototypePatchModelAsPatch, asPatchErr := toolchainToolPrototypePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateToolOptions model
				updateToolOptionsModel := new(cdtoolchainv2.UpdateToolOptions)
				updateToolOptionsModel.ToolchainID = core.StringPtr("testString")
				updateToolOptionsModel.ToolID = core.StringPtr("testString")
				updateToolOptionsModel.ToolchainToolPrototypePatch = toolchainToolPrototypePatchModelAsPatch
				updateToolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdToolchainService.UpdateToolWithContext(ctx, updateToolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdToolchainService.DisableRetries()
				result, response, operationErr := cdToolchainService.UpdateTool(updateToolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdToolchainService.UpdateToolWithContext(ctx, updateToolOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateToolPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "tool_type_id": "ToolTypeID", "toolchain_id": "ToolchainID", "toolchain_crn": "ToolchainCRN", "href": "Href", "referent": {"ui_href": "UIHref", "api_href": "APIHref"}, "name": "Name", "updated_at": "2019-01-01T12:00:00.000Z", "parameters": {"anyKey": "anyValue"}, "state": "configured"}`)
				}))
			})
			It(`Invoke UpdateTool successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdToolchainService.UpdateTool(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ToolchainToolPrototypePatch model
				toolchainToolPrototypePatchModel := new(cdtoolchainv2.ToolchainToolPrototypePatch)
				toolchainToolPrototypePatchModel.Name = core.StringPtr("MyTool")
				toolchainToolPrototypePatchModel.ToolTypeID = core.StringPtr("draservicebroker")
				toolchainToolPrototypePatchModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				toolchainToolPrototypePatchModelAsPatch, asPatchErr := toolchainToolPrototypePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateToolOptions model
				updateToolOptionsModel := new(cdtoolchainv2.UpdateToolOptions)
				updateToolOptionsModel.ToolchainID = core.StringPtr("testString")
				updateToolOptionsModel.ToolID = core.StringPtr("testString")
				updateToolOptionsModel.ToolchainToolPrototypePatch = toolchainToolPrototypePatchModelAsPatch
				updateToolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdToolchainService.UpdateTool(updateToolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTool with error: Operation validation and request error`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the ToolchainToolPrototypePatch model
				toolchainToolPrototypePatchModel := new(cdtoolchainv2.ToolchainToolPrototypePatch)
				toolchainToolPrototypePatchModel.Name = core.StringPtr("MyTool")
				toolchainToolPrototypePatchModel.ToolTypeID = core.StringPtr("draservicebroker")
				toolchainToolPrototypePatchModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				toolchainToolPrototypePatchModelAsPatch, asPatchErr := toolchainToolPrototypePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateToolOptions model
				updateToolOptionsModel := new(cdtoolchainv2.UpdateToolOptions)
				updateToolOptionsModel.ToolchainID = core.StringPtr("testString")
				updateToolOptionsModel.ToolID = core.StringPtr("testString")
				updateToolOptionsModel.ToolchainToolPrototypePatch = toolchainToolPrototypePatchModelAsPatch
				updateToolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdToolchainService.UpdateTool(updateToolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateToolOptions model with no property values
				updateToolOptionsModelNew := new(cdtoolchainv2.UpdateToolOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdToolchainService.UpdateTool(updateToolOptionsModelNew)
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
			It(`Invoke UpdateTool successfully`, func() {
				cdToolchainService, serviceErr := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdToolchainService).ToNot(BeNil())

				// Construct an instance of the ToolchainToolPrototypePatch model
				toolchainToolPrototypePatchModel := new(cdtoolchainv2.ToolchainToolPrototypePatch)
				toolchainToolPrototypePatchModel.Name = core.StringPtr("MyTool")
				toolchainToolPrototypePatchModel.ToolTypeID = core.StringPtr("draservicebroker")
				toolchainToolPrototypePatchModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				toolchainToolPrototypePatchModelAsPatch, asPatchErr := toolchainToolPrototypePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateToolOptions model
				updateToolOptionsModel := new(cdtoolchainv2.UpdateToolOptions)
				updateToolOptionsModel.ToolchainID = core.StringPtr("testString")
				updateToolOptionsModel.ToolID = core.StringPtr("testString")
				updateToolOptionsModel.ToolchainToolPrototypePatch = toolchainToolPrototypePatchModelAsPatch
				updateToolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdToolchainService.UpdateTool(updateToolOptionsModel)
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
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			cdToolchainService, _ := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
				URL:           "http://cdtoolchainv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateToolOptions successfully`, func() {
				// Construct an instance of the CreateToolOptions model
				toolchainID := "testString"
				createToolOptionsToolTypeID := "draservicebroker"
				createToolOptionsModel := cdToolchainService.NewCreateToolOptions(toolchainID, createToolOptionsToolTypeID)
				createToolOptionsModel.SetToolchainID("testString")
				createToolOptionsModel.SetToolTypeID("draservicebroker")
				createToolOptionsModel.SetName("testString")
				createToolOptionsModel.SetParameters(map[string]interface{}{"anyKey": "anyValue"})
				createToolOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createToolOptionsModel).ToNot(BeNil())
				Expect(createToolOptionsModel.ToolchainID).To(Equal(core.StringPtr("testString")))
				Expect(createToolOptionsModel.ToolTypeID).To(Equal(core.StringPtr("draservicebroker")))
				Expect(createToolOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createToolOptionsModel.Parameters).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(createToolOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewDeleteToolOptions successfully`, func() {
				// Construct an instance of the DeleteToolOptions model
				toolchainID := "testString"
				toolID := "testString"
				deleteToolOptionsModel := cdToolchainService.NewDeleteToolOptions(toolchainID, toolID)
				deleteToolOptionsModel.SetToolchainID("testString")
				deleteToolOptionsModel.SetToolID("testString")
				deleteToolOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteToolOptionsModel).ToNot(BeNil())
				Expect(deleteToolOptionsModel.ToolchainID).To(Equal(core.StringPtr("testString")))
				Expect(deleteToolOptionsModel.ToolID).To(Equal(core.StringPtr("testString")))
				Expect(deleteToolOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewGetToolByIDOptions successfully`, func() {
				// Construct an instance of the GetToolByIDOptions model
				toolchainID := "testString"
				toolID := "testString"
				getToolByIDOptionsModel := cdToolchainService.NewGetToolByIDOptions(toolchainID, toolID)
				getToolByIDOptionsModel.SetToolchainID("testString")
				getToolByIDOptionsModel.SetToolID("testString")
				getToolByIDOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getToolByIDOptionsModel).ToNot(BeNil())
				Expect(getToolByIDOptionsModel.ToolchainID).To(Equal(core.StringPtr("testString")))
				Expect(getToolByIDOptionsModel.ToolID).To(Equal(core.StringPtr("testString")))
				Expect(getToolByIDOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewListToolchainsOptions successfully`, func() {
				// Construct an instance of the ListToolchainsOptions model
				resourceGroupID := "testString"
				listToolchainsOptionsModel := cdToolchainService.NewListToolchainsOptions(resourceGroupID)
				listToolchainsOptionsModel.SetResourceGroupID("testString")
				listToolchainsOptionsModel.SetLimit(int64(10))
				listToolchainsOptionsModel.SetStart("testString")
				listToolchainsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listToolchainsOptionsModel).ToNot(BeNil())
				Expect(listToolchainsOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listToolchainsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listToolchainsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listToolchainsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListToolsOptions successfully`, func() {
				// Construct an instance of the ListToolsOptions model
				toolchainID := "testString"
				listToolsOptionsModel := cdToolchainService.NewListToolsOptions(toolchainID)
				listToolsOptionsModel.SetToolchainID("testString")
				listToolsOptionsModel.SetLimit(int64(10))
				listToolsOptionsModel.SetStart("testString")
				listToolsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listToolsOptionsModel).ToNot(BeNil())
				Expect(listToolsOptionsModel.ToolchainID).To(Equal(core.StringPtr("testString")))
				Expect(listToolsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listToolsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listToolsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateToolOptions successfully`, func() {
				// Construct an instance of the UpdateToolOptions model
				toolchainID := "testString"
				toolID := "testString"
				toolchainToolPrototypePatch := map[string]interface{}{"anyKey": "anyValue"}
				updateToolOptionsModel := cdToolchainService.NewUpdateToolOptions(toolchainID, toolID, toolchainToolPrototypePatch)
				updateToolOptionsModel.SetToolchainID("testString")
				updateToolOptionsModel.SetToolID("testString")
				updateToolOptionsModel.SetToolchainToolPrototypePatch(map[string]interface{}{"anyKey": "anyValue"})
				updateToolOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateToolOptionsModel).ToNot(BeNil())
				Expect(updateToolOptionsModel.ToolchainID).To(Equal(core.StringPtr("testString")))
				Expect(updateToolOptionsModel.ToolID).To(Equal(core.StringPtr("testString")))
				Expect(updateToolOptionsModel.ToolchainToolPrototypePatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateToolOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateToolchainOptions successfully`, func() {
				// Construct an instance of the UpdateToolchainOptions model
				toolchainID := "testString"
				toolchainPrototypePatch := map[string]interface{}{"anyKey": "anyValue"}
				updateToolchainOptionsModel := cdToolchainService.NewUpdateToolchainOptions(toolchainID, toolchainPrototypePatch)
				updateToolchainOptionsModel.SetToolchainID("testString")
				updateToolchainOptionsModel.SetToolchainPrototypePatch(map[string]interface{}{"anyKey": "anyValue"})
				updateToolchainOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateToolchainOptionsModel).ToNot(BeNil())
				Expect(updateToolchainOptionsModel.ToolchainID).To(Equal(core.StringPtr("testString")))
				Expect(updateToolchainOptionsModel.ToolchainPrototypePatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
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
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
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
