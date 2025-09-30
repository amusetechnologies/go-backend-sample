package repo

import (
	"fmt"
	"theatre-management-system/src/interfaces"
	"theatre-management-system/src/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// theatreRepository implements the TheatreRepository interface
type theatreRepository struct {
	db *gorm.DB
}

// NewTheatreRepository creates a new theatre repository
func NewTheatreRepository(db *gorm.DB) interfaces.TheatreRepository {
	return &theatreRepository{db: db}
}

// Create creates a new theatre
func (r *theatreRepository) Create(theatre *models.Theatre) error {
	return r.db.Create(theatre).Error
}

// GetByID retrieves a theatre by ID with all relationships
func (r *theatreRepository) GetByID(id uuid.UUID) (*models.Theatre, error) {
	var theatre models.Theatre
	err := r.db.Preload("Location").Preload("TheatreType").Preload("Shows").First(&theatre, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &theatre, nil
}

// GetAll retrieves all theatres with pagination
func (r *theatreRepository) GetAll(limit, offset int) ([]*models.Theatre, error) {
	var theatres []*models.Theatre
	err := r.db.Preload("Location").Preload("TheatreType").Limit(limit).Offset(offset).Find(&theatres).Error
	if err != nil {
		return nil, err
	}
	return theatres, nil
}

// Update updates an existing theatre
func (r *theatreRepository) Update(theatre *models.Theatre) error {
	return r.db.Save(theatre).Error
}

// Delete soft deletes a theatre
func (r *theatreRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Theatre{}, "id = ?", id).Error
}

// GetByLocationID retrieves theatres by location ID
func (r *theatreRepository) GetByLocationID(locationID uuid.UUID) ([]*models.Theatre, error) {
	var theatres []*models.Theatre
	err := r.db.Preload("Location").Preload("TheatreType").Where("location_id = ?", locationID).Find(&theatres).Error
	if err != nil {
		return nil, err
	}
	return theatres, nil
}

// GetByTheatreTypeID retrieves theatres by theatre type ID
func (r *theatreRepository) GetByTheatreTypeID(theatreTypeID uuid.UUID) ([]*models.Theatre, error) {
	var theatres []*models.Theatre
	err := r.db.Preload("Location").Preload("TheatreType").Where("theatre_type_id = ?", theatreTypeID).Find(&theatres).Error
	if err != nil {
		return nil, err
	}
	return theatres, nil
}

// GetFeaturedTheatres retrieves all featured theatres
func (r *theatreRepository) GetFeaturedTheatres() ([]*models.Theatre, error) {
	var theatres []*models.Theatre
	err := r.db.Preload("Location").Preload("TheatreType").Where("is_featured = ?", true).Find(&theatres).Error
	if err != nil {
		return nil, err
	}
	return theatres, nil
}

// GetActiveTheatres retrieves all active theatres
func (r *theatreRepository) GetActiveTheatres() ([]*models.Theatre, error) {
	var theatres []*models.Theatre
	err := r.db.Preload("Location").Preload("TheatreType").Where("is_active = ?", true).Find(&theatres).Error
	if err != nil {
		return nil, err
	}
	return theatres, nil
}

// Search searches theatres by name or description
func (r *theatreRepository) Search(query string) ([]*models.Theatre, error) {
	var theatres []*models.Theatre
	searchPattern := fmt.Sprintf("%%%s%%", query)

	err := r.db.Preload("Location").Preload("TheatreType").Where(
		"name ILIKE ? OR description ILIKE ?",
		searchPattern, searchPattern,
	).Find(&theatres).Error

	if err != nil {
		return nil, err
	}

	return theatres, nil
}

// GetNearbyTheatres finds theatres within a radius (in kilometers) of given coordinates
func (r *theatreRepository) GetNearbyTheatres(latitude, longitude, radius float64) ([]*models.Theatre, error) {
	var theatres []*models.Theatre

	// Using PostGIS ST_DWithin function for geographic queries
	// ST_DWithin uses meters, so convert radius from km to meters
	radiusMeters := radius * 1000

	query := `
		SELECT t.* FROM theatres t
		INNER JOIN locations l ON t.location_id = l.id
		WHERE l.latitude IS NOT NULL 
		AND l.longitude IS NOT NULL 
		AND ST_DWithin(
			ST_MakePoint(l.longitude, l.latitude)::geography,
			ST_MakePoint(?, ?)::geography,
			?
		)
		AND t.deleted_at IS NULL
		AND l.deleted_at IS NULL
	`

	err := r.db.Preload("Location").Preload("TheatreType").Raw(query, longitude, latitude, radiusMeters).Scan(&theatres).Error
	if err != nil {
		return nil, err
	}

	return theatres, nil
}
