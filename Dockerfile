# Use the official Golang image as the base image
FROM golang:1.21.6-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Copy .env file to working directory
RUN ls -la
COPY .env .

# Build the Go application
RUN go build -o league-api ./cmd/league-api

# Expose the port on which the application will run
EXPOSE 3000

# Run the Go application
CMD ["./league-api"]

