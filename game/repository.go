package game

import (
	"github.com/gofiber/fiber/v2"
)

type Repository struct {
}

func NewBookRepository() *Repository {
	return &Repository{}
}

type BookRepository interface {
	NewGame(ctx *fiber.Ctx) error
	GetGame(ctx *fiber.Ctx) error
	UpdateGame(ctx *fiber.Ctx) error
	GetGames(ctx *fiber.Ctx) error
	DeleteGame(ctx *fiber.Ctx) error
}
