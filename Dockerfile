FROM golang:1.7

RUN go get github.com/tools/godep && \
    mkdir -p /go/src/app
CMD ["go-wrapper", "run"]

WORKDIR /go/src/app
COPY . /go/src/app
RUN go-wrapper download && \
		go test && \
		go-wrapper install
