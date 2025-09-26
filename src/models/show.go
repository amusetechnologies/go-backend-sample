package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Show represents a show/performance at a theatre
type Show struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title       string         `json:"title" gorm:"type:varchar(255);not null" validate:"required,min=1,max=255"`
	Description string         `json:"description" gorm:"type:text" validate:"max=2000"`
	Director    string         `json:"director" gorm:"type:varchar(255)" validate:"max=255"`
	Cast        string         `json:"cast" gorm:"type:text" validate:"max=1000"`
	Duration    int            `json:"duration" gorm:"type:integer" validate:"omitempty,min=1,max=600"` // Duration in minutes
	StartDate   *time.Time     `json:"start_date" gorm:"type:date"`
	EndDate     *time.Time     `json:"end_date" gorm:"type:date"`
	Price       float64        `json:"price" gorm:"type:decimal(10,2)" validate:"omitempty,min=0"`
	ImageURL    string         `json:"image_url" gorm:"type:varchar(500)" validate:"omitempty,url,max=500"`
	TrailerURL  string         `json:"trailer_url" gorm:"type:varchar(500)" validate:"omitempty,url,max=500"`
	IsFeatured  bool           `json:"is_featured" gorm:"default:false"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Foreign Keys
	TheatreID  uuid.UUID `json:"theatre_id" gorm:"type:uuid;not null" validate:"required"`
	ShowTypeID uuid.UUID `json:"show_type_id" gorm:"type:uuid;not null" validate:"required"`

	// Relationships
	Theatre  Theatre  `json:"theatre" gorm:"foreignKey:TheatreID"`
	ShowType ShowType `json:"show_type" gorm:"foreignKey:ShowTypeID"`
}

// BeforeCreate hook to generate UUID if not set
func (s *Show) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}
