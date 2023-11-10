# [1.3.0](https://github.com/IBM/continuous-delivery-go-sdk/compare/v1.2.1...v1.3.0) (2023-11-10)


### Features

* support eu-es region ([#65](https://github.com/IBM/continuous-delivery-go-sdk/issues/65)) ([3d6eaa0](https://github.com/IBM/continuous-delivery-go-sdk/commit/3d6eaa07a42b34c1c2ebc0982f2c5e99dec00f98)), closes [#63](https://github.com/IBM/continuous-delivery-go-sdk/issues/63)

## [1.2.1](https://github.com/IBM/continuous-delivery-go-sdk/compare/v1.2.0...v1.2.1) (2023-10-27)


### Bug Fixes

* **tekton:** add error message to PipelineRun ([4b65026](https://github.com/IBM/continuous-delivery-go-sdk/commit/4b6502699a6d1cd90b89c87165b2a478ecf7deff))

# [1.2.0](https://github.com/IBM/continuous-delivery-go-sdk/compare/v1.1.3...v1.2.0) (2023-09-19)


### Bug Fixes

* upgrade deploy go version in travis build ([#60](https://github.com/IBM/continuous-delivery-go-sdk/issues/60)) ([0dd3c40](https://github.com/IBM/continuous-delivery-go-sdk/commit/0dd3c406565b6815941e2cf4c09bf6adfdb5079e))


### Features

* **toolchain:** add name query param to list toolchains ([a6b0a30](https://github.com/IBM/continuous-delivery-go-sdk/commit/a6b0a30a16849c12584ff78ff8f316b32e922cfb))

## [1.1.3](https://github.com/IBM/continuous-delivery-go-sdk/compare/v1.1.2...v1.1.3) (2023-08-25)


### Bug Fixes

* **tekton:** Fix PipelineRun trigger properties ([a0b0932](https://github.com/IBM/continuous-delivery-go-sdk/commit/a0b09327c82a19f0e87db11f293e1ce5674213df))

## [1.1.2](https://github.com/IBM/continuous-delivery-go-sdk/compare/v1.1.1...v1.1.2) (2023-08-01)


### Bug Fixes

* **tekton:** add trigger favorite field ([cd01eba](https://github.com/IBM/continuous-delivery-go-sdk/commit/cd01ebad52adaccec81a3620889b86dc6c39b165))

## [1.1.1](https://github.com/IBM/continuous-delivery-go-sdk/compare/v1.1.0...v1.1.1) (2023-05-30)


### Bug Fixes

* **tekton:** update nextBuildNumber ([6b59938](https://github.com/IBM/continuous-delivery-go-sdk/commit/6b59938a786b052408c70b81c194164aefd24676))

# [1.1.0](https://github.com/IBM/continuous-delivery-go-sdk/compare/v1.0.5...v1.1.0) (2023-03-20)


### Features

* **tekton:** add next build number support ([3dbec03](https://github.com/IBM/continuous-delivery-go-sdk/commit/3dbec03f6bb5dbf90354a744ac214ce42bf45fae))

## [1.0.5](https://github.com/IBM/continuous-delivery-go-sdk/compare/v1.0.4...v1.0.5) (2023-02-27)


### Bug Fixes

* sync with latest template repo updates ([38ad7ff](https://github.com/IBM/continuous-delivery-go-sdk/commit/38ad7ff1dd94bb49ad260936ae764001e32b503e))

## [1.0.4](https://github.com/IBM/continuous-delivery-go-sdk/compare/v1.0.3...v1.0.4) (2023-01-23)


### Bug Fixes

* **tekton:** fix for empty tags and events ([9976e87](https://github.com/IBM/continuous-delivery-go-sdk/commit/9976e878600e37a9d454d9edf22244d4fdaed196))

## [1.0.3](https://github.com/IBM/continuous-delivery-go-sdk/compare/v1.0.2...v1.0.3) (2023-01-17)


### Bug Fixes

* **tekton:** WorkerIdentity and href updates ([5482480](https://github.com/IBM/continuous-delivery-go-sdk/commit/54824807d0e8ce5b43ae627081eaddae705efb51))

## [1.0.2](https://github.com/IBM/continuous-delivery-go-sdk/compare/v1.0.1...v1.0.2) (2023-01-11)


### Bug Fixes

* **build:** fix semVer issue ([3e7a4a8](https://github.com/IBM/continuous-delivery-go-sdk/commit/3e7a4a852e1f5e991bf35ac2cfe0aa81def666f8))
* **build:** remove release line from package.json ([3513289](https://github.com/IBM/continuous-delivery-go-sdk/commit/35132899f607e6248b1e018a421ca4acb6ce2252))
* **tekton:** fix integration tests ([d725ad4](https://github.com/IBM/continuous-delivery-go-sdk/commit/d725ad4141e00271b7059f78d8ab98e3d19f7841))
* **tekton:** post GA updates for tekton APIs ([53b6220](https://github.com/IBM/continuous-delivery-go-sdk/commit/53b6220e93571ac894076142234f5b33e143c01f))
* **tekton:** updates based on API linting ([1a3fd13](https://github.com/IBM/continuous-delivery-go-sdk/commit/1a3fd137fc4d0dca3903ee4ab0e753a7d57a05a4))

## [1.0.1](https://github.com/IBM/continuous-delivery-go-sdk/compare/v1.0.0...v1.0.1) (2022-11-29)


### Bug Fixes

* **doc:** Promote Open Beta to GA ([4d5a2f6](https://github.com/IBM/continuous-delivery-go-sdk/commit/4d5a2f6d6e28924a0fefce614f07e4c86f8a52f2))

# [1.0.0](https://github.com/IBM/continuous-delivery-go-sdk/compare/v0.1.9...v1.0.0) (2022-11-29)


### Features

* **build:** dummy commit to trigger major version build ([d176144](https://github.com/IBM/continuous-delivery-go-sdk/commit/d176144866e81c690843a3a780729d614f8daf6a))


### BREAKING CHANGES

* **build:** retract v1.0.0 and create dummy v1.0.0 tag/release in github

Signed-off-by: Brian Gleeson <brian.gleeson@ie.ibm.com>

## [0.1.9](https://github.com/IBM/continuous-delivery-go-sdk/compare/v0.1.8...v0.1.9) (2022-11-28)


### Bug Fixes

* **tekton:** final GA review changes ([7dcdb4b](https://github.com/IBM/continuous-delivery-go-sdk/commit/7dcdb4b7f39539c0c0e3af6b84908f7e28951bb1))

## [0.1.8](https://github.com/IBM/continuous-delivery-go-sdk/compare/v0.1.7...v0.1.8) (2022-11-24)


### Bug Fixes

* retract package version 1.0.0 ([3071c19](https://github.com/IBM/continuous-delivery-go-sdk/commit/3071c19dbdf0a50698ed08471fe251ebe0f0fe6a))
* **toolchain:** remove offset pagination parameter ([b67da5a](https://github.com/IBM/continuous-delivery-go-sdk/commit/b67da5a75ce7456e7f37c670b52f2186bfbed870))

## [0.1.7](https://github.com/IBM/continuous-delivery-go-sdk/compare/v0.1.6...v0.1.7) (2022-11-23)


### Bug Fixes

* **toolchain:** remove offset pagination parameter ([#30](https://github.com/IBM/continuous-delivery-go-sdk/issues/30)) ([5ffec28](https://github.com/IBM/continuous-delivery-go-sdk/commit/5ffec2827bf178d6a18076dd5b12a1b0c9ed70e0))

## [0.1.6](https://github.com/IBM/continuous-delivery-go-sdk/compare/v0.1.5...v0.1.6) (2022-11-21)


### Bug Fixes

* **toolchain:** remove 'offset' and 'limit' from list examples ([#29](https://github.com/IBM/continuous-delivery-go-sdk/issues/29)) ([cc1de6a](https://github.com/IBM/continuous-delivery-go-sdk/commit/cc1de6a1e47e5002f7e7c01f6d227da649622ac8))

## [0.1.5](https://github.com/IBM/continuous-delivery-go-sdk/compare/v0.1.4...v0.1.5) (2022-11-11)


### Bug Fixes

* **tekton:** fix up integration tests ([711d693](https://github.com/IBM/continuous-delivery-go-sdk/commit/711d69331ebad33afaf59b6088fc8bbbdae39f90))
* **tekton:** ga review updates ([d28353a](https://github.com/IBM/continuous-delivery-go-sdk/commit/d28353a420143e559ba93c17d0b267ff77eab241))
* **tekton:** latest GA changes ([4e237f1](https://github.com/IBM/continuous-delivery-go-sdk/commit/4e237f1b143ff4231899f2988e36f1cfc1f25414))
* **tekton:** reshape service_instance_id ([1307184](https://github.com/IBM/continuous-delivery-go-sdk/commit/1307184862f1b7ca71844c6a325e6ff773e98187))

## [0.1.4](https://github.com/IBM/continuous-delivery-go-sdk/compare/v0.1.3...v0.1.4) (2022-10-12)


### Bug Fixes

* **tekton:** documentation change ([d3692c5](https://github.com/IBM/continuous-delivery-go-sdk/commit/d3692c50baf6d40b826af83c60547070f0039374))

## [0.1.3](https://github.com/IBM/continuous-delivery-go-sdk/compare/v0.1.2...v0.1.3) (2022-10-06)


### Bug Fixes

* **toolchain:** Added UI href property in toolchain response ([#20](https://github.com/IBM/continuous-delivery-go-sdk/issues/20)) ([b5397a6](https://github.com/IBM/continuous-delivery-go-sdk/commit/b5397a6d3e84f6f62f7dbfe2c8c240b134fe6fcc))

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
