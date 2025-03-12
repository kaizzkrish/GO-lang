# Step 1: Build the Go binary
FROM golang:1.23 AS builder
# Set the working directory inside the container to /app
WORKDIR /app
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o go-tutorial .

# Step 2: Create a minimal image from scratch
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/go-tutorial .
COPY ./views ./views

# Command to run the application
CMD ["/go-tutorial"]


# Instructions:
# 1. Ensure you have Docker installed on your machine.
# 2. Navigate to the directory containing the Dockerfile:
#    cd /Users/krishnakumarchandran/development/go-tutorial
# 3. Build the Docker image using the following command:
#    docker build -t go-tutorial .
# 4. Once the image is built, you can run the container using:
#    docker run --rm go-tutorial
# 5. The application should now be running inside the container.
# 6. Alternatively, if you have a docker-compose.yml file, you can use Docker Compose to run the application:
#    docker-compose up