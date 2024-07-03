# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

# Set the working directory inside the builder container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go application
RUN go build -o ./bin/readme-merge

# Stage 2: Create a minimal image with the Go binary
FROM alpine:3.20.1

# Install Git
RUN apk add --no-cache git

# Set the working directory inside the final container
WORKDIR /app

# Create the bin directory in the final container
RUN mkdir -p /app/bin

# Copy the Go binary and entrypoint script from the builder container
COPY --from=builder /app/bin/* /app/bin/

# Ensure the entrypoint script is executable
RUN chmod +x /app/bin/entrypoint.sh

# Command to run your application
ENTRYPOINT ["/app/bin/entrypoint.sh"]
