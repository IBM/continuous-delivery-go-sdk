/**
 * (C) Copyright IBM Corp. 2021.
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

package exampleservicev1_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/CloudEngineering/go-sdk-template/exampleservicev1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`ExampleServiceV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(exampleServiceService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(exampleServiceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
				URL: "https://exampleservicev1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(exampleServiceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"EXAMPLE_SERVICE_URL": "https://exampleservicev1/api",
				"EXAMPLE_SERVICE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1UsingExternalConfig(&exampleservicev1.ExampleServiceV1Options{
				})
				Expect(exampleServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := exampleServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != exampleServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(exampleServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(exampleServiceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1UsingExternalConfig(&exampleservicev1.ExampleServiceV1Options{
					URL: "https://testService/api",
				})
				Expect(exampleServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(exampleServiceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := exampleServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != exampleServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(exampleServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(exampleServiceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1UsingExternalConfig(&exampleservicev1.ExampleServiceV1Options{
				})
				err := exampleServiceService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(exampleServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(exampleServiceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := exampleServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != exampleServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(exampleServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(exampleServiceService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"EXAMPLE_SERVICE_URL": "https://exampleservicev1/api",
				"EXAMPLE_SERVICE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1UsingExternalConfig(&exampleservicev1.ExampleServiceV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(exampleServiceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"EXAMPLE_SERVICE_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1UsingExternalConfig(&exampleservicev1.ExampleServiceV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(exampleServiceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = exampleservicev1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListResources(listResourcesOptions *ListResourcesOptions) - Operation response error`, func() {
		listResourcesPath := "/resources"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourcesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListResources with error: Operation response processing error`, func() {
				exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(exampleServiceService).ToNot(BeNil())

				// Construct an instance of the ListResourcesOptions model
				listResourcesOptionsModel := new(exampleservicev1.ListResourcesOptions)
				listResourcesOptionsModel.Limit = core.Int64Ptr(int64(38))
				listResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := exampleServiceService.ListResources(listResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				exampleServiceService.EnableRetries(0, 0)
				result, response, operationErr = exampleServiceService.ListResources(listResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListResources(listResourcesOptions *ListResourcesOptions)`, func() {
		listResourcesPath := "/resources"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourcesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "resources": [{"resource_id": "ResourceID", "name": "Name", "tag": "Tag", "read_only": "ReadOnly"}]}`)
				}))
			})
			It(`Invoke ListResources successfully with retries`, func() {
				exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(exampleServiceService).ToNot(BeNil())
				exampleServiceService.EnableRetries(0, 0)

				// Construct an instance of the ListResourcesOptions model
				listResourcesOptionsModel := new(exampleservicev1.ListResourcesOptions)
				listResourcesOptionsModel.Limit = core.Int64Ptr(int64(38))
				listResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := exampleServiceService.ListResourcesWithContext(ctx, listResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				exampleServiceService.DisableRetries()
				result, response, operationErr := exampleServiceService.ListResources(listResourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = exampleServiceService.ListResourcesWithContext(ctx, listResourcesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listResourcesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "resources": [{"resource_id": "ResourceID", "name": "Name", "tag": "Tag", "read_only": "ReadOnly"}]}`)
				}))
			})
			It(`Invoke ListResources successfully`, func() {
				exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(exampleServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := exampleServiceService.ListResources(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourcesOptions model
				listResourcesOptionsModel := new(exampleservicev1.ListResourcesOptions)
				listResourcesOptionsModel.Limit = core.Int64Ptr(int64(38))
				listResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = exampleServiceService.ListResources(listResourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListResources with error: Operation request error`, func() {
				exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(exampleServiceService).ToNot(BeNil())

				// Construct an instance of the ListResourcesOptions model
				listResourcesOptionsModel := new(exampleservicev1.ListResourcesOptions)
				listResourcesOptionsModel.Limit = core.Int64Ptr(int64(38))
				listResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := exampleServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := exampleServiceService.ListResources(listResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateResource(createResourceOptions *CreateResourceOptions) - Operation response error`, func() {
		createResourcePath := "/resources"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createResourcePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateResource with error: Operation response processing error`, func() {
				exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(exampleServiceService).ToNot(BeNil())

				// Construct an instance of the CreateResourceOptions model
				createResourceOptionsModel := new(exampleservicev1.CreateResourceOptions)
				createResourceOptionsModel.ResourceID = core.StringPtr("testString")
				createResourceOptionsModel.Name = core.StringPtr("testString")
				createResourceOptionsModel.Tag = core.StringPtr("testString")
				createResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := exampleServiceService.CreateResource(createResourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				exampleServiceService.EnableRetries(0, 0)
				result, response, operationErr = exampleServiceService.CreateResource(createResourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateResource(createResourceOptions *CreateResourceOptions)`, func() {
		createResourcePath := "/resources"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createResourcePath))
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
					fmt.Fprintf(res, "%s", `{"resource_id": "ResourceID", "name": "Name", "tag": "Tag", "read_only": "ReadOnly"}`)
				}))
			})
			It(`Invoke CreateResource successfully with retries`, func() {
				exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(exampleServiceService).ToNot(BeNil())
				exampleServiceService.EnableRetries(0, 0)

				// Construct an instance of the CreateResourceOptions model
				createResourceOptionsModel := new(exampleservicev1.CreateResourceOptions)
				createResourceOptionsModel.ResourceID = core.StringPtr("testString")
				createResourceOptionsModel.Name = core.StringPtr("testString")
				createResourceOptionsModel.Tag = core.StringPtr("testString")
				createResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := exampleServiceService.CreateResourceWithContext(ctx, createResourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				exampleServiceService.DisableRetries()
				result, response, operationErr := exampleServiceService.CreateResource(createResourceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = exampleServiceService.CreateResourceWithContext(ctx, createResourceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createResourcePath))
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
					fmt.Fprintf(res, "%s", `{"resource_id": "ResourceID", "name": "Name", "tag": "Tag", "read_only": "ReadOnly"}`)
				}))
			})
			It(`Invoke CreateResource successfully`, func() {
				exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(exampleServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := exampleServiceService.CreateResource(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateResourceOptions model
				createResourceOptionsModel := new(exampleservicev1.CreateResourceOptions)
				createResourceOptionsModel.ResourceID = core.StringPtr("testString")
				createResourceOptionsModel.Name = core.StringPtr("testString")
				createResourceOptionsModel.Tag = core.StringPtr("testString")
				createResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = exampleServiceService.CreateResource(createResourceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateResource with error: Operation validation and request error`, func() {
				exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(exampleServiceService).ToNot(BeNil())

				// Construct an instance of the CreateResourceOptions model
				createResourceOptionsModel := new(exampleservicev1.CreateResourceOptions)
				createResourceOptionsModel.ResourceID = core.StringPtr("testString")
				createResourceOptionsModel.Name = core.StringPtr("testString")
				createResourceOptionsModel.Tag = core.StringPtr("testString")
				createResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := exampleServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := exampleServiceService.CreateResource(createResourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateResourceOptions model with no property values
				createResourceOptionsModelNew := new(exampleservicev1.CreateResourceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = exampleServiceService.CreateResource(createResourceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResource(getResourceOptions *GetResourceOptions) - Operation response error`, func() {
		getResourcePath := "/resources/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourcePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResource with error: Operation response processing error`, func() {
				exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(exampleServiceService).ToNot(BeNil())

				// Construct an instance of the GetResourceOptions model
				getResourceOptionsModel := new(exampleservicev1.GetResourceOptions)
				getResourceOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := exampleServiceService.GetResource(getResourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				exampleServiceService.EnableRetries(0, 0)
				result, response, operationErr = exampleServiceService.GetResource(getResourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetResource(getResourceOptions *GetResourceOptions)`, func() {
		getResourcePath := "/resources/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourcePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resource_id": "ResourceID", "name": "Name", "tag": "Tag", "read_only": "ReadOnly"}`)
				}))
			})
			It(`Invoke GetResource successfully with retries`, func() {
				exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(exampleServiceService).ToNot(BeNil())
				exampleServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetResourceOptions model
				getResourceOptionsModel := new(exampleservicev1.GetResourceOptions)
				getResourceOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := exampleServiceService.GetResourceWithContext(ctx, getResourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				exampleServiceService.DisableRetries()
				result, response, operationErr := exampleServiceService.GetResource(getResourceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = exampleServiceService.GetResourceWithContext(ctx, getResourceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getResourcePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resource_id": "ResourceID", "name": "Name", "tag": "Tag", "read_only": "ReadOnly"}`)
				}))
			})
			It(`Invoke GetResource successfully`, func() {
				exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(exampleServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := exampleServiceService.GetResource(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceOptions model
				getResourceOptionsModel := new(exampleservicev1.GetResourceOptions)
				getResourceOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = exampleServiceService.GetResource(getResourceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetResource with error: Operation validation and request error`, func() {
				exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(exampleServiceService).ToNot(BeNil())

				// Construct an instance of the GetResourceOptions model
				getResourceOptionsModel := new(exampleservicev1.GetResourceOptions)
				getResourceOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := exampleServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := exampleServiceService.GetResource(getResourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceOptions model with no property values
				getResourceOptionsModelNew := new(exampleservicev1.GetResourceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = exampleServiceService.GetResource(getResourceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceEncoded(getResourceEncodedOptions *GetResourceEncodedOptions) - Operation response error`, func() {
		getResourceEncodedPath := "/resources/encoded/url%253encoded%253resource%253id"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceEncodedPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceEncoded with error: Operation response processing error`, func() {
				exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(exampleServiceService).ToNot(BeNil())

				// Construct an instance of the GetResourceEncodedOptions model
				getResourceEncodedOptionsModel := new(exampleservicev1.GetResourceEncodedOptions)
				getResourceEncodedOptionsModel.UrlEncodedResourceID = core.StringPtr("url%3encoded%3resource%3id")
				getResourceEncodedOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := exampleServiceService.GetResourceEncoded(getResourceEncodedOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				exampleServiceService.EnableRetries(0, 0)
				result, response, operationErr = exampleServiceService.GetResourceEncoded(getResourceEncodedOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetResourceEncoded(getResourceEncodedOptions *GetResourceEncodedOptions)`, func() {
		getResourceEncodedPath := "/resources/encoded/url%253encoded%253resource%253id"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceEncodedPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resource_id": "ResourceID", "name": "Name", "tag": "Tag", "read_only": "ReadOnly"}`)
				}))
			})
			It(`Invoke GetResourceEncoded successfully with retries`, func() {
				exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(exampleServiceService).ToNot(BeNil())
				exampleServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetResourceEncodedOptions model
				getResourceEncodedOptionsModel := new(exampleservicev1.GetResourceEncodedOptions)
				getResourceEncodedOptionsModel.UrlEncodedResourceID = core.StringPtr("url%3encoded%3resource%3id")
				getResourceEncodedOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := exampleServiceService.GetResourceEncodedWithContext(ctx, getResourceEncodedOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				exampleServiceService.DisableRetries()
				result, response, operationErr := exampleServiceService.GetResourceEncoded(getResourceEncodedOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = exampleServiceService.GetResourceEncodedWithContext(ctx, getResourceEncodedOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getResourceEncodedPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resource_id": "ResourceID", "name": "Name", "tag": "Tag", "read_only": "ReadOnly"}`)
				}))
			})
			It(`Invoke GetResourceEncoded successfully`, func() {
				exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(exampleServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := exampleServiceService.GetResourceEncoded(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceEncodedOptions model
				getResourceEncodedOptionsModel := new(exampleservicev1.GetResourceEncodedOptions)
				getResourceEncodedOptionsModel.UrlEncodedResourceID = core.StringPtr("url%3encoded%3resource%3id")
				getResourceEncodedOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = exampleServiceService.GetResourceEncoded(getResourceEncodedOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetResourceEncoded with error: Operation validation and request error`, func() {
				exampleServiceService, serviceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(exampleServiceService).ToNot(BeNil())

				// Construct an instance of the GetResourceEncodedOptions model
				getResourceEncodedOptionsModel := new(exampleservicev1.GetResourceEncodedOptions)
				getResourceEncodedOptionsModel.UrlEncodedResourceID = core.StringPtr("url%3encoded%3resource%3id")
				getResourceEncodedOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := exampleServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := exampleServiceService.GetResourceEncoded(getResourceEncodedOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceEncodedOptions model with no property values
				getResourceEncodedOptionsModelNew := new(exampleservicev1.GetResourceEncodedOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = exampleServiceService.GetResourceEncoded(getResourceEncodedOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			exampleServiceService, _ := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
				URL:           "http://exampleservicev1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateResourceOptions successfully`, func() {
				// Construct an instance of the CreateResourceOptions model
				createResourceOptionsResourceID := "testString"
				createResourceOptionsName := "testString"
				createResourceOptionsModel := exampleServiceService.NewCreateResourceOptions(createResourceOptionsResourceID, createResourceOptionsName)
				createResourceOptionsModel.SetResourceID("testString")
				createResourceOptionsModel.SetName("testString")
				createResourceOptionsModel.SetTag("testString")
				createResourceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createResourceOptionsModel).ToNot(BeNil())
				Expect(createResourceOptionsModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(createResourceOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createResourceOptionsModel.Tag).To(Equal(core.StringPtr("testString")))
				Expect(createResourceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceEncodedOptions successfully`, func() {
				// Construct an instance of the GetResourceEncodedOptions model
				urlEncodedResourceID := "url%3encoded%3resource%3id"
				getResourceEncodedOptionsModel := exampleServiceService.NewGetResourceEncodedOptions(urlEncodedResourceID)
				getResourceEncodedOptionsModel.SetUrlEncodedResourceID("url%3encoded%3resource%3id")
				getResourceEncodedOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceEncodedOptionsModel).ToNot(BeNil())
				Expect(getResourceEncodedOptionsModel.UrlEncodedResourceID).To(Equal(core.StringPtr("url%3encoded%3resource%3id")))
				Expect(getResourceEncodedOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceOptions successfully`, func() {
				// Construct an instance of the GetResourceOptions model
				resourceID := "testString"
				getResourceOptionsModel := exampleServiceService.NewGetResourceOptions(resourceID)
				getResourceOptionsModel.SetResourceID("testString")
				getResourceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceOptionsModel).ToNot(BeNil())
				Expect(getResourceOptionsModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListResourcesOptions successfully`, func() {
				// Construct an instance of the ListResourcesOptions model
				listResourcesOptionsModel := exampleServiceService.NewListResourcesOptions()
				listResourcesOptionsModel.SetLimit(int64(38))
				listResourcesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listResourcesOptionsModel).ToNot(BeNil())
				Expect(listResourcesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listResourcesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResource successfully`, func() {
				resourceID := "testString"
				name := "testString"
				model, err := exampleServiceService.NewResource(resourceID, name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
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

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
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
