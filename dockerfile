# Use Golang as the base image
FROM golang:1.21 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go files and dependencies
COPY . .

# Build the Go application
RUN go build -o my-go-app

# Create a minimal runtime image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the compiled binary from the builder
COPY --from=builder /app/my-go-app .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./my-go-app"]
