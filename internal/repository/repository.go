package repository

import (
	"github.com/hasanalay/insider-go-task/helpers"
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

func PlayMatch(week uint) ([]models.Match, error) {
	matches := []models.Match{}

	if err := db.DB.Preload("HomeTeam").Preload("AwayTeam").Where("week = ?", week).Find(&matches).Error; err != nil {
		return nil, err
	}
	for i := range matches {
		if !matches[i].IsPlayed {
			var homeTeam, awayTeam models.Team
			db.DB.First(&homeTeam, matches[i].HomeID)
			db.DB.First(&awayTeam, matches[i].AwayID)

			homeGoals, awayGoals := helpers.PlayMatch(homeTeam.Power, awayTeam.Power)
			matches[i].HomeGoals = uint(homeGoals)
			matches[i].AwayGoals = uint(awayGoals)
			matches[i].IsPlayed = true

			db.DB.Save(&matches[i])

			homeTeam.GoalDifference += homeGoals - awayGoals
			awayTeam.GoalDifference += awayGoals - homeGoals

			if homeGoals > awayGoals {
				homeTeam.Win++
				homeTeam.Points += 3
				awayTeam.Lose++
			} else if awayGoals > homeGoals {
				awayTeam.Win++
				awayTeam.Points += 3
				homeTeam.Lose++
			} else {
				homeTeam.Draw++
				homeTeam.Points++
				awayTeam.Draw++
				awayTeam.Points++
			}
			db.DB.Save(&homeTeam)
			db.DB.Save(&awayTeam)
		}
	}
	return matches, nil
}

func PlayAllMatches() ([]models.Match, error) {
	matches := []models.Match{}

	if err := db.DB.Preload("HomeTeam").Preload("AwayTeam").Find(&matches).Error; err != nil {
		return nil, err
	}
	for i := range matches {
		if !matches[i].IsPlayed {
			var homeTeam, awayTeam models.Team
			db.DB.First(&homeTeam, matches[i].HomeID)
			db.DB.First(&awayTeam, matches[i].AwayID)

			homeGoals, awayGoals := helpers.PlayMatch(homeTeam.Power, awayTeam.Power)
			matches[i].HomeGoals = uint(homeGoals)
			matches[i].AwayGoals = uint(awayGoals)
			matches[i].IsPlayed = true
			db.DB.Save(&matches[i])

			homeTeam.GoalDifference += homeGoals - awayGoals
			awayTeam.GoalDifference += awayGoals - homeGoals
			if homeGoals > awayGoals {
				homeTeam.Win++
				homeTeam.Points += 3
				awayTeam.Lose++
			} else if awayGoals > homeGoals {
				awayTeam.Win++
				awayTeam.Points += 3
				homeTeam.Lose++
			} else {
				homeTeam.Draw++
				homeTeam.Points++
				awayTeam.Draw++
				awayTeam.Points++
			}
			db.DB.Save(&homeTeam)
			db.DB.Save(&awayTeam)
		}

	}
	return matches, nil
}
