# base image
FROM golang:1.16-alpine

WORKDIR /app

# Copy all the application files to the working directory
COPY . .  

 # Download the Go module dependencies
RUN go mod download 

# Document that the container listens on port 8080
EXPOSE 8080  

# Specify the command to run the Go application
CMD ["go", "run", "main.go"] 


