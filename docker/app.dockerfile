FROM golang:alpine

WORKDIR /fiberboilerplate

ADD . .

RUN go mod download

ENTRYPOINT go build  && ./fiberboilerplate