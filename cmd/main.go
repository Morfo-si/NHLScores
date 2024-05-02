package main

import (
	"log"
	"os"
	"time"

	"github.com/Morfo-si/NHLScores.git/internal/db"
	"github.com/Morfo-si/NHLScores.git/internal/handler"
	"github.com/Morfo-si/NHLScores.git/internal/model"
	"github.com/Morfo-si/NHLScores.git/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	// Connect to the database
	db, err := db.Connect()
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Game{})

	// Initialize your repository
	gameRepository := repository.NewGameRepository(db)

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

	// Setup product handlers
	handler.SetupProductHandlers(app, gameRepository)

	log.Fatal(app.Listen(":8080"))
}
