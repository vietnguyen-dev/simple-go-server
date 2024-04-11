# Use the official Go image as a base image
FROM golang:1.22.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Start with an Ubuntu Linux-based image
FROM ubuntu:22.04

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage into this stage
COPY --from=builder /app/main .

# Copy the frontend directory
COPY frontend /app/frontend

# Allow the ain directory to be run
RUN chmod +x /app/main

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]


















