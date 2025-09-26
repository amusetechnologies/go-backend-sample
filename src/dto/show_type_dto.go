package dto

import (
	"time"

	"github.com/google/uuid"
)

// ShowTypeBase contains basic show type information for creation/updates
type ShowTypeBase struct {
	Name        string `json:"name" validate:"required,min=1,max=100"`
	Description string `json:"description" validate:"max=1000"`
	IsActive    *bool  `json:"is_active,omitempty"`
}

// ShowTypeDetails contains detailed show type information including relationships
type ShowTypeDetails struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relationships
	Shows []ShowSummary `json:"shows,omitempty"`
}

// ShowTypeSummary contains summary show type information for lists
type ShowTypeSummary struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
}
