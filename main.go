package main

import (
	"log"
	"os"
	"time"

	"github.com/Morfo-si/NHLScores.git/database"
	"github.com/Morfo-si/NHLScores.git/game"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitDB creates a SQLite DB and populates it.
func InitDB() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("hockey.db"))
	if err != nil {
		log.Fatal(err)
	}

	database.DBConn.AutoMigrate(&game.Game{})
}

func SetupRoutes(app *fiber.App) {
	app.Get("/", Status)

	api := app.Group("/api")

	// routes
	api.Get("/", Api)
	api.Get("/game", game.GetGames)
	api.Get("/game/:id", game.GetGame)
	api.Post("/game", game.NewGame)
	api.Delete("/game/:id", game.DeleteGame)
	api.Put("/game/:id", game.UpdateGame)
}

func Status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

func Api(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(c.App().GetRoutes())
}

func main() {

	InitDB()

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
