// +build integration

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
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.ibm.com/CloudEngineering/go-sdk-template/exampleservicev1"
)

/**
 * This class contains an integration test for the example service.
 *
 * Notes:
 *
 * 1. This example integration test shows how to automatically skip tests if the required config file
 *    is not available.
 *
 * 2. Before running this test:
 *    a. "cp example-service.env.hide example-service.env"
 *    b. start up the ExampleService service by following the instructions here:
 *    https://github.ibm.com/CloudEngineering/go-sdk-template/blob/master/README_FIRST.md#integration-tests
 */
const externalConfigFile = "../example-service.env"

var (
	err          error
	service      *exampleservicev1.ExampleServiceV1
	serviceURL   string
	config       map[string]string
	configLoaded bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`ExampleServiceV1 Integration Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(exampleservicev1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			configLoaded = true
			fmt.Printf("Service URL: %s\n", serviceURL)
		})
	})

	Describe(`Service-level tests`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			service, err = exampleservicev1.NewExampleServiceV1UsingExternalConfig(
				&exampleservicev1.ExampleServiceV1Options{})
			Expect(err).To(BeNil())
			Expect(service).ToNot(BeNil())
			Expect(service.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`CreateResource() - create a resource`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Successfully invoke CreateResource()`, func() {
			createResourceOptions := &exampleservicev1.CreateResourceOptions{
				ResourceID: core.StringPtr("3"),
				Name:       core.StringPtr("To Kill a Mockingbird"),
				Tag:        core.StringPtr("Book"),
			}
			result, detailedResponse, err := service.CreateResource(createResourceOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(201))
			Expect(result).ToNot(BeNil())

			Expect(*result.ResourceID).To(Equal("3"))
			Expect(*result.Name).To(Equal("To Kill a Mockingbird"))
			Expect(*result.Tag).To(Equal("Book"))
		})
	})

	Describe(`GetResource() - get a resource`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Successfully invoke GetResource()`, func() {
			getResourceOptions := &exampleservicev1.GetResourceOptions{
				ResourceID: core.StringPtr("1"),
			}
			result, detailedResponse, err := service.GetResource(getResourceOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			Expect(*result.ResourceID).To(Equal("1"))
			Expect(*result.Name).To(Equal("The Great Gatsby"))
			Expect(*result.Tag).To(Equal("Book"))
		})

		It(`Negative test - invoke GetResource() with error`, func() {
			resourceID := "BAD_RESOURCE_ID"
			getResourceOptions := &exampleservicev1.GetResourceOptions{
				ResourceID: &resourceID,
			}
			result, detailedResponse, err := service.GetResource(getResourceOptions)
			Expect(err).ToNot(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(404))
			Expect(result).To(BeNil())
		})
	})

	Describe(`ListResources() - list resources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Successfully invoke ListResources()`, func() {
			listResourcesOptions := &exampleservicev1.ListResourcesOptions{}
			result, detailedResponse, err := service.ListResources(listResourcesOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})
})
