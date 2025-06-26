package main

import (
	"github.com/gofiber/fiber"
	"github.com/varnitha0415/GoLearnings/bookstore_api_fiber/config"
)

func main() {

	app := fiber.New()

	collection := config.ConnectToMongoDB()
	bookstoreHandler := handlers.bookStoreHandlerImpl(collection)

	app.Get("/books", func(c *fiber.Ctx) error {
		return bookstoreHandler.GetAllBookStore(c, dbClient)
	})

	app.Listen(":3000")

}
