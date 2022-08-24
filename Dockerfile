## Build

FROM golang:1.18-alpine AS build

RUN mkdir /app
WORKDIR /app
COPY ./src /app

RUN apk update
RUN apk upgrade
RUN apk add gcc
RUN apk add g++

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go install github.com/mattn/go-sqlite3
RUN go build -o /bin

## Deploy

FROM alpine:latest
MAINTAINER Help-14 [mail@help14.com]
LABEL maintainer="mail@help14.com"

RUN mkdir /app

COPY --from=build /bin/ocean /app/

EXPOSE 8000

WORKDIR /app
ENTRYPOINT ./ocean