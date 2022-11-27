# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o ./session_management_service ./pkg/session_management_service/main

EXPOSE 4040
CMD [ "./session_management_service" ]