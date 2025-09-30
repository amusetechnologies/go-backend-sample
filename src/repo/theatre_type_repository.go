package repo

import (
	"theatre-management-system/src/interfaces"
	"theatre-management-system/src/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// theatreTypeRepository implements the TheatreTypeRepository interface
type theatreTypeRepository struct {
	db *gorm.DB
}

// NewTheatreTypeRepository creates a new theatre type repository
func NewTheatreTypeRepository(db *gorm.DB) interfaces.TheatreTypeRepository {
	return &theatreTypeRepository{db: db}
}

// Create creates a new theatre type
func (r *theatreTypeRepository) Create(theatreType *models.TheatreType) error {
	return r.db.Create(theatreType).Error
}

// GetByID retrieves a theatre type by ID
func (r *theatreTypeRepository) GetByID(id uuid.UUID) (*models.TheatreType, error) {
	var theatreType models.TheatreType
	err := r.db.Preload("Theatres").First(&theatreType, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &theatreType, nil
}

// GetAll retrieves all theatre types with pagination
func (r *theatreTypeRepository) GetAll(limit, offset int) ([]*models.TheatreType, error) {
	var theatreTypes []*models.TheatreType
	err := r.db.Limit(limit).Offset(offset).Find(&theatreTypes).Error
	if err != nil {
		return nil, err
	}
	return theatreTypes, nil
}

// Update updates an existing theatre type
func (r *theatreTypeRepository) Update(theatreType *models.TheatreType) error {
	return r.db.Save(theatreType).Error
}

// Delete soft deletes a theatre type
func (r *theatreTypeRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.TheatreType{}, "id = ?", id).Error
}

// GetByName retrieves a theatre type by name
func (r *theatreTypeRepository) GetByName(name string) (*models.TheatreType, error) {
	var theatreType models.TheatreType
	err := r.db.Where("name = ?", name).First(&theatreType).Error
	if err != nil {
		return nil, err
	}
	return &theatreType, nil
}

// GetActiveTypes retrieves all active theatre types
func (r *theatreTypeRepository) GetActiveTypes() ([]*models.TheatreType, error) {
	var theatreTypes []*models.TheatreType
	err := r.db.Where("is_active = ?", true).Find(&theatreTypes).Error
	if err != nil {
		return nil, err
	}
	return theatreTypes, nil
}
