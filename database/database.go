package database

import (
	"github.com/Morfo-si/NHLScores.git/domain"
	"github.com/Morfo-si/NHLScores.git/entity"
	"gorm.io/gorm"
)

var DBConn *gorm.DB
var Repo entity.BookRepository

// InitDB creates a SQLite DB and populates it.
func InitDB() {
	Repo = domain.NewSqliteRepository()
}
