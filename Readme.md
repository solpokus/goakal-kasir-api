# Kasir API

A simple RESTful API for a Point of Sale (POS) system, built with Go (Golang).

## Overview

Kasir API provides endpoints to manage categories and verify the system health. It uses an **in-memory database** for storage, making it lightweight and easy to run without external dependencies. It also includes **Swagger** documentation for easy API exploration.

## Features

- **Category Management (CRUD)**: Create, Read, Update, and Delete categories.
- **In-Memory Storage**: Thread-safe storage using Go maps and mutexes.
- **Health Check**: Endpoint to monitor server status.
- **Swagger UI**: Interactive API documentation.

## Prerequisites

- **Go 1.22** or higher.

## Getting Started

### 1. Clone the Repository

```bash
git clone <repository-url>
cd goakal-kasir-api
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Run the Application

```bash
go run main.go
```

The server will start on port `8080`.

## API Documentation

### Swagger UI

Access the interactive API documentation at:

[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

### Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/categories` | Get all categories |
| `POST` | `/categories` | Create a new category |
| `GET` | `/categories/{id}` | Get a specific category by ID |
| `PUT` | `/categories/{id}` | Update a category by ID |
| `DELETE` | `/categories/{id}` | Delete a category by ID |
| `GET` | `/health` | Check server health status |

### Example Usage (cURL)

**Create a Category:**
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"Beverages", "description":"Drinks"}' http://localhost:8080/categories
```

**Get All Categories:**
```bash
curl http://localhost:8080/categories
```

## Project Structure

```
goakal-kasir-api/
├── docs/           # Swagger documentation files
├── handler/        # HTTP Handlers
├── model/          # Data structures
├── repository/     # Data access layer (In-Memory)
├── main.go         # Application entry point
├── go.mod          # Go module file
└── Readme.md       # Project documentation
```

## License

Apache 2.0
