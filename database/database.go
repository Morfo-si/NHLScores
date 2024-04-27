package database

import (
	"github.com/Morfo-si/NHLScores.git/game"
	"gorm.io/gorm"
)

var DBConn *gorm.DB
var Repo game.BookRepository

// InitDB creates a SQLite DB and populates it.
func InitDB() {
	Repo = game.NewSqliteRepository()
}
