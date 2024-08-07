package helpers

import (
	"fmt"
	"math/rand"
)

// PlayMatch determines match score
//
//	@param homeTeamPower
//	@param awayTeamPower
//	@return homeGoals
//	@return awayGoals
func PlayMatch(homeTeamPower int, awayTeamPower int) (homeGoals int, awayGoals int) {
	homeAdwantage := 3
	awayDisadwantage := 4

	homeTeamPower = (homeTeamPower + homeAdwantage) / 10
	awayTeamPower = (awayTeamPower + awayDisadwantage) / 8

	if homeTeamPower > awayTeamPower {
		homeGoals = rand.Intn(8) + 1
		awayGoals = rand.Intn(homeGoals) + 1
	} else if awayTeamPower > homeTeamPower {
		awayGoals = rand.Intn(8) + 1
		homeGoals = rand.Intn(awayGoals) + 1
	} else {
		homeGoals = rand.Intn(8) + 1
		awayGoals = rand.Intn(8) + 1
	}
	return
}
func PredictChampion(points, totalPoints uint) (percentage string) {
	if totalPoints == 0 || points == 0 {
		return "0%" 
	}
	return fmt.Sprintf("%d%%", (points * 100) / totalPoints)
}

