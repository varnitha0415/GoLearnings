package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/varnitha0415/GoLearnings/books_api_crud/config"
	"github.com/varnitha0415/GoLearnings/bookstore_api_fiber/handlers"
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

	app := fiber.New()

	bookstoreHandler := &handlers.BookStoreHandlerImpl{}
	var err error
	dbClient, err = ConnectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/books", func(w http.ResponseWriter, r *http.Request) {
		bookstoreHandler.
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

	app.Listen(":3000")

}
