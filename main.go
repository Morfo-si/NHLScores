package main

import (
	"log"
	"time"

	"github.com/Morfo-si/NHLScores.git/database"
	"github.com/Morfo-si/NHLScores.git/game"
	"github.com/gofiber/fiber/v2"
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

func main() {
	app := fiber.New(fiber.Config{
		AppName:           "NHLScores",
		BodyLimit:         fiber.DefaultBodyLimit,
		EnablePrintRoutes: true,
		ServerHeader:      "NHLScores",
		StrictRouting:     true,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       10 * time.Second,
	},
	)
	// app.Use(cors.New())

	InitDB()

	app.Get("/game", game.GetGames)
	app.Get("/game/:id", game.GetGame)
	app.Post("/game", game.NewGame)
	app.Delete("/game/:id", game.DeleteGame)

	log.Fatal(app.Listen(":8080"))
}
