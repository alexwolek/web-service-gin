# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /web-service-gin

EXPOSE 8080

CMD [ "/web-service-gin" ]