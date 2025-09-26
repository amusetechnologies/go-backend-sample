package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ShowType represents different types of shows (Musical, Opera, Concert, Play, etc.)
type ShowType struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string         `json:"name" gorm:"type:varchar(100);not null;uniqueIndex" validate:"required,min=1,max=100"`
	Description string         `json:"description" gorm:"type:text" validate:"max=1000"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Relationships
	Shows []Show `json:"shows,omitempty" gorm:"foreignKey:ShowTypeID"`
}

// BeforeCreate hook to generate UUID if not set
func (st *ShowType) BeforeCreate(tx *gorm.DB) error {
	if st.ID == uuid.Nil {
		st.ID = uuid.New()
	}
	return nil
}
