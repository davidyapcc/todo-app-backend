# ToDo API

## Overview

This API application is designed to handle the CRUD functionalities of the ToDo app.

## Prerequisites

To run this application locally, ensure you have the following installed:

- Go (version >=1.21.4)
- MySQL (version >=8.0.32) or Docker for containerized MySQL

Install Dependencies

    go mod tidy

Run the Application

    go run main.go

Database Migrations

    goose -dir /path/to/migrations mysql "user:password@tcp(host:port)/your_database" up

Make sure to replace placeholders (user, password, host, port, your_database) with your MySQL credentials and database information.

The application should start and be accessible at http://localhost:8080.