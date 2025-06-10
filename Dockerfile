# Build stage
FROM --platform=$TARGETPLATFORM golang:1.23.2 as builder

WORKDIR /app

# Enable CGO and set Linux as target OS
ENV CGO_ENABLED=0 GOOS=linux

# Install only necessary build dependencies
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Copy go mod files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -ldflags="-s -w" -o main .

# Production stage
FROM --platform=$TARGETPLATFORM alpine:latest

WORKDIR /app

# Install CA certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy only necessary files from builder
COPY --from=builder /app/docs /app/docs
COPY --from=builder /app/main /app/main

# Create a non-root user
RUN adduser -D appuser
USER appuser

# Expose the port the app runs on
EXPOSE 8080

CMD ["/app/main"]