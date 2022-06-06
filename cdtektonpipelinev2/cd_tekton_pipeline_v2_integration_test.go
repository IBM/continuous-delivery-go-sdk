// +build integration

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

package cdtektonpipelinev2_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/org-ids/tekton-pipeline-go-sdk/cdtektonpipelinev2"
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
		err          error
		cdTektonPipelineService *cdtektonpipelinev2.CdTektonPipelineV2
		serviceURL   string
		config       map[string]string
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
			config, err = core.GetServiceProperties(cdtektonpipelinev2.DefaultServiceName)
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

			cdTektonPipelineServiceOptions := &cdtektonpipelinev2.CdTektonPipelineV2Options{}

			cdTektonPipelineService, err = cdtektonpipelinev2.NewCdTektonPipelineV2UsingExternalConfig(cdTektonPipelineServiceOptions)

			Expect(err).To(BeNil())
			Expect(cdTektonPipelineService).ToNot(BeNil())
			Expect(cdTektonPipelineService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			cdTektonPipelineService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateTektonPipeline - Create tekton pipeline`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTektonPipeline(createTektonPipelineOptions *CreateTektonPipelineOptions)`, func() {

			workerWithIDModel := &cdtektonpipelinev2.WorkerWithID{
				ID: core.StringPtr("public"),
			}

			createTektonPipelineOptions := &cdtektonpipelinev2.CreateTektonPipelineOptions{
				ID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				Worker: workerWithIDModel,
			}

			tektonPipeline, response, err := cdTektonPipelineService.CreateTektonPipeline(createTektonPipelineOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(tektonPipeline).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 404
			//
		})
	})

	Describe(`GetTektonPipeline - Get tekton pipeline data`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTektonPipeline(getTektonPipelineOptions *GetTektonPipelineOptions)`, func() {

			getTektonPipelineOptions := &cdtektonpipelinev2.GetTektonPipelineOptions{
				ID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
			}

			tektonPipeline, response, err := cdTektonPipelineService.GetTektonPipeline(getTektonPipelineOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tektonPipeline).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})

	Describe(`UpdateTektonPipeline - Update tekton pipeline data`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateTektonPipeline(updateTektonPipelineOptions *UpdateTektonPipelineOptions)`, func() {

			workerWithIDModel := &cdtektonpipelinev2.WorkerWithID{
				ID: core.StringPtr("public"),
			}

			updateTektonPipelineOptions := &cdtektonpipelinev2.UpdateTektonPipelineOptions{
				ID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				Worker: workerWithIDModel,
			}

			tektonPipeline, response, err := cdTektonPipelineService.UpdateTektonPipeline(updateTektonPipelineOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tektonPipeline).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})

	Describe(`ListTektonPipelineRuns - List pipeline run records`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTektonPipelineRuns(listTektonPipelineRunsOptions *ListTektonPipelineRunsOptions) with pagination`, func(){
			var result cdtektonpipelinev2.PipelineRuns

			listTektonPipelineRunsOptions := &cdtektonpipelinev2.ListTektonPipelineRunsOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				Limit: core.Int64Ptr(int64(1)),
				Offset: core.Int64Ptr(int64(38)),
				Status: core.StringPtr("succeeded"),
				TriggerName: core.StringPtr("manual-trigger"),
			}

			listTektonPipelineRunsOptions.Offset = nil
			listTektonPipelineRunsOptions.Limit = core.Int64Ptr(1)

			for {
				pipelineRuns, response, err := cdTektonPipelineService.ListTektonPipelineRuns(listTektonPipelineRunsOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(pipelineRuns).ToNot(BeNil())
				result.PipelineRuns = append(result.PipelineRuns, pipelineRuns.PipelineRuns...)

				listTektonPipelineRunsOptions.Offset, err = pipelineRuns.GetNextOffset()
				Expect(err).To(BeNil())

				if listTektonPipelineRunsOptions.Offset == nil {
					break
				}

				//
				// The following status codes aren't covered by tests.
				// Please provide integration tests for these too.
				//
				// 401
				// 404
				//
			}
		})
	})

	Describe(`CreateTektonPipelineRun - Start a trigger to create a pipelineRun`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTektonPipelineRun(createTektonPipelineRunOptions *CreateTektonPipelineRunOptions)`, func() {

			createTektonPipelineRunOptions := &cdtektonpipelinev2.CreateTektonPipelineRunOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				TriggerName: core.StringPtr("Generic Webhook Trigger - 0"),
				TriggerProperties: make(map[string]interface{}),
				SecureTriggerProperties: make(map[string]interface{}),
				TriggerHeader: make(map[string]interface{}),
				TriggerBody: make(map[string]interface{}),
			}

			pipelineRun, response, err := cdTektonPipelineService.CreateTektonPipelineRun(createTektonPipelineRunOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(pipelineRun).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 404
			//
		})
	})

	Describe(`GetTektonPipelineRun - Get a pipeline run record`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTektonPipelineRun(getTektonPipelineRunOptions *GetTektonPipelineRunOptions)`, func() {

			getTektonPipelineRunOptions := &cdtektonpipelinev2.GetTektonPipelineRunOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				ID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				Includes: core.StringPtr("definitions"),
			}

			pipelineRun, response, err := cdTektonPipelineService.GetTektonPipelineRun(getTektonPipelineRunOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pipelineRun).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})

	Describe(`CancelTektonPipelineRun - Cancel a pipeline run`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CancelTektonPipelineRun(cancelTektonPipelineRunOptions *CancelTektonPipelineRunOptions)`, func() {

			cancelTektonPipelineRunOptions := &cdtektonpipelinev2.CancelTektonPipelineRunOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				ID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				Force: core.BoolPtr(true),
			}

			pipelineRun, response, err := cdTektonPipelineService.CancelTektonPipelineRun(cancelTektonPipelineRunOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pipelineRun).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 404
			//
		})
	})

	Describe(`RerunTektonPipelineRun - Rerun a pipeline run`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RerunTektonPipelineRun(rerunTektonPipelineRunOptions *RerunTektonPipelineRunOptions)`, func() {

			rerunTektonPipelineRunOptions := &cdtektonpipelinev2.RerunTektonPipelineRunOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				ID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
			}

			pipelineRun, response, err := cdTektonPipelineService.RerunTektonPipelineRun(rerunTektonPipelineRunOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(pipelineRun).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})

	Describe(`GetTektonPipelineRunLogs - Get a list of pipeline run log IDs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTektonPipelineRunLogs(getTektonPipelineRunLogsOptions *GetTektonPipelineRunLogsOptions)`, func() {

			getTektonPipelineRunLogsOptions := &cdtektonpipelinev2.GetTektonPipelineRunLogsOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				ID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
			}

			pipelineRunLogs, response, err := cdTektonPipelineService.GetTektonPipelineRunLogs(getTektonPipelineRunLogsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pipelineRunLogs).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})

	Describe(`GetTektonPipelineRunLogContent - Get the log content of a pipeline run step by using the step log ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTektonPipelineRunLogContent(getTektonPipelineRunLogContentOptions *GetTektonPipelineRunLogContentOptions)`, func() {

			getTektonPipelineRunLogContentOptions := &cdtektonpipelinev2.GetTektonPipelineRunLogContentOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				PipelineRunID: core.StringPtr("bf4b3abd-0c93-416b-911e-9cf42f1a1085"),
				ID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
			}

			stepLog, response, err := cdTektonPipelineService.GetTektonPipelineRunLogContent(getTektonPipelineRunLogContentOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(stepLog).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})

	Describe(`ListTektonPipelineDefinitions - List pipeline definitions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTektonPipelineDefinitions(listTektonPipelineDefinitionsOptions *ListTektonPipelineDefinitionsOptions)`, func() {

			listTektonPipelineDefinitionsOptions := &cdtektonpipelinev2.ListTektonPipelineDefinitionsOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
			}

			definitions, response, err := cdTektonPipelineService.ListTektonPipelineDefinitions(listTektonPipelineDefinitionsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(definitions).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})

	Describe(`CreateTektonPipelineDefinition - Create a single definition`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptions *CreateTektonPipelineDefinitionOptions)`, func() {

			definitionScmSourceModel := &cdtektonpipelinev2.DefinitionScmSource{
				URL: core.StringPtr("https://github.com/IBM/tekton-tutorial.git"),
				Branch: core.StringPtr("master"),
				Tag: core.StringPtr("testString"),
				Path: core.StringPtr(".tekton"),
			}

			createTektonPipelineDefinitionOptions := &cdtektonpipelinev2.CreateTektonPipelineDefinitionOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				ScmSource: definitionScmSourceModel,
			}

			definition, response, err := cdTektonPipelineService.CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(definition).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 404
			//
		})
	})

	Describe(`GetTektonPipelineDefinition - Retrieve a single definition entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTektonPipelineDefinition(getTektonPipelineDefinitionOptions *GetTektonPipelineDefinitionOptions)`, func() {

			getTektonPipelineDefinitionOptions := &cdtektonpipelinev2.GetTektonPipelineDefinitionOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				DefinitionID: core.StringPtr("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada"),
			}

			definition, response, err := cdTektonPipelineService.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(definition).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})

	Describe(`ReplaceTektonPipelineDefinition - Edit a single definition entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceTektonPipelineDefinition(replaceTektonPipelineDefinitionOptions *ReplaceTektonPipelineDefinitionOptions)`, func() {

			definitionScmSourceModel := &cdtektonpipelinev2.DefinitionScmSource{
				URL: core.StringPtr("https://github.com/IBM/tekton-tutorial.git"),
				Branch: core.StringPtr("master"),
				Tag: core.StringPtr("testString"),
				Path: core.StringPtr(".tekton"),
			}

			replaceTektonPipelineDefinitionOptions := &cdtektonpipelinev2.ReplaceTektonPipelineDefinitionOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				DefinitionID: core.StringPtr("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada"),
				ScmSource: definitionScmSourceModel,
				ServiceInstanceID: core.StringPtr("071d8049-d984-4feb-a2ed-2a1e938918ba"),
				ID: core.StringPtr("22f92ab1-e0ac-4c65-84e7-8a4cb32dba0f"),
			}

			definition, response, err := cdTektonPipelineService.ReplaceTektonPipelineDefinition(replaceTektonPipelineDefinitionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(definition).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 404
			//
		})
	})

	Describe(`ListTektonPipelineProperties - List pipeline environment properties`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTektonPipelineProperties(listTektonPipelinePropertiesOptions *ListTektonPipelinePropertiesOptions)`, func() {

			listTektonPipelinePropertiesOptions := &cdtektonpipelinev2.ListTektonPipelinePropertiesOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				Name: core.StringPtr("prod"),
				Type: []string{"SECURE", "TEXT"},
				Sort: core.StringPtr("name"),
			}

			envProperties, response, err := cdTektonPipelineService.ListTektonPipelineProperties(listTektonPipelinePropertiesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(envProperties).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})

	Describe(`CreateTektonPipelineProperties - Create pipeline environment property`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTektonPipelineProperties(createTektonPipelinePropertiesOptions *CreateTektonPipelinePropertiesOptions)`, func() {

			createTektonPipelinePropertiesOptions := &cdtektonpipelinev2.CreateTektonPipelinePropertiesOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				Name: core.StringPtr("key1"),
				Value: core.StringPtr("https://github.com/IBM/tekton-tutorial.git"),
				Enum: []string{"testString"},
				Default: core.StringPtr("testString"),
				Type: core.StringPtr("TEXT"),
				Path: core.StringPtr("testString"),
			}

			property, response, err := cdTektonPipelineService.CreateTektonPipelineProperties(createTektonPipelinePropertiesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(property).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 404
			//
		})
	})

	Describe(`GetTektonPipelineProperty - Get a single pipeline environment property`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTektonPipelineProperty(getTektonPipelinePropertyOptions *GetTektonPipelinePropertyOptions)`, func() {

			getTektonPipelinePropertyOptions := &cdtektonpipelinev2.GetTektonPipelinePropertyOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				PropertyName: core.StringPtr("debug-pipeline"),
			}

			property, response, err := cdTektonPipelineService.GetTektonPipelineProperty(getTektonPipelinePropertyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(property).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})

	Describe(`ReplaceTektonPipelineProperty - Replace a single pipeline environment property value`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceTektonPipelineProperty(replaceTektonPipelinePropertyOptions *ReplaceTektonPipelinePropertyOptions)`, func() {

			replaceTektonPipelinePropertyOptions := &cdtektonpipelinev2.ReplaceTektonPipelinePropertyOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				PropertyName: core.StringPtr("debug-pipeline"),
				Name: core.StringPtr("key1"),
				Value: core.StringPtr("https://github.com/IBM/tekton-tutorial.git"),
				Enum: []string{"testString"},
				Default: core.StringPtr("testString"),
				Type: core.StringPtr("TEXT"),
				Path: core.StringPtr("testString"),
			}

			property, response, err := cdTektonPipelineService.ReplaceTektonPipelineProperty(replaceTektonPipelinePropertyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(property).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 404
			//
		})
	})

	Describe(`ListTektonPipelineTriggers - List pipeline triggers`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTektonPipelineTriggers(listTektonPipelineTriggersOptions *ListTektonPipelineTriggersOptions)`, func() {

			listTektonPipelineTriggersOptions := &cdtektonpipelinev2.ListTektonPipelineTriggersOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				Type: core.StringPtr("manual,scm"),
				Name: core.StringPtr("testString"),
				EventListener: core.StringPtr("testString"),
				WorkerID: core.StringPtr("testString"),
				WorkerName: core.StringPtr("testString"),
				Disabled: core.StringPtr("true"),
				Tags: core.StringPtr("tag1,tag2"),
			}

			triggers, response, err := cdTektonPipelineService.ListTektonPipelineTriggers(listTektonPipelineTriggersOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(triggers).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})

	Describe(`CreateTektonPipelineTrigger - Create a trigger or duplicate a trigger`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTektonPipelineTrigger(createTektonPipelineTriggerOptions *CreateTektonPipelineTriggerOptions)`, func() {

			triggerModel := &cdtektonpipelinev2.TriggerDuplicateTrigger{
				SourceTriggerID: core.StringPtr("b3a8228f-1c82-409b-b249-7639166a0300"),
				Name: core.StringPtr("Generic Trigger- duplicated"),
			}

			createTektonPipelineTriggerOptions := &cdtektonpipelinev2.CreateTektonPipelineTriggerOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				Trigger: triggerModel,
			}

			trigger, response, err := cdTektonPipelineService.CreateTektonPipelineTrigger(createTektonPipelineTriggerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(trigger).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 404
			//
		})
	})

	Describe(`GetTektonPipelineTrigger - Get a single trigger`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTektonPipelineTrigger(getTektonPipelineTriggerOptions *GetTektonPipelineTriggerOptions)`, func() {

			getTektonPipelineTriggerOptions := &cdtektonpipelinev2.GetTektonPipelineTriggerOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				TriggerID: core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147"),
			}

			trigger, response, err := cdTektonPipelineService.GetTektonPipelineTrigger(getTektonPipelineTriggerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trigger).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})

	Describe(`UpdateTektonPipelineTrigger - Edit a single trigger entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateTektonPipelineTrigger(updateTektonPipelineTriggerOptions *UpdateTektonPipelineTriggerOptions)`, func() {

			workerModel := &cdtektonpipelinev2.Worker{
				Name: core.StringPtr("testString"),
				Type: core.StringPtr("private"),
				ID: core.StringPtr("testString"),
			}

			concurrencyModel := &cdtektonpipelinev2.Concurrency{
				MaxConcurrentRuns: core.Int64Ptr(int64(20)),
			}

			genericSecretModel := &cdtektonpipelinev2.GenericSecret{
				Type: core.StringPtr("tokenMatches"),
				Value: core.StringPtr("testString"),
				Source: core.StringPtr("header"),
				KeyName: core.StringPtr("testString"),
				Algorithm: core.StringPtr("md4"),
			}

			triggerScmSourceModel := &cdtektonpipelinev2.TriggerScmSource{
				URL: core.StringPtr("testString"),
				Branch: core.StringPtr("testString"),
				Pattern: core.StringPtr("testString"),
				BlindConnection: core.BoolPtr(true),
				HookID: core.StringPtr("testString"),
			}

			eventsModel := &cdtektonpipelinev2.Events{
				Push: core.BoolPtr(true),
				PullRequestClosed: core.BoolPtr(true),
				PullRequest: core.BoolPtr(true),
			}

			updateTektonPipelineTriggerOptions := &cdtektonpipelinev2.UpdateTektonPipelineTriggerOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				TriggerID: core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147"),
				Type: core.StringPtr("manual"),
				Name: core.StringPtr("start-deploy"),
				EventListener: core.StringPtr("testString"),
				Tags: []string{"testString"},
				Worker: workerModel,
				Concurrency: concurrencyModel,
				Disabled: core.BoolPtr(true),
				Secret: genericSecretModel,
				Cron: core.StringPtr("testString"),
				Timezone: core.StringPtr("Africa/Abidjan"),
				ScmSource: triggerScmSourceModel,
				Events: eventsModel,
			}

			trigger, response, err := cdTektonPipelineService.UpdateTektonPipelineTrigger(updateTektonPipelineTriggerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trigger).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 404
			//
		})
	})

	Describe(`ListTektonPipelineTriggerProperties - List trigger environment properties`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTektonPipelineTriggerProperties(listTektonPipelineTriggerPropertiesOptions *ListTektonPipelineTriggerPropertiesOptions)`, func() {

			listTektonPipelineTriggerPropertiesOptions := &cdtektonpipelinev2.ListTektonPipelineTriggerPropertiesOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				TriggerID: core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147"),
				Name: core.StringPtr("prod"),
				Type: core.StringPtr("SECURE,TEXT"),
				Sort: core.StringPtr("name"),
			}

			triggerProperties, response, err := cdTektonPipelineService.ListTektonPipelineTriggerProperties(listTektonPipelineTriggerPropertiesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(triggerProperties).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})

	Describe(`CreateTektonPipelineTriggerProperties - Create trigger's environment property`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTektonPipelineTriggerProperties(createTektonPipelineTriggerPropertiesOptions *CreateTektonPipelineTriggerPropertiesOptions)`, func() {

			createTektonPipelineTriggerPropertiesOptions := &cdtektonpipelinev2.CreateTektonPipelineTriggerPropertiesOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				TriggerID: core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147"),
				Name: core.StringPtr("key1"),
				Value: core.StringPtr("https://github.com/IBM/tekton-tutorial.git"),
				Enum: []string{"testString"},
				Default: core.StringPtr("testString"),
				Type: core.StringPtr("TEXT"),
				Path: core.StringPtr("testString"),
			}

			triggerProperty, response, err := cdTektonPipelineService.CreateTektonPipelineTriggerProperties(createTektonPipelineTriggerPropertiesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(triggerProperty).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 404
			//
		})
	})

	Describe(`GetTektonPipelineTriggerProperty - Get a trigger's environment property`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTektonPipelineTriggerProperty(getTektonPipelineTriggerPropertyOptions *GetTektonPipelineTriggerPropertyOptions)`, func() {

			getTektonPipelineTriggerPropertyOptions := &cdtektonpipelinev2.GetTektonPipelineTriggerPropertyOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				TriggerID: core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147"),
				PropertyName: core.StringPtr("debug-pipeline"),
			}

			triggerProperty, response, err := cdTektonPipelineService.GetTektonPipelineTriggerProperty(getTektonPipelineTriggerPropertyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(triggerProperty).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})

	Describe(`ReplaceTektonPipelineTriggerProperty - Replace a trigger's environment property value`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceTektonPipelineTriggerProperty(replaceTektonPipelineTriggerPropertyOptions *ReplaceTektonPipelineTriggerPropertyOptions)`, func() {

			replaceTektonPipelineTriggerPropertyOptions := &cdtektonpipelinev2.ReplaceTektonPipelineTriggerPropertyOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				TriggerID: core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147"),
				PropertyName: core.StringPtr("debug-pipeline"),
				Name: core.StringPtr("key1"),
				Value: core.StringPtr("https://github.com/IBM/tekton-tutorial.git"),
				Enum: []string{"testString"},
				Default: core.StringPtr("testString"),
				Type: core.StringPtr("TEXT"),
				Path: core.StringPtr("testString"),
			}

			triggerProperty, response, err := cdTektonPipelineService.ReplaceTektonPipelineTriggerProperty(replaceTektonPipelineTriggerPropertyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(triggerProperty).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 404
			//
		})
	})

	Describe(`DeleteTektonPipelineTriggerProperty - Delete a trigger's property`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTektonPipelineTriggerProperty(deleteTektonPipelineTriggerPropertyOptions *DeleteTektonPipelineTriggerPropertyOptions)`, func() {

			deleteTektonPipelineTriggerPropertyOptions := &cdtektonpipelinev2.DeleteTektonPipelineTriggerPropertyOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				TriggerID: core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147"),
				PropertyName: core.StringPtr("debug-pipeline"),
			}

			response, err := cdTektonPipelineService.DeleteTektonPipelineTriggerProperty(deleteTektonPipelineTriggerPropertyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})

	Describe(`DeleteTektonPipelineTrigger - Delete a single trigger`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTektonPipelineTrigger(deleteTektonPipelineTriggerOptions *DeleteTektonPipelineTriggerOptions)`, func() {

			deleteTektonPipelineTriggerOptions := &cdtektonpipelinev2.DeleteTektonPipelineTriggerOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				TriggerID: core.StringPtr("1bb892a1-2e04-4768-a369-b1159eace147"),
			}

			response, err := cdTektonPipelineService.DeleteTektonPipelineTrigger(deleteTektonPipelineTriggerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})

	Describe(`DeleteTektonPipelineRun - Delete a pipeline run record`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTektonPipelineRun(deleteTektonPipelineRunOptions *DeleteTektonPipelineRunOptions)`, func() {

			deleteTektonPipelineRunOptions := &cdtektonpipelinev2.DeleteTektonPipelineRunOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				ID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
			}

			response, err := cdTektonPipelineService.DeleteTektonPipelineRun(deleteTektonPipelineRunOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})

	Describe(`DeleteTektonPipelineProperty - Delete a single pipeline environment property`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTektonPipelineProperty(deleteTektonPipelinePropertyOptions *DeleteTektonPipelinePropertyOptions)`, func() {

			deleteTektonPipelinePropertyOptions := &cdtektonpipelinev2.DeleteTektonPipelinePropertyOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				PropertyName: core.StringPtr("debug-pipeline"),
			}

			response, err := cdTektonPipelineService.DeleteTektonPipelineProperty(deleteTektonPipelinePropertyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})

	Describe(`DeleteTektonPipelineDefinition - Delete a single definition entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTektonPipelineDefinition(deleteTektonPipelineDefinitionOptions *DeleteTektonPipelineDefinitionOptions)`, func() {

			deleteTektonPipelineDefinitionOptions := &cdtektonpipelinev2.DeleteTektonPipelineDefinitionOptions{
				PipelineID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
				DefinitionID: core.StringPtr("94299034-d45f-4e9a-8ed5-6bd5c7bb7ada"),
			}

			response, err := cdTektonPipelineService.DeleteTektonPipelineDefinition(deleteTektonPipelineDefinitionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})

	Describe(`DeleteTektonPipeline - Delete tekton pipeline instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTektonPipeline(deleteTektonPipelineOptions *DeleteTektonPipelineOptions)`, func() {

			deleteTektonPipelineOptions := &cdtektonpipelinev2.DeleteTektonPipelineOptions{
				ID: core.StringPtr("94619026-912b-4d92-8f51-6c74f0692d90"),
			}

			response, err := cdTektonPipelineService.DeleteTektonPipeline(deleteTektonPipelineOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			//
		})
	})
})

//
// Utility functions are declared in the unit test file
//
