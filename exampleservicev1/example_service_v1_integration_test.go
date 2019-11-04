// +build integration

package exampleservicev1_test

import (
    "testing"
    "github.com/IBM/go-sdk-core/core"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "github.ibm.com/CloudEngineering/go-sdk-template/exampleservicev1"
)

func TestExampleServiceIntegration(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Example Service Integration Test")
}

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
            result, detailedResponse, err := service.ListResources(listResourcesOptions)
            It(`Successfully list all resources`, func() {

                Expect(err).To(BeNil())
                Expect(detailedResponse.StatusCode).To(Equal(200))

                firstResource := result.Resources[0]
                Expect(*firstResource.ResourceID).To(Equal(int64(1)))
                Expect(*firstResource.Name).To(Equal("The Great Gatsby"))
                Expect(*firstResource.Tag).To(Equal("Book"))

                secondResource := result.Resources[1]
                Expect(*secondResource.ResourceID).To(Equal(int64(2)))
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
                result, detailedResponse, err := service.GetResource(getResourceOptions)
                It(`Successfully get resource by ResourceID`, func() {
                    Expect(err).To(BeNil())
                    Expect(detailedResponse.StatusCode).To(Equal(200))

                    Expect(*result.ResourceID).To(Equal(int64(1)))
                    Expect(*result.Name).To(Equal("The Great Gatsby"))
                    Expect(*result.Tag).To(Equal("Book"))
                })
            })

            Context(`Failed to get resource by ResourceID`, func() {
                resourceID := "111"
                getResourceOptions := &exampleservicev1.GetResourceOptions{
                    ResourceID: &resourceID,
                }
                result, detailedResponse, err := service.GetResource(getResourceOptions)
                It(`Failed to get resource by ResourceID`, func() {
                    Expect(result).To(BeNil())
                    Expect(detailedResponse.StatusCode).To(Equal(404))
                    Expect(err).Should(HaveOccurred())
                })
            })
        })
        Describe(`CreateResource(createResourceOptions *CreateResourceOptions)`, func() {
            createResourceOptions := service.NewCreateResourceOptions().
                                            SetResourceID(int64(3)).
                                            SetName("To Kill a Mockingbird").
                                            SetTag("Book")

            result, detailedResponse, err := service.CreateResource(createResourceOptions)
            It(`Successfully create new resource`, func() {
                Expect(err).To(BeNil())
                Expect(detailedResponse.StatusCode).To(Equal(201))
                Expect(*result.ResourceID).To(Equal(int64(3)))
                Expect(*result.Name).To(Equal("To Kill a Mockingbird"))
                Expect(*result.Tag).To(Equal("Book"))
            })
        })
    })


