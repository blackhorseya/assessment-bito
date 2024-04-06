FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY adapter ./adapter
COPY pkg ./pkg
COPY entity ./entity
COPY app ./app

RUN go build -o ./bin/bito ./adapter

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/bito .

ENTRYPOINT ["./bito"]
