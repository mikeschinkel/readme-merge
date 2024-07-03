# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

# Set the working directory inside the builder container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go application
RUN go build -o ./readme-merge

# Stage 2: Create a minimal image with the Go binary
FROM alpine:3.20.1

# Set the working directory inside the final container
WORKDIR /app

# Copy the Go binary from the builder container
COPY --from=builder /app/readme-merge .

# Command to run your application
ENTRYPOINT ["./entrypoint.sh"]

