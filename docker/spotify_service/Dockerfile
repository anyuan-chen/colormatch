# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o ./spotify-service ./pkg/spotify_service/main

EXPOSE 3030
CMD [ "./spotify-service" ]