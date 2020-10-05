# syntax=docker/dockerfile:experimental

# Build Image
ARG GO_VERSION=1.14
FROM golang:${GO_VERSION}-alpine AS builder
WORKDIR /go/src/github.com/eahrend/iot-auth-orchestrate/device
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app .

# Application layer
FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates
RUN apk add bash
RUN mkdir /app
COPY --from=builder /app /app
WORKDIR /app
CMD ["./app"]
