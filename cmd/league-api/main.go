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

	// Ana dizindeki .env dosyasının tam yolunu belirleme
	envPath, err := filepath.Abs("../../.env")
	if err != nil {
		log.Fatal("Error determining .env file path: ", err)
	}

	// Ortam değişkenlerini yükle
	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	app := fiber.New()

	// Veritabanı bağlantısı
	if err := db.ConnectDB(); err != nil {
		log.Fatal(err)
	}
	// Route'ları ekleme
	app.Get("/teams", handlers.GetTeams)
	app.Post("/teams", handlers.CreateTeam)
	app.Get("/teams/:id", handlers.GetTeamByID)
	app.Put("/teams/:id", handlers.UpdateTeam)
	app.Delete("/teams/:id", handlers.DeleteTeam)

	//TODO doğru metodları güncelle
	app.Get("/matches", handlers.GetMatches)
	app.Post("/matches", handlers.CreateMatch)
	app.Get("/matches/:id", handlers.GetMatchByID)
	app.Put("/matches/:id", handlers.UpdateMatch)
	app.Delete("/matches/:id", handlers.DeleteMatch)
	app.Get("matches/week/:week", handlers.GetMatchesByWeek)

	app.Listen(":3000")
}
