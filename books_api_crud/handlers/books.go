package handlers

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// AddBook adds a new book to the database
func AddBook(book http.ResponseWriter, r *http.Request) {
	client, err := config.ConnectToMongoDB()
	if err != nil {
		http.Error(book, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(r.Context())
	collection := client.Database("bookstore").Collection("books")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		http.Error(book, "Failed to retrieve books", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())
	var books []models.Book
	for cursor.Next(context.Background()) {
		var book models.Book
		if err := cursor.Decode(&book); err != nil {
			http.Error(book, "Failed to decode book", http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}
}
