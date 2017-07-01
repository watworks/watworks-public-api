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

## Misc notes ##

An earlier version of the app used a config file, and a custom docker entrypoint script to generate the config from ENV variables when the file was not present.  Generally I'd like to avoid that approach, but if it's needed, see examples in this post: https://medium.com/@kelseyhightower/12-fractured-apps-1080c73d481c
