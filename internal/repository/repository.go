package repository

import (
	"github.com/hasanalay/insider-go-task/internal/db"
	"github.com/hasanalay/insider-go-task/internal/models"
)

func GetTeams() ([]models.Team, error) {
	teams := []models.Team{}
	if err := db.DB.Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}

func CreateTeam(team *models.Team) error {
	return db.DB.Create(team).Error
}

func GetTeamByID(id uint) (*models.Team, error) {
	team := models.Team{}
	if err := db.DB.First(&team, id).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

func UpdateTeam(id uint, updatedTeam *models.Team) error {
	return db.DB.Model(&models.Team{}).Where("id = ?", id).Updates(updatedTeam).Error
}

func DeleteTeam(id uint) error {
	return db.DB.Delete(&models.Team{}, id).Error
}

func GetMatches() ([]models.Match, error) {
	matches := []models.Match{}
	if err := db.DB.Preload("HomeTeam").Preload("AwayTeam").Find(&matches).Error; err != nil {
		return nil, err
	}
	return matches, nil
}

func CreateMatch(match *models.Match) error {
	return db.DB.Create(match).Error
}

func GetMatchByID(id uint) (*models.Match, error) {
	match := models.Match{}
	if err := db.DB.Preload("HomeTeam").Preload("AwayTeam").First(&match, id).Error; err != nil {
		return nil, err
	}
	return &match, nil
}

func UpdateMatch(id uint, updatedMatch *models.Match) error {
	return db.DB.Model(&models.Match{}).Where("id = ?", id).Updates(updatedMatch).Error
}

func DeleteMatch(id uint) error {
	return db.DB.Delete(&models.Match{}, id).Error
}

func GetMatchesByWeek(week uint) ([]models.Match, error) {
	matches := []models.Match{}

	if err := db.DB.Preload("HomeTeam").Preload("AwayTeam").Where("week = ?", week).Find(&matches).Error; err != nil {
		return nil, err
	}
	return matches, nil
}
