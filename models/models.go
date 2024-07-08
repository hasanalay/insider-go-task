package models

import (
	"gorm.io/gorm"
)

type Team struct {
	ID             int    `gorm:"primary key; autoIncrement" json:"id"`
	TeamName       string `json:"team_name"`
	Points         int    `json:"points"`
	Win            int    `json:"win"`
	Draw           int    `json:"lose"`
	Lose           int    `json:"draw"`
	GoalDifference int    `json:"goal_difference"`
	Power          int    `json:"power"`
}

type Match struct {
	ID        int `gorm:"primary key; autoIncrement" json:"id"`
	Week      int `json:"week"`
	HomeID    int `json:"home_id"`
	AwayID    int `json:"away_id"`
	HomeGoals int `json:"home_goals"`
	AwayGoals int `json:"away_goals"`
}

func MigrateTeams(db *gorm.DB) error {
	err := db.AutoMigrate(&Team{})
	return err
}

func MigrateMatches(db *gorm.DB) error {
	err := db.AutoMigrate(&Match{})
	return err
}
