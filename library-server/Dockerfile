# Use the official Golang image as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 3000

# Create a .env file with default values
RUN echo "DB_HOST=db" > .env && \
    echo "DB_PORT=5432" >> .env && \
    echo "DB_USER=postgres" >> .env && \
    echo "DB_PASSWORD=postgres" >> .env && \
    echo "DB_NAME=library_db" >> .env && \
    echo "PORT=3000" >> .env

# Command to run the executable
CMD ["./main"]
