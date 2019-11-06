FROM golang:1.11

WORKDIR /go/src/github.com/omio-labs/myke
COPY . /go/src/github.com/omio-labs/myke
CMD ["bin/cross-compile.sh"]
