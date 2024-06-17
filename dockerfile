# Use an official Golang image as a build stage
FROM golang:1.18 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-simple-bank ./cmd/.

# Use a minimal image for the final stage
FROM alpine:latest

# Copy the built binary from the builder stage
COPY --from=builder /docker-simple-bank /docker-simple-bank

# Expose the port your application runs on
EXPOSE 8080

# Command to run the binary
CMD ["/docker-simple-bank"]
