package repository

import "github.com/Morfo-si/NHLScores.git/internal/model"

type GameRepository interface {
	GetGame(id uint) (model.Game, error)
	GetGames() ([]model.Game, error)
	NewGame(game *model.Game) error
	DeleteGame(id uint) error
	UpdateGame(game model.Game) error
}
