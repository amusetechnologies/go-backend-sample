package dto

import (
	"time"

	"github.com/google/uuid"
)

// TheatreBase contains basic theatre information for creation/updates
type TheatreBase struct {
	Name          string    `json:"name" validate:"required,min=1,max=255"`
	Description   string    `json:"description" validate:"max=2000"`
	Capacity      *int      `json:"capacity" validate:"omitempty,min=1,max=100000"`
	Address       string    `json:"address" validate:"max=500"`
	Phone         string    `json:"phone" validate:"max=20"`
	Email         string    `json:"email" validate:"omitempty,email,max=255"`
	Website       string    `json:"website" validate:"omitempty,url,max=500"`
	ImageURL      string    `json:"image_url" validate:"omitempty,url,max=500"`
	IsFeatured    *bool     `json:"is_featured,omitempty"`
	IsActive      *bool     `json:"is_active,omitempty"`
	LocationID    uuid.UUID `json:"location_id" validate:"required"`
	TheatreTypeID uuid.UUID `json:"theatre_type_id" validate:"required"`
}

// TheatreDetails contains detailed theatre information including relationships
type TheatreDetails struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Capacity      int       `json:"capacity"`
	Address       string    `json:"address"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	Website       string    `json:"website"`
	ImageURL      string    `json:"image_url"`
	IsFeatured    bool      `json:"is_featured"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	LocationID    uuid.UUID `json:"location_id"`
	TheatreTypeID uuid.UUID `json:"theatre_type_id"`

	// Relationships
	Location    LocationSummary    `json:"location"`
	TheatreType TheatreTypeSummary `json:"theatre_type"`
	Shows       []ShowSummary      `json:"shows,omitempty"`
}

// TheatreSummary contains summary theatre information for lists
type TheatreSummary struct {
	ID            uuid.UUID          `json:"id"`
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	Capacity      int                `json:"capacity"`
	ImageURL      string             `json:"image_url"`
	IsFeatured    bool               `json:"is_featured"`
	IsActive      bool               `json:"is_active"`
	LocationID    uuid.UUID          `json:"location_id"`
	TheatreTypeID uuid.UUID          `json:"theatre_type_id"`
	Location      LocationSummary    `json:"location"`
	TheatreType   TheatreTypeSummary `json:"theatre_type"`
}
