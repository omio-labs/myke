# Contributing

We love pull requests from everyone. By participating in this project, you
agree to abide by the [code of conduct](https://github.com/omio-labs/myke/blob/master/CODE_OF_CONDUCT.md).

## Contribution Guide

* Make sure there are no other open issues or PRs
* Fork, then clone the repo
* Setup your machine
  * Install docker and docker-compose
  * Run `docker-compose build` to build the app
  * Run `docker-compose run --rm default /bin/bash` for a terminal inside the container
  * Make sure existing tests pass: `go test ./...`
* Make your changes and add test cases
  * We prefer testing-by-example, every package inside `examples/` folder has its own `package_test.go` in testing table style
* Run new tests again: `go test ./...`
* Run and test the program: `go run main.go`
* Push to your fork and [submit a pull request](https://github.com/omio-labs/myke/compare)
* Make sure tests pass in the CI
* Make sure you sign the [Developer Certificate of Origin](https://developercertificate.org/)
