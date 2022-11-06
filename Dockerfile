FROM golang:latest

WORKDIR /backend

COPY ./ /backend

RUN go mod download

ENTRYPOINT go run cmd/main.go