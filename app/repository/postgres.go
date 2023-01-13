package repository

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"maply/config"
	"maply/models"
)

var DB *gorm.DB

func InitPostgres(cfg config.PostgresConfig) {
	var err error

	// Connect to database
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.DBName,
	)
	//DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
	//	Logger: logger.Default.LogMode(logger.Info),
	//})
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Fatalf("Failed to initialize PostgreSQL: %s", err.Error())
	}
	log.Info("Connection opened to PostgreSQL")

	// Make auto migrations
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Request{})
	log.Info("PostgreSQL migrated")
}
