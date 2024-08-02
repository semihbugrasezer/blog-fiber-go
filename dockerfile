FROM golang:1.19-alpine

# Install gcc and musl-dev for cgo
RUN apk update && apk add --no-cache gcc musl-dev

# Set working directory
WORKDIR /app

# Copy source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Ensure the binary is executable
RUN chmod +x ./main

# List files and permissions for debugging
RUN ls -al

# Define the default command to run the application
CMD ["./main"]
