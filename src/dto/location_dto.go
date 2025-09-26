package dto

import (
	"time"

	"github.com/google/uuid"
)

// LocationBase contains basic location information for creation/updates
type LocationBase struct {
	Name        string   `json:"name" validate:"required,min=1,max=255"`
	City        string   `json:"city" validate:"required,min=1,max=100"`
	State       string   `json:"state" validate:"max=100"`
	Country     string   `json:"country" validate:"required,min=1,max=100"`
	Latitude    *float64 `json:"latitude" validate:"omitempty,min=-90,max=90"`
	Longitude   *float64 `json:"longitude" validate:"omitempty,min=-180,max=180"`
	PostalCode  string   `json:"postal_code" validate:"max=20"`
	Address     string   `json:"address" validate:"max=500"`
	Description string   `json:"description" validate:"max=1000"`
	IsActive    *bool    `json:"is_active,omitempty"`
}

// LocationDetails contains detailed location information including relationships
type LocationDetails struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Country     string    `json:"country"`
	Latitude    *float64  `json:"latitude"`
	Longitude   *float64  `json:"longitude"`
	PostalCode  string    `json:"postal_code"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relationships
	Theatres []TheatreSummary `json:"theatres,omitempty"`
}

// LocationSummary contains summary location information for lists
type LocationSummary struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Country   string    `json:"country"`
	Latitude  *float64  `json:"latitude"`
	Longitude *float64  `json:"longitude"`
	IsActive  bool      `json:"is_active"`
}
