package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TheatreType represents different types of theatres (Broadway, Off-Broadway, Regional, etc.)
type TheatreType struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string         `json:"name" gorm:"type:varchar(100);not null;uniqueIndex" validate:"required,min=1,max=100"`
	Description string         `json:"description" gorm:"type:text" validate:"max=1000"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Relationships
	Theatres []Theatre `json:"theatres,omitempty" gorm:"foreignKey:TheatreTypeID"`
}

// BeforeCreate hook to generate UUID if not set
func (tt *TheatreType) BeforeCreate(tx *gorm.DB) error {
	if tt.ID == uuid.Nil {
		tt.ID = uuid.New()
	}
	return nil
}
