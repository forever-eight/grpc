FROM golang:1.12

ADD . /go/src/github.com/forever-eight/grpc/server

RUN go get github.com/forever-eight/grpc/server

ENTRYPOINT ["/go/bin/server"]

EXPOSE 5300