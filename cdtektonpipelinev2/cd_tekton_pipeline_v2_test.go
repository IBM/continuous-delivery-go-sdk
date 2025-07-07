/**
 * (C) Copyright IBM Corp. 2025.
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

package cdtektonpipelinev2_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/continuous-delivery-go-sdk/v2/cdtektonpipelinev2"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`CdTektonPipelineV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cdTektonPipelineService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cdTektonPipelineService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
				URL: "https://cdtektonpipelinev2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cdTektonPipelineService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CD_TEKTON_PIPELINE_URL": "https://cdtektonpipelinev2/api",
				"CD_TEKTON_PIPELINE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2UsingExternalConfig(&cdtektonpipelinev2.CdTektonPipelineV2Options{
				})
				Expect(cdTektonPipelineService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := cdTektonPipelineService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != cdTektonPipelineService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(cdTektonPipelineService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(cdTektonPipelineService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2UsingExternalConfig(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL: "https://testService/api",
				})
				Expect(cdTektonPipelineService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := cdTektonPipelineService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != cdTektonPipelineService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(cdTektonPipelineService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(cdTektonPipelineService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2UsingExternalConfig(&cdtektonpipelinev2.CdTektonPipelineV2Options{
				})
				err := cdTektonPipelineService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := cdTektonPipelineService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != cdTektonPipelineService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(cdTektonPipelineService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(cdTektonPipelineService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CD_TEKTON_PIPELINE_URL": "https://cdtektonpipelinev2/api",
				"CD_TEKTON_PIPELINE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2UsingExternalConfig(&cdtektonpipelinev2.CdTektonPipelineV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cdTektonPipelineService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CD_TEKTON_PIPELINE_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2UsingExternalConfig(&cdtektonpipelinev2.CdTektonPipelineV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cdTektonPipelineService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = cdtektonpipelinev2.GetServiceURLForRegion("us-south")
			Expect(url).To(Equal("https://api.us-south.devops.cloud.ibm.com/pipeline/v2"))
			Expect(err).To(BeNil())

			url, err = cdtektonpipelinev2.GetServiceURLForRegion("us-east")
			Expect(url).To(Equal("https://api.us-east.devops.cloud.ibm.com/pipeline/v2"))
			Expect(err).To(BeNil())

			url, err = cdtektonpipelinev2.GetServiceURLForRegion("eu-de")
			Expect(url).To(Equal("https://api.eu-de.devops.cloud.ibm.com/pipeline/v2"))
			Expect(err).To(BeNil())

			url, err = cdtektonpipelinev2.GetServiceURLForRegion("eu-gb")
			Expect(url).To(Equal("https://api.eu-gb.devops.cloud.ibm.com/pipeline/v2"))
			Expect(err).To(BeNil())

			url, err = cdtektonpipelinev2.GetServiceURLForRegion("eu-es")
			Expect(url).To(Equal("https://api.eu-es.devops.cloud.ibm.com/pipeline/v2"))
			Expect(err).To(BeNil())

			url, err = cdtektonpipelinev2.GetServiceURLForRegion("jp-osa")
			Expect(url).To(Equal("https://api.jp-osa.devops.cloud.ibm.com/pipeline/v2"))
			Expect(err).To(BeNil())

			url, err = cdtektonpipelinev2.GetServiceURLForRegion("jp-tok")
			Expect(url).To(Equal("https://api.jp-tok.devops.cloud.ibm.com/pipeline/v2"))
			Expect(err).To(BeNil())

			url, err = cdtektonpipelinev2.GetServiceURLForRegion("au-syd")
			Expect(url).To(Equal("https://api.au-syd.devops.cloud.ibm.com/pipeline/v2"))
			Expect(err).To(BeNil())

			url, err = cdtektonpipelinev2.GetServiceURLForRegion("ca-tor")
			Expect(url).To(Equal("https://api.ca-tor.devops.cloud.ibm.com/pipeline/v2"))
			Expect(err).To(BeNil())

			url, err = cdtektonpipelinev2.GetServiceURLForRegion("br-sao")
			Expect(url).To(Equal("https://api.br-sao.devops.cloud.ibm.com/pipeline/v2"))
			Expect(err).To(BeNil())

			url, err = cdtektonpipelinev2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateTektonPipeline(createTektonPipelineOptions *CreateTektonPipelineOptions) - Operation response error`, func() {
		createTektonPipelinePath := "/tekton_pipelines"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelinePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTektonPipeline with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("public")

				// Construct an instance of the CreateTektonPipelineOptions model
				createTektonPipelineOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineOptions)
				createTektonPipelineOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineOptionsModel.NextBuildNumber = core.Int64Ptr(int64(1))
				createTektonPipelineOptionsModel.EnableNotifications = core.BoolPtr(false)
				createTektonPipelineOptionsModel.EnablePartialCloning = core.BoolPtr(false)
				createTektonPipelineOptionsModel.Worker = workerIdentityModel
				createTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipeline(createTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.CreateTektonPipeline(createTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTektonPipeline(createTektonPipelineOptions *CreateTektonPipelineOptions)`, func() {
		createTektonPipelinePath := "/tekton_pipelines"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelinePath))
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "status": "configured", "resource_group": {"id": "ID"}, "toolchain": {"id": "ID", "crn": "crn:v1:staging:public:toolchain:us-south:a/0ba224679d6c697f9baee5e14ade83ac:bf5fa00f-ddef-4298-b87b-aa8b6da0e1a6::"}, "id": "ID", "definitions": [{"source": {"type": "Type", "properties": {"url": "URL", "branch": "Branch", "tag": "Tag", "path": "Path", "tool": {"id": "ID"}}}, "href": "Href", "id": "ID"}], "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}], "updated_at": "2019-01-01T12:00:00.000Z", "created_at": "2019-01-01T12:00:00.000Z", "triggers": [{"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "runs_url": "RunsURL", "href": "Href", "build_number": 1, "next_build_number": 1, "enable_notifications": false, "enable_partial_cloning": true, "enabled": true}`)
				}))
			})
			It(`Invoke CreateTektonPipeline successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("public")

				// Construct an instance of the CreateTektonPipelineOptions model
				createTektonPipelineOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineOptions)
				createTektonPipelineOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineOptionsModel.NextBuildNumber = core.Int64Ptr(int64(1))
				createTektonPipelineOptionsModel.EnableNotifications = core.BoolPtr(false)
				createTektonPipelineOptionsModel.EnablePartialCloning = core.BoolPtr(false)
				createTektonPipelineOptionsModel.Worker = workerIdentityModel
				createTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.CreateTektonPipelineWithContext(ctx, createTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipeline(createTektonPipelineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.CreateTektonPipelineWithContext(ctx, createTektonPipelineOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelinePath))
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "status": "configured", "resource_group": {"id": "ID"}, "toolchain": {"id": "ID", "crn": "crn:v1:staging:public:toolchain:us-south:a/0ba224679d6c697f9baee5e14ade83ac:bf5fa00f-ddef-4298-b87b-aa8b6da0e1a6::"}, "id": "ID", "definitions": [{"source": {"type": "Type", "properties": {"url": "URL", "branch": "Branch", "tag": "Tag", "path": "Path", "tool": {"id": "ID"}}}, "href": "Href", "id": "ID"}], "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}], "updated_at": "2019-01-01T12:00:00.000Z", "created_at": "2019-01-01T12:00:00.000Z", "triggers": [{"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "runs_url": "RunsURL", "href": "Href", "build_number": 1, "next_build_number": 1, "enable_notifications": false, "enable_partial_cloning": true, "enabled": true}`)
				}))
			})
			It(`Invoke CreateTektonPipeline successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipeline(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("public")

				// Construct an instance of the CreateTektonPipelineOptions model
				createTektonPipelineOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineOptions)
				createTektonPipelineOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineOptionsModel.NextBuildNumber = core.Int64Ptr(int64(1))
				createTektonPipelineOptionsModel.EnableNotifications = core.BoolPtr(false)
				createTektonPipelineOptionsModel.EnablePartialCloning = core.BoolPtr(false)
				createTektonPipelineOptionsModel.Worker = workerIdentityModel
				createTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.CreateTektonPipeline(createTektonPipelineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTektonPipeline with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("public")

				// Construct an instance of the CreateTektonPipelineOptions model
				createTektonPipelineOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineOptions)
				createTektonPipelineOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineOptionsModel.NextBuildNumber = core.Int64Ptr(int64(1))
				createTektonPipelineOptionsModel.EnableNotifications = core.BoolPtr(false)
				createTektonPipelineOptionsModel.EnablePartialCloning = core.BoolPtr(false)
				createTektonPipelineOptionsModel.Worker = workerIdentityModel
				createTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipeline(createTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTektonPipelineOptions model with no property values
				createTektonPipelineOptionsModelNew := new(cdtektonpipelinev2.CreateTektonPipelineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.CreateTektonPipeline(createTektonPipelineOptionsModelNew)
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
			It(`Invoke CreateTektonPipeline successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("public")

				// Construct an instance of the CreateTektonPipelineOptions model
				createTektonPipelineOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineOptions)
				createTektonPipelineOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineOptionsModel.NextBuildNumber = core.Int64Ptr(int64(1))
				createTektonPipelineOptionsModel.EnableNotifications = core.BoolPtr(false)
				createTektonPipelineOptionsModel.EnablePartialCloning = core.BoolPtr(false)
				createTektonPipelineOptionsModel.Worker = workerIdentityModel
				createTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipeline(createTektonPipelineOptionsModel)
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
	Describe(`GetTektonPipeline(getTektonPipelineOptions *GetTektonPipelineOptions) - Operation response error`, func() {
		getTektonPipelinePath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelinePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTektonPipeline with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineOptions model
				getTektonPipelineOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineOptions)
				getTektonPipelineOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.GetTektonPipeline(getTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipeline(getTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTektonPipeline(getTektonPipelineOptions *GetTektonPipelineOptions)`, func() {
		getTektonPipelinePath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelinePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "status": "configured", "resource_group": {"id": "ID"}, "toolchain": {"id": "ID", "crn": "crn:v1:staging:public:toolchain:us-south:a/0ba224679d6c697f9baee5e14ade83ac:bf5fa00f-ddef-4298-b87b-aa8b6da0e1a6::"}, "id": "ID", "definitions": [{"source": {"type": "Type", "properties": {"url": "URL", "branch": "Branch", "tag": "Tag", "path": "Path", "tool": {"id": "ID"}}}, "href": "Href", "id": "ID"}], "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}], "updated_at": "2019-01-01T12:00:00.000Z", "created_at": "2019-01-01T12:00:00.000Z", "triggers": [{"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "runs_url": "RunsURL", "href": "Href", "build_number": 1, "next_build_number": 1, "enable_notifications": false, "enable_partial_cloning": true, "enabled": true}`)
				}))
			})
			It(`Invoke GetTektonPipeline successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the GetTektonPipelineOptions model
				getTektonPipelineOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineOptions)
				getTektonPipelineOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.GetTektonPipelineWithContext(ctx, getTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.GetTektonPipeline(getTektonPipelineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.GetTektonPipelineWithContext(ctx, getTektonPipelineOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelinePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "status": "configured", "resource_group": {"id": "ID"}, "toolchain": {"id": "ID", "crn": "crn:v1:staging:public:toolchain:us-south:a/0ba224679d6c697f9baee5e14ade83ac:bf5fa00f-ddef-4298-b87b-aa8b6da0e1a6::"}, "id": "ID", "definitions": [{"source": {"type": "Type", "properties": {"url": "URL", "branch": "Branch", "tag": "Tag", "path": "Path", "tool": {"id": "ID"}}}, "href": "Href", "id": "ID"}], "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}], "updated_at": "2019-01-01T12:00:00.000Z", "created_at": "2019-01-01T12:00:00.000Z", "triggers": [{"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "runs_url": "RunsURL", "href": "Href", "build_number": 1, "next_build_number": 1, "enable_notifications": false, "enable_partial_cloning": true, "enabled": true}`)
				}))
			})
			It(`Invoke GetTektonPipeline successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.GetTektonPipeline(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTektonPipelineOptions model
				getTektonPipelineOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineOptions)
				getTektonPipelineOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipeline(getTektonPipelineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTektonPipeline with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineOptions model
				getTektonPipelineOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineOptions)
				getTektonPipelineOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.GetTektonPipeline(getTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTektonPipelineOptions model with no property values
				getTektonPipelineOptionsModelNew := new(cdtektonpipelinev2.GetTektonPipelineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipeline(getTektonPipelineOptionsModelNew)
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
			It(`Invoke GetTektonPipeline successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineOptions model
				getTektonPipelineOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineOptions)
				getTektonPipelineOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.GetTektonPipeline(getTektonPipelineOptionsModel)
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
	Describe(`UpdateTektonPipeline(updateTektonPipelineOptions *UpdateTektonPipelineOptions) - Operation response error`, func() {
		updateTektonPipelinePath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTektonPipelinePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTektonPipeline with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("public")

				// Construct an instance of the TektonPipelinePatch model
				tektonPipelinePatchModel := new(cdtektonpipelinev2.TektonPipelinePatch)
				tektonPipelinePatchModel.NextBuildNumber = core.Int64Ptr(int64(1))
				tektonPipelinePatchModel.EnableNotifications = core.BoolPtr(true)
				tektonPipelinePatchModel.EnablePartialCloning = core.BoolPtr(true)
				tektonPipelinePatchModel.Worker = workerIdentityModel
				tektonPipelinePatchModelAsPatch, asPatchErr := tektonPipelinePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTektonPipelineOptions model
				updateTektonPipelineOptionsModel := new(cdtektonpipelinev2.UpdateTektonPipelineOptions)
				updateTektonPipelineOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				updateTektonPipelineOptionsModel.TektonPipelinePatch = tektonPipelinePatchModelAsPatch
				updateTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.UpdateTektonPipeline(updateTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.UpdateTektonPipeline(updateTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTektonPipeline(updateTektonPipelineOptions *UpdateTektonPipelineOptions)`, func() {
		updateTektonPipelinePath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTektonPipelinePath))
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "status": "configured", "resource_group": {"id": "ID"}, "toolchain": {"id": "ID", "crn": "crn:v1:staging:public:toolchain:us-south:a/0ba224679d6c697f9baee5e14ade83ac:bf5fa00f-ddef-4298-b87b-aa8b6da0e1a6::"}, "id": "ID", "definitions": [{"source": {"type": "Type", "properties": {"url": "URL", "branch": "Branch", "tag": "Tag", "path": "Path", "tool": {"id": "ID"}}}, "href": "Href", "id": "ID"}], "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}], "updated_at": "2019-01-01T12:00:00.000Z", "created_at": "2019-01-01T12:00:00.000Z", "triggers": [{"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "runs_url": "RunsURL", "href": "Href", "build_number": 1, "next_build_number": 1, "enable_notifications": false, "enable_partial_cloning": true, "enabled": true}`)
				}))
			})
			It(`Invoke UpdateTektonPipeline successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("public")

				// Construct an instance of the TektonPipelinePatch model
				tektonPipelinePatchModel := new(cdtektonpipelinev2.TektonPipelinePatch)
				tektonPipelinePatchModel.NextBuildNumber = core.Int64Ptr(int64(1))
				tektonPipelinePatchModel.EnableNotifications = core.BoolPtr(true)
				tektonPipelinePatchModel.EnablePartialCloning = core.BoolPtr(true)
				tektonPipelinePatchModel.Worker = workerIdentityModel
				tektonPipelinePatchModelAsPatch, asPatchErr := tektonPipelinePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTektonPipelineOptions model
				updateTektonPipelineOptionsModel := new(cdtektonpipelinev2.UpdateTektonPipelineOptions)
				updateTektonPipelineOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				updateTektonPipelineOptionsModel.TektonPipelinePatch = tektonPipelinePatchModelAsPatch
				updateTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.UpdateTektonPipelineWithContext(ctx, updateTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.UpdateTektonPipeline(updateTektonPipelineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.UpdateTektonPipelineWithContext(ctx, updateTektonPipelineOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateTektonPipelinePath))
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "status": "configured", "resource_group": {"id": "ID"}, "toolchain": {"id": "ID", "crn": "crn:v1:staging:public:toolchain:us-south:a/0ba224679d6c697f9baee5e14ade83ac:bf5fa00f-ddef-4298-b87b-aa8b6da0e1a6::"}, "id": "ID", "definitions": [{"source": {"type": "Type", "properties": {"url": "URL", "branch": "Branch", "tag": "Tag", "path": "Path", "tool": {"id": "ID"}}}, "href": "Href", "id": "ID"}], "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}], "updated_at": "2019-01-01T12:00:00.000Z", "created_at": "2019-01-01T12:00:00.000Z", "triggers": [{"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "runs_url": "RunsURL", "href": "Href", "build_number": 1, "next_build_number": 1, "enable_notifications": false, "enable_partial_cloning": true, "enabled": true}`)
				}))
			})
			It(`Invoke UpdateTektonPipeline successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.UpdateTektonPipeline(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("public")

				// Construct an instance of the TektonPipelinePatch model
				tektonPipelinePatchModel := new(cdtektonpipelinev2.TektonPipelinePatch)
				tektonPipelinePatchModel.NextBuildNumber = core.Int64Ptr(int64(1))
				tektonPipelinePatchModel.EnableNotifications = core.BoolPtr(true)
				tektonPipelinePatchModel.EnablePartialCloning = core.BoolPtr(true)
				tektonPipelinePatchModel.Worker = workerIdentityModel
				tektonPipelinePatchModelAsPatch, asPatchErr := tektonPipelinePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTektonPipelineOptions model
				updateTektonPipelineOptionsModel := new(cdtektonpipelinev2.UpdateTektonPipelineOptions)
				updateTektonPipelineOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				updateTektonPipelineOptionsModel.TektonPipelinePatch = tektonPipelinePatchModelAsPatch
				updateTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.UpdateTektonPipeline(updateTektonPipelineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTektonPipeline with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("public")

				// Construct an instance of the TektonPipelinePatch model
				tektonPipelinePatchModel := new(cdtektonpipelinev2.TektonPipelinePatch)
				tektonPipelinePatchModel.NextBuildNumber = core.Int64Ptr(int64(1))
				tektonPipelinePatchModel.EnableNotifications = core.BoolPtr(true)
				tektonPipelinePatchModel.EnablePartialCloning = core.BoolPtr(true)
				tektonPipelinePatchModel.Worker = workerIdentityModel
				tektonPipelinePatchModelAsPatch, asPatchErr := tektonPipelinePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTektonPipelineOptions model
				updateTektonPipelineOptionsModel := new(cdtektonpipelinev2.UpdateTektonPipelineOptions)
				updateTektonPipelineOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				updateTektonPipelineOptionsModel.TektonPipelinePatch = tektonPipelinePatchModelAsPatch
				updateTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.UpdateTektonPipeline(updateTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTektonPipelineOptions model with no property values
				updateTektonPipelineOptionsModelNew := new(cdtektonpipelinev2.UpdateTektonPipelineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.UpdateTektonPipeline(updateTektonPipelineOptionsModelNew)
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
			It(`Invoke UpdateTektonPipeline successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("public")

				// Construct an instance of the TektonPipelinePatch model
				tektonPipelinePatchModel := new(cdtektonpipelinev2.TektonPipelinePatch)
				tektonPipelinePatchModel.NextBuildNumber = core.Int64Ptr(int64(1))
				tektonPipelinePatchModel.EnableNotifications = core.BoolPtr(true)
				tektonPipelinePatchModel.EnablePartialCloning = core.BoolPtr(true)
				tektonPipelinePatchModel.Worker = workerIdentityModel
				tektonPipelinePatchModelAsPatch, asPatchErr := tektonPipelinePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTektonPipelineOptions model
				updateTektonPipelineOptionsModel := new(cdtektonpipelinev2.UpdateTektonPipelineOptions)
				updateTektonPipelineOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				updateTektonPipelineOptionsModel.TektonPipelinePatch = tektonPipelinePatchModelAsPatch
				updateTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.UpdateTektonPipeline(updateTektonPipelineOptionsModel)
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
	Describe(`DeleteTektonPipeline(deleteTektonPipelineOptions *DeleteTektonPipelineOptions)`, func() {
		deleteTektonPipelinePath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTektonPipelinePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTektonPipeline successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := cdTektonPipelineService.DeleteTektonPipeline(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTektonPipelineOptions model
				deleteTektonPipelineOptionsModel := new(cdtektonpipelinev2.DeleteTektonPipelineOptions)
				deleteTektonPipelineOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = cdTektonPipelineService.DeleteTektonPipeline(deleteTektonPipelineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTektonPipeline with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the DeleteTektonPipelineOptions model
				deleteTektonPipelineOptionsModel := new(cdtektonpipelinev2.DeleteTektonPipelineOptions)
				deleteTektonPipelineOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := cdTektonPipelineService.DeleteTektonPipeline(deleteTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTektonPipelineOptions model with no property values
				deleteTektonPipelineOptionsModelNew := new(cdtektonpipelinev2.DeleteTektonPipelineOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = cdTektonPipelineService.DeleteTektonPipeline(deleteTektonPipelineOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTektonPipelineRuns(listTektonPipelineRunsOptions *ListTektonPipelineRunsOptions) - Operation response error`, func() {
		listTektonPipelineRunsPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/pipeline_runs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTektonPipelineRunsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"succeeded"}))
					Expect(req.URL.Query()["trigger.name"]).To(Equal([]string{"manual-trigger"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTektonPipelineRuns with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ListTektonPipelineRunsOptions model
				listTektonPipelineRunsOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineRunsOptions)
				listTektonPipelineRunsOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineRunsOptionsModel.Start = core.StringPtr("testString")
				listTektonPipelineRunsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTektonPipelineRunsOptionsModel.Status = core.StringPtr("succeeded")
				listTektonPipelineRunsOptionsModel.TriggerName = core.StringPtr("manual-trigger")
				listTektonPipelineRunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineRuns(listTektonPipelineRunsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.ListTektonPipelineRuns(listTektonPipelineRunsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTektonPipelineRuns(listTektonPipelineRunsOptions *ListTektonPipelineRunsOptions)`, func() {
		listTektonPipelineRunsPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/pipeline_runs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTektonPipelineRunsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"succeeded"}))
					Expect(req.URL.Query()["trigger.name"]).To(Equal([]string{"manual-trigger"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"pipeline_runs": [{"id": "ID", "href": "Href", "user_info": {"iam_id": "IamID", "sub": "Sub"}, "status": "pending", "definition_id": "DefinitionID", "definition": {"id": "ID"}, "description": "Description", "worker": {"name": "Name", "agent_id": "AgentID", "service_id": "ServiceID", "id": "ID"}, "pipeline_id": "PipelineID", "pipeline": {"id": "ID"}, "listener_name": "ListenerName", "trigger": {"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}, "event_params_blob": "EventParamsBlob", "trigger_headers": "TriggerHeaders", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}], "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "run_url": "RunURL", "error_message": "ErrorMessage"}], "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListTektonPipelineRuns successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the ListTektonPipelineRunsOptions model
				listTektonPipelineRunsOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineRunsOptions)
				listTektonPipelineRunsOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineRunsOptionsModel.Start = core.StringPtr("testString")
				listTektonPipelineRunsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTektonPipelineRunsOptionsModel.Status = core.StringPtr("succeeded")
				listTektonPipelineRunsOptionsModel.TriggerName = core.StringPtr("manual-trigger")
				listTektonPipelineRunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.ListTektonPipelineRunsWithContext(ctx, listTektonPipelineRunsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineRuns(listTektonPipelineRunsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.ListTektonPipelineRunsWithContext(ctx, listTektonPipelineRunsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTektonPipelineRunsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"succeeded"}))
					Expect(req.URL.Query()["trigger.name"]).To(Equal([]string{"manual-trigger"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"pipeline_runs": [{"id": "ID", "href": "Href", "user_info": {"iam_id": "IamID", "sub": "Sub"}, "status": "pending", "definition_id": "DefinitionID", "definition": {"id": "ID"}, "description": "Description", "worker": {"name": "Name", "agent_id": "AgentID", "service_id": "ServiceID", "id": "ID"}, "pipeline_id": "PipelineID", "pipeline": {"id": "ID"}, "listener_name": "ListenerName", "trigger": {"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}, "event_params_blob": "EventParamsBlob", "trigger_headers": "TriggerHeaders", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}], "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "run_url": "RunURL", "error_message": "ErrorMessage"}], "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListTektonPipelineRuns successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineRuns(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTektonPipelineRunsOptions model
				listTektonPipelineRunsOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineRunsOptions)
				listTektonPipelineRunsOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineRunsOptionsModel.Start = core.StringPtr("testString")
				listTektonPipelineRunsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTektonPipelineRunsOptionsModel.Status = core.StringPtr("succeeded")
				listTektonPipelineRunsOptionsModel.TriggerName = core.StringPtr("manual-trigger")
				listTektonPipelineRunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.ListTektonPipelineRuns(listTektonPipelineRunsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTektonPipelineRuns with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ListTektonPipelineRunsOptions model
				listTektonPipelineRunsOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineRunsOptions)
				listTektonPipelineRunsOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineRunsOptionsModel.Start = core.StringPtr("testString")
				listTektonPipelineRunsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTektonPipelineRunsOptionsModel.Status = core.StringPtr("succeeded")
				listTektonPipelineRunsOptionsModel.TriggerName = core.StringPtr("manual-trigger")
				listTektonPipelineRunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineRuns(listTektonPipelineRunsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTektonPipelineRunsOptions model with no property values
				listTektonPipelineRunsOptionsModelNew := new(cdtektonpipelinev2.ListTektonPipelineRunsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.ListTektonPipelineRuns(listTektonPipelineRunsOptionsModelNew)
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
			It(`Invoke ListTektonPipelineRuns successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ListTektonPipelineRunsOptions model
				listTektonPipelineRunsOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineRunsOptions)
				listTektonPipelineRunsOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineRunsOptionsModel.Start = core.StringPtr("testString")
				listTektonPipelineRunsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTektonPipelineRunsOptionsModel.Status = core.StringPtr("succeeded")
				listTektonPipelineRunsOptionsModel.TriggerName = core.StringPtr("manual-trigger")
				listTektonPipelineRunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineRuns(listTektonPipelineRunsOptionsModel)
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
				responseObject := new(cdtektonpipelinev2.PipelineRunsCollection)
				nextObject := new(cdtektonpipelinev2.RunsNextPage)
				nextObject.Href = core.StringPtr("ibm.com?start=abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(cdtektonpipelinev2.PipelineRunsCollection)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "Next" URL`, func() {
				responseObject := new(cdtektonpipelinev2.PipelineRunsCollection)
				nextObject := new(cdtektonpipelinev2.RunsNextPage)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

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
					Expect(req.URL.EscapedPath()).To(Equal(listTektonPipelineRunsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?start=1"},"total_count":2,"limit":1,"pipeline_runs":[{"id":"ID","href":"Href","user_info":{"iam_id":"IamID","sub":"Sub"},"status":"pending","definition_id":"DefinitionID","definition":{"id":"ID"},"description":"Description","worker":{"name":"Name","agent_id":"AgentID","service_id":"ServiceID","id":"ID"},"pipeline_id":"PipelineID","pipeline":{"id":"ID"},"listener_name":"ListenerName","trigger":{"type":"Type","name":"start-deploy","href":"Href","event_listener":"EventListener","id":"ID","properties":[{"name":"Name","value":"Value","href":"Href","enum":["Enum"],"type":"secure","path":"Path","locked":true}],"tags":["Tags"],"worker":{"name":"Name","type":"Type","id":"ID"},"max_concurrent_runs":4,"enabled":true,"favorite":false},"event_params_blob":"EventParamsBlob","trigger_headers":"TriggerHeaders","properties":[{"name":"Name","value":"Value","href":"Href","enum":["Enum"],"type":"secure","locked":true,"path":"Path"}],"created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z","run_url":"RunURL","error_message":"ErrorMessage"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"pipeline_runs":[{"id":"ID","href":"Href","user_info":{"iam_id":"IamID","sub":"Sub"},"status":"pending","definition_id":"DefinitionID","definition":{"id":"ID"},"description":"Description","worker":{"name":"Name","agent_id":"AgentID","service_id":"ServiceID","id":"ID"},"pipeline_id":"PipelineID","pipeline":{"id":"ID"},"listener_name":"ListenerName","trigger":{"type":"Type","name":"start-deploy","href":"Href","event_listener":"EventListener","id":"ID","properties":[{"name":"Name","value":"Value","href":"Href","enum":["Enum"],"type":"secure","path":"Path","locked":true}],"tags":["Tags"],"worker":{"name":"Name","type":"Type","id":"ID"},"max_concurrent_runs":4,"enabled":true,"favorite":false},"event_params_blob":"EventParamsBlob","trigger_headers":"TriggerHeaders","properties":[{"name":"Name","value":"Value","href":"Href","enum":["Enum"],"type":"secure","locked":true,"path":"Path"}],"created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z","run_url":"RunURL","error_message":"ErrorMessage"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use TektonPipelineRunsPager.GetNext successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				listTektonPipelineRunsOptionsModel := &cdtektonpipelinev2.ListTektonPipelineRunsOptions{
					PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
					Limit: core.Int64Ptr(int64(10)),
					Status: core.StringPtr("succeeded"),
					TriggerName: core.StringPtr("manual-trigger"),
				}

				pager, err := cdTektonPipelineService.NewTektonPipelineRunsPager(listTektonPipelineRunsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []cdtektonpipelinev2.PipelineRun
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use TektonPipelineRunsPager.GetAll successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				listTektonPipelineRunsOptionsModel := &cdtektonpipelinev2.ListTektonPipelineRunsOptions{
					PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
					Limit: core.Int64Ptr(int64(10)),
					Status: core.StringPtr("succeeded"),
					TriggerName: core.StringPtr("manual-trigger"),
				}

				pager, err := cdTektonPipelineService.NewTektonPipelineRunsPager(listTektonPipelineRunsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateTektonPipelineRun(createTektonPipelineRunOptions *CreateTektonPipelineRunOptions) - Operation response error`, func() {
		createTektonPipelineRunPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/pipeline_runs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelineRunPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTektonPipelineRun with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the PipelineRunTrigger model
				pipelineRunTriggerModel := new(cdtektonpipelinev2.PipelineRunTrigger)
				pipelineRunTriggerModel.Name = core.StringPtr("Manual Trigger 1")
				pipelineRunTriggerModel.Properties = map[string]interface{}{"anyKey": "anyValue"}
				pipelineRunTriggerModel.SecureProperties = map[string]interface{}{"anyKey": "anyValue"}
				pipelineRunTriggerModel.HeadersVar = map[string]interface{}{"anyKey": "anyValue"}
				pipelineRunTriggerModel.Body = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the CreateTektonPipelineRunOptions model
				createTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineRunOptions)
				createTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineRunOptionsModel.Description = core.StringPtr("My custom manual PipelineRun")
				createTektonPipelineRunOptionsModel.TriggerName = core.StringPtr("testString")
				createTektonPipelineRunOptionsModel.TriggerProperties = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.SecureTriggerProperties = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.TriggerHeaders = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.TriggerBody = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.Trigger = pipelineRunTriggerModel
				createTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineRun(createTektonPipelineRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.CreateTektonPipelineRun(createTektonPipelineRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTektonPipelineRun(createTektonPipelineRunOptions *CreateTektonPipelineRunOptions)`, func() {
		createTektonPipelineRunPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/pipeline_runs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelineRunPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "href": "Href", "user_info": {"iam_id": "IamID", "sub": "Sub"}, "status": "pending", "definition_id": "DefinitionID", "definition": {"id": "ID"}, "description": "Description", "worker": {"name": "Name", "agent_id": "AgentID", "service_id": "ServiceID", "id": "ID"}, "pipeline_id": "PipelineID", "pipeline": {"id": "ID"}, "listener_name": "ListenerName", "trigger": {"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}, "event_params_blob": "EventParamsBlob", "trigger_headers": "TriggerHeaders", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}], "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "run_url": "RunURL", "error_message": "ErrorMessage"}`)
				}))
			})
			It(`Invoke CreateTektonPipelineRun successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the PipelineRunTrigger model
				pipelineRunTriggerModel := new(cdtektonpipelinev2.PipelineRunTrigger)
				pipelineRunTriggerModel.Name = core.StringPtr("Manual Trigger 1")
				pipelineRunTriggerModel.Properties = map[string]interface{}{"anyKey": "anyValue"}
				pipelineRunTriggerModel.SecureProperties = map[string]interface{}{"anyKey": "anyValue"}
				pipelineRunTriggerModel.HeadersVar = map[string]interface{}{"anyKey": "anyValue"}
				pipelineRunTriggerModel.Body = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the CreateTektonPipelineRunOptions model
				createTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineRunOptions)
				createTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineRunOptionsModel.Description = core.StringPtr("My custom manual PipelineRun")
				createTektonPipelineRunOptionsModel.TriggerName = core.StringPtr("testString")
				createTektonPipelineRunOptionsModel.TriggerProperties = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.SecureTriggerProperties = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.TriggerHeaders = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.TriggerBody = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.Trigger = pipelineRunTriggerModel
				createTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.CreateTektonPipelineRunWithContext(ctx, createTektonPipelineRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineRun(createTektonPipelineRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.CreateTektonPipelineRunWithContext(ctx, createTektonPipelineRunOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelineRunPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "href": "Href", "user_info": {"iam_id": "IamID", "sub": "Sub"}, "status": "pending", "definition_id": "DefinitionID", "definition": {"id": "ID"}, "description": "Description", "worker": {"name": "Name", "agent_id": "AgentID", "service_id": "ServiceID", "id": "ID"}, "pipeline_id": "PipelineID", "pipeline": {"id": "ID"}, "listener_name": "ListenerName", "trigger": {"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}, "event_params_blob": "EventParamsBlob", "trigger_headers": "TriggerHeaders", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}], "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "run_url": "RunURL", "error_message": "ErrorMessage"}`)
				}))
			})
			It(`Invoke CreateTektonPipelineRun successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineRun(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PipelineRunTrigger model
				pipelineRunTriggerModel := new(cdtektonpipelinev2.PipelineRunTrigger)
				pipelineRunTriggerModel.Name = core.StringPtr("Manual Trigger 1")
				pipelineRunTriggerModel.Properties = map[string]interface{}{"anyKey": "anyValue"}
				pipelineRunTriggerModel.SecureProperties = map[string]interface{}{"anyKey": "anyValue"}
				pipelineRunTriggerModel.HeadersVar = map[string]interface{}{"anyKey": "anyValue"}
				pipelineRunTriggerModel.Body = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the CreateTektonPipelineRunOptions model
				createTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineRunOptions)
				createTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineRunOptionsModel.Description = core.StringPtr("My custom manual PipelineRun")
				createTektonPipelineRunOptionsModel.TriggerName = core.StringPtr("testString")
				createTektonPipelineRunOptionsModel.TriggerProperties = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.SecureTriggerProperties = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.TriggerHeaders = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.TriggerBody = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.Trigger = pipelineRunTriggerModel
				createTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.CreateTektonPipelineRun(createTektonPipelineRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTektonPipelineRun with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the PipelineRunTrigger model
				pipelineRunTriggerModel := new(cdtektonpipelinev2.PipelineRunTrigger)
				pipelineRunTriggerModel.Name = core.StringPtr("Manual Trigger 1")
				pipelineRunTriggerModel.Properties = map[string]interface{}{"anyKey": "anyValue"}
				pipelineRunTriggerModel.SecureProperties = map[string]interface{}{"anyKey": "anyValue"}
				pipelineRunTriggerModel.HeadersVar = map[string]interface{}{"anyKey": "anyValue"}
				pipelineRunTriggerModel.Body = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the CreateTektonPipelineRunOptions model
				createTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineRunOptions)
				createTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineRunOptionsModel.Description = core.StringPtr("My custom manual PipelineRun")
				createTektonPipelineRunOptionsModel.TriggerName = core.StringPtr("testString")
				createTektonPipelineRunOptionsModel.TriggerProperties = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.SecureTriggerProperties = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.TriggerHeaders = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.TriggerBody = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.Trigger = pipelineRunTriggerModel
				createTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineRun(createTektonPipelineRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTektonPipelineRunOptions model with no property values
				createTektonPipelineRunOptionsModelNew := new(cdtektonpipelinev2.CreateTektonPipelineRunOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.CreateTektonPipelineRun(createTektonPipelineRunOptionsModelNew)
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
			It(`Invoke CreateTektonPipelineRun successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the PipelineRunTrigger model
				pipelineRunTriggerModel := new(cdtektonpipelinev2.PipelineRunTrigger)
				pipelineRunTriggerModel.Name = core.StringPtr("Manual Trigger 1")
				pipelineRunTriggerModel.Properties = map[string]interface{}{"anyKey": "anyValue"}
				pipelineRunTriggerModel.SecureProperties = map[string]interface{}{"anyKey": "anyValue"}
				pipelineRunTriggerModel.HeadersVar = map[string]interface{}{"anyKey": "anyValue"}
				pipelineRunTriggerModel.Body = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the CreateTektonPipelineRunOptions model
				createTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineRunOptions)
				createTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineRunOptionsModel.Description = core.StringPtr("My custom manual PipelineRun")
				createTektonPipelineRunOptionsModel.TriggerName = core.StringPtr("testString")
				createTektonPipelineRunOptionsModel.TriggerProperties = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.SecureTriggerProperties = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.TriggerHeaders = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.TriggerBody = map[string]interface{}{"anyKey": "anyValue"}
				createTektonPipelineRunOptionsModel.Trigger = pipelineRunTriggerModel
				createTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineRun(createTektonPipelineRunOptionsModel)
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
	Describe(`GetTektonPipelineRun(getTektonPipelineRunOptions *GetTektonPipelineRunOptions) - Operation response error`, func() {
		getTektonPipelineRunPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/pipeline_runs/94619026-912b-4d92-8f51-6c74f0692d90"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineRunPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["includes"]).To(Equal([]string{"definitions"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTektonPipelineRun with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineRunOptions model
				getTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineRunOptions)
				getTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunOptionsModel.Includes = core.StringPtr("definitions")
				getTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineRun(getTektonPipelineRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineRun(getTektonPipelineRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTektonPipelineRun(getTektonPipelineRunOptions *GetTektonPipelineRunOptions)`, func() {
		getTektonPipelineRunPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/pipeline_runs/94619026-912b-4d92-8f51-6c74f0692d90"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineRunPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["includes"]).To(Equal([]string{"definitions"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "href": "Href", "user_info": {"iam_id": "IamID", "sub": "Sub"}, "status": "pending", "definition_id": "DefinitionID", "definition": {"id": "ID"}, "description": "Description", "worker": {"name": "Name", "agent_id": "AgentID", "service_id": "ServiceID", "id": "ID"}, "pipeline_id": "PipelineID", "pipeline": {"id": "ID"}, "listener_name": "ListenerName", "trigger": {"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}, "event_params_blob": "EventParamsBlob", "trigger_headers": "TriggerHeaders", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}], "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "run_url": "RunURL", "error_message": "ErrorMessage"}`)
				}))
			})
			It(`Invoke GetTektonPipelineRun successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the GetTektonPipelineRunOptions model
				getTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineRunOptions)
				getTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunOptionsModel.Includes = core.StringPtr("definitions")
				getTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.GetTektonPipelineRunWithContext(ctx, getTektonPipelineRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineRun(getTektonPipelineRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.GetTektonPipelineRunWithContext(ctx, getTektonPipelineRunOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineRunPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["includes"]).To(Equal([]string{"definitions"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "href": "Href", "user_info": {"iam_id": "IamID", "sub": "Sub"}, "status": "pending", "definition_id": "DefinitionID", "definition": {"id": "ID"}, "description": "Description", "worker": {"name": "Name", "agent_id": "AgentID", "service_id": "ServiceID", "id": "ID"}, "pipeline_id": "PipelineID", "pipeline": {"id": "ID"}, "listener_name": "ListenerName", "trigger": {"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}, "event_params_blob": "EventParamsBlob", "trigger_headers": "TriggerHeaders", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}], "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "run_url": "RunURL", "error_message": "ErrorMessage"}`)
				}))
			})
			It(`Invoke GetTektonPipelineRun successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineRun(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTektonPipelineRunOptions model
				getTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineRunOptions)
				getTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunOptionsModel.Includes = core.StringPtr("definitions")
				getTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineRun(getTektonPipelineRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTektonPipelineRun with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineRunOptions model
				getTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineRunOptions)
				getTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunOptionsModel.Includes = core.StringPtr("definitions")
				getTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineRun(getTektonPipelineRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTektonPipelineRunOptions model with no property values
				getTektonPipelineRunOptionsModelNew := new(cdtektonpipelinev2.GetTektonPipelineRunOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineRun(getTektonPipelineRunOptionsModelNew)
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
			It(`Invoke GetTektonPipelineRun successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineRunOptions model
				getTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineRunOptions)
				getTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunOptionsModel.Includes = core.StringPtr("definitions")
				getTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineRun(getTektonPipelineRunOptionsModel)
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
	Describe(`DeleteTektonPipelineRun(deleteTektonPipelineRunOptions *DeleteTektonPipelineRunOptions)`, func() {
		deleteTektonPipelineRunPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/pipeline_runs/94619026-912b-4d92-8f51-6c74f0692d90"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTektonPipelineRunPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTektonPipelineRun successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := cdTektonPipelineService.DeleteTektonPipelineRun(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTektonPipelineRunOptions model
				deleteTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.DeleteTektonPipelineRunOptions)
				deleteTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelineRunOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = cdTektonPipelineService.DeleteTektonPipelineRun(deleteTektonPipelineRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTektonPipelineRun with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the DeleteTektonPipelineRunOptions model
				deleteTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.DeleteTektonPipelineRunOptions)
				deleteTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelineRunOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := cdTektonPipelineService.DeleteTektonPipelineRun(deleteTektonPipelineRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTektonPipelineRunOptions model with no property values
				deleteTektonPipelineRunOptionsModelNew := new(cdtektonpipelinev2.DeleteTektonPipelineRunOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = cdTektonPipelineService.DeleteTektonPipelineRun(deleteTektonPipelineRunOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CancelTektonPipelineRun(cancelTektonPipelineRunOptions *CancelTektonPipelineRunOptions) - Operation response error`, func() {
		cancelTektonPipelineRunPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/pipeline_runs/94619026-912b-4d92-8f51-6c74f0692d90/cancel"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(cancelTektonPipelineRunPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CancelTektonPipelineRun with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the CancelTektonPipelineRunOptions model
				cancelTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.CancelTektonPipelineRunOptions)
				cancelTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				cancelTektonPipelineRunOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				cancelTektonPipelineRunOptionsModel.Force = core.BoolPtr(true)
				cancelTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.CancelTektonPipelineRun(cancelTektonPipelineRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.CancelTektonPipelineRun(cancelTektonPipelineRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CancelTektonPipelineRun(cancelTektonPipelineRunOptions *CancelTektonPipelineRunOptions)`, func() {
		cancelTektonPipelineRunPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/pipeline_runs/94619026-912b-4d92-8f51-6c74f0692d90/cancel"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(cancelTektonPipelineRunPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "href": "Href", "user_info": {"iam_id": "IamID", "sub": "Sub"}, "status": "pending", "definition_id": "DefinitionID", "definition": {"id": "ID"}, "description": "Description", "worker": {"name": "Name", "agent_id": "AgentID", "service_id": "ServiceID", "id": "ID"}, "pipeline_id": "PipelineID", "pipeline": {"id": "ID"}, "listener_name": "ListenerName", "trigger": {"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}, "event_params_blob": "EventParamsBlob", "trigger_headers": "TriggerHeaders", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}], "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "run_url": "RunURL", "error_message": "ErrorMessage"}`)
				}))
			})
			It(`Invoke CancelTektonPipelineRun successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the CancelTektonPipelineRunOptions model
				cancelTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.CancelTektonPipelineRunOptions)
				cancelTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				cancelTektonPipelineRunOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				cancelTektonPipelineRunOptionsModel.Force = core.BoolPtr(true)
				cancelTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.CancelTektonPipelineRunWithContext(ctx, cancelTektonPipelineRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.CancelTektonPipelineRun(cancelTektonPipelineRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.CancelTektonPipelineRunWithContext(ctx, cancelTektonPipelineRunOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(cancelTektonPipelineRunPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "href": "Href", "user_info": {"iam_id": "IamID", "sub": "Sub"}, "status": "pending", "definition_id": "DefinitionID", "definition": {"id": "ID"}, "description": "Description", "worker": {"name": "Name", "agent_id": "AgentID", "service_id": "ServiceID", "id": "ID"}, "pipeline_id": "PipelineID", "pipeline": {"id": "ID"}, "listener_name": "ListenerName", "trigger": {"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}, "event_params_blob": "EventParamsBlob", "trigger_headers": "TriggerHeaders", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}], "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "run_url": "RunURL", "error_message": "ErrorMessage"}`)
				}))
			})
			It(`Invoke CancelTektonPipelineRun successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.CancelTektonPipelineRun(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CancelTektonPipelineRunOptions model
				cancelTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.CancelTektonPipelineRunOptions)
				cancelTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				cancelTektonPipelineRunOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				cancelTektonPipelineRunOptionsModel.Force = core.BoolPtr(true)
				cancelTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.CancelTektonPipelineRun(cancelTektonPipelineRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CancelTektonPipelineRun with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the CancelTektonPipelineRunOptions model
				cancelTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.CancelTektonPipelineRunOptions)
				cancelTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				cancelTektonPipelineRunOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				cancelTektonPipelineRunOptionsModel.Force = core.BoolPtr(true)
				cancelTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.CancelTektonPipelineRun(cancelTektonPipelineRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CancelTektonPipelineRunOptions model with no property values
				cancelTektonPipelineRunOptionsModelNew := new(cdtektonpipelinev2.CancelTektonPipelineRunOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.CancelTektonPipelineRun(cancelTektonPipelineRunOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CancelTektonPipelineRun successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the CancelTektonPipelineRunOptions model
				cancelTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.CancelTektonPipelineRunOptions)
				cancelTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				cancelTektonPipelineRunOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				cancelTektonPipelineRunOptionsModel.Force = core.BoolPtr(true)
				cancelTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.CancelTektonPipelineRun(cancelTektonPipelineRunOptionsModel)
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
	Describe(`RerunTektonPipelineRun(rerunTektonPipelineRunOptions *RerunTektonPipelineRunOptions) - Operation response error`, func() {
		rerunTektonPipelineRunPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/pipeline_runs/94619026-912b-4d92-8f51-6c74f0692d90/rerun"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(rerunTektonPipelineRunPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RerunTektonPipelineRun with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the RerunTektonPipelineRunOptions model
				rerunTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.RerunTektonPipelineRunOptions)
				rerunTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				rerunTektonPipelineRunOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				rerunTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.RerunTektonPipelineRun(rerunTektonPipelineRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.RerunTektonPipelineRun(rerunTektonPipelineRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RerunTektonPipelineRun(rerunTektonPipelineRunOptions *RerunTektonPipelineRunOptions)`, func() {
		rerunTektonPipelineRunPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/pipeline_runs/94619026-912b-4d92-8f51-6c74f0692d90/rerun"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(rerunTektonPipelineRunPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "href": "Href", "user_info": {"iam_id": "IamID", "sub": "Sub"}, "status": "pending", "definition_id": "DefinitionID", "definition": {"id": "ID"}, "description": "Description", "worker": {"name": "Name", "agent_id": "AgentID", "service_id": "ServiceID", "id": "ID"}, "pipeline_id": "PipelineID", "pipeline": {"id": "ID"}, "listener_name": "ListenerName", "trigger": {"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}, "event_params_blob": "EventParamsBlob", "trigger_headers": "TriggerHeaders", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}], "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "run_url": "RunURL", "error_message": "ErrorMessage"}`)
				}))
			})
			It(`Invoke RerunTektonPipelineRun successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the RerunTektonPipelineRunOptions model
				rerunTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.RerunTektonPipelineRunOptions)
				rerunTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				rerunTektonPipelineRunOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				rerunTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.RerunTektonPipelineRunWithContext(ctx, rerunTektonPipelineRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.RerunTektonPipelineRun(rerunTektonPipelineRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.RerunTektonPipelineRunWithContext(ctx, rerunTektonPipelineRunOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(rerunTektonPipelineRunPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "href": "Href", "user_info": {"iam_id": "IamID", "sub": "Sub"}, "status": "pending", "definition_id": "DefinitionID", "definition": {"id": "ID"}, "description": "Description", "worker": {"name": "Name", "agent_id": "AgentID", "service_id": "ServiceID", "id": "ID"}, "pipeline_id": "PipelineID", "pipeline": {"id": "ID"}, "listener_name": "ListenerName", "trigger": {"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}, "event_params_blob": "EventParamsBlob", "trigger_headers": "TriggerHeaders", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}], "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "run_url": "RunURL", "error_message": "ErrorMessage"}`)
				}))
			})
			It(`Invoke RerunTektonPipelineRun successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.RerunTektonPipelineRun(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RerunTektonPipelineRunOptions model
				rerunTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.RerunTektonPipelineRunOptions)
				rerunTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				rerunTektonPipelineRunOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				rerunTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.RerunTektonPipelineRun(rerunTektonPipelineRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RerunTektonPipelineRun with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the RerunTektonPipelineRunOptions model
				rerunTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.RerunTektonPipelineRunOptions)
				rerunTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				rerunTektonPipelineRunOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				rerunTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.RerunTektonPipelineRun(rerunTektonPipelineRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RerunTektonPipelineRunOptions model with no property values
				rerunTektonPipelineRunOptionsModelNew := new(cdtektonpipelinev2.RerunTektonPipelineRunOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.RerunTektonPipelineRun(rerunTektonPipelineRunOptionsModelNew)
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
			It(`Invoke RerunTektonPipelineRun successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the RerunTektonPipelineRunOptions model
				rerunTektonPipelineRunOptionsModel := new(cdtektonpipelinev2.RerunTektonPipelineRunOptions)
				rerunTektonPipelineRunOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				rerunTektonPipelineRunOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				rerunTektonPipelineRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.RerunTektonPipelineRun(rerunTektonPipelineRunOptionsModel)
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
	Describe(`GetTektonPipelineRunLogs(getTektonPipelineRunLogsOptions *GetTektonPipelineRunLogsOptions) - Operation response error`, func() {
		getTektonPipelineRunLogsPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/pipeline_runs/94619026-912b-4d92-8f51-6c74f0692d90/logs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineRunLogsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTektonPipelineRunLogs with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineRunLogsOptions model
				getTektonPipelineRunLogsOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineRunLogsOptions)
				getTektonPipelineRunLogsOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogsOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineRunLogs(getTektonPipelineRunLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineRunLogs(getTektonPipelineRunLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTektonPipelineRunLogs(getTektonPipelineRunLogsOptions *GetTektonPipelineRunLogsOptions)`, func() {
		getTektonPipelineRunLogsPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/pipeline_runs/94619026-912b-4d92-8f51-6c74f0692d90/logs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineRunLogsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"logs": [{"href": "Href", "id": "ID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke GetTektonPipelineRunLogs successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the GetTektonPipelineRunLogsOptions model
				getTektonPipelineRunLogsOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineRunLogsOptions)
				getTektonPipelineRunLogsOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogsOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.GetTektonPipelineRunLogsWithContext(ctx, getTektonPipelineRunLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineRunLogs(getTektonPipelineRunLogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.GetTektonPipelineRunLogsWithContext(ctx, getTektonPipelineRunLogsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineRunLogsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"logs": [{"href": "Href", "id": "ID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke GetTektonPipelineRunLogs successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineRunLogs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTektonPipelineRunLogsOptions model
				getTektonPipelineRunLogsOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineRunLogsOptions)
				getTektonPipelineRunLogsOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogsOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineRunLogs(getTektonPipelineRunLogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTektonPipelineRunLogs with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineRunLogsOptions model
				getTektonPipelineRunLogsOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineRunLogsOptions)
				getTektonPipelineRunLogsOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogsOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineRunLogs(getTektonPipelineRunLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTektonPipelineRunLogsOptions model with no property values
				getTektonPipelineRunLogsOptionsModelNew := new(cdtektonpipelinev2.GetTektonPipelineRunLogsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineRunLogs(getTektonPipelineRunLogsOptionsModelNew)
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
			It(`Invoke GetTektonPipelineRunLogs successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineRunLogsOptions model
				getTektonPipelineRunLogsOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineRunLogsOptions)
				getTektonPipelineRunLogsOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogsOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineRunLogs(getTektonPipelineRunLogsOptionsModel)
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
	Describe(`GetTektonPipelineRunLogContent(getTektonPipelineRunLogContentOptions *GetTektonPipelineRunLogContentOptions) - Operation response error`, func() {
		getTektonPipelineRunLogContentPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/pipeline_runs/bf4b3abd-0c93-416b-911e-9cf42f1a1085/logs/94619026-912b-4d92-8f51-6c74f0692d90"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineRunLogContentPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTektonPipelineRunLogContent with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineRunLogContentOptions model
				getTektonPipelineRunLogContentOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineRunLogContentOptions)
				getTektonPipelineRunLogContentOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogContentOptionsModel.PipelineRunID = core.StringPtr("bf4b3abd-0c93-416b-911e-9cf42f1a1085")
				getTektonPipelineRunLogContentOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogContentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineRunLogContent(getTektonPipelineRunLogContentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineRunLogContent(getTektonPipelineRunLogContentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTektonPipelineRunLogContent(getTektonPipelineRunLogContentOptions *GetTektonPipelineRunLogContentOptions)`, func() {
		getTektonPipelineRunLogContentPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/pipeline_runs/bf4b3abd-0c93-416b-911e-9cf42f1a1085/logs/94619026-912b-4d92-8f51-6c74f0692d90"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineRunLogContentPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"data": "Data", "id": "ID"}`)
				}))
			})
			It(`Invoke GetTektonPipelineRunLogContent successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the GetTektonPipelineRunLogContentOptions model
				getTektonPipelineRunLogContentOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineRunLogContentOptions)
				getTektonPipelineRunLogContentOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogContentOptionsModel.PipelineRunID = core.StringPtr("bf4b3abd-0c93-416b-911e-9cf42f1a1085")
				getTektonPipelineRunLogContentOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogContentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.GetTektonPipelineRunLogContentWithContext(ctx, getTektonPipelineRunLogContentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineRunLogContent(getTektonPipelineRunLogContentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.GetTektonPipelineRunLogContentWithContext(ctx, getTektonPipelineRunLogContentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineRunLogContentPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"data": "Data", "id": "ID"}`)
				}))
			})
			It(`Invoke GetTektonPipelineRunLogContent successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineRunLogContent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTektonPipelineRunLogContentOptions model
				getTektonPipelineRunLogContentOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineRunLogContentOptions)
				getTektonPipelineRunLogContentOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogContentOptionsModel.PipelineRunID = core.StringPtr("bf4b3abd-0c93-416b-911e-9cf42f1a1085")
				getTektonPipelineRunLogContentOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogContentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineRunLogContent(getTektonPipelineRunLogContentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTektonPipelineRunLogContent with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineRunLogContentOptions model
				getTektonPipelineRunLogContentOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineRunLogContentOptions)
				getTektonPipelineRunLogContentOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogContentOptionsModel.PipelineRunID = core.StringPtr("bf4b3abd-0c93-416b-911e-9cf42f1a1085")
				getTektonPipelineRunLogContentOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogContentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineRunLogContent(getTektonPipelineRunLogContentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTektonPipelineRunLogContentOptions model with no property values
				getTektonPipelineRunLogContentOptionsModelNew := new(cdtektonpipelinev2.GetTektonPipelineRunLogContentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineRunLogContent(getTektonPipelineRunLogContentOptionsModelNew)
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
			It(`Invoke GetTektonPipelineRunLogContent successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineRunLogContentOptions model
				getTektonPipelineRunLogContentOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineRunLogContentOptions)
				getTektonPipelineRunLogContentOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogContentOptionsModel.PipelineRunID = core.StringPtr("bf4b3abd-0c93-416b-911e-9cf42f1a1085")
				getTektonPipelineRunLogContentOptionsModel.ID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogContentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineRunLogContent(getTektonPipelineRunLogContentOptionsModel)
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
	Describe(`ListTektonPipelineDefinitions(listTektonPipelineDefinitionsOptions *ListTektonPipelineDefinitionsOptions) - Operation response error`, func() {
		listTektonPipelineDefinitionsPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/definitions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTektonPipelineDefinitionsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTektonPipelineDefinitions with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ListTektonPipelineDefinitionsOptions model
				listTektonPipelineDefinitionsOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineDefinitionsOptions)
				listTektonPipelineDefinitionsOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineDefinitionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineDefinitions(listTektonPipelineDefinitionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.ListTektonPipelineDefinitions(listTektonPipelineDefinitionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTektonPipelineDefinitions(listTektonPipelineDefinitionsOptions *ListTektonPipelineDefinitionsOptions)`, func() {
		listTektonPipelineDefinitionsPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/definitions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTektonPipelineDefinitionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"definitions": [{"source": {"type": "Type", "properties": {"url": "URL", "branch": "Branch", "tag": "Tag", "path": "Path", "tool": {"id": "ID"}}}, "href": "Href", "id": "ID"}]}`)
				}))
			})
			It(`Invoke ListTektonPipelineDefinitions successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the ListTektonPipelineDefinitionsOptions model
				listTektonPipelineDefinitionsOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineDefinitionsOptions)
				listTektonPipelineDefinitionsOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineDefinitionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.ListTektonPipelineDefinitionsWithContext(ctx, listTektonPipelineDefinitionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineDefinitions(listTektonPipelineDefinitionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.ListTektonPipelineDefinitionsWithContext(ctx, listTektonPipelineDefinitionsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTektonPipelineDefinitionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"definitions": [{"source": {"type": "Type", "properties": {"url": "URL", "branch": "Branch", "tag": "Tag", "path": "Path", "tool": {"id": "ID"}}}, "href": "Href", "id": "ID"}]}`)
				}))
			})
			It(`Invoke ListTektonPipelineDefinitions successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineDefinitions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTektonPipelineDefinitionsOptions model
				listTektonPipelineDefinitionsOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineDefinitionsOptions)
				listTektonPipelineDefinitionsOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineDefinitionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.ListTektonPipelineDefinitions(listTektonPipelineDefinitionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTektonPipelineDefinitions with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ListTektonPipelineDefinitionsOptions model
				listTektonPipelineDefinitionsOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineDefinitionsOptions)
				listTektonPipelineDefinitionsOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineDefinitionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineDefinitions(listTektonPipelineDefinitionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTektonPipelineDefinitionsOptions model with no property values
				listTektonPipelineDefinitionsOptionsModelNew := new(cdtektonpipelinev2.ListTektonPipelineDefinitionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.ListTektonPipelineDefinitions(listTektonPipelineDefinitionsOptionsModelNew)
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
			It(`Invoke ListTektonPipelineDefinitions successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ListTektonPipelineDefinitionsOptions model
				listTektonPipelineDefinitionsOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineDefinitionsOptions)
				listTektonPipelineDefinitionsOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineDefinitionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineDefinitions(listTektonPipelineDefinitionsOptionsModel)
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
	Describe(`CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptions *CreateTektonPipelineDefinitionOptions) - Operation response error`, func() {
		createTektonPipelineDefinitionPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/definitions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelineDefinitionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTektonPipelineDefinition with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the Tool model
				toolModel := new(cdtektonpipelinev2.Tool)
				toolModel.ID = core.StringPtr("testString")

				// Construct an instance of the DefinitionSourceProperties model
				definitionSourcePropertiesModel := new(cdtektonpipelinev2.DefinitionSourceProperties)
				definitionSourcePropertiesModel.URL = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				definitionSourcePropertiesModel.Branch = core.StringPtr("master")
				definitionSourcePropertiesModel.Tag = core.StringPtr("testString")
				definitionSourcePropertiesModel.Path = core.StringPtr(".tekton")
				definitionSourcePropertiesModel.Tool = toolModel

				// Construct an instance of the DefinitionSource model
				definitionSourceModel := new(cdtektonpipelinev2.DefinitionSource)
				definitionSourceModel.Type = core.StringPtr("git")
				definitionSourceModel.Properties = definitionSourcePropertiesModel

				// Construct an instance of the CreateTektonPipelineDefinitionOptions model
				createTektonPipelineDefinitionOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineDefinitionOptions)
				createTektonPipelineDefinitionOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineDefinitionOptionsModel.Source = definitionSourceModel
				createTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptions *CreateTektonPipelineDefinitionOptions)`, func() {
		createTektonPipelineDefinitionPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/definitions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelineDefinitionPath))
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
					fmt.Fprintf(res, "%s", `{"source": {"type": "Type", "properties": {"url": "URL", "branch": "Branch", "tag": "Tag", "path": "Path", "tool": {"id": "ID"}}}, "href": "Href", "id": "ID"}`)
				}))
			})
			It(`Invoke CreateTektonPipelineDefinition successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the Tool model
				toolModel := new(cdtektonpipelinev2.Tool)
				toolModel.ID = core.StringPtr("testString")

				// Construct an instance of the DefinitionSourceProperties model
				definitionSourcePropertiesModel := new(cdtektonpipelinev2.DefinitionSourceProperties)
				definitionSourcePropertiesModel.URL = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				definitionSourcePropertiesModel.Branch = core.StringPtr("master")
				definitionSourcePropertiesModel.Tag = core.StringPtr("testString")
				definitionSourcePropertiesModel.Path = core.StringPtr(".tekton")
				definitionSourcePropertiesModel.Tool = toolModel

				// Construct an instance of the DefinitionSource model
				definitionSourceModel := new(cdtektonpipelinev2.DefinitionSource)
				definitionSourceModel.Type = core.StringPtr("git")
				definitionSourceModel.Properties = definitionSourcePropertiesModel

				// Construct an instance of the CreateTektonPipelineDefinitionOptions model
				createTektonPipelineDefinitionOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineDefinitionOptions)
				createTektonPipelineDefinitionOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineDefinitionOptionsModel.Source = definitionSourceModel
				createTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.CreateTektonPipelineDefinitionWithContext(ctx, createTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.CreateTektonPipelineDefinitionWithContext(ctx, createTektonPipelineDefinitionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelineDefinitionPath))
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
					fmt.Fprintf(res, "%s", `{"source": {"type": "Type", "properties": {"url": "URL", "branch": "Branch", "tag": "Tag", "path": "Path", "tool": {"id": "ID"}}}, "href": "Href", "id": "ID"}`)
				}))
			})
			It(`Invoke CreateTektonPipelineDefinition successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineDefinition(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Tool model
				toolModel := new(cdtektonpipelinev2.Tool)
				toolModel.ID = core.StringPtr("testString")

				// Construct an instance of the DefinitionSourceProperties model
				definitionSourcePropertiesModel := new(cdtektonpipelinev2.DefinitionSourceProperties)
				definitionSourcePropertiesModel.URL = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				definitionSourcePropertiesModel.Branch = core.StringPtr("master")
				definitionSourcePropertiesModel.Tag = core.StringPtr("testString")
				definitionSourcePropertiesModel.Path = core.StringPtr(".tekton")
				definitionSourcePropertiesModel.Tool = toolModel

				// Construct an instance of the DefinitionSource model
				definitionSourceModel := new(cdtektonpipelinev2.DefinitionSource)
				definitionSourceModel.Type = core.StringPtr("git")
				definitionSourceModel.Properties = definitionSourcePropertiesModel

				// Construct an instance of the CreateTektonPipelineDefinitionOptions model
				createTektonPipelineDefinitionOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineDefinitionOptions)
				createTektonPipelineDefinitionOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineDefinitionOptionsModel.Source = definitionSourceModel
				createTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTektonPipelineDefinition with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the Tool model
				toolModel := new(cdtektonpipelinev2.Tool)
				toolModel.ID = core.StringPtr("testString")

				// Construct an instance of the DefinitionSourceProperties model
				definitionSourcePropertiesModel := new(cdtektonpipelinev2.DefinitionSourceProperties)
				definitionSourcePropertiesModel.URL = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				definitionSourcePropertiesModel.Branch = core.StringPtr("master")
				definitionSourcePropertiesModel.Tag = core.StringPtr("testString")
				definitionSourcePropertiesModel.Path = core.StringPtr(".tekton")
				definitionSourcePropertiesModel.Tool = toolModel

				// Construct an instance of the DefinitionSource model
				definitionSourceModel := new(cdtektonpipelinev2.DefinitionSource)
				definitionSourceModel.Type = core.StringPtr("git")
				definitionSourceModel.Properties = definitionSourcePropertiesModel

				// Construct an instance of the CreateTektonPipelineDefinitionOptions model
				createTektonPipelineDefinitionOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineDefinitionOptions)
				createTektonPipelineDefinitionOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineDefinitionOptionsModel.Source = definitionSourceModel
				createTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTektonPipelineDefinitionOptions model with no property values
				createTektonPipelineDefinitionOptionsModelNew := new(cdtektonpipelinev2.CreateTektonPipelineDefinitionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptionsModelNew)
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
			It(`Invoke CreateTektonPipelineDefinition successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the Tool model
				toolModel := new(cdtektonpipelinev2.Tool)
				toolModel.ID = core.StringPtr("testString")

				// Construct an instance of the DefinitionSourceProperties model
				definitionSourcePropertiesModel := new(cdtektonpipelinev2.DefinitionSourceProperties)
				definitionSourcePropertiesModel.URL = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				definitionSourcePropertiesModel.Branch = core.StringPtr("master")
				definitionSourcePropertiesModel.Tag = core.StringPtr("testString")
				definitionSourcePropertiesModel.Path = core.StringPtr(".tekton")
				definitionSourcePropertiesModel.Tool = toolModel

				// Construct an instance of the DefinitionSource model
				definitionSourceModel := new(cdtektonpipelinev2.DefinitionSource)
				definitionSourceModel.Type = core.StringPtr("git")
				definitionSourceModel.Properties = definitionSourcePropertiesModel

				// Construct an instance of the CreateTektonPipelineDefinitionOptions model
				createTektonPipelineDefinitionOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineDefinitionOptions)
				createTektonPipelineDefinitionOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineDefinitionOptionsModel.Source = definitionSourceModel
				createTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptionsModel)
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
	Describe(`GetTektonPipelineDefinition(getTektonPipelineDefinitionOptions *GetTektonPipelineDefinitionOptions) - Operation response error`, func() {
		getTektonPipelineDefinitionPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/definitions/94299034-d45f-4e9a-8ed5-6bd5c7bb7ada"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineDefinitionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTektonPipelineDefinition with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineDefinitionOptions model
				getTektonPipelineDefinitionOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineDefinitionOptions)
				getTektonPipelineDefinitionOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineDefinitionOptionsModel.DefinitionID = core.StringPtr("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada")
				getTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTektonPipelineDefinition(getTektonPipelineDefinitionOptions *GetTektonPipelineDefinitionOptions)`, func() {
		getTektonPipelineDefinitionPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/definitions/94299034-d45f-4e9a-8ed5-6bd5c7bb7ada"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineDefinitionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"source": {"type": "Type", "properties": {"url": "URL", "branch": "Branch", "tag": "Tag", "path": "Path", "tool": {"id": "ID"}}}, "href": "Href", "id": "ID"}`)
				}))
			})
			It(`Invoke GetTektonPipelineDefinition successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the GetTektonPipelineDefinitionOptions model
				getTektonPipelineDefinitionOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineDefinitionOptions)
				getTektonPipelineDefinitionOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineDefinitionOptionsModel.DefinitionID = core.StringPtr("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada")
				getTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.GetTektonPipelineDefinitionWithContext(ctx, getTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.GetTektonPipelineDefinitionWithContext(ctx, getTektonPipelineDefinitionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineDefinitionPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"source": {"type": "Type", "properties": {"url": "URL", "branch": "Branch", "tag": "Tag", "path": "Path", "tool": {"id": "ID"}}}, "href": "Href", "id": "ID"}`)
				}))
			})
			It(`Invoke GetTektonPipelineDefinition successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineDefinition(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTektonPipelineDefinitionOptions model
				getTektonPipelineDefinitionOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineDefinitionOptions)
				getTektonPipelineDefinitionOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineDefinitionOptionsModel.DefinitionID = core.StringPtr("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada")
				getTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTektonPipelineDefinition with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineDefinitionOptions model
				getTektonPipelineDefinitionOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineDefinitionOptions)
				getTektonPipelineDefinitionOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineDefinitionOptionsModel.DefinitionID = core.StringPtr("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada")
				getTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTektonPipelineDefinitionOptions model with no property values
				getTektonPipelineDefinitionOptionsModelNew := new(cdtektonpipelinev2.GetTektonPipelineDefinitionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptionsModelNew)
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
			It(`Invoke GetTektonPipelineDefinition successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineDefinitionOptions model
				getTektonPipelineDefinitionOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineDefinitionOptions)
				getTektonPipelineDefinitionOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineDefinitionOptionsModel.DefinitionID = core.StringPtr("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada")
				getTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptionsModel)
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
	Describe(`ReplaceTektonPipelineDefinition(replaceTektonPipelineDefinitionOptions *ReplaceTektonPipelineDefinitionOptions) - Operation response error`, func() {
		replaceTektonPipelineDefinitionPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/definitions/94299034-d45f-4e9a-8ed5-6bd5c7bb7ada"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceTektonPipelineDefinitionPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceTektonPipelineDefinition with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the Tool model
				toolModel := new(cdtektonpipelinev2.Tool)
				toolModel.ID = core.StringPtr("testString")

				// Construct an instance of the DefinitionSourceProperties model
				definitionSourcePropertiesModel := new(cdtektonpipelinev2.DefinitionSourceProperties)
				definitionSourcePropertiesModel.URL = core.StringPtr("testString")
				definitionSourcePropertiesModel.Branch = core.StringPtr("testString")
				definitionSourcePropertiesModel.Tag = core.StringPtr("testString")
				definitionSourcePropertiesModel.Path = core.StringPtr("testString")
				definitionSourcePropertiesModel.Tool = toolModel

				// Construct an instance of the DefinitionSource model
				definitionSourceModel := new(cdtektonpipelinev2.DefinitionSource)
				definitionSourceModel.Type = core.StringPtr("testString")
				definitionSourceModel.Properties = definitionSourcePropertiesModel

				// Construct an instance of the ReplaceTektonPipelineDefinitionOptions model
				replaceTektonPipelineDefinitionOptionsModel := new(cdtektonpipelinev2.ReplaceTektonPipelineDefinitionOptions)
				replaceTektonPipelineDefinitionOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				replaceTektonPipelineDefinitionOptionsModel.DefinitionID = core.StringPtr("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada")
				replaceTektonPipelineDefinitionOptionsModel.Source = definitionSourceModel
				replaceTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.ReplaceTektonPipelineDefinition(replaceTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.ReplaceTektonPipelineDefinition(replaceTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceTektonPipelineDefinition(replaceTektonPipelineDefinitionOptions *ReplaceTektonPipelineDefinitionOptions)`, func() {
		replaceTektonPipelineDefinitionPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/definitions/94299034-d45f-4e9a-8ed5-6bd5c7bb7ada"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceTektonPipelineDefinitionPath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"source": {"type": "Type", "properties": {"url": "URL", "branch": "Branch", "tag": "Tag", "path": "Path", "tool": {"id": "ID"}}}, "href": "Href", "id": "ID"}`)
				}))
			})
			It(`Invoke ReplaceTektonPipelineDefinition successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the Tool model
				toolModel := new(cdtektonpipelinev2.Tool)
				toolModel.ID = core.StringPtr("testString")

				// Construct an instance of the DefinitionSourceProperties model
				definitionSourcePropertiesModel := new(cdtektonpipelinev2.DefinitionSourceProperties)
				definitionSourcePropertiesModel.URL = core.StringPtr("testString")
				definitionSourcePropertiesModel.Branch = core.StringPtr("testString")
				definitionSourcePropertiesModel.Tag = core.StringPtr("testString")
				definitionSourcePropertiesModel.Path = core.StringPtr("testString")
				definitionSourcePropertiesModel.Tool = toolModel

				// Construct an instance of the DefinitionSource model
				definitionSourceModel := new(cdtektonpipelinev2.DefinitionSource)
				definitionSourceModel.Type = core.StringPtr("testString")
				definitionSourceModel.Properties = definitionSourcePropertiesModel

				// Construct an instance of the ReplaceTektonPipelineDefinitionOptions model
				replaceTektonPipelineDefinitionOptionsModel := new(cdtektonpipelinev2.ReplaceTektonPipelineDefinitionOptions)
				replaceTektonPipelineDefinitionOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				replaceTektonPipelineDefinitionOptionsModel.DefinitionID = core.StringPtr("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada")
				replaceTektonPipelineDefinitionOptionsModel.Source = definitionSourceModel
				replaceTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.ReplaceTektonPipelineDefinitionWithContext(ctx, replaceTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.ReplaceTektonPipelineDefinition(replaceTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.ReplaceTektonPipelineDefinitionWithContext(ctx, replaceTektonPipelineDefinitionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceTektonPipelineDefinitionPath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"source": {"type": "Type", "properties": {"url": "URL", "branch": "Branch", "tag": "Tag", "path": "Path", "tool": {"id": "ID"}}}, "href": "Href", "id": "ID"}`)
				}))
			})
			It(`Invoke ReplaceTektonPipelineDefinition successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.ReplaceTektonPipelineDefinition(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Tool model
				toolModel := new(cdtektonpipelinev2.Tool)
				toolModel.ID = core.StringPtr("testString")

				// Construct an instance of the DefinitionSourceProperties model
				definitionSourcePropertiesModel := new(cdtektonpipelinev2.DefinitionSourceProperties)
				definitionSourcePropertiesModel.URL = core.StringPtr("testString")
				definitionSourcePropertiesModel.Branch = core.StringPtr("testString")
				definitionSourcePropertiesModel.Tag = core.StringPtr("testString")
				definitionSourcePropertiesModel.Path = core.StringPtr("testString")
				definitionSourcePropertiesModel.Tool = toolModel

				// Construct an instance of the DefinitionSource model
				definitionSourceModel := new(cdtektonpipelinev2.DefinitionSource)
				definitionSourceModel.Type = core.StringPtr("testString")
				definitionSourceModel.Properties = definitionSourcePropertiesModel

				// Construct an instance of the ReplaceTektonPipelineDefinitionOptions model
				replaceTektonPipelineDefinitionOptionsModel := new(cdtektonpipelinev2.ReplaceTektonPipelineDefinitionOptions)
				replaceTektonPipelineDefinitionOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				replaceTektonPipelineDefinitionOptionsModel.DefinitionID = core.StringPtr("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada")
				replaceTektonPipelineDefinitionOptionsModel.Source = definitionSourceModel
				replaceTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.ReplaceTektonPipelineDefinition(replaceTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceTektonPipelineDefinition with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the Tool model
				toolModel := new(cdtektonpipelinev2.Tool)
				toolModel.ID = core.StringPtr("testString")

				// Construct an instance of the DefinitionSourceProperties model
				definitionSourcePropertiesModel := new(cdtektonpipelinev2.DefinitionSourceProperties)
				definitionSourcePropertiesModel.URL = core.StringPtr("testString")
				definitionSourcePropertiesModel.Branch = core.StringPtr("testString")
				definitionSourcePropertiesModel.Tag = core.StringPtr("testString")
				definitionSourcePropertiesModel.Path = core.StringPtr("testString")
				definitionSourcePropertiesModel.Tool = toolModel

				// Construct an instance of the DefinitionSource model
				definitionSourceModel := new(cdtektonpipelinev2.DefinitionSource)
				definitionSourceModel.Type = core.StringPtr("testString")
				definitionSourceModel.Properties = definitionSourcePropertiesModel

				// Construct an instance of the ReplaceTektonPipelineDefinitionOptions model
				replaceTektonPipelineDefinitionOptionsModel := new(cdtektonpipelinev2.ReplaceTektonPipelineDefinitionOptions)
				replaceTektonPipelineDefinitionOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				replaceTektonPipelineDefinitionOptionsModel.DefinitionID = core.StringPtr("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada")
				replaceTektonPipelineDefinitionOptionsModel.Source = definitionSourceModel
				replaceTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.ReplaceTektonPipelineDefinition(replaceTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceTektonPipelineDefinitionOptions model with no property values
				replaceTektonPipelineDefinitionOptionsModelNew := new(cdtektonpipelinev2.ReplaceTektonPipelineDefinitionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.ReplaceTektonPipelineDefinition(replaceTektonPipelineDefinitionOptionsModelNew)
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
			It(`Invoke ReplaceTektonPipelineDefinition successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the Tool model
				toolModel := new(cdtektonpipelinev2.Tool)
				toolModel.ID = core.StringPtr("testString")

				// Construct an instance of the DefinitionSourceProperties model
				definitionSourcePropertiesModel := new(cdtektonpipelinev2.DefinitionSourceProperties)
				definitionSourcePropertiesModel.URL = core.StringPtr("testString")
				definitionSourcePropertiesModel.Branch = core.StringPtr("testString")
				definitionSourcePropertiesModel.Tag = core.StringPtr("testString")
				definitionSourcePropertiesModel.Path = core.StringPtr("testString")
				definitionSourcePropertiesModel.Tool = toolModel

				// Construct an instance of the DefinitionSource model
				definitionSourceModel := new(cdtektonpipelinev2.DefinitionSource)
				definitionSourceModel.Type = core.StringPtr("testString")
				definitionSourceModel.Properties = definitionSourcePropertiesModel

				// Construct an instance of the ReplaceTektonPipelineDefinitionOptions model
				replaceTektonPipelineDefinitionOptionsModel := new(cdtektonpipelinev2.ReplaceTektonPipelineDefinitionOptions)
				replaceTektonPipelineDefinitionOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				replaceTektonPipelineDefinitionOptionsModel.DefinitionID = core.StringPtr("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada")
				replaceTektonPipelineDefinitionOptionsModel.Source = definitionSourceModel
				replaceTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.ReplaceTektonPipelineDefinition(replaceTektonPipelineDefinitionOptionsModel)
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
	Describe(`DeleteTektonPipelineDefinition(deleteTektonPipelineDefinitionOptions *DeleteTektonPipelineDefinitionOptions)`, func() {
		deleteTektonPipelineDefinitionPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/definitions/94299034-d45f-4e9a-8ed5-6bd5c7bb7ada"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTektonPipelineDefinitionPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTektonPipelineDefinition successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := cdTektonPipelineService.DeleteTektonPipelineDefinition(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTektonPipelineDefinitionOptions model
				deleteTektonPipelineDefinitionOptionsModel := new(cdtektonpipelinev2.DeleteTektonPipelineDefinitionOptions)
				deleteTektonPipelineDefinitionOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelineDefinitionOptionsModel.DefinitionID = core.StringPtr("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada")
				deleteTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = cdTektonPipelineService.DeleteTektonPipelineDefinition(deleteTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTektonPipelineDefinition with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the DeleteTektonPipelineDefinitionOptions model
				deleteTektonPipelineDefinitionOptionsModel := new(cdtektonpipelinev2.DeleteTektonPipelineDefinitionOptions)
				deleteTektonPipelineDefinitionOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelineDefinitionOptionsModel.DefinitionID = core.StringPtr("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada")
				deleteTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := cdTektonPipelineService.DeleteTektonPipelineDefinition(deleteTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTektonPipelineDefinitionOptions model with no property values
				deleteTektonPipelineDefinitionOptionsModelNew := new(cdtektonpipelinev2.DeleteTektonPipelineDefinitionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = cdTektonPipelineService.DeleteTektonPipelineDefinition(deleteTektonPipelineDefinitionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTektonPipelineProperties(listTektonPipelinePropertiesOptions *ListTektonPipelinePropertiesOptions) - Operation response error`, func() {
		listTektonPipelinePropertiesPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/properties"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTektonPipelinePropertiesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"prod"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"name"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTektonPipelineProperties with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ListTektonPipelinePropertiesOptions model
				listTektonPipelinePropertiesOptionsModel := new(cdtektonpipelinev2.ListTektonPipelinePropertiesOptions)
				listTektonPipelinePropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelinePropertiesOptionsModel.Name = core.StringPtr("prod")
				listTektonPipelinePropertiesOptionsModel.Type = []string{"secure", "text"}
				listTektonPipelinePropertiesOptionsModel.Sort = core.StringPtr("name")
				listTektonPipelinePropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineProperties(listTektonPipelinePropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.ListTektonPipelineProperties(listTektonPipelinePropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTektonPipelineProperties(listTektonPipelinePropertiesOptions *ListTektonPipelinePropertiesOptions)`, func() {
		listTektonPipelinePropertiesPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/properties"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTektonPipelinePropertiesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"prod"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"name"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}]}`)
				}))
			})
			It(`Invoke ListTektonPipelineProperties successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the ListTektonPipelinePropertiesOptions model
				listTektonPipelinePropertiesOptionsModel := new(cdtektonpipelinev2.ListTektonPipelinePropertiesOptions)
				listTektonPipelinePropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelinePropertiesOptionsModel.Name = core.StringPtr("prod")
				listTektonPipelinePropertiesOptionsModel.Type = []string{"secure", "text"}
				listTektonPipelinePropertiesOptionsModel.Sort = core.StringPtr("name")
				listTektonPipelinePropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.ListTektonPipelinePropertiesWithContext(ctx, listTektonPipelinePropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineProperties(listTektonPipelinePropertiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.ListTektonPipelinePropertiesWithContext(ctx, listTektonPipelinePropertiesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTektonPipelinePropertiesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"prod"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"name"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}]}`)
				}))
			})
			It(`Invoke ListTektonPipelineProperties successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineProperties(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTektonPipelinePropertiesOptions model
				listTektonPipelinePropertiesOptionsModel := new(cdtektonpipelinev2.ListTektonPipelinePropertiesOptions)
				listTektonPipelinePropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelinePropertiesOptionsModel.Name = core.StringPtr("prod")
				listTektonPipelinePropertiesOptionsModel.Type = []string{"secure", "text"}
				listTektonPipelinePropertiesOptionsModel.Sort = core.StringPtr("name")
				listTektonPipelinePropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.ListTektonPipelineProperties(listTektonPipelinePropertiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTektonPipelineProperties with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ListTektonPipelinePropertiesOptions model
				listTektonPipelinePropertiesOptionsModel := new(cdtektonpipelinev2.ListTektonPipelinePropertiesOptions)
				listTektonPipelinePropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelinePropertiesOptionsModel.Name = core.StringPtr("prod")
				listTektonPipelinePropertiesOptionsModel.Type = []string{"secure", "text"}
				listTektonPipelinePropertiesOptionsModel.Sort = core.StringPtr("name")
				listTektonPipelinePropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineProperties(listTektonPipelinePropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTektonPipelinePropertiesOptions model with no property values
				listTektonPipelinePropertiesOptionsModelNew := new(cdtektonpipelinev2.ListTektonPipelinePropertiesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.ListTektonPipelineProperties(listTektonPipelinePropertiesOptionsModelNew)
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
			It(`Invoke ListTektonPipelineProperties successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ListTektonPipelinePropertiesOptions model
				listTektonPipelinePropertiesOptionsModel := new(cdtektonpipelinev2.ListTektonPipelinePropertiesOptions)
				listTektonPipelinePropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelinePropertiesOptionsModel.Name = core.StringPtr("prod")
				listTektonPipelinePropertiesOptionsModel.Type = []string{"secure", "text"}
				listTektonPipelinePropertiesOptionsModel.Sort = core.StringPtr("name")
				listTektonPipelinePropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineProperties(listTektonPipelinePropertiesOptionsModel)
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
	Describe(`CreateTektonPipelineProperties(createTektonPipelinePropertiesOptions *CreateTektonPipelinePropertiesOptions) - Operation response error`, func() {
		createTektonPipelinePropertiesPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/properties"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelinePropertiesPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTektonPipelineProperties with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the CreateTektonPipelinePropertiesOptions model
				createTektonPipelinePropertiesOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelinePropertiesOptions)
				createTektonPipelinePropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelinePropertiesOptionsModel.Name = core.StringPtr("prop1")
				createTektonPipelinePropertiesOptionsModel.Type = core.StringPtr("text")
				createTektonPipelinePropertiesOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				createTektonPipelinePropertiesOptionsModel.Enum = []string{"testString"}
				createTektonPipelinePropertiesOptionsModel.Locked = core.BoolPtr(false)
				createTektonPipelinePropertiesOptionsModel.Path = core.StringPtr("testString")
				createTektonPipelinePropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineProperties(createTektonPipelinePropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.CreateTektonPipelineProperties(createTektonPipelinePropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTektonPipelineProperties(createTektonPipelinePropertiesOptions *CreateTektonPipelinePropertiesOptions)`, func() {
		createTektonPipelinePropertiesPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/properties"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelinePropertiesPath))
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}`)
				}))
			})
			It(`Invoke CreateTektonPipelineProperties successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the CreateTektonPipelinePropertiesOptions model
				createTektonPipelinePropertiesOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelinePropertiesOptions)
				createTektonPipelinePropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelinePropertiesOptionsModel.Name = core.StringPtr("prop1")
				createTektonPipelinePropertiesOptionsModel.Type = core.StringPtr("text")
				createTektonPipelinePropertiesOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				createTektonPipelinePropertiesOptionsModel.Enum = []string{"testString"}
				createTektonPipelinePropertiesOptionsModel.Locked = core.BoolPtr(false)
				createTektonPipelinePropertiesOptionsModel.Path = core.StringPtr("testString")
				createTektonPipelinePropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.CreateTektonPipelinePropertiesWithContext(ctx, createTektonPipelinePropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineProperties(createTektonPipelinePropertiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.CreateTektonPipelinePropertiesWithContext(ctx, createTektonPipelinePropertiesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelinePropertiesPath))
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}`)
				}))
			})
			It(`Invoke CreateTektonPipelineProperties successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineProperties(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateTektonPipelinePropertiesOptions model
				createTektonPipelinePropertiesOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelinePropertiesOptions)
				createTektonPipelinePropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelinePropertiesOptionsModel.Name = core.StringPtr("prop1")
				createTektonPipelinePropertiesOptionsModel.Type = core.StringPtr("text")
				createTektonPipelinePropertiesOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				createTektonPipelinePropertiesOptionsModel.Enum = []string{"testString"}
				createTektonPipelinePropertiesOptionsModel.Locked = core.BoolPtr(false)
				createTektonPipelinePropertiesOptionsModel.Path = core.StringPtr("testString")
				createTektonPipelinePropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.CreateTektonPipelineProperties(createTektonPipelinePropertiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTektonPipelineProperties with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the CreateTektonPipelinePropertiesOptions model
				createTektonPipelinePropertiesOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelinePropertiesOptions)
				createTektonPipelinePropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelinePropertiesOptionsModel.Name = core.StringPtr("prop1")
				createTektonPipelinePropertiesOptionsModel.Type = core.StringPtr("text")
				createTektonPipelinePropertiesOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				createTektonPipelinePropertiesOptionsModel.Enum = []string{"testString"}
				createTektonPipelinePropertiesOptionsModel.Locked = core.BoolPtr(false)
				createTektonPipelinePropertiesOptionsModel.Path = core.StringPtr("testString")
				createTektonPipelinePropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineProperties(createTektonPipelinePropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTektonPipelinePropertiesOptions model with no property values
				createTektonPipelinePropertiesOptionsModelNew := new(cdtektonpipelinev2.CreateTektonPipelinePropertiesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.CreateTektonPipelineProperties(createTektonPipelinePropertiesOptionsModelNew)
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
			It(`Invoke CreateTektonPipelineProperties successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the CreateTektonPipelinePropertiesOptions model
				createTektonPipelinePropertiesOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelinePropertiesOptions)
				createTektonPipelinePropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelinePropertiesOptionsModel.Name = core.StringPtr("prop1")
				createTektonPipelinePropertiesOptionsModel.Type = core.StringPtr("text")
				createTektonPipelinePropertiesOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				createTektonPipelinePropertiesOptionsModel.Enum = []string{"testString"}
				createTektonPipelinePropertiesOptionsModel.Locked = core.BoolPtr(false)
				createTektonPipelinePropertiesOptionsModel.Path = core.StringPtr("testString")
				createTektonPipelinePropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineProperties(createTektonPipelinePropertiesOptionsModel)
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
	Describe(`GetTektonPipelineProperty(getTektonPipelinePropertyOptions *GetTektonPipelinePropertyOptions) - Operation response error`, func() {
		getTektonPipelinePropertyPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/properties/debug-pipeline"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelinePropertyPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTektonPipelineProperty with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelinePropertyOptions model
				getTektonPipelinePropertyOptionsModel := new(cdtektonpipelinev2.GetTektonPipelinePropertyOptions)
				getTektonPipelinePropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelinePropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				getTektonPipelinePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineProperty(getTektonPipelinePropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineProperty(getTektonPipelinePropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTektonPipelineProperty(getTektonPipelinePropertyOptions *GetTektonPipelinePropertyOptions)`, func() {
		getTektonPipelinePropertyPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/properties/debug-pipeline"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelinePropertyPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}`)
				}))
			})
			It(`Invoke GetTektonPipelineProperty successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the GetTektonPipelinePropertyOptions model
				getTektonPipelinePropertyOptionsModel := new(cdtektonpipelinev2.GetTektonPipelinePropertyOptions)
				getTektonPipelinePropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelinePropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				getTektonPipelinePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.GetTektonPipelinePropertyWithContext(ctx, getTektonPipelinePropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineProperty(getTektonPipelinePropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.GetTektonPipelinePropertyWithContext(ctx, getTektonPipelinePropertyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelinePropertyPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}`)
				}))
			})
			It(`Invoke GetTektonPipelineProperty successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineProperty(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTektonPipelinePropertyOptions model
				getTektonPipelinePropertyOptionsModel := new(cdtektonpipelinev2.GetTektonPipelinePropertyOptions)
				getTektonPipelinePropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelinePropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				getTektonPipelinePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineProperty(getTektonPipelinePropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTektonPipelineProperty with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelinePropertyOptions model
				getTektonPipelinePropertyOptionsModel := new(cdtektonpipelinev2.GetTektonPipelinePropertyOptions)
				getTektonPipelinePropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelinePropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				getTektonPipelinePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineProperty(getTektonPipelinePropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTektonPipelinePropertyOptions model with no property values
				getTektonPipelinePropertyOptionsModelNew := new(cdtektonpipelinev2.GetTektonPipelinePropertyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineProperty(getTektonPipelinePropertyOptionsModelNew)
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
			It(`Invoke GetTektonPipelineProperty successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelinePropertyOptions model
				getTektonPipelinePropertyOptionsModel := new(cdtektonpipelinev2.GetTektonPipelinePropertyOptions)
				getTektonPipelinePropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelinePropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				getTektonPipelinePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineProperty(getTektonPipelinePropertyOptionsModel)
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
	Describe(`ReplaceTektonPipelineProperty(replaceTektonPipelinePropertyOptions *ReplaceTektonPipelinePropertyOptions) - Operation response error`, func() {
		replaceTektonPipelinePropertyPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/properties/debug-pipeline"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceTektonPipelinePropertyPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceTektonPipelineProperty with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ReplaceTektonPipelinePropertyOptions model
				replaceTektonPipelinePropertyOptionsModel := new(cdtektonpipelinev2.ReplaceTektonPipelinePropertyOptions)
				replaceTektonPipelinePropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				replaceTektonPipelinePropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				replaceTektonPipelinePropertyOptionsModel.Name = core.StringPtr("prop1")
				replaceTektonPipelinePropertyOptionsModel.Type = core.StringPtr("text")
				replaceTektonPipelinePropertyOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				replaceTektonPipelinePropertyOptionsModel.Enum = []string{"testString"}
				replaceTektonPipelinePropertyOptionsModel.Locked = core.BoolPtr(false)
				replaceTektonPipelinePropertyOptionsModel.Path = core.StringPtr("testString")
				replaceTektonPipelinePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.ReplaceTektonPipelineProperty(replaceTektonPipelinePropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.ReplaceTektonPipelineProperty(replaceTektonPipelinePropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceTektonPipelineProperty(replaceTektonPipelinePropertyOptions *ReplaceTektonPipelinePropertyOptions)`, func() {
		replaceTektonPipelinePropertyPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/properties/debug-pipeline"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceTektonPipelinePropertyPath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}`)
				}))
			})
			It(`Invoke ReplaceTektonPipelineProperty successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceTektonPipelinePropertyOptions model
				replaceTektonPipelinePropertyOptionsModel := new(cdtektonpipelinev2.ReplaceTektonPipelinePropertyOptions)
				replaceTektonPipelinePropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				replaceTektonPipelinePropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				replaceTektonPipelinePropertyOptionsModel.Name = core.StringPtr("prop1")
				replaceTektonPipelinePropertyOptionsModel.Type = core.StringPtr("text")
				replaceTektonPipelinePropertyOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				replaceTektonPipelinePropertyOptionsModel.Enum = []string{"testString"}
				replaceTektonPipelinePropertyOptionsModel.Locked = core.BoolPtr(false)
				replaceTektonPipelinePropertyOptionsModel.Path = core.StringPtr("testString")
				replaceTektonPipelinePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.ReplaceTektonPipelinePropertyWithContext(ctx, replaceTektonPipelinePropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.ReplaceTektonPipelineProperty(replaceTektonPipelinePropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.ReplaceTektonPipelinePropertyWithContext(ctx, replaceTektonPipelinePropertyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceTektonPipelinePropertyPath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "locked": true, "path": "Path"}`)
				}))
			})
			It(`Invoke ReplaceTektonPipelineProperty successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.ReplaceTektonPipelineProperty(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceTektonPipelinePropertyOptions model
				replaceTektonPipelinePropertyOptionsModel := new(cdtektonpipelinev2.ReplaceTektonPipelinePropertyOptions)
				replaceTektonPipelinePropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				replaceTektonPipelinePropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				replaceTektonPipelinePropertyOptionsModel.Name = core.StringPtr("prop1")
				replaceTektonPipelinePropertyOptionsModel.Type = core.StringPtr("text")
				replaceTektonPipelinePropertyOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				replaceTektonPipelinePropertyOptionsModel.Enum = []string{"testString"}
				replaceTektonPipelinePropertyOptionsModel.Locked = core.BoolPtr(false)
				replaceTektonPipelinePropertyOptionsModel.Path = core.StringPtr("testString")
				replaceTektonPipelinePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.ReplaceTektonPipelineProperty(replaceTektonPipelinePropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceTektonPipelineProperty with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ReplaceTektonPipelinePropertyOptions model
				replaceTektonPipelinePropertyOptionsModel := new(cdtektonpipelinev2.ReplaceTektonPipelinePropertyOptions)
				replaceTektonPipelinePropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				replaceTektonPipelinePropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				replaceTektonPipelinePropertyOptionsModel.Name = core.StringPtr("prop1")
				replaceTektonPipelinePropertyOptionsModel.Type = core.StringPtr("text")
				replaceTektonPipelinePropertyOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				replaceTektonPipelinePropertyOptionsModel.Enum = []string{"testString"}
				replaceTektonPipelinePropertyOptionsModel.Locked = core.BoolPtr(false)
				replaceTektonPipelinePropertyOptionsModel.Path = core.StringPtr("testString")
				replaceTektonPipelinePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.ReplaceTektonPipelineProperty(replaceTektonPipelinePropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceTektonPipelinePropertyOptions model with no property values
				replaceTektonPipelinePropertyOptionsModelNew := new(cdtektonpipelinev2.ReplaceTektonPipelinePropertyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.ReplaceTektonPipelineProperty(replaceTektonPipelinePropertyOptionsModelNew)
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
			It(`Invoke ReplaceTektonPipelineProperty successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ReplaceTektonPipelinePropertyOptions model
				replaceTektonPipelinePropertyOptionsModel := new(cdtektonpipelinev2.ReplaceTektonPipelinePropertyOptions)
				replaceTektonPipelinePropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				replaceTektonPipelinePropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				replaceTektonPipelinePropertyOptionsModel.Name = core.StringPtr("prop1")
				replaceTektonPipelinePropertyOptionsModel.Type = core.StringPtr("text")
				replaceTektonPipelinePropertyOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				replaceTektonPipelinePropertyOptionsModel.Enum = []string{"testString"}
				replaceTektonPipelinePropertyOptionsModel.Locked = core.BoolPtr(false)
				replaceTektonPipelinePropertyOptionsModel.Path = core.StringPtr("testString")
				replaceTektonPipelinePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.ReplaceTektonPipelineProperty(replaceTektonPipelinePropertyOptionsModel)
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
	Describe(`DeleteTektonPipelineProperty(deleteTektonPipelinePropertyOptions *DeleteTektonPipelinePropertyOptions)`, func() {
		deleteTektonPipelinePropertyPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/properties/debug-pipeline"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTektonPipelinePropertyPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTektonPipelineProperty successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := cdTektonPipelineService.DeleteTektonPipelineProperty(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTektonPipelinePropertyOptions model
				deleteTektonPipelinePropertyOptionsModel := new(cdtektonpipelinev2.DeleteTektonPipelinePropertyOptions)
				deleteTektonPipelinePropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelinePropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				deleteTektonPipelinePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = cdTektonPipelineService.DeleteTektonPipelineProperty(deleteTektonPipelinePropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTektonPipelineProperty with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the DeleteTektonPipelinePropertyOptions model
				deleteTektonPipelinePropertyOptionsModel := new(cdtektonpipelinev2.DeleteTektonPipelinePropertyOptions)
				deleteTektonPipelinePropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelinePropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				deleteTektonPipelinePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := cdTektonPipelineService.DeleteTektonPipelineProperty(deleteTektonPipelinePropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTektonPipelinePropertyOptions model with no property values
				deleteTektonPipelinePropertyOptionsModelNew := new(cdtektonpipelinev2.DeleteTektonPipelinePropertyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = cdTektonPipelineService.DeleteTektonPipelineProperty(deleteTektonPipelinePropertyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTektonPipelineTriggers(listTektonPipelineTriggersOptions *ListTektonPipelineTriggersOptions) - Operation response error`, func() {
		listTektonPipelineTriggersPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTektonPipelineTriggersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"manual,scm"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["event_listener"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["worker.id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["worker.name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["disabled"]).To(Equal([]string{"true"}))
					Expect(req.URL.Query()["tags"]).To(Equal([]string{"tag1,tag2"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTektonPipelineTriggers with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ListTektonPipelineTriggersOptions model
				listTektonPipelineTriggersOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineTriggersOptions)
				listTektonPipelineTriggersOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineTriggersOptionsModel.Type = core.StringPtr("manual,scm")
				listTektonPipelineTriggersOptionsModel.Name = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.EventListener = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.WorkerID = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.WorkerName = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.Disabled = core.StringPtr("true")
				listTektonPipelineTriggersOptionsModel.Tags = core.StringPtr("tag1,tag2")
				listTektonPipelineTriggersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineTriggers(listTektonPipelineTriggersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.ListTektonPipelineTriggers(listTektonPipelineTriggersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTektonPipelineTriggers(listTektonPipelineTriggersOptions *ListTektonPipelineTriggersOptions)`, func() {
		listTektonPipelineTriggersPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTektonPipelineTriggersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"manual,scm"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["event_listener"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["worker.id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["worker.name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["disabled"]).To(Equal([]string{"true"}))
					Expect(req.URL.Query()["tags"]).To(Equal([]string{"tag1,tag2"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"triggers": [{"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}]}`)
				}))
			})
			It(`Invoke ListTektonPipelineTriggers successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the ListTektonPipelineTriggersOptions model
				listTektonPipelineTriggersOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineTriggersOptions)
				listTektonPipelineTriggersOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineTriggersOptionsModel.Type = core.StringPtr("manual,scm")
				listTektonPipelineTriggersOptionsModel.Name = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.EventListener = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.WorkerID = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.WorkerName = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.Disabled = core.StringPtr("true")
				listTektonPipelineTriggersOptionsModel.Tags = core.StringPtr("tag1,tag2")
				listTektonPipelineTriggersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.ListTektonPipelineTriggersWithContext(ctx, listTektonPipelineTriggersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineTriggers(listTektonPipelineTriggersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.ListTektonPipelineTriggersWithContext(ctx, listTektonPipelineTriggersOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTektonPipelineTriggersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"manual,scm"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["event_listener"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["worker.id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["worker.name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["disabled"]).To(Equal([]string{"true"}))
					Expect(req.URL.Query()["tags"]).To(Equal([]string{"tag1,tag2"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"triggers": [{"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}]}`)
				}))
			})
			It(`Invoke ListTektonPipelineTriggers successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineTriggers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTektonPipelineTriggersOptions model
				listTektonPipelineTriggersOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineTriggersOptions)
				listTektonPipelineTriggersOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineTriggersOptionsModel.Type = core.StringPtr("manual,scm")
				listTektonPipelineTriggersOptionsModel.Name = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.EventListener = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.WorkerID = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.WorkerName = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.Disabled = core.StringPtr("true")
				listTektonPipelineTriggersOptionsModel.Tags = core.StringPtr("tag1,tag2")
				listTektonPipelineTriggersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.ListTektonPipelineTriggers(listTektonPipelineTriggersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTektonPipelineTriggers with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ListTektonPipelineTriggersOptions model
				listTektonPipelineTriggersOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineTriggersOptions)
				listTektonPipelineTriggersOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineTriggersOptionsModel.Type = core.StringPtr("manual,scm")
				listTektonPipelineTriggersOptionsModel.Name = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.EventListener = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.WorkerID = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.WorkerName = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.Disabled = core.StringPtr("true")
				listTektonPipelineTriggersOptionsModel.Tags = core.StringPtr("tag1,tag2")
				listTektonPipelineTriggersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineTriggers(listTektonPipelineTriggersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTektonPipelineTriggersOptions model with no property values
				listTektonPipelineTriggersOptionsModelNew := new(cdtektonpipelinev2.ListTektonPipelineTriggersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.ListTektonPipelineTriggers(listTektonPipelineTriggersOptionsModelNew)
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
			It(`Invoke ListTektonPipelineTriggers successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ListTektonPipelineTriggersOptions model
				listTektonPipelineTriggersOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineTriggersOptions)
				listTektonPipelineTriggersOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineTriggersOptionsModel.Type = core.StringPtr("manual,scm")
				listTektonPipelineTriggersOptionsModel.Name = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.EventListener = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.WorkerID = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.WorkerName = core.StringPtr("testString")
				listTektonPipelineTriggersOptionsModel.Disabled = core.StringPtr("true")
				listTektonPipelineTriggersOptionsModel.Tags = core.StringPtr("tag1,tag2")
				listTektonPipelineTriggersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineTriggers(listTektonPipelineTriggersOptionsModel)
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
	Describe(`CreateTektonPipelineTrigger(createTektonPipelineTriggerOptions *CreateTektonPipelineTriggerOptions) - Operation response error`, func() {
		createTektonPipelineTriggerPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelineTriggerPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTektonPipelineTrigger with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("public")

				// Construct an instance of the GenericSecret model
				genericSecretModel := new(cdtektonpipelinev2.GenericSecret)
				genericSecretModel.Type = core.StringPtr("token_matches")
				genericSecretModel.Value = core.StringPtr("testString")
				genericSecretModel.Source = core.StringPtr("header")
				genericSecretModel.KeyName = core.StringPtr("testString")
				genericSecretModel.Algorithm = core.StringPtr("md4")

				// Construct an instance of the TriggerSourcePropertiesPrototype model
				triggerSourcePropertiesPrototypeModel := new(cdtektonpipelinev2.TriggerSourcePropertiesPrototype)
				triggerSourcePropertiesPrototypeModel.URL = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Branch = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Pattern = core.StringPtr("testString")

				// Construct an instance of the TriggerSourcePrototype model
				triggerSourcePrototypeModel := new(cdtektonpipelinev2.TriggerSourcePrototype)
				triggerSourcePrototypeModel.Type = core.StringPtr("testString")
				triggerSourcePrototypeModel.Properties = triggerSourcePropertiesPrototypeModel

				// Construct an instance of the CreateTektonPipelineTriggerOptions model
				createTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineTriggerOptions)
				createTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineTriggerOptionsModel.Type = core.StringPtr("manual")
				createTektonPipelineTriggerOptionsModel.Name = core.StringPtr("Manual Trigger")
				createTektonPipelineTriggerOptionsModel.EventListener = core.StringPtr("pr-listener")
				createTektonPipelineTriggerOptionsModel.Tags = []string{"testString"}
				createTektonPipelineTriggerOptionsModel.Worker = workerIdentityModel
				createTektonPipelineTriggerOptionsModel.MaxConcurrentRuns = core.Int64Ptr(int64(3))
				createTektonPipelineTriggerOptionsModel.LimitWaitingRuns = core.BoolPtr(false)
				createTektonPipelineTriggerOptionsModel.Enabled = core.BoolPtr(true)
				createTektonPipelineTriggerOptionsModel.Secret = genericSecretModel
				createTektonPipelineTriggerOptionsModel.Cron = core.StringPtr("testString")
				createTektonPipelineTriggerOptionsModel.Timezone = core.StringPtr("testString")
				createTektonPipelineTriggerOptionsModel.Source = triggerSourcePrototypeModel
				createTektonPipelineTriggerOptionsModel.Events = []string{"push"}
				createTektonPipelineTriggerOptionsModel.Filter = core.StringPtr("testString")
				createTektonPipelineTriggerOptionsModel.Favorite = core.BoolPtr(false)
				createTektonPipelineTriggerOptionsModel.EnableEventsFromForks = core.BoolPtr(false)
				createTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineTrigger(createTektonPipelineTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.CreateTektonPipelineTrigger(createTektonPipelineTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTektonPipelineTrigger(createTektonPipelineTriggerOptions *CreateTektonPipelineTriggerOptions)`, func() {
		createTektonPipelineTriggerPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelineTriggerPath))
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
					fmt.Fprintf(res, "%s", `{"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}`)
				}))
			})
			It(`Invoke CreateTektonPipelineTrigger successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("public")

				// Construct an instance of the GenericSecret model
				genericSecretModel := new(cdtektonpipelinev2.GenericSecret)
				genericSecretModel.Type = core.StringPtr("token_matches")
				genericSecretModel.Value = core.StringPtr("testString")
				genericSecretModel.Source = core.StringPtr("header")
				genericSecretModel.KeyName = core.StringPtr("testString")
				genericSecretModel.Algorithm = core.StringPtr("md4")

				// Construct an instance of the TriggerSourcePropertiesPrototype model
				triggerSourcePropertiesPrototypeModel := new(cdtektonpipelinev2.TriggerSourcePropertiesPrototype)
				triggerSourcePropertiesPrototypeModel.URL = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Branch = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Pattern = core.StringPtr("testString")

				// Construct an instance of the TriggerSourcePrototype model
				triggerSourcePrototypeModel := new(cdtektonpipelinev2.TriggerSourcePrototype)
				triggerSourcePrototypeModel.Type = core.StringPtr("testString")
				triggerSourcePrototypeModel.Properties = triggerSourcePropertiesPrototypeModel

				// Construct an instance of the CreateTektonPipelineTriggerOptions model
				createTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineTriggerOptions)
				createTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineTriggerOptionsModel.Type = core.StringPtr("manual")
				createTektonPipelineTriggerOptionsModel.Name = core.StringPtr("Manual Trigger")
				createTektonPipelineTriggerOptionsModel.EventListener = core.StringPtr("pr-listener")
				createTektonPipelineTriggerOptionsModel.Tags = []string{"testString"}
				createTektonPipelineTriggerOptionsModel.Worker = workerIdentityModel
				createTektonPipelineTriggerOptionsModel.MaxConcurrentRuns = core.Int64Ptr(int64(3))
				createTektonPipelineTriggerOptionsModel.LimitWaitingRuns = core.BoolPtr(false)
				createTektonPipelineTriggerOptionsModel.Enabled = core.BoolPtr(true)
				createTektonPipelineTriggerOptionsModel.Secret = genericSecretModel
				createTektonPipelineTriggerOptionsModel.Cron = core.StringPtr("testString")
				createTektonPipelineTriggerOptionsModel.Timezone = core.StringPtr("testString")
				createTektonPipelineTriggerOptionsModel.Source = triggerSourcePrototypeModel
				createTektonPipelineTriggerOptionsModel.Events = []string{"push"}
				createTektonPipelineTriggerOptionsModel.Filter = core.StringPtr("testString")
				createTektonPipelineTriggerOptionsModel.Favorite = core.BoolPtr(false)
				createTektonPipelineTriggerOptionsModel.EnableEventsFromForks = core.BoolPtr(false)
				createTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.CreateTektonPipelineTriggerWithContext(ctx, createTektonPipelineTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineTrigger(createTektonPipelineTriggerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.CreateTektonPipelineTriggerWithContext(ctx, createTektonPipelineTriggerOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelineTriggerPath))
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
					fmt.Fprintf(res, "%s", `{"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}`)
				}))
			})
			It(`Invoke CreateTektonPipelineTrigger successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineTrigger(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("public")

				// Construct an instance of the GenericSecret model
				genericSecretModel := new(cdtektonpipelinev2.GenericSecret)
				genericSecretModel.Type = core.StringPtr("token_matches")
				genericSecretModel.Value = core.StringPtr("testString")
				genericSecretModel.Source = core.StringPtr("header")
				genericSecretModel.KeyName = core.StringPtr("testString")
				genericSecretModel.Algorithm = core.StringPtr("md4")

				// Construct an instance of the TriggerSourcePropertiesPrototype model
				triggerSourcePropertiesPrototypeModel := new(cdtektonpipelinev2.TriggerSourcePropertiesPrototype)
				triggerSourcePropertiesPrototypeModel.URL = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Branch = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Pattern = core.StringPtr("testString")

				// Construct an instance of the TriggerSourcePrototype model
				triggerSourcePrototypeModel := new(cdtektonpipelinev2.TriggerSourcePrototype)
				triggerSourcePrototypeModel.Type = core.StringPtr("testString")
				triggerSourcePrototypeModel.Properties = triggerSourcePropertiesPrototypeModel

				// Construct an instance of the CreateTektonPipelineTriggerOptions model
				createTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineTriggerOptions)
				createTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineTriggerOptionsModel.Type = core.StringPtr("manual")
				createTektonPipelineTriggerOptionsModel.Name = core.StringPtr("Manual Trigger")
				createTektonPipelineTriggerOptionsModel.EventListener = core.StringPtr("pr-listener")
				createTektonPipelineTriggerOptionsModel.Tags = []string{"testString"}
				createTektonPipelineTriggerOptionsModel.Worker = workerIdentityModel
				createTektonPipelineTriggerOptionsModel.MaxConcurrentRuns = core.Int64Ptr(int64(3))
				createTektonPipelineTriggerOptionsModel.LimitWaitingRuns = core.BoolPtr(false)
				createTektonPipelineTriggerOptionsModel.Enabled = core.BoolPtr(true)
				createTektonPipelineTriggerOptionsModel.Secret = genericSecretModel
				createTektonPipelineTriggerOptionsModel.Cron = core.StringPtr("testString")
				createTektonPipelineTriggerOptionsModel.Timezone = core.StringPtr("testString")
				createTektonPipelineTriggerOptionsModel.Source = triggerSourcePrototypeModel
				createTektonPipelineTriggerOptionsModel.Events = []string{"push"}
				createTektonPipelineTriggerOptionsModel.Filter = core.StringPtr("testString")
				createTektonPipelineTriggerOptionsModel.Favorite = core.BoolPtr(false)
				createTektonPipelineTriggerOptionsModel.EnableEventsFromForks = core.BoolPtr(false)
				createTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.CreateTektonPipelineTrigger(createTektonPipelineTriggerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTektonPipelineTrigger with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("public")

				// Construct an instance of the GenericSecret model
				genericSecretModel := new(cdtektonpipelinev2.GenericSecret)
				genericSecretModel.Type = core.StringPtr("token_matches")
				genericSecretModel.Value = core.StringPtr("testString")
				genericSecretModel.Source = core.StringPtr("header")
				genericSecretModel.KeyName = core.StringPtr("testString")
				genericSecretModel.Algorithm = core.StringPtr("md4")

				// Construct an instance of the TriggerSourcePropertiesPrototype model
				triggerSourcePropertiesPrototypeModel := new(cdtektonpipelinev2.TriggerSourcePropertiesPrototype)
				triggerSourcePropertiesPrototypeModel.URL = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Branch = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Pattern = core.StringPtr("testString")

				// Construct an instance of the TriggerSourcePrototype model
				triggerSourcePrototypeModel := new(cdtektonpipelinev2.TriggerSourcePrototype)
				triggerSourcePrototypeModel.Type = core.StringPtr("testString")
				triggerSourcePrototypeModel.Properties = triggerSourcePropertiesPrototypeModel

				// Construct an instance of the CreateTektonPipelineTriggerOptions model
				createTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineTriggerOptions)
				createTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineTriggerOptionsModel.Type = core.StringPtr("manual")
				createTektonPipelineTriggerOptionsModel.Name = core.StringPtr("Manual Trigger")
				createTektonPipelineTriggerOptionsModel.EventListener = core.StringPtr("pr-listener")
				createTektonPipelineTriggerOptionsModel.Tags = []string{"testString"}
				createTektonPipelineTriggerOptionsModel.Worker = workerIdentityModel
				createTektonPipelineTriggerOptionsModel.MaxConcurrentRuns = core.Int64Ptr(int64(3))
				createTektonPipelineTriggerOptionsModel.LimitWaitingRuns = core.BoolPtr(false)
				createTektonPipelineTriggerOptionsModel.Enabled = core.BoolPtr(true)
				createTektonPipelineTriggerOptionsModel.Secret = genericSecretModel
				createTektonPipelineTriggerOptionsModel.Cron = core.StringPtr("testString")
				createTektonPipelineTriggerOptionsModel.Timezone = core.StringPtr("testString")
				createTektonPipelineTriggerOptionsModel.Source = triggerSourcePrototypeModel
				createTektonPipelineTriggerOptionsModel.Events = []string{"push"}
				createTektonPipelineTriggerOptionsModel.Filter = core.StringPtr("testString")
				createTektonPipelineTriggerOptionsModel.Favorite = core.BoolPtr(false)
				createTektonPipelineTriggerOptionsModel.EnableEventsFromForks = core.BoolPtr(false)
				createTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineTrigger(createTektonPipelineTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTektonPipelineTriggerOptions model with no property values
				createTektonPipelineTriggerOptionsModelNew := new(cdtektonpipelinev2.CreateTektonPipelineTriggerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.CreateTektonPipelineTrigger(createTektonPipelineTriggerOptionsModelNew)
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
			It(`Invoke CreateTektonPipelineTrigger successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("public")

				// Construct an instance of the GenericSecret model
				genericSecretModel := new(cdtektonpipelinev2.GenericSecret)
				genericSecretModel.Type = core.StringPtr("token_matches")
				genericSecretModel.Value = core.StringPtr("testString")
				genericSecretModel.Source = core.StringPtr("header")
				genericSecretModel.KeyName = core.StringPtr("testString")
				genericSecretModel.Algorithm = core.StringPtr("md4")

				// Construct an instance of the TriggerSourcePropertiesPrototype model
				triggerSourcePropertiesPrototypeModel := new(cdtektonpipelinev2.TriggerSourcePropertiesPrototype)
				triggerSourcePropertiesPrototypeModel.URL = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Branch = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Pattern = core.StringPtr("testString")

				// Construct an instance of the TriggerSourcePrototype model
				triggerSourcePrototypeModel := new(cdtektonpipelinev2.TriggerSourcePrototype)
				triggerSourcePrototypeModel.Type = core.StringPtr("testString")
				triggerSourcePrototypeModel.Properties = triggerSourcePropertiesPrototypeModel

				// Construct an instance of the CreateTektonPipelineTriggerOptions model
				createTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineTriggerOptions)
				createTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineTriggerOptionsModel.Type = core.StringPtr("manual")
				createTektonPipelineTriggerOptionsModel.Name = core.StringPtr("Manual Trigger")
				createTektonPipelineTriggerOptionsModel.EventListener = core.StringPtr("pr-listener")
				createTektonPipelineTriggerOptionsModel.Tags = []string{"testString"}
				createTektonPipelineTriggerOptionsModel.Worker = workerIdentityModel
				createTektonPipelineTriggerOptionsModel.MaxConcurrentRuns = core.Int64Ptr(int64(3))
				createTektonPipelineTriggerOptionsModel.LimitWaitingRuns = core.BoolPtr(false)
				createTektonPipelineTriggerOptionsModel.Enabled = core.BoolPtr(true)
				createTektonPipelineTriggerOptionsModel.Secret = genericSecretModel
				createTektonPipelineTriggerOptionsModel.Cron = core.StringPtr("testString")
				createTektonPipelineTriggerOptionsModel.Timezone = core.StringPtr("testString")
				createTektonPipelineTriggerOptionsModel.Source = triggerSourcePrototypeModel
				createTektonPipelineTriggerOptionsModel.Events = []string{"push"}
				createTektonPipelineTriggerOptionsModel.Filter = core.StringPtr("testString")
				createTektonPipelineTriggerOptionsModel.Favorite = core.BoolPtr(false)
				createTektonPipelineTriggerOptionsModel.EnableEventsFromForks = core.BoolPtr(false)
				createTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineTrigger(createTektonPipelineTriggerOptionsModel)
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
	Describe(`GetTektonPipelineTrigger(getTektonPipelineTriggerOptions *GetTektonPipelineTriggerOptions) - Operation response error`, func() {
		getTektonPipelineTriggerPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers/1bb892a1-2e04-4768-a369-b1159eace147"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineTriggerPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTektonPipelineTrigger with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineTriggerOptions model
				getTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineTriggerOptions)
				getTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineTriggerOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				getTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineTrigger(getTektonPipelineTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineTrigger(getTektonPipelineTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTektonPipelineTrigger(getTektonPipelineTriggerOptions *GetTektonPipelineTriggerOptions)`, func() {
		getTektonPipelineTriggerPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers/1bb892a1-2e04-4768-a369-b1159eace147"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineTriggerPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}`)
				}))
			})
			It(`Invoke GetTektonPipelineTrigger successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the GetTektonPipelineTriggerOptions model
				getTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineTriggerOptions)
				getTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineTriggerOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				getTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.GetTektonPipelineTriggerWithContext(ctx, getTektonPipelineTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineTrigger(getTektonPipelineTriggerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.GetTektonPipelineTriggerWithContext(ctx, getTektonPipelineTriggerOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineTriggerPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}`)
				}))
			})
			It(`Invoke GetTektonPipelineTrigger successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineTrigger(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTektonPipelineTriggerOptions model
				getTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineTriggerOptions)
				getTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineTriggerOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				getTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineTrigger(getTektonPipelineTriggerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTektonPipelineTrigger with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineTriggerOptions model
				getTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineTriggerOptions)
				getTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineTriggerOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				getTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineTrigger(getTektonPipelineTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTektonPipelineTriggerOptions model with no property values
				getTektonPipelineTriggerOptionsModelNew := new(cdtektonpipelinev2.GetTektonPipelineTriggerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineTrigger(getTektonPipelineTriggerOptionsModelNew)
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
			It(`Invoke GetTektonPipelineTrigger successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineTriggerOptions model
				getTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineTriggerOptions)
				getTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineTriggerOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				getTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineTrigger(getTektonPipelineTriggerOptionsModel)
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
	Describe(`UpdateTektonPipelineTrigger(updateTektonPipelineTriggerOptions *UpdateTektonPipelineTriggerOptions) - Operation response error`, func() {
		updateTektonPipelineTriggerPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers/1bb892a1-2e04-4768-a369-b1159eace147"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTektonPipelineTriggerPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTektonPipelineTrigger with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the GenericSecret model
				genericSecretModel := new(cdtektonpipelinev2.GenericSecret)
				genericSecretModel.Type = core.StringPtr("token_matches")
				genericSecretModel.Value = core.StringPtr("testString")
				genericSecretModel.Source = core.StringPtr("header")
				genericSecretModel.KeyName = core.StringPtr("testString")
				genericSecretModel.Algorithm = core.StringPtr("md4")

				// Construct an instance of the TriggerSourcePropertiesPrototype model
				triggerSourcePropertiesPrototypeModel := new(cdtektonpipelinev2.TriggerSourcePropertiesPrototype)
				triggerSourcePropertiesPrototypeModel.URL = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Branch = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Pattern = core.StringPtr("testString")

				// Construct an instance of the TriggerSourcePrototype model
				triggerSourcePrototypeModel := new(cdtektonpipelinev2.TriggerSourcePrototype)
				triggerSourcePrototypeModel.Type = core.StringPtr("testString")
				triggerSourcePrototypeModel.Properties = triggerSourcePropertiesPrototypeModel

				// Construct an instance of the TriggerPatch model
				triggerPatchModel := new(cdtektonpipelinev2.TriggerPatch)
				triggerPatchModel.Type = core.StringPtr("manual")
				triggerPatchModel.Name = core.StringPtr("start-deploy")
				triggerPatchModel.EventListener = core.StringPtr("testString")
				triggerPatchModel.Tags = []string{"testString"}
				triggerPatchModel.Worker = workerIdentityModel
				triggerPatchModel.MaxConcurrentRuns = core.Int64Ptr(int64(4))
				triggerPatchModel.LimitWaitingRuns = core.BoolPtr(false)
				triggerPatchModel.Enabled = core.BoolPtr(true)
				triggerPatchModel.Secret = genericSecretModel
				triggerPatchModel.Cron = core.StringPtr("testString")
				triggerPatchModel.Timezone = core.StringPtr("America/Los_Angeles, CET, Europe/London, GMT, US/Eastern, or UTC")
				triggerPatchModel.Source = triggerSourcePrototypeModel
				triggerPatchModel.Events = []string{"push", "pull_request"}
				triggerPatchModel.Filter = core.StringPtr("header['x-github-event'] == 'push' && body.ref == 'refs/heads/main'")
				triggerPatchModel.Favorite = core.BoolPtr(false)
				triggerPatchModel.EnableEventsFromForks = core.BoolPtr(false)
				triggerPatchModelAsPatch, asPatchErr := triggerPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTektonPipelineTriggerOptions model
				updateTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.UpdateTektonPipelineTriggerOptions)
				updateTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				updateTektonPipelineTriggerOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				updateTektonPipelineTriggerOptionsModel.TriggerPatch = triggerPatchModelAsPatch
				updateTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.UpdateTektonPipelineTrigger(updateTektonPipelineTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.UpdateTektonPipelineTrigger(updateTektonPipelineTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTektonPipelineTrigger(updateTektonPipelineTriggerOptions *UpdateTektonPipelineTriggerOptions)`, func() {
		updateTektonPipelineTriggerPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers/1bb892a1-2e04-4768-a369-b1159eace147"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTektonPipelineTriggerPath))
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
					fmt.Fprintf(res, "%s", `{"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}`)
				}))
			})
			It(`Invoke UpdateTektonPipelineTrigger successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the GenericSecret model
				genericSecretModel := new(cdtektonpipelinev2.GenericSecret)
				genericSecretModel.Type = core.StringPtr("token_matches")
				genericSecretModel.Value = core.StringPtr("testString")
				genericSecretModel.Source = core.StringPtr("header")
				genericSecretModel.KeyName = core.StringPtr("testString")
				genericSecretModel.Algorithm = core.StringPtr("md4")

				// Construct an instance of the TriggerSourcePropertiesPrototype model
				triggerSourcePropertiesPrototypeModel := new(cdtektonpipelinev2.TriggerSourcePropertiesPrototype)
				triggerSourcePropertiesPrototypeModel.URL = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Branch = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Pattern = core.StringPtr("testString")

				// Construct an instance of the TriggerSourcePrototype model
				triggerSourcePrototypeModel := new(cdtektonpipelinev2.TriggerSourcePrototype)
				triggerSourcePrototypeModel.Type = core.StringPtr("testString")
				triggerSourcePrototypeModel.Properties = triggerSourcePropertiesPrototypeModel

				// Construct an instance of the TriggerPatch model
				triggerPatchModel := new(cdtektonpipelinev2.TriggerPatch)
				triggerPatchModel.Type = core.StringPtr("manual")
				triggerPatchModel.Name = core.StringPtr("start-deploy")
				triggerPatchModel.EventListener = core.StringPtr("testString")
				triggerPatchModel.Tags = []string{"testString"}
				triggerPatchModel.Worker = workerIdentityModel
				triggerPatchModel.MaxConcurrentRuns = core.Int64Ptr(int64(4))
				triggerPatchModel.LimitWaitingRuns = core.BoolPtr(false)
				triggerPatchModel.Enabled = core.BoolPtr(true)
				triggerPatchModel.Secret = genericSecretModel
				triggerPatchModel.Cron = core.StringPtr("testString")
				triggerPatchModel.Timezone = core.StringPtr("America/Los_Angeles, CET, Europe/London, GMT, US/Eastern, or UTC")
				triggerPatchModel.Source = triggerSourcePrototypeModel
				triggerPatchModel.Events = []string{"push", "pull_request"}
				triggerPatchModel.Filter = core.StringPtr("header['x-github-event'] == 'push' && body.ref == 'refs/heads/main'")
				triggerPatchModel.Favorite = core.BoolPtr(false)
				triggerPatchModel.EnableEventsFromForks = core.BoolPtr(false)
				triggerPatchModelAsPatch, asPatchErr := triggerPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTektonPipelineTriggerOptions model
				updateTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.UpdateTektonPipelineTriggerOptions)
				updateTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				updateTektonPipelineTriggerOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				updateTektonPipelineTriggerOptionsModel.TriggerPatch = triggerPatchModelAsPatch
				updateTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.UpdateTektonPipelineTriggerWithContext(ctx, updateTektonPipelineTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.UpdateTektonPipelineTrigger(updateTektonPipelineTriggerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.UpdateTektonPipelineTriggerWithContext(ctx, updateTektonPipelineTriggerOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateTektonPipelineTriggerPath))
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
					fmt.Fprintf(res, "%s", `{"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}`)
				}))
			})
			It(`Invoke UpdateTektonPipelineTrigger successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.UpdateTektonPipelineTrigger(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the GenericSecret model
				genericSecretModel := new(cdtektonpipelinev2.GenericSecret)
				genericSecretModel.Type = core.StringPtr("token_matches")
				genericSecretModel.Value = core.StringPtr("testString")
				genericSecretModel.Source = core.StringPtr("header")
				genericSecretModel.KeyName = core.StringPtr("testString")
				genericSecretModel.Algorithm = core.StringPtr("md4")

				// Construct an instance of the TriggerSourcePropertiesPrototype model
				triggerSourcePropertiesPrototypeModel := new(cdtektonpipelinev2.TriggerSourcePropertiesPrototype)
				triggerSourcePropertiesPrototypeModel.URL = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Branch = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Pattern = core.StringPtr("testString")

				// Construct an instance of the TriggerSourcePrototype model
				triggerSourcePrototypeModel := new(cdtektonpipelinev2.TriggerSourcePrototype)
				triggerSourcePrototypeModel.Type = core.StringPtr("testString")
				triggerSourcePrototypeModel.Properties = triggerSourcePropertiesPrototypeModel

				// Construct an instance of the TriggerPatch model
				triggerPatchModel := new(cdtektonpipelinev2.TriggerPatch)
				triggerPatchModel.Type = core.StringPtr("manual")
				triggerPatchModel.Name = core.StringPtr("start-deploy")
				triggerPatchModel.EventListener = core.StringPtr("testString")
				triggerPatchModel.Tags = []string{"testString"}
				triggerPatchModel.Worker = workerIdentityModel
				triggerPatchModel.MaxConcurrentRuns = core.Int64Ptr(int64(4))
				triggerPatchModel.LimitWaitingRuns = core.BoolPtr(false)
				triggerPatchModel.Enabled = core.BoolPtr(true)
				triggerPatchModel.Secret = genericSecretModel
				triggerPatchModel.Cron = core.StringPtr("testString")
				triggerPatchModel.Timezone = core.StringPtr("America/Los_Angeles, CET, Europe/London, GMT, US/Eastern, or UTC")
				triggerPatchModel.Source = triggerSourcePrototypeModel
				triggerPatchModel.Events = []string{"push", "pull_request"}
				triggerPatchModel.Filter = core.StringPtr("header['x-github-event'] == 'push' && body.ref == 'refs/heads/main'")
				triggerPatchModel.Favorite = core.BoolPtr(false)
				triggerPatchModel.EnableEventsFromForks = core.BoolPtr(false)
				triggerPatchModelAsPatch, asPatchErr := triggerPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTektonPipelineTriggerOptions model
				updateTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.UpdateTektonPipelineTriggerOptions)
				updateTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				updateTektonPipelineTriggerOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				updateTektonPipelineTriggerOptionsModel.TriggerPatch = triggerPatchModelAsPatch
				updateTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.UpdateTektonPipelineTrigger(updateTektonPipelineTriggerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTektonPipelineTrigger with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the GenericSecret model
				genericSecretModel := new(cdtektonpipelinev2.GenericSecret)
				genericSecretModel.Type = core.StringPtr("token_matches")
				genericSecretModel.Value = core.StringPtr("testString")
				genericSecretModel.Source = core.StringPtr("header")
				genericSecretModel.KeyName = core.StringPtr("testString")
				genericSecretModel.Algorithm = core.StringPtr("md4")

				// Construct an instance of the TriggerSourcePropertiesPrototype model
				triggerSourcePropertiesPrototypeModel := new(cdtektonpipelinev2.TriggerSourcePropertiesPrototype)
				triggerSourcePropertiesPrototypeModel.URL = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Branch = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Pattern = core.StringPtr("testString")

				// Construct an instance of the TriggerSourcePrototype model
				triggerSourcePrototypeModel := new(cdtektonpipelinev2.TriggerSourcePrototype)
				triggerSourcePrototypeModel.Type = core.StringPtr("testString")
				triggerSourcePrototypeModel.Properties = triggerSourcePropertiesPrototypeModel

				// Construct an instance of the TriggerPatch model
				triggerPatchModel := new(cdtektonpipelinev2.TriggerPatch)
				triggerPatchModel.Type = core.StringPtr("manual")
				triggerPatchModel.Name = core.StringPtr("start-deploy")
				triggerPatchModel.EventListener = core.StringPtr("testString")
				triggerPatchModel.Tags = []string{"testString"}
				triggerPatchModel.Worker = workerIdentityModel
				triggerPatchModel.MaxConcurrentRuns = core.Int64Ptr(int64(4))
				triggerPatchModel.LimitWaitingRuns = core.BoolPtr(false)
				triggerPatchModel.Enabled = core.BoolPtr(true)
				triggerPatchModel.Secret = genericSecretModel
				triggerPatchModel.Cron = core.StringPtr("testString")
				triggerPatchModel.Timezone = core.StringPtr("America/Los_Angeles, CET, Europe/London, GMT, US/Eastern, or UTC")
				triggerPatchModel.Source = triggerSourcePrototypeModel
				triggerPatchModel.Events = []string{"push", "pull_request"}
				triggerPatchModel.Filter = core.StringPtr("header['x-github-event'] == 'push' && body.ref == 'refs/heads/main'")
				triggerPatchModel.Favorite = core.BoolPtr(false)
				triggerPatchModel.EnableEventsFromForks = core.BoolPtr(false)
				triggerPatchModelAsPatch, asPatchErr := triggerPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTektonPipelineTriggerOptions model
				updateTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.UpdateTektonPipelineTriggerOptions)
				updateTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				updateTektonPipelineTriggerOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				updateTektonPipelineTriggerOptionsModel.TriggerPatch = triggerPatchModelAsPatch
				updateTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.UpdateTektonPipelineTrigger(updateTektonPipelineTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTektonPipelineTriggerOptions model with no property values
				updateTektonPipelineTriggerOptionsModelNew := new(cdtektonpipelinev2.UpdateTektonPipelineTriggerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.UpdateTektonPipelineTrigger(updateTektonPipelineTriggerOptionsModelNew)
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
			It(`Invoke UpdateTektonPipelineTrigger successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				workerIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the GenericSecret model
				genericSecretModel := new(cdtektonpipelinev2.GenericSecret)
				genericSecretModel.Type = core.StringPtr("token_matches")
				genericSecretModel.Value = core.StringPtr("testString")
				genericSecretModel.Source = core.StringPtr("header")
				genericSecretModel.KeyName = core.StringPtr("testString")
				genericSecretModel.Algorithm = core.StringPtr("md4")

				// Construct an instance of the TriggerSourcePropertiesPrototype model
				triggerSourcePropertiesPrototypeModel := new(cdtektonpipelinev2.TriggerSourcePropertiesPrototype)
				triggerSourcePropertiesPrototypeModel.URL = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Branch = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Pattern = core.StringPtr("testString")

				// Construct an instance of the TriggerSourcePrototype model
				triggerSourcePrototypeModel := new(cdtektonpipelinev2.TriggerSourcePrototype)
				triggerSourcePrototypeModel.Type = core.StringPtr("testString")
				triggerSourcePrototypeModel.Properties = triggerSourcePropertiesPrototypeModel

				// Construct an instance of the TriggerPatch model
				triggerPatchModel := new(cdtektonpipelinev2.TriggerPatch)
				triggerPatchModel.Type = core.StringPtr("manual")
				triggerPatchModel.Name = core.StringPtr("start-deploy")
				triggerPatchModel.EventListener = core.StringPtr("testString")
				triggerPatchModel.Tags = []string{"testString"}
				triggerPatchModel.Worker = workerIdentityModel
				triggerPatchModel.MaxConcurrentRuns = core.Int64Ptr(int64(4))
				triggerPatchModel.LimitWaitingRuns = core.BoolPtr(false)
				triggerPatchModel.Enabled = core.BoolPtr(true)
				triggerPatchModel.Secret = genericSecretModel
				triggerPatchModel.Cron = core.StringPtr("testString")
				triggerPatchModel.Timezone = core.StringPtr("America/Los_Angeles, CET, Europe/London, GMT, US/Eastern, or UTC")
				triggerPatchModel.Source = triggerSourcePrototypeModel
				triggerPatchModel.Events = []string{"push", "pull_request"}
				triggerPatchModel.Filter = core.StringPtr("header['x-github-event'] == 'push' && body.ref == 'refs/heads/main'")
				triggerPatchModel.Favorite = core.BoolPtr(false)
				triggerPatchModel.EnableEventsFromForks = core.BoolPtr(false)
				triggerPatchModelAsPatch, asPatchErr := triggerPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTektonPipelineTriggerOptions model
				updateTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.UpdateTektonPipelineTriggerOptions)
				updateTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				updateTektonPipelineTriggerOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				updateTektonPipelineTriggerOptionsModel.TriggerPatch = triggerPatchModelAsPatch
				updateTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.UpdateTektonPipelineTrigger(updateTektonPipelineTriggerOptionsModel)
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
	Describe(`DeleteTektonPipelineTrigger(deleteTektonPipelineTriggerOptions *DeleteTektonPipelineTriggerOptions)`, func() {
		deleteTektonPipelineTriggerPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers/1bb892a1-2e04-4768-a369-b1159eace147"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTektonPipelineTriggerPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTektonPipelineTrigger successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := cdTektonPipelineService.DeleteTektonPipelineTrigger(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTektonPipelineTriggerOptions model
				deleteTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.DeleteTektonPipelineTriggerOptions)
				deleteTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelineTriggerOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				deleteTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = cdTektonPipelineService.DeleteTektonPipelineTrigger(deleteTektonPipelineTriggerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTektonPipelineTrigger with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the DeleteTektonPipelineTriggerOptions model
				deleteTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.DeleteTektonPipelineTriggerOptions)
				deleteTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelineTriggerOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				deleteTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := cdTektonPipelineService.DeleteTektonPipelineTrigger(deleteTektonPipelineTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTektonPipelineTriggerOptions model with no property values
				deleteTektonPipelineTriggerOptionsModelNew := new(cdtektonpipelinev2.DeleteTektonPipelineTriggerOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = cdTektonPipelineService.DeleteTektonPipelineTrigger(deleteTektonPipelineTriggerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DuplicateTektonPipelineTrigger(duplicateTektonPipelineTriggerOptions *DuplicateTektonPipelineTriggerOptions) - Operation response error`, func() {
		duplicateTektonPipelineTriggerPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers/1bb892a1-2e04-4768-a369-b1159eace147/duplicate"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(duplicateTektonPipelineTriggerPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DuplicateTektonPipelineTrigger with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the DuplicateTektonPipelineTriggerOptions model
				duplicateTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.DuplicateTektonPipelineTriggerOptions)
				duplicateTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				duplicateTektonPipelineTriggerOptionsModel.SourceTriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				duplicateTektonPipelineTriggerOptionsModel.Name = core.StringPtr("triggerName")
				duplicateTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.DuplicateTektonPipelineTrigger(duplicateTektonPipelineTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.DuplicateTektonPipelineTrigger(duplicateTektonPipelineTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DuplicateTektonPipelineTrigger(duplicateTektonPipelineTriggerOptions *DuplicateTektonPipelineTriggerOptions)`, func() {
		duplicateTektonPipelineTriggerPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers/1bb892a1-2e04-4768-a369-b1159eace147/duplicate"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(duplicateTektonPipelineTriggerPath))
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
					fmt.Fprintf(res, "%s", `{"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}`)
				}))
			})
			It(`Invoke DuplicateTektonPipelineTrigger successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the DuplicateTektonPipelineTriggerOptions model
				duplicateTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.DuplicateTektonPipelineTriggerOptions)
				duplicateTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				duplicateTektonPipelineTriggerOptionsModel.SourceTriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				duplicateTektonPipelineTriggerOptionsModel.Name = core.StringPtr("triggerName")
				duplicateTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.DuplicateTektonPipelineTriggerWithContext(ctx, duplicateTektonPipelineTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.DuplicateTektonPipelineTrigger(duplicateTektonPipelineTriggerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.DuplicateTektonPipelineTriggerWithContext(ctx, duplicateTektonPipelineTriggerOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(duplicateTektonPipelineTriggerPath))
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
					fmt.Fprintf(res, "%s", `{"type": "Type", "name": "start-deploy", "href": "Href", "event_listener": "EventListener", "id": "ID", "properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}], "tags": ["Tags"], "worker": {"name": "Name", "type": "Type", "id": "ID"}, "max_concurrent_runs": 4, "enabled": true, "favorite": false}`)
				}))
			})
			It(`Invoke DuplicateTektonPipelineTrigger successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.DuplicateTektonPipelineTrigger(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DuplicateTektonPipelineTriggerOptions model
				duplicateTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.DuplicateTektonPipelineTriggerOptions)
				duplicateTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				duplicateTektonPipelineTriggerOptionsModel.SourceTriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				duplicateTektonPipelineTriggerOptionsModel.Name = core.StringPtr("triggerName")
				duplicateTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.DuplicateTektonPipelineTrigger(duplicateTektonPipelineTriggerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DuplicateTektonPipelineTrigger with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the DuplicateTektonPipelineTriggerOptions model
				duplicateTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.DuplicateTektonPipelineTriggerOptions)
				duplicateTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				duplicateTektonPipelineTriggerOptionsModel.SourceTriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				duplicateTektonPipelineTriggerOptionsModel.Name = core.StringPtr("triggerName")
				duplicateTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.DuplicateTektonPipelineTrigger(duplicateTektonPipelineTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DuplicateTektonPipelineTriggerOptions model with no property values
				duplicateTektonPipelineTriggerOptionsModelNew := new(cdtektonpipelinev2.DuplicateTektonPipelineTriggerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.DuplicateTektonPipelineTrigger(duplicateTektonPipelineTriggerOptionsModelNew)
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
			It(`Invoke DuplicateTektonPipelineTrigger successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the DuplicateTektonPipelineTriggerOptions model
				duplicateTektonPipelineTriggerOptionsModel := new(cdtektonpipelinev2.DuplicateTektonPipelineTriggerOptions)
				duplicateTektonPipelineTriggerOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				duplicateTektonPipelineTriggerOptionsModel.SourceTriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				duplicateTektonPipelineTriggerOptionsModel.Name = core.StringPtr("triggerName")
				duplicateTektonPipelineTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.DuplicateTektonPipelineTrigger(duplicateTektonPipelineTriggerOptionsModel)
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
	Describe(`ListTektonPipelineTriggerProperties(listTektonPipelineTriggerPropertiesOptions *ListTektonPipelineTriggerPropertiesOptions) - Operation response error`, func() {
		listTektonPipelineTriggerPropertiesPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers/1bb892a1-2e04-4768-a369-b1159eace147/properties"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTektonPipelineTriggerPropertiesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"prod"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"secure,text"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"name"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTektonPipelineTriggerProperties with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ListTektonPipelineTriggerPropertiesOptions model
				listTektonPipelineTriggerPropertiesOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineTriggerPropertiesOptions)
				listTektonPipelineTriggerPropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineTriggerPropertiesOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				listTektonPipelineTriggerPropertiesOptionsModel.Name = core.StringPtr("prod")
				listTektonPipelineTriggerPropertiesOptionsModel.Type = core.StringPtr("secure,text")
				listTektonPipelineTriggerPropertiesOptionsModel.Sort = core.StringPtr("name")
				listTektonPipelineTriggerPropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineTriggerProperties(listTektonPipelineTriggerPropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.ListTektonPipelineTriggerProperties(listTektonPipelineTriggerPropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTektonPipelineTriggerProperties(listTektonPipelineTriggerPropertiesOptions *ListTektonPipelineTriggerPropertiesOptions)`, func() {
		listTektonPipelineTriggerPropertiesPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers/1bb892a1-2e04-4768-a369-b1159eace147/properties"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTektonPipelineTriggerPropertiesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"prod"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"secure,text"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"name"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}]}`)
				}))
			})
			It(`Invoke ListTektonPipelineTriggerProperties successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the ListTektonPipelineTriggerPropertiesOptions model
				listTektonPipelineTriggerPropertiesOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineTriggerPropertiesOptions)
				listTektonPipelineTriggerPropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineTriggerPropertiesOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				listTektonPipelineTriggerPropertiesOptionsModel.Name = core.StringPtr("prod")
				listTektonPipelineTriggerPropertiesOptionsModel.Type = core.StringPtr("secure,text")
				listTektonPipelineTriggerPropertiesOptionsModel.Sort = core.StringPtr("name")
				listTektonPipelineTriggerPropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.ListTektonPipelineTriggerPropertiesWithContext(ctx, listTektonPipelineTriggerPropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineTriggerProperties(listTektonPipelineTriggerPropertiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.ListTektonPipelineTriggerPropertiesWithContext(ctx, listTektonPipelineTriggerPropertiesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTektonPipelineTriggerPropertiesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"prod"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"secure,text"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"name"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"properties": [{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}]}`)
				}))
			})
			It(`Invoke ListTektonPipelineTriggerProperties successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineTriggerProperties(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTektonPipelineTriggerPropertiesOptions model
				listTektonPipelineTriggerPropertiesOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineTriggerPropertiesOptions)
				listTektonPipelineTriggerPropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineTriggerPropertiesOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				listTektonPipelineTriggerPropertiesOptionsModel.Name = core.StringPtr("prod")
				listTektonPipelineTriggerPropertiesOptionsModel.Type = core.StringPtr("secure,text")
				listTektonPipelineTriggerPropertiesOptionsModel.Sort = core.StringPtr("name")
				listTektonPipelineTriggerPropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.ListTektonPipelineTriggerProperties(listTektonPipelineTriggerPropertiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTektonPipelineTriggerProperties with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ListTektonPipelineTriggerPropertiesOptions model
				listTektonPipelineTriggerPropertiesOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineTriggerPropertiesOptions)
				listTektonPipelineTriggerPropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineTriggerPropertiesOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				listTektonPipelineTriggerPropertiesOptionsModel.Name = core.StringPtr("prod")
				listTektonPipelineTriggerPropertiesOptionsModel.Type = core.StringPtr("secure,text")
				listTektonPipelineTriggerPropertiesOptionsModel.Sort = core.StringPtr("name")
				listTektonPipelineTriggerPropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineTriggerProperties(listTektonPipelineTriggerPropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTektonPipelineTriggerPropertiesOptions model with no property values
				listTektonPipelineTriggerPropertiesOptionsModelNew := new(cdtektonpipelinev2.ListTektonPipelineTriggerPropertiesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.ListTektonPipelineTriggerProperties(listTektonPipelineTriggerPropertiesOptionsModelNew)
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
			It(`Invoke ListTektonPipelineTriggerProperties successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ListTektonPipelineTriggerPropertiesOptions model
				listTektonPipelineTriggerPropertiesOptionsModel := new(cdtektonpipelinev2.ListTektonPipelineTriggerPropertiesOptions)
				listTektonPipelineTriggerPropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineTriggerPropertiesOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				listTektonPipelineTriggerPropertiesOptionsModel.Name = core.StringPtr("prod")
				listTektonPipelineTriggerPropertiesOptionsModel.Type = core.StringPtr("secure,text")
				listTektonPipelineTriggerPropertiesOptionsModel.Sort = core.StringPtr("name")
				listTektonPipelineTriggerPropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.ListTektonPipelineTriggerProperties(listTektonPipelineTriggerPropertiesOptionsModel)
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
	Describe(`CreateTektonPipelineTriggerProperties(createTektonPipelineTriggerPropertiesOptions *CreateTektonPipelineTriggerPropertiesOptions) - Operation response error`, func() {
		createTektonPipelineTriggerPropertiesPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers/1bb892a1-2e04-4768-a369-b1159eace147/properties"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelineTriggerPropertiesPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTektonPipelineTriggerProperties with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the CreateTektonPipelineTriggerPropertiesOptions model
				createTektonPipelineTriggerPropertiesOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineTriggerPropertiesOptions)
				createTektonPipelineTriggerPropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineTriggerPropertiesOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				createTektonPipelineTriggerPropertiesOptionsModel.Name = core.StringPtr("prop1")
				createTektonPipelineTriggerPropertiesOptionsModel.Type = core.StringPtr("text")
				createTektonPipelineTriggerPropertiesOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				createTektonPipelineTriggerPropertiesOptionsModel.Enum = []string{"testString"}
				createTektonPipelineTriggerPropertiesOptionsModel.Path = core.StringPtr("testString")
				createTektonPipelineTriggerPropertiesOptionsModel.Locked = core.BoolPtr(false)
				createTektonPipelineTriggerPropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineTriggerProperties(createTektonPipelineTriggerPropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.CreateTektonPipelineTriggerProperties(createTektonPipelineTriggerPropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTektonPipelineTriggerProperties(createTektonPipelineTriggerPropertiesOptions *CreateTektonPipelineTriggerPropertiesOptions)`, func() {
		createTektonPipelineTriggerPropertiesPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers/1bb892a1-2e04-4768-a369-b1159eace147/properties"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelineTriggerPropertiesPath))
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}`)
				}))
			})
			It(`Invoke CreateTektonPipelineTriggerProperties successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the CreateTektonPipelineTriggerPropertiesOptions model
				createTektonPipelineTriggerPropertiesOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineTriggerPropertiesOptions)
				createTektonPipelineTriggerPropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineTriggerPropertiesOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				createTektonPipelineTriggerPropertiesOptionsModel.Name = core.StringPtr("prop1")
				createTektonPipelineTriggerPropertiesOptionsModel.Type = core.StringPtr("text")
				createTektonPipelineTriggerPropertiesOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				createTektonPipelineTriggerPropertiesOptionsModel.Enum = []string{"testString"}
				createTektonPipelineTriggerPropertiesOptionsModel.Path = core.StringPtr("testString")
				createTektonPipelineTriggerPropertiesOptionsModel.Locked = core.BoolPtr(false)
				createTektonPipelineTriggerPropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.CreateTektonPipelineTriggerPropertiesWithContext(ctx, createTektonPipelineTriggerPropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineTriggerProperties(createTektonPipelineTriggerPropertiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.CreateTektonPipelineTriggerPropertiesWithContext(ctx, createTektonPipelineTriggerPropertiesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelineTriggerPropertiesPath))
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}`)
				}))
			})
			It(`Invoke CreateTektonPipelineTriggerProperties successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineTriggerProperties(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateTektonPipelineTriggerPropertiesOptions model
				createTektonPipelineTriggerPropertiesOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineTriggerPropertiesOptions)
				createTektonPipelineTriggerPropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineTriggerPropertiesOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				createTektonPipelineTriggerPropertiesOptionsModel.Name = core.StringPtr("prop1")
				createTektonPipelineTriggerPropertiesOptionsModel.Type = core.StringPtr("text")
				createTektonPipelineTriggerPropertiesOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				createTektonPipelineTriggerPropertiesOptionsModel.Enum = []string{"testString"}
				createTektonPipelineTriggerPropertiesOptionsModel.Path = core.StringPtr("testString")
				createTektonPipelineTriggerPropertiesOptionsModel.Locked = core.BoolPtr(false)
				createTektonPipelineTriggerPropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.CreateTektonPipelineTriggerProperties(createTektonPipelineTriggerPropertiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTektonPipelineTriggerProperties with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the CreateTektonPipelineTriggerPropertiesOptions model
				createTektonPipelineTriggerPropertiesOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineTriggerPropertiesOptions)
				createTektonPipelineTriggerPropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineTriggerPropertiesOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				createTektonPipelineTriggerPropertiesOptionsModel.Name = core.StringPtr("prop1")
				createTektonPipelineTriggerPropertiesOptionsModel.Type = core.StringPtr("text")
				createTektonPipelineTriggerPropertiesOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				createTektonPipelineTriggerPropertiesOptionsModel.Enum = []string{"testString"}
				createTektonPipelineTriggerPropertiesOptionsModel.Path = core.StringPtr("testString")
				createTektonPipelineTriggerPropertiesOptionsModel.Locked = core.BoolPtr(false)
				createTektonPipelineTriggerPropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineTriggerProperties(createTektonPipelineTriggerPropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTektonPipelineTriggerPropertiesOptions model with no property values
				createTektonPipelineTriggerPropertiesOptionsModelNew := new(cdtektonpipelinev2.CreateTektonPipelineTriggerPropertiesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.CreateTektonPipelineTriggerProperties(createTektonPipelineTriggerPropertiesOptionsModelNew)
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
			It(`Invoke CreateTektonPipelineTriggerProperties successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the CreateTektonPipelineTriggerPropertiesOptions model
				createTektonPipelineTriggerPropertiesOptionsModel := new(cdtektonpipelinev2.CreateTektonPipelineTriggerPropertiesOptions)
				createTektonPipelineTriggerPropertiesOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineTriggerPropertiesOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				createTektonPipelineTriggerPropertiesOptionsModel.Name = core.StringPtr("prop1")
				createTektonPipelineTriggerPropertiesOptionsModel.Type = core.StringPtr("text")
				createTektonPipelineTriggerPropertiesOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				createTektonPipelineTriggerPropertiesOptionsModel.Enum = []string{"testString"}
				createTektonPipelineTriggerPropertiesOptionsModel.Path = core.StringPtr("testString")
				createTektonPipelineTriggerPropertiesOptionsModel.Locked = core.BoolPtr(false)
				createTektonPipelineTriggerPropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.CreateTektonPipelineTriggerProperties(createTektonPipelineTriggerPropertiesOptionsModel)
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
	Describe(`GetTektonPipelineTriggerProperty(getTektonPipelineTriggerPropertyOptions *GetTektonPipelineTriggerPropertyOptions) - Operation response error`, func() {
		getTektonPipelineTriggerPropertyPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers/1bb892a1-2e04-4768-a369-b1159eace147/properties/debug-pipeline"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineTriggerPropertyPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTektonPipelineTriggerProperty with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineTriggerPropertyOptions model
				getTektonPipelineTriggerPropertyOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineTriggerPropertyOptions)
				getTektonPipelineTriggerPropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineTriggerPropertyOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				getTektonPipelineTriggerPropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				getTektonPipelineTriggerPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineTriggerProperty(getTektonPipelineTriggerPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineTriggerProperty(getTektonPipelineTriggerPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTektonPipelineTriggerProperty(getTektonPipelineTriggerPropertyOptions *GetTektonPipelineTriggerPropertyOptions)`, func() {
		getTektonPipelineTriggerPropertyPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers/1bb892a1-2e04-4768-a369-b1159eace147/properties/debug-pipeline"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineTriggerPropertyPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}`)
				}))
			})
			It(`Invoke GetTektonPipelineTriggerProperty successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the GetTektonPipelineTriggerPropertyOptions model
				getTektonPipelineTriggerPropertyOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineTriggerPropertyOptions)
				getTektonPipelineTriggerPropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineTriggerPropertyOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				getTektonPipelineTriggerPropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				getTektonPipelineTriggerPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.GetTektonPipelineTriggerPropertyWithContext(ctx, getTektonPipelineTriggerPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineTriggerProperty(getTektonPipelineTriggerPropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.GetTektonPipelineTriggerPropertyWithContext(ctx, getTektonPipelineTriggerPropertyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineTriggerPropertyPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}`)
				}))
			})
			It(`Invoke GetTektonPipelineTriggerProperty successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineTriggerProperty(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTektonPipelineTriggerPropertyOptions model
				getTektonPipelineTriggerPropertyOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineTriggerPropertyOptions)
				getTektonPipelineTriggerPropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineTriggerPropertyOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				getTektonPipelineTriggerPropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				getTektonPipelineTriggerPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineTriggerProperty(getTektonPipelineTriggerPropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTektonPipelineTriggerProperty with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineTriggerPropertyOptions model
				getTektonPipelineTriggerPropertyOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineTriggerPropertyOptions)
				getTektonPipelineTriggerPropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineTriggerPropertyOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				getTektonPipelineTriggerPropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				getTektonPipelineTriggerPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineTriggerProperty(getTektonPipelineTriggerPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTektonPipelineTriggerPropertyOptions model with no property values
				getTektonPipelineTriggerPropertyOptionsModelNew := new(cdtektonpipelinev2.GetTektonPipelineTriggerPropertyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.GetTektonPipelineTriggerProperty(getTektonPipelineTriggerPropertyOptionsModelNew)
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
			It(`Invoke GetTektonPipelineTriggerProperty successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineTriggerPropertyOptions model
				getTektonPipelineTriggerPropertyOptionsModel := new(cdtektonpipelinev2.GetTektonPipelineTriggerPropertyOptions)
				getTektonPipelineTriggerPropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineTriggerPropertyOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				getTektonPipelineTriggerPropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				getTektonPipelineTriggerPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.GetTektonPipelineTriggerProperty(getTektonPipelineTriggerPropertyOptionsModel)
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
	Describe(`ReplaceTektonPipelineTriggerProperty(replaceTektonPipelineTriggerPropertyOptions *ReplaceTektonPipelineTriggerPropertyOptions) - Operation response error`, func() {
		replaceTektonPipelineTriggerPropertyPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers/1bb892a1-2e04-4768-a369-b1159eace147/properties/debug-pipeline"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceTektonPipelineTriggerPropertyPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceTektonPipelineTriggerProperty with error: Operation response processing error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ReplaceTektonPipelineTriggerPropertyOptions model
				replaceTektonPipelineTriggerPropertyOptionsModel := new(cdtektonpipelinev2.ReplaceTektonPipelineTriggerPropertyOptions)
				replaceTektonPipelineTriggerPropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				replaceTektonPipelineTriggerPropertyOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				replaceTektonPipelineTriggerPropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				replaceTektonPipelineTriggerPropertyOptionsModel.Name = core.StringPtr("prop1")
				replaceTektonPipelineTriggerPropertyOptionsModel.Type = core.StringPtr("text")
				replaceTektonPipelineTriggerPropertyOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				replaceTektonPipelineTriggerPropertyOptionsModel.Enum = []string{"testString"}
				replaceTektonPipelineTriggerPropertyOptionsModel.Path = core.StringPtr("testString")
				replaceTektonPipelineTriggerPropertyOptionsModel.Locked = core.BoolPtr(false)
				replaceTektonPipelineTriggerPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cdTektonPipelineService.ReplaceTektonPipelineTriggerProperty(replaceTektonPipelineTriggerPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cdTektonPipelineService.EnableRetries(0, 0)
				result, response, operationErr = cdTektonPipelineService.ReplaceTektonPipelineTriggerProperty(replaceTektonPipelineTriggerPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceTektonPipelineTriggerProperty(replaceTektonPipelineTriggerPropertyOptions *ReplaceTektonPipelineTriggerPropertyOptions)`, func() {
		replaceTektonPipelineTriggerPropertyPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers/1bb892a1-2e04-4768-a369-b1159eace147/properties/debug-pipeline"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceTektonPipelineTriggerPropertyPath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}`)
				}))
			})
			It(`Invoke ReplaceTektonPipelineTriggerProperty successfully with retries`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())
				cdTektonPipelineService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceTektonPipelineTriggerPropertyOptions model
				replaceTektonPipelineTriggerPropertyOptionsModel := new(cdtektonpipelinev2.ReplaceTektonPipelineTriggerPropertyOptions)
				replaceTektonPipelineTriggerPropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				replaceTektonPipelineTriggerPropertyOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				replaceTektonPipelineTriggerPropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				replaceTektonPipelineTriggerPropertyOptionsModel.Name = core.StringPtr("prop1")
				replaceTektonPipelineTriggerPropertyOptionsModel.Type = core.StringPtr("text")
				replaceTektonPipelineTriggerPropertyOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				replaceTektonPipelineTriggerPropertyOptionsModel.Enum = []string{"testString"}
				replaceTektonPipelineTriggerPropertyOptionsModel.Path = core.StringPtr("testString")
				replaceTektonPipelineTriggerPropertyOptionsModel.Locked = core.BoolPtr(false)
				replaceTektonPipelineTriggerPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cdTektonPipelineService.ReplaceTektonPipelineTriggerPropertyWithContext(ctx, replaceTektonPipelineTriggerPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cdTektonPipelineService.DisableRetries()
				result, response, operationErr := cdTektonPipelineService.ReplaceTektonPipelineTriggerProperty(replaceTektonPipelineTriggerPropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cdTektonPipelineService.ReplaceTektonPipelineTriggerPropertyWithContext(ctx, replaceTektonPipelineTriggerPropertyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceTektonPipelineTriggerPropertyPath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"name": "Name", "value": "Value", "href": "Href", "enum": ["Enum"], "type": "secure", "path": "Path", "locked": true}`)
				}))
			})
			It(`Invoke ReplaceTektonPipelineTriggerProperty successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cdTektonPipelineService.ReplaceTektonPipelineTriggerProperty(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceTektonPipelineTriggerPropertyOptions model
				replaceTektonPipelineTriggerPropertyOptionsModel := new(cdtektonpipelinev2.ReplaceTektonPipelineTriggerPropertyOptions)
				replaceTektonPipelineTriggerPropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				replaceTektonPipelineTriggerPropertyOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				replaceTektonPipelineTriggerPropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				replaceTektonPipelineTriggerPropertyOptionsModel.Name = core.StringPtr("prop1")
				replaceTektonPipelineTriggerPropertyOptionsModel.Type = core.StringPtr("text")
				replaceTektonPipelineTriggerPropertyOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				replaceTektonPipelineTriggerPropertyOptionsModel.Enum = []string{"testString"}
				replaceTektonPipelineTriggerPropertyOptionsModel.Path = core.StringPtr("testString")
				replaceTektonPipelineTriggerPropertyOptionsModel.Locked = core.BoolPtr(false)
				replaceTektonPipelineTriggerPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cdTektonPipelineService.ReplaceTektonPipelineTriggerProperty(replaceTektonPipelineTriggerPropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceTektonPipelineTriggerProperty with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ReplaceTektonPipelineTriggerPropertyOptions model
				replaceTektonPipelineTriggerPropertyOptionsModel := new(cdtektonpipelinev2.ReplaceTektonPipelineTriggerPropertyOptions)
				replaceTektonPipelineTriggerPropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				replaceTektonPipelineTriggerPropertyOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				replaceTektonPipelineTriggerPropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				replaceTektonPipelineTriggerPropertyOptionsModel.Name = core.StringPtr("prop1")
				replaceTektonPipelineTriggerPropertyOptionsModel.Type = core.StringPtr("text")
				replaceTektonPipelineTriggerPropertyOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				replaceTektonPipelineTriggerPropertyOptionsModel.Enum = []string{"testString"}
				replaceTektonPipelineTriggerPropertyOptionsModel.Path = core.StringPtr("testString")
				replaceTektonPipelineTriggerPropertyOptionsModel.Locked = core.BoolPtr(false)
				replaceTektonPipelineTriggerPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cdTektonPipelineService.ReplaceTektonPipelineTriggerProperty(replaceTektonPipelineTriggerPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceTektonPipelineTriggerPropertyOptions model with no property values
				replaceTektonPipelineTriggerPropertyOptionsModelNew := new(cdtektonpipelinev2.ReplaceTektonPipelineTriggerPropertyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cdTektonPipelineService.ReplaceTektonPipelineTriggerProperty(replaceTektonPipelineTriggerPropertyOptionsModelNew)
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
			It(`Invoke ReplaceTektonPipelineTriggerProperty successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the ReplaceTektonPipelineTriggerPropertyOptions model
				replaceTektonPipelineTriggerPropertyOptionsModel := new(cdtektonpipelinev2.ReplaceTektonPipelineTriggerPropertyOptions)
				replaceTektonPipelineTriggerPropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				replaceTektonPipelineTriggerPropertyOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				replaceTektonPipelineTriggerPropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				replaceTektonPipelineTriggerPropertyOptionsModel.Name = core.StringPtr("prop1")
				replaceTektonPipelineTriggerPropertyOptionsModel.Type = core.StringPtr("text")
				replaceTektonPipelineTriggerPropertyOptionsModel.Value = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				replaceTektonPipelineTriggerPropertyOptionsModel.Enum = []string{"testString"}
				replaceTektonPipelineTriggerPropertyOptionsModel.Path = core.StringPtr("testString")
				replaceTektonPipelineTriggerPropertyOptionsModel.Locked = core.BoolPtr(false)
				replaceTektonPipelineTriggerPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cdTektonPipelineService.ReplaceTektonPipelineTriggerProperty(replaceTektonPipelineTriggerPropertyOptionsModel)
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
	Describe(`DeleteTektonPipelineTriggerProperty(deleteTektonPipelineTriggerPropertyOptions *DeleteTektonPipelineTriggerPropertyOptions)`, func() {
		deleteTektonPipelineTriggerPropertyPath := "/tekton_pipelines/94619026-912b-4d92-8f51-6c74f0692d90/triggers/1bb892a1-2e04-4768-a369-b1159eace147/properties/debug-pipeline"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTektonPipelineTriggerPropertyPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTektonPipelineTriggerProperty successfully`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := cdTektonPipelineService.DeleteTektonPipelineTriggerProperty(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTektonPipelineTriggerPropertyOptions model
				deleteTektonPipelineTriggerPropertyOptionsModel := new(cdtektonpipelinev2.DeleteTektonPipelineTriggerPropertyOptions)
				deleteTektonPipelineTriggerPropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelineTriggerPropertyOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				deleteTektonPipelineTriggerPropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				deleteTektonPipelineTriggerPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = cdTektonPipelineService.DeleteTektonPipelineTriggerProperty(deleteTektonPipelineTriggerPropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTektonPipelineTriggerProperty with error: Operation validation and request error`, func() {
				cdTektonPipelineService, serviceErr := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cdTektonPipelineService).ToNot(BeNil())

				// Construct an instance of the DeleteTektonPipelineTriggerPropertyOptions model
				deleteTektonPipelineTriggerPropertyOptionsModel := new(cdtektonpipelinev2.DeleteTektonPipelineTriggerPropertyOptions)
				deleteTektonPipelineTriggerPropertyOptionsModel.PipelineID = core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelineTriggerPropertyOptionsModel.TriggerID = core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")
				deleteTektonPipelineTriggerPropertyOptionsModel.PropertyName = core.StringPtr("debug-pipeline")
				deleteTektonPipelineTriggerPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cdTektonPipelineService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := cdTektonPipelineService.DeleteTektonPipelineTriggerProperty(deleteTektonPipelineTriggerPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTektonPipelineTriggerPropertyOptions model with no property values
				deleteTektonPipelineTriggerPropertyOptionsModelNew := new(cdtektonpipelinev2.DeleteTektonPipelineTriggerPropertyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = cdTektonPipelineService.DeleteTektonPipelineTriggerProperty(deleteTektonPipelineTriggerPropertyOptionsModelNew)
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
			cdTektonPipelineService, _ := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
				URL:           "http://cdtektonpipelinev2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCancelTektonPipelineRunOptions successfully`, func() {
				// Construct an instance of the CancelTektonPipelineRunOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				id := "94619026-912b-4d92-8f51-6c74f0692d90"
				cancelTektonPipelineRunOptionsModel := cdTektonPipelineService.NewCancelTektonPipelineRunOptions(pipelineID, id)
				cancelTektonPipelineRunOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				cancelTektonPipelineRunOptionsModel.SetID("94619026-912b-4d92-8f51-6c74f0692d90")
				cancelTektonPipelineRunOptionsModel.SetForce(true)
				cancelTektonPipelineRunOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(cancelTektonPipelineRunOptionsModel).ToNot(BeNil())
				Expect(cancelTektonPipelineRunOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(cancelTektonPipelineRunOptionsModel.ID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(cancelTektonPipelineRunOptionsModel.Force).To(Equal(core.BoolPtr(true)))
				Expect(cancelTektonPipelineRunOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTektonPipelineDefinitionOptions successfully`, func() {
				// Construct an instance of the Tool model
				toolModel := new(cdtektonpipelinev2.Tool)
				Expect(toolModel).ToNot(BeNil())
				toolModel.ID = core.StringPtr("testString")
				Expect(toolModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DefinitionSourceProperties model
				definitionSourcePropertiesModel := new(cdtektonpipelinev2.DefinitionSourceProperties)
				Expect(definitionSourcePropertiesModel).ToNot(BeNil())
				definitionSourcePropertiesModel.URL = core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")
				definitionSourcePropertiesModel.Branch = core.StringPtr("master")
				definitionSourcePropertiesModel.Tag = core.StringPtr("testString")
				definitionSourcePropertiesModel.Path = core.StringPtr(".tekton")
				definitionSourcePropertiesModel.Tool = toolModel
				Expect(definitionSourcePropertiesModel.URL).To(Equal(core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")))
				Expect(definitionSourcePropertiesModel.Branch).To(Equal(core.StringPtr("master")))
				Expect(definitionSourcePropertiesModel.Tag).To(Equal(core.StringPtr("testString")))
				Expect(definitionSourcePropertiesModel.Path).To(Equal(core.StringPtr(".tekton")))
				Expect(definitionSourcePropertiesModel.Tool).To(Equal(toolModel))

				// Construct an instance of the DefinitionSource model
				definitionSourceModel := new(cdtektonpipelinev2.DefinitionSource)
				Expect(definitionSourceModel).ToNot(BeNil())
				definitionSourceModel.Type = core.StringPtr("git")
				definitionSourceModel.Properties = definitionSourcePropertiesModel
				Expect(definitionSourceModel.Type).To(Equal(core.StringPtr("git")))
				Expect(definitionSourceModel.Properties).To(Equal(definitionSourcePropertiesModel))

				// Construct an instance of the CreateTektonPipelineDefinitionOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				var createTektonPipelineDefinitionOptionsSource *cdtektonpipelinev2.DefinitionSource = nil
				createTektonPipelineDefinitionOptionsModel := cdTektonPipelineService.NewCreateTektonPipelineDefinitionOptions(pipelineID, createTektonPipelineDefinitionOptionsSource)
				createTektonPipelineDefinitionOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineDefinitionOptionsModel.SetSource(definitionSourceModel)
				createTektonPipelineDefinitionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTektonPipelineDefinitionOptionsModel).ToNot(BeNil())
				Expect(createTektonPipelineDefinitionOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(createTektonPipelineDefinitionOptionsModel.Source).To(Equal(definitionSourceModel))
				Expect(createTektonPipelineDefinitionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTektonPipelineOptions successfully`, func() {
				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				Expect(workerIdentityModel).ToNot(BeNil())
				workerIdentityModel.ID = core.StringPtr("public")
				Expect(workerIdentityModel.ID).To(Equal(core.StringPtr("public")))

				// Construct an instance of the CreateTektonPipelineOptions model
				createTektonPipelineOptionsID := "94619026-912b-4d92-8f51-6c74f0692d90"
				createTektonPipelineOptionsModel := cdTektonPipelineService.NewCreateTektonPipelineOptions(createTektonPipelineOptionsID)
				createTektonPipelineOptionsModel.SetID("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineOptionsModel.SetNextBuildNumber(int64(1))
				createTektonPipelineOptionsModel.SetEnableNotifications(false)
				createTektonPipelineOptionsModel.SetEnablePartialCloning(false)
				createTektonPipelineOptionsModel.SetWorker(workerIdentityModel)
				createTektonPipelineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTektonPipelineOptionsModel).ToNot(BeNil())
				Expect(createTektonPipelineOptionsModel.ID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(createTektonPipelineOptionsModel.NextBuildNumber).To(Equal(core.Int64Ptr(int64(1))))
				Expect(createTektonPipelineOptionsModel.EnableNotifications).To(Equal(core.BoolPtr(false)))
				Expect(createTektonPipelineOptionsModel.EnablePartialCloning).To(Equal(core.BoolPtr(false)))
				Expect(createTektonPipelineOptionsModel.Worker).To(Equal(workerIdentityModel))
				Expect(createTektonPipelineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTektonPipelinePropertiesOptions successfully`, func() {
				// Construct an instance of the CreateTektonPipelinePropertiesOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				createTektonPipelinePropertiesOptionsName := "prop1"
				createTektonPipelinePropertiesOptionsType := "text"
				createTektonPipelinePropertiesOptionsModel := cdTektonPipelineService.NewCreateTektonPipelinePropertiesOptions(pipelineID, createTektonPipelinePropertiesOptionsName, createTektonPipelinePropertiesOptionsType)
				createTektonPipelinePropertiesOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelinePropertiesOptionsModel.SetName("prop1")
				createTektonPipelinePropertiesOptionsModel.SetType("text")
				createTektonPipelinePropertiesOptionsModel.SetValue("https://github.com/open-toolchain/hello-tekton.git")
				createTektonPipelinePropertiesOptionsModel.SetEnum([]string{"testString"})
				createTektonPipelinePropertiesOptionsModel.SetLocked(false)
				createTektonPipelinePropertiesOptionsModel.SetPath("testString")
				createTektonPipelinePropertiesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTektonPipelinePropertiesOptionsModel).ToNot(BeNil())
				Expect(createTektonPipelinePropertiesOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(createTektonPipelinePropertiesOptionsModel.Name).To(Equal(core.StringPtr("prop1")))
				Expect(createTektonPipelinePropertiesOptionsModel.Type).To(Equal(core.StringPtr("text")))
				Expect(createTektonPipelinePropertiesOptionsModel.Value).To(Equal(core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")))
				Expect(createTektonPipelinePropertiesOptionsModel.Enum).To(Equal([]string{"testString"}))
				Expect(createTektonPipelinePropertiesOptionsModel.Locked).To(Equal(core.BoolPtr(false)))
				Expect(createTektonPipelinePropertiesOptionsModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(createTektonPipelinePropertiesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTektonPipelineRunOptions successfully`, func() {
				// Construct an instance of the PipelineRunTrigger model
				pipelineRunTriggerModel := new(cdtektonpipelinev2.PipelineRunTrigger)
				Expect(pipelineRunTriggerModel).ToNot(BeNil())
				pipelineRunTriggerModel.Name = core.StringPtr("Manual Trigger 1")
				pipelineRunTriggerModel.Properties = map[string]interface{}{"anyKey": "anyValue"}
				pipelineRunTriggerModel.SecureProperties = map[string]interface{}{"anyKey": "anyValue"}
				pipelineRunTriggerModel.HeadersVar = map[string]interface{}{"anyKey": "anyValue"}
				pipelineRunTriggerModel.Body = map[string]interface{}{"anyKey": "anyValue"}
				Expect(pipelineRunTriggerModel.Name).To(Equal(core.StringPtr("Manual Trigger 1")))
				Expect(pipelineRunTriggerModel.Properties).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(pipelineRunTriggerModel.SecureProperties).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(pipelineRunTriggerModel.HeadersVar).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(pipelineRunTriggerModel.Body).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))

				// Construct an instance of the CreateTektonPipelineRunOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				createTektonPipelineRunOptionsModel := cdTektonPipelineService.NewCreateTektonPipelineRunOptions(pipelineID)
				createTektonPipelineRunOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineRunOptionsModel.SetDescription("My custom manual PipelineRun")
				createTektonPipelineRunOptionsModel.SetTriggerName("testString")
				createTektonPipelineRunOptionsModel.SetTriggerProperties(map[string]interface{}{"anyKey": "anyValue"})
				createTektonPipelineRunOptionsModel.SetSecureTriggerProperties(map[string]interface{}{"anyKey": "anyValue"})
				createTektonPipelineRunOptionsModel.SetTriggerHeaders(map[string]interface{}{"anyKey": "anyValue"})
				createTektonPipelineRunOptionsModel.SetTriggerBody(map[string]interface{}{"anyKey": "anyValue"})
				createTektonPipelineRunOptionsModel.SetTrigger(pipelineRunTriggerModel)
				createTektonPipelineRunOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTektonPipelineRunOptionsModel).ToNot(BeNil())
				Expect(createTektonPipelineRunOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(createTektonPipelineRunOptionsModel.Description).To(Equal(core.StringPtr("My custom manual PipelineRun")))
				Expect(createTektonPipelineRunOptionsModel.TriggerName).To(Equal(core.StringPtr("testString")))
				Expect(createTektonPipelineRunOptionsModel.TriggerProperties).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(createTektonPipelineRunOptionsModel.SecureTriggerProperties).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(createTektonPipelineRunOptionsModel.TriggerHeaders).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(createTektonPipelineRunOptionsModel.TriggerBody).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(createTektonPipelineRunOptionsModel.Trigger).To(Equal(pipelineRunTriggerModel))
				Expect(createTektonPipelineRunOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTektonPipelineTriggerOptions successfully`, func() {
				// Construct an instance of the WorkerIdentity model
				workerIdentityModel := new(cdtektonpipelinev2.WorkerIdentity)
				Expect(workerIdentityModel).ToNot(BeNil())
				workerIdentityModel.ID = core.StringPtr("public")
				Expect(workerIdentityModel.ID).To(Equal(core.StringPtr("public")))

				// Construct an instance of the GenericSecret model
				genericSecretModel := new(cdtektonpipelinev2.GenericSecret)
				Expect(genericSecretModel).ToNot(BeNil())
				genericSecretModel.Type = core.StringPtr("token_matches")
				genericSecretModel.Value = core.StringPtr("testString")
				genericSecretModel.Source = core.StringPtr("header")
				genericSecretModel.KeyName = core.StringPtr("testString")
				genericSecretModel.Algorithm = core.StringPtr("md4")
				Expect(genericSecretModel.Type).To(Equal(core.StringPtr("token_matches")))
				Expect(genericSecretModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(genericSecretModel.Source).To(Equal(core.StringPtr("header")))
				Expect(genericSecretModel.KeyName).To(Equal(core.StringPtr("testString")))
				Expect(genericSecretModel.Algorithm).To(Equal(core.StringPtr("md4")))

				// Construct an instance of the TriggerSourcePropertiesPrototype model
				triggerSourcePropertiesPrototypeModel := new(cdtektonpipelinev2.TriggerSourcePropertiesPrototype)
				Expect(triggerSourcePropertiesPrototypeModel).ToNot(BeNil())
				triggerSourcePropertiesPrototypeModel.URL = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Branch = core.StringPtr("testString")
				triggerSourcePropertiesPrototypeModel.Pattern = core.StringPtr("testString")
				Expect(triggerSourcePropertiesPrototypeModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(triggerSourcePropertiesPrototypeModel.Branch).To(Equal(core.StringPtr("testString")))
				Expect(triggerSourcePropertiesPrototypeModel.Pattern).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TriggerSourcePrototype model
				triggerSourcePrototypeModel := new(cdtektonpipelinev2.TriggerSourcePrototype)
				Expect(triggerSourcePrototypeModel).ToNot(BeNil())
				triggerSourcePrototypeModel.Type = core.StringPtr("testString")
				triggerSourcePrototypeModel.Properties = triggerSourcePropertiesPrototypeModel
				Expect(triggerSourcePrototypeModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(triggerSourcePrototypeModel.Properties).To(Equal(triggerSourcePropertiesPrototypeModel))

				// Construct an instance of the CreateTektonPipelineTriggerOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				createTektonPipelineTriggerOptionsType := "manual"
				createTektonPipelineTriggerOptionsName := "Manual Trigger"
				createTektonPipelineTriggerOptionsEventListener := "pr-listener"
				createTektonPipelineTriggerOptionsModel := cdTektonPipelineService.NewCreateTektonPipelineTriggerOptions(pipelineID, createTektonPipelineTriggerOptionsType, createTektonPipelineTriggerOptionsName, createTektonPipelineTriggerOptionsEventListener)
				createTektonPipelineTriggerOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineTriggerOptionsModel.SetType("manual")
				createTektonPipelineTriggerOptionsModel.SetName("Manual Trigger")
				createTektonPipelineTriggerOptionsModel.SetEventListener("pr-listener")
				createTektonPipelineTriggerOptionsModel.SetTags([]string{"testString"})
				createTektonPipelineTriggerOptionsModel.SetWorker(workerIdentityModel)
				createTektonPipelineTriggerOptionsModel.SetMaxConcurrentRuns(int64(3))
				createTektonPipelineTriggerOptionsModel.SetLimitWaitingRuns(false)
				createTektonPipelineTriggerOptionsModel.SetEnabled(true)
				createTektonPipelineTriggerOptionsModel.SetSecret(genericSecretModel)
				createTektonPipelineTriggerOptionsModel.SetCron("testString")
				createTektonPipelineTriggerOptionsModel.SetTimezone("testString")
				createTektonPipelineTriggerOptionsModel.SetSource(triggerSourcePrototypeModel)
				createTektonPipelineTriggerOptionsModel.SetEvents([]string{"push"})
				createTektonPipelineTriggerOptionsModel.SetFilter("testString")
				createTektonPipelineTriggerOptionsModel.SetFavorite(false)
				createTektonPipelineTriggerOptionsModel.SetEnableEventsFromForks(false)
				createTektonPipelineTriggerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTektonPipelineTriggerOptionsModel).ToNot(BeNil())
				Expect(createTektonPipelineTriggerOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(createTektonPipelineTriggerOptionsModel.Type).To(Equal(core.StringPtr("manual")))
				Expect(createTektonPipelineTriggerOptionsModel.Name).To(Equal(core.StringPtr("Manual Trigger")))
				Expect(createTektonPipelineTriggerOptionsModel.EventListener).To(Equal(core.StringPtr("pr-listener")))
				Expect(createTektonPipelineTriggerOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createTektonPipelineTriggerOptionsModel.Worker).To(Equal(workerIdentityModel))
				Expect(createTektonPipelineTriggerOptionsModel.MaxConcurrentRuns).To(Equal(core.Int64Ptr(int64(3))))
				Expect(createTektonPipelineTriggerOptionsModel.LimitWaitingRuns).To(Equal(core.BoolPtr(false)))
				Expect(createTektonPipelineTriggerOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(createTektonPipelineTriggerOptionsModel.Secret).To(Equal(genericSecretModel))
				Expect(createTektonPipelineTriggerOptionsModel.Cron).To(Equal(core.StringPtr("testString")))
				Expect(createTektonPipelineTriggerOptionsModel.Timezone).To(Equal(core.StringPtr("testString")))
				Expect(createTektonPipelineTriggerOptionsModel.Source).To(Equal(triggerSourcePrototypeModel))
				Expect(createTektonPipelineTriggerOptionsModel.Events).To(Equal([]string{"push"}))
				Expect(createTektonPipelineTriggerOptionsModel.Filter).To(Equal(core.StringPtr("testString")))
				Expect(createTektonPipelineTriggerOptionsModel.Favorite).To(Equal(core.BoolPtr(false)))
				Expect(createTektonPipelineTriggerOptionsModel.EnableEventsFromForks).To(Equal(core.BoolPtr(false)))
				Expect(createTektonPipelineTriggerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTektonPipelineTriggerPropertiesOptions successfully`, func() {
				// Construct an instance of the CreateTektonPipelineTriggerPropertiesOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				triggerID := "1bb892a1-2e04-4768-a369-b1159eace147"
				createTektonPipelineTriggerPropertiesOptionsName := "prop1"
				createTektonPipelineTriggerPropertiesOptionsType := "text"
				createTektonPipelineTriggerPropertiesOptionsModel := cdTektonPipelineService.NewCreateTektonPipelineTriggerPropertiesOptions(pipelineID, triggerID, createTektonPipelineTriggerPropertiesOptionsName, createTektonPipelineTriggerPropertiesOptionsType)
				createTektonPipelineTriggerPropertiesOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				createTektonPipelineTriggerPropertiesOptionsModel.SetTriggerID("1bb892a1-2e04-4768-a369-b1159eace147")
				createTektonPipelineTriggerPropertiesOptionsModel.SetName("prop1")
				createTektonPipelineTriggerPropertiesOptionsModel.SetType("text")
				createTektonPipelineTriggerPropertiesOptionsModel.SetValue("https://github.com/open-toolchain/hello-tekton.git")
				createTektonPipelineTriggerPropertiesOptionsModel.SetEnum([]string{"testString"})
				createTektonPipelineTriggerPropertiesOptionsModel.SetPath("testString")
				createTektonPipelineTriggerPropertiesOptionsModel.SetLocked(false)
				createTektonPipelineTriggerPropertiesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTektonPipelineTriggerPropertiesOptionsModel).ToNot(BeNil())
				Expect(createTektonPipelineTriggerPropertiesOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(createTektonPipelineTriggerPropertiesOptionsModel.TriggerID).To(Equal(core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")))
				Expect(createTektonPipelineTriggerPropertiesOptionsModel.Name).To(Equal(core.StringPtr("prop1")))
				Expect(createTektonPipelineTriggerPropertiesOptionsModel.Type).To(Equal(core.StringPtr("text")))
				Expect(createTektonPipelineTriggerPropertiesOptionsModel.Value).To(Equal(core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")))
				Expect(createTektonPipelineTriggerPropertiesOptionsModel.Enum).To(Equal([]string{"testString"}))
				Expect(createTektonPipelineTriggerPropertiesOptionsModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(createTektonPipelineTriggerPropertiesOptionsModel.Locked).To(Equal(core.BoolPtr(false)))
				Expect(createTektonPipelineTriggerPropertiesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDefinitionSource successfully`, func() {
				typeVar := "testString"
				var properties *cdtektonpipelinev2.DefinitionSourceProperties = nil
				_, err := cdTektonPipelineService.NewDefinitionSource(typeVar, properties)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewDefinitionSourceProperties successfully`, func() {
				url := "testString"
				path := "testString"
				_model, err := cdTektonPipelineService.NewDefinitionSourceProperties(url, path)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDeleteTektonPipelineDefinitionOptions successfully`, func() {
				// Construct an instance of the DeleteTektonPipelineDefinitionOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				definitionID := "94299034-d45f-4e9a-8ed5-6bd5c7bb7ada"
				deleteTektonPipelineDefinitionOptionsModel := cdTektonPipelineService.NewDeleteTektonPipelineDefinitionOptions(pipelineID, definitionID)
				deleteTektonPipelineDefinitionOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelineDefinitionOptionsModel.SetDefinitionID("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada")
				deleteTektonPipelineDefinitionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTektonPipelineDefinitionOptionsModel).ToNot(BeNil())
				Expect(deleteTektonPipelineDefinitionOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(deleteTektonPipelineDefinitionOptionsModel.DefinitionID).To(Equal(core.StringPtr("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada")))
				Expect(deleteTektonPipelineDefinitionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTektonPipelineOptions successfully`, func() {
				// Construct an instance of the DeleteTektonPipelineOptions model
				id := "94619026-912b-4d92-8f51-6c74f0692d90"
				deleteTektonPipelineOptionsModel := cdTektonPipelineService.NewDeleteTektonPipelineOptions(id)
				deleteTektonPipelineOptionsModel.SetID("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTektonPipelineOptionsModel).ToNot(BeNil())
				Expect(deleteTektonPipelineOptionsModel.ID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(deleteTektonPipelineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTektonPipelinePropertyOptions successfully`, func() {
				// Construct an instance of the DeleteTektonPipelinePropertyOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				propertyName := "debug-pipeline"
				deleteTektonPipelinePropertyOptionsModel := cdTektonPipelineService.NewDeleteTektonPipelinePropertyOptions(pipelineID, propertyName)
				deleteTektonPipelinePropertyOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelinePropertyOptionsModel.SetPropertyName("debug-pipeline")
				deleteTektonPipelinePropertyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTektonPipelinePropertyOptionsModel).ToNot(BeNil())
				Expect(deleteTektonPipelinePropertyOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(deleteTektonPipelinePropertyOptionsModel.PropertyName).To(Equal(core.StringPtr("debug-pipeline")))
				Expect(deleteTektonPipelinePropertyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTektonPipelineRunOptions successfully`, func() {
				// Construct an instance of the DeleteTektonPipelineRunOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				id := "94619026-912b-4d92-8f51-6c74f0692d90"
				deleteTektonPipelineRunOptionsModel := cdTektonPipelineService.NewDeleteTektonPipelineRunOptions(pipelineID, id)
				deleteTektonPipelineRunOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelineRunOptionsModel.SetID("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelineRunOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTektonPipelineRunOptionsModel).ToNot(BeNil())
				Expect(deleteTektonPipelineRunOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(deleteTektonPipelineRunOptionsModel.ID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(deleteTektonPipelineRunOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTektonPipelineTriggerOptions successfully`, func() {
				// Construct an instance of the DeleteTektonPipelineTriggerOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				triggerID := "1bb892a1-2e04-4768-a369-b1159eace147"
				deleteTektonPipelineTriggerOptionsModel := cdTektonPipelineService.NewDeleteTektonPipelineTriggerOptions(pipelineID, triggerID)
				deleteTektonPipelineTriggerOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelineTriggerOptionsModel.SetTriggerID("1bb892a1-2e04-4768-a369-b1159eace147")
				deleteTektonPipelineTriggerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTektonPipelineTriggerOptionsModel).ToNot(BeNil())
				Expect(deleteTektonPipelineTriggerOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(deleteTektonPipelineTriggerOptionsModel.TriggerID).To(Equal(core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")))
				Expect(deleteTektonPipelineTriggerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTektonPipelineTriggerPropertyOptions successfully`, func() {
				// Construct an instance of the DeleteTektonPipelineTriggerPropertyOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				triggerID := "1bb892a1-2e04-4768-a369-b1159eace147"
				propertyName := "debug-pipeline"
				deleteTektonPipelineTriggerPropertyOptionsModel := cdTektonPipelineService.NewDeleteTektonPipelineTriggerPropertyOptions(pipelineID, triggerID, propertyName)
				deleteTektonPipelineTriggerPropertyOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				deleteTektonPipelineTriggerPropertyOptionsModel.SetTriggerID("1bb892a1-2e04-4768-a369-b1159eace147")
				deleteTektonPipelineTriggerPropertyOptionsModel.SetPropertyName("debug-pipeline")
				deleteTektonPipelineTriggerPropertyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTektonPipelineTriggerPropertyOptionsModel).ToNot(BeNil())
				Expect(deleteTektonPipelineTriggerPropertyOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(deleteTektonPipelineTriggerPropertyOptionsModel.TriggerID).To(Equal(core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")))
				Expect(deleteTektonPipelineTriggerPropertyOptionsModel.PropertyName).To(Equal(core.StringPtr("debug-pipeline")))
				Expect(deleteTektonPipelineTriggerPropertyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDuplicateTektonPipelineTriggerOptions successfully`, func() {
				// Construct an instance of the DuplicateTektonPipelineTriggerOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				sourceTriggerID := "1bb892a1-2e04-4768-a369-b1159eace147"
				duplicateTektonPipelineTriggerOptionsName := "triggerName"
				duplicateTektonPipelineTriggerOptionsModel := cdTektonPipelineService.NewDuplicateTektonPipelineTriggerOptions(pipelineID, sourceTriggerID, duplicateTektonPipelineTriggerOptionsName)
				duplicateTektonPipelineTriggerOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				duplicateTektonPipelineTriggerOptionsModel.SetSourceTriggerID("1bb892a1-2e04-4768-a369-b1159eace147")
				duplicateTektonPipelineTriggerOptionsModel.SetName("triggerName")
				duplicateTektonPipelineTriggerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(duplicateTektonPipelineTriggerOptionsModel).ToNot(BeNil())
				Expect(duplicateTektonPipelineTriggerOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(duplicateTektonPipelineTriggerOptionsModel.SourceTriggerID).To(Equal(core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")))
				Expect(duplicateTektonPipelineTriggerOptionsModel.Name).To(Equal(core.StringPtr("triggerName")))
				Expect(duplicateTektonPipelineTriggerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTektonPipelineDefinitionOptions successfully`, func() {
				// Construct an instance of the GetTektonPipelineDefinitionOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				definitionID := "94299034-d45f-4e9a-8ed5-6bd5c7bb7ada"
				getTektonPipelineDefinitionOptionsModel := cdTektonPipelineService.NewGetTektonPipelineDefinitionOptions(pipelineID, definitionID)
				getTektonPipelineDefinitionOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineDefinitionOptionsModel.SetDefinitionID("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada")
				getTektonPipelineDefinitionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTektonPipelineDefinitionOptionsModel).ToNot(BeNil())
				Expect(getTektonPipelineDefinitionOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(getTektonPipelineDefinitionOptionsModel.DefinitionID).To(Equal(core.StringPtr("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada")))
				Expect(getTektonPipelineDefinitionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTektonPipelineOptions successfully`, func() {
				// Construct an instance of the GetTektonPipelineOptions model
				id := "94619026-912b-4d92-8f51-6c74f0692d90"
				getTektonPipelineOptionsModel := cdTektonPipelineService.NewGetTektonPipelineOptions(id)
				getTektonPipelineOptionsModel.SetID("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTektonPipelineOptionsModel).ToNot(BeNil())
				Expect(getTektonPipelineOptionsModel.ID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(getTektonPipelineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTektonPipelinePropertyOptions successfully`, func() {
				// Construct an instance of the GetTektonPipelinePropertyOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				propertyName := "debug-pipeline"
				getTektonPipelinePropertyOptionsModel := cdTektonPipelineService.NewGetTektonPipelinePropertyOptions(pipelineID, propertyName)
				getTektonPipelinePropertyOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelinePropertyOptionsModel.SetPropertyName("debug-pipeline")
				getTektonPipelinePropertyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTektonPipelinePropertyOptionsModel).ToNot(BeNil())
				Expect(getTektonPipelinePropertyOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(getTektonPipelinePropertyOptionsModel.PropertyName).To(Equal(core.StringPtr("debug-pipeline")))
				Expect(getTektonPipelinePropertyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTektonPipelineRunLogContentOptions successfully`, func() {
				// Construct an instance of the GetTektonPipelineRunLogContentOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				pipelineRunID := "bf4b3abd-0c93-416b-911e-9cf42f1a1085"
				id := "94619026-912b-4d92-8f51-6c74f0692d90"
				getTektonPipelineRunLogContentOptionsModel := cdTektonPipelineService.NewGetTektonPipelineRunLogContentOptions(pipelineID, pipelineRunID, id)
				getTektonPipelineRunLogContentOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogContentOptionsModel.SetPipelineRunID("bf4b3abd-0c93-416b-911e-9cf42f1a1085")
				getTektonPipelineRunLogContentOptionsModel.SetID("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogContentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTektonPipelineRunLogContentOptionsModel).ToNot(BeNil())
				Expect(getTektonPipelineRunLogContentOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(getTektonPipelineRunLogContentOptionsModel.PipelineRunID).To(Equal(core.StringPtr("bf4b3abd-0c93-416b-911e-9cf42f1a1085")))
				Expect(getTektonPipelineRunLogContentOptionsModel.ID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(getTektonPipelineRunLogContentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTektonPipelineRunLogsOptions successfully`, func() {
				// Construct an instance of the GetTektonPipelineRunLogsOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				id := "94619026-912b-4d92-8f51-6c74f0692d90"
				getTektonPipelineRunLogsOptionsModel := cdTektonPipelineService.NewGetTektonPipelineRunLogsOptions(pipelineID, id)
				getTektonPipelineRunLogsOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogsOptionsModel.SetID("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunLogsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTektonPipelineRunLogsOptionsModel).ToNot(BeNil())
				Expect(getTektonPipelineRunLogsOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(getTektonPipelineRunLogsOptionsModel.ID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(getTektonPipelineRunLogsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTektonPipelineRunOptions successfully`, func() {
				// Construct an instance of the GetTektonPipelineRunOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				id := "94619026-912b-4d92-8f51-6c74f0692d90"
				getTektonPipelineRunOptionsModel := cdTektonPipelineService.NewGetTektonPipelineRunOptions(pipelineID, id)
				getTektonPipelineRunOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunOptionsModel.SetID("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineRunOptionsModel.SetIncludes("definitions")
				getTektonPipelineRunOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTektonPipelineRunOptionsModel).ToNot(BeNil())
				Expect(getTektonPipelineRunOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(getTektonPipelineRunOptionsModel.ID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(getTektonPipelineRunOptionsModel.Includes).To(Equal(core.StringPtr("definitions")))
				Expect(getTektonPipelineRunOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTektonPipelineTriggerOptions successfully`, func() {
				// Construct an instance of the GetTektonPipelineTriggerOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				triggerID := "1bb892a1-2e04-4768-a369-b1159eace147"
				getTektonPipelineTriggerOptionsModel := cdTektonPipelineService.NewGetTektonPipelineTriggerOptions(pipelineID, triggerID)
				getTektonPipelineTriggerOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineTriggerOptionsModel.SetTriggerID("1bb892a1-2e04-4768-a369-b1159eace147")
				getTektonPipelineTriggerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTektonPipelineTriggerOptionsModel).ToNot(BeNil())
				Expect(getTektonPipelineTriggerOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(getTektonPipelineTriggerOptionsModel.TriggerID).To(Equal(core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")))
				Expect(getTektonPipelineTriggerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTektonPipelineTriggerPropertyOptions successfully`, func() {
				// Construct an instance of the GetTektonPipelineTriggerPropertyOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				triggerID := "1bb892a1-2e04-4768-a369-b1159eace147"
				propertyName := "debug-pipeline"
				getTektonPipelineTriggerPropertyOptionsModel := cdTektonPipelineService.NewGetTektonPipelineTriggerPropertyOptions(pipelineID, triggerID, propertyName)
				getTektonPipelineTriggerPropertyOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				getTektonPipelineTriggerPropertyOptionsModel.SetTriggerID("1bb892a1-2e04-4768-a369-b1159eace147")
				getTektonPipelineTriggerPropertyOptionsModel.SetPropertyName("debug-pipeline")
				getTektonPipelineTriggerPropertyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTektonPipelineTriggerPropertyOptionsModel).ToNot(BeNil())
				Expect(getTektonPipelineTriggerPropertyOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(getTektonPipelineTriggerPropertyOptionsModel.TriggerID).To(Equal(core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")))
				Expect(getTektonPipelineTriggerPropertyOptionsModel.PropertyName).To(Equal(core.StringPtr("debug-pipeline")))
				Expect(getTektonPipelineTriggerPropertyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTektonPipelineDefinitionsOptions successfully`, func() {
				// Construct an instance of the ListTektonPipelineDefinitionsOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				listTektonPipelineDefinitionsOptionsModel := cdTektonPipelineService.NewListTektonPipelineDefinitionsOptions(pipelineID)
				listTektonPipelineDefinitionsOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineDefinitionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTektonPipelineDefinitionsOptionsModel).ToNot(BeNil())
				Expect(listTektonPipelineDefinitionsOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(listTektonPipelineDefinitionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTektonPipelinePropertiesOptions successfully`, func() {
				// Construct an instance of the ListTektonPipelinePropertiesOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				listTektonPipelinePropertiesOptionsModel := cdTektonPipelineService.NewListTektonPipelinePropertiesOptions(pipelineID)
				listTektonPipelinePropertiesOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelinePropertiesOptionsModel.SetName("prod")
				listTektonPipelinePropertiesOptionsModel.SetType([]string{"secure", "text"})
				listTektonPipelinePropertiesOptionsModel.SetSort("name")
				listTektonPipelinePropertiesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTektonPipelinePropertiesOptionsModel).ToNot(BeNil())
				Expect(listTektonPipelinePropertiesOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(listTektonPipelinePropertiesOptionsModel.Name).To(Equal(core.StringPtr("prod")))
				Expect(listTektonPipelinePropertiesOptionsModel.Type).To(Equal([]string{"secure", "text"}))
				Expect(listTektonPipelinePropertiesOptionsModel.Sort).To(Equal(core.StringPtr("name")))
				Expect(listTektonPipelinePropertiesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTektonPipelineRunsOptions successfully`, func() {
				// Construct an instance of the ListTektonPipelineRunsOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				listTektonPipelineRunsOptionsModel := cdTektonPipelineService.NewListTektonPipelineRunsOptions(pipelineID)
				listTektonPipelineRunsOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineRunsOptionsModel.SetStart("testString")
				listTektonPipelineRunsOptionsModel.SetLimit(int64(10))
				listTektonPipelineRunsOptionsModel.SetStatus("succeeded")
				listTektonPipelineRunsOptionsModel.SetTriggerName("manual-trigger")
				listTektonPipelineRunsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTektonPipelineRunsOptionsModel).ToNot(BeNil())
				Expect(listTektonPipelineRunsOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(listTektonPipelineRunsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listTektonPipelineRunsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listTektonPipelineRunsOptionsModel.Status).To(Equal(core.StringPtr("succeeded")))
				Expect(listTektonPipelineRunsOptionsModel.TriggerName).To(Equal(core.StringPtr("manual-trigger")))
				Expect(listTektonPipelineRunsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTektonPipelineTriggerPropertiesOptions successfully`, func() {
				// Construct an instance of the ListTektonPipelineTriggerPropertiesOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				triggerID := "1bb892a1-2e04-4768-a369-b1159eace147"
				listTektonPipelineTriggerPropertiesOptionsModel := cdTektonPipelineService.NewListTektonPipelineTriggerPropertiesOptions(pipelineID, triggerID)
				listTektonPipelineTriggerPropertiesOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineTriggerPropertiesOptionsModel.SetTriggerID("1bb892a1-2e04-4768-a369-b1159eace147")
				listTektonPipelineTriggerPropertiesOptionsModel.SetName("prod")
				listTektonPipelineTriggerPropertiesOptionsModel.SetType("secure,text")
				listTektonPipelineTriggerPropertiesOptionsModel.SetSort("name")
				listTektonPipelineTriggerPropertiesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTektonPipelineTriggerPropertiesOptionsModel).ToNot(BeNil())
				Expect(listTektonPipelineTriggerPropertiesOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(listTektonPipelineTriggerPropertiesOptionsModel.TriggerID).To(Equal(core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")))
				Expect(listTektonPipelineTriggerPropertiesOptionsModel.Name).To(Equal(core.StringPtr("prod")))
				Expect(listTektonPipelineTriggerPropertiesOptionsModel.Type).To(Equal(core.StringPtr("secure,text")))
				Expect(listTektonPipelineTriggerPropertiesOptionsModel.Sort).To(Equal(core.StringPtr("name")))
				Expect(listTektonPipelineTriggerPropertiesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTektonPipelineTriggersOptions successfully`, func() {
				// Construct an instance of the ListTektonPipelineTriggersOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				listTektonPipelineTriggersOptionsModel := cdTektonPipelineService.NewListTektonPipelineTriggersOptions(pipelineID)
				listTektonPipelineTriggersOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				listTektonPipelineTriggersOptionsModel.SetType("manual,scm")
				listTektonPipelineTriggersOptionsModel.SetName("testString")
				listTektonPipelineTriggersOptionsModel.SetEventListener("testString")
				listTektonPipelineTriggersOptionsModel.SetWorkerID("testString")
				listTektonPipelineTriggersOptionsModel.SetWorkerName("testString")
				listTektonPipelineTriggersOptionsModel.SetDisabled("true")
				listTektonPipelineTriggersOptionsModel.SetTags("tag1,tag2")
				listTektonPipelineTriggersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTektonPipelineTriggersOptionsModel).ToNot(BeNil())
				Expect(listTektonPipelineTriggersOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(listTektonPipelineTriggersOptionsModel.Type).To(Equal(core.StringPtr("manual,scm")))
				Expect(listTektonPipelineTriggersOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listTektonPipelineTriggersOptionsModel.EventListener).To(Equal(core.StringPtr("testString")))
				Expect(listTektonPipelineTriggersOptionsModel.WorkerID).To(Equal(core.StringPtr("testString")))
				Expect(listTektonPipelineTriggersOptionsModel.WorkerName).To(Equal(core.StringPtr("testString")))
				Expect(listTektonPipelineTriggersOptionsModel.Disabled).To(Equal(core.StringPtr("true")))
				Expect(listTektonPipelineTriggersOptionsModel.Tags).To(Equal(core.StringPtr("tag1,tag2")))
				Expect(listTektonPipelineTriggersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPipelineRunTrigger successfully`, func() {
				name := "start-deploy"
				_model, err := cdTektonPipelineService.NewPipelineRunTrigger(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewReplaceTektonPipelineDefinitionOptions successfully`, func() {
				// Construct an instance of the Tool model
				toolModel := new(cdtektonpipelinev2.Tool)
				Expect(toolModel).ToNot(BeNil())
				toolModel.ID = core.StringPtr("testString")
				Expect(toolModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DefinitionSourceProperties model
				definitionSourcePropertiesModel := new(cdtektonpipelinev2.DefinitionSourceProperties)
				Expect(definitionSourcePropertiesModel).ToNot(BeNil())
				definitionSourcePropertiesModel.URL = core.StringPtr("testString")
				definitionSourcePropertiesModel.Branch = core.StringPtr("testString")
				definitionSourcePropertiesModel.Tag = core.StringPtr("testString")
				definitionSourcePropertiesModel.Path = core.StringPtr("testString")
				definitionSourcePropertiesModel.Tool = toolModel
				Expect(definitionSourcePropertiesModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(definitionSourcePropertiesModel.Branch).To(Equal(core.StringPtr("testString")))
				Expect(definitionSourcePropertiesModel.Tag).To(Equal(core.StringPtr("testString")))
				Expect(definitionSourcePropertiesModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(definitionSourcePropertiesModel.Tool).To(Equal(toolModel))

				// Construct an instance of the DefinitionSource model
				definitionSourceModel := new(cdtektonpipelinev2.DefinitionSource)
				Expect(definitionSourceModel).ToNot(BeNil())
				definitionSourceModel.Type = core.StringPtr("testString")
				definitionSourceModel.Properties = definitionSourcePropertiesModel
				Expect(definitionSourceModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(definitionSourceModel.Properties).To(Equal(definitionSourcePropertiesModel))

				// Construct an instance of the ReplaceTektonPipelineDefinitionOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				definitionID := "94299034-d45f-4e9a-8ed5-6bd5c7bb7ada"
				var replaceTektonPipelineDefinitionOptionsSource *cdtektonpipelinev2.DefinitionSource = nil
				replaceTektonPipelineDefinitionOptionsModel := cdTektonPipelineService.NewReplaceTektonPipelineDefinitionOptions(pipelineID, definitionID, replaceTektonPipelineDefinitionOptionsSource)
				replaceTektonPipelineDefinitionOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				replaceTektonPipelineDefinitionOptionsModel.SetDefinitionID("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada")
				replaceTektonPipelineDefinitionOptionsModel.SetSource(definitionSourceModel)
				replaceTektonPipelineDefinitionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceTektonPipelineDefinitionOptionsModel).ToNot(BeNil())
				Expect(replaceTektonPipelineDefinitionOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(replaceTektonPipelineDefinitionOptionsModel.DefinitionID).To(Equal(core.StringPtr("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada")))
				Expect(replaceTektonPipelineDefinitionOptionsModel.Source).To(Equal(definitionSourceModel))
				Expect(replaceTektonPipelineDefinitionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceTektonPipelinePropertyOptions successfully`, func() {
				// Construct an instance of the ReplaceTektonPipelinePropertyOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				propertyName := "debug-pipeline"
				replaceTektonPipelinePropertyOptionsName := "prop1"
				replaceTektonPipelinePropertyOptionsType := "text"
				replaceTektonPipelinePropertyOptionsModel := cdTektonPipelineService.NewReplaceTektonPipelinePropertyOptions(pipelineID, propertyName, replaceTektonPipelinePropertyOptionsName, replaceTektonPipelinePropertyOptionsType)
				replaceTektonPipelinePropertyOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				replaceTektonPipelinePropertyOptionsModel.SetPropertyName("debug-pipeline")
				replaceTektonPipelinePropertyOptionsModel.SetName("prop1")
				replaceTektonPipelinePropertyOptionsModel.SetType("text")
				replaceTektonPipelinePropertyOptionsModel.SetValue("https://github.com/open-toolchain/hello-tekton.git")
				replaceTektonPipelinePropertyOptionsModel.SetEnum([]string{"testString"})
				replaceTektonPipelinePropertyOptionsModel.SetLocked(false)
				replaceTektonPipelinePropertyOptionsModel.SetPath("testString")
				replaceTektonPipelinePropertyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceTektonPipelinePropertyOptionsModel).ToNot(BeNil())
				Expect(replaceTektonPipelinePropertyOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(replaceTektonPipelinePropertyOptionsModel.PropertyName).To(Equal(core.StringPtr("debug-pipeline")))
				Expect(replaceTektonPipelinePropertyOptionsModel.Name).To(Equal(core.StringPtr("prop1")))
				Expect(replaceTektonPipelinePropertyOptionsModel.Type).To(Equal(core.StringPtr("text")))
				Expect(replaceTektonPipelinePropertyOptionsModel.Value).To(Equal(core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")))
				Expect(replaceTektonPipelinePropertyOptionsModel.Enum).To(Equal([]string{"testString"}))
				Expect(replaceTektonPipelinePropertyOptionsModel.Locked).To(Equal(core.BoolPtr(false)))
				Expect(replaceTektonPipelinePropertyOptionsModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(replaceTektonPipelinePropertyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceTektonPipelineTriggerPropertyOptions successfully`, func() {
				// Construct an instance of the ReplaceTektonPipelineTriggerPropertyOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				triggerID := "1bb892a1-2e04-4768-a369-b1159eace147"
				propertyName := "debug-pipeline"
				replaceTektonPipelineTriggerPropertyOptionsName := "prop1"
				replaceTektonPipelineTriggerPropertyOptionsType := "text"
				replaceTektonPipelineTriggerPropertyOptionsModel := cdTektonPipelineService.NewReplaceTektonPipelineTriggerPropertyOptions(pipelineID, triggerID, propertyName, replaceTektonPipelineTriggerPropertyOptionsName, replaceTektonPipelineTriggerPropertyOptionsType)
				replaceTektonPipelineTriggerPropertyOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				replaceTektonPipelineTriggerPropertyOptionsModel.SetTriggerID("1bb892a1-2e04-4768-a369-b1159eace147")
				replaceTektonPipelineTriggerPropertyOptionsModel.SetPropertyName("debug-pipeline")
				replaceTektonPipelineTriggerPropertyOptionsModel.SetName("prop1")
				replaceTektonPipelineTriggerPropertyOptionsModel.SetType("text")
				replaceTektonPipelineTriggerPropertyOptionsModel.SetValue("https://github.com/open-toolchain/hello-tekton.git")
				replaceTektonPipelineTriggerPropertyOptionsModel.SetEnum([]string{"testString"})
				replaceTektonPipelineTriggerPropertyOptionsModel.SetPath("testString")
				replaceTektonPipelineTriggerPropertyOptionsModel.SetLocked(false)
				replaceTektonPipelineTriggerPropertyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceTektonPipelineTriggerPropertyOptionsModel).ToNot(BeNil())
				Expect(replaceTektonPipelineTriggerPropertyOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(replaceTektonPipelineTriggerPropertyOptionsModel.TriggerID).To(Equal(core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")))
				Expect(replaceTektonPipelineTriggerPropertyOptionsModel.PropertyName).To(Equal(core.StringPtr("debug-pipeline")))
				Expect(replaceTektonPipelineTriggerPropertyOptionsModel.Name).To(Equal(core.StringPtr("prop1")))
				Expect(replaceTektonPipelineTriggerPropertyOptionsModel.Type).To(Equal(core.StringPtr("text")))
				Expect(replaceTektonPipelineTriggerPropertyOptionsModel.Value).To(Equal(core.StringPtr("https://github.com/open-toolchain/hello-tekton.git")))
				Expect(replaceTektonPipelineTriggerPropertyOptionsModel.Enum).To(Equal([]string{"testString"}))
				Expect(replaceTektonPipelineTriggerPropertyOptionsModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(replaceTektonPipelineTriggerPropertyOptionsModel.Locked).To(Equal(core.BoolPtr(false)))
				Expect(replaceTektonPipelineTriggerPropertyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRerunTektonPipelineRunOptions successfully`, func() {
				// Construct an instance of the RerunTektonPipelineRunOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				id := "94619026-912b-4d92-8f51-6c74f0692d90"
				rerunTektonPipelineRunOptionsModel := cdTektonPipelineService.NewRerunTektonPipelineRunOptions(pipelineID, id)
				rerunTektonPipelineRunOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				rerunTektonPipelineRunOptionsModel.SetID("94619026-912b-4d92-8f51-6c74f0692d90")
				rerunTektonPipelineRunOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(rerunTektonPipelineRunOptionsModel).ToNot(BeNil())
				Expect(rerunTektonPipelineRunOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(rerunTektonPipelineRunOptionsModel.ID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(rerunTektonPipelineRunOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTool successfully`, func() {
				id := "testString"
				_model, err := cdTektonPipelineService.NewTool(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTriggerSourcePropertiesPrototype successfully`, func() {
				url := "testString"
				_model, err := cdTektonPipelineService.NewTriggerSourcePropertiesPrototype(url)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTriggerSourcePrototype successfully`, func() {
				typeVar := "testString"
				var properties *cdtektonpipelinev2.TriggerSourcePropertiesPrototype = nil
				_, err := cdTektonPipelineService.NewTriggerSourcePrototype(typeVar, properties)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewUpdateTektonPipelineOptions successfully`, func() {
				// Construct an instance of the UpdateTektonPipelineOptions model
				id := "94619026-912b-4d92-8f51-6c74f0692d90"
				updateTektonPipelineOptionsModel := cdTektonPipelineService.NewUpdateTektonPipelineOptions(id)
				updateTektonPipelineOptionsModel.SetID("94619026-912b-4d92-8f51-6c74f0692d90")
				updateTektonPipelineOptionsModel.SetTektonPipelinePatch(map[string]interface{}{"anyKey": "anyValue"})
				updateTektonPipelineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTektonPipelineOptionsModel).ToNot(BeNil())
				Expect(updateTektonPipelineOptionsModel.ID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(updateTektonPipelineOptionsModel.TektonPipelinePatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateTektonPipelineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTektonPipelineTriggerOptions successfully`, func() {
				// Construct an instance of the UpdateTektonPipelineTriggerOptions model
				pipelineID := "94619026-912b-4d92-8f51-6c74f0692d90"
				triggerID := "1bb892a1-2e04-4768-a369-b1159eace147"
				updateTektonPipelineTriggerOptionsModel := cdTektonPipelineService.NewUpdateTektonPipelineTriggerOptions(pipelineID, triggerID)
				updateTektonPipelineTriggerOptionsModel.SetPipelineID("94619026-912b-4d92-8f51-6c74f0692d90")
				updateTektonPipelineTriggerOptionsModel.SetTriggerID("1bb892a1-2e04-4768-a369-b1159eace147")
				updateTektonPipelineTriggerOptionsModel.SetTriggerPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateTektonPipelineTriggerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTektonPipelineTriggerOptionsModel).ToNot(BeNil())
				Expect(updateTektonPipelineTriggerOptionsModel.PipelineID).To(Equal(core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90")))
				Expect(updateTektonPipelineTriggerOptionsModel.TriggerID).To(Equal(core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147")))
				Expect(updateTektonPipelineTriggerOptionsModel.TriggerPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateTektonPipelineTriggerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewWorkerIdentity successfully`, func() {
				id := "testString"
				_model, err := cdTektonPipelineService.NewWorkerIdentity(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalDefinitionSource successfully`, func() {
			// Construct an instance of the model.
			model := new(cdtektonpipelinev2.DefinitionSource)
			model.Type = core.StringPtr("testString")
			model.Properties = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *cdtektonpipelinev2.DefinitionSource
			err = cdtektonpipelinev2.UnmarshalDefinitionSource(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDefinitionSourceProperties successfully`, func() {
			// Construct an instance of the model.
			model := new(cdtektonpipelinev2.DefinitionSourceProperties)
			model.URL = core.StringPtr("testString")
			model.Branch = core.StringPtr("testString")
			model.Tag = core.StringPtr("testString")
			model.Path = core.StringPtr("testString")
			model.Tool = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *cdtektonpipelinev2.DefinitionSourceProperties
			err = cdtektonpipelinev2.UnmarshalDefinitionSourceProperties(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGenericSecret successfully`, func() {
			// Construct an instance of the model.
			model := new(cdtektonpipelinev2.GenericSecret)
			model.Type = core.StringPtr("token_matches")
			model.Value = core.StringPtr("testString")
			model.Source = core.StringPtr("header")
			model.KeyName = core.StringPtr("testString")
			model.Algorithm = core.StringPtr("md4")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *cdtektonpipelinev2.GenericSecret
			err = cdtektonpipelinev2.UnmarshalGenericSecret(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalPipelineRunTrigger successfully`, func() {
			// Construct an instance of the model.
			model := new(cdtektonpipelinev2.PipelineRunTrigger)
			model.Name = core.StringPtr("start-deploy")
			model.Properties = map[string]interface{}{"anyKey": "anyValue"}
			model.SecureProperties = map[string]interface{}{"anyKey": "anyValue"}
			model.HeadersVar = map[string]interface{}{"anyKey": "anyValue"}
			model.Body = map[string]interface{}{"anyKey": "anyValue"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *cdtektonpipelinev2.PipelineRunTrigger
			err = cdtektonpipelinev2.UnmarshalPipelineRunTrigger(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTektonPipelinePatch successfully`, func() {
			// Construct an instance of the model.
			model := new(cdtektonpipelinev2.TektonPipelinePatch)
			model.NextBuildNumber = core.Int64Ptr(int64(1))
			model.EnableNotifications = core.BoolPtr(true)
			model.EnablePartialCloning = core.BoolPtr(true)
			model.Worker = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *cdtektonpipelinev2.TektonPipelinePatch
			err = cdtektonpipelinev2.UnmarshalTektonPipelinePatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTool successfully`, func() {
			// Construct an instance of the model.
			model := new(cdtektonpipelinev2.Tool)
			model.ID = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *cdtektonpipelinev2.Tool
			err = cdtektonpipelinev2.UnmarshalTool(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTriggerPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(cdtektonpipelinev2.TriggerPatch)
			model.Type = core.StringPtr("manual")
			model.Name = core.StringPtr("start-deploy")
			model.EventListener = core.StringPtr("testString")
			model.Tags = []string{"testString"}
			model.Worker = nil
			model.MaxConcurrentRuns = core.Int64Ptr(int64(4))
			model.LimitWaitingRuns = core.BoolPtr(false)
			model.Enabled = core.BoolPtr(true)
			model.Secret = nil
			model.Cron = core.StringPtr("testString")
			model.Timezone = core.StringPtr("America/Los_Angeles, CET, Europe/London, GMT, US/Eastern, or UTC")
			model.Source = nil
			model.Events = []string{"push", "pull_request"}
			model.Filter = core.StringPtr("header['x-github-event'] == 'push' && body.ref == 'refs/heads/main'")
			model.Favorite = core.BoolPtr(false)
			model.EnableEventsFromForks = core.BoolPtr(false)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *cdtektonpipelinev2.TriggerPatch
			err = cdtektonpipelinev2.UnmarshalTriggerPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTriggerSourcePropertiesPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(cdtektonpipelinev2.TriggerSourcePropertiesPrototype)
			model.URL = core.StringPtr("testString")
			model.Branch = core.StringPtr("testString")
			model.Pattern = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *cdtektonpipelinev2.TriggerSourcePropertiesPrototype
			err = cdtektonpipelinev2.UnmarshalTriggerSourcePropertiesPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTriggerSourcePrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(cdtektonpipelinev2.TriggerSourcePrototype)
			model.Type = core.StringPtr("testString")
			model.Properties = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *cdtektonpipelinev2.TriggerSourcePrototype
			err = cdtektonpipelinev2.UnmarshalTriggerSourcePrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalWorkerIdentity successfully`, func() {
			// Construct an instance of the model.
			model := new(cdtektonpipelinev2.WorkerIdentity)
			model.ID = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *cdtektonpipelinev2.WorkerIdentity
			err = cdtektonpipelinev2.UnmarshalWorkerIdentity(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIGVtZXJnZW5jeSBicm9hZGNhc3Qgc3lzdGVt")
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

func CreateMockByteArray(encodedString string) *[]byte {
	ba, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}
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
