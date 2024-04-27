package database

import (
	"github.com/Morfo-si/NHLScores.git/domain"
	"gorm.io/gorm"
)

var DBConn *gorm.DB
var Repo domain.BookRepository

// InitDB creates a SQLite DB and populates it.
func InitDB() {
	Repo = domain.NewSqliteRepository()
}
