// +build examples

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

package exampleservicev1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/CloudEngineering/go-sdk-template/exampleservicev1"
)

//
// This file provides an example of how to use the ExampleService service.
//
// The following configuration properties are assumed to be defined:
// EXAMPLE_SERVICE_URL=<service base url>
// EXAMPLE_SERVICE_AUTH_TYPE=iam
// EXAMPLE_SERVICE_APIKEY=<IAM apikey>
// EXAMPLE_SERVICE_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../example_service_v1.env"

var (
	exampleServiceService *exampleservicev1.ExampleServiceV1
	config       map[string]string
	configLoaded bool = false
)

// Globlal variables to hold link values
var (
	getResourceLink string
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`ExampleServiceV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(exampleservicev1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			exampleServiceServiceOptions := &exampleservicev1.ExampleServiceV1Options{}

			exampleServiceService, err = exampleservicev1.NewExampleServiceV1UsingExternalConfig(exampleServiceServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(exampleServiceService).ToNot(BeNil())
		})
	})

	Describe(`ExampleServiceV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateResource request example`, func() {
			fmt.Println("\nCreateResource() result:")
			// begin-create_resource

			createResourceOptions := exampleServiceService.NewCreateResourceOptions(
				"The Hunt for Red October",
				"Book",
			)

			resource, response, err := exampleServiceService.CreateResource(createResourceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resource, "", "  ")
			fmt.Println(string(b))

			// end-create_resource

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(resource).ToNot(BeNil())

			getResourceLink = *resource.ResourceID

		})
		It(`ListResources request example`, func() {
			fmt.Println("\nListResources() result:")
			// begin-list_resources

			listResourcesOptions := exampleServiceService.NewListResourcesOptions()
			listResourcesOptions.SetLimit(int64(1))

			resources, response, err := exampleServiceService.ListResources(listResourcesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resources, "", "  ")
			fmt.Println(string(b))

			// end-list_resources

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resources).ToNot(BeNil())

		})
		It(`GetResource request example`, func() {
			fmt.Println("\nGetResource() result:")
			// begin-get_resource

			getResourceOptions := exampleServiceService.NewGetResourceOptions(
				getResourceLink,
			)

			resource, response, err := exampleServiceService.GetResource(getResourceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resource, "", "  ")
			fmt.Println(string(b))

			// end-get_resource

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resource).ToNot(BeNil())

		})
		It(`GetResourceEncoded request example`, func() {
			fmt.Println("\nGetResourceEncoded() result:")
			// begin-get_resource_encoded

			getResourceEncodedOptions := exampleServiceService.NewGetResourceEncodedOptions(
				"url%3encoded%3resource%3id",
			)

			resource, response, err := exampleServiceService.GetResourceEncoded(getResourceEncodedOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resource, "", "  ")
			fmt.Println(string(b))

			// end-get_resource_encoded

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resource).ToNot(BeNil())

		})
	})
})
