FROM golang:1.9.2

RUN echo "deb http://ftp.debian.org/debian jessie-backports main" > /etc/apt/sources.list.d/backports.list
RUN apt-get update && apt-get install -y protobuf-compiler protobuf-c-compiler capnproto

ENV GOPATH /go
WORKDIR    /go

RUN go get -u github.com/golang/protobuf/...
RUN go get -u zombiezen.com/go/capnproto2/...
RUN mv /go/bin/capnpc-go /go/bin/capnpc-go2
RUN go get -u github.com/glycerine/go-capnproto/...
RUN go get -u github.com/tinylib/msgp/...

ADD .   /proto
WORKDIR /proto

CMD bash
