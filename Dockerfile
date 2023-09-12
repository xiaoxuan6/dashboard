FROM golang:1.20.4-alpine3.18 AS builder

COPY . /src
WORKDIR /src

RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct

RUN go build -ldflags "-s -w" -o dashboard .

FROM alpine

COPY --from=builder /src/dashboard /app/dashboard
COPY --from=builder /src/static /app/static
COPY --from=builder /src/index.html /app/index.html

WORKDIR /app

ENTRYPOINT 80

ENTRYPOINT ["./dashboard"]