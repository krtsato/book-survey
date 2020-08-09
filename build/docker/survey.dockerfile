FROM golang:1.14.7-alpine3.12

WORKDIR /book-survey

ENV GO111MODULE on

RUN set -ox pipefail \
  && apk update \
  && apk add --no-cache bash vim \
  && rm -rf /var/cache/apk/* \
  && go get -u github.com/go-sql-driver/mysql

COPY . .

# Multi-stage Build Example
# FROM golang:1.14.7 as builder
# ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
# WORKDIR /book-servey
# COPY . .
# RUN go build -o cmd/survey cmd/main.go
#
# FROM alpine:3.12
# WORKDIR /book-survey
# COPY --from=builder /book-survey/cmd/survey ./cmd/survey
