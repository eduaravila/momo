FROM golang:1.20-alpine as base

LABEL maintainer="EduardoAvila <me@eduaravila.com>"

COPY go.* /github.com/eduaravila/momo/
COPY packages/router /github.com/eduaravila/momo/packages/router
COPY packages/postgres /github.com/eduaravila/momo/packages/postgres
COPY packages/server /github.com/eduaravila/momo/packages/server
COPY packages/decorators /github.com/eduaravila/momo/packages/decorators
COPY packages/metrics /github.com/eduaravila/momo/packages/metrics
COPY packages/genproto /github.com/eduaravila/momo/packages/genproto

COPY apps/auth /github.com/eduaravila/momo/apps/auth

WORKDIR /github.com/eduaravila/momo/apps/auth

RUN go mod download

# DEBUG MODE
FROM base as debug

RUN apk add build-base
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN ls || cat 

FROM base as dev

RUN go install github.com/cosmtrek/air@latest

FROM base as prod

ENV ENV=production
# create dir for the binary
WORKDIR /github.com/eduaravila/momo/apps/auth/cmd/http

RUN  go build -o .

# execute the binary
CMD [ "./server" ]