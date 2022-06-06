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

package cdtektonpipelinev2_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/org-ids/tekton-pipeline-go-sdk/cdtektonpipelinev2"
)

//
// This file provides an example of how to use the CD Tekton Pipeline service.
//
// The following configuration properties are assumed to be defined:
// CD_TEKTON_PIPELINE_URL=<service base url>
// CD_TEKTON_PIPELINE_AUTH_TYPE=iam
// CD_TEKTON_PIPELINE_APIKEY=<IAM apikey>
// CD_TEKTON_PIPELINE_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`CdTektonPipelineV2 Examples Tests`, func() {

	const externalConfigFile = "../cd_tekton_pipeline_v2.env"

	var (
		cdTektonPipelineService *cdtektonpipelinev2.CdTektonPipelineV2
		config       map[string]string
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
			config, err = core.GetServiceProperties(cdtektonpipelinev2.DefaultServiceName)
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

			cdTektonPipelineServiceOptions := &cdtektonpipelinev2.CdTektonPipelineV2Options{}

			cdTektonPipelineService, err = cdtektonpipelinev2.NewCdTektonPipelineV2UsingExternalConfig(cdTektonPipelineServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(cdTektonPipelineService).ToNot(BeNil())
		})
	})

	Describe(`CdTektonPipelineV2 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTektonPipeline request example`, func() {
			fmt.Println("\nCreateTektonPipeline() result:")
			// begin-create_tekton_pipeline

			workerWithIDModel := &cdtektonpipelinev2.WorkerWithID{
				ID: core.StringPtr("public"),
			}

			createTektonPipelineOptions := cdTektonPipelineService.NewCreateTektonPipelineOptions()
			createTektonPipelineOptions.SetID("94619026-912b-4d92-8f51-6c74f0692d90")
			createTektonPipelineOptions.SetWorker(workerWithIDModel)

			tektonPipeline, response, err := cdTektonPipelineService.CreateTektonPipeline(createTektonPipelineOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tektonPipeline, "", "  ")
			fmt.Println(string(b))

			// end-create_tekton_pipeline

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(tektonPipeline).ToNot(BeNil())

		})
		It(`GetTektonPipeline request example`, func() {
			fmt.Println("\nGetTektonPipeline() result:")
			// begin-get_tekton_pipeline

			getTektonPipelineOptions := cdTektonPipelineService.NewGetTektonPipelineOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
			)

			tektonPipeline, response, err := cdTektonPipelineService.GetTektonPipeline(getTektonPipelineOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tektonPipeline, "", "  ")
			fmt.Println(string(b))

			// end-get_tekton_pipeline

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tektonPipeline).ToNot(BeNil())

		})
		It(`UpdateTektonPipeline request example`, func() {
			fmt.Println("\nUpdateTektonPipeline() result:")
			// begin-update_tekton_pipeline

			workerWithIDModel := &cdtektonpipelinev2.WorkerWithID{
				ID: core.StringPtr("public"),
			}

			updateTektonPipelineOptions := cdTektonPipelineService.NewUpdateTektonPipelineOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
			)
			updateTektonPipelineOptions.SetWorker(workerWithIDModel)

			tektonPipeline, response, err := cdTektonPipelineService.UpdateTektonPipeline(updateTektonPipelineOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tektonPipeline, "", "  ")
			fmt.Println(string(b))

			// end-update_tekton_pipeline

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tektonPipeline).ToNot(BeNil())

		})
		It(`ListTektonPipelineRuns request example`, func() {
			fmt.Println("\nListTektonPipelineRuns() result:")
			// begin-list_tekton_pipeline_runs

			listTektonPipelineRunsOptions := cdTektonPipelineService.NewListTektonPipelineRunsOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
			)
			listTektonPipelineRunsOptions.SetStatus("succeeded")
			listTektonPipelineRunsOptions.SetTriggerName("manual-trigger")

			pipelineRuns, response, err := cdTektonPipelineService.ListTektonPipelineRuns(listTektonPipelineRunsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(pipelineRuns, "", "  ")
			fmt.Println(string(b))

			// end-list_tekton_pipeline_runs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pipelineRuns).ToNot(BeNil())

		})
		It(`CreateTektonPipelineRun request example`, func() {
			fmt.Println("\nCreateTektonPipelineRun() result:")
			// begin-create_tekton_pipeline_run

			createTektonPipelineRunOptions := cdTektonPipelineService.NewCreateTektonPipelineRunOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
			)
			createTektonPipelineRunOptions.SetTriggerName("Generic Webhook Trigger - 0")

			pipelineRun, response, err := cdTektonPipelineService.CreateTektonPipelineRun(createTektonPipelineRunOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(pipelineRun, "", "  ")
			fmt.Println(string(b))

			// end-create_tekton_pipeline_run

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(pipelineRun).ToNot(BeNil())

		})
		It(`GetTektonPipelineRun request example`, func() {
			fmt.Println("\nGetTektonPipelineRun() result:")
			// begin-get_tekton_pipeline_run

			getTektonPipelineRunOptions := cdTektonPipelineService.NewGetTektonPipelineRunOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"94619026-912b-4d92-8f51-6c74f0692d90",
			)
			getTektonPipelineRunOptions.SetIncludes("definitions")

			pipelineRun, response, err := cdTektonPipelineService.GetTektonPipelineRun(getTektonPipelineRunOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(pipelineRun, "", "  ")
			fmt.Println(string(b))

			// end-get_tekton_pipeline_run

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pipelineRun).ToNot(BeNil())

		})
		It(`CancelTektonPipelineRun request example`, func() {
			fmt.Println("\nCancelTektonPipelineRun() result:")
			// begin-cancel_tekton_pipeline_run

			cancelTektonPipelineRunOptions := cdTektonPipelineService.NewCancelTektonPipelineRunOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"94619026-912b-4d92-8f51-6c74f0692d90",
			)
			cancelTektonPipelineRunOptions.SetForce(true)

			pipelineRun, response, err := cdTektonPipelineService.CancelTektonPipelineRun(cancelTektonPipelineRunOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(pipelineRun, "", "  ")
			fmt.Println(string(b))

			// end-cancel_tekton_pipeline_run

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pipelineRun).ToNot(BeNil())

		})
		It(`RerunTektonPipelineRun request example`, func() {
			fmt.Println("\nRerunTektonPipelineRun() result:")
			// begin-rerun_tekton_pipeline_run

			rerunTektonPipelineRunOptions := cdTektonPipelineService.NewRerunTektonPipelineRunOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"94619026-912b-4d92-8f51-6c74f0692d90",
			)

			pipelineRun, response, err := cdTektonPipelineService.RerunTektonPipelineRun(rerunTektonPipelineRunOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(pipelineRun, "", "  ")
			fmt.Println(string(b))

			// end-rerun_tekton_pipeline_run

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(pipelineRun).ToNot(BeNil())

		})
		It(`GetTektonPipelineRunLogs request example`, func() {
			fmt.Println("\nGetTektonPipelineRunLogs() result:")
			// begin-get_tekton_pipeline_run_logs

			getTektonPipelineRunLogsOptions := cdTektonPipelineService.NewGetTektonPipelineRunLogsOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"94619026-912b-4d92-8f51-6c74f0692d90",
			)

			pipelineRunLogs, response, err := cdTektonPipelineService.GetTektonPipelineRunLogs(getTektonPipelineRunLogsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(pipelineRunLogs, "", "  ")
			fmt.Println(string(b))

			// end-get_tekton_pipeline_run_logs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pipelineRunLogs).ToNot(BeNil())

		})
		It(`GetTektonPipelineRunLogContent request example`, func() {
			fmt.Println("\nGetTektonPipelineRunLogContent() result:")
			// begin-get_tekton_pipeline_run_log_content

			getTektonPipelineRunLogContentOptions := cdTektonPipelineService.NewGetTektonPipelineRunLogContentOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"bf4b3abd-0c93-416b-911e-9cf42f1a1085",
				"94619026-912b-4d92-8f51-6c74f0692d90",
			)

			stepLog, response, err := cdTektonPipelineService.GetTektonPipelineRunLogContent(getTektonPipelineRunLogContentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(stepLog, "", "  ")
			fmt.Println(string(b))

			// end-get_tekton_pipeline_run_log_content

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(stepLog).ToNot(BeNil())

		})
		It(`ListTektonPipelineDefinitions request example`, func() {
			fmt.Println("\nListTektonPipelineDefinitions() result:")
			// begin-list_tekton_pipeline_definitions

			listTektonPipelineDefinitionsOptions := cdTektonPipelineService.NewListTektonPipelineDefinitionsOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
			)

			definitions, response, err := cdTektonPipelineService.ListTektonPipelineDefinitions(listTektonPipelineDefinitionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(definitions, "", "  ")
			fmt.Println(string(b))

			// end-list_tekton_pipeline_definitions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(definitions).ToNot(BeNil())

		})
		It(`CreateTektonPipelineDefinition request example`, func() {
			fmt.Println("\nCreateTektonPipelineDefinition() result:")
			// begin-create_tekton_pipeline_definition

			definitionScmSourceModel := &cdtektonpipelinev2.DefinitionScmSource{
				URL: core.StringPtr("https://github.com/IBM/tekton-tutorial.git"),
				Branch: core.StringPtr("master"),
				Path: core.StringPtr(".tekton"),
			}

			createTektonPipelineDefinitionOptions := cdTektonPipelineService.NewCreateTektonPipelineDefinitionOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
			)
			createTektonPipelineDefinitionOptions.SetScmSource(definitionScmSourceModel)

			definition, response, err := cdTektonPipelineService.CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(definition, "", "  ")
			fmt.Println(string(b))

			// end-create_tekton_pipeline_definition

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(definition).ToNot(BeNil())

		})
		It(`GetTektonPipelineDefinition request example`, func() {
			fmt.Println("\nGetTektonPipelineDefinition() result:")
			// begin-get_tekton_pipeline_definition

			getTektonPipelineDefinitionOptions := cdTektonPipelineService.NewGetTektonPipelineDefinitionOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"94299034-d45f-4e9a-8ed5-6bd5c7bb7ada",
			)

			definition, response, err := cdTektonPipelineService.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(definition, "", "  ")
			fmt.Println(string(b))

			// end-get_tekton_pipeline_definition

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(definition).ToNot(BeNil())

		})
		It(`ReplaceTektonPipelineDefinition request example`, func() {
			fmt.Println("\nReplaceTektonPipelineDefinition() result:")
			// begin-replace_tekton_pipeline_definition

			definitionScmSourceModel := &cdtektonpipelinev2.DefinitionScmSource{
				URL: core.StringPtr("https://github.com/IBM/tekton-tutorial.git"),
				Branch: core.StringPtr("master"),
				Path: core.StringPtr(".tekton"),
			}

			replaceTektonPipelineDefinitionOptions := cdTektonPipelineService.NewReplaceTektonPipelineDefinitionOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"94299034-d45f-4e9a-8ed5-6bd5c7bb7ada",
			)
			replaceTektonPipelineDefinitionOptions.SetScmSource(definitionScmSourceModel)
			replaceTektonPipelineDefinitionOptions.SetServiceInstanceID("071d8049-d984-4feb-a2ed-2a1e938918ba")
			replaceTektonPipelineDefinitionOptions.SetID("22f92ab1-e0ac-4c65-84e7-8a4cb32dba0f")

			definition, response, err := cdTektonPipelineService.ReplaceTektonPipelineDefinition(replaceTektonPipelineDefinitionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(definition, "", "  ")
			fmt.Println(string(b))

			// end-replace_tekton_pipeline_definition

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(definition).ToNot(BeNil())

		})
		It(`ListTektonPipelineProperties request example`, func() {
			fmt.Println("\nListTektonPipelineProperties() result:")
			// begin-list_tekton_pipeline_properties

			listTektonPipelinePropertiesOptions := cdTektonPipelineService.NewListTektonPipelinePropertiesOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
			)
			listTektonPipelinePropertiesOptions.SetName("prod")
			listTektonPipelinePropertiesOptions.SetType([]string{"SECURE", "TEXT"})
			listTektonPipelinePropertiesOptions.SetSort("name")

			envProperties, response, err := cdTektonPipelineService.ListTektonPipelineProperties(listTektonPipelinePropertiesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(envProperties, "", "  ")
			fmt.Println(string(b))

			// end-list_tekton_pipeline_properties

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(envProperties).ToNot(BeNil())

		})
		It(`CreateTektonPipelineProperties request example`, func() {
			fmt.Println("\nCreateTektonPipelineProperties() result:")
			// begin-create_tekton_pipeline_properties

			createTektonPipelinePropertiesOptions := cdTektonPipelineService.NewCreateTektonPipelinePropertiesOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
			)
			createTektonPipelinePropertiesOptions.SetName("key1")
			createTektonPipelinePropertiesOptions.SetValue("https://github.com/IBM/tekton-tutorial.git")
			createTektonPipelinePropertiesOptions.SetType("TEXT")

			property, response, err := cdTektonPipelineService.CreateTektonPipelineProperties(createTektonPipelinePropertiesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(property, "", "  ")
			fmt.Println(string(b))

			// end-create_tekton_pipeline_properties

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(property).ToNot(BeNil())

		})
		It(`GetTektonPipelineProperty request example`, func() {
			fmt.Println("\nGetTektonPipelineProperty() result:")
			// begin-get_tekton_pipeline_property

			getTektonPipelinePropertyOptions := cdTektonPipelineService.NewGetTektonPipelinePropertyOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"debug-pipeline",
			)

			property, response, err := cdTektonPipelineService.GetTektonPipelineProperty(getTektonPipelinePropertyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(property, "", "  ")
			fmt.Println(string(b))

			// end-get_tekton_pipeline_property

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(property).ToNot(BeNil())

		})
		It(`ReplaceTektonPipelineProperty request example`, func() {
			fmt.Println("\nReplaceTektonPipelineProperty() result:")
			// begin-replace_tekton_pipeline_property

			replaceTektonPipelinePropertyOptions := cdTektonPipelineService.NewReplaceTektonPipelinePropertyOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"debug-pipeline",
			)
			replaceTektonPipelinePropertyOptions.SetName("key1")
			replaceTektonPipelinePropertyOptions.SetValue("https://github.com/IBM/tekton-tutorial.git")
			replaceTektonPipelinePropertyOptions.SetType("TEXT")

			property, response, err := cdTektonPipelineService.ReplaceTektonPipelineProperty(replaceTektonPipelinePropertyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(property, "", "  ")
			fmt.Println(string(b))

			// end-replace_tekton_pipeline_property

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(property).ToNot(BeNil())

		})
		It(`ListTektonPipelineTriggers request example`, func() {
			fmt.Println("\nListTektonPipelineTriggers() result:")
			// begin-list_tekton_pipeline_triggers

			listTektonPipelineTriggersOptions := cdTektonPipelineService.NewListTektonPipelineTriggersOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
			)
			listTektonPipelineTriggersOptions.SetType("manual,scm")
			listTektonPipelineTriggersOptions.SetDisabled("true")
			listTektonPipelineTriggersOptions.SetTags("tag1,tag2")

			triggers, response, err := cdTektonPipelineService.ListTektonPipelineTriggers(listTektonPipelineTriggersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(triggers, "", "  ")
			fmt.Println(string(b))

			// end-list_tekton_pipeline_triggers

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(triggers).ToNot(BeNil())

		})
		It(`CreateTektonPipelineTrigger request example`, func() {
			fmt.Println("\nCreateTektonPipelineTrigger() result:")
			// begin-create_tekton_pipeline_trigger

			triggerModel := &cdtektonpipelinev2.TriggerDuplicateTrigger{
				SourceTriggerID: core.StringPtr("b3a8228f-1c82-409b-b249-7639166a0300"),
				Name: core.StringPtr("Generic Trigger- duplicated"),
			}

			createTektonPipelineTriggerOptions := cdTektonPipelineService.NewCreateTektonPipelineTriggerOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
			)
			createTektonPipelineTriggerOptions.SetTrigger(triggerModel)

			trigger, response, err := cdTektonPipelineService.CreateTektonPipelineTrigger(createTektonPipelineTriggerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trigger, "", "  ")
			fmt.Println(string(b))

			// end-create_tekton_pipeline_trigger

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(trigger).ToNot(BeNil())

		})
		It(`GetTektonPipelineTrigger request example`, func() {
			fmt.Println("\nGetTektonPipelineTrigger() result:")
			// begin-get_tekton_pipeline_trigger

			getTektonPipelineTriggerOptions := cdTektonPipelineService.NewGetTektonPipelineTriggerOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"1bb892a1-2e04-4768-a369-b1159eace147",
			)

			trigger, response, err := cdTektonPipelineService.GetTektonPipelineTrigger(getTektonPipelineTriggerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trigger, "", "  ")
			fmt.Println(string(b))

			// end-get_tekton_pipeline_trigger

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trigger).ToNot(BeNil())

		})
		It(`UpdateTektonPipelineTrigger request example`, func() {
			fmt.Println("\nUpdateTektonPipelineTrigger() result:")
			// begin-update_tekton_pipeline_trigger

			updateTektonPipelineTriggerOptions := cdTektonPipelineService.NewUpdateTektonPipelineTriggerOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"1bb892a1-2e04-4768-a369-b1159eace147",
			)
			updateTektonPipelineTriggerOptions.SetName("start-deploy")
			updateTektonPipelineTriggerOptions.SetDisabled(true)

			trigger, response, err := cdTektonPipelineService.UpdateTektonPipelineTrigger(updateTektonPipelineTriggerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trigger, "", "  ")
			fmt.Println(string(b))

			// end-update_tekton_pipeline_trigger

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trigger).ToNot(BeNil())

		})
		It(`ListTektonPipelineTriggerProperties request example`, func() {
			fmt.Println("\nListTektonPipelineTriggerProperties() result:")
			// begin-list_tekton_pipeline_trigger_properties

			listTektonPipelineTriggerPropertiesOptions := cdTektonPipelineService.NewListTektonPipelineTriggerPropertiesOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"1bb892a1-2e04-4768-a369-b1159eace147",
				"prod",
				"SECURE,TEXT",
				"name",
			)

			triggerProperties, response, err := cdTektonPipelineService.ListTektonPipelineTriggerProperties(listTektonPipelineTriggerPropertiesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(triggerProperties, "", "  ")
			fmt.Println(string(b))

			// end-list_tekton_pipeline_trigger_properties

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(triggerProperties).ToNot(BeNil())

		})
		It(`CreateTektonPipelineTriggerProperties request example`, func() {
			fmt.Println("\nCreateTektonPipelineTriggerProperties() result:")
			// begin-create_tekton_pipeline_trigger_properties

			createTektonPipelineTriggerPropertiesOptions := cdTektonPipelineService.NewCreateTektonPipelineTriggerPropertiesOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"1bb892a1-2e04-4768-a369-b1159eace147",
			)
			createTektonPipelineTriggerPropertiesOptions.SetName("key1")
			createTektonPipelineTriggerPropertiesOptions.SetValue("https://github.com/IBM/tekton-tutorial.git")
			createTektonPipelineTriggerPropertiesOptions.SetType("TEXT")

			triggerProperty, response, err := cdTektonPipelineService.CreateTektonPipelineTriggerProperties(createTektonPipelineTriggerPropertiesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(triggerProperty, "", "  ")
			fmt.Println(string(b))

			// end-create_tekton_pipeline_trigger_properties

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(triggerProperty).ToNot(BeNil())

		})
		It(`GetTektonPipelineTriggerProperty request example`, func() {
			fmt.Println("\nGetTektonPipelineTriggerProperty() result:")
			// begin-get_tekton_pipeline_trigger_property

			getTektonPipelineTriggerPropertyOptions := cdTektonPipelineService.NewGetTektonPipelineTriggerPropertyOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"1bb892a1-2e04-4768-a369-b1159eace147",
				"debug-pipeline",
			)

			triggerProperty, response, err := cdTektonPipelineService.GetTektonPipelineTriggerProperty(getTektonPipelineTriggerPropertyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(triggerProperty, "", "  ")
			fmt.Println(string(b))

			// end-get_tekton_pipeline_trigger_property

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(triggerProperty).ToNot(BeNil())

		})
		It(`ReplaceTektonPipelineTriggerProperty request example`, func() {
			fmt.Println("\nReplaceTektonPipelineTriggerProperty() result:")
			// begin-replace_tekton_pipeline_trigger_property

			replaceTektonPipelineTriggerPropertyOptions := cdTektonPipelineService.NewReplaceTektonPipelineTriggerPropertyOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"1bb892a1-2e04-4768-a369-b1159eace147",
				"debug-pipeline",
			)
			replaceTektonPipelineTriggerPropertyOptions.SetName("key1")
			replaceTektonPipelineTriggerPropertyOptions.SetValue("https://github.com/IBM/tekton-tutorial.git")
			replaceTektonPipelineTriggerPropertyOptions.SetType("TEXT")

			triggerProperty, response, err := cdTektonPipelineService.ReplaceTektonPipelineTriggerProperty(replaceTektonPipelineTriggerPropertyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(triggerProperty, "", "  ")
			fmt.Println(string(b))

			// end-replace_tekton_pipeline_trigger_property

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(triggerProperty).ToNot(BeNil())

		})
		It(`DeleteTektonPipelineTriggerProperty request example`, func() {
			// begin-delete_tekton_pipeline_trigger_property

			deleteTektonPipelineTriggerPropertyOptions := cdTektonPipelineService.NewDeleteTektonPipelineTriggerPropertyOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"1bb892a1-2e04-4768-a369-b1159eace147",
				"debug-pipeline",
			)

			response, err := cdTektonPipelineService.DeleteTektonPipelineTriggerProperty(deleteTektonPipelineTriggerPropertyOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteTektonPipelineTriggerProperty(): %d\n", response.StatusCode)
			}

			// end-delete_tekton_pipeline_trigger_property

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteTektonPipelineTrigger request example`, func() {
			// begin-delete_tekton_pipeline_trigger

			deleteTektonPipelineTriggerOptions := cdTektonPipelineService.NewDeleteTektonPipelineTriggerOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"1bb892a1-2e04-4768-a369-b1159eace147",
			)

			response, err := cdTektonPipelineService.DeleteTektonPipelineTrigger(deleteTektonPipelineTriggerOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteTektonPipelineTrigger(): %d\n", response.StatusCode)
			}

			// end-delete_tekton_pipeline_trigger

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteTektonPipelineRun request example`, func() {
			// begin-delete_tekton_pipeline_run

			deleteTektonPipelineRunOptions := cdTektonPipelineService.NewDeleteTektonPipelineRunOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"94619026-912b-4d92-8f51-6c74f0692d90",
			)

			response, err := cdTektonPipelineService.DeleteTektonPipelineRun(deleteTektonPipelineRunOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteTektonPipelineRun(): %d\n", response.StatusCode)
			}

			// end-delete_tekton_pipeline_run

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteTektonPipelineProperty request example`, func() {
			// begin-delete_tekton_pipeline_property

			deleteTektonPipelinePropertyOptions := cdTektonPipelineService.NewDeleteTektonPipelinePropertyOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"debug-pipeline",
			)

			response, err := cdTektonPipelineService.DeleteTektonPipelineProperty(deleteTektonPipelinePropertyOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteTektonPipelineProperty(): %d\n", response.StatusCode)
			}

			// end-delete_tekton_pipeline_property

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteTektonPipelineDefinition request example`, func() {
			// begin-delete_tekton_pipeline_definition

			deleteTektonPipelineDefinitionOptions := cdTektonPipelineService.NewDeleteTektonPipelineDefinitionOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
				"94299034-d45f-4e9a-8ed5-6bd5c7bb7ada",
			)

			response, err := cdTektonPipelineService.DeleteTektonPipelineDefinition(deleteTektonPipelineDefinitionOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteTektonPipelineDefinition(): %d\n", response.StatusCode)
			}

			// end-delete_tekton_pipeline_definition

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteTektonPipeline request example`, func() {
			// begin-delete_tekton_pipeline

			deleteTektonPipelineOptions := cdTektonPipelineService.NewDeleteTektonPipelineOptions(
				"94619026-912b-4d92-8f51-6c74f0692d90",
			)

			response, err := cdTektonPipelineService.DeleteTektonPipeline(deleteTektonPipelineOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteTektonPipeline(): %d\n", response.StatusCode)
			}

			// end-delete_tekton_pipeline

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
