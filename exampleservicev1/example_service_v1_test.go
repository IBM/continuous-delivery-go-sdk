/**
 * (C) Copyright IBM Corp. 2020.
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
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
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
			testService, testServiceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
				URL: "https://exampleservicev1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
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
				testService, testServiceErr := exampleservicev1.NewExampleServiceV1UsingExternalConfig(&exampleservicev1.ExampleServiceV1Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := exampleservicev1.NewExampleServiceV1UsingExternalConfig(&exampleservicev1.ExampleServiceV1Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := exampleservicev1.NewExampleServiceV1UsingExternalConfig(&exampleservicev1.ExampleServiceV1Options{
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"EXAMPLE_SERVICE_URL": "https://exampleservicev1/api",
				"EXAMPLE_SERVICE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := exampleservicev1.NewExampleServiceV1UsingExternalConfig(&exampleservicev1.ExampleServiceV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"EXAMPLE_SERVICE_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := exampleservicev1.NewExampleServiceV1UsingExternalConfig(&exampleservicev1.ExampleServiceV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListResources(listResourcesOptions *ListResourcesOptions) - Operation response error`, func() {
		listResourcesPath := "/resources"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			 		defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listResourcesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListResources with error: Operation response processing error`, func() {
				testService, testServiceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListResourcesOptions model
				listResourcesOptionsModel := new(exampleservicev1.ListResourcesOptions)
				listResourcesOptionsModel.Limit = core.Int64Ptr(int64(38))
				listResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListResources(listResourcesOptionsModel)
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listResourcesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"offset": 6, "limit": 5, "resources": [{"resource_id": "ResourceID", "name": "Name", "tag": "Tag", "read_only": "ReadOnly"}]}`)
				}))
			})
			It(`Invoke ListResources successfully`, func() {
				testService, testServiceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListResources(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourcesOptions model
				listResourcesOptionsModel := new(exampleservicev1.ListResourcesOptions)
				listResourcesOptionsModel.Limit = core.Int64Ptr(int64(38))
 				listResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListResources(listResourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListResources with error: Operation request error`, func() {
				testService, testServiceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListResourcesOptions model
				listResourcesOptionsModel := new(exampleservicev1.ListResourcesOptions)
				listResourcesOptionsModel.Limit = core.Int64Ptr(int64(38))
				listResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListResources(listResourcesOptionsModel)
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
					Expect(req.URL.Path).To(Equal(createResourcePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateResource with error: Operation response processing error`, func() {
				testService, testServiceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateResourceOptions model
				createResourceOptionsModel := new(exampleservicev1.CreateResourceOptions)
				createResourceOptionsModel.ResourceID = core.StringPtr("testString")
				createResourceOptionsModel.Name = core.StringPtr("testString")
				createResourceOptionsModel.Tag = core.StringPtr("testString")
				createResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateResource(createResourceOptionsModel)
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createResourcePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `{"resource_id": "ResourceID", "name": "Name", "tag": "Tag", "read_only": "ReadOnly"}`)
				}))
			})
			It(`Invoke CreateResource successfully`, func() {
				testService, testServiceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateResource(nil)
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
				result, response, operationErr = testService.CreateResource(createResourceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateResource with error: Operation validation and request error`, func() {
				testService, testServiceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateResourceOptions model
				createResourceOptionsModel := new(exampleservicev1.CreateResourceOptions)
				createResourceOptionsModel.ResourceID = core.StringPtr("testString")
				createResourceOptionsModel.Name = core.StringPtr("testString")
				createResourceOptionsModel.Tag = core.StringPtr("testString")
				createResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateResource(createResourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateResourceOptions model with no property values
				createResourceOptionsModelNew := new(exampleservicev1.CreateResourceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateResource(createResourceOptionsModelNew)
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
					Expect(req.URL.Path).To(Equal(getResourcePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResource with error: Operation response processing error`, func() {
				testService, testServiceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetResourceOptions model
				getResourceOptionsModel := new(exampleservicev1.GetResourceOptions)
				getResourceOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetResource(getResourceOptionsModel)
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getResourcePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"resource_id": "ResourceID", "name": "Name", "tag": "Tag", "read_only": "ReadOnly"}`)
				}))
			})
			It(`Invoke GetResource successfully`, func() {
				testService, testServiceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetResource(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceOptions model
				getResourceOptionsModel := new(exampleservicev1.GetResourceOptions)
				getResourceOptionsModel.ResourceID = core.StringPtr("testString")
 				getResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetResource(getResourceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetResource with error: Operation validation and request error`, func() {
				testService, testServiceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetResourceOptions model
				getResourceOptionsModel := new(exampleservicev1.GetResourceOptions)
				getResourceOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetResource(getResourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceOptions model with no property values
				getResourceOptionsModelNew := new(exampleservicev1.GetResourceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetResource(getResourceOptionsModelNew)
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
			testService, _ := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
				URL:           "http://exampleservicev1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateResourceOptions successfully`, func() {
				// Construct an instance of the CreateResourceOptions model
				createResourceOptionsResourceID := "testString"
				createResourceOptionsName := "testString"
				createResourceOptionsModel := testService.NewCreateResourceOptions(createResourceOptionsResourceID, createResourceOptionsName)
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
			It(`Invoke NewGetResourceOptions successfully`, func() {
				// Construct an instance of the GetResourceOptions model
				resourceID := "testString"
				getResourceOptionsModel := testService.NewGetResourceOptions(resourceID)
				getResourceOptionsModel.SetResourceID("testString")
				getResourceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceOptionsModel).ToNot(BeNil())
				Expect(getResourceOptionsModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListResourcesOptions successfully`, func() {
				// Construct an instance of the ListResourcesOptions model
				listResourcesOptionsModel := testService.NewListResourcesOptions()
				listResourcesOptionsModel.SetLimit(int64(38))
				listResourcesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listResourcesOptionsModel).ToNot(BeNil())
				Expect(listResourcesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listResourcesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResource successfully`, func() {
				resourceID := "testString"
				name := "testString"
				model, err := testService.NewResource(resourceID, name)
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
