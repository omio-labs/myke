#!/usr/bin/env bash
set -xe
go get -u -v github.com/golang/lint/golint
go get -u -v github.com/tools/godep
go get -u -v github.com/mitchellh/gox
go get -u -v github.com/jteeuwen/go-bindata/...
godep restore -v
