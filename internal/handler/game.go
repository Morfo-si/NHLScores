package handler

import (
	"github.com/Morfo-si/NHLScores.git/internal/model"
	"github.com/Morfo-si/NHLScores.git/internal/repository"
	"github.com/gofiber/fiber/v2"
)

func Status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

func Api(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(c.App().GetRoutes())
}

func SetupProductHandlers(app *fiber.App, gameRepository repository.GameRepository) {
	app.Get("/", Status)

	api := app.Group("/api")

	// routes
	api.Get("/", Api)
	api.Get("/game", func(c *fiber.Ctx) error {
		games, err := gameRepository.GetGames()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				&fiber.Map{
					"success": false,
					"message": err.Error(),
					"data":    games,
				},
			)
		}
		return c.Status(fiber.StatusOK).JSON(
			&fiber.Map{
				"success": true,
				"message": "",
				"data":    games,
			},
		)
	})

	api.Get("/game/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				&fiber.Map{
					"success": false,
					"message": "id cannot be empty",
					"data":    nil,
				})
		}
		game, err := gameRepository.GetGame(uint(id))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(
				&fiber.Map{
					"success": false,
					"message": fiber.ErrNotFound.Error(),
					"data":    nil,
				},
			)
		}
		return c.Status(fiber.StatusOK).JSON(
			&fiber.Map{
				"success": true,
				"message": "",
				"data":    game,
			},
		)
	})

	api.Post("/game", func(c *fiber.Ctx) error {
		var game model.Game
		if err := c.BodyParser(&game); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				&fiber.Map{
					"success": false,
					"message": fiber.ErrBadRequest.Error(),
					"data":    game,
				})
		}
		if err := gameRepository.NewGame(&game); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				&fiber.Map{
					"success": false,
					"message": fiber.ErrBadRequest.Error(),
					"data":    game,
				})
		}
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"success": true,
				"message": "Game created successfully",
				"data":    game,
			})
	})

	api.Put("/game/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				&fiber.Map{
					"success": false,
					"message": "id cannot be empty",
					"data":    id,
				})
		}

		var game model.Game
		if err := c.BodyParser(&game); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				&fiber.Map{
					"success": false,
					"message": fiber.ErrBadRequest.Error(),
					"data":    game,
				})
		}

		game.ID = uint(id)
		if err := gameRepository.UpdateGame(game); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				&fiber.Map{
					"success": false,
					"message": "id cannot be empty",
					"data":    id,
				})
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"success": true,
			"message": "Game updated successfully",
			"data":    game,
		})
	})

	api.Delete("/game/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				&fiber.Map{
					"success": false,
					"message": "id cannot be empty",
					"data":    nil,
				})
		}
		if err := gameRepository.DeleteGame(uint(id)); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(
				&fiber.Map{
					"success": false,
					"message": fiber.ErrNotFound.Error(),
					"data":    nil,
				})
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"success": true,
			"message": "Game successfully deleted",
			"data":    id,
		})
	})
}
