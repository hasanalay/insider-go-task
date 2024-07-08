package models

import (
	"gorm.io/gorm"
)

type Team struct {
	ID             uint   `gorm:"primary key; autoIncrement" json:"id"`
	TeamName       string `json:"team_name"`
	Points         uint   `json:"points"`
	Win            uint   `json:"win"`
	Draw           uint   `json:"lose"`
	Lose           uint   `json:"draw"`
	GoalDifference int    `json:"goal_difference"`
	Power          int    `json:"power"`
}

type Match struct {
	ID        uint `gorm:"primary key; autoIncrement" json:"id"`
	Week      uint `gorm:"default:0" json:"week"`
	HomeID    uint `gorm:"default:0" json:"home_id"`
	AwayID    uint `gorm:"default:0" json:"away_id"`
	HomeGoals uint `gorm:"default:0" json:"home_goals"`
	AwayGoals uint `gorm:"default:0" json:"away_goals"`
	IsPlayed  bool `gorm:"default:false" json:"is_played"`
}

func MigrateTeams(db *gorm.DB) error {
	err := db.AutoMigrate(&Team{})
	return err
}

func MigrateMatches(db *gorm.DB) error {
	err := db.AutoMigrate(&Match{})
	return err
}
