FROM golang:1.7

RUN go get github.com/tools/godep && \
		go get -v github.com/onsi/ginkgo/ginkgo && \
    mkdir -p /go/src/myke
CMD ["go", "run"]

WORKDIR /go/src/myke
COPY . /go/src/myke
RUN godep restore && \
		ginkgo -r && \
		go-wrapper install
