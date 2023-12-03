package repositories

import (
	"gorm.io/gorm"

	"service1/modules/entities/models"
)

type consumerRepo struct {
	db *gorm.DB
}

func NewsConsumerRepository(db *gorm.DB) models.ConsumerRepository {
	db.AutoMigrate(models.UserReadDog{})
	return &consumerRepo{db: db}
}

func (r *consumerRepo) CreateUserReadedDog(arg *models.UserReadDog) error {
	if err := r.db.Create(arg).Error; err != nil {
		return err
	}
	return nil
}
