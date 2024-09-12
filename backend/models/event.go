package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Location     string    `json:"location"`
	Date         string    `json:"date"`
	Category     string    `json:"category"`
	MaxAttendees string    `json:"maxAttendees"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
