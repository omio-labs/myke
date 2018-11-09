FROM golang:1.11

WORKDIR /go/src/github.com/goeuro/myke
COPY . /go/src/github.com/goeuro/myke
CMD ["bin/cross-compile.sh"]
