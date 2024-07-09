package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hasanalay/insider-go-task/internal/models"
	"github.com/hasanalay/insider-go-task/internal/repository"
)

// region Team
func GetTeams(c *fiber.Ctx) error {
	teams, err := repository.GetTeams()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Teams fetched successfully!",
		"data":    teams,
	})
}

func CreateTeam(c *fiber.Ctx) error {
	team := new(models.Team)
	if err := c.BodyParser(team); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := repository.CreateTeam(team); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Team created successfully!",
		"data":    team,
	})
}

func GetTeamByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid team ID",
		})
	}
	team, err := repository.GetTeamByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if team == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Team not found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Team fetched successfully!",
		"data":    team,
	})
}

func UpdateTeam(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid team ID",
		})
	}
	team := new(models.Team)
	if err := c.BodyParser(team); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := repository.UpdateTeam(uint(id), team); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Team updated successfully!",
		"data":    team,
	})
}

func DeleteTeam(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid team ID",
		})
	}
	if err := repository.DeleteTeam(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

//endregion Team

// region Match
func GetMatches(c *fiber.Ctx) error {
	matches, err := repository.GetMatches()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Matches fetched successfully!",
		"data":    matches,
	})
}

func CreateMatch(c *fiber.Ctx) error {
	match := new(models.Match)
	if err := c.BodyParser(match); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := repository.CreateMatch(match); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Match created successfully!",
		"data":    match,
	})
}

func GetMatchByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid match ID",
		})
	}
	match, err := repository.GetMatchByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if match == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Match not found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Match fetched successfully!",
		"data":    match,
	})
}

func UpdateMatch(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid match ID",
		})
	}
	match := new(models.Match)
	if err := c.BodyParser(match); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := repository.UpdateMatch(uint(id), match); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Match updated successfully!",
		"data":    match,
	})
}

func DeleteMatch(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid match ID",
		})
	}
	if err := repository.DeleteMatch(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func GetMatchesByWeek(c *fiber.Ctx) error {
	week, err := strconv.Atoi(c.Params("week"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid match week number",
		})
	}
	matches, err := repository.GetMatchesByWeek(uint(week))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if len(matches) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Matches not found",
		})
	}
	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Matches of week:%s fetched successfully!", c.Params("week")),
		"data":    matches,
	})
}

func PlayMatch(c *fiber.Ctx) error {
	week, err := strconv.Atoi(c.Params("week"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid match week number",
		})
	}
	matches, err := repository.PlayMatch(uint(week))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if len(matches) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Matches not found",
		})
	}
	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Matches of week:%s played successfully!", c.Params("week")),
		"data":    matches,
	})
}

func PlayAllMatches(c *fiber.Ctx) error {
	matches, err := repository.PlayAllMatches()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if len(matches) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Matches not found",
		})
	}
	return c.JSON(fiber.Map{
		"message": " All Matches played successfully!",
		"data":    matches,
	})
}

//endregion Match
