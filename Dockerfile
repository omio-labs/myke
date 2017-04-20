FROM golang:1.8

WORKDIR /go/src/github.com/goeuro/myke
COPY Godeps Godeps
COPY bin/init.sh bin/
RUN bin/init.sh

COPY . /go/src/github.com/goeuro/myke
CMD ["bin/cross-compile.sh"]
