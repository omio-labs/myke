FROM golang:1.7

WORKDIR /go/src/myke
COPY Godeps /go/src/myke/Godeps
RUN go get -u -v github.com/tools/godep && \
		go get -u -v github.com/mitchellh/gox && \
		go get -u -v github.com/jteeuwen/go-bindata/... && \
    godep restore && \
    godep get -v github.com/onsi/ginkgo/ginkgo

COPY . /go/src/myke
RUN ginkgo -r -v --trace --keepGoing && \
		gofmt -l .

CMD ["bin/cross-compile.sh"]
