# Stage 1: Build the Go binary
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy the entire workspace root
COPY . .

# Build the API specifically
# Using go 1.24 features for workspace building
RUN go build -o /app/api ./apps/api/cmd

# Stage 2: Create the production image
FROM alpine:latest

# Install CA certificates for secure connections
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/api .

# Run the API
ENTRYPOINT ["./api"]