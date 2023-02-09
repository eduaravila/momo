FROM golang:1.19-alpine as base

LABEL maintainer="EduardoAvila <me@eduaravila.com>"

COPY go.* /github.com/eduaravila/momo/
COPY packages/router /github.com/eduaravila/momo/packages/router
COPY packages/db /github.com/eduaravila/momo/packages/db

COPY apps/auth /github.com/eduaravila/momo/apps/auth

WORKDIR /github.com/eduaravila/momo/apps/auth

RUN go install github.com/cosmtrek/air@latest

RUN go mod download

FROM base as dev


FROM base as prod

ENV ENV=production

# create dir for the binary
WORKDIR /github.com/eduaravila/momo/apps/auth/cmd/server

RUN  go build -o .

# execute the binary
CMD [ "./server" ]