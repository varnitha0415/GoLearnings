package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/varnitha0415/GoLearnings/books_api_crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookHandlerImpl struct{}

// GetAllBooks retrieves all books from the database
func (h *BookHandlerImpl) GetAllBooks(w http.ResponseWriter, r *http.Request, dbClient *mongo.Client) {
	collection := dbClient.Database("bookstore").Collection("books")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		http.Error(w, "Failed to retrieve books", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())
	var books []models.Book
	for cursor.Next(context.Background()) {
		var book models.Book
		if err := cursor.Decode(&book); err != nil {
			http.Error(w, "Failed to decode book", http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// AddBook adds a new book to the database
func (h *BookHandlerImpl) AddBook(w http.ResponseWriter, r *http.Request, dbClient *mongo.Client) {

	collection := dbClient.Database("bookstore").Collection("books")

	// Decode the request body into a Book struct
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Insert the book into the collection
	book.ID = primitive.NewObjectID() // Generate a new ObjectID
	_, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		http.Error(w, "Failed to insert book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// GetBook retrieves a book by its ID
func (h *BookHandlerImpl) GetBook(w http.ResponseWriter, r *http.Request, dbClient *mongo.Client) {
	collection := dbClient.Database("bookstore").Collection("books")

	idParam := r.URL.Query().Get("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book models.Book
	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// UpdateBook updates an existing book in the database
func (h *BookHandlerImpl) UpdateBook(w http.ResponseWriter, r *http.Request, dbClient *mongo.Client) {

	collection := dbClient.Database("bookstore").Collection("books")

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	idParam := r.URL.Query().Get("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	update := bson.M{
		"$set": bson.M{
			"title":        book.Title,
			"author":       book.Author,
			"published_at": book.PublishedAt,
		},
	}

	result, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *BookHandlerImpl) DeleteBook(w http.ResponseWriter, r *http.Request, dBclient *mongo.Client) {
	collection := dBclient.Database("bookstore").Collection("books")

	idParam := r.URL.Query().Get("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
