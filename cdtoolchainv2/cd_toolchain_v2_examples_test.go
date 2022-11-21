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

package cdtoolchainv2_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/continuous-delivery-go-sdk/cdtoolchainv2"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the CD Toolchain service.
//
// The following configuration properties are assumed to be defined:
// CD_TOOLCHAIN_URL=<service base url>
// CD_TOOLCHAIN_AUTH_TYPE=iam
// CD_TOOLCHAIN_APIKEY=<IAM apikey>
//
// These configuration properties can be exported as environment variables, or stored
// in the "../cd_toolchain_v2.env" configuration file as defined above
//
var _ = Describe(`CdToolchainV2 Examples Tests`, func() {

	const externalConfigFile = "../cd_toolchain_v2.env"

	var (
		cdToolchainService *cdtoolchainv2.CdToolchainV2
		config       map[string]string

		// Variables to hold link values
		toolIDLink      string
		toolchainIDLink string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(cdtoolchainv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			cdToolchainServiceOptions := &cdtoolchainv2.CdToolchainV2Options{}

			cdToolchainService, err = cdtoolchainv2.NewCdToolchainV2UsingExternalConfig(cdToolchainServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(cdToolchainService).ToNot(BeNil())
		})
	})

	Describe(`CdToolchainV2 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateToolchain request example`, func() {
			fmt.Println("\nCreateToolchain() result:")
			// begin-create_toolchain

			createToolchainOptions := cdToolchainService.NewCreateToolchainOptions(
				"TestToolchainV2",
				"6a9a01f2cff54a7f966f803d92877123",
			)

			toolchainPost, response, err := cdToolchainService.CreateToolchain(createToolchainOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(toolchainPost, "", "  ")
			fmt.Println(string(b))

			// end-create_toolchain

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(toolchainPost).ToNot(BeNil())

			toolchainIDLink = *toolchainPost.ID
			fmt.Fprintf(GinkgoWriter, "Saved toolchainIDLink value: %v\n", toolchainIDLink)
		})
		It(`CreateTool request example`, func() {
			fmt.Println("\nCreateTool() result:")
			// begin-create_tool

			createToolOptions := cdToolchainService.NewCreateToolOptions(
				toolchainIDLink,
				"draservicebroker",
			)

			toolchainToolPost, response, err := cdToolchainService.CreateTool(createToolOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(toolchainToolPost, "", "  ")
			fmt.Println(string(b))

			// end-create_tool

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(toolchainToolPost).ToNot(BeNil())

			toolIDLink = *toolchainToolPost.ID
			fmt.Fprintf(GinkgoWriter, "Saved toolIDLink value: %v\n", toolIDLink)
		})
		It(`ListToolchains request example`, func() {
			fmt.Println("\nListToolchains() result:")
			// begin-list_toolchains
			listToolchainsOptions := &cdtoolchainv2.ListToolchainsOptions{
				ResourceGroupID: core.StringPtr("testString"),
			}

			pager, err := cdToolchainService.NewToolchainsPager(listToolchainsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []cdtoolchainv2.ToolchainModel
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_toolchains
		})
		It(`GetToolchainByID request example`, func() {
			fmt.Println("\nGetToolchainByID() result:")
			// begin-get_toolchain_by_id

			getToolchainByIDOptions := cdToolchainService.NewGetToolchainByIDOptions(
				toolchainIDLink,
			)

			toolchain, response, err := cdToolchainService.GetToolchainByID(getToolchainByIDOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(toolchain, "", "  ")
			fmt.Println(string(b))

			// end-get_toolchain_by_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(toolchain).ToNot(BeNil())

		})
		It(`UpdateToolchain request example`, func() {
			fmt.Println("\nUpdateToolchain() result:")
			// begin-update_toolchain

			toolchainPrototypePatchModel := &cdtoolchainv2.ToolchainPrototypePatch{
			}
			toolchainPrototypePatchModelAsPatch, asPatchErr := toolchainPrototypePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateToolchainOptions := cdToolchainService.NewUpdateToolchainOptions(
				toolchainIDLink,
				toolchainPrototypePatchModelAsPatch,
			)

			toolchainPatch, response, err := cdToolchainService.UpdateToolchain(updateToolchainOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(toolchainPatch, "", "  ")
			fmt.Println(string(b))

			// end-update_toolchain

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(toolchainPatch).ToNot(BeNil())

		})
		It(`ListTools request example`, func() {
			fmt.Println("\nListTools() result:")
			// begin-list_tools
			listToolsOptions := &cdtoolchainv2.ListToolsOptions{
				ToolchainID: &toolchainIDLink,
			}

			pager, err := cdToolchainService.NewToolsPager(listToolsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []cdtoolchainv2.ToolModel
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_tools
		})
		It(`GetToolByID request example`, func() {
			fmt.Println("\nGetToolByID() result:")
			// begin-get_tool_by_id

			getToolByIDOptions := cdToolchainService.NewGetToolByIDOptions(
				toolchainIDLink,
				toolIDLink,
			)

			toolchainTool, response, err := cdToolchainService.GetToolByID(getToolByIDOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(toolchainTool, "", "  ")
			fmt.Println(string(b))

			// end-get_tool_by_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(toolchainTool).ToNot(BeNil())

		})
		It(`UpdateTool request example`, func() {
			fmt.Println("\nUpdateTool() result:")
			// begin-update_tool

			toolchainToolPrototypePatchModel := &cdtoolchainv2.ToolchainToolPrototypePatch{
			}
			toolchainToolPrototypePatchModelAsPatch, asPatchErr := toolchainToolPrototypePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateToolOptions := cdToolchainService.NewUpdateToolOptions(
				toolchainIDLink,
				toolIDLink,
				toolchainToolPrototypePatchModelAsPatch,
			)

			toolchainToolPatch, response, err := cdToolchainService.UpdateTool(updateToolOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(toolchainToolPatch, "", "  ")
			fmt.Println(string(b))

			// end-update_tool

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(toolchainToolPatch).ToNot(BeNil())

		})
		It(`DeleteTool request example`, func() {
			// begin-delete_tool

			deleteToolOptions := cdToolchainService.NewDeleteToolOptions(
				toolchainIDLink,
				toolIDLink,
			)

			response, err := cdToolchainService.DeleteTool(deleteToolOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteTool(): %d\n", response.StatusCode)
			}

			// end-delete_tool

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteToolchain request example`, func() {
			// begin-delete_toolchain

			deleteToolchainOptions := cdToolchainService.NewDeleteToolchainOptions(
				toolchainIDLink,
			)

			response, err := cdToolchainService.DeleteToolchain(deleteToolchainOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteToolchain(): %d\n", response.StatusCode)
			}

			// end-delete_toolchain

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
