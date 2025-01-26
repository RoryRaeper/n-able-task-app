# Use the official golang image as the base image
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /n-able-task-app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o /n-able-task-app/main .

# Start a new stage from scratch
FROM golang:1.23
WORKDIR /n-able-task-app-app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /n-able-task-app/main .
EXPOSE 8080

# Command to run the executable
CMD ["./main"]