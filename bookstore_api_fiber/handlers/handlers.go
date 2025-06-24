package handlers

import (
	"github.com/gofiber/fiber"
)

type BookStoreHandler interface {
	GetAllBookStore(c *fiber.Ctx)
	AddBookStore(c *fiber.Ctx)
	GetBookStore(c *fiber.Ctx)
	UpdateBookStore(c *fiber.Ctx)
	DeleteBookStore(c *fiber.Ctx)
}
