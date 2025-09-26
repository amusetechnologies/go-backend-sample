package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Theatre represents a theatre venue
type Theatre struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string         `json:"name" gorm:"type:varchar(255);not null" validate:"required,min=1,max=255"`
	Description string         `json:"description" gorm:"type:text" validate:"max=2000"`
	Capacity    int            `json:"capacity" gorm:"type:integer" validate:"omitempty,min=1,max=100000"`
	Address     string         `json:"address" gorm:"type:text" validate:"max=500"`
	Phone       string         `json:"phone" gorm:"type:varchar(20)" validate:"max=20"`
	Email       string         `json:"email" gorm:"type:varchar(255)" validate:"omitempty,email,max=255"`
	Website     string         `json:"website" gorm:"type:varchar(500)" validate:"omitempty,url,max=500"`
	ImageURL    string         `json:"image_url" gorm:"type:varchar(500)" validate:"omitempty,url,max=500"`
	IsFeatured  bool           `json:"is_featured" gorm:"default:false"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Foreign Keys
	LocationID    uuid.UUID `json:"location_id" gorm:"type:uuid;not null" validate:"required"`
	TheatreTypeID uuid.UUID `json:"theatre_type_id" gorm:"type:uuid;not null" validate:"required"`

	// Relationships
	Location    Location    `json:"location" gorm:"foreignKey:LocationID"`
	TheatreType TheatreType `json:"theatre_type" gorm:"foreignKey:TheatreTypeID"`
	Shows       []Show      `json:"shows,omitempty" gorm:"foreignKey:TheatreID"`
}

// BeforeCreate hook to generate UUID if not set
func (t *Theatre) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}
