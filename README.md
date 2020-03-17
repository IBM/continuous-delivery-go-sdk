[![Build Status](https://travis.ibm.com/CloudEngineering/go-sdk-template.svg?token=eW5FVD71iyte6tTby8gr&branch=master)](https://travis.ibm.com/CloudEngineering/go-sdk-template.svg?token=eW5FVD71iyte6tTby8gr&branch=master)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

# IBM Cloud MySDK Go SDK Version 0.0.1
Go client library to interact with the various [MySDK Service APIs](https://cloud.ibm.com/apidocs?category=platform_services).

## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
    + [`go get` command](#go-get-command)
    + [Go modules](#go-modules)
    + [`dep` dependency manager](#dep-dependency-manager)
- [Using the SDK](#using-the-sdk)
  * [Constructing service clients](#constructing-service-clients)
    + [Setting service client options programmatically](#setting-service-client-options-programmatically)
    + [Constructing a service client using external configuration](#constructing-a-service-client-using-external-configuration)
      - [Define configuration properties](#define-configuration-properties)
      - [Construct service client](#construct-service-client)
      - [Storing configuration properties in a file](#storing-configuration-properties-in-a-file)
      - [Complete configuration-loading process](#complete-configuration-loading-process)
  * [Authentication](#authentication)
    + [Example: construct IamAuthenticator with an IAM api key](#example-construct-iamauthenticator-with-an-iam-api-key)
    + [Example: construct BearerTokenAuthenticator with an access token](#example-construct-bearertokenauthenticator-with-an-access-token)
  * [Passing operation parameters via an "options" struct](#passing-operation-parameters-via-an-options-struct)
  * [Receiving operation responses](#receiving-operation-responses)
      - [Examples](#examples)
  * [Error Handling](#error-handling)
  * [Sending HTTP headers](#sending-http-headers)
    + [Sending HTTP headers with all requests](#sending-http-headers-with-all-requests)
    + [Sending request HTTP headers](#sending-request-http-headers)
- [Questions](#questions)
- [Open source @ IBM](#open-source--ibm)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud MySDK Go SDK allows developers to programmatically interact with the following IBM Cloud services:

Service Name | Package name 
--- | --- 
[Example Service](https://cloud.ibm.com/apidocs/example-service) | exampleservicev1

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration?target=%2Fdeveloper%2Fwatson&

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.12 or above.

## Installation
The current version of this SDK: 0.0.1

There are a few different ways to download and install the MySDK Go SDK project for use by your
Go application:

#### `go get` command  
Use this command to download and install the MySDK Go SDK project to allow your Go application to
use it:

```
go get -u github.ibm.com/CloudEngineering/go-sdk-template
```

#### Go modules  
If your application is using Go modules, you can add a suitable import to your
Go application, like this:

```go
import (
	"github.ibm.com/CloudEngineering/go-sdk-template/exampleservicev1"
)
```

then run `go mod tidy` to download and install the new dependency and update your Go application's
`go.mod` file.

#### `dep` dependency manager  
If your application is using the `dep` dependency management tool, you can add a dependency
to your `Gopkg.toml` file.  Here is an example:

```
[[constraint]]
  name = "github.ibm.com/CloudEngineering/go-sdk-template"
  version = "0.0.1"

```

then run `dep ensure`.

## Using the SDK
This section provides general information on how to use the services contained in this SDK.

### Constructing service clients
Each service is implemented in its own package (e.g. `exampleservicev1`).
The package will contain a "service client" struct (a client-side representation of the service), as well as an "options" struct that is used to construct instances of the service client.

#### Setting service client options programmatically
Here's an example of how to construct an instance of a service (ExampleServiceV1) while specifying service client options
(authenticator, service endpoint URL, etc.) programmatically:

```go
import {
    "github.com/IBM/go-sdk-core/v3/core"
    "github.ibm.com/CloudEngineering/go-sdk-template/exampleservicev1"
}

// Create an authenticator.
authenticator := &core.IamAuthenticator{
    ApiKey: "my-iam-apikey",
}

// Set our custom endpoint base URL (optional)
serviceURL := "https://myservice.cloud.ibm.com/v1"

// Create an instance of the "ExampleServiceV1Options"  struct.
options := &exampleservicev1.ExampleServiceV1Options{
    URL: serviceURL,
    Authenticator: authenticator,
}

// Create an instance of the "ExampleServiceV1" service client.
myservice, err := exampleservicev1.NewExampleServiceV1(options)
if err != nil {
    // handle error
}

// Service operations can now be called using the "myservice" variable.

```

#### Constructing a service client using external configuration
For a typical application deployed to the IBM Cloud, it might be convenient to avoid hard-coding
certain service client options (IAM API Key, service endpoint URL, etc.).
Instead, the SDK allows you to store these values in configuration properties external to your
application.

##### Define configuration properties
First, define the configuration properties to be used by your application.  These properties
can be implemented as either (1) exported environment variables or (2) stored in a *credentials* file.
In the examples that follow, we'll use environment variables to implement our configuration properties.
Each property name is of the form: `<serviceName>_<propertyKey>`.
Here is an example of some configuration properties for the Resource Controller service:

```
export EXAMPLE_SERVICE_URL=https://myservice.cloud.ibm.com/v2
export EXAMPLE_SERVICE_AUTH_TYPE=iam
export EXAMPLE_SERVICE_APIKEY=my-iam-apikey
```

The service name "example_service" is the default service name for the "Example" service,
so the SDK will (by default) look for properties that start with this prefix folded to upper case.

##### Construct service client
After you have defined the configuration properties for your application, you can
construct an instance of the service client like this:

```go
myservice, err := exampleservicev1.NewExampleServiceV1UsingExternalConfig(
        &exampleservicev1.ExampleServiceV1Options{})
if err != nil {
        // handle error
}
```

The `NewExampleServiceV1UsingExternalConfig` function will:
1. construct an authenticator using the environment variables above (an IAM authenticator using "my-iam-apikey" as the api key).
2. initialize the service client to use a base endpoint URL of "https://myservice.cloud.ibm.com/v2" rather than the default URL.

##### Storing configuration properties in a file
Instead of exporting your configuration properties as environment variables, you can store the properties
in a *credentials* file.   Here is an example of a credentials file that contains the properties from the example above:

```
# Contents of "example-service.env"
EXAMPLE_SERVICE_URL=https://myservice.cloud.ibm.com/v1
EXAMPLE_SERVICE_AUTH_TYPE=iam
EXAMPLE_SERVICE_APIKEY=my-iam-apikey

```

You would then provide the name of the credentials file via the `IBM_CREDENTIALS_FILE` environment variable:

```
export IBM_CREDENTIALS_FILE=/myfolder/example-service.env
```

When the SDK needs to look for configuration properties, it will detect the `IBM_CREDENTIALS_FILE` environment
variable, then load the properties from the specified file.

##### Complete configuration-loading process
The above examples provide a glimpse of two specific ways to provide external configuration to the SDK
(environment variables and credentials file specified via the `IBM_CREDENTIALS_FILE` environment variable).

The complete configuration-loading process supported by the SDK is as follows:
1. Look for a credentials file whose name is specified by the `IBM_CREDENTIALS_FILE` environment variable
2. Look for a credentials file at `<current-working-director>/ibm-credentials.env`
3. Look for a credentials file at `<user-home-directory>/ibm-credentials.env`
4. Look for environment variables whose names start with `<upper-case-service-name>_` (e.g. `RESOURCE_CONTROLLER_`)

At each of the above steps, if one or more configuration properties are found for the specified service,
those properties are then returned to the SDK and any subsequent steps are bypassed.

### Authentication
IBM Cloud "MySDK" Services use token-based Identity and Access Management (IAM) authentication.

IAM authentication uses an API key to obtain an access token, which is then used to authenticate
each API request.  Access tokens are valid for a limited amount of time and must be refreshed or reacquired.

To provide credentials to the SDK, you can do one of the following:
1. Construct or configure an `IamAuthenticator` instance with your IAM api key - in this case,
the SDK's IamAuthenticator implementation will use your API key to obtain an access token, ensure that it is valid,
and will then include the access token in each outgoing request, refreshing it as needed.

2. Construct or configure a `BearerTokenAuthenticator` instance using an access token that you obtain yourself -
in this case, you are responsible for obtaining the access token and refreshing it as needed.

For more details about authentication, including the full set of authentication schemes supported by
the underlying Go Core library, see
[Authentication](https://github.com/IBM/go-sdk-core/blob/master/Authentication.md)

#### Example: construct IamAuthenticator with an IAM api key

```go
// Letting the SDK manage the IAM access token
import {
    "github.com/IBM/go-sdk-core/v3/core"
    "github.ibm.com/CloudEngineering/go-sdk-template/exampleservicev1"
}

...

// Create the IamAuthenticator instance.
authenticator := &core.IamAuthenticator{
    ApiKey: "myapikey",
}

// Create the service options struct.
options := &exampleservicev1.ExampleServiceV1Options{
    Authenticator: authenticator,
}

// Construct the service instance.
myservice, err := exampleservicev1.NewExampleServiceV1(options)

```

#### Example: construct BearerTokenAuthenticator with an access token

```go
// Manage the IAM access token within the application
import {
    "github.com/IBM/go-sdk-core/v3/core"
    "github.ibm.com/CloudEngineering/go-sdk-template/exampleservicev1"
}

...

// Create the BearerTokenAuthenticator instance.
authenticator := &core.BearerTokenAuthenticator{
    BearerToken: "my IAM access token",
}

// Create the service options struct.
options := &exampleservicev1.ExampleServiceV1Options{
    Authenticator: authenticator,
}

// Construct the service instance.
myservice, err := exampleservicev1.NewExampleServiceV1(options)

...

// Later when the access token expires, the application must refresh the access token,
// then set the new access token on the authenticator.
// Subsequent request invocations will include the new access token.
authenticator.BearerToken = /* new access token */
```

### Passing operation parameters via an "options" struct
For each operation belonging to a service, an "options" struct is defined as a container for
the parameters associated with the operation.
The name of the struct will be `<operation-name>Options` and it will contain a field for each
operation parameter.  
Here's an example of an options struct for the `GetResource` operation:

```go
// GetResourceOptions : The GetResource options.
type GetResourceOptions struct {
	// The id of the resource to retrieve.
	ResourceID *string `json:"resource_id" validate:"required"`

	...
}
```

In this example, the `GetResource` operation has one parameter - `ResourceID`.
When invoking this operation, the application first creates an instance of the `GetResourceOptions`
struct and then sets the parameter value within it.  Along with the "options" struct, a constructor
function is also provided.  
Here's an example:

```go
options := myservice.NewGetResourceOptions("resource-id-1")
```

Then the operation can be called like this:

```go
result, detailedResponse, err := myservice.GetResource(options)
```

The use of the "options" struct pattern (instead of listing each operation parameter within the
argument list of the service methods) allows for future expansion of the API (within certain
guidelines) without impacting applications.

### Receiving operation responses

Each operation will return the following values:
1. `result` - An operation-specific response object (if the operation is defined as returning a response object).
2. `detailedResponse` - An instance of the `core.DetailedResponse` struct. This will contain the following fields:  
  * `StatusCode` - the HTTP status code returned in the response message
  * `Headers` - the HTTP headers returned in the response message
  * `Result` - the operation result (if available). This is the same value returned in the `result` return value mentioned above.
3. 'err' - An error object.  This return value will be nil if the operation was successful, or non-nil
if unsuccessful.

##### Examples
1. Here's an example of calling the `GetResource` operation which returns an instance of the `Resource`
struct as its result:

```go
// Construct the service instance
myservice := ... 

// Call the GetResource operation and receive the returned Resource.
options := myservice.NewGetResourceOptions("resource-id-1")
result, detailedResponse, err := myservice.GetResource(options)

// Now use 'result' which should be an instance of 'Resource'.
```

2. Here's an example of calling the `DeleteResource` operation which does not return a response object:

```
// Construct the service instance
myservice := ...

// Call the DeleteResource operation.
options := service.DeleteResourceOptions("resource-id-1")
detailedResponse, err := service.DeleteResource(options)
```

### Error Handling

In the case of an error response from the server endpoint, the Platform Services Go SDK will do the following:
1. The service method (operation) will return a non-nil `error` object.  This `error` object will
contain the error message retrieved from the HTTP response if possible, or a generic error message
otherwise.
2. The `detailedResponse.Result` field will contain the unmarshalled response (in the form of a
`map[string]interface{}`) if the operation returned a JSON response.  
This allows the application to examine all of the error information returned in the HTTP
response message.
3. The `detailedResponse.RawResult` field will contain the raw response body as a `[]byte` if the
operation returned a non-JSON response.

Here's an example of checking the `error` object after invoking the `GetResource` operation:

```go
// Call the GetResource operation and receive the returned Resource.
options := myservice.NewGetResourceOptions("bad-resource-id")
result, detailedResponse, err := myservice.GetResource(options)
if err != nil {
    fmt.Println("Error retrieving the resource: ", err.Error())
    fmt.Println("   full error response: ", detailedResponse.Result)
}
```

### Sending HTTP headers

#### Sending HTTP headers with all requests
A set of default HTTP headers can be included with all requests by using the `SetDefaultHeaders(http.Header)`
method of the service client.

Here's an example that includes `Custom-Header` with each request invocation:

```go
// Construct the service instance.
myservice, err := exampleservicev1.ExampleServiceV1(
    &exampleservicev1.ExampleServiceV1Options{
        Authenticator: authenticator,
    })

customHeaders := http.Header{}
customHeaders.Add("Custom-Header", "custom_value")
myservice.Service.SetDefaultHeaders(customHeaders)

// "Custom-Header" will now be included with all subsequent requests invoked from "myservice".
```

#### Sending request HTTP headers
Custom HTTP headers can also be passed with any individual request.
Just add the custom headers to the "options" struct prior to calling the operation.

Here's an example that includes `Custom-Header` along with the `GetResourceInstance` operation invocation:

```go

// Call the GetResource operation, passing our Custom-Header.
options := myservice.NewGetResourceOptions("resource-id-1")
customHeaders := make(map[string]interface{})
customHeaders["Custom-Header"] = "custom_value"
options.SetHeaders(customHeaders)
result, detailedResponse, err := myservice.GetResource(options)
// "Custom-Header" will be sent along with the "GetResource" request.
```

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at [dW Answers](https://developer.ibm.com/answers/questions/ask/?topics=ibm-cloud) or
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing
See [CONTRIBUTING](https://github.ibm.com/ibmcloud/platform-services-go-sdk/blob/master/CONTRIBUTING.md).

## License

The IBM Cloud "MySDK" Go SDK is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](https://github.ibm.com/CloudEngineering/go-sdk-template/blob/master/LICENSE).
