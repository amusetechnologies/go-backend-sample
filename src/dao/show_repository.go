package dao

import (
	"fmt"
	"theatre-management-system/src/interfaces"
	"theatre-management-system/src/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// showRepository implements the ShowRepository interface
type showRepository struct {
	db *gorm.DB
}

// NewShowRepository creates a new show repository
func NewShowRepository(db *gorm.DB) interfaces.ShowRepository {
	return &showRepository{db: db}
}

// Create creates a new show
func (r *showRepository) Create(show *models.Show) error {
	return r.db.Create(show).Error
}

// GetByID retrieves a show by ID with all relationships
func (r *showRepository) GetByID(id uuid.UUID) (*models.Show, error) {
	var show models.Show
	err := r.db.Preload("Theatre").Preload("Theatre.Location").Preload("ShowType").First(&show, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &show, nil
}

// GetAll retrieves all shows with pagination
func (r *showRepository) GetAll(limit, offset int) ([]*models.Show, error) {
	var shows []*models.Show
	err := r.db.Preload("Theatre").Preload("Theatre.Location").Preload("ShowType").Limit(limit).Offset(offset).Find(&shows).Error
	if err != nil {
		return nil, err
	}
	return shows, nil
}

// Update updates an existing show
func (r *showRepository) Update(show *models.Show) error {
	return r.db.Save(show).Error
}

// Delete soft deletes a show
func (r *showRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Show{}, "id = ?", id).Error
}

// GetByTheatreID retrieves shows by theatre ID
func (r *showRepository) GetByTheatreID(theatreID uuid.UUID) ([]*models.Show, error) {
	var shows []*models.Show
	err := r.db.Preload("Theatre").Preload("ShowType").Where("theatre_id = ?", theatreID).Find(&shows).Error
	if err != nil {
		return nil, err
	}
	return shows, nil
}

// GetByShowTypeID retrieves shows by show type ID
func (r *showRepository) GetByShowTypeID(showTypeID uuid.UUID) ([]*models.Show, error) {
	var shows []*models.Show
	err := r.db.Preload("Theatre").Preload("ShowType").Where("show_type_id = ?", showTypeID).Find(&shows).Error
	if err != nil {
		return nil, err
	}
	return shows, nil
}

// GetFeaturedShows retrieves all featured shows
func (r *showRepository) GetFeaturedShows() ([]*models.Show, error) {
	var shows []*models.Show
	err := r.db.Preload("Theatre").Preload("ShowType").Where("is_featured = ?", true).Find(&shows).Error
	if err != nil {
		return nil, err
	}
	return shows, nil
}

// GetActiveShows retrieves all active shows
func (r *showRepository) GetActiveShows() ([]*models.Show, error) {
	var shows []*models.Show
	err := r.db.Preload("Theatre").Preload("ShowType").Where("is_active = ?", true).Find(&shows).Error
	if err != nil {
		return nil, err
	}
	return shows, nil
}

// GetCurrentShows retrieves shows that are currently running
func (r *showRepository) GetCurrentShows() ([]*models.Show, error) {
	var shows []*models.Show
	now := time.Now()

	err := r.db.Preload("Theatre").Preload("ShowType").Where(
		"is_active = ? AND start_date <= ? AND (end_date IS NULL OR end_date >= ?)",
		true, now, now,
	).Find(&shows).Error

	if err != nil {
		return nil, err
	}
	return shows, nil
}

// GetUpcomingShows retrieves shows that will start in the future
func (r *showRepository) GetUpcomingShows() ([]*models.Show, error) {
	var shows []*models.Show
	now := time.Now()

	err := r.db.Preload("Theatre").Preload("ShowType").Where(
		"is_active = ? AND start_date > ?",
		true, now,
	).Find(&shows).Error

	if err != nil {
		return nil, err
	}
	return shows, nil
}

// Search searches shows by title, description, director, or cast
func (r *showRepository) Search(query string) ([]*models.Show, error) {
	var shows []*models.Show
	searchPattern := fmt.Sprintf("%%%s%%", query)

	err := r.db.Preload("Theatre").Preload("ShowType").Where(
		"title ILIKE ? OR description ILIKE ? OR director ILIKE ? OR cast ILIKE ?",
		searchPattern, searchPattern, searchPattern, searchPattern,
	).Find(&shows).Error

	if err != nil {
		return nil, err
	}

	return shows, nil
}
