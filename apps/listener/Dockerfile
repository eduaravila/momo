# Choose whatever you want, version >= 1.16
FROM golang:1.19-alpine as base

LABEL author="EduardoAvila <me@eduaravila.com>"

WORKDIR /app/src/app


RUN go install github.com/cosmtrek/air@latest
COPY go.mod go.sum ./
RUN go mod download



FROM base as builder

WORKDIR /app/src/app

# create dir for the binary
RUN mkdir -p /app/bin
RUN  go build -o bin/app main .

# execute the binary
CMD [ "./main" ]