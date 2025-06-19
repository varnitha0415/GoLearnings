package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/varnitha0415/GoLearnings/books_api_crud/config"
	"github.com/varnitha0415/GoLearnings/books_api_crud/handlers"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbClient *mongo.Client

// ConnectToDB initializes the database connection
func ConnectToDB() (*mongo.Client, error) {
	client, err := config.ConnectToMongoDB()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func main() {
	bookHandler := &handlers.BookHandlerImpl{}
	var err error
	dbClient, err = ConnectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		bookHandler.GetAllBooks(w, r, dbClient)
	})
	http.HandleFunc("/books/add", func(w http.ResponseWriter, r *http.Request) {
		bookHandler.AddBook(w, r, dbClient)
	})
	http.HandleFunc("/books/update", func(w http.ResponseWriter, r *http.Request) {
		bookHandler.UpdateBook(w, r, dbClient)
	})
	http.HandleFunc("/books/delete", func(w http.ResponseWriter, r *http.Request) {
		bookHandler.DeleteBook(w, r, dbClient)
	})

	fmt.Println("Starting server on :8080")
}
