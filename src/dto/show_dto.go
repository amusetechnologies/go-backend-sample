package dto

import (
	"time"

	"github.com/google/uuid"
)

// ShowBase contains basic show information for creation/updates
type ShowBase struct {
	Title       string     `json:"title" validate:"required,min=1,max=255"`
	Description string     `json:"description" validate:"max=2000"`
	Director    string     `json:"director" validate:"max=255"`
	Cast        string     `json:"cast" validate:"max=1000"`
	Duration    *int       `json:"duration" validate:"omitempty,min=1,max=600"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Price       *float64   `json:"price" validate:"omitempty,min=0"`
	ImageURL    string     `json:"image_url" validate:"omitempty,url,max=500"`
	TrailerURL  string     `json:"trailer_url" validate:"omitempty,url,max=500"`
	IsFeatured  *bool      `json:"is_featured,omitempty"`
	IsActive    *bool      `json:"is_active,omitempty"`
	TheatreID   uuid.UUID  `json:"theatre_id" validate:"required"`
	ShowTypeID  uuid.UUID  `json:"show_type_id" validate:"required"`
}

// ShowDetails contains detailed show information including relationships
type ShowDetails struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Director    string     `json:"director"`
	Cast        string     `json:"cast"`
	Duration    int        `json:"duration"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Price       float64    `json:"price"`
	ImageURL    string     `json:"image_url"`
	TrailerURL  string     `json:"trailer_url"`
	IsFeatured  bool       `json:"is_featured"`
	IsActive    bool       `json:"is_active"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	TheatreID   uuid.UUID  `json:"theatre_id"`
	ShowTypeID  uuid.UUID  `json:"show_type_id"`

	// Relationships
	Theatre  TheatreSummary  `json:"theatre"`
	ShowType ShowTypeSummary `json:"show_type"`
}

// ShowSummary contains summary show information for lists
type ShowSummary struct {
	ID          uuid.UUID       `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Director    string          `json:"director"`
	Duration    int             `json:"duration"`
	StartDate   *time.Time      `json:"start_date"`
	EndDate     *time.Time      `json:"end_date"`
	Price       float64         `json:"price"`
	ImageURL    string          `json:"image_url"`
	IsFeatured  bool            `json:"is_featured"`
	IsActive    bool            `json:"is_active"`
	TheatreID   uuid.UUID       `json:"theatre_id"`
	ShowTypeID  uuid.UUID       `json:"show_type_id"`
	Theatre     TheatreSummary  `json:"theatre"`
	ShowType    ShowTypeSummary `json:"show_type"`
}
