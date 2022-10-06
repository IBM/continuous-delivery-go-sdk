This document contains details about [changes](CHANGELOG.md) to IBM Cloud Continuous Delivery (CD) APIs, SDKs, and Terraform resources.

# [0.1.2 (2022-09-09)](CHANGELOG.md#012-2022-09-09) Beta

## Schedule

- October 3, 2022 - IBM Cloud Terraform Provider version 1.46.0 is published: https://registry.terraform.io/providers/IBM-Cloud/ibm/1.46.0. The CD Tekton Pipeline and CD Toolchain Beta resources and data sources within this Provider contain a number of breaking changes. These changes will not work with the IBM Cloud Continuous Delivery service APIs until we deliver corresponding changes to the APIs. Until then (see next bullet), you should use IBM Cloud Terraform Provider version 1.45.1 highest.
- October 5, 2022 - Breaking changes to the IBM Cloud Continuous Delivery service’s “v2” Beta HTTP APIs and Go SDKs are deployed to production. These changes are compatible with IBM Cloud Terraform Provider versions 1.46.0-beta0 and 1.46.0. Once these changes are deployed, you should use IBM Cloud Terraform Provider version 1.46.0.

## Breaking changes to APIs and SDKs

General changes
- Toolchain base URLs changed from https://api.{region}.devops.cloud.ibm.com/v2 to https://api.{region}.devops.cloud.ibm.com/toolchain/v2
- Tekton Pipeline base URLs changed from https://api.{region}.devops.cloud.ibm.com/v2 to https://api.{region}.devops.cloud.ibm.com/pipeline/v2

`toolchains` and `toolchains/../tools`
- Removed `parameters_references` in POST/PATCH tool request body. Calling the endpoint with this property will now result in a 400 error.
- Added of cursor-based pagination in GET `toolchains` and GET `toolchains/../tools` collection endpoints, controlled using the `start` and `limit` query parameters. While offset/limit pagination is still supported, cursor-based pagination is now the default behaviour.
- Calling the POST `toolchains/../tools` endpoint with all required broker parameters will now return an error in the event of broker failure during provision or bind operation or any other internal error, as opposed to the previous behaviour which would return a status 200 and a 'misconfigured' tool
- Calling the POST `toolchains/../tools` endpoint without all required broker parameters will now return an error, as opposed to the previous behaviour which would return a status 200 and an 'unconfigured' tool

`tekton_pipelines`
- Field `created` renamed to `created_at`
- Field `html_url` renamed to `runs_url`

`tekton_pipelines/../pipeline_runs`
- Field `html_url` renamed to `run_url` for individual runs

`tekton_pipelines/../definitions`
- Field `service_instance_id` relocated inside `scm_source`

`tekton_pipelines/../triggers`
- `trigger` object removed and replaced by its nested properties
- `concurrency` object replaced by `max_concurrent_runs` integer
- For triggers of type `scm`, field `service_instance_id` relocated inside `scm_source`

`tekton_pipelines/../properties` and `tekton_pipelines/../triggers/../properties`
- `enum` property for `single_select` type reshaped to remove `default` param and use `value` param instead
- Non-snake_case property types renamed to conform to snake_case
  - `SECURE` --> `secure`
  - `TEXT` --> `text`
  - `INTEGRATION` --> `integration`
  - `SINGLE_SELECT` --> `single_select`
  - `APPCONFiG` --> `appconfig`
- Non-snake_case secret types renamed to conform to snake_case, for generic webhook triggers only
  - `tokenMatches` --> `token_matches`
  - `digestMatches` --> `digest_matches`
  - `internalValidation` --> `internal_validation`

## Backward compatible changes to APIs and SDKs

### Feature: Cursor-based pagination for pipeline_runs

Added support for token based pagination when fetching the collection of `pipeline_runs`. In the returned data object, the `next` object contains a link with a generated token, which can be used to fetch the next page of data from the `pipeline_runs` endpoint. Note: offset and limit pagination continue to be supported when provided in the request query parameters.

Example: Response from GET pipeline_runs endpoint

