# Use the official Golang image as the parent image
FROM golang:1.16

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go app
RUN go build -o main .

# Expose port 8080 for the container
EXPOSE 3000

# Define the command to run the executable
CMD ["./main"]
