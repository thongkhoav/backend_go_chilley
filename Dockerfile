# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod & download deps before to cache layer
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -o server main.go

# Run stage
FROM alpine:latest

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/server .

# Expose port
EXPOSE 8080

# Run app
CMD ["./server"]
