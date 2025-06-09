# Use official Go image
FROM golang:1.22 as builder

WORKDIR /app

# Copy go mod and download deps
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy rest of the code
COPY . .

# Build binary
RUN go build -o weather-api ./cmd/main.go

# Final lightweight image
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/weather-api .

# Copy .env file
COPY .env .

EXPOSE 8080

CMD ["./weather-api"]
