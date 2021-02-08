// +build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/CloudEngineering/go-sdk-template/exampleservicev1"
)

/**
 * This file contains an integration test for the exampleservicev1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 *
 * Before running this test:
 * a. "cp example-service.env.hide example-service.env"
 * b. start up the ExampleService service by following the instructions here:
 * https://github.ibm.com/CloudEngineering/go-sdk-template/blob/main/README_FIRST.md#integration-tests
 */

var _ = Describe(`ExampleServiceV1 Integration Tests`, func() {

	const externalConfigFile = "../example-service.env"

	var (
		err                   error
		exampleServiceService *exampleservicev1.ExampleServiceV1
		serviceURL            string
		config                map[string]string
	)

	// Globlal variables to hold link values
	var (
		getResourceLink string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

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

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			exampleServiceServiceOptions := &exampleservicev1.ExampleServiceV1Options{}

			exampleServiceService, err = exampleservicev1.NewExampleServiceV1UsingExternalConfig(exampleServiceServiceOptions)

			Expect(err).To(BeNil())
			Expect(exampleServiceService).ToNot(BeNil())
			Expect(exampleServiceService.Service.Options.URL).To(Equal(serviceURL))

			goLogger := log.New(GinkgoWriter, "", log.LstdFlags)
			core.SetLogger(core.NewLogger(core.LevelDebug, goLogger, goLogger))
			exampleServiceService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateResource - Create a resource`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateResource(createResourceOptions *CreateResourceOptions)`, func() {

			createResourceOptions := &exampleservicev1.CreateResourceOptions{
				ResourceID: core.StringPtr("3"),
				Name:       core.StringPtr("To Kill a MockingBird"),
				Tag:        core.StringPtr("Book"),
			}

			resource, response, err := exampleServiceService.CreateResource(createResourceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(resource).ToNot(BeNil())

			Expect(*resource.ResourceID).To(Equal("3"))
			Expect(*resource.Name).To(Equal("To Kill a MockingBird"))
			Expect(*resource.Tag).To(Equal("Book"))

			getResourceLink = *resource.ResourceID
		})
	})

	Describe(`ListResources - List all resources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListResources(listResourcesOptions *ListResourcesOptions)`, func() {

			listResourcesOptions := &exampleservicev1.ListResourcesOptions{
				Limit: core.Int64Ptr(int64(100)),
			}

			resources, response, err := exampleServiceService.ListResources(listResourcesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resources).ToNot(BeNil())
		})
	})

	Describe(`GetResource - Info for a specific resource`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResource(getResourceOptions *GetResourceOptions)`, func() {
			Expect(getResourceLink).ToNot(BeEmpty())

			getResourceOptions := &exampleservicev1.GetResourceOptions{
				ResourceID: core.StringPtr(getResourceLink),
			}

			resource, response, err := exampleServiceService.GetResource(getResourceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resource).ToNot(BeNil())

			Expect(*resource.ResourceID).To(Equal("3"))
			Expect(*resource.Name).To(Equal("To Kill a MockingBird"))
			Expect(*resource.Tag).To(Equal("Book"))
		})
		It(`Negative test - invoke GetResource() with error`, func() {
			resourceID := "BAD_RESOURCE_ID"
			getResourceOptions := &exampleservicev1.GetResourceOptions{
				ResourceID: &resourceID,
			}
			result, detailedResponse, err := exampleServiceService.GetResource(getResourceOptions)
			Expect(err).ToNot(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(404))
			Expect(result).To(BeNil())
		})
	})
})
