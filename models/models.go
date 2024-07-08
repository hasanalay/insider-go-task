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
	Week      uint `json:"week"`
	HomeID    uint `json:"home_id"`
	AwayID    uint `json:"away_id"`
	HomeGoals uint `json:"home_goals"`
	AwayGoals uint `json:"away_goals"`
}

func MigrateTeams(db *gorm.DB) error {
	err := db.AutoMigrate(&Team{})
	return err
}

func MigrateMatches(db *gorm.DB) error {
	err := db.AutoMigrate(&Match{})
	return err
}
