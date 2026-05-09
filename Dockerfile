# syntax=docker/dockerfile:1

FROM golang:1.24-alpine AS builder
WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/api ./cmd/api

FROM alpine:3.20
WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /bin/api /usr/local/bin/api

EXPOSE 8080

CMD ["/usr/local/bin/api"]
