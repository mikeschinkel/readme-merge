# Use the official Golang image as a base
FROM golang:1.22-alpine3.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o /app/readme-merge

# Command to run your application
ENTRYPOINT ["/app/readme-merge"]
