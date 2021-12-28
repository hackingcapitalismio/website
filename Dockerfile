# syntax = docker/dockerfile:1.2

FROM golang:latest
WORKDIR /go/src/github.com/kris-nova/hcio
COPY app app
COPY main.go main.go
COPY go.mod go.sum ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /hcio

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY public /public
COPY --from=0 /hcio /hcio
CMD ["/hcio"]
