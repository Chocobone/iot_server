# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Verify dependencies
RUN go mod verify

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o iot_server ./cmd/...

# Final stage
FROM alpine:latest

WORKDIR /app

# Install necessary runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Copy the binary from builder
COPY --from=builder /app/iot_server .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./iot_server"]

# Development stage
FROM golang:1.21-alpine AS dev

WORKDIR /app

# Install air for hot reload
RUN go install github.com/cosmtrek/air@latest

# Copy go mod and sum files
COPY go.mod go.sum ./

# Verify dependencies
RUN go mod verify

# Copy the source code
COPY . .

# Expose the port
EXPOSE 8080

# Command to run air for development
CMD ["air", "-c", ".air.toml"]