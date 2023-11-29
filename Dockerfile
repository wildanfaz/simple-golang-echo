# === Stage 1: Build Stage ===
FROM golang:1.18-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o simple-golang-echo .

# === Stage 2: Final Stage ===
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the build stage to the final stage
COPY --from=builder /app/simple-golang-echo .

# Expose the port the app runs on
EXPOSE 1323

# Command to run the executable
CMD ["./simple-golang-echo", "start"]
