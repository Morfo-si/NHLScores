package game

import (
	"fmt"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	HomeTeam     string `json:"home_team"`
	VisitorTeam  string `json:"visitor_team"`
	HomeScore    int    `json:"home_score"`
	VisitorScore int    `json:"visitor_score"`
	Date         string `json:"date"`
}

func (g *Game) Validate() error {
	if g.ID < 1 {
		return fmt.Errorf("user ID cannot be zero or less")
	}
	if g.HomeScore < 1 {
		return fmt.Errorf("home score cannot be zero or less")
	}
	if g.VisitorScore < 1 {
		return fmt.Errorf("visitor score cannot be zero or less")
	}
	return nil
}
