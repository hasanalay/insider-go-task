package models

//"gorm.io/gorm"

type Team struct {
	ID             uint   `gorm:"primary key; autoIncrement" json:"id"`
	TeamName       string `json:"team_name"`
	Points         uint   `json:"points"`
	Win            uint   `json:"win"`
	Draw           uint   `json:"draw"`
	Lose           uint   `json:"lose"`
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

	HomeTeam Team `gorm:"foreignkey:home_id" json:"home_team"`
	AwayTeam Team `gorm:"foreignkey:away_id" json:"away_team"`
}

type Prediction struct {
	Team       string
	Percentage string
}
