# Issues

If you encounter an issue with the project, you are welcome to submit a [bug report](<github-repo-url>/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

# Pull Requests

If you want to contribute to the repository, here's a quick guide:
  1. Fork the repository
  2. Develop and test your code changes:
      * To build/test: `go test ./...`   
      * Please add one or more tests to validate your changes.
  3. Make sure everything builds/tests cleanly
  4. Commit your changes  
  5. Push to your fork and submit a pull request to the **master** branch
  
# Coding Style

This SDK follows the conventions from the [Golang "Effective Go" style guide](). 

You should use [golint](https://github.com/golang/lint#installation) to check style and code coverage. `golint ./...` will perform the style checks on your changes.

# Running the Tests

Out of the box, `go test ./...` runs unit tests and integration tests (which require credentials).
To run only the unit tests (sufficient for most cases), use `go test -run 'Unit'`.

To run the integration tests, you need to provide credentials to the integration test framework.
The integration test framework will skip integration tests for any service that does not have credentials,

To provide credentials for the integration tests, copy `test/resources/.config.properties` to `test/resources/auth.js`
and fill in credentials for the service(s) you wish to test.

To run the tests in a specific test class, use the `-run` flag when invoking `mvn test`, e.g.:

```
go test -run NameOfTest
```

You can run a specific test by adding the name of the file, e.g.:

```
go test my_test.go
```

# Code Coverage

This repo uses [Cover](https://golang.org/cmd/cover/) to measure code coverage. To obtain a code coverage report, run `go test -coverprofile=cover.out` from the root of the project.

# Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
   have the right to submit it under the open source license
   indicated in the file; or

(b) The contribution is based upon previous work that, to the best
   of my knowledge, is covered under an appropriate open source
   license and I have the right under that license to submit that
   work with modifications, whether created in whole or in part
   by me, under the same open source license (unless I am
   permitted to submit under a different license), as indicated
   in the file; or

(c) The contribution was provided directly to me by some other
   person who certified (a), (b) or (c) and I have not modified
   it.

(d) I understand and agree that this project and the contribution
   are public and that a record of the contribution (including all
   personal information I submit with it, including my sign-off) is
   maintained indefinitely and may be redistributed consistent with
   this project or the open source license(s) involved.

## Additional Resources
+ [General GitHub documentation](https://help.github.com/)
+ [GitHub pull request documentation](https://help.github.com/send-pull-requests/)

[Maven]: https://maven.apache.org/guides/getting-started/maven-in-five-minutes.html
