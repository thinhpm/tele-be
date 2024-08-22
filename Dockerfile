# Stage 1: Build the Go binary
FROM golang:1.21-alpine AS builder

# Install Delve
# RUN go install github.com/go-delve/delve/cmd/dlv@latest


# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first, to cache dependency installation
COPY go.mod ./

RUN go mod tidy

# Download all dependencies; dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go binary, specifying the correct path to the main.go file
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/app/main.go

# Command to run Delve in headless mode, listening on all interfaces
# CMD ["dlv", "debug", "--headless", "--listen=:2345", "--api-version=2", "--accept-multiclient", "--log", "./cmd/app/main.go"]


# Stage 2: Create a minimal image with the Go binary
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .


# Expose the application's port
EXPOSE 8080

# Command to run the Go application
CMD ["./main"]


