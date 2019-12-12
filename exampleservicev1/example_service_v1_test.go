/**
 * (C) Copyright IBM Corp. 2019.
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
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/CloudEngineering/go-sdk-template/exampleservicev1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

var _ = Describe(`ExampleServiceV1`, func() {
	Describe(`ListResources(listResourcesOptions *ListResourcesOptions)`, func() {
		listResourcesPath := "/resources"
		Context(`Successfully - List all resources`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listResourcesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"offset": 6, "limit": 5, "resources": []}`)
			}))
			It(`Succeed to call ListResources`, func() {
				defer testServer.Close()

				testService, testServiceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL: testServer.URL,
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

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListResources(listResourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateResource(createResourceOptions *CreateResourceOptions)`, func() {
		createResourcePath := "/resources"
		Context(`Successfully - Create a resource`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createResourcePath))
				Expect(req.Method).To(Equal("POST"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"resource_id": "fake_ResourceID", "name": "fake_Name", "tag": "fake_Tag"}`)
			}))
			It(`Succeed to call CreateResource`, func() {
				defer testServer.Close()

				testService, testServiceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL: testServer.URL,
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

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateResource(createResourceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetResource(getResourceOptions *GetResourceOptions)`, func() {
		getResourcePath := "/resources/{resource_id}"
		getResourcePath = strings.Replace(getResourcePath, "{resource_id}", "testString", 1)
		Context(`Successfully - Info for a specific resource`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getResourcePath))
				Expect(req.Method).To(Equal("GET"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"resource_id": "fake_ResourceID", "name": "fake_Name", "tag": "fake_Tag"}`)
			}))
			It(`Succeed to call GetResource`, func() {
				defer testServer.Close()

				testService, testServiceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL: testServer.URL,
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

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetResource(getResourceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("Model constructor tests", func () {
		Context("with a sample service", func () {
			testService, _ := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
				URL: "http://exampleservicev1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It("should call NewResource successfully", func () {
				resourceID := "testString"
				name := "testString"
				model, err := testService.NewResource(resourceID, name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe("Utility function tests", func () {
		It("Call CreateMockMap successfully", func () {
			mockMap := CreateMockMap()
			Expect(mockMap).ToNot(BeNil())
		})
		It("Call CreateMockByteArray successfully", func () {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It("Call CreateMockUUID successfully", func () {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It("Call CreateMockReader successfully", func () {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It("Call CreateMockDate successfully", func () {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It("Call CreateMockDateTime successfully", func () {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockMap() map[string]interface{} {
	m := make(map[string]interface{})
	return m
}

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, len(mockData))
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
	d := strfmt.Date(time.Now())
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Now())
	return &d
}
