# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN go build -o ./image-service ./pkg/image_service/main

EXPOSE 5050
CMD [ "./image-service" ]