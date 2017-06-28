# Watworks Public Api #

[![Build Status](https://travis-ci.org/watworks/watworks-public-api.svg?branch=master)](https://travis-ci.org/watworks/watworks-public-api)

> Serving DHTML since 1342!

This is the primary api backend for watworks.  It serves YOUR needs.  Whatever those are.

## Developing ##

You can run `make dev-up` to get started.

	make dev-up

There are some extra steps you'll need to take to enable code completion in your IDE.  You will need to have `go` installed on your host system, with `GOPATH` properly configured.  Once your `$GOPATH` is already set on your host system, run `make dev-init` to link this project into your gopath, and your IDE should be able to do code completion and import resolution properly.

To test and/or run the code, use the following:

* `make test` - runs all tests
* `make run` - compiles then runs the binary