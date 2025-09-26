package dto

import (
	"time"

	"github.com/google/uuid"
)

// TheatreTypeBase contains basic theatre type information for creation/updates
type TheatreTypeBase struct {
	Name        string `json:"name" validate:"required,min=1,max=100"`
	Description string `json:"description" validate:"max=1000"`
	IsActive    *bool  `json:"is_active,omitempty"`
}

// TheatreTypeDetails contains detailed theatre type information including relationships
type TheatreTypeDetails struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relationships
	Theatres []TheatreSummary `json:"theatres,omitempty"`
}

// TheatreTypeSummary contains summary theatre type information for lists
type TheatreTypeSummary struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
}
