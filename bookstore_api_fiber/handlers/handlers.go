package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type BookStoreHandler interface {
	GetAllBookStore(c *fiber.Ctx) error
	AddBookStore(c *fiber.Ctx) error
	GetBookStore(c *fiber.Ctx) error
	UpdateBookStore(c *fiber.Ctx) error
	DeleteBookStore(c *fiber.Ctx) error
}
