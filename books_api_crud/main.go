package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/varnitha0415/GoLearnings/books_api_crud/handlers"
)

func main() {
	http.HandleFunc("/books", handlers.GetAllBooks)
	http.HandleFunc("/books/add", handlers.AddBook)
	http.HandleFunc("/books/update", handlers.UpdateBook)
	http.HandleFunc("/books/delete", handlers.DeleteBook)
	http.HandleFunc("/books/id", handlers.GetBook)

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
