MySDK Go SDK

Go client library to use the MySDK Services.

<details>
<summary>Table of Contents</summary>

* [Overview](#overview)
* [Prerequisites](#prerequisites)
* [Installation](#installation)
* [Authentication](#authentication)
* [Sample Code](#sample-code)
* [License](#license)

</details>

## Overview

The IBM Cloud MySDK Go SDK allows developers to programmatically interact with the 
MySDK IBM Cloud services.

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration?target=%2Fdeveloper%2Fwatson&

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* An installation of Go (version 1.12 or above) on your local machine.

## Installation
There are a few different ways to download and install the MySDK Go SDK project for use by your
Go application:
##### 1. `go get` command
Use this command to download and install the MySDK Go SDK project to allow your Go application to 
use it:
```
go get -u github.com/my-org/my-sdk
```
##### 2. Go modules
If your application is using Go modules, you can add a suitable import to your
Go application, like this:
```go
import (
	"github.com/my-org/my-sdk/myservicev1"
)
```
then run `go mod tidy` to download and install the new dependency and update your Go application's 
`go.mod` file.
##### 2. `dep` dependency manager
If your application is using the `dep` dependency management tool, you can add a dependency
to your `Gopkg.toml` file.  Here is an example:
```
[[constraint]]
  name = "github.com/my-org/my-sdk/myservicev1"
  version = "0.0.1"

```
then run `dep ensure`.

## Authentication

MySDK services use token-based Identity and Access Management (IAM) authentication [IAM](#iam).

IAM authentication uses a service API key to obtain an access token that is used to authenticate
each API request.
Access tokens are valid for a limited amount of time and must be regenerated.

To provide credentials to the SDK, you supply either an IAM service **API key** or an **access token**:

- Use the API key to have the SDK manage the lifecycle of the access token. The SDK requests an access token, ensures that the access token is valid, and refreshes it if necessary.
- Use the access token if you want to manage the lifecycle yourself. For details, see [Authenticating with IAM tokens](https://cloud.ibm.com/docs/services/watson/getting-started-iam.html).


Supplying the IAM API key:

```go
// letting the SDK manage the IAM access token
import {
    "github.com/IBM/go-sdk-core/core"
    "github.com/my-org/my-sdk/myservicev1"
}
...
// Create the IAM authenticator.
authenticator := &core.IamAuthenticator{
    ApiKey: "myapikey",
}

// Create the service options struct.
options := &myservicev1.MyServiceV1Options{
    Authenticator: authenticator,
}

// Construct the service instance.
service, err := myservicev1.NewMyServiceV1(options)

```

Supplying the access token (a bearer token):

```go
import {
    "github.com/IBM/go-sdk-core/core"
    "github.com/my-org/my-sdk/myservicev1"
}
...
// Create the BearerToken authenticator.
authenticator := &core.BearerTokenAuthenticator{
    BearerToken: "my IAM access token",
}

// Create the service options struct.
options := &myservicev1.MyServiceV1Options{
    Authenticator: authenticator,
}

// Construct the service instance.
service, err := myservicev1.NewMyServiceV1(options)
```

## Using the SDK
### Passing operation parameters via an "options" struct
For each operation belonging to a service, an options struct is defined as a container for 
the parameters associated with the operation.
The name of the struct will be <operation-name>Options and it will contain a field for each 
operation parameter.  Here's an example:
```go
// GetResourceOptions : The GetResource options.
type GetResourceOptions struct {

	// The id of the resource to retrieve.
	ResourceID *string `json:"resource_id" validate:"required"`
}
```
In this example, the `GetResource` operation has one parameter, named `ResourceID`.
When invoking this operation, the application first creates an instance of the `GetResourceOptions`
struct and then sets the ResourceID field within it.  Along with the options struct, a constructor
function is also provided.
Here's an example:
```go
options := service.NewGetResourceOptions("resource-id-1")
```
Then the operation can be called like this:
```go
result, detailedResponse, err := service.GetResource(options)
```
This pattern allows for future expansion of the API (within certain guidelines) without impacting
applications.

### Receiving operation responses

Each service method (operation) will return the following values:
1. `result` - An operation-specific result (if the operation is defined as returning a result).
2. `detailedResponse` - An instance of the `core.DetailedResponse` struct.
This will contain response information such as:
* the HTTP status code returned in the response message
* the HTTP headers returned in the response message
* the operation result (if available). This is the same value returned in the `result` return value
mentioned above.
3. 'err' - An error object.  This return value will be non-nil if the operation was not successful,
or nil if it was successful.

##### Example:
1. Here's an example of calling the `GetResource` operation which returns an instance of the `Resource`
struct as its result:
```
// Construct the service instance.
service, err := myservicev1.NewMyServiceV1(
    &myservicev1.MyServiceV1Options{
        Authenticator: authenticator,
    })

// Call the GetResource operation and receive the returned Resource.
options := service.NewGetResourceOptions("resource-id-1")
result, detailedResponse, err := service.GetResource(options)

// Now use 'result' which should be an instance of 'Resource'.
```
2. Here's an example of calling the `DeleteResource` operation which does not return a response object:
```
// Construct the service instance.
service, err := myservicev1.NewMyServiceV1(
    &myservicev1.MyServiceV1Options{
        Authenticator: authenticator,
    })

// Call the DeleteResource operation and receive the returned Resource.
options := service.NewDeleteResourceOptions("resource-id-1")
detailedResponse, err := service.DeleteResource(options)
```

### Error Handling

In the case of an error response from the server endpoint:, the MySDK Go SDK will do the following:
1. The service method (operation) will return a non-nil `error` object.  This `error` object will
contain the error message retrieved from the HTTP response if possible, or a generic error message
otherwise.
2. The `detailedResponse.Result` field will contain the unmarshalled response (in the form of a
`map[string]interface{}`) if the operation returned a JSON response.  
This allows the application to examine all of the error information returned in the HTTP 
response message.
2. The `detailedResponse.RawResult` field will contain the raw response body as a `[]byte` if the
operation returned a non-JSON response.

##### Example:
Here's an example of checking the `error` object after invoking the `GetResource` operation
```go
// Call the GetResource operation and receive the returned Resource.
options := service.NewGetResourceOptions("bad-resource-id")
result, detailedResponse, err := service.GetResource(options)
if err != nil {
    fmt.Println("Error retrieving the resource: ", err.Error())
    fmt.Println("   full error response: ", detailedResponse.Result)
}
```

### Default headers

Default HTTP headers can be specified by using the `SetDefaultHeaders(http.Header)`
method of the client instance.  Once set on the client instance, default headers are sent with
every outbound request.  
##### Example:
The example below sets the header `Custom-Header` with the value "custom_value" as the default 
header:
```go
// Construct the service instance.
service, err := myservicev1.NewMyServiceV1(
    &myservicev1.MyServiceV1Options{
        Authenticator: authenticator,
    })

customHeaders := http.Header{}
customHeaders.Add("Custom-Header", "custom_value")
service.Service.SetDefaultHeaders(customHeaders)

// "Custom-Header" will now be included with all subsequent operation invocations.
```

### Sending request headers

Custom HTTP headers can also be passed with any request.
To do so, add the headers to the options object passed to the service method.
##### Example:
Here's an example that sets "Custom-Header" on the `GetResourceOptions` instance and then
invokes the `GetResource` operation:
```go

// Call the GetResource operation, passing our Custom-Header.
options := service.NewGetResourceOptions("resource-id-1")
customHeaders := make(map[string]interface{})
customHeaders["Custom-Header"] = "custom_value"
options.SetHeaders(customHeaders)
result, detailedResponse, err := service.GetResource(options)
```

## Sample Code

See [Samples](Samples).

## License

The IBM Cloud MySDK Go SDK is released under the Apache 2.0 license. The license's full text can be found in [LICENSE](LICENSE).
