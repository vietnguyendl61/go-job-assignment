# Start from the official golang image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /go/src/app

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Expose port to the outside world
EXPOSE 4000
EXPOSE 10000

# Command to run the executable
CMD ["go", "run", "main.go"]
