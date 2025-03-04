# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy module dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main cmd/main.go

# Runtime stage
FROM alpine:3.17

WORKDIR /root/

# Copy the compiled binary
COPY --from=builder /app/main .

# Copy the .env file
COPY .env .env

# Expose application port
EXPOSE 8080 

# Run the application
CMD ["./main"]
