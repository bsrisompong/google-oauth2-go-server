# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/server/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY .env .env

# Ensure migration files are copied to the final image
COPY db/migrations /app/db/migrations


EXPOSE 8080

CMD ["./main"]
