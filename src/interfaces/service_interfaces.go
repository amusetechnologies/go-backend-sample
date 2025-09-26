package interfaces

import (
	"theatre-management-system/src/dto"

	"github.com/google/uuid"
)

// LocationService defines the interface for location business logic
type LocationService interface {
	CreateLocation(location *dto.LocationBase) (*dto.LocationDetails, error)
	GetLocationByID(id uuid.UUID) (*dto.LocationDetails, error)
	GetAllLocations(limit, offset int) ([]*dto.LocationSummary, error)
	UpdateLocation(id uuid.UUID, location *dto.LocationBase) (*dto.LocationDetails, error)
	DeleteLocation(id uuid.UUID) error
	GetLocationsByCoordinates(latitude, longitude, radius float64) ([]*dto.LocationSummary, error)
	GetActiveLocations() ([]*dto.LocationSummary, error)
	SearchLocations(query string) ([]*dto.LocationSummary, error)
}

// TheatreTypeService defines the interface for theatre type business logic
type TheatreTypeService interface {
	CreateTheatreType(theatreType *dto.TheatreTypeBase) (*dto.TheatreTypeDetails, error)
	GetTheatreTypeByID(id uuid.UUID) (*dto.TheatreTypeDetails, error)
	GetAllTheatreTypes(limit, offset int) ([]*dto.TheatreTypeSummary, error)
	UpdateTheatreType(id uuid.UUID, theatreType *dto.TheatreTypeBase) (*dto.TheatreTypeDetails, error)
	DeleteTheatreType(id uuid.UUID) error
	GetTheatreTypeByName(name string) (*dto.TheatreTypeDetails, error)
	GetActiveTheatreTypes() ([]*dto.TheatreTypeSummary, error)
}

// ShowTypeService defines the interface for show type business logic
type ShowTypeService interface {
	CreateShowType(showType *dto.ShowTypeBase) (*dto.ShowTypeDetails, error)
	GetShowTypeByID(id uuid.UUID) (*dto.ShowTypeDetails, error)
	GetAllShowTypes(limit, offset int) ([]*dto.ShowTypeSummary, error)
	UpdateShowType(id uuid.UUID, showType *dto.ShowTypeBase) (*dto.ShowTypeDetails, error)
	DeleteShowType(id uuid.UUID) error
	GetShowTypeByName(name string) (*dto.ShowTypeDetails, error)
	GetActiveShowTypes() ([]*dto.ShowTypeSummary, error)
}

// TheatreService defines the interface for theatre business logic
type TheatreService interface {
	CreateTheatre(theatre *dto.TheatreBase) (*dto.TheatreDetails, error)
	GetTheatreByID(id uuid.UUID) (*dto.TheatreDetails, error)
	GetAllTheatres(limit, offset int) ([]*dto.TheatreSummary, error)
	UpdateTheatre(id uuid.UUID, theatre *dto.TheatreBase) (*dto.TheatreDetails, error)
	DeleteTheatre(id uuid.UUID) error
	GetTheatresByLocationID(locationID uuid.UUID) ([]*dto.TheatreSummary, error)
	GetTheatresByTheatreTypeID(theatreTypeID uuid.UUID) ([]*dto.TheatreSummary, error)
	GetFeaturedTheatres() ([]*dto.TheatreSummary, error)
	GetActiveTheatres() ([]*dto.TheatreSummary, error)
	SearchTheatres(query string) ([]*dto.TheatreSummary, error)
	GetNearbyTheatres(latitude, longitude, radius float64) ([]*dto.TheatreSummary, error)
}

// ShowService defines the interface for show business logic
type ShowService interface {
	CreateShow(show *dto.ShowBase) (*dto.ShowDetails, error)
	GetShowByID(id uuid.UUID) (*dto.ShowDetails, error)
	GetAllShows(limit, offset int) ([]*dto.ShowSummary, error)
	UpdateShow(id uuid.UUID, show *dto.ShowBase) (*dto.ShowDetails, error)
	DeleteShow(id uuid.UUID) error
	GetShowsByTheatreID(theatreID uuid.UUID) ([]*dto.ShowSummary, error)
	GetShowsByShowTypeID(showTypeID uuid.UUID) ([]*dto.ShowSummary, error)
	GetFeaturedShows() ([]*dto.ShowSummary, error)
	GetActiveShows() ([]*dto.ShowSummary, error)
	GetCurrentShows() ([]*dto.ShowSummary, error)
	GetUpcomingShows() ([]*dto.ShowSummary, error)
	SearchShows(query string) ([]*dto.ShowSummary, error)
}
