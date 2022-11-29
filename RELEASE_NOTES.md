This document contains details about [changes](CHANGELOG.md) to IBM Cloud Continuous Delivery (CD) APIs, SDKs, and Terraform resources.

# Pending release

This section outlines changes to the CD APIs, SDKs, and Terraform resources accrued since release [1.0.1](#101-2022-11-29).

**No pending changes**

# [1.0.1 (2022-11-29)](CHANGELOG.md#101-2022-11-29)

Open Beta [0.1.9](#019-2022-11-28-beta) promoted to GA.

# [0.1.9 (2022-11-28)](CHANGELOG.md#019-2022-11-28) Beta

This section outlines changes to the CD APIs, SDKs, and Terraform resources accrued since release [0.1.5](#015-2022-11-11-beta).

## Schedule

- November 28-30, 2022 - Changes to CD APIs and Go SDK will go live. This release includes breaking changes.
- December 1, 2022 - IBM Cloud Terraform Provider version `1.48.0` will be published. Date is approximate.

## Breaking changes to APIs and SDKs

### Offset-based pagination removed

`/toolchain/v2/toolchains`
`/toolhains/v2/toolchains/{toolchain_id}/tools`
`/pipeline/v2/tekton_pipelines/{pipeline_id}/pipeline_runs`

- The `offset` parameter for offset-based pagination is removed. Pagination is exclusively cursor based.

### Toolchain tags removed

`/toolhain/v2/toolchains/{toolchain_id}`

- The `tags` parameter is removed from toolchain resources. To work with tags, use [IBM Cloud Global Tagging](https://cloud.ibm.com/apidocs/tagging).

# [0.1.5 (2022-11-11)](CHANGELOG.md#015-2022-11-11) Beta

This section outlines changes to the CD APIs, SDKs, and Terraform resources accrued since release [0.1.2](#012-2022-09-09-beta). **This release includes breaking changes**.

## Schedule

- November 17, 2022 - IBM Cloud Terraform Provider version `1.48.0-beta0` is published, which corresponds to and uses this release of the Go SDK. This release includes breaking changes.
- November 18, 2022 - Corresponding changes to the IBM Cloud Continuous Delivery APIs are live. You should use this release of the Go SDK and IBM Cloud Terraform Provider version `1.48.0-beta0`.
- November 30, 2022 - IBM Cloud Terraform Provider version `1.48.0` will be published. Date is approximate.

## Breaking changes to APIs and SDKs

### Tool integrations

`/toolchain/v2/toolchains/{toolchain_id}/tools`

- The `limit` for GET is reduced from 200 to 150.

### Tekton pipelines

`/pipeline/v2/tekton_pipelines/{pipeline_id}`

- `enable_slack_notifications` renamed to `enable_notifications`
- `resource_group_id` restructured into `resource_group` object

### Tekton pipeline definitions

`/pipeline/v2/tekton_pipelines/{pipeline_id}/definitions/{definition_id}`

- `scm_source` object restructured to `source` object
- `service_instance_id` restructured to `tool` object

### Tekton pipeline triggers

`/pipeline/v2/tekton_pipelines/{pipeline_id}/triggers/{trigger_id}`

- For triggers of type `scm`, `scm_source` object restructured to `source` object
- For triggers of type `scm`, `service_instance_id` restructured to `tool` object
- `disabled` changed to `enabled`
- `events` reshaped from object of booleans into array of strings
- `scm_source` object restructured to `source` object

### Tekton pipeline runs

`/pipeline/v2/tekton_pipelines/{pipeline_id}/pipeline_runs`

- When fetching pipeline runs, in the response data `worker.agent` has been renamed to `worker.agent_id`
- When triggering a pipeline run, in the request body `trigger_header` has been renamed to `trigger_headers`

`/pipeline/v2/tekton_pipelines/{pipeline_id}/pipeline_runs/{run_id}`

- When fetching a pipeline run, in the response data `worker.agent` has been renamed to `worker.agent_id`
- When fetching a pipeline run, in the request body `event_header_params_blob` has been renamed to `trigger_headers`, matching with the property of the same name that was passed in the request body when triggering the run

## Backward compatible changes to APIs and SDKs

### Toolchain API authentication and authorization

- The error code for authentication failures is changed from HTTP 404 Not Found to HTTP 401 Unauthenticated.
- The error code when attempting a POST or PATCH of a toolchain or tool integration to which the caller has only Viewer access is changed from HTTP 404 Not Found to HTTP 403 Forbidden. The error code in which the caller has no access remains unchanged as HTTP 404 Not Found.

### Toolchains

`/toolchain/v2/toolchains/{toolchain_id}`

- New `ui_href` property added in responses to GET, POST, and PATCH toolchain calls.

## Breaking changes to Terraform resources

### ibm_cd_toolchain_tool_githubintegrated

Resource and datasource is removed. Use `ibm_cd_toolchain_tool_githubconsolidated` instead.

### ibm_cd_toolchain_tool_appconfig

Some nested properties within the `parameters` object are renamed:

- `region` --> `location` - Prefix `ibm:yp:` is no longer required on values
- `resource_group` --> `resource_group_name`
- `instance_name` --> `instance_id`
- `environment_name` --> `environment_id`
- `collection_name` --> `collection_id`

Example:

#### v1.46.0

```
resource "ibm_cd_toolchain_tool_appconfig" "my_appconfig_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    name = "appconfig_tool_01"
    region = "ibm:yp:us-south"
    resource_group = "Default"
    instance_name = "2a9e3c79-3595-45df-824d-9250aeb598c8"
    environment_name = "environment_01"
    collection_name = "collection_01"
  }
}
```

#### v1.48.0

```
resource "ibm_cd_toolchain_tool_appconfig" "my_appconfig_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    name = "appconfig_tool_01"
    location = "us-south"
    resource_group_name = "Default"
    instance_id = "2a9e3c79-3595-45df-824d-9250aeb598c8"
    environment_id = "environment_01"
    collection_id = "collection_01"
  }
}
```

### ibm_cd_toolchain_tool_jenkins

A nested property within the `parameters` object is changed:

- `webhook_url` is now a computed property

### ibm_cd_toolchain_tool_keyprotect

Some nested properties within the `parameters` object are renamed:

- `region` --> `location` - Prefix `ibm:yp:` is no longer required on values
- `resource_group` --> `resource_group_name`

Example:

#### v1.46.0

```
resource "ibm_cd_toolchain_tool_keyprotect" "my_keyprotect_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    name = "kp_tool_01"
    instance_name = "Key Protect-XX"
    region = "ibm:yp:us-south"
    resource_group = "Default"
  }
}
```

#### v1.48.0

```
resource "ibm_cd_toolchain_tool_keyprotect" "my_keyprotect_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    name = "kp_tool_01"
    instance_name = "Key Protect-XX"
    location = "us-south"
    resource_group_name = "Default"
  }
}
```

### ibm_cd_toolchain_tool_nexus

A nested property within the `parameters` object is renamed:

- `dashboard_url` --> `server_url`

Example:

#### v1.46.0

```
resource "ibm_cd_toolchain_tool_nexus" "my_nexus_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    name = "my-nexus"
    type = "npm"
    user_id = "<user_id>"
    token = "<token>"
    dashboard_url = "https://my.nexus.server.com/"
  }
}
```

#### v1.48.0

```
resource "ibm_cd_toolchain_tool_nexus" "my_nexus_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    name = "my-nexus"
    type = "npm"
    user_id = "<user_id>"
    token = "<token>"
    server_url = "https://my.nexus.server.com/"
  }
}
```

### ibm_cd_toolchain_tool_pagerduty

Creating a new PagerDuty service is no longer supported. The PagerDuty service must already exist. As a result, some nested properties within the `parameters1 object are removed:

- `key_type`
- `api_key`
- `service_name`
- `user_email`
- `user_phone`
- `service_url`

Some nested properties within the `parameters` object are changed:

- `service_key` is now a required property
- `service_url` is now a required property

Example:

#### v1.46.0

```
resource "ibm_cd_toolchain_tool_pagerduty" "my_pagerduty_toolchain" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    key_type = "api"
    api_key = "<api_key>"
    service_name = "AS34FR4"
    user_email = "<user_email>"
    user_phone = "<user_phone>"
  }
}
```

#### v1.48.0

```
resource "ibm_cd_toolchain_tool_pagerduty" "my_pagerduty_toolchain" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    service_url = "https://mycompany.example.pagerduty.com/services/AS34FR4"
    service_key = "<service_key>"
  }
}
```

### ibm_cd_toolchain_tool_pipeline

Some nested properties within the `parameters` object are removed:

- `type`
- `ui_pipeline`

Example:

#### v1.46.0

```
resource "ibm_cd_toolchain_tool_pipeline" "my_pipeline_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    name = "pipeline-tool-01"
    type = "tekton"
    ui_pipeline = false
  }
}
```

#### v1.48.0

```
resource "ibm_cd_toolchain_tool_pipeline" "my_pipeline_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    name = "pipeline-tool-01"
  }
}
```

### ibm_cd_toolchain_tool_privateworker

A nested property within the `parameters` object is changed:

- `worker_queue_identifier` is now a computed property

### ibm_cd_toolchain_tool_saucelabs

A nested property within the `parameters` object is renamed:

- `key` --> `access_key`

Example:

#### v1.46.0

```
resource "ibm_cd_toolchain_tool_saucelabs" "my_saucelabs_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    username = "<username>"
    key = "<access_key>"
  }
}
```

#### v1.48.0

```
resource "ibm_cd_toolchain_tool_saucelabs" "my_saucelabs_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    username = "<username>"
    access_key = "<access_key>"
  }
}
```

### ibm_cd_toolchain_tool_secretsmanager

Some nested properties within the `parameters` object are renamed:

- `region` --> `location` - Prefix `ibm:yp:` is no longer required on values
- `resource_group` --> `resource_group_name`

Example:

#### v1.46.0

```
resource "ibm_cd_toolchain_tool_secretsmanager" "my_secretsmanager_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    name = "sm_tool_01"
    instance_name = "Secrets Manager-XX"
    region = "ibm:yp:us-south"
    resource_group = "Default"
  }
}
```

#### v1.48.0

```
resource "ibm_cd_toolchain_tool_secretsmanager" "my_secretsmanager_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    name = "sm_tool_01"
    instance_name = "Secrets Manager-XX"
    location = "us-south"
    resource_group_name = "Default"
  }
}
```

### ibm_cd_toolchain_tool_securitycompliance

A nested property within the `parameters` object are renamed:

- `evidence_repo_name` --> `evidence_repo_url`

Some nested properties within the `parameters` object are removed:

- `location`
- `trigger_info`

Example:

#### v1.46.0

```
resource "ibm_cd_toolchain_tool_securitycompliance" "my_securitycompliance_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    name = "compliance"
    evidence_namespace = "cd"
    trigger_scan = "disabled"
    scope = "my-scope"
    profile = "IBM Cloud Security Best Practices v1.0.0"
    location = "https://cloud.ibm.com"
    evidence_repo_name = "https://github.example.com/<username>/compliance-evidence-<datestamp>"
  }
}
```

#### v1.48.0

```
resource "ibm_cd_toolchain_tool_securitycompliance" "my_securitycompliance_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    name = "compliance"
    evidence_namespace = "cd"
    trigger_scan = "disabled"
    scope = "my-scope"
    profile = "IBM Cloud Security Best Practices v1.0.0"
    evidence_repo_url = "https://github.example.com/<username>/compliance-evidence-<datestamp>"
  }
}
```

### ibm_cd_toolchain_tool_slack

Some nested properties within the `parameters` object are renamed:

- `api_token` --> `webhook`
- `team_url` --> `team_name`

Example:

#### v1.46.0

```
resource "ibm_cd_toolchain_tool_slack" "my_slack_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    channel_name = "#my_channel"
    pipeline_start = true
    pipeline_success = true
    pipeline_fail = true
    toolchain_bind = true
    toolchain_unbind = true
    api_token = "https://hooks.slack.com/services/A5EWRN5WK/A726ZQWT68G/TsdTjp6q4i6wFQTICTasjkE8"
    team_url = "my_team"
  }
}
```

#### v1.48.0

```
resource "ibm_cd_toolchain_tool_slack" "my_slack_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    channel_name = "#my_channel"
    pipeline_start = true
    pipeline_success = true
    pipeline_fail = true
    toolchain_bind = true
    toolchain_unbind = true
    webhook = "https://hooks.slack.com/services/A5EWRN5WK/A726ZQWT68G/TsdTjp6q4i6wFQTICTasjkE8"
    team_name = "my_team"
  }
}
```

### ibm_cd_toolchain_tool_sonarqube

A nested property within the `parameters` object is renamed:

- `dashboard_url` --> `server_url`

Example:

#### v1.46.0

```
resource "ibm_cd_toolchain_tool_sonarqube" "my_sonarqube_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    name = "my-sonarqube"
    user_login = "<user_login>"
    user_password = "<user_password>"
    blind_connection = true
    dashboard_url = "https://my.sonarqube.server.com/"
  }
}
```

#### v1.48.0

```
resource "ibm_cd_toolchain_tool_sonarqube" "my_sonarqube_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    name = "my-sonarqube"
    user_login = "<user_login>"
    user_password = "<user_password>"
    blind_connection = true
    server_url = "https://my.sonarqube.server.com/"
  }
}
```

### ibm_cd_toolchain_tool_bitbucket

A nested property within the `parameters` object is renamed:

- `has_issues` --> `toolchain_issues_enabled`

Example:

#### v1.46.0

```
resource "ibm_cd_toolchain_tool_bitbucket" "my_bitbucket_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    initialization {
        git_id = "bitbucketgit"
        owner_id = "<bitbucket-user-id>"
        repo_name = "myrepo"
        source_repo_url = "https://bitbucket.org/source-repo-owner/source-repo"
        type = "clone"
        private_repo = true
    }
    parameters {
        enable_traceability = false
        has_issues = true
    }
  }
}
```

#### v1.48.0

```
resource "ibm_cd_toolchain_tool_bitbucket" "my_bitbucket_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    initialization {
        git_id = "bitbucketgit"
        owner_id = "<bitbucket-user-id>"
        repo_name = "myrepo"
        source_repo_url = "https://bitbucket.org/source-repo-owner/source-repo"
        type = "clone"
        private_repo = true
    }
    parameters {
        enable_traceability = false
        integration_owner = "my-userid"
        toolchain_issues_enabled = true
    }
  }
}
```

### ibm_cd_toolchain_tool_githubconsolidated

A nested property within the `parameters` object si renamed:

- `has_issues` --> `toolchain_issues_enabled`

Example:

#### v1.46.0

```
resource "ibm_cd_toolchain_tool_githubconsolidated" "my_githubconsolidated_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    initialization {
        git_id = "github"
        owner_id = "<github-user-id>"
        repo_name = "myrepo"
        source_repo_url = "https://github.com/source-repo-owner/source-repo"
        type = "clone"
        private_repo = true
        auto_init = false
    }
    parameters {
        enable_traceability = false
        has_issues = true
    }
  }
}
```

#### v1.48.0

```
resource "ibm_cd_toolchain_tool_githubconsolidated" "my_githubconsolidated_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    initialization {
        git_id = "github"
        owner_id = "<github-user-id>"
        repo_name = "myrepo"
        source_repo_url = "https://github.com/source-repo-owner/source-repo"
        type = "clone"
        private_repo = true
        auto_init = false
    }
    parameters {
        enable_traceability = false
        integration_owner = "my-userid"
        toolchain_issues_enabled = true
    }
  }
}
```

### ibm_cd_toolchain_tool_gitlab

A nested property within the `parameters` object is renamed:

- `has_issues` --> `toolchain_issues_enabled`

Example:

#### v1.46.0

```
resource "ibm_cd_toolchain_tool_gitlab" "my_gitlab_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    initialization {
        git_id = "gitlab"
        owner_id = "<gitlab-user-id>"
        repo_name = "myrepo"
        source_repo_url = "https://gitlab.com/source-repo-owner/source-repo"
        type = "clone"
        private_repo = true
    }
    parameters {
        enable_traceability = false
        has_issues = true
    }
  }
}
```

#### v1.48.0

```
resource "ibm_cd_toolchain_tool_gitlab" "my_gitlab_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    initialization {
        git_id = "gitlab"
        owner_id = "<gitlab-user-id>"
        repo_name = "myrepo"
        source_repo_url = "https://gitlab.com/source-repo-owner/source-repo"
        type = "clone"
        private_repo = true
    }
    parameters {
        enable_traceability = false
        integration_owner = "my-userid"
        toolchain_issues_enabled = true
    }
  }
}
```

### ibm_cd_toolchain_tool_hostedgit

A nested property within the `parameters` object is renamed:

- `has_issues` --> `toolchain_issues_enabled`

Example:

#### v1.46.0

```
resource "ibm_cd_toolchain_tool_hostedgit" "my_hostedgit_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    initialization {
        owner_id = "<gitlab-user-id>"
        repo_name = "myrepo"
        source_repo_url = "https://us-south.git.cloud.ibm.com/source-repo-owner/source-repo"
        type = "clone"
        private_repo = true
    }
    parameters {
        enable_traceability = false
        has_issues = true
    }
  }
}
```

#### v1.48.0

```
resource "ibm_cd_toolchain_tool_hostedgit" "my_hostedgit_tool" {
  toolchain_id = ibm_cd_toolchain.toolchain.id
  parameters {
    initialization {
        git_id = "hostedgit"
        owner_id = "<gitlab-user-id>"
        repo_name = "myrepo"
        source_repo_url = "https://us-south.git.cloud.ibm.com/source-repo-owner/source-repo"
        type = "clone"
        private_repo = true
    }
    parameters {
        enable_traceability = false
        integration_owner = "my-userid"
        toolchain_issues_enabled = true
    }
  }
}
```

### ibm_cd_tekton_pipeline resource

- `enable_slack_notifications` is renamed to `enable_notifications`

Example:

#### v1.46.0

```
resource "ibm_cd_tekton_pipeline" "my_pipeline" {
  enable_slack_notifications = true
  pipeline_id = ibm_cd_toolchain_tool_pipeline.pipeline_tool.tool_id
  worker {
    id = "public"
  }
}
```

#### v1.48.0

```
resource "ibm_cd_tekton_pipeline" "my_pipeline" {
  enable_notifications = true
  pipeline_id = ibm_cd_toolchain_tool_pipeline.pipeline_tool.tool_id
  worker {
    id = "public"
  }
}
```

### ibm_cd_tekton_pipeline data source

- `resource_group_id` restructured into `resource_group` object

Example:

#### v1.46.0

```
output "data_pipeline_output" { value = data.ibm_cd_tekton_pipeline.get_pipeline }
```

Result:

```
data_pipeline_output = {
  "created_at" = "2022-09-22T12:50:53.041Z"
  "id" = "11112222-3333-4444-5555-666677778888"
  "name" = "my-sample-pipeline"
  "pipeline_id" = "11112222-3333-4444-5555-666677778888"
  "resource_group_id" = "99992222-3333-4444-5555-666677778888"
}
```

#### v1.48.0

```
output "data_pipeline_output" { value = data.ibm_cd_tekton_pipeline.get_pipeline }
```

Result:

```
data_pipeline_output = {
  "created_at" = "2022-09-22T12:50:53.041Z"
  "id" = "11112222-3333-4444-5555-666677778888"
  "name" = "my-sample-pipeline"
  "pipeline_id" = "11112222-3333-4444-5555-666677778888"
  "resource_group" {
      "id" = "99992222-3333-4444-5555-666677778888"
  }
}
```

### ibm_cd_tekton_pipeline_definition resource

- `scm_source` object restructured

Example:

#### v1.46.0

```
resource "ibm_cd_tekton_pipeline_definition" "my_definition" {
  pipeline_id = ibm_cd_toolchain_tool_pipeline.pipeline_tool.tool_id
  scm_source {
    branch = "master"
    path = ".tekton"
    url = "https://github.com/IBM/tekton-tutorial.git"
  }
}
```

#### v1.48.0

```
resource "ibm_cd_tekton_pipeline_definition" "my_definition" {
  pipeline_id = ibm_cd_toolchain_tool_pipeline.pipeline_tool.tool_id
  source {
    type = "git"
    properties {
      branch = "master"
      path = ".tekton"
      url = "https://github.com/IBM/tekton-tutorial.git"
    }
  }
}
```

### ibm_cd_tekton_pipeline_definition data source output

- `scm_source` object restructured
- `service_instance_id` restructured to `tool` object

Example:

#### v1.46.0

```
output "data_definition_output" { value=data.ibm_cd_tekton_pipeline_definition.my_definition }
```

Result:

```
data_definition_output = {
  "definition_id" = "11112222-3333-4444-5555-666677778888"
  "id" = "aaaabbbb-cccc-dddd-eeee-ffff00001111/11112222-3333-4444-5555-666677778888"
  "pipeline_id" = "aaaabbbb-cccc-dddd-eeee-ffff00001111"
  "scm_source" = tolist([{
      "branch" = "master"
      "path" = ".tekton"
      "service_instance_id" = "99990000-aaaa-ffff-1111-8888bbbbcccc"
      "url" = "https://github.com/IBM/tekton-tutorial.git"
  }])
}
```

#### v1.48.0

```
output "data_definition_output" { value=data.ibm_cd_tekton_pipeline_definition.my_definition }
```

Result:

```
data_definition_output = {
  "definition_id" = "11112222-3333-4444-5555-666677778888"
  "id" = "aaaabbbb-cccc-dddd-eeee-ffff00001111/11112222-3333-4444-5555-666677778888"
  "pipeline_id" = "aaaabbbb-cccc-dddd-eeee-ffff00001111"
  "source" = tolist([{
    "type" = "git"
    "properties" = tolist([{
      "branch" = "master"
      "path" = ".tekton"
      "url" = "https://github.com/IBM/tekton-tutorial.git"
      "tool" = tolist([{
        "id": "99990000-aaaa-ffff-1111-8888bbbbcccc"
      }])
    }])
  }])
}
```

### ibm_cd_tekton_pipeline_trigger resource

- `disabled` changed to `enabled`
- `events` reshaped from object of booleans to array of strings
- `scm_source` object restructured

Example:

#### v1.46.0

```
resource "ibm_cd_tekton_pipeline_trigger" "my_trigger" {
  pipeline_id = ibm_cd_toolchain_tool_pipeline.pipeline_tool.tool_id
  type = "scm"
  name = "my_trigger"
  event_listener = "listener"
  disabled = "false"
  scm_source {
      url = "https://github.com/IBM/tekton-tutorial.git"
      branch = "master"
  }
  events = {
    push: "true"
    pull_request: "true"
  }
}
```

#### v1.48.0

```
resource "ibm_cd_tekton_pipeline_trigger" "my_trigger" {
  pipeline_id = ibm_cd_toolchain_tool_pipeline.pipeline_tool.tool_id
  type = "scm"
  name = "my_trigger"
  event_listener = "listener"
  enabled = "true"
  source {
    type = "git"
    properties {
      url = "https://github.com/IBM/tekton-tutorial.git"
      branch = "master"
    }
  }
  events = ["push", "pull_request"]
}
```

### ibm_cd_tekton_pipeline_trigger data source output

- For triggers of type `scm`, `scm_source` object restructured to `source` object
- For triggers of type `scm`, `service_instance_id` restructured to `tool` object

Example:

#### v1.46.0

```
output "data_trigger_output" { value=data.ibm_cd_tekton_pipeline_trigger.my_trigger }
```

Result:

```
data_trigger_output = {
  "type" = "scm"
  "trigger_id" = "11112222-3333-4444-5555-666677778888"
  "event_listener" = "my-listener"
  "id" = "aaaabbbb-cccc-dddd-eeee-ffff00001111/11112222-3333-4444-5555-666677778888"
  "name" = "my-scm-trigger"
  "pipeline_id" = "aaaabbbb-cccc-dddd-eeee-ffff00001111"
  "scm_source" = tolist([{
      "branch" = "master"
      "service_instance_id" = "99990000-aaaa-ffff-1111-8888bbbbcccc"
      "url" = "https://github.com/IBM/tekton-tutorial.git"
    }])
}
```

#### v1.48.0

```
output "data_definition_output" { value=data.ibm_cd_tekton_pipeline_definition.my_definition }
```

Result:

```
data_trigger_output = {
  "type" = "scm"
  "trigger_id" = "11112222-3333-4444-5555-666677778888"
  "event_listener" = "my-listener"
  "id" = "aaaabbbb-cccc-dddd-eeee-ffff00001111/11112222-3333-4444-5555-666677778888"
  "name" = "my-scm-trigger"
  "pipeline_id" = "aaaabbbb-cccc-dddd-eeee-ffff00001111"
  "source" = tolist([{
    "type" = "git"
    "properties" = tolist([{
      "branch" = "master"
      "tool" = tolist([{
        "id": "99990000-aaaa-ffff-1111-8888bbbbcccc"
      }])
      "url" = "https://github.com/IBM/tekton-tutorial.git"
    }])
  }])
}
```

## Backward compatible changes to Terraform resources

### ibm_cd_toolchain

A new computed property is added:

- `ui_href`

### ibm_cd_toolchain_tool_...

A new computed property is added to all tool integration resources:

- `ui_href`

### ibm_cd_toolchain_tool_jira

This is a new resource and data source

### ibm_cd_toolchain_tool_bitbucket

The `type` property within the `initialization` object supports new values:
- `type` can be set to `clone_if_not_exists`
- `type` can be set to `fork_if_not_exists`
- `type` can be set to `new_if_not_exists`

A nested property within the `parameters` object is changed:

- `integration_owner` is now a configurable property

New nested computed properties are added to the `parameters` object:

- `default_branch`
- `repo_id`

### ibm_cd_toolchain_tool_githubconsolidated

The `type` property within the `initialization` object supports new values:
- `type` can be set to `clone_if_not_exists`
- `type` can be set to `fork_if_not_exists`
- `type` can be set to `new_if_not_exists`

A nested property within the `parameters` object is changed:

- `integration_owner` is now a configurable property

New nested computed properties are added to the `parameters` object:

- `default_branch`
- `repo_id`

### ibm_cd_toolchain_tool_gitlab

The `type` property within the `initialization` object supports new values:

- `type` can be set to `clone_if_not_exists`
- `type` can be set to `fork_if_not_exists`
- `type` can be set to `new_if_not_exists`

A nested property within the `parameters` object is changed:

- `integration_owner` is now a configurable property

New nested computed properties are added to the `parameters` object:

- `default_branch`
- `repo_id`

### ibm_cd_toolchain_tool_hostedgit

The `type` property within the `initialization` object supports new values:

- `type` can be set to `clone_if_not_exists`
- `type` can be set to `fork_if_not_exists`
- `type` can be set to `new_if_not_exists`

A nested property within the `parameters` object is changed:

- `integration_owner` is now a configurable property

New nested computed properties are added to the `parameters` object:

- `default_branch`
- `repo_id`

A new nested property is added to the `initialization` object:

- `git_id`

# [0.1.2 (2022-09-09)](CHANGELOG.md#012-2022-09-09) Beta

## Schedule

- October 3, 2022 - IBM Cloud Terraform Provider version 1.46.0 is published: https://registry.terraform.io/providers/IBM-Cloud/ibm/1.46.0. The CD Tekton Pipeline and CD Toolchain Beta resources and data sources within this Provider contain a number of breaking changes. These changes will not work with the IBM Cloud Continuous Delivery service APIs until we deliver corresponding changes to the APIs. Until then (see next bullet), you should use IBM Cloud Terraform Provider version 1.45.1 highest.
- October 5, 2022 - Breaking changes to the IBM Cloud Continuous Delivery service’s v2 Beta HTTP APIs and Go SDKs are deployed to production. These changes are compatible with IBM Cloud Terraform Provider versions 1.46.0-beta0 and 1.46.0. Once these changes are deployed, you should use IBM Cloud Terraform Provider version 1.46.0.

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

```json
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

