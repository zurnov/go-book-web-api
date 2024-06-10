# Book API

This is a simple Web API for Books implemented in Golang. It provides RESTful endpoints for CRUD operations on book
resources.

## Features

- **Create, Read, Update, Delete (CRUD)**: Perform basic CRUD operations on book resources.

## Usage

- **Get All Books**: `GET /api/books`
- **Get a Single Book**: `GET /api/books/{id}`
- **Create a Book**: `POST /api/books`
- **Update a Book**: `PUT /api/books/{id}`
- **Delete a Book**: `DELETE /api/books/{id}`

## Technologies Used

- **Golang**: Backend implementation.
- **Gorilla Mux**: Routing and HTTP request handling.
- **UUID**: Generating unique identifiers.

## How to Run

1. Ensure Go v1.22.2 is installed.
2. Clone the repository:
    ```sh
    git clone https://github.com/zurnov/go-book-web-api
    ```
3. Install dependencies:
    ```sh
    go mod download
    ```
4. Start the server:
    ```sh
    go run main.go
    ```
6. The server will be running on port 8000.

## Endpoints

### GET /api/books

Returns a collection of all books recorded in the system.

### GET /api/books/{id}

Returns data for a single book by the given identifier.

### POST /api/books

Creates a new book.

### PUT /api/books/{id}

Updates the information for an already added book in the database.

### DELETE /api/books/{id}

Deletes a book by identifier.

