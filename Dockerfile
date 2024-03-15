# Use a specific version of golang alpine image for reproducible builds
FROM golang:1.21-alpine AS builder

# Set a working directory inside the container
WORKDIR /app

# Copy go mod and sum files to leverage Docker cache layering
COPY go.mod go.sum ./

# Download dependencies in advance; improves caching on rebuilds
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/backend-lxd

# Use a minimal alpine image for the runtime
FROM alpine:latest

# Add ca-certificates in case your application makes external HTTPS requests
RUN apk --no-cache add ca-certificates

# Set the working directory in the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/backend-lxd .

# Run the binary
CMD ["./backend-lxd"]
