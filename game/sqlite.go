package game

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SqliteDBRepository fulfills the Repository interface
type SqliteRepository struct {
	Db *gorm.DB
}

func NewSqliteRepository() *SqliteRepository {
	db, err := gorm.Open(
		sqlite.Open("hockey.db"),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	if err != nil {
		log.Fatal("Failed to connect to the database. \n", err)
		return nil
	}

	err = db.AutoMigrate(&Game{})
	if err != nil {
		log.Fatal("Failed to migrate the database schema. \n", err)
		return nil
	}

	return &SqliteRepository{
		Db: db,
	}
}

func (s *SqliteRepository) GetGames(c *fiber.Ctx) error {
	var games []Game
	s.Db.Find(&games)
	return c.Status(fiber.StatusOK).JSON(
		&fiber.Map{
			"success": true,
			"message": "",
			"data":    games,
		},
	)
}

func (s *SqliteRepository) GetGame(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"success": false,
				"message": "id cannot be empty",
				"data":    nil,
			})
	}

	var game Game
	s.Db.Find(&game, id)
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

func (s *SqliteRepository) NewGame(c *fiber.Ctx) error {
	game := new(Game)
	if err := c.BodyParser(game); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"success": false,
				"message": fiber.ErrBadRequest.Error(),
				"data":    game,
			})
	}
	s.Db.Create(&game)
	return c.Status(fiber.StatusCreated).JSON(game)
}

func (s *SqliteRepository) DeleteGame(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"success": false,
				"message": "id cannot be empty",
				"data":    nil,
			})
	}

	var game Game
	s.Db.First(&game, id)
	if game.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(
			&fiber.Map{
				"success": false,
				"message": fiber.ErrNotFound.Error(),
				"data":    nil,
			})
	}

	s.Db.Delete(&game)
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "Game successfully deleted",
		"data":    game,
	})
}

func (s *SqliteRepository) UpdateGame(c *fiber.Ctx) error {
	id := c.Params("id")

	game := new(Game)
	if err := c.BodyParser(game); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"success": false,
				"message": fiber.ErrBadRequest.Error(),
				"data":    game,
			})
	}
	if err := s.Db.Model(Game{}).Where("id = ?", id).Updates(game).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"success": false,
				"message": fiber.ErrNotFound.Error(),
				"data":    game,
			})
	}
	s.Db.Find(&game, id)
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "Game successfully updated",
		"data":    game,
	})
}
