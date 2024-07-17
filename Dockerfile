# Use the official Go image as the build environment
FROM golang:1.23rc2-alpine3.20 as builder

# Set the working directory inside the container
WORKDIR /app

# Expose port 8081 for the application
EXPOSE 8080

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# # Copy the entire project to the working directory
COPY . .

# # Build the Go application and output the binary to /app
RUN go build -o ./bin/raven ./cmd/main.go

# # Use the official Alpine Linux image for the final container
FROM golang:1.23rc2-alpine3.20

# # Copy the built binary from the builder stage to /bin/app
COPY --from=builder /app/bin/raven /opt/raven

# # Set the command to run the application
CMD ["/opt/raven"]