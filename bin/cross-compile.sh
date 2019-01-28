#!/usr/bin/env bash
#Usage: build.sh <version>
set -e

# Prepare tmp folder
rm -rf tmp
mkdir -p tmp

# Insert version
version=${1?"version is required"}
echo $version >> "tmp/version"

# Generate license notices
go mod vendor
deps="github.com/goeuro/myke $(go list -f '{{ join .Deps "\n"}}' ./... | grep -v 'goeuro/myke')"
out="tmp/LICENSES"
echo -e "OPEN SOURCE LICENSES\n" > $out

for dep in $deps; do
	if [ -d "vendor/$dep" ]; then
		notices=$(ls -d vendor/$dep/* 2>/dev/null | grep -i -e "license" -e "licence" -e "copying" -e "notice" || echo)
		if [ ! -z "$notices" ]; then
			echo -e "BEGIN LICENSE FOR $dep\n\n" >> $out
			for notice in $notices; do
				echo "Adding license: $notice"
				cat $notice >> $out
			done
			echo -e "\nEND LICENSE FOR $dep\n\n" >> $out
		fi
	fi
done

# Compile resources
go run github.com/omeid/go-resources/cmd/resources \
	-declare -var=FS -output core/bindata.go -package core tmp/*

# Cross compile
export CGO_ENABLED=0
go run github.com/mitchellh/gox \
	-osarch="darwin/amd64 linux/amd64 windows/amd64" \
	-ldflags="-s -w" \
	-output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}"