```
{
    "limit": 10,
    "pipeline_runs": [...],
    "first": {
        "href": "https://devops-api.us-south.devops.cloud.ibm.com/pipeline/v2/tekton_pipelines//pipeline_runs?limit=10"
    },
    "next": {
        "href": "https://devops-api.us-south.devops.cloud.ibm.com/pipeline/v2/tekton_pipelines/11112222-3333-4444-5555-666677778888/pipeline_runs?limit=10&start=MTY0MzI4ODYwOTg2Nw==",
        "start": "MTY0MzI4ODYwOTg2Nw=="
    }
}
```

### Feature: New Pipeline creation options

Added support for `enable_slack_notifications` and `enable_partial_cloning` settings in the pipeline creation options. Allows enabling or disabling of these two pipeline configuration options that appear in the "Other Settings" page of the pipeline UI.

Example: Create Tekton Pipeline with optional settings enabled

```
curl -X POST \
--location \
--header "Authorization: Bearer $IAM_TOKEN" \
--header "Accept: application/json" \
--header "Content-Type: application/json" \
--data '{ "id": "11112222-3333-4444-5555-666677778888", "enable_slack_notifications": true, "enable_partial_cloning": true }' \
"https://devops-api.us-south.devops.cloud.ibm.com/pipeline/v2/tekton_pipelines"
```

### Feature: New webhook_url field

Added `webhook_url` in generic webhook trigger response payload

Example: Response from GET triggers endpoint, trigger type "generic"

```
{
    "type": "generic",
    "name": "my-generic-trigger",
    "event_listener": "my-listener",
    "href": "https://devops-api.us-south.devops.dev.cloud.ibm.com/pipeline/v2/tekton_pipelines/002cedad-b5a4-46c0-b1ac-17ecd8288bbd/triggers/80ff3cc7-b1c1-4b35-8494-db0e78d16607"
    "id": "99990000-aaaa-ffff-1111-8888bbbbcccc",
    "secret": {
        "type": "token_matches",
        "value": "[secret]",
        "source": "header",
        "key_name": "header"
    },
    "webhook_url": "https://devops-api.us-south.devops.dev.cloud.ibm.com/v1/tekton-webhook/11112222-3333-4444-5555-666677778888/run/99990000-aaaa-ffff-1111-8888bbbbcccc"
}
```

## Breaking changes to Terraform resources

`ibm_cd_tekton_pipeline` data source output
- `created` renamed `created_at`
- `html_url` renamed `runs_url`

`ibm_cd_tekton_pipeline_definition` data source output
- `service_instance_id` relocated inside `scm_source`

`ibm_cd_tekton_pipeline_trigger` resource
- `trigger` object removed and replaced by its nested properties
- `concurrency` object replaced by `max_concurrent_runs` integer

`ibm_cd_tekton_pipeline_trigger` data source output
- For triggers of type `scm`, `service_instance_id` relocated inside `scm_source`

`ibm_cd_tekton_pipeline_property` and `ibm_cd_tekton_pipeline_trigger_property` resources
- `enum` property for `single_select` type reshaped to remove `default` param and use `value` param instead
- Non-snake_case property types renamed to conform to snake_case
  - `SECURE` --> `secure`
  - `TEXT` --> `text`
  - `INTEGRATION` --> `integration`
  - `SINGLE_SELECT` --> `single_select`
  - `APPCONFiG` --> `appconfig`
- Non-snake_case secret types renamed to conform to snake_case, for generic webhook triggers only
  - `tokenMatches` --> `token_matches`
  - `digestMatches` --> `digest_matches`
  - `internalValidation` --> `internal_validation`

**Note:** In examples that follow, some fields are excluded for brevity.

### Example: `created_at`, `runs_url`

Version 1.45.0
```
data "ibm_cd_tekton_pipeline" "get_pipeline" {
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
output "data_pipeline_output" { value = data.ibm_cd_tekton_pipeline.get_pipeline }
```
Result
```
data_pipeline_output = {
  "created" = "2022-09-22T12:50:53.041Z"
  "id" = "11112222-3333-4444-5555-666677778888"
  "name" = "my-sample-pipeline"
  "pipeline_id" = "11112222-3333-4444-5555-666677778888"
  "html_url" = "https://cloud.ibm.com/devops/pipelines/tekton/11112222-3333-4444-5555-666677778888?env_id=ibm:yp:us-south"
  "updated_at" = "2022-09-22T12:53:03.720Z"  
}
```

Version 1.46.0-beta0, 1.46.0
```
data "ibm_cd_tekton_pipeline" "get_pipeline" {
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
output "data_pipeline_output" { value = data.ibm_cd_tekton_pipeline.get_pipeline }
```
Result
```
data_pipeline_output = {
  "created_at" = "2022-09-22T12:50:53.041Z"
  "id" = "11112222-3333-4444-5555-666677778888"
  "name" = "my-sample-pipeline"
  "pipeline_id" = "11112222-3333-4444-5555-666677778888"
  "runs_url" = "https://cloud.ibm.com/devops/pipelines/tekton/11112222-3333-4444-5555-666677778888?env_id=ibm:yp:us-south"
  "updated_at" = "2022-09-22T12:53:03.720Z"  
}
```

### Example: `service_instance_id` (pipeline definition)

Version 1.45.0
```
data "ibm_cd_tekton_pipeline_definition" "get_definition" {
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
  definition_id = ibm_cd_tekton_pipeline_definition.my_definition.definition_id
}
output "data_definition_output" { value=data.ibm_cd_tekton_pipeline_definition.get_definition }
```
Result
```
data_definition_output = {
  "definition_id" = "11112222-3333-4444-5555-666677778888"
  "id" = "aaaabbbb-cccc-dddd-eeee-ffff00001111/11112222-3333-4444-5555-666677778888"
  "pipeline_id" = "aaaabbbb-cccc-dddd-eeee-ffff00001111"
  "service_instance_id" = "99990000-aaaa-ffff-1111-8888bbbbcccc"
  "scm_source" = tolist([
    {
      "branch" = "master"
      "path" = ".tekton"
      "tag" = ""
      "url" = "https://github.com/IBM/tekton-tutorial.git"
    },
  ])
}
```

Version 1.46.0-beta0, 1.46.0
```
data "ibm_cd_tekton_pipeline_definition" "get_definition" {
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
  definition_id = ibm_cd_tekton_pipeline_definition.my_definition.definition_id
}
output "data_definition_output" { value=data.ibm_cd_tekton_pipeline_definition.get_definition }
```
Result
```
data_definition_output = {
  "definition_id" = "11112222-3333-4444-5555-666677778888"
  "id" = "aaaabbbb-cccc-dddd-eeee-ffff00001111/11112222-3333-4444-5555-666677778888"
  "pipeline_id" = "aaaabbbb-cccc-dddd-eeee-ffff00001111"
  "scm_source" = tolist([
    {
      "branch" = "master"
      "path" = ".tekton"
      "service_instance_id" = "99990000-aaaa-ffff-1111-8888bbbbcccc"
      "tag" = ""
      "url" = "https://github.com/IBM/tekton-tutorial.git"
    },
  ])
}
```

### Example: `trigger`

Version 1.45.0
```
resource "ibm_cd_tekton_pipeline_trigger" "manual_trigger" {
  trigger {
    type = "manual"
    name = "manual-trigger-1"
    event_listener  = "my-listener"
    concurrency {
      max_concurrent_runs = 2
    }
  }
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
```

Version 1.46.0-beta0, 1.46.0
```
resource "ibm_cd_tekton_pipeline_trigger" "manual_trigger" {
  type = "manual"
  name = "manual-trigger-1"
  event_listener  = "my-listener"
  max_concurrent_runs = 2
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
```

### Example: `service_instance_id` (pipeline trigger)

Version 1.45.0
```
data "ibm_cd_tekton_pipeline_trigger" "get_scm_trigger" {
 pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
 trigger_id = ibm_cd_tekton_pipeline_trigger.my_scm_trigger.trigger_id
}
output "data_output_scm_trigger" { value=data.ibm_cd_tekton_pipeline_trigger.get_scm_trigger }
```
Result
```
data_output_scm_trigger = {
  "event_listener" = "my-listener"
  "id" = "aaaabbbb-cccc-dddd-eeee-ffff00001111/11112222-3333-4444-5555-666677778888"
  "name" = "my-scm-trigger"
  "pipeline_id" = "aaaabbbb-cccc-dddd-eeee-ffff00001111"
  "service_instance_id" = "99990000-aaaa-ffff-1111-8888bbbbcccc"
  "scm_source" = tolist([
    {
      "blind_connection" = false
      "branch" = "master"
      "hook_id" = "380736321"
      "pattern" = ""
      "url" = "https://github.com/IBM/tekton-tutorial.git"
    },
  ])
  "trigger_id" = "11112222-3333-4444-5555-666677778888"
  "type" = "scm"
}
```

Version 1.46.0-beta0, 1.46.0
```
data "ibm_cd_tekton_pipeline_trigger" "get_scm_trigger" {
 pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
 trigger_id = ibm_cd_tekton_pipeline_trigger.my_scm_trigger.trigger_id
}
output "data_output_scm_trigger" { value=data.ibm_cd_tekton_pipeline_trigger.get_scm_trigger }
```
Result
```
data_output_scm_trigger = {
  "event_listener" = "my-listener"
  "id" = "aaaabbbb-cccc-dddd-eeee-ffff00001111/11112222-3333-4444-5555-666677778888"
  "name" = "my-scm-trigger"
  "pipeline_id" = "aaaabbbb-cccc-dddd-eeee-ffff00001111"
  "scm_source" = tolist([
    {
      "blind_connection" = false
      "branch" = "master"
      "hook_id" = "380736321"
      "pattern" = ""
      "service_instance_id" = "99990000-aaaa-ffff-1111-8888bbbbcccc"
      "url" = "https://github.com/IBM/tekton-tutorial.git"
    },
  ])
  "trigger_id" = "11112222-3333-4444-5555-666677778888"
  "type" = "scm"
}
```

### Example: type `text`

Version 1.45.0
```
resource "ibm_cd_tekton_pipeline_property" "my-text-prop" {
  name = "evidence-repo"
  type = "TEXT"
  value = "sample-text-value"
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
```

Version 1.46.0-beta0, 1.46.0
```
resource "ibm_cd_tekton_pipeline_property" "my-text-prop" {
  name = "evidence-repo"
  type = "text"
  value = "sample-text-value"
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
```

### Example: type `secure`

Version 1.45.0
```
resource "ibm_cd_tekton_pipeline_property" "my-secure-prop" {
  name = "my-api-key"
  type = "SECURE"
  value = "my-secret-value"
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
```

Version 1.46.0-beta0, 1.46.0
```
resource "ibm_cd_tekton_pipeline_property" "my-secure-prop" {
  name = "my-api-key"
  type = "secure"
  value = "my-secret-value"
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id      
}
```

### Example: type `single_select`

Version 1.45.0
```
resource "ibm_cd_tekton_pipeline_property" "regions" {
  name = "regions"
  type = "SINGLE_SELECT"
  enum = ["us-south", "eu-gb", "eu-de"]
  default = "eu-de"
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
```

Version 1.46.0-beta0, 1.46.0
```
resource "ibm_cd_tekton_pipeline_property" "regions" {
  name = "regions"
  type = "single_select"
  enum = ["us-south", "eu-gb", "eu-de"]
  value = "eu-de"
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
```

### Example: type `integration`
Version 1.45.0
```
resource "ibm_cd_tekton_pipeline_property" "integration_property" {
  name = "my_integration_prop"
  value = ibm_cd_toolchain_tool_githubconsolidated.repo1.tool_id
  type = "INTEGRATION"
  path = "parameters.repo_url"
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
```

Version 1.46.0-beta0, 1.46.0
```
resource "ibm_cd_tekton_pipeline_property" "integration_property" {
  name = "my_integration_prop"
  value = ibm_cd_toolchain_tool_githubconsolidated.repo1.tool_id
  type = "integration"
  path = "parameters.repo_url"
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
```

### Example: type `appconfig`
Version 1.45.0
```
resource "ibm_cd_tekton_pipeline_property" "appconfig_property" {
  name = "myAppConfig"
  value = "{appconfig::11112222-3333-4444-5555-666677778888.prop.myproperty}"
  type = "APPCONFIG"
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
```

Version 1.46.0-beta0, 1.46.0
```
resource "ibm_cd_tekton_pipeline_property" "appconfig_property" {
  name = "myAppConfig"
  value = "{appconfig::11112222-3333-4444-5555-666677778888.prop.myproperty}"
  type = "appconfig"
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
```

### Example: type `digest_matches`

Version 1.45.0
```
resource "ibm_cd_tekton_pipeline_trigger" "generic_trigger_digest" {
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
  type = "generic"
  name = "trigger_digest"
  event_listener = "my-listener"
  secret {
    type = "digestMatches"
    value = "123"
    source = "header"
    key_name = "header"
    algorithm = "sha512"
  }
}
```

Version 1.46.0-beta0, 1.46.0
```
resource "ibm_cd_tekton_pipeline_trigger" "generic_trigger_digest" {
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
  type = "generic"
  name = "trigger_digest"
  event_listener = "my-listener"
  secret {
    type = "digest_matches"
    value = "123"
    source = "header"
    key_name = "header"
    algorithm = "sha512"
  }
}
```

### Example: type `internal_validation`

Version 1.45.0
```
resource "ibm_cd_tekton_pipeline_trigger" "generic_trigger_internal" {
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
  type = "generic"
  name = "trigger_internal"
  event_listener = "my-listener"
  secret {
    type = "internalValidation"
  }
}
```

Version 1.46.0-beta0, 1.46.0
```
resource "ibm_cd_tekton_pipeline_trigger" "generic_trigger_internal" {
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
  type = "generic"
  name = "trigger_internal"
  event_listener = "my-listener"
  secret {
    type = "internal_validation"
  }
}
```

### Example: type `token_matches`

Version 1.45.0
```
resource "ibm_cd_tekton_pipeline_trigger" "generic_trigger_token" {
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
  type = "generic"
  name = "trigger_token"
  event_listener = "my-listener"
  secret {
    type = "tokenMatches"
    value = "123"
    source = "header"
    key_name = "header"
  }
}
```

Version 1.46.0-beta0, 1.46.0
```
resource "ibm_cd_tekton_pipeline_trigger" "generic_trigger_token" {
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
  type = "generic"
  name = "trigger_token"
  event_listener = "my-listener"
  secret {
    type = "token_matches"
    value = "123"
    source = "header"
    key_name = "header"
  }
}
```

## Backward compatible changes to Terraform resources

### Feature: New Pipeline creation options

Added support for `enable_slack_notifications` and `enable_partial_cloning` settings in the `ibm_cd_tekton_pipeline` resource. Allows enabling or disabling of these two pipeline configuration options that appear in the "Other Settings" page of the pipeline UI.

Example

```
resource "ibm_cd_tekton_pipeline" "my_pipeline" {
  enable_slack_notifications = true
  enable_partial_cloning = true
  pipeline_id = ibm_cd_toolchain_tool_pipeline.pipeline_tool.tool_id
  worker {
    id = "public"
  }
}
```

### Feature: New webhook_url field

Added `webhook_url` in generic webhook trigger response payload

Example

```
data_trigger_generic_output = {
  "event_listener" = "my-listener"
  "id" = "11112222-3333-4444-5555-666677778888/99990000-aaaa-ffff-1111-8888bbbbcccc"
  "name" = "my-generic-trigger"
  "pipeline_id" = "11112222-3333-4444-5555-666677778888"
  "trigger_id" = "99990000-aaaa-ffff-1111-8888bbbbcccc"
  "type" = "generic"
  "webhook_url" = "https://devops-api.us-south.devops.dev.cloud.ibm.com/v1/tekton-webhook/11112222-3333-4444-5555-666677778888/run/99990000-aaaa-ffff-1111-8888bbbbcccc"
}
```
