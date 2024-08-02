# Use the official Golang image to create a build artifact.
FROM golang:1.22.5 AS builder

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Use a minimal Docker image for the final build
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder / .

# Copy views and static files
COPY --from=builder /views ./views
COPY --from=builder /static ./static

# Expose port
EXPOSE 8000

# Command to run the executable
CMD ["./main"]
