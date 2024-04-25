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
	return c.JSON(games)
}

func GetGame(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var game Game
	db.Find(&game, id)
	if game.ID <= 0 {
		return c.Status(500).SendString("No game was found")
	}
	return c.JSON(game)
}

func NewGame(c *fiber.Ctx) error {
	db := database.DBConn
	game := new(Game)
	if err := c.BodyParser(game); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&game)
	return c.JSON(game)
}

func DeleteGame(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var game Game
	db.First(&game, id)
	if game.ID <= 0 {
		return c.Status(500).SendString("No game was found")
	}
	db.Delete(&game)
	return c.SendString("Game successfully deleted")
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
