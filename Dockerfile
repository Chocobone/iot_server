# Build stage
FROM golang:1.24.3 AS builder

WORKDIR /app

# Copy go mod and sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .
# Binary build
RUN go build -buildvcs=false -trimpath  -ldflags "-w -s" -o iot_server

# Final stage
FROM alpine:latest AS deploy

RUN apk update

COPY --from=builder /app/iot_server .

CMD ["./iot_server"]

# Development stage using Air
FROM golang:1.24.3 AS dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["air"]