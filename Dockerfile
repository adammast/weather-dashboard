# Use an official lightweight Go image as the base image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application code into the container
COPY . .

# Build the application
RUN go build -o weather-dashboard

# Command to run the application
CMD ["./weather-dashboard"]
