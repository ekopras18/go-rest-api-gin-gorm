# Stage 1: Build
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build main.go dari cmd/api/
RUN go build -o main ./cmd/api

# Stage 2: Run
FROM alpine:latest

# Certificate untuk HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /apps/

COPY --from=builder /app/main .
COPY .env .env

EXPOSE 8080

CMD ["./main"]
