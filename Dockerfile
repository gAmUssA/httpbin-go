# Build stage
FROM golang:1.17 AS build-env

WORKDIR /app

# Copy all the files from current directory to Docker environment
COPY . .

# Enable Go modules
ENV GO111MODULE=on

# Install dependencies
RUN go mod download

# Build the Go app binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o httpbin .

# Final stage
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=build-env /app/httpbin .

# Expose port 8080 for the app
EXPOSE 8080

ENV GIN_MODE=release

# Command to run the binary
CMD ["./httpbin"]
