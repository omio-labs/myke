FROM golang:1.7

WORKDIR /go/src/myke
COPY Godeps /go/src/myke/Godeps
RUN go get github.com/tools/godep && \
		go get github.com/mitchellh/gox && \
    godep restore && \
    godep get github.com/onsi/ginkgo/ginkgo

COPY . /go/src/myke
RUN ginkgo -r -v --trace --keepGoing && \
		gofmt -l .

CMD ["./cross-compile.sh"]
