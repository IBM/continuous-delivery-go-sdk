MySDK Go SDK

Go client library to use the MySDK Services.

<details>
<summary>Table of Contents</summary>

* [Overview](#overview)
* [Prerequisites](#prerequisites)
* [Installation](#installation)
* [Using the SDK](#using-the-sdk)
  * [Constructing service clients](#constructing-service-clients)
  * [Authentication](#authentication)
  * [Passing operation parameters via an options struct](#passing-operation-parameters-via-an-options-struct)
  * [Receiving operation responses](#receiving-operation-responses)
  * [Error Handling](#error-handling)
  * [Default headers](#default-headers)
  * [Sending request headers](#sending-request-headers)
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
##### 3. `dep` dependency manager
If your application is using the `dep` dependency management tool, you can add a dependency
to your `Gopkg.toml` file.  Here is an example:
```
[[constraint]]
  name = "github.com/my-org/my-sdk/myservicev1"
  version = "0.0.1"

```
then run `dep ensure`.

## Using the SDK
This section provides general information on how to use the services contained in this SDK.
### Constructing service clients
Each service is implemented in its own package (e.g. `myservicev1`).
The package will contain a "service client"
struct (a client-side representation of the service), as well as an "options" struct that is used to
construct instances of the service client.  
Here's an example of how to construct an instance of "My Service":
```go
import {
    "github.com/IBM/go-sdk-core/core"
    "github.com/my-org/my-sdk/myservicev1"
}

// Create an authenticator.
authenticator := /* create an authenticator - see examples below */

// Create an instance of the "MyServiceV1Options"  struct.
myserviceURL := "https://myservice.cloud.ibm.com/api"
options := &myservicev1.MyServiceV1Options{
    Authenticator: authenticator,
    URL: myserviceURL,
}

// Create an instance of the "MyServiceV1" service client.
service, err := NewMyServiceV1(options)
if err != nil {
    // handle error
}

// Service operations can now be called using the "service" variable.

```

### Authentication
MySDK services use token-based Identity and Access Management (IAM) authentication.

IAM authentication uses an API key to obtain an access token, which is then used to authenticate
each API request.  Access tokens are valid for a limited amount of time and must be regenerated.

To provide credentials to the SDK, you supply either an IAM service **API key** or an **access token**:

- Specify the IAM API key to have the SDK manage the lifecycle of the access token.
The SDK requests an access token, ensures that the access token is valid, and refreshes it when
necessary.
- Specify the access token if you want to manage the lifecycle yourself.
For details, see [Authenticating with IAM tokens](https://cloud.ibm.com/docs/services/watson/getting-started-iam.html).

##### Examples:
* Supplying the IAM API key and letting the SDK manage the access token for you:

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

* Supplying the access token (a bearer token) and managing it yourself:

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

...
// Later when the access token expires, the application must refresh the access token,
// then set the new access token on the authenticator.
// Subsequent request invocations will include the new access token.
authenticator.BearerToken = /* new access token */
```
For more information on authentication, including the full set of authentication schemes supported by
the underlying Go Core library, see
[this page](https://github.com/IBM/go-sdk-core/blob/master/Authentication.md)

### Passing operation parameters via an options struct
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

    // The type of the resource to retrieve.
    ResourceType *string `json:"resource_type" validate:"required"`

    ...
}
```
In this example, the `GetResource` operation has two parameters - `ResourceID` and `ResourceType`.
When invoking this operation, the application first creates an instance of the `GetResourceOptions`
struct and then sets the parameter values within it.  Along with the "options" struct, a constructor
function is also provided.  
Here's an example:
```go
options := service.NewGetResourceOptions("resource-id-1", "resource-type-1")
```
Then the operation can be called like this:
```go
result, detailedResponse, err := service.GetResource(options)
```
This use of the "options" struct pattern (instead of listing each operation parameter within the
argument list of the service method) allows for future expansion of the API (within certain
guidelines) without impacting applications.

### Receiving operation responses

Each service method (operation) will return the following values:
1. `result` - An operation-specific result (if the operation is defined as returning a result).
2. `detailedResponse` - An instance of the `core.DetailedResponse` struct.
This will contain the following fields:
* `StatusCode` - the HTTP status code returned in the response message
* `Headers` - the HTTP headers returned in the response message
* `Result` - the operation result (if available). This is the same value returned in the `result` return value
mentioned above.
3. `err` - An error object.  This return value will be nil if the operation was successful, or non-nil
if unsuccessful.

##### Example:
1. Here's an example of calling the `GetResource` operation which returns an instance of the `Resource`
struct as its result:
```go
// Construct the service instance.
service, err := myservicev1.NewMyServiceV1(
    &myservicev1.MyServiceV1Options{
        Authenticator: authenticator,
    })

// Call the GetResource operation and receive the returned Resource.
options := service.NewGetResourceOptions("resource-id-1", "resource-type-1")
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

In the case of an error response from the server endpoint, the MySDK Go SDK will do the following:
1. The service method (operation) will return a non-nil `error` object.  This `error` object will
contain the error message retrieved from the HTTP response if possible, or a generic error message
otherwise.
2. The `detailedResponse.Result` field will contain the unmarshalled response (in the form of a
`map[string]interface{}`) if the operation returned a JSON response.  
This allows the application to examine all of the error information returned in the HTTP
response message.
3. The `detailedResponse.RawResult` field will contain the raw response body as a `[]byte` if the
operation returned a non-JSON response.

##### Example:
Here's an example of checking the `error` object after invoking the `GetResource` operation:
```go
// Call the GetResource operation and receive the returned Resource.
options := service.NewGetResourceOptions("bad-resource-id", "bad-resource-type")
result, detailedResponse, err := service.GetResource(options)
if err != nil {
    fmt.Println("Error retrieving the resource: ", err.Error())
    fmt.Println("   full error response: ", detailedResponse.Result)
}
```

### Default headers
Default HTTP headers can be specified by using the `SetDefaultHeaders(http.Header)`
method of the client instance.  Once set on the service client, default headers are sent with
every outbound request.  
##### Example:
The example below sets the header `Custom-Header` with the value "custom_value" as a default
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

// "Custom-Header" will now be included with all subsequent requests invoked from "service".
```

### Sending request headers
Custom HTTP headers can also be passed with any individual request.
To do so, add the headers to the "options" struct passed to the service method.
##### Example:
Here's an example that sets "Custom-Header" on the `GetResourceOptions` instance and then
invokes the `GetResource` operation:
```go

// Call the GetResource operation, passing our Custom-Header.
options := service.NewGetResourceOptions("resource-id-1", "resource-type-1")
customHeaders := make(map[string]interface{})
customHeaders["Custom-Header"] = "custom_value"
options.SetHeaders(customHeaders)
result, detailedResponse, err := service.GetResource(options)
// "Custom-Header" will be sent along with the "GetResource" request.
```

### Transaction IDs

Every call from the SDK will receive a response which will contain a transaction ID, accessible via the `x-global-transaction-id` header.  This transaction ID is useful for troubleshooting and accessing relevant logs from your service instance.

## License

The IBM Cloud MySDK Go SDK is released under the Apache 2.0 license. The license's full text can be found in [LICENSE](LICENSE).
