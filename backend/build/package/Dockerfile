FROM --platform=linux/x86_64 golang:1.23.3

RUN mkdir /go/src/boons-drone.com/chatting -p

WORKDIR /go/src/boons-drone.com/chatting

ADD . /go/src/boons-drone.com/chatting

RUN go install golang.org/x/tools/cmd/goimports@latest
RUN go install github.com/air-verse/air@latest
