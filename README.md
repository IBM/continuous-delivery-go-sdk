# IBM Cloud Go SDK Template
This repository serves as a template for Go SDKs that are produced with the 
[IBM OpenAPI SDK Generator](https://github.ibm.com/CloudEngineering/openapi-sdkgen).

You can use the contents of this repository to create your own Go SDKs.

## How to use this repository

##### 1. Copy the repository
Copy the files contained in this repository as a starting point when building your own Go SDK
for one or more IBM Cloud services.

##### 2. Modify the copied files to reflect your SDK
The following files will need to be modified after copying them from this template repository:
* .travis.yml - Update this file as needed to incorporate any required steps for your SDK


* headers.go - Go SDKs built with the IBM OpenAPI SDK Generator 
need to implement a package called "common" which contains a function called `GetSdkHeaders`.  
The `GetSdkHeaders` function is invoked by the generated service methods and should be modified to suit the
needs of your particular SDK.

##### 3. Generate the Go code with the IBM OpenAPI SDK Generator
This is the step that you've been waiting for!

In this step, you'll invoke the IBM OpenAPI SDK Generator to process your API definition.

This will generate a collection of Go source files which will be included in your SDK project.
You'll find instructions on how to do this [here](https://github.ibm.com/CloudEngineering/openapi-sdkgen/wiki).