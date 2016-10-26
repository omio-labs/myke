FROM golang:1.7

WORKDIR /go/src/myke
COPY Godeps /go/src/myke/Godeps
RUN go get github.com/tools/godep && \
    godep restore && \
    godep get github.com/onsi/ginkgo/ginkgo

COPY . /go/src/myke
RUN ginkgo -r -v --trace --keepGoing && \
		gofmt -l .

CMD ["go", "build"]
