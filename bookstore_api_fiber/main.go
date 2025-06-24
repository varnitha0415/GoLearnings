package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/varnitha0415/GoLearnings/bookstore_api_fiber/config"
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

	app.Get("/books", func(c *fiber.Ctx) error {
		bookstoreHandler.GetAllBookStore(c, dbClient)

	})

	app.Listen(":3000")

}
