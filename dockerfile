FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/server/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

COPY .env .env

EXPOSE 8080

CMD ["./main"]
