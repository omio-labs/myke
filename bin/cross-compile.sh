#!/usr/bin/env bash
set -ex

# Run tests
ginkgo -r -v --trace --keepGoing

# Generate license notices
deps="github.com/goeuro/myke $(go list -f '{{ join .Deps "\n"}}' . | grep -v 'goeuro/myke')"
rm -rf tmp
mkdir -p tmp
out="tmp/LICENSES"
echo -e "OPEN SOURCE LICENSES\n" > $out

for dep in $deps; do
	if [ -d "$GOPATH/src/$dep" ]; then
		notices=$(ls -d $GOPATH/src/$dep/* 2>/dev/null | grep -i -e "license" -e "licence" -e "copying" -e "notice")
		if [ ! -z "$notices" ]; then
			echo -e "$dep\n\n" >> $out
			for notice in $notices; do
				cat $notice >> $out
			done
			echo -e "\n\n" >> $out
		fi
	fi
done

# Compile bindata
go-bindata -o core/bindata.go -nometadata -pkg core tmp/

# Cross compile
gox \
	-osarch="darwin/amd64 linux/amd64 windows/amd64" \
	-output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}"
