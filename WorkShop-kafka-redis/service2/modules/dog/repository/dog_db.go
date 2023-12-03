package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"service2/modules/entities/models"
)

type dogRepositoryDB struct {
	db *gorm.DB
}

func NewDogRepositoryDB(db *gorm.DB) dogRepositoryDB {
	db.AutoMigrate(models.Dog{})
	return dogRepositoryDB{db: db}
}

func (r dogRepositoryDB) CreateDog(dog *models.Dog) error {
	// Create dog record in the database
	if err := r.db.Create(dog).Error; err != nil {
		return fmt.Errorf("failed to create dog: %v", err)
	}

	return nil
}

func (r dogRepositoryDB) GetDogs() ([]models.Dog, error) {
	var dogs []models.Dog
	err := r.db.Find(&dogs).Error
	if err != nil {
		return nil, errors.New("Get dog not found")
	}
	return dogs, nil
}

func (r dogRepositoryDB) HasDog() (bool, error) {
	var count int64
	err := r.db.Model(&models.Dog{}).Count(&count).Error
	return count > 0, err
}

func (r dogRepositoryDB) DogExists(dogID uint) bool {
	var dog models.Dog
	result := r.db.First(&dog, dogID)

	if result.Error == nil && result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func (r dogRepositoryDB) FindDogByID(dogID uint) (*models.Dog, error) {
	var dog models.Dog
	result := r.db.First(&dog, dogID)

	if result.Error == nil && result.RowsAffected > 0 {
		return &dog, nil
	} else {
		return nil, errors.New("find dog by ID not found")
	}
}
