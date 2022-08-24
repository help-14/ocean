## Build

FROM golang:1.18-alpine AS build

RUN mkdir /app
WORKDIR /app
COPY ./ /app

RUN go mod download
RUN go build -o /bin

## Deploy

FROM alpine:latest
MAINTAINER Help-14 [mail@help14.com]
LABEL maintainer="mail@help14.com"

RUN mkdir /app

COPY --from=build /bin/ocean /app/
COPY ./sample-config.yaml /config.yaml

EXPOSE 8000

WORKDIR /app
ENTRYPOINT ./ocean