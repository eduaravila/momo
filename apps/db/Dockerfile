# Use the official PostgreSQL image as the base image
FROM postgres:15.1-alpine

LABEL maintainer="EduardoAvila <me@eduaravila.com>"

# Set the working directory
WORKDIR /usr/src/app


ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_USER=postgres
ENV POSTGRES_DB=postgres

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Apply all available migrations.
RUN  goose postgres "user=postgres password=postgres dbname=postgres sslmode=disable" up