```curl
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

```json
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

```terraform
data "ibm_cd_tekton_pipeline" "get_pipeline" {
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
output "data_pipeline_output" { value = data.ibm_cd_tekton_pipeline.get_pipeline }
```

Result

```terraform
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

```terraform
data "ibm_cd_tekton_pipeline" "get_pipeline" {
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
output "data_pipeline_output" { value = data.ibm_cd_tekton_pipeline.get_pipeline }
```

Result

```terraform
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

```terraform
data "ibm_cd_tekton_pipeline_definition" "get_definition" {
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
  definition_id = ibm_cd_tekton_pipeline_definition.my_definition.definition_id
}
output "data_definition_output" { value=data.ibm_cd_tekton_pipeline_definition.get_definition }
```

Result

```terraform
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

```terraform
data "ibm_cd_tekton_pipeline_definition" "get_definition" {
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
  definition_id = ibm_cd_tekton_pipeline_definition.my_definition.definition_id
}
output "data_definition_output" { value=data.ibm_cd_tekton_pipeline_definition.get_definition }
```

Result

```terraform
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

```terraform
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

```terraform
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

```terraform
data "ibm_cd_tekton_pipeline_trigger" "get_scm_trigger" {
 pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
 trigger_id = ibm_cd_tekton_pipeline_trigger.my_scm_trigger.trigger_id
}
output "data_output_scm_trigger" { value=data.ibm_cd_tekton_pipeline_trigger.get_scm_trigger }
```

Result

```terraform
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

```terraform
data "ibm_cd_tekton_pipeline_trigger" "get_scm_trigger" {
 pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
 trigger_id = ibm_cd_tekton_pipeline_trigger.my_scm_trigger.trigger_id
}
output "data_output_scm_trigger" { value=data.ibm_cd_tekton_pipeline_trigger.get_scm_trigger }
```

Result

```terraform
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

```terraform
resource "ibm_cd_tekton_pipeline_property" "my-text-prop" {
  name = "evidence-repo"
  type = "TEXT"
  value = "sample-text-value"
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
```

Version 1.46.0-beta0, 1.46.0

```terraform
resource "ibm_cd_tekton_pipeline_property" "my-text-prop" {
  name = "evidence-repo"
  type = "text"
  value = "sample-text-value"
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
```

### Example: type `secure`

Version 1.45.0

```terraform
resource "ibm_cd_tekton_pipeline_property" "my-secure-prop" {
  name = "my-api-key"
  type = "SECURE"
  value = "my-secret-value"
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
```

Version 1.46.0-beta0, 1.46.0

```terraform
resource "ibm_cd_tekton_pipeline_property" "my-secure-prop" {
  name = "my-api-key"
  type = "secure"
  value = "my-secret-value"
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id      
}
```

### Example: type `single_select`

Version 1.45.0

```terraform
resource "ibm_cd_tekton_pipeline_property" "regions" {
  name = "regions"
  type = "SINGLE_SELECT"
  enum = ["us-south", "eu-gb", "eu-de"]
  default = "eu-de"
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
```

Version 1.46.0-beta0, 1.46.0

```terraform
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

```terraform
resource "ibm_cd_tekton_pipeline_property" "integration_property" {
  name = "my_integration_prop"
  value = ibm_cd_toolchain_tool_githubconsolidated.repo1.tool_id
  type = "INTEGRATION"
  path = "parameters.repo_url"
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
```

Version 1.46.0-beta0, 1.46.0

```terraform
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

```terraform
resource "ibm_cd_tekton_pipeline_property" "appconfig_property" {
  name = "myAppConfig"
  value = "{appconfig::11112222-3333-4444-5555-666677778888.prop.myproperty}"
  type = "APPCONFIG"
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
```

Version 1.46.0-beta0, 1.46.0

```terraform
resource "ibm_cd_tekton_pipeline_property" "appconfig_property" {
  name = "myAppConfig"
  value = "{appconfig::11112222-3333-4444-5555-666677778888.prop.myproperty}"
  type = "appconfig"
  pipeline_id = ibm_cd_tekton_pipeline.my_pipeline.pipeline_id
}
```

### Example: type `digest_matches`

Version 1.45.0

```terraform
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

```terraform
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

```terraform
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

```terraform
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

```terraform
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

```terraform
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

```terraform
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

```terraform
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
