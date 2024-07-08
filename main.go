package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/hasanalay/insider-go-task/db"
	"github.com/hasanalay/insider-go-task/models"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type Team struct {
	TeamName       string `json:"team_name"`
	Points         uint   `json:"points"`
	Win            uint   `json:"win"`
	Draw           uint   `json:"lose"`
	Lose           uint   `json:"draw"`
	GoalDifference int    `json:"goal_difference"`
	Power          int    `json:"power"`
}
type Match struct {
	Week      uint `json:"week"`
	HomeID    uint `json:"home_id"`
	AwayID    uint `json:"away_id"`
	HomeGoals uint `json:"home_goals"`
	AwayGoals uint `json:"away_goals"`
	IsPlayed  bool `json:"is_played"`
}

func (r *Repository) createTeam(context *fiber.Ctx) error {
	team := Team{}
	err := context.BodyParser(&team)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	err = r.DB.Create(&team).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not able to create team!"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Team has been added"})
	return nil
}

func (r *Repository) getTeams(context *fiber.Ctx) error {
	teamModels := &[]models.Team{}
	err := r.DB.Find(teamModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not get the teams!"})
		return err
	}

	if len(*teamModels) == 0 {
		context.Status(http.StatusNotFound).JSON(
			&fiber.Map{"message": "There is no team in the database"})
		return nil
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Teams fetched succesfully",
			"data":    teamModels,
		})
	return nil
}

func (r *Repository) getTeamByID(context *fiber.Ctx) error {
	teamModel := &models.Team{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id cannot be empty!"})
		return nil
	}
	err := r.DB.Where("id = ?", id).First(teamModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			context.Status(http.StatusNotFound).JSON(
				&fiber.Map{"message": "Team not found!"})
		} else {
			context.Status(http.StatusBadRequest).JSON(
				&fiber.Map{"message": "Could not get the team!"})
		}
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Team fetched successfully!",
			"data":    teamModel,
		})
	return nil
}

func (r *Repository) deleteTeam(context *fiber.Ctx) error {
	teamModel := models.Team{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id cannot be empty!"})
		return nil
	}
	err := r.DB.Delete(teamModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not delete the team!"})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Team deleted successfully!"})
	return nil
}

func (r *Repository) updateTeam(context *fiber.Ctx) error {
	teamModel := models.Team{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "id cannot be empty!"})
		return nil
	}

	team := Team{}
	if err := context.BodyParser(&team); err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed!"})
		return err
	}

	result := r.DB.Model(&teamModel).Where("id = ?", id).Updates(team)
	if result.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not update the team!"})
		return result.Error
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Team updated successfully!"})
	return nil
}

func (r *Repository) getMatches(context *fiber.Ctx) error {
	matchModels := &[]models.Match{}
	err := r.DB.Find(matchModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not get the matchs!"})
		return err
	}

	if len(*matchModels) == 0 {
		context.Status(http.StatusNotFound).JSON(
			&fiber.Map{"message": "No matches found for the given week!"})
		return nil
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Matchs fetched succesfully",
			"data":    matchModels,
		})
	return nil
}

func (r *Repository) getMatchesByWeek(context *fiber.Ctx) error {
	matchModels := &[]models.Match{}
	week := context.Params("week")
	if week == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "week cannot be empty!"})
		return nil
	}
	err := r.DB.Where("week  = ?", week).Find(matchModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not get the matches of week!"})
		return err
	}

	if len(*matchModels) == 0 {
		context.Status(http.StatusNotFound).JSON(
			&fiber.Map{"message": "No matches found for the given week!"})
		return nil
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Match fetched successfully!",
			"data":    matchModels,
		})
	return nil
}

func (r *Repository) updateMatch(context *fiber.Ctx) error {
	matchModel := models.Match{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "id cannot be empty!"})
		return nil
	}

	match := Match{}
	if err := context.BodyParser(&match); err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed!"})
		return err
	}

	result := r.DB.Model(&matchModel).Where("id = ?", id).Updates(match)
	if result.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not update the match!"})
		return result.Error
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Match updated successfully!"})
	return nil
}


func (r *Repository) setupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_team", r.createTeam)
	api.Get("/teams", r.getTeams)
	api.Get("teams/:id", r.getTeamByID)
	api.Delete("delete_team/:id", r.deleteTeam)
	api.Put("update_team/:id", r.updateTeam)

	api.Get("matches", r.getMatches)
	api.Get("matches/:week", r.getMatchesByWeek)
	api.Put("update_match/:id", r.updateMatch)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := db.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}

	err = models.MigrateTeams(db)
	if err != nil {
		log.Fatal("could not migrate the Teams")
	}

	err = models.MigrateMatches(db)
	if err != nil {
		log.Fatal("could not migrate the Matches")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.setupRoutes(app)
	app.Listen(":8080")
}
