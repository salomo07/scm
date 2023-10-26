# Stage 1: Build the Go application
FROM golang:1.21.3 as builder
WORKDIR /app
COPY . .
RUN go build -o myapp

# Stage 2: Create the final image with Redis
FROM redis:latest

# Set the Redis password
ENV REDIS_PASSWORD your-redis-password

# Copy the compiled Go application from the previous stage
COPY --from=builder /app/myapp /myapp

# Expose the port for your Go application
EXPOSE 8080

# Start Redis with the configured password
CMD ["redis-server", "--requirepass", "$REDIS_PASSWORD"]

# Start your Go application
CMD ["/myapp"]
