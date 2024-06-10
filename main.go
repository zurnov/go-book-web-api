package uni_book_web_api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

package main

import (
"encoding/json"
"fmt"
"net/http"

"github.com/google/uuid"
"github.com/gorilla/mux"
)

type Book struct {
	ID                string  `json:"id"`
	Isbn              string  `json:"isbn"`
	Title             string  `json:"title"`
	Author            *Author `json:"author"`
	YearOfPublication int     `json:"yearOfPublication"`
}

type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var library []Book

func main() {
	library = []Book{
		{ID: uuid.New().String(), Isbn: "12345", Title: "Book One", Author: &Author{FirstName: "John", LastName: "Doe"}, YearOfPublication: 2001},
		{ID: uuid.New().String(), Isbn: "67890", Title: "Book Two", Author: &Author{FirstName: "Steve", LastName: "Smith"}, YearOfPublication: 2005},
		{ID: uuid.New().String(), Isbn: "13579", Title: "Book Three", Author: &Author{FirstName: "Jane", LastName: "Doe"}, YearOfPublication: 2010},
		{ID: uuid.New().String(), Isbn: "4214", Title: "Book Four", Author: &Author{FirstName: "Janny", LastName: "Dash"}, YearOfPublication: 2010},
	}

	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Server is running successfully on port 8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		return
	}
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(library)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, book := range library {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Book not found"})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBook Book
	json.NewDecoder(r.Body).Decode(&newBook)
	newBook.ID = uuid.New().String()
	library = append(library, newBook)
	json.NewEncoder(w).Encode(newBook)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var updatedBook Book
	json.NewDecoder(r.Body).Decode(&updatedBook)
	for i, book := range library {
		if book.ID == params["id"] {
			updatedBook.ID = book.ID
			library[i] = updatedBook
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Book not found"})
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, book := range library {
		if book.ID == params["id"] {
			library = append(library[:i], library[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Book not found"})
}

