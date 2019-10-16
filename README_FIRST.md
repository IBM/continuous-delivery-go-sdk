# IBM Cloud Go SDK Template
This repository serves as a template for Go SDKs that are produced with the
[IBM OpenAPI SDK Generator](https://github.ibm.com/CloudEngineering/openapi-sdkgen).

You can use the contents of this repository to create your own Go SDKs.

## How to use this repository
##### 1. Copy or clone the repository
Copy the files contained in this repository as a starting point when building your own Go SDK
for one or more IBM Cloud services.
You can copy the repository using one of the following methods:
1. Download the zip file from the git repository main page
(click on the `Clone or download` button, then click on `Download ZIP`).
After downloading the zip file, unzip it into a suitable location where you want the project to exist.
The unzip command will create the `go-sdk-template-master` directory, which you might want to rename 
to something that reflects the Go SDK project that you're trying to build.
2. Clone this git repository using a command like this:
```
git clone git@github.ibm.com:CloudEngineering/go-sdk-template.git my-sdk
```
where `my-sdk` is the name of the directory to clone the repository into.  If you omit that option, git
will clone the repo into the `go-sdk-template` directory.  In that case, you might want to rename the 
directory to something that reflects the Go SDK project that you're trying to build.

Note: If you use the `git clone` method of copying the repository, be sure to change the URL
associated with the `origin` remote, like this:
```
git remote set-url origin <your project's git URL>
```
If you don't do this, you might inadvertently try to push your changes back to the `go-sdk-template` repository.

##### 2. Sanity-check your copy of the `go-sdk-template` repository
After copying or cloning this repository, you can do a quick sanity check by running this command in
the project root directory:
```
go test ./...
```
You should see output like this:
```
$ go test ./...
go: finding github.com/IBM/go-sdk-core v1.0.0
go: finding github.com/go-playground/locales v0.12.1
go: finding github.com/stretchr/testify v1.4.0
.
.
.
ok  	github.ibm.com/CloudEngineering/go-sdk-template/common	0.002s
ok  	github.ibm.com/CloudEngineering/go-sdk-template/exampleservicev1	0.006s
```

Note: the first time you build and test the project, you'll see output showing
that the Go engine is downloading the dependencies needed by the project since
they're not yet cached in your environment.

Note: this project uses go "modules" for dependency management.
For this reason, make sure the `GOPATH` environment variable is not set in
your shell when executing the `go` commands above.

##### 3. Modify your copy of the repository to reflect your SDK project
Once you have verified that your copy of this template repository builds and tests cleanly, it's time to
transform it into your new Go SDK project.
Make sure you have done the following:
1. Rename the root directory to reflect your Go SDK project (e.g. `mv go-sdk-template my-go-sdk`)
2. If you previously cloned the `go-sdk-template` git repository, be sure to change the url associated
with the remote named `origin`, like this: `git remote set-url origin <your project's git URL>`

Next, make modifications to various files as instructed below:
* __README.md__ - This file is intended to be customized to form the `README.md` file for your new
Go SDK project, so modify it as needed to reflect your project.  
* __.travis.yml__ - This file contains a basic set of commands to automate Travis builds for your Go SDK
project.  Modify this file as needed to incorporate any other build steps needed by your project.
* __common/headers.go__ - Go SDKs built with the IBM OpenAPI SDK Generator
need to implement a package called "common" which contains a function called `GetSdkHeaders`.  
The `GetSdkHeaders` function is invoked by the generated service methods and should be modified to 
suit the needs of your particular SDK. The default implementation of `GetSdkHeaders` will return a map
containing the "User-Agent" header with a value of the form 
`go-sdk-template-0.0.1 (arch=x86_64; os=Linux; go.version=1.12.9)`.
You can modify this function to customize the value of the `User-Agent` header or add additional 
headers to the map that is returned.  The headers returned in the map will be added to each
outgoing request invoked by applications using your SDK.
After modifying the `common/headers.go` file, be sure to also update the accompanying testcase
(`common/headers_test.go`) to properly test the new version of `headers.go`.
* __common/version.go__ - this file contains a single constant named `Version` which represents the
version of the SDK project.  The initial value is `0.0.1` which implies a pre-release version.
Be sure to update this value as appropriate to reflect the correct version of your Go SDK.
We highly recommend the use of [semantic versioning](https://semver.org/), which works nicely with the
Go engine's module and dependency management.

##### 4. Generate the Go code with the IBM OpenAPI SDK Generator
This is the step that you've been waiting for!

In this step, you'll invoke the IBM OpenAPI SDK Generator to process your API definition(s).

###### Generator setup
1. Install an official release of the OpenAPI SDK Generator.
Details are [here](https://github.ibm.com/CloudEngineering/openapi-sdkgen/blob/master/README.md#using-a-pre-built-installer).
You might want to also add the installation directory to your shell PATH environment variable.
2. Determine the correct API package prefix to use for your Go SDK project.  This value will also be
used as the Go module prefix in step 6 below.
This is typically a string of the form `github.com/my-org/my-sdk` or `github.ibm.com/my-org/my-sdk`,
depending on the github server where your Go SDK project will be located.
Suppose your SDK project (named "cloud-go-sdk") is going to be housed in the IBM internal github 
server under the "ibmcloud" github organization.  In that case, the API package prefix would be `github.ibm.com/ibmcloud/cloud-go-sdk`.
3. Modify your API definition(s) to include the API package prefix.  Details on the `apiPackage` 
configuration property can be found [here](https://github.ibm.com/CloudEngineering/openapi-sdkgen/wiki/Config-Options).  
Here's an example of an API definition that has this property defined:
```
openapi: "3.0.0"
info:
  version: 1.0.0
  title: Example service
  x-alternate-name: ExampleService
  license:
    name: MIT
  x-codegen-config:
    go:
      apiPackage: 'github.ibm.com/ibmcloud/cloud-go-sdk'
```
###### Generating the code for your service(s)
For each service that you want to include in your Go SDK project, process the 
service's API definition with the SDK Generator, like this:
```
openapi-sdkgen.sh generate -i <API-definition-filename> -g watson-go -o <output-directory>
```
For the output directory, you can specify the root directory of your Go SDK project, and the generator
will create a directory underneath that to represent the package associated with the generated service,
then it will write the generated source files to that package directory.

##### 5. Remove `exampleservicev1` package
The `go-sdk-template` repository includes an example service in the `exampleservicev1` package.
This is an example of a service that was generated with the IBM OpenAPI SDK Generator.
Once you have generated the Go code for your own service(s) and added those packages to your SDK
project, you'll want to remove the `exampleservice1` package so that it is not included in your SDK.
To remove the package, simply remove the `exampleservicev1` directory and the files contained within it:
```
rm -fr exampleservicev1
```
##### 6. Update the dependencies specified in the go.mod file
The `go-sdk-template` repository uses Go modules to manage dependencies.
[For more information on Go modules, see [this](https://github.com/golang/go/wiki/Modules)].

The `go.mod` file supplied with this repository reflects the dependencies associated with the actual
code delivered with the repository (i.e. the "common" and "exampleservicev1" packages).
After you have generated the Go code for your own service(s) and added those packages to the project
(and removed the `exampleservicev1` package), you'll need to update the `go.mod` file to reflect 
your own project's module prefix and the dependencies of the new code in your project.
The easiest way to do this is to simply remove the existing `go.mod` and `go.sum` files, 
and then use the Go engine to re-create them:
```
rm go.mod
rm go.sum
go mod init github.com/my-org/my-sdk-repo
go mod tidy
```
In the `go mod init...` command above, be sure to use the correct module prefix for your own Go SDK 
project.
##### 7. Build and test the project
If you made it this far, congratulate yourself!

After modifying the template repository to form your new Go SDK project and then generating the Go
code for your service(s) and adding the resulting package(s) to your project, it's time to build
and test your project.

The OpenAPI SDK Generator will generate unit tests for your service(s) in addition to the client 
SDK code, so you should have generated test cases in each of your service package(s).

To build and test all of the code within your project, you can run these commands in the project
root directory:
```
go build ./...
go test ./...
```
Technically, you only need to run the second command because the `go test` command will build
all the code (non-test and test code) as needed before running the tests.
If everything builds and tests cleanly, you should see output like this:
```
$ go test ./...
ok  	github.ibm.com/CloudEngineering/go-sdk-template/common	0.002s
ok  	github.ibm.com/CloudEngineering/go-sdk-template/exampleservicev1	0.006s
```
Note: The above output reflects the module prefix for the `go-sdk-template` repository and the
example service that is shipped with it.  Your output should reflect your Go SDK project's
module prefix and your project's set of packages.

If you encounter compile issues with either the client SDK or test code generated by the SDK Generator,
please let us know by posting on the `#wcp-sdk-generation` slack channel or by opening an issue
in the [`github.ibm.com/arf/arf-planning-sdk`](https://github.ibm.com/arf/planning-sdk-squad/)
issue repository.

Our goal is to generate the client SDK and test code that can be built and tested without manual intervention.  If we fall short of that goal, we'd love to hear about it.
