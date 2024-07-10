package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"os"

	"github.com/hasanalay/insider-go-task/internal/models"
)

var DB *gorm.DB

func ConnectDB() error {
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASS") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=" + os.Getenv("DB_SSLMODE")

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// auto migrate models to the database
	if err := DB.AutoMigrate(&models.Team{}, &models.Match{}); err != nil {
		return err
	}

	return nil
}
