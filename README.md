# IBM Cloud Continuous Delivery Go SDK 2.0.2

[![Build Status](https://app.travis-ci.com/IBM/continuous-delivery-go-sdk.svg?branch=main)](https://app.travis-ci.com/github/IBM/continuous-delivery-go-sdk)
[![Release](https://img.shields.io/github/v/release/IBM/continuous-delivery-go-sdk)](https://github.com/IBM/continuous-delivery-go-sdk/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/IBM/continuous-delivery-go-sdk.svg)](https://pkg.go.dev/github.com/IBM/continuous-delivery-go-sdk)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/IBM/continuous-delivery-go-sdk)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

The Go client library to interact with the [IBM Cloud Continuous Delivery Toolchain and Tekton Pipeline APIs](https://cloud.ibm.com/docs?tab=api-docs&category=devops).

## Table of Contents

<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [IBM Cloud Continuous Delivery Go SDK 2.0.2](#ibm-cloud-continuous-delivery-go-sdk-201)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
    - [Go modules](#go-modules)
    - [`go get` command](#go-get-command)
  - [Using the SDK](#using-the-sdk)
  - [Questions](#questions)
  - [Issues](#issues)
  - [Open source @ IBM](#open-source--ibm)
  - [Contributing](#contributing)
  - [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud Continuous Delivery Go SDK allows developers to programmatically interact with the IBM Cloud services that are listed in Table 1:

Service name | Package name
--- | ---
[Toolchain API](https://cloud.ibm.com/apidocs/toolchain?code=go) | cdtoolchainv2
[Tekton Pipeline API](https://cloud.ibm.com/apidocs/tekton-pipeline?code=go) | cdtektonpipelinev2

Table 1. IBM Cloud services

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

- An [IBM Cloud][ibm-cloud-onboarding] account.
- An IAM API key to allow the SDK to access your account. Create an [API key](https://cloud.ibm.com/iam/apikeys).
- Go version 1.23 or above.

## Installation

The current version of this SDK: 2.0.2

### Go modules

If your application uses Go modules for dependency management (recommended), add an import for each service that you use in your application. 

Example:

```go
import (
  "github.com/IBM/continuous-delivery-go-sdk/v2/cdtoolchainv2"
)
```

Next, run `go build` or `go mod tidy` to download and install the new dependencies and update your application's `go.mod` file.  

In the example, the `cdtoolchainv2` part of the import path is the package name that is associated with the Toolchain service. See Table 1 to find the appropriate package name for the services that your application uses.

### `go get` command

Alternatively, you can use the `go get` command to download and install the appropriate packages that your application uses:

```sh
go get -u github.com/IBM/continuous-delivery-go-sdk/v2/cdtoolchainv2
```

Be sure to use the appropriate package name from Table 1 for the services that your application uses.

## Using the SDK

For general SDK usage information, see [IBM Cloud SDK Common README](https://github.com/IBM/ibm-cloud-sdk-common/blob/main/README.md).

## Questions

If you are having difficulties using this SDK or you have a question about the IBM Cloud services, ask a question at [Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).
Alternatively, you can reach out to the IBM Cloud Continuous Delivery development team by joining us on [Slack](https://ic-devops-slack-invite.us-south.devops.cloud.ibm.com/).

## Issues

If you have a problem with the project, submit a [bug report](https://github.com/IBM/continuous-delivery-go-sdk/issues), but before you do, search for similar problems. Someone else might have already reported the problem.

## Open source @ IBM

Find more open source projects on the [IBM GitHub Page](http://ibm.github.io/).

## Contributing

For more information, see [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license. You can find the license's full text in [LICENSE](LICENSE).
