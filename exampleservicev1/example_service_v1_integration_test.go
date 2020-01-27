// +build integration

package exampleservicev1_test

import (
	"github.com/IBM/go-sdk-core/v3/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/CloudEngineering/go-sdk-template/exampleservicev1"
)

var _ = Describe(`ExampleServiceV1`, func() {

	authenticator := &core.NoAuthAuthenticator{}
	options := &exampleservicev1.ExampleServiceV1Options{
		Authenticator: authenticator,
	}
	service, err := exampleservicev1.NewExampleServiceV1(options)
	It(`Successfully created ExampleServiceV1 service instance`, func() {
		Expect(err).To(BeNil())
	})

	err = service.SetServiceURL("http://localhost:3000")
	It(`Successfully change default service URL to point to localhost:3000`, func() {
		Expect(err).To(BeNil())
	})

	Describe(`ListResources(listResourcesOptions *ListResourcesOptions)`, func() {
		listResourcesOptions := &exampleservicev1.ListResourcesOptions{}

		It(`Successfully list all resources`, func() {
			result, detailedResponse, err := service.ListResources(listResourcesOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))

			firstResource := result.Resources[0]
			Expect(*firstResource.ResourceID).To(Equal("1"))
			Expect(*firstResource.Name).To(Equal("The Great Gatsby"))
			Expect(*firstResource.Tag).To(Equal("Book"))

			secondResource := result.Resources[1]
			Expect(*secondResource.ResourceID).To(Equal("2"))
			Expect(*secondResource.Name).To(Equal("Pride and Prejudice"))
			Expect(*secondResource.Tag).To(Equal("Book"))
		})
	})
	Describe(`GetResource(getResourceOptions *GetResourceOptions)`, func() {
		Context(`Successfully get resource by ResourceID`, func() {
			resourceID := "1"
			getResourceOptions := &exampleservicev1.GetResourceOptions{
				ResourceID: &resourceID,
			}
			It(`Successfully get resource by ResourceID`, func() {
				result, detailedResponse, err := service.GetResource(getResourceOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ResourceID).To(Equal("1"))
				Expect(*result.Name).To(Equal("The Great Gatsby"))
				Expect(*result.Tag).To(Equal("Book"))
			})
		})

		Context(`Failed to get resource by ResourceID`, func() {
			resourceID := "111"
			getResourceOptions := &exampleservicev1.GetResourceOptions{
				ResourceID: &resourceID,
			}
			It(`Failed to get resource by ResourceID`, func() {
				result, detailedResponse, err := service.GetResource(getResourceOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})
	Describe(`CreateResource(createResourceOptions *CreateResourceOptions)`, func() {
		createResourceOptions := service.NewCreateResourceOptions().
			SetResourceID("3").
			SetName("To Kill a Mockingbird").
			SetTag("Book")

		It(`Successfully create new resource`, func() {
			result, detailedResponse, err := service.CreateResource(createResourceOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(201))
			Expect(*result.ResourceID).To(Equal("3"))
			Expect(*result.Name).To(Equal("To Kill a Mockingbird"))
			Expect(*result.Tag).To(Equal("Book"))
		})
	})
})
