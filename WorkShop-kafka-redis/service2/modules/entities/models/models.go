package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type MI struct {
	Imperial string
	Metric   string
}

type DogResponse struct {
	ID               uint
	Name             string
	Temperament      string
	LifeSpan         string
	Origin           string
	Weight           MI
	Height           MI
	BredFor          string
	BreedGroup       string
	ReferenceImageID string
}

// data base
type Dog struct {
	gorm.Model
	ID               uint   `json:"id" gorm:"primaryKey"`
	Name             string `json:"name" gorm:"not null"`
	Temperament      string `json:"temperament"`
	LifeSpan         string `json:"life_span"`
	Origin           string `json:"origin"`
	WeightImperial   string `json:"weight_imperial" gorm:"column:weight_imperial"`
	WeightMetric     string `json:"weight_metric" gorm:"column:weight_metric"`
	HeightImperial   string `json:"height_imperial" gorm:"column:height_imperial"`
	HeightMetric     string `json:"height_metric" gorm:"column:height_metric"`
	BredFor          string `json:"bred_for"`
	BreedGroup       string `json:"breed_group"`
	ReferenceImageID string `json:"reference_image_id"`
}

type User struct {
	ID       uint
	CreateAt time.Time
	UpdateAt time.Time
	Name     string
	Email    string
}


type DogRow struct {
	ID      uint
	RawData json.RawMessage
}

type ResponseError struct {
	Message    string `json:"message"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
}

type ResponseData struct {
	Message    string      `json:"message"`
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}