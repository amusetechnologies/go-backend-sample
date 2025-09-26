package interfaces

import (
	"theatre-management-system/src/models"

	"github.com/google/uuid"
)

// LocationRepository defines the interface for location data access
type LocationRepository interface {
	Create(location *models.Location) error
	GetByID(id uuid.UUID) (*models.Location, error)
	GetAll(limit, offset int) ([]*models.Location, error)
	Update(location *models.Location) error
	Delete(id uuid.UUID) error
	GetByCoordinates(latitude, longitude, radius float64) ([]*models.Location, error)
	GetActiveLocations() ([]*models.Location, error)
	Search(query string) ([]*models.Location, error)
}

// TheatreTypeRepository defines the interface for theatre type data access
type TheatreTypeRepository interface {
	Create(theatreType *models.TheatreType) error
	GetByID(id uuid.UUID) (*models.TheatreType, error)
	GetAll(limit, offset int) ([]*models.TheatreType, error)
	Update(theatreType *models.TheatreType) error
	Delete(id uuid.UUID) error
	GetByName(name string) (*models.TheatreType, error)
	GetActiveTypes() ([]*models.TheatreType, error)
}

// ShowTypeRepository defines the interface for show type data access
type ShowTypeRepository interface {
	Create(showType *models.ShowType) error
	GetByID(id uuid.UUID) (*models.ShowType, error)
	GetAll(limit, offset int) ([]*models.ShowType, error)
	Update(showType *models.ShowType) error
	Delete(id uuid.UUID) error
	GetByName(name string) (*models.ShowType, error)
	GetActiveTypes() ([]*models.ShowType, error)
}

// TheatreRepository defines the interface for theatre data access
type TheatreRepository interface {
	Create(theatre *models.Theatre) error
	GetByID(id uuid.UUID) (*models.Theatre, error)
	GetAll(limit, offset int) ([]*models.Theatre, error)
	Update(theatre *models.Theatre) error
	Delete(id uuid.UUID) error
	GetByLocationID(locationID uuid.UUID) ([]*models.Theatre, error)
	GetByTheatreTypeID(theatreTypeID uuid.UUID) ([]*models.Theatre, error)
	GetFeaturedTheatres() ([]*models.Theatre, error)
	GetActiveTheatres() ([]*models.Theatre, error)
	Search(query string) ([]*models.Theatre, error)
	GetNearbyTheatres(latitude, longitude, radius float64) ([]*models.Theatre, error)
}

// ShowRepository defines the interface for show data access
type ShowRepository interface {
	Create(show *models.Show) error
	GetByID(id uuid.UUID) (*models.Show, error)
	GetAll(limit, offset int) ([]*models.Show, error)
	Update(show *models.Show) error
	Delete(id uuid.UUID) error
	GetByTheatreID(theatreID uuid.UUID) ([]*models.Show, error)
	GetByShowTypeID(showTypeID uuid.UUID) ([]*models.Show, error)
	GetFeaturedShows() ([]*models.Show, error)
	GetActiveShows() ([]*models.Show, error)
	GetCurrentShows() ([]*models.Show, error)
	GetUpcomingShows() ([]*models.Show, error)
	Search(query string) ([]*models.Show, error)
}
