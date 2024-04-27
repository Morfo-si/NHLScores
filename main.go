package main

import (
	"log"
	"os"
	"time"

	"github.com/Morfo-si/NHLScores.git/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", Status)

	api := app.Group("/api")

	// routes
	api.Get("/", Api)
	api.Get("/game", database.Repo.GetGames)
	api.Get("/game/:id", database.Repo.GetGame)
	api.Post("/game", database.Repo.NewGame)
	api.Delete("/game/:id", database.Repo.DeleteGame)
	api.Put("/game/:id", database.Repo.UpdateGame)
}

func Status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

func Api(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(c.App().GetRoutes())
}

func main() {

	database.InitDB()

	app := fiber.New(fiber.Config{
		AppName:       "NHLScores",
		BodyLimit:     fiber.DefaultBodyLimit,
		ServerHeader:  "NHLScores",
		StrictRouting: true,
		ReadTimeout:   1 * time.Second,
		WriteTimeout:  1 * time.Second,
		IdleTimeout:   10 * time.Second,
	})

	app.Use(logger.New(logger.Config{
		Format:        "${time} [${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeZone:      "UTC",
		Output:        os.Stdout,
		DisableColors: false,
	}))

	SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
