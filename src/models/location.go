package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Location represents a geographic location for theatres
type Location struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string         `json:"name" gorm:"type:varchar(255);not null" validate:"required,min=1,max=255"`
	City        string         `json:"city" gorm:"type:varchar(100);not null" validate:"required,min=1,max=100"`
	State       string         `json:"state" gorm:"type:varchar(100)" validate:"max=100"`
	Country     string         `json:"country" gorm:"type:varchar(100);not null" validate:"required,min=1,max=100"`
	Latitude    *float64       `json:"latitude" gorm:"type:double precision" validate:"omitempty,min=-90,max=90"`
	Longitude   *float64       `json:"longitude" gorm:"type:double precision" validate:"omitempty,min=-180,max=180"`
	PostalCode  string         `json:"postal_code" gorm:"type:varchar(20)" validate:"max=20"`
	Address     string         `json:"address" gorm:"type:text" validate:"max=500"`
	Description string         `json:"description" gorm:"type:text" validate:"max=1000"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Relationships
	Theatres []Theatre `json:"theatres,omitempty" gorm:"foreignKey:LocationID"`
}

// BeforeCreate hook to generate UUID if not set
func (l *Location) BeforeCreate(tx *gorm.DB) error {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	return nil
}
