# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app
COPY .env .
# Download and install any required Go dependencies
RUN go mod download

# Run the command to start the application
CMD ["go", "run", "main.go"]