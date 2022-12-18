FROM golang:1.18.2 AS builder
MAINTAINER "Sigrid Jin"

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

# download the required Go dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN go build go-backend

EXPOSE 8080
EXPOSE 80

CMD ["go", "run", "main.go", "serve"]
