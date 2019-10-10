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
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/CloudEngineering/go-sdk-template/exampleservicev1"
	"net/http"
	"net/http/httptest"
	"strings"
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
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call ListResources`, func() {
				defer testServer.Close()

				testService, testServiceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL: testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListResources(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listResourcesOptions := testService.NewListResourcesOptions()
				result, response, operationErr = testService.ListResources(listResourcesOptions)
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
				fmt.Fprintf(res, `{"resource_id": 10, "name": "fake Name"}`)
				res.WriteHeader(201)
			}))
			It(`Succeed to call CreateResource`, func() {
				defer testServer.Close()

				testService, testServiceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL: testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateResource(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createResourceOptions := testService.NewCreateResourceOptions()
				result, response, operationErr = testService.CreateResource(createResourceOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetResource(getResourceOptions *GetResourceOptions)`, func() {
		getResourcePath := "/resources/{resource_id}"
		resourceID := "exampleString"
		getResourcePath = strings.Replace(getResourcePath, "{resource_id}", resourceID, 1)
		Context(`Successfully - Info for a specific resource`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getResourcePath))
				Expect(req.Method).To(Equal("GET"))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"resource_id": 10, "name": "fake Name"}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call GetResource`, func() {
				defer testServer.Close()

				testService, testServiceErr := exampleservicev1.NewExampleServiceV1(&exampleservicev1.ExampleServiceV1Options{
					URL: testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetResource(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getResourceOptions := testService.NewGetResourceOptions(resourceID)
				result, response, operationErr = testService.GetResource(getResourceOptions)
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
				resourceID := int64(1234)
				name := "exampleString"
				model, err := testService.NewResource(resourceID, name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
})
