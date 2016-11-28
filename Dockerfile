FROM golang:1.7

WORKDIR /go/src/github.com/goeuro/myke
COPY Godeps /go/src/github.com/goeuro/myke/Godeps
RUN go get -u -v github.com/tools/godep && \
		go get -u -v github.com/mitchellh/gox && \
		go get -u -v github.com/jteeuwen/go-bindata/... && \
    godep restore -v && \
    godep get -v github.com/onsi/ginkgo/ginkgo

COPY . /go/src/github.com/goeuro/myke
CMD ["bin/cross-compile.sh"]
