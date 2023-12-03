package databases

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"service2/configs"
	"service2/pkg/utils"
)

func NewPostgresConnection(cfg *configs.Config) (*gorm.DB, error) {

	dsn, err := utils.UrlBuilder("postgres", cfg)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	log.Println("postgreSQL database has been connected 🐘")
	return db, nil
}
