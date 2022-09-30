## [0.1.2](https://github.com/IBM/continuous-delivery-go-sdk/compare/v0.1.1...v0.1.2) (2022-09-09)

### Features

* feat(toolchain): Latest toolchain SDK produced from otc-api swagger ([#14](https://github.com/IBM/continuous-delivery-go-sdk/pull/14))

## [0.1.1](https://github.com/IBM/continuous-delivery-go-sdk/compare/v0.1.0...v0.1.1) (2022-09-09)

### Bug Fixes

* **tekton:** fix validator warnings ([#13](https://github.com/IBM/continuous-delivery-go-sdk/pull/13)) ([dfda869](https://github.com/IBM/continuous-delivery-go-sdk/commit/dfda869446493414aef36a616d5755d73a9274fc))

### Features

* feat(tekton): add webhookUrl to generic trigger data ([#13](https://github.com/IBM/continuous-delivery-go-sdk/pull/13)) ([dfda869](https://github.com/IBM/continuous-delivery-go-sdk/commit/dfda869446493414aef36a616d5755d73a9274fc))
* feat(tekton): added support for "enable_slack_notifications" and "enable_partial_cloning” settings in pipeline
### BREAKING CHANGES

* feat(tekton): rename pipeline output values
  * "created" renamed to "created_at"
  * "html_url" renamed to "runs_url"
* feat(tekton): refactor definition object
  * "service_instance_id" relocated inside "scm_source" object
* feat(tekton): refactor triggers of type "scm"
  * "service_instance_id" relocated inside "scm_source" object

# [0.1.0](https://github.com/IBM/continuous-delivery-go-sdk/compare/v0.0.8...v0.1.0) (2022-09-05)

### Features

* feat(tekton): add token pagination when fetching PipelineRuns

### Bug Fixes

* fix(tekton): string and schema updates ([#9](https://github.com/IBM/continuous-delivery-go-sdk/issues/9)) ([58a498a](https://github.com/IBM/continuous-delivery-go-sdk/commit/58a498a4ee9e031dcd0fa7bc141a3b69c95d8f34))
* fix(tekton): remove enabled flag

### BREAKING CHANGES

* feat(tekton): The schema of pipeline properties, trigger properties and generic trigger secrets have been refactored
  * rename property types to lowercase ["TEXT", "SINGLE_SELECT", "SECURE", "INTEGRATION", "APPCONFIG"] -> ["text", "single_select", "secure", "integration", "appconfig"]
  * rename secret types to snake case ["digestMatches", "tokenMatches", "internalValidation"] -> ["digest_matches", "token_matches", "internal_validation"]
* fix(tekton): update patch requests to use json-merge content type
* fix(tekton): refactor single_select type property, remove the `default` property
* feat(tekton): decouple create trigger and duplicate trigger APIs
* fix(tekton): refactor Trigger
  * remove wrapping `trigger` object
  * replace `concurrency` object with `max_concurrent_runs` integer property

## [0.0.8](https://github.com/IBM/continuous-delivery-go-sdk/compare/v0.0.7...v0.0.8) (2022-08-02)

### Bug Fixes

* chore(semantic versioning): enable automatic semantic versioning ([#8](https://github.com/IBM/continuous-delivery-go-sdk/issues/8)) ([30c5ee5](https://github.com/IBM/continuous-delivery-go-sdk/commit/30c5ee58454ed56bd220a60bb239c92514869037))
