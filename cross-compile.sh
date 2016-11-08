#!/usr/bin/env sh

gox \
	-osarch="darwin/amd64 linux/amd64 windows/amd64" \
	-output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}"
