FROM golang:latest

WORKDIR /go/src/

ENV GO111MODULE=on

RUN mkdir -p /go/src/greeting

ADD ./src /go/src/greeting

EXPOSE 9000

CMD [ "go", "run", "/go/src/greeting/greeting.go"]
