package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/varnitha0415/GoLearnings/bookstore_api_fiber/config"
	"github.com/varnitha0415/GoLearnings/bookstore_api_fiber/handlers"
)

func main() {

	app := fiber.New()

	collection := config.ConnectToMongoDB()
	bookstoreHandler := handlers.NewBookHandler(collection)

	app.Get("/bookstore", bookstoreHandler.GetAllBookStore)

	app.Listen(":3000")

}
