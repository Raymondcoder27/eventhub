package models

import "time"

type Event struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Location     string    `json:"location"`
	Date         time.Time `json:"date"`
	Category     string    `json:"category"`
	MaxAttendees string    `json:"maxAttendees"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
