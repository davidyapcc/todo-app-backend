# Use the base image with Go installed
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Initialize Go module
RUN go mod init github.com/davidyapcc/todo-app-backend

# Install Goose
RUN go install github.com/pressly/goose/cmd/goose@latest

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the application runs on
EXPOSE 8080

# Command to run Goose migrations and the application
CMD ["sh", "-c", "goose -dir /app/migrations mysql 'admin:admin@tcp(mysql-db:3306)/todo' up && ./main"]