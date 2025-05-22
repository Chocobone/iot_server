FROM golang:1.24-alpine AS deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -trimpath -ldflags "-w -s" -o cmd/main.go

FROM alpine:latest AS deploy

RUN apt-get update && apt-get upgrade -y

COPY --from=deploy-builder /app/cmd/main .

CMD ["./cmd/main"]

FROM golang:1.24-alpine AS dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["air"]