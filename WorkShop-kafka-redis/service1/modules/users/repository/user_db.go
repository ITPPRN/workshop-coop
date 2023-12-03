package repository

import (
	"fmt"

	"gorm.io/gorm"

	"service1/modules/entities/models"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) userRepositoryDB {
	db.AutoMigrate(models.User{})
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) CreateUser(name string, email string) (*models.User, error) {
	newUser := &models.User{
		Name:  name,
		Email: email,
	}
	if err := r.db.Create(newUser).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	return newUser, nil
}

func (r userRepositoryDB) UpdateUser(id uint, name string, email string) (*models.User, error) {
	upreq := &models.User{Name: name, Email: email}
	if err := r.db.Model(&models.User{}).Where("id = ?", id).Updates(upreq).Error; err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}
	return upreq, nil
}

func (r userRepositoryDB) DeleteUser(id uint) (*string, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}

	if err := r.db.Delete(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to delete user: %v", err)
	}

	message := fmt.Sprintf("Delete ID %d Successfuly", id)
	return &message, nil
}
