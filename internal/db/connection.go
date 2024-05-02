package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("hockey.db"),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		})
	if err != nil {
		log.Fatal("Failed to connect to the database. \n", err)
		return nil, nil
	}
	return db, nil
}
