FROM golang:1.7

RUN go get github.com/tools/godep && \
    mkdir -p /go/src/app
CMD ["go", "run"]

WORKDIR /go/src/app
COPY . /go/src/app
RUN go-wrapper download -t && \
		go test && \
		go-wrapper install
