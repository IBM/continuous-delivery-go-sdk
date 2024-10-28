// +build integration

/**
 * (C) Copyright IBM Corp. 2022, 2023.
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

package cdtektonpipelinev2_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/continuous-delivery-go-sdk/v2/cdtoolchainv2"
	"github.com/IBM/continuous-delivery-go-sdk/v2/cdtektonpipelinev2"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the cdtektonpipelinev2 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`CdTektonPipelineV2 Integration Tests`, func() {
	const externalConfigFile = "../cd_tekton_pipeline_v2.env"

	var (
		err              error
		cdToolchainService *cdtoolchainv2.CdToolchainV2
		cdTektonPipelineService *cdtektonpipelinev2.CdTektonPipelineV2
		toolchainURL     string
		pipelineURL      string
		toolchainConfig  map[string]string
		pipelineConfig   map[string]string

		// Variables to hold link values
		toolchainIDLink        string
		pipelineIDLink         string
		definitionIDLink       string
		triggerIDLink          string
		duplicateTriggerIDLink string
		runIDLink              string
		rerunIDLink            string
		runLogLink             cdtektonpipelinev2.Log
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
			toolchainConfig, err = core.GetServiceProperties(cdtoolchainv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			pipelineConfig, err = core.GetServiceProperties(cdtektonpipelinev2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			toolchainURL = toolchainConfig["URL"]
			// log.Printf("[INFO] === toolchainURL: %v\n", toolchainURL)
			if toolchainURL == "" {
				Skip("Unable to load toolchain service URL configuration property, skipping tests")
			}
			pipelineURL = pipelineConfig["URL"]
			// log.Printf("[INFO] === pipelineURL: %v\n", pipelineURL)
			if pipelineURL == "" {
				Skip("Unable to load pipeline service URL configuration property, skipping tests")
			}
			fmt.Fprintf(GinkgoWriter, "Toolchain Service URL: %v\n", toolchainURL)
			fmt.Fprintf(GinkgoWriter, "Pipeline Service URL: %v\n", pipelineURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Toolchain Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			cdToolchainServiceOptions := &cdtoolchainv2.CdToolchainV2Options{
				URL: toolchainURL,
			}

			cdToolchainService, err = cdtoolchainv2.NewCdToolchainV2UsingExternalConfig(cdToolchainServiceOptions)
			Expect(err).To(BeNil())
			Expect(cdToolchainService).ToNot(BeNil())
			Expect(cdToolchainService.Service.Options.URL).To(Equal(toolchainURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			cdToolchainService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateToolchain - Create a toolchain`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateToolchain(createToolchainOptions *CreateToolchainOptions)`, func() {
			rgId := pipelineConfig["RESOURCE_GROUP_ID"]
			now := time.Now()
			timeString := now.Format("20060102-1504")
			if rgId == "" {
				Skip("Unable to load resource group ID configuration property, skipping tests")
			}
			createToolchainOptions := &cdtoolchainv2.CreateToolchainOptions{
				Name: core.StringPtr("TestGoSDK-" + timeString),
				ResourceGroupID: &rgId,
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

	Describe(`Tekton Pipeline Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			cdTektonPipelineServiceOptions := &cdtektonpipelinev2.CdTektonPipelineV2Options{}

			cdTektonPipelineService, err = cdtektonpipelinev2.NewCdTektonPipelineV2UsingExternalConfig(cdTektonPipelineServiceOptions)
			Expect(err).To(BeNil())
			Expect(cdTektonPipelineService).ToNot(BeNil())
			Expect(cdTektonPipelineService.Service.Options.URL).To(Equal(pipelineURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			cdTektonPipelineService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateTool - Create a github integration tool`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTool(createToolOptions *CreateToolOptions)`, func() {
			params := make(map[string]interface{})
			params["repo_url"] = "https://github.com/open-toolchain/hello-tekton.git"
			params["type"] = "link"
			createToolOptions := &cdtoolchainv2.CreateToolOptions{
				ToolchainID: &toolchainIDLink,
				ToolTypeID: core.StringPtr("githubconsolidated"),
				Name: core.StringPtr("testString"),
				Parameters: params,
			}
			toolchainToolPost, response, err := cdToolchainService.CreateTool(createToolOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(toolchainToolPost).ToNot(BeNil())
		})
	})

	Describe(`CreateTool - Create a pipeline tool`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTool(createToolOptions *CreateToolOptions)`, func() {
			params := make(map[string]interface{})
			params["name"] =  "sdkIntegrationTestPipeline"
			params["type"] =  "tekton"
			createToolOptions := &cdtoolchainv2.CreateToolOptions{
				ToolchainID: &toolchainIDLink,
				ToolTypeID: core.StringPtr("pipeline"),
				Parameters: params,
			}

			toolchainPipelinePost, response, err := cdToolchainService.CreateTool(createToolOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(toolchainPipelinePost).ToNot(BeNil())

			pipelineIDLink = *toolchainPipelinePost.ID
			fmt.Fprintf(GinkgoWriter, "Saved pipelineIDLink value: %v\n", pipelineIDLink)
		})
	})

	Describe(`CreateTektonPipeline - Create Tekton pipeline`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTektonPipeline(createTektonPipelineOptions *CreateTektonPipelineOptions)`, func() {
			workerIdentityModel := &cdtektonpipelinev2.WorkerIdentity{
				ID: core.StringPtr("public"),
			}

			createTektonPipelineOptions := &cdtektonpipelinev2.CreateTektonPipelineOptions{
				EnableNotifications: core.BoolPtr(false),
				EnablePartialCloning: core.BoolPtr(false),
				ID: core.StringPtr(pipelineIDLink),
				Worker: workerIdentityModel,
			}

			tektonPipeline, response, err := cdTektonPipelineService.CreateTektonPipeline(createTektonPipelineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(tektonPipeline).ToNot(BeNil())
		})
	})

	Describe(`GetTektonPipeline - Get Tekton pipeline data`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTektonPipeline(getTektonPipelineOptions *GetTektonPipelineOptions)`, func() {
			getTektonPipelineOptions := &cdtektonpipelinev2.GetTektonPipelineOptions{
				ID: core.StringPtr(pipelineIDLink),
			}

			tektonPipeline, response, err := cdTektonPipelineService.GetTektonPipeline(getTektonPipelineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tektonPipeline).ToNot(BeNil())
		})
	})

	Describe(`UpdateTektonPipeline - Update Tekton pipeline data`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateTektonPipeline(updateTektonPipelineOptions *UpdateTektonPipelineOptions)`, func() {
			workerIdentityModel := &cdtektonpipelinev2.WorkerIdentity{
				ID: core.StringPtr("public"),
			}

			tektonPipelinePatchModel := &cdtektonpipelinev2.TektonPipelinePatch{
				EnableNotifications: core.BoolPtr(true),
				EnablePartialCloning: core.BoolPtr(true),
				Worker: workerIdentityModel,
			}
			tektonPipelinePatchModelAsPatch, asPatchErr := tektonPipelinePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateTektonPipelineOptions := &cdtektonpipelinev2.UpdateTektonPipelineOptions{
				ID: core.StringPtr(pipelineIDLink),
				TektonPipelinePatch: tektonPipelinePatchModelAsPatch,
			}

			tektonPipeline, response, err := cdTektonPipelineService.UpdateTektonPipeline(updateTektonPipelineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tektonPipeline).ToNot(BeNil())
		})
	})

	Describe(`ListTektonPipelineDefinitions - List pipeline definitions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTektonPipelineDefinitions(listTektonPipelineDefinitionsOptions *ListTektonPipelineDefinitionsOptions)`, func() {
			listTektonPipelineDefinitionsOptions := &cdtektonpipelinev2.ListTektonPipelineDefinitionsOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
			}

			definitionsCollection, response, err := cdTektonPipelineService.ListTektonPipelineDefinitions(listTektonPipelineDefinitionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(definitionsCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateTektonPipelineDefinition - Create a single definition`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptions *CreateTektonPipelineDefinitionOptions)`, func() {
			definitionSourcePropertiesModel := &cdtektonpipelinev2.DefinitionSourceProperties{
				URL: core.StringPtr("https://github.com/open-toolchain/hello-tekton.git"),
				Branch: core.StringPtr("testString"),
				Path: core.StringPtr("testString"),
			}

			definitionSourceModel := &cdtektonpipelinev2.DefinitionSource{
				Type: core.StringPtr("git"),
				Properties: definitionSourcePropertiesModel,
			}

			createTektonPipelineDefinitionOptions := cdTektonPipelineService.NewCreateTektonPipelineDefinitionOptions(
				pipelineIDLink,
				definitionSourceModel,
			)

			definition, response, err := cdTektonPipelineService.CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(definition).ToNot(BeNil())
			definitionIDLink = *definition.ID
			definitionSource := *definition.Source
			definitionSourceProps := *definitionSource.Properties
			definitionURL := *definitionSourceProps.URL
			definitionBranch := *definitionSourceProps.Branch
			definitionPath := *definitionSourceProps.Path
			Expect(definitionPath).To(Equal("testString"))
			Expect(definitionBranch).To(Equal("testString"))
			Expect(definitionURL).To(Equal("https://github.com/open-toolchain/hello-tekton.git"))
		})
	})

	Describe(`GetTektonPipelineDefinition - Retrieve a single definition entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTektonPipelineDefinition(getTektonPipelineDefinitionOptions *GetTektonPipelineDefinitionOptions)`, func() {
			getTektonPipelineDefinitionOptions := &cdtektonpipelinev2.GetTektonPipelineDefinitionOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				DefinitionID: core.StringPtr(definitionIDLink),
			}

			definition, response, err := cdTektonPipelineService.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(definition).ToNot(BeNil())
			definitionSource := *definition.Source
			definitionSourceProps := *definitionSource.Properties
			definitionURL := *definitionSourceProps.URL
			definitionBranch := *definitionSourceProps.Branch
			definitionPath := *definitionSourceProps.Path
			Expect(definitionPath).To(Equal("testString"))
			Expect(definitionBranch).To(Equal("testString"))
			Expect(definitionURL).To(Equal("https://github.com/open-toolchain/hello-tekton.git"))
		})
	})

	Describe(`ReplaceTektonPipelineDefinition - Edit a single definition entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceTektonPipelineDefinition(replaceTektonPipelineDefinitionOptions *ReplaceTektonPipelineDefinitionOptions)`, func() {
			definitionSourcePropertiesModel := &cdtektonpipelinev2.DefinitionSourceProperties{
				URL: core.StringPtr("https://github.com/open-toolchain/hello-tekton.git"),
				Branch: core.StringPtr("master"),
				Path: core.StringPtr(".tekton"),
			}

			definitionSourceModel := &cdtektonpipelinev2.DefinitionSource{
				Type: core.StringPtr("git"),
				Properties: definitionSourcePropertiesModel,
			}

			replaceTektonPipelineDefinitionOptions := &cdtektonpipelinev2.ReplaceTektonPipelineDefinitionOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				DefinitionID: core.StringPtr(definitionIDLink),
				Source: definitionSourceModel,
			}

			definition, response, err := cdTektonPipelineService.ReplaceTektonPipelineDefinition(replaceTektonPipelineDefinitionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(definition).ToNot(BeNil())
			definitionSource := *definition.Source
			definitionSourceProps := *definitionSource.Properties
			definitionURL := *definitionSourceProps.URL
			definitionBranch := *definitionSourceProps.Branch
			definitionPath := *definitionSourceProps.Path
			Expect(definitionPath).To(Equal(".tekton"))
			Expect(definitionBranch).To(Equal("master"))
			Expect(definitionURL).To(Equal("https://github.com/open-toolchain/hello-tekton.git"))
		})
	})

	Describe(`ListTektonPipelineProperties - List the pipeline's environment properties`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTektonPipelineProperties(listTektonPipelinePropertiesOptions *ListTektonPipelinePropertiesOptions)`, func() {
			listTektonPipelinePropertiesOptions := &cdtektonpipelinev2.ListTektonPipelinePropertiesOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				Sort: core.StringPtr("name"),
			}

			propertiesCollection, response, err := cdTektonPipelineService.ListTektonPipelineProperties(listTektonPipelinePropertiesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(propertiesCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateTektonPipelineProperties - Create a pipeline environment property`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTektonPipelineProperties(createTektonPipelinePropertiesOptions *CreateTektonPipelinePropertiesOptions)`, func() {
			createTektonPipelinePropertiesOptions := &cdtektonpipelinev2.CreateTektonPipelinePropertiesOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				Name: core.StringPtr("prop1"),
				Value: core.StringPtr("https://github.com/open-toolchain/hello-tekton.git"),
				Type: core.StringPtr("text"),
			}

			property, response, err := cdTektonPipelineService.CreateTektonPipelineProperties(createTektonPipelinePropertiesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(property).ToNot(BeNil())
			propName := *property.Name
			propType := *property.Type
			propValue := *property.Value
			Expect(propName).To(Equal("prop1"))
			Expect(propType).To(Equal("text"))
			Expect(propValue).To(Equal("https://github.com/open-toolchain/hello-tekton.git"))
		})
	})

	Describe(`GetTektonPipelineProperty - Get a pipeline environment property`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTektonPipelineProperty(getTektonPipelinePropertyOptions *GetTektonPipelinePropertyOptions)`, func() {
			getTektonPipelinePropertyOptions := &cdtektonpipelinev2.GetTektonPipelinePropertyOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				PropertyName: core.StringPtr("prop1"),
			}

			property, response, err := cdTektonPipelineService.GetTektonPipelineProperty(getTektonPipelinePropertyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(property).ToNot(BeNil())
		})
	})

	Describe(`ReplaceTektonPipelineProperty - Replace the value of an environment property`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceTektonPipelineProperty(replaceTektonPipelinePropertyOptions *ReplaceTektonPipelinePropertyOptions)`, func() {
			replaceTektonPipelinePropertyOptions := &cdtektonpipelinev2.ReplaceTektonPipelinePropertyOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				PropertyName: core.StringPtr("prop1"),
				Name: core.StringPtr("prop1"),
				Type: core.StringPtr("text"),
				Value: core.StringPtr("editedValue"),
			}

			property, response, err := cdTektonPipelineService.ReplaceTektonPipelineProperty(replaceTektonPipelinePropertyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(property).ToNot(BeNil())
			propName := *property.Name
			propType := *property.Type
			propValue := *property.Value
			Expect(propName).To(Equal("prop1"))
			Expect(propType).To(Equal("text"))
			Expect(propValue).To(Equal("editedValue"))
		})
	})

	Describe(`ListTektonPipelineTriggers - List pipeline triggers`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTektonPipelineTriggers(listTektonPipelineTriggersOptions *ListTektonPipelineTriggersOptions)`, func() {
			listTektonPipelineTriggersOptions := &cdtektonpipelinev2.ListTektonPipelineTriggersOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
			}

			triggersCollection, response, err := cdTektonPipelineService.ListTektonPipelineTriggers(listTektonPipelineTriggersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(triggersCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateTektonPipelineTrigger - Create a trigger`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTektonPipelineTrigger(createTektonPipelineTriggerOptions *CreateTektonPipelineTriggerOptions)`, func() {
			workerModel := &cdtektonpipelinev2.WorkerIdentity{
				ID: core.StringPtr("public"),
			}

			createTektonPipelineTriggerOptions := &cdtektonpipelinev2.CreateTektonPipelineTriggerOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				Type: core.StringPtr("manual"),
				Name: core.StringPtr("ManualTrigger"),
				EventListener: core.StringPtr("listener"),
				Tags: []string{"tag1"},
				Worker: workerModel,
				MaxConcurrentRuns: core.Int64Ptr(int64(2 )),
			}

			trigger, response, err := cdTektonPipelineService.CreateTektonPipelineTrigger(createTektonPipelineTriggerOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(trigger).ToNot(BeNil())
			triggerModel := trigger.(*cdtektonpipelinev2.Trigger)
			triggerIDLink = *triggerModel.ID
			triggerName := *triggerModel.Name
			triggerType := *triggerModel.Type
			triggerEventListener := *triggerModel.EventListener
			triggerMaxConcurrentRuns := *triggerModel.MaxConcurrentRuns
			triggerWorker := *triggerModel.Worker
			triggerWorkerId := *triggerWorker.ID
			Expect(triggerName).To(Equal("ManualTrigger"))
			Expect(triggerType).To(Equal("manual"))
			Expect(triggerEventListener).To(Equal("listener"))
			Expect(triggerMaxConcurrentRuns).To(Equal(int64(2)))
			Expect(triggerWorkerId).To(Equal("public"))
		})
	})

	Describe(`CreateTektonPipelineRun - Trigger a pipeline run`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTektonPipelineRun(createTektonPipelineRunOptions *CreateTektonPipelineRunOptions)`, func() {
			runTriggerModel := new(cdtektonpipelinev2.PipelineRunTrigger)
			runTriggerModel.Name = core.StringPtr("ManualTrigger")
			runTriggerModel.Properties = map[string]interface{}{"addedProp1": "addedValue", "triggerProp1": "overrideValue"}
			runTriggerModel.SecureProperties = map[string]interface{}{"triggerSecProp1": "overrideValue"}
			runTriggerModel.HeadersVar = map[string]interface{}{"x-custom-header": "x-custom-value"}
			runTriggerModel.Body = map[string]interface{}{"bodyKey": "bodyValue"}
			createTektonPipelineRunOptions := &cdtektonpipelinev2.CreateTektonPipelineRunOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				Trigger: runTriggerModel,
			}
			pipelineRun, response, err := cdTektonPipelineService.CreateTektonPipelineRun(createTektonPipelineRunOptions)
			runIDLink = *pipelineRun.ID
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(pipelineRun).ToNot(BeNil())
			trigger := pipelineRun.Trigger
			triggerModel := *trigger.(*cdtektonpipelinev2.Trigger)
			triggerID := *triggerModel.ID
			triggerName := *triggerModel.Name
			runStatus := *pipelineRun.Status
			runWorkerName := *pipelineRun.Worker.Name
			runPipelineID := *pipelineRun.PipelineID
			runListenerName := *pipelineRun.ListenerName
			runCreatedAt := *pipelineRun.CreatedAt
			runUpdatedAt := *pipelineRun.UpdatedAt
			runURL := *pipelineRun.RunURL
			properties := pipelineRun.Properties
			eventParams := pipelineRun.EventParamsBlob
			triggerHeaders := pipelineRun.TriggerHeaders
			Expect(runPipelineID).To(Equal(pipelineIDLink))
			Expect(triggerID).To(Equal(triggerIDLink))
			Expect(triggerName).To(Equal("ManualTrigger"))
			Expect(runStatus).ToNot(BeNil())
			Expect(runWorkerName).To(Equal("public"))
			Expect(runListenerName).To(Equal("listener"))
			Expect(runCreatedAt).ToNot(BeNil())
			Expect(runUpdatedAt).ToNot(BeNil())
			Expect(runURL).ToNot(BeNil())
			Expect(properties).ToNot(BeNil())
			Expect(eventParams).ToNot(BeNil())
			Expect(triggerHeaders).ToNot(BeNil())
		})
	})

	Describe(`GetTektonPipelineTrigger - Get a single trigger`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTektonPipelineTrigger(getTektonPipelineTriggerOptions *GetTektonPipelineTriggerOptions)`, func() {
			getTektonPipelineTriggerOptions := &cdtektonpipelinev2.GetTektonPipelineTriggerOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				TriggerID: core.StringPtr(triggerIDLink),
			}

			trigger, response, err := cdTektonPipelineService.GetTektonPipelineTrigger(getTektonPipelineTriggerOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trigger).ToNot(BeNil())
			triggerModel := trigger.(*cdtektonpipelinev2.Trigger)
			triggerName := *triggerModel.Name
			triggerType := *triggerModel.Type
			triggerEventListener := *triggerModel.EventListener
			triggerMaxConcurrentRuns := *triggerModel.MaxConcurrentRuns
			triggerWorker := *triggerModel.Worker
			triggerWorkerId := *triggerWorker.ID
			Expect(triggerName).To(Equal("ManualTrigger"))
			Expect(triggerType).To(Equal("manual"))
			Expect(triggerEventListener).To(Equal("listener"))
			Expect(triggerMaxConcurrentRuns).To(Equal(int64(2)))
			Expect(triggerWorkerId).To(Equal("public"))
		})
	})

	Describe(`UpdateTektonPipelineTrigger - Edit a trigger`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateTektonPipelineTrigger(updateTektonPipelineTriggerOptions *UpdateTektonPipelineTriggerOptions)`, func() {
			triggerPatchModel := &cdtektonpipelinev2.TriggerPatch{
				Type: core.StringPtr("manual"),
				Name: core.StringPtr("start-deploy"),
				EventListener: core.StringPtr("listener"),
				Tags: []string{"edit1"},
				MaxConcurrentRuns: core.Int64Ptr(int64(4)),
			}
			triggerPatchModelAsPatch, asPatchErr := triggerPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateTektonPipelineTriggerOptions := &cdtektonpipelinev2.UpdateTektonPipelineTriggerOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				TriggerID: core.StringPtr(triggerIDLink),
				TriggerPatch: triggerPatchModelAsPatch,
			}

			trigger, response, err := cdTektonPipelineService.UpdateTektonPipelineTrigger(updateTektonPipelineTriggerOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trigger).ToNot(BeNil())
			triggerModel := trigger.(*cdtektonpipelinev2.Trigger)
			triggerName := *triggerModel.Name
			triggerType := *triggerModel.Type
			triggerEventListener := *triggerModel.EventListener
			triggerMaxConcurrentRuns := *triggerModel.MaxConcurrentRuns
			triggerWorker := *triggerModel.Worker
			triggerWorkerId := *triggerWorker.ID
			Expect(triggerName).To(Equal("start-deploy"))
			Expect(triggerType).To(Equal("manual"))
			Expect(triggerEventListener).To(Equal("listener"))
			Expect(triggerMaxConcurrentRuns).To(Equal(int64(4)))
			Expect(triggerWorkerId).To(Equal("public"))
		})
	})

	Describe(`DuplicateTektonPipelineTrigger - Duplicate a trigger`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DuplicateTektonPipelineTrigger(duplicateTektonPipelineTriggerOptions *DuplicateTektonPipelineTriggerOptions)`, func() {
			duplicateTektonPipelineTriggerOptions := &cdtektonpipelinev2.DuplicateTektonPipelineTriggerOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				SourceTriggerID: core.StringPtr(triggerIDLink),
				Name: core.StringPtr("start-deploy-copy"),
			}

			trigger, response, err := cdTektonPipelineService.DuplicateTektonPipelineTrigger(duplicateTektonPipelineTriggerOptions)
			triggerModel := trigger.(*cdtektonpipelinev2.Trigger)
			duplicateTriggerIDLink = *triggerModel.ID
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(trigger).ToNot(BeNil())
		})
	})

	Describe(`ListTektonPipelineTriggerProperties - List trigger properties`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTektonPipelineTriggerProperties(listTektonPipelineTriggerPropertiesOptions *ListTektonPipelineTriggerPropertiesOptions)`, func() {
			listTektonPipelineTriggerPropertiesOptions := &cdtektonpipelinev2.ListTektonPipelineTriggerPropertiesOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				TriggerID: core.StringPtr(triggerIDLink),
				Name: core.StringPtr("prop"),
				Sort: core.StringPtr("name"),
				Type: core.StringPtr("text"),
			}

			triggerPropertiesCollection, response, err := cdTektonPipelineService.ListTektonPipelineTriggerProperties(listTektonPipelineTriggerPropertiesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(triggerPropertiesCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateTektonPipelineTriggerProperties - Create a trigger property`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTektonPipelineTriggerProperties(createTektonPipelineTriggerPropertiesOptions *CreateTektonPipelineTriggerPropertiesOptions)`, func() {
			createTektonPipelineTriggerPropertiesOptions := &cdtektonpipelinev2.CreateTektonPipelineTriggerPropertiesOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				TriggerID: core.StringPtr(triggerIDLink),
				Name: core.StringPtr("triggerProp1"),
				Value: core.StringPtr("triggerPropValue1"),
				Type: core.StringPtr("text"),
			}

			triggerProperty, response, err := cdTektonPipelineService.CreateTektonPipelineTriggerProperties(createTektonPipelineTriggerPropertiesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(triggerProperty).ToNot(BeNil())
			triggerPropName := *triggerProperty.Name
			triggerPropType := *triggerProperty.Type
			triggerPropValue := *triggerProperty.Value
			Expect(triggerPropName).To(Equal("triggerProp1"))
			Expect(triggerPropType).To(Equal("text"))
			Expect(triggerPropValue).To(Equal("triggerPropValue1"))
		})
	})

	Describe(`GetTektonPipelineTriggerProperty - Get a trigger property`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTektonPipelineTriggerProperty(getTektonPipelineTriggerPropertyOptions *GetTektonPipelineTriggerPropertyOptions)`, func() {
			getTektonPipelineTriggerPropertyOptions := &cdtektonpipelinev2.GetTektonPipelineTriggerPropertyOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				TriggerID: core.StringPtr(triggerIDLink),
				PropertyName: core.StringPtr("triggerProp1"),
			}

			triggerProperty, response, err := cdTektonPipelineService.GetTektonPipelineTriggerProperty(getTektonPipelineTriggerPropertyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(triggerProperty).ToNot(BeNil())
			triggerPropName := *triggerProperty.Name
			triggerPropType := *triggerProperty.Type
			triggerPropValue := *triggerProperty.Value
			Expect(triggerPropName).To(Equal("triggerProp1"))
			Expect(triggerPropType).To(Equal("text"))
			Expect(triggerPropValue).To(Equal("triggerPropValue1"))
		})
	})

	Describe(`ReplaceTektonPipelineTriggerProperty - Replace a trigger property value`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceTektonPipelineTriggerProperty(replaceTektonPipelineTriggerPropertyOptions *ReplaceTektonPipelineTriggerPropertyOptions)`, func() {
			replaceTektonPipelineTriggerPropertyOptions := &cdtektonpipelinev2.ReplaceTektonPipelineTriggerPropertyOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				TriggerID: core.StringPtr(triggerIDLink),
				PropertyName: core.StringPtr("triggerProp1"),
				Name: core.StringPtr("triggerProp1"),
				Value: core.StringPtr("editedValue"),
				Type: core.StringPtr("text"),
			}

			triggerProperty, response, err := cdTektonPipelineService.ReplaceTektonPipelineTriggerProperty(replaceTektonPipelineTriggerPropertyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(triggerProperty).ToNot(BeNil())
			triggerPropName := *triggerProperty.Name
			triggerPropType := *triggerProperty.Type
			triggerPropValue := *triggerProperty.Value
			Expect(triggerPropName).To(Equal("triggerProp1"))
			Expect(triggerPropType).To(Equal("text"))
			Expect(triggerPropValue).To(Equal("editedValue"))
		})
	})

	Describe(`DeleteTektonPipelineTriggerProperty - Delete a trigger property`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTektonPipelineTriggerProperty(deleteTektonPipelineTriggerPropertyOptions *DeleteTektonPipelineTriggerPropertyOptions)`, func() {
			deleteTektonPipelineTriggerPropertyOptions := &cdtektonpipelinev2.DeleteTektonPipelineTriggerPropertyOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				TriggerID: core.StringPtr(triggerIDLink),
				PropertyName: core.StringPtr("triggerProp1"),
			}

			response, err := cdTektonPipelineService.DeleteTektonPipelineTriggerProperty(deleteTektonPipelineTriggerPropertyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteTektonPipelineTrigger - Delete a single trigger`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTektonPipelineTrigger(deleteTektonPipelineTriggerOptions *DeleteTektonPipelineTriggerOptions)`, func() {
			deleteTektonPipelineTriggerOptions := &cdtektonpipelinev2.DeleteTektonPipelineTriggerOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				TriggerID: core.StringPtr(duplicateTriggerIDLink),
			}

			response, err := cdTektonPipelineService.DeleteTektonPipelineTrigger(deleteTektonPipelineTriggerOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`GetTektonPipelineRun - Get a pipeline run record`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTektonPipelineRun(getTektonPipelineRunOptions *GetTektonPipelineRunOptions)`, func() {
			getTektonPipelineRunOptions := &cdtektonpipelinev2.GetTektonPipelineRunOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				ID: core.StringPtr(runIDLink),
			}

			pipelineRun, response, err := cdTektonPipelineService.GetTektonPipelineRun(getTektonPipelineRunOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pipelineRun).ToNot(BeNil())
			trigger := pipelineRun.Trigger
			triggerModel := *trigger.(*cdtektonpipelinev2.Trigger)
			triggerID := *triggerModel.ID
			triggerName := *triggerModel.Name
			runStatus := *pipelineRun.Status
			runWorkerName := *pipelineRun.Worker.Name
			runPipelineID := *pipelineRun.PipelineID
			runListenerName := *pipelineRun.ListenerName
			runCreatedAt := *pipelineRun.CreatedAt
			runUpdatedAt := *pipelineRun.UpdatedAt
			runURL := *pipelineRun.RunURL
			Expect(runPipelineID).To(Equal(pipelineIDLink))
			Expect(triggerID).To(Equal(triggerIDLink))
			Expect(triggerName).To(Equal("ManualTrigger"))
			Expect(runStatus).ToNot(BeNil())
			Expect(runWorkerName).To(Equal("public"))
			Expect(runListenerName).To(Equal("listener"))
			Expect(runCreatedAt).ToNot(BeNil())
			Expect(runUpdatedAt).ToNot(BeNil())
			Expect(runURL).ToNot(BeNil())
		})
	})

	Describe(`ListTektonPipelineRuns - List pipeline run records`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTektonPipelineRuns(listTektonPipelineRunsOptions *ListTektonPipelineRunsOptions) with pagination`, func(){
			listTektonPipelineRunsOptions := &cdtektonpipelinev2.ListTektonPipelineRunsOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
			}

			listTektonPipelineRunsOptions.Start = nil
			listTektonPipelineRunsOptions.Limit = core.Int64Ptr(1)

			var allResults []cdtektonpipelinev2.PipelineRun
			for {
				pipelineRunsCollection, response, err := cdTektonPipelineService.ListTektonPipelineRuns(listTektonPipelineRunsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(pipelineRunsCollection).ToNot(BeNil())
				runsFirst := *pipelineRunsCollection.First
				runsLast := *pipelineRunsCollection.Last
				runsLimit := *pipelineRunsCollection.Limit
				Expect(runsFirst).ToNot(BeNil())
				Expect(runsLast).ToNot(BeNil())
				Expect(runsLimit).To(Equal(int64(1)))
				allResults = append(allResults, pipelineRunsCollection.PipelineRuns...)

				listTektonPipelineRunsOptions.Start, err = pipelineRunsCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listTektonPipelineRunsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})

		It(`ListTektonPipelineRuns(listTektonPipelineRunsOptions *ListTektonPipelineRunsOptions) using TektonPipelineRunsPager`, func(){
			listTektonPipelineRunsOptions := &cdtektonpipelinev2.ListTektonPipelineRunsOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
			}

			// Test GetNext().
			pager, err := cdTektonPipelineService.NewTektonPipelineRunsPager(listTektonPipelineRunsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []cdtektonpipelinev2.PipelineRun
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = cdTektonPipelineService.NewTektonPipelineRunsPager(listTektonPipelineRunsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListTektonPipelineRuns() returned a total of %d item(s) using TektonPipelineRunsPager.\n", len(allResults))
		})
	})
	
	Describe(`RerunTektonPipelineRun - Rerun a pipeline run`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RerunTektonPipelineRun(rerunTektonPipelineRunOptions *RerunTektonPipelineRunOptions)`, func() {
			rerunTektonPipelineRunOptions := &cdtektonpipelinev2.RerunTektonPipelineRunOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				ID: core.StringPtr(runIDLink),
			}

			pipelineRun, response, err := cdTektonPipelineService.RerunTektonPipelineRun(rerunTektonPipelineRunOptions)
			rerunIDLink = *pipelineRun.ID
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(pipelineRun).ToNot(BeNil())
			trigger := pipelineRun.Trigger
			triggerModel := *trigger.(*cdtektonpipelinev2.Trigger)
			triggerID := *triggerModel.ID
			triggerName := *triggerModel.Name
			runStatus := *pipelineRun.Status
			runWorkerName := *pipelineRun.Worker.Name
			runPipelineID := *pipelineRun.PipelineID
			runListenerName := *pipelineRun.ListenerName
			runCreatedAt := *pipelineRun.CreatedAt
			runUpdatedAt := *pipelineRun.UpdatedAt
			runURL := *pipelineRun.RunURL
			Expect(runPipelineID).To(Equal(pipelineIDLink))
			Expect(triggerID).To(Equal(triggerIDLink))
			Expect(triggerName).To(Equal("ManualTrigger"))
			Expect(runStatus).ToNot(BeNil())
			Expect(runWorkerName).To(Equal("public"))
			Expect(runListenerName).To(Equal("listener"))
			Expect(runCreatedAt).ToNot(BeNil())
			Expect(runUpdatedAt).ToNot(BeNil())
			Expect(runURL).ToNot(BeNil())
		})
	})

	Describe(`CancelTektonPipelineRun - Cancel a pipeline run`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CancelTektonPipelineRun(cancelTektonPipelineRunOptions *CancelTektonPipelineRunOptions)`, func() {
			cancelTektonPipelineRunOptions := &cdtektonpipelinev2.CancelTektonPipelineRunOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				ID: core.StringPtr(rerunIDLink),
				Force: core.BoolPtr(true),
			}

			pipelineRun, response, err := cdTektonPipelineService.CancelTektonPipelineRun(cancelTektonPipelineRunOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(pipelineRun).ToNot(BeNil())
		})
	})

	Describe(`DeleteTektonPipelineProperty - Delete a single pipeline environment property`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTektonPipelineProperty(deleteTektonPipelinePropertyOptions *DeleteTektonPipelinePropertyOptions)`, func() {
			deleteTektonPipelinePropertyOptions := &cdtektonpipelinev2.DeleteTektonPipelinePropertyOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				PropertyName: core.StringPtr("prop1"),
			}

			response, err := cdTektonPipelineService.DeleteTektonPipelineProperty(deleteTektonPipelinePropertyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`GetTektonPipelineRunLogs - Get a list of pipeline run log objects`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTektonPipelineRunLogs(getTektonPipelineRunLogsOptions *GetTektonPipelineRunLogsOptions)`, func() {
			getTektonPipelineRunLogsOptions := &cdtektonpipelinev2.GetTektonPipelineRunLogsOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				ID: core.StringPtr(runIDLink),
			}

			logsCollection, response, err := cdTektonPipelineService.GetTektonPipelineRunLogs(getTektonPipelineRunLogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(logsCollection).ToNot(BeNil())
			logs := logsCollection.Logs
			Expect(logs).ToNot(BeNil())
			runLogLink = logs[1]
		})
	})

	Describe(`GetTektonPipelineRunLogContent - Get the log content of a pipeline run step`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTektonPipelineRunLogContent(getTektonPipelineRunLogContentOptions *GetTektonPipelineRunLogContentOptions)`, func() {
			getTektonPipelineRunLogContentOptions := &cdtektonpipelinev2.GetTektonPipelineRunLogContentOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				PipelineRunID: core.StringPtr(runIDLink),
				ID: core.StringPtr(*runLogLink.ID),
			}

			stepLog, response, err := cdTektonPipelineService.GetTektonPipelineRunLogContent(getTektonPipelineRunLogContentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(stepLog).ToNot(BeNil())
			logData := *stepLog.Data
			// log.Printf("[DEBUG] === logData: %+v\n", logData)
			Expect(logData).ToNot(BeNil())
		})
	})

	Describe(`DeleteTektonPipelineRun - Delete a pipeline run record`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTektonPipelineRun(deleteTektonPipelineRunOptions *DeleteTektonPipelineRunOptions)`, func() {
			deleteTektonPipelineRunOptions := &cdtektonpipelinev2.DeleteTektonPipelineRunOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				ID: core.StringPtr(runIDLink),
			}

			response, err := cdTektonPipelineService.DeleteTektonPipelineRun(deleteTektonPipelineRunOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteTektonPipelineDefinition - Delete a single definition entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTektonPipelineDefinition(deleteTektonPipelineDefinitionOptions *DeleteTektonPipelineDefinitionOptions)`, func() {
			deleteTektonPipelineDefinitionOptions := &cdtektonpipelinev2.DeleteTektonPipelineDefinitionOptions{
				PipelineID: core.StringPtr(pipelineIDLink),
				DefinitionID: core.StringPtr(definitionIDLink),
			}

			response, err := cdTektonPipelineService.DeleteTektonPipelineDefinition(deleteTektonPipelineDefinitionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteTektonPipeline - Delete Tekton pipeline instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTektonPipeline(deleteTektonPipelineOptions *DeleteTektonPipelineOptions)`, func() {
			deleteTektonPipelineOptions := &cdtektonpipelinev2.DeleteTektonPipelineOptions{
				ID: core.StringPtr(pipelineIDLink),
			}

			response, err := cdTektonPipelineService.DeleteTektonPipeline(deleteTektonPipelineOptions)
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
