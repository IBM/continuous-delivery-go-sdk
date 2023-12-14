// +build integration

/**
 * (C) Copyright IBM Corp. 2023.
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

package cdtoolchainv2_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/continuous-delivery-go-sdk/cdtoolchainv2"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the cdtoolchainv2 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`CdToolchainV2 Integration Tests`, func() {
	const externalConfigFile = "../cd_toolchain_v2.env"

	var (
		err          error
		cdToolchainService *cdtoolchainv2.CdToolchainV2
		serviceURL   string
		config       map[string]string

		// Variables to hold link values
		toolIDLink string
		toolchainIDLink string
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
			config, err = core.GetServiceProperties(cdtoolchainv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			cdToolchainServiceOptions := &cdtoolchainv2.CdToolchainV2Options{}

			cdToolchainService, err = cdtoolchainv2.NewCdToolchainV2UsingExternalConfig(cdToolchainServiceOptions)
			Expect(err).To(BeNil())
			Expect(cdToolchainService).ToNot(BeNil())
			Expect(cdToolchainService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			cdToolchainService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateToolchain - Create a toolchain`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateToolchain(createToolchainOptions *CreateToolchainOptions)`, func() {
			createToolchainOptions := &cdtoolchainv2.CreateToolchainOptions{
				Name: core.StringPtr("TestToolchainV2"),
				ResourceGroupID: core.StringPtr("6a9a01f2cff54a7f966f803d92877123"),
				Description: core.StringPtr("A sample toolchain to test the API"),
			}

			toolchainPost, response, err := cdToolchainService.CreateToolchain(createToolchainOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(toolchainPost).ToNot(BeNil())

			toolchainIDLink = *toolchainPost.ID
			fmt.Fprintf(GinkgoWriter, "Saved toolchainIDLink value: %v\n", toolchainIDLink)
		})
	})

	Describe(`CreateTool - Create a tool`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTool(createToolOptions *CreateToolOptions)`, func() {
			createToolOptions := &cdtoolchainv2.CreateToolOptions{
				ToolchainID: &toolchainIDLink,
				ToolTypeID: core.StringPtr("draservicebroker"),
				Name: core.StringPtr("testString"),
				Parameters: map[string]interface{}{"anyKey": "anyValue"},
			}

			toolchainToolPost, response, err := cdToolchainService.CreateTool(createToolOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(toolchainToolPost).ToNot(BeNil())

			toolIDLink = *toolchainToolPost.ID
			fmt.Fprintf(GinkgoWriter, "Saved toolIDLink value: %v\n", toolIDLink)
		})
	})

	Describe(`ListToolchains - Get a list of toolchains`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListToolchains(listToolchainsOptions *ListToolchainsOptions) with pagination`, func(){
			listToolchainsOptions := &cdtoolchainv2.ListToolchainsOptions{
				ResourceGroupID: core.StringPtr("6a9a01f2cff54a7f966f803d92877123"),
				Limit: core.Int64Ptr(int64(10)),
				Start: core.StringPtr("testString"),
				Name: core.StringPtr("TestToolchainV2"),
			}

			listToolchainsOptions.Start = nil
			listToolchainsOptions.Limit = core.Int64Ptr(1)

			var allResults []cdtoolchainv2.ToolchainModel
			for {
				toolchainCollection, response, err := cdToolchainService.ListToolchains(listToolchainsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(toolchainCollection).ToNot(BeNil())
				allResults = append(allResults, toolchainCollection.Toolchains...)

				listToolchainsOptions.Start, err = toolchainCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listToolchainsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListToolchains(listToolchainsOptions *ListToolchainsOptions) using ToolchainsPager`, func(){
			listToolchainsOptions := &cdtoolchainv2.ListToolchainsOptions{
				ResourceGroupID: core.StringPtr("6a9a01f2cff54a7f966f803d92877123"),
				Limit: core.Int64Ptr(int64(10)),
				Name: core.StringPtr("TestToolchainV2"),
			}

			// Test GetNext().
			pager, err := cdToolchainService.NewToolchainsPager(listToolchainsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []cdtoolchainv2.ToolchainModel
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = cdToolchainService.NewToolchainsPager(listToolchainsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListToolchains() returned a total of %d item(s) using ToolchainsPager.\n", len(allResults))
		})
	})

	Describe(`GetToolchainByID - Get a toolchain`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetToolchainByID(getToolchainByIDOptions *GetToolchainByIDOptions)`, func() {
			getToolchainByIDOptions := &cdtoolchainv2.GetToolchainByIDOptions{
				ToolchainID: &toolchainIDLink,
			}

			toolchain, response, err := cdToolchainService.GetToolchainByID(getToolchainByIDOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(toolchain).ToNot(BeNil())
		})
	})

	Describe(`UpdateToolchain - Update a toolchain`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateToolchain(updateToolchainOptions *UpdateToolchainOptions)`, func() {
			toolchainPrototypePatchModel := &cdtoolchainv2.ToolchainPrototypePatch{
				Name: core.StringPtr("newToolchainName"),
				Description: core.StringPtr("New toolchain description"),
			}
			toolchainPrototypePatchModelAsPatch, asPatchErr := toolchainPrototypePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateToolchainOptions := &cdtoolchainv2.UpdateToolchainOptions{
				ToolchainID: &toolchainIDLink,
				ToolchainPrototypePatch: toolchainPrototypePatchModelAsPatch,
			}

			toolchainPatch, response, err := cdToolchainService.UpdateToolchain(updateToolchainOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(toolchainPatch).ToNot(BeNil())
		})
	})

	Describe(`CreateToolchainEvent - Create a toolchain event`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateToolchainEvent(createToolchainEventOptions *CreateToolchainEventOptions)`, func() {
			toolchainEventPrototypeDataApplicationJSONModel := &cdtoolchainv2.ToolchainEventPrototypeDataApplicationJSON{
				Content: map[string]interface{}{"anyKey": "anyValue"},
			}

			toolchainEventPrototypeDataModel := &cdtoolchainv2.ToolchainEventPrototypeData{
				ApplicationJSON: toolchainEventPrototypeDataApplicationJSONModel,
				TextPlain: core.StringPtr("This event is dispatched because the pipeline failed"),
			}

			createToolchainEventOptions := &cdtoolchainv2.CreateToolchainEventOptions{
				ToolchainID: &toolchainIDLink,
				Title: core.StringPtr("My-custom-event"),
				Description: core.StringPtr("This is my custom event"),
				ContentType: core.StringPtr("application/json"),
				Data: toolchainEventPrototypeDataModel,
			}

			toolchainEventPost, response, err := cdToolchainService.CreateToolchainEvent(createToolchainEventOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(toolchainEventPost).ToNot(BeNil())
		})
	})

	Describe(`ListTools - Get a list of tools bound to a toolchain`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTools(listToolsOptions *ListToolsOptions) with pagination`, func(){
			listToolsOptions := &cdtoolchainv2.ListToolsOptions{
				ToolchainID: &toolchainIDLink,
				Limit: core.Int64Ptr(int64(10)),
				Start: core.StringPtr("testString"),
			}

			listToolsOptions.Start = nil
			listToolsOptions.Limit = core.Int64Ptr(1)

			var allResults []cdtoolchainv2.ToolModel
			for {
				toolchainToolCollection, response, err := cdToolchainService.ListTools(listToolsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(toolchainToolCollection).ToNot(BeNil())
				allResults = append(allResults, toolchainToolCollection.Tools...)

				listToolsOptions.Start, err = toolchainToolCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listToolsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListTools(listToolsOptions *ListToolsOptions) using ToolsPager`, func(){
			listToolsOptions := &cdtoolchainv2.ListToolsOptions{
				ToolchainID: &toolchainIDLink,
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := cdToolchainService.NewToolsPager(listToolsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []cdtoolchainv2.ToolModel
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = cdToolchainService.NewToolsPager(listToolsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListTools() returned a total of %d item(s) using ToolsPager.\n", len(allResults))
		})
	})

	Describe(`GetToolByID - Get a tool`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetToolByID(getToolByIDOptions *GetToolByIDOptions)`, func() {
			getToolByIDOptions := &cdtoolchainv2.GetToolByIDOptions{
				ToolchainID: &toolchainIDLink,
				ToolID: &toolIDLink,
			}

			toolchainTool, response, err := cdToolchainService.GetToolByID(getToolByIDOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(toolchainTool).ToNot(BeNil())
		})
	})

	Describe(`UpdateTool - Update a tool`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateTool(updateToolOptions *UpdateToolOptions)`, func() {
			toolchainToolPrototypePatchModel := &cdtoolchainv2.ToolchainToolPrototypePatch{
				Name: core.StringPtr("MyTool"),
				ToolTypeID: core.StringPtr("draservicebroker"),
				Parameters: map[string]interface{}{"anyKey": "anyValue"},
			}
			toolchainToolPrototypePatchModelAsPatch, asPatchErr := toolchainToolPrototypePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateToolOptions := &cdtoolchainv2.UpdateToolOptions{
				ToolchainID: &toolchainIDLink,
				ToolID: &toolIDLink,
				ToolchainToolPrototypePatch: toolchainToolPrototypePatchModelAsPatch,
			}

			toolchainToolPatch, response, err := cdToolchainService.UpdateTool(updateToolOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(toolchainToolPatch).ToNot(BeNil())
		})
	})

	Describe(`DeleteTool - Delete a tool`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTool(deleteToolOptions *DeleteToolOptions)`, func() {
			deleteToolOptions := &cdtoolchainv2.DeleteToolOptions{
				ToolchainID: &toolchainIDLink,
				ToolID: &toolIDLink,
			}

			response, err := cdToolchainService.DeleteTool(deleteToolOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteToolchain - Delete a toolchain`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteToolchain(deleteToolchainOptions *DeleteToolchainOptions)`, func() {
			deleteToolchainOptions := &cdtoolchainv2.DeleteToolchainOptions{
				ToolchainID: &toolchainIDLink,
			}

			response, err := cdToolchainService.DeleteToolchain(deleteToolchainOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
