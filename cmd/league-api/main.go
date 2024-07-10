package main

import (
	"log"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/hasanalay/insider-go-task/internal/db"
	"github.com/hasanalay/insider-go-task/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {

	// load .env file
	envPath, err := filepath.Abs("../../.env")
	if err != nil {
		log.Fatal("Error determining .env file path: ", err)
	}

	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	app := fiber.New()

	// connect database
	if err := db.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	// routes
	api := app.Group("/api")
	api.Get("/teams", handlers.GetTeams)
	api.Post("/teams", handlers.CreateTeam)
	api.Get("/teams/:id", handlers.GetTeamByID)
	api.Put("/teams/:id", handlers.UpdateTeam)
	app.Delete("/teams/:id", handlers.DeleteTeam)

	api.Get("/matches", handlers.GetMatches)
	api.Post("/matches", handlers.CreateMatch)
	api.Get("/matches/:id", handlers.GetMatchByID)
	api.Put("/matches/:id", handlers.UpdateMatch)
	api.Delete("/matches/:id", handlers.DeleteMatch)
	api.Get("/matches/week/:week", handlers.GetMatchesByWeek)

	api.Get("/league/:week", handlers.PlayMatch)
	api.Get("/league", handlers.PlayAllMatches)
	api.Put("/league/change-match/:id", handlers.ChangeMatchResult)

	app.Listen(":3000")
}
