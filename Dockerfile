FROM golang:1.18.0-alpine AS builder

COPY . /src
WORKDIR /src

RUN go build -ldflags "-s -w" -o dashboard .

FROM alpine

COPY --builder /src/dashboard /app/dashboard

WORKDIR /app

ENTRYPOINT 80

ENTRYPOINT ["./dashboard"]