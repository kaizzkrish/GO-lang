# Step 1: Build the Go binary
FROM golang:1.23 AS builder

WORKDIR /app
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o go-tutorial .

# Step 2: Create a minimal image from scratch
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/go-tutorial .

# Command to run the application
CMD ["/go-tutorial"]