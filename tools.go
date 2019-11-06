// +build tools
// Declaring build-time dependencies, these are ignored at compile-time
// Refer https://github.com/golang/go/issues/25922

package main

import (
	_ "github.com/mitchellh/gox"
	_ "github.com/omeid/go-resources"
	_ "golang.org/x/lint/golint"
)
