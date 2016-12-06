FROM golang:1.7
MAINTAINER Andreev Vladislav <andreevlad@gmail.com>

ADD . /go/src/github.com/maddevsio/instagram-agent

WORKDIR /go/src/github.com/maddevsio/instagram-agent

RUN go get -v && go build -v

CMD ["./instagram-agent"]
