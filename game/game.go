package game

import (
	"fmt"

	"github.com/Morfo-si/NHLScores.git/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	HomeTeam     string `json:"home_team"`
	VisitorTeam  string `json:"visitor_team"`
	HomeScore    int    `json:"home_score"`
	VisitorScore int    `json:"visitor_score"`
	Date         string `json:"date"`
}

func GetGames(c *fiber.Ctx) error {
	db := database.DBConn
	var games []Game
	db.Find(&games)
	return c.Status(fiber.StatusOK).JSON(
		&fiber.Map{
			"success": true,
			"message": "",
			"data":    games,
		},
	)
}

func GetGame(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"success": false,
				"message": "id cannot be empty",
				"data":    nil,
			})
	}
	db := database.DBConn

	var game Game
	db.Find(&game, id)
	if game.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(
			&fiber.Map{
				"success": false,
				"message": fiber.ErrNotFound.Error(),
				"data":    nil,
			})
	}
	return c.Status(fiber.StatusOK).JSON(
		&fiber.Map{
			"success": true,
			"message": "",
			"data":    game,
		},
	)
}

func NewGame(c *fiber.Ctx) error {
	db := database.DBConn
	game := new(Game)
	if err := c.BodyParser(game); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"success": false,
				"message": fiber.ErrBadRequest.Error(),
				"data":    game,
			})
	}
	db.Create(&game)
	return c.Status(fiber.StatusCreated).JSON(game)
}

func DeleteGame(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"success": false,
				"message": "id cannot be empty",
				"data":    nil,
			})
	}

	db := database.DBConn

	var game Game
	db.First(&game, id)
	if game.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(
			&fiber.Map{
				"success": false,
				"message": fiber.ErrNotFound.Error(),
				"data":    nil,
			})
	}

	db.Delete(&game)
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "Game successfully deleted",
		"data":    game,
	})
}

func UpdateGame(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	game := new(Game)
	if err := c.BodyParser(game); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"success": false,
				"message": fiber.ErrBadRequest.Error(),
				"data":    game,
			})
	}
	if err := db.Model(Game{}).Where("id = ?", id).Updates(game).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"success": false,
				"message": fiber.ErrNotFound.Error(),
				"data":    game,
			})
	}
	db.Find(&game, id)
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "Game successfully updated",
		"data":    game,
	})
}

func (g *Game) Validate() error {
	if g.ID < 1 {
		return fmt.Errorf("user ID cannot be zero or less")
	}
	if g.HomeScore < 1 {
		return fmt.Errorf("home score cannot be zero or less")
	}
	if g.VisitorScore < 1 {
		return fmt.Errorf("visitor score cannot be zero or less")
	}
	return nil
}
