package handlers

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/varnitha0415/GoLearnings/bookstore_api_fiber/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookStoreHandlerImpl struct{}

// GetAllBooks retrieves all books from the database
func (h *BookStoreHandlerImpl) GetAllBooks(c *fiber.Ctx, dbClient *mongo.Client) {
	collection := dbClient.Database("bookstore").Collection("bookstore")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to retrieve books",
			"data":    err.Error(),
		})
	}
	defer cursor.Close(context.Background())
	var books []models.BookStore
	for cursor.Next(context.Background()) {
		var book models.BookStore
		if err := cursor.Decode(&book); err != nil {

			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to decode book",
				"data":    err.Error(),
			})

		}
		books = append(books, book)
	}
	return c.json(fiber.Map{
		"status":  "success",
		"message": "Books retrieved successfully",
		"data":    books,
	})
}

// AddBook adds a new book to the database
func (h *BookStoreHandlerImpl) AddBook(c *fiber.Ctx, dbClient *mongo.Client) {

	collection := dbClient.Database("bookstore").Collection("bookstore")

	// Decode the request body into a Book struct
	var bookstore models.BookStore
	if err := c.BodyParse(bookstore); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request payload",
			"data":    err.Error(),
		})
	}

	// Insert the book into the collection
	bookstore.ID = primitive.NewObjectID() // Generate a new ObjectID
	_, err := collection.InsertOne(context.Background(), bookstore)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to insert book",
			"data":    err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Book added successfully",
		"data":    bookstore,
	})
}

// GetBook retrieves a book by its ID
func (h *BookStoreHandlerImpl) GetBook(c *fiber.Ctx, dbClient *mongo.Client) {
	collection := dbClient.Database("bookstore").Collection("bookstore")

	idParam := c.URL.Query().Get("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid bookstore ID",
			"data":    err.Error(),
		})
	}

	var bookstore models.BookStore
	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&bookstore)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "BookStore not found",
				"data":    err.Error(),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to retrieve book",
			"data":    err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "BookStore retrieved successfully",
		"data":    bookstore,
	})
}

// UpdateBook updates an existing book in the database
func (h *BookStoreHandlerImpl) UpdateBook(c *fiber.Ctx, dbClient *mongo.Client) {

	collection := dbClient.Database("bookstore").Collection("bookstore")

	var bookstore models.BookStore
	if err := c.BodyParse(bookstore); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request payload",
			"data":    err.Error(),
		})
	}

	idParam := c.URL.Query().Get("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid bookstore ID",
			"data":    err.Error(),
		})
	}

	update := bson.M{
		"$set": bson.M{
			"title":        bookstore.Name,
			"author":       bookstore.Area,
			"published_at": bookstore.City,
		},
	}

	result, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to update bookstore",
			"data":    err.Error(),
		})
	}

	if result.MatchedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "BookStore not found",
			"data":    "No book found with the provided ID",
		})
	}

	c.json(fiber.Map{
		"status":  "success",
		"message": "BookStore updated successfully",
		"data":    bookstore,
	})
}

func (h *BookStoreHandlerImpl) DeleteBook(c *fiber.Ctx, dbClient *mongo.Client) {
	collection := dbClient.Database("bookstore").Collection("bookstore")

	idParam := c.URL.Query().Get("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid bookstore ID",
			"data":    err.Error(),
		})
	}

	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete bookstore",
			"data":    err.Error(),
		})
	}

	if result.DeletedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "BookStore not found",
			"data":    "No book found with the provided ID",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "BookStore deleted successfully",
		"data":    nil,
	})
}
