# Go Application with Simple UI and MongoDB Integration

This is a simple Go application that provides a user interface (UI) for entering a person's name and age. The entered data is then saved to a MongoDB database. The application is packaged with Docker and includes the MongoExpress UI for easy management of the MongoDB database.

## Prerequisites

Before running the application, ensure that you have the following dependencies installed:

- Docker
- Docker Compose

## Getting Started

1. Clone the repository to your local machine
2. Start the application using 
```shell 
   docker-compose up
```

This command will build the Docker image, start the Go application, MongoDB, and MongoExpress. Wait for the containers to be fully initialized. \
\
3. Access the application:

The Go application with the simple UI can be accessed at: http://localhost:8080.
The MongoExpress UI for managing the MongoDB database can be accessed at: http://localhost:8081.

## Usage
Open your web browser and go to http://localhost:8080.

In the provided form, enter the person's name and age.

Click the "Submit" button to save the entered data to the MongoDB database.

Success! The person's data has been saved to the database.

## Cleaning Up
To stop and remove the Docker containers created for this application, run the following command in the project directory:
```
docker-compose down
```

## Blog Post

You can find the detailed blog post [here](https://kamnagarg-10157.medium.com/understanding-mutex-in-go-5f41199085b9).
