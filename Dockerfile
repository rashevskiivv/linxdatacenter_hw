FROM golang:latest

WORKDIR /app

COPY . /app

RUN go mod download

RUN go build -o ./.bin cmd/main.go

EXPOSE 8080

ENTRYPOINT ["./.bin"]