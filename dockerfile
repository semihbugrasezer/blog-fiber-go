# Use the official Golang image to create a build artifact.
FROM golang:1.22.5 AS builder

# Set the working directory inside the builder container
WORKDIR /app

# Copy go mod and sum files to the container
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the application binary
RUN go build -o main .

# Use a minimal Docker image for the final build
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Set the working directory inside the final container
WORKDIR /root/

# Copy the binary from the builder stage to the final container
COPY --from=builder /app/main .

# Copy views and static files from the builder stage to the final container
COPY --from=builder /app/views ./views
COPY --from=builder /app/static ./static

# Expose the port that the application will run on
EXPOSE 8000

# Command to run the executable when the container starts
CMD ["./main"]
