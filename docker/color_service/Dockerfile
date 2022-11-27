# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN go build -o ./image-service ./pkg/color_service/main

EXPOSE 6060
CMD [ "./color_service" ]