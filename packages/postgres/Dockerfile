FROM golang:1.19-alpine as base
LABEL maintainer="EduardoAvila <me@eduaravila.com>"

WORKDIR /app
COPY . .
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

