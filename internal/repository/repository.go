package repository

import (
	"errors"

	"github.com/hasanalay/insider-go-task/helpers"
	"github.com/hasanalay/insider-go-task/internal/db"
	"github.com/hasanalay/insider-go-task/internal/models"
)

// GetTeams  retrieves all teams
//
//	@return []models.Team
//	@return error
func GetTeams() ([]models.Team, error) {
	teams := []models.Team{}
	if err := db.DB.Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}

// CreateTeam  creates team
//
//	@param team
//	@return error
func CreateTeam(team *models.Team) error {
	return db.DB.Create(team).Error
}

// GetTeamByID get team information with an id
//
//	@param id
//	@return *models.Team
//	@return error
func GetTeamByID(id uint) (*models.Team, error) {
	team := models.Team{}
	if err := db.DB.First(&team, id).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

// UpdateTeam updates team
//
//	@param id
//	@param updatedTeam
//	@return error
func UpdateTeam(id uint, updatedTeam *models.Team) error {
	return db.DB.Model(&models.Team{}).Where("id = ?", id).Updates(updatedTeam).Error
}

// DeleteTeam deletes team
//
//	@param id
//	@return error
func DeleteTeam(id uint) error {
	return db.DB.Delete(&models.Team{}, id).Error
}

// GetMatches retrieves all matches
//
//	@return []models.Match
//	@return error
func GetMatches() ([]models.Match, error) {
	matches := []models.Match{}
	if err := db.DB.Preload("HomeTeam").Preload("AwayTeam").Find(&matches).Error; err != nil {
		return nil, err
	}
	return matches, nil
}

// CreateMatch creates match
//
//	@param match
//	@return error
func CreateMatch(match *models.Match) error {
	return db.DB.Create(match).Error
}

// GetMatchByID get team information with an id
//
//	@param id
//	@return *models.Match
//	@return error
func GetMatchByID(id uint) (*models.Match, error) {
	match := models.Match{}
	if err := db.DB.Preload("HomeTeam").Preload("AwayTeam").First(&match, id).Error; err != nil {
		return nil, err
	}
	return &match, nil
}

// UpdateMatch updates match
//
//	@param id
//	@param updatedMatch
//	@return error
func UpdateMatch(id uint, updatedMatch *models.Match) error {
	match := models.Match{}
	if err := db.DB.Preload("HomeTeam").Preload("AwayTeam").First(&match, id).Error; err != nil {
		return err
	}
	if !match.IsPlayed {
		return db.DB.Model(&models.Match{}).Where("id = ?", id).Updates(updatedMatch).Error
	}
	return errors.New("the match has already been played and cannot be updated")
}


// DeleteMatch deletes match
//
//	@param id
//	@return error
func DeleteMatch(id uint) error {
	return db.DB.Delete(&models.Match{}, id).Error
}

// GetMatchesByWeek retrieves given week's matches
//
//	@param week
//	@return []models.Match
//	@return error
func GetMatchesByWeek(week uint) ([]models.Match, error) {
	matches := []models.Match{}

	if err := db.DB.Preload("HomeTeam").Preload("AwayTeam").Where("week = ?", week).Find(&matches).Error; err != nil {
		return nil, err
	}
	return matches, nil
}

// PlayMatch plays given week's matches
//
//	@param week
//	@return []models.Match
//	@return error
func PlayMatch(week uint) ([]models.Match,[]models.Team, error) {
	matches := []models.Match{}

	if err := db.DB.Preload("HomeTeam").Preload("AwayTeam").Where("week = ?", week).Find(&matches).Error; err != nil {
		return nil, nil, err
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
	teams := []models.Team{}
	if err := db.DB.Order("points DESC").Order("goal_difference DESC").Find(&teams).Error; err != nil {
		return nil, nil, err
	}
	return matches, teams, nil
}

// PlayAllMatches  plays all matches at once
//
//	@return []models.Match
//	@return error
func PlayAllMatches() ([]models.Match,[]models.Team, error) {
	matches := []models.Match{}

	if err := db.DB.Preload("HomeTeam").Preload("AwayTeam").Find(&matches).Error; err != nil {
		return nil, nil, err
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
	teams := []models.Team{}
	if err := db.DB.Order("points DESC").Order("goal_difference DESC").Find(&teams).Error; err != nil {
		return nil, nil, err
	}
	return matches, teams, nil
}

// ChangeMatchResult changes match results of given id's match
//
//	@param id
//	@param updatedMatch
//	@return *models.Match
//	@return error
func ChangeMatchResult(id uint, updatedMatch *models.Match) (*models.Match, *[]models.Team, error) {
	match := models.Match{}
	if err := db.DB.Preload("HomeTeam").Preload("AwayTeam").First(&match, id).Error; err != nil {
		return nil, nil, err
	}

	var homeTeam, awayTeam models.Team
	db.DB.First(&homeTeam, match.HomeID)
	db.DB.First(&awayTeam, match.AwayID)

	if match.HomeGoals > match.AwayGoals {
		homeTeam.Win--
		homeTeam.Points -= 3
		awayTeam.Lose--
	} else if match.AwayGoals > match.HomeGoals {
		awayTeam.Win--
		awayTeam.Points -= 3
		homeTeam.Lose--
	} else {
		homeTeam.Draw--
		homeTeam.Points--
		awayTeam.Draw--
		awayTeam.Points--
	}

	match.HomeGoals = updatedMatch.HomeGoals
	match.AwayGoals = updatedMatch.AwayGoals
	match.IsPlayed = updatedMatch.IsPlayed

	homeTeam.GoalDifference = int(updatedMatch.HomeGoals) - int(updatedMatch.AwayGoals)
	awayTeam.GoalDifference = int(updatedMatch.AwayGoals) - int(updatedMatch.HomeGoals)

	if updatedMatch.HomeGoals > updatedMatch.AwayGoals {
		homeTeam.Win++
		homeTeam.Points += 3
		awayTeam.Lose++
	} else if updatedMatch.AwayGoals > updatedMatch.HomeGoals {
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
	db.DB.Save(&match)

	matchResult := models.Match{}
	if err := db.DB.Preload("HomeTeam").Preload("AwayTeam").First(&matchResult, id).Error; err != nil {
		return nil, nil, err
	}
	teams := []models.Team{}
	if err := db.DB.Order("points DESC").Order("goal_difference DESC").Find(&teams).Error; err != nil {
		return nil, nil, err
	}

	return &matchResult, &teams, nil
}
