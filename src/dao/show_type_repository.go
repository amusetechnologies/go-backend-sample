package dao

import (
	"theatre-management-system/src/interfaces"
	"theatre-management-system/src/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// showTypeRepository implements the ShowTypeRepository interface
type showTypeRepository struct {
	db *gorm.DB
}

// NewShowTypeRepository creates a new show type repository
func NewShowTypeRepository(db *gorm.DB) interfaces.ShowTypeRepository {
	return &showTypeRepository{db: db}
}

// Create creates a new show type
func (r *showTypeRepository) Create(showType *models.ShowType) error {
	return r.db.Create(showType).Error
}

// GetByID retrieves a show type by ID
func (r *showTypeRepository) GetByID(id uuid.UUID) (*models.ShowType, error) {
	var showType models.ShowType
	err := r.db.Preload("Shows").First(&showType, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &showType, nil
}

// GetAll retrieves all show types with pagination
func (r *showTypeRepository) GetAll(limit, offset int) ([]*models.ShowType, error) {
	var showTypes []*models.ShowType
	err := r.db.Limit(limit).Offset(offset).Find(&showTypes).Error
	if err != nil {
		return nil, err
	}
	return showTypes, nil
}

// Update updates an existing show type
func (r *showTypeRepository) Update(showType *models.ShowType) error {
	return r.db.Save(showType).Error
}

// Delete soft deletes a show type
func (r *showTypeRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.ShowType{}, "id = ?", id).Error
}

// GetByName retrieves a show type by name
func (r *showTypeRepository) GetByName(name string) (*models.ShowType, error) {
	var showType models.ShowType
	err := r.db.Where("name = ?", name).First(&showType).Error
	if err != nil {
		return nil, err
	}
	return &showType, nil
}

// GetActiveTypes retrieves all active show types
func (r *showTypeRepository) GetActiveTypes() ([]*models.ShowType, error) {
	var showTypes []*models.ShowType
	err := r.db.Where("is_active = ?", true).Find(&showTypes).Error
	if err != nil {
		return nil, err
	}
	return showTypes, nil
}
