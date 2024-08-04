FROM golang:1.20-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Install make and any other dependencies, build the Go application
RUN apk update && apk add --no-cache make curl && \
    make build && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz && \
    ls -la /app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY ./internal/db/migrate ./migration
COPY app.env .
COPY start.sh .
COPY wait-for.sh .

# List the files to ensure they are correctly copied
RUN ls && chmod +x start.sh

# Make port 8080 available to the world outside this container
EXPOSE 8080

# Run the executable
ENTRYPOINT ["/app/start.sh"]
CMD ["/app/main"]
