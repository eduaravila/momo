FROM golang:1.19-alpine as base

LABEL maintainer="EduardoAvila <me@eduaravila.com>"

WORKDIR /app/src/app

COPY . .
RUN go install github.com/cosmtrek/air@latest
RUN go mod download

FROM base as dev


FROM base as prod

ENV ENV=production

# create dir for the binary
RUN mkdir -p /app/bin
RUN  go build -o bin/app

# execute the binary
CMD [ "./app" ]