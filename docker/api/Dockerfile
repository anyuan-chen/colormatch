# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN go build -o ./api ./pkg/api/main

EXPOSE 8080
CMD [ "./api" ]