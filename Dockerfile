FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build API
RUN go build -o /app/api ./cmd/api

# Build Worker
RUN go build -o /app/worker ./cmd/worker

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binaries from builder
COPY --from=builder /app/api /app/api
COPY --from=builder /app/worker /app/worker
COPY .env.example .env

# Expose port for API
EXPOSE 8000

# Default command (can be overridden in docker-compose)
CMD ["./api"]
