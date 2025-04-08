FROM golang:1.24-alpine

# Install Air (hot reload) & git (for modules)
RUN apk add --no-cache git
RUN go install github.com/air-verse/air@latest

WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
