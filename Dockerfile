FROM golang:1.20-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Install make and any other dependencies
RUN apk update && apk add --no-cache make

# Build the Go application
RUN make build

RUN ls -la /app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY app.env .

# Make port 8080 available to the world outside this container
EXPOSE 8080

# Run the executable
CMD ["/app/main"]
