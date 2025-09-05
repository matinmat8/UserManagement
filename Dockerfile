# ---- Build Stage ----
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Download dependencies first (better cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build
RUN go build -o app .

# ---- Run Stage ----
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/app .

# Use environment variables for Redis config
ENV REDIS_HOST=redis
ENV REDIS_PORT=6379

CMD ["./app"]
