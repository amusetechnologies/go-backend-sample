package repo

import (
	"fmt"
	"theatre-management-system/src/interfaces"
	"theatre-management-system/src/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// locationRepository implements the LocationRepository interface
type locationRepository struct {
	db *gorm.DB
}

// NewLocationRepository creates a new location repository
func NewLocationRepository(db *gorm.DB) interfaces.LocationRepository {
	return &locationRepository{db: db}
}

// Create creates a new location
func (r *locationRepository) Create(location *models.Location) error {
	return r.db.Create(location).Error
}

// GetByID retrieves a location by ID
func (r *locationRepository) GetByID(id uuid.UUID) (*models.Location, error) {
	var location models.Location
	err := r.db.Preload("Theatres").First(&location, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &location, nil
}

// GetAll retrieves all locations with pagination
func (r *locationRepository) GetAll(limit, offset int) ([]*models.Location, error) {
	var locations []*models.Location
	err := r.db.Limit(limit).Offset(offset).Find(&locations).Error
	if err != nil {
		return nil, err
	}
	return locations, nil
}

// Update updates an existing location
func (r *locationRepository) Update(location *models.Location) error {
	return r.db.Save(location).Error
}

// Delete soft deletes a location
func (r *locationRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Location{}, "id = ?", id).Error
}

// GetByCoordinates finds locations within a radius (in kilometers) of given coordinates
func (r *locationRepository) GetByCoordinates(latitude, longitude, radius float64) ([]*models.Location, error) {
	var locations []*models.Location

	// Using PostGIS ST_DWithin function for geographic queries
	// ST_DWithin uses meters, so convert radius from km to meters
	radiusMeters := radius * 1000

	query := `
		SELECT * FROM locations 
		WHERE latitude IS NOT NULL 
		AND longitude IS NOT NULL 
		AND ST_DWithin(
			ST_MakePoint(longitude, latitude)::geography,
			ST_MakePoint(?, ?)::geography,
			?
		)
		AND deleted_at IS NULL
	`

	err := r.db.Raw(query, longitude, latitude, radiusMeters).Scan(&locations).Error
	if err != nil {
		return nil, err
	}

	return locations, nil
}

// GetActiveLocations retrieves all active locations
func (r *locationRepository) GetActiveLocations() ([]*models.Location, error) {
	var locations []*models.Location
	err := r.db.Where("is_active = ?", true).Find(&locations).Error
	if err != nil {
		return nil, err
	}
	return locations, nil
}

// Search searches locations by name, city, or country
func (r *locationRepository) Search(query string) ([]*models.Location, error) {
	var locations []*models.Location
	searchPattern := fmt.Sprintf("%%%s%%", query)

	err := r.db.Where(
		"name ILIKE ? OR city ILIKE ? OR country ILIKE ?",
		searchPattern, searchPattern, searchPattern,
	).Find(&locations).Error

	if err != nil {
		return nil, err
	}

	return locations, nil
}
