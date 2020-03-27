# IBM Cloud Go SDK Template Usage Instructions

This repository serves as a template for Go SDKs that are produced with the
[IBM OpenAPI SDK Generator](https://github.ibm.com/CloudEngineering/openapi-sdkgen).

You can use the contents of this repository to create your own Go SDK repository.

## How to use this repository

### 1. Create your new github repository from this template
This SDK template repository is implemented as a
[github template](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template),
which makes it easy to create new projects from it.

To create a new SDK repository from this template, follow these instructions:  
1. In your browser, open the link for this
[template repository](https://github.ibm.com/CloudEngineering/go-sdk-template).

2. Click on the `Use this template` button that appears next to the `Clone or download` button.

3. In the next window:  
- Select the `Owner`. This is the github id or organization where the new repository should be created
- Enter the respository name (e.g. `platform-services-go-sdk`):  
  - Recommendation: use a name of the form `<service-category>-<language>-sdk`, where:  
    - `<service-category>` refers to the IBM Cloud service category associated with the services that
	  will be included in the project (e.g. `platform-services`)
    - `<language>` is the language associated with the SDK project (e.g. `go`)
	
4. Click the `Create repository from template` button to create the new repository  

If your goal is to create the new SDK repository on the `Github Enterprise` server (github.ibm.com),
then you are finished creating the new repository and you can proceed to section 2.

On the other hand, if your goal is to create the new SDK repository on the `Public Github` server (github.com),
then perform these additional steps:

5. Create a new **EMPTY** repository on the Public Github server:  
- Select "No template" for the "Repository template" option
- Select the `Owner` (your personal id or an organization)
- Enter the same respository name that you used when creating the new repository above (e.g. my-go-sdk)
- Do NOT select the `Initialize this repository with a README` option
- Select `None` for the `Add .gitignore` and `Add a license` options
- Click the `Create repository` button.
- After the new empty repository has been created, you will be at the main page
of your new repository, which will include this text:
```
...or push an existing repository from the command line

git remote add origin git@github.com:padamstx/my-go-sdk.git
git push -u origin master
```
- Take note of the two git commands listed above for your new repository, as we'll execute these later

6. Clone your new `Github Enterprise` repository (created in steps 1-3 above)
to your local development environment:  

```sh
[/work/demos]
$ git clone git@github.ibm.com:phil-adams/my-go-sdk.git
Cloning into 'my-go-sdk'...
remote: Enumerating objects: 36, done.
remote: Counting objects: 100% (36/36), done.
remote: Compressing objects: 100% (32/32), done.
remote: Total 36 (delta 1), reused 0 (delta 0), pack-reused 0
Receiving objects: 100% (36/36), 28.74 KiB | 577.00 KiB/s, done.
Resolving deltas: 100% (1/1), done.
```

7. "cd" into your project's root directory:

```sh
[/work/demos]
$ cd my-go-sdk
[/work/demos/my-go-sdk]
$ 
```

8. Remove the existing remote:  
```sh
[/work/demos/my-go-sdk]
$ git remote remove origin
```

9. Add a new remote which reflects your new `Public Github` repository:

```sh
[/work/demos/my-go-sdk]
$ git remote add origin git@github.com:padamstx/my-go-sdk.git
```

10. Push your local repository to the new remote (Public Github):  

```sh
[/work/demos/my-go-sdk]
$ git push -u origin master
Enumerating objects: 36, done.
Counting objects: 100% (36/36), done.
Delta compression using up to 12 threads
Compressing objects: 100% (31/31), done.
Writing objects: 100% (36/36), 28.74 KiB | 28.74 MiB/s, done.
Total 36 (delta 1), reused 36 (delta 1)
remote: Resolving deltas: 100% (1/1), done.
To github.com:padamstx/my-go-sdk.git
 * [new branch]      master -> master
Branch 'master' set up to track remote branch 'master' from 'origin'.
```

You have now created your new SDK repository on the `Public Github` server.

You may want to now delete the new SDK repository that you created on the `Github Enterprise`
server since it will no longer be used now that you have created your repository on `Public Github`.


### 2. Sanity-check your new repository

After creating your new SDK repository from the template repository, and cloning it
into your local development environment, you can do a quick sanity check by
running this command in the project root directory:
```
go test ./...
```
You should see output like this:
```
$ go test ./...
go: finding github.com/IBM/go-sdk-core/v3 v3.2.4
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


### 3. Modify selected files

- In this section, you'll modify various files within your new SDK repository to reflect
the proper names and settings for your specific project.

- The template repository comes with an example service included, but this should be removed
from your project.  Remove the following directory and its contents:
  - exampleservicev1

- Next, here is a list of the various files within the project with comments
that will guide you in the required modifications:

  - `common/headers.go`:
    - modify the `sdkName` constant to reflect your project name (e.g. `platform-services-go-sdk`)
    - read the comments in the `GetSdkHeaders()` function

  - `common/version.go`:
    - make sure the `Version` constant is set to "0.0.1", as this will be the starting version
      number (release) of the project.

  - `go.mod`/`go.sum`:
    - Remove the `go.mod` and `go.sum` files
    - Run this command to create a new `go.mod` file which will contain your project's
      github url as the module import path:
      ```sh
         go mod init <module-import-path>
      ```
      where `<module-import-path>` should be the correct module import path for your project.
      This will be the github repository URL without the `https` scheme
      (e.g. `github.ibm.com/ibmcloud/platform-services-go-sdk`).

  - `README.md`:
    - Change the title to reflect your project; leave the version in the title as `0.0.1`
    - Change the `cloud.ibm.com/apidocs` link to reflect the correct service category
      (e.g. `platform-services`)
    - In the Overview section, modify `IBM Cloud MySDK Python SDK` to reflect your project
      (e.g. `IBM Cloud Platform Services Python SDK`)
    - In the table of services, remove the entry for the example service; later you'll list each
      service contained in your SDK project in this table, along with a link to the online reference docs
      and the name of the generated service struct.
    - In the Installation section, update the examples to reflect your new
      project's module import path (e.g. `github.ibm.com/ibmcloud/platform-services-go-sdk`).
    - In the "Issues" section, modify `<github-repo-url>` to reflect the Github URL for your project.
    - Note that the README.md file contains a link to a common README document where general
      SDK usage information can be found.
    - When finished read through the document and make any other changes that might be necessary.

  - `CONTRIBUTING.md`:
    - In the "Issues" section, modify `<github-repo-url>` to reflect the Github URL for your project.

At this point, it's probably a good idea to commit the changes that you have made so far.
Be sure to use proper commit messages when committing changes (follow the link in `CONTRIBUTING.md`
to the common CONTRIBUTING document).  
Example:
```sh
cd <project-root>
git add .
git commit -m "chore: initial SDK project setup"
```


##### 4. Generate the Go code with the IBM OpenAPI SDK Generator
This is the step that you've been waiting for!

In this step, you'll invoke the IBM OpenAPI SDK Generator to process your API definition(s).

###### Generator setup
1. Install an official release of the OpenAPI SDK Generator.
Details are [here](https://github.ibm.com/CloudEngineering/openapi-sdkgen/blob/master/README.md#using-a-pre-built-installer).
You might want to also add the installation directory to your shell PATH environment variable.
2. Determine the correct API package prefix to use for your Go SDK project.  This value will also be
used as the Go module prefix in step 6 below.
This is typically a string of the form `github.com/my-org/my-sdk` or `github.ibm.com/my-org/my-sdk` (the url of your SDK repository),
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

##### 8. Running Example Service Integration Test

To set up and run the integration tests, clone the [Example Service repo](https://github.ibm.com/CloudEngineering/example-service) and follow the instructions there for how to start up an instance of the example service

Integration test code can be found [here](exampleservicev1/example_service_v1_integration_test.go)

To run only the integration tests, run this command from the project root directory:
```sh
go test exampleservicev1/example_service_v1_integration_test.go
```

##### 9. Continuous Integration
This repository is set up to use [Travis](https://travis-ci.org/) for continuous integration.

Note - to pass credentials to Travis and run integration tests, create a file named `ibm-credentials.env` at the root of the project directory, then encrypt the `ibm-credentials.env` file with the Travis CLI to store the decryption keys in Travis setting. For more information on the format of the `ibm-credentials.env` file, see (example credentials file here)[https://github.com/IBM/go-sdk-core/blob/master/resources/ibm-credentials.env]. To encrypt `ibm-credentials.env` and set decryption keys on Travis config:

1. Enable Travis-CI for your repository in Travis.
2. Make sure Ruby and Ruby Gem are installed and up to date on your local machine. You can [install Ruby here](https://www.ruby-lang.org/en/documentation/installation/)
3. Install Travis CLI (`gem install travis`). To verify installation, type `travis -v`
4. Log into Travis through CLI. Depending on whether you're trying to connect to Travis Enterprise, or Public Travis, the commands will be different.

Here's the command for logging into Travis Enterprise:
```sh
travis login -X --github-token <your-github-enterprise-token> --api-endpoint https://travis.ibm.com/api
```

Here's the command for logging into Public Travis
```sh
travis login --github-token <your-public-github-token> --com
```

5. From the root of the project, run the command `travis encrypt-file ibm-credentials.env`
6. The command will generate a file called `ibm-credentials.env.enc` in the project folder root directory. Commit the file to your repository
7. Terminal should print out a command to add to your build script. In that command is a string with the format similar to `encrypted_12345_key`. Copy that string
8. Open `.travis.yml` from root directory. Replace the string `encrypted_12345_key` with the name of your generated environment variable from the last step
9. Also replace the string `encrypted_12345_iv` with the name of your generated environment variable, but modify the string from `_key` to `_iv`
10. Commit the changes you made to the `.travis.yml` file and push to Github. Travis-CI pipeline should automatically start running

The config file `.travis.yml` contains all the instructions necessary to run the recommended build. Each step is described below.

The `before_install` step runs the instructions to decrypt the `ibm-credentials.env.enc` file. It only does for *pushes* to a branch. This is done so that integration tests only run on *push* builds and not on *pull request* builds.

The `script` section runs the generated unit tests for the generated SDK. It will also run the command to lint go files. To run integration tests, first tag integration tests with the tag `integration` by adding the line `// +build integration` to the [top of the integration test file](exampleservicev1/example_service_v1_integration_test.go). The, uncomment the line ```go test `go list ./...` -tags integration``` to run both unit tests and integration tests in `.travis.yml`.

## Setting the ``User-Agent`` Header In Preparation for SDK Metrics Gathering

If you plan to gather metrics for your SDK, the `User-Agent` header value must be
a string similar to the following:
   `my-go-sdk/0.0.1 (lang=go; arch=x86_64; os=Linux; go.version=1.12.9)`

The key parts are the sdk name (`my-go-sdk`), version (`0.0.1`) and the
language name (`lang=go`).
This is required because the analytics data collector uses the User-Agent header included
with each request to gather usage data for IBM Cloud services.

The default implementation of the `common.GetSDKHeaders()` method provided in this SDK template
repository will need to be modified slightly for your SDK.
Replace the `my-go-sdk/0.0.1` part with the name and version of your
Go SDK. The rest of the system information should remain as-is.

For example, suppose your Go SDK project is called `platform-services-go-sdk` and its
version is `2.3.1`.
The `User-Agent` header value should be:
   `platform-services-go-sdk/2.3.1 (lang=go; arch=x86_64; os=Linux; go.version=1.12.9)`

__Note__: It is very important that the sdk name ends with the string `-sdk`,
as the analytics data collector uses this to gather usage data.

More information about the analytics tool, and other steps you should take to start gathering
metrics for your SDK can be found [here](https://github.ibm.com/CloudEngineering/sdk-analytics).
