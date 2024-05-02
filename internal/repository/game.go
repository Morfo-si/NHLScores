package repository

import (
	"github.com/Morfo-si/NHLScores.git/internal/model"
	"gorm.io/gorm"
)

type GameRepositoryImpl struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) GameRepository {
	return &GameRepositoryImpl{db: db}
}

func (r *GameRepositoryImpl) GetGames() ([]model.Game, error) {
	var games []model.Game
	if err := r.db.Find(&games).Error; err != nil {
		return nil, err
	}
	return games, nil
}

func (r *GameRepositoryImpl) GetGame(id uint) (model.Game, error) {
	var game model.Game
	if err := r.db.First(&game, id).Error; err != nil {
		return model.Game{}, err
	}
	return game, nil
}

func (r *GameRepositoryImpl) NewGame(game *model.Game) error {
	if err := r.db.Create(&game).Error; err != nil {
		return err
	}
	return nil
}

func (r *GameRepositoryImpl) UpdateGame(game model.Game) error {
	if err := r.db.Save(&game).Error; err != nil {
		return err
	}
	return nil
}

func (r *GameRepositoryImpl) DeleteGame(id uint) error {
	if err := r.db.Delete(&model.Game{}, id).Error; err != nil {
		return err
	}
	return nil
}
