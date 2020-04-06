// +build integration

package exampleservicev1_test

import (
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

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
 *    https://github.ibm.com/CloudEngineering/java-sdk-template/blob/master/README_FIRST.md#integration-tests
 */
const externalConfigFile = "../example-service.env"

var (
	service      *exampleservicev1.ExampleServiceV1
	err          error
	configLoaded bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`ExampleServiceV1`, func() {
	It("Successfully load the configuration", func() {
		err = godotenv.Load(externalConfigFile)
		if err == nil {
			//
			// Retrieve any test-specific properties from the environment.
			// ...

			// Set the flag to allow tests to execute.
			configLoaded = true
		}

		if !configLoaded {
			Skip("External configuration could not be loaded, skipping...")
		}
	})

	It("Successfully create the service client instance", func() {
		shouldSkipTest()

		service, err = exampleservicev1.NewExampleServiceV1UsingExternalConfig(
			&exampleservicev1.ExampleServiceV1Options{})
		Expect(err).To(BeNil())
		Expect(service).ToNot(BeNil())
		Expect(service.Service.Options.URL).To(Not(Equal("")))
	})

	Describe(`ListResources(listResourcesOptions *ListResourcesOptions)`, func() {
		It(`Successfully list all resources`, func() {
			shouldSkipTest()

			result, detailedResponse, err := service.ListResources(&exampleservicev1.ListResourcesOptions{})
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			resources := result.Resources
			Expect(resources).ToNot(BeNil())
			Expect(len(resources)).Should(BeNumerically(">=", 2))

			firstResource := resources[0]
			Expect(*firstResource.ResourceID).To(Equal("1"))
			Expect(*firstResource.Name).To(Equal("The Great Gatsby"))
			Expect(*firstResource.Tag).To(Equal("Book"))

			secondResource := resources[1]
			Expect(*secondResource.ResourceID).To(Equal("2"))
			Expect(*secondResource.Name).To(Equal("Pride and Prejudice"))
			Expect(*secondResource.Tag).To(Equal("Book"))
		})
	})

	Describe(`GetResource(getResourceOptions *GetResourceOptions)`, func() {
		It(`Successfully get resource by ResourceID`, func() {
			shouldSkipTest()

			resourceID := "1"
			getResourceOptions := &exampleservicev1.GetResourceOptions{
				ResourceID: &resourceID,
			}
			result, detailedResponse, err := service.GetResource(getResourceOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))

			Expect(*result.ResourceID).To(Equal("1"))
			Expect(*result.Name).To(Equal("The Great Gatsby"))
			Expect(*result.Tag).To(Equal("Book"))
		})

		It(`Negative test - retrieve resource with incorrect resource id`, func() {
			shouldSkipTest()

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

	Describe(`CreateResource(createResourceOptions *CreateResourceOptions)`, func() {
		It(`Successfully create new resource`, func() {
			shouldSkipTest()

			createResourceOptions := service.NewCreateResourceOptions().
				SetResourceID("3").
				SetName("To Kill a Mockingbird").
				SetTag("Book")
			result, detailedResponse, err := service.CreateResource(createResourceOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(201))
			Expect(*result.ResourceID).To(Equal("3"))
			Expect(*result.Name).To(Equal("To Kill a Mockingbird"))
			Expect(*result.Tag).To(Equal("Book"))
		})
	})
})
