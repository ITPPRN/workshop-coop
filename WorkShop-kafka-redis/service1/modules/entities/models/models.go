package models

import (
	"encoding/json"
	"time"
)

type UserResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRequest struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// db
type User struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
}

type UserReadDog struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	CreateAt   time.Time
	UserID     uint
	DogID      uint
	DogDetails json.RawMessage
}
