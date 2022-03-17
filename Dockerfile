# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18-buster AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /web-service-gin

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /web-service-gin /web-service-gin

ENV GIN_MODE=release
ENV PORT=8080
EXPOSE 80
EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/web-service-gin"]