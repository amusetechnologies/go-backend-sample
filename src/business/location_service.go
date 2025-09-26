package business

import (
	"errors"
	"theatre-management-system/src/constants"
	"theatre-management-system/src/dto"
	"theatre-management-system/src/interfaces"
	"theatre-management-system/src/mappers"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// locationService implements the LocationService interface
type locationService struct {
	locationRepo interfaces.LocationRepository
	mapper       *mappers.LocationMapper
	validator    *validator.Validate
}

// NewLocationService creates a new location service
func NewLocationService(locationRepo interfaces.LocationRepository) interfaces.LocationService {
	return &locationService{
		locationRepo: locationRepo,
		mapper:       mappers.NewLocationMapper(),
		validator:    validator.New(),
	}
}

// CreateLocation creates a new location
func (s *locationService) CreateLocation(locationDTO *dto.LocationBase) (*dto.LocationDetails, error) {
	// Validate input
	if err := s.validator.Struct(locationDTO); err != nil {
		return nil, errors.New(constants.ErrorValidationFailed + ": " + err.Error())
	}

	// Convert DTO to model
	location := s.mapper.ToModel(locationDTO)

	// Create in database
	if err := s.locationRepo.Create(location); err != nil {
		return nil, err
	}

	// Return created location
	return s.mapper.ToDetailsDTO(location), nil
}

// GetLocationByID retrieves a location by ID
func (s *locationService) GetLocationByID(id uuid.UUID) (*dto.LocationDetails, error) {
	location, err := s.locationRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.ErrorLocationNotFound)
		}
		return nil, err
	}

	return s.mapper.ToDetailsDTO(location), nil
}

// GetAllLocations retrieves all locations with pagination
func (s *locationService) GetAllLocations(limit, offset int) ([]*dto.LocationSummary, error) {
	// Apply default and max limits
	if limit <= 0 || limit > constants.MaxLimit {
		limit = constants.DefaultLimit
	}
	if offset < 0 {
		offset = constants.DefaultOffset
	}

	locations, err := s.locationRepo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(locations), nil
}

// UpdateLocation updates an existing location
func (s *locationService) UpdateLocation(id uuid.UUID, locationDTO *dto.LocationBase) (*dto.LocationDetails, error) {
	// Validate input
	if err := s.validator.Struct(locationDTO); err != nil {
		return nil, errors.New(constants.ErrorValidationFailed + ": " + err.Error())
	}

	// Get existing location
	location, err := s.locationRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.ErrorLocationNotFound)
		}
		return nil, err
	}

	// Update model with new data
	s.mapper.UpdateModel(location, locationDTO)

	// Save to database
	if err := s.locationRepo.Update(location); err != nil {
		return nil, err
	}

	return s.mapper.ToDetailsDTO(location), nil
}

// DeleteLocation soft deletes a location
func (s *locationService) DeleteLocation(id uuid.UUID) error {
	// Check if location exists
	_, err := s.locationRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constants.ErrorLocationNotFound)
		}
		return err
	}

	return s.locationRepo.Delete(id)
}

// GetLocationsByCoordinates finds locations within a radius of given coordinates
func (s *locationService) GetLocationsByCoordinates(latitude, longitude, radius float64) ([]*dto.LocationSummary, error) {
	// Validate coordinates
	if latitude < -90 || latitude > 90 {
		return nil, errors.New("invalid latitude: must be between -90 and 90")
	}
	if longitude < -180 || longitude > 180 {
		return nil, errors.New("invalid longitude: must be between -180 and 180")
	}
	if radius <= 0 {
		radius = constants.DefaultRadius
	}

	locations, err := s.locationRepo.GetByCoordinates(latitude, longitude, radius)
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(locations), nil
}

// GetActiveLocations retrieves all active locations
func (s *locationService) GetActiveLocations() ([]*dto.LocationSummary, error) {
	locations, err := s.locationRepo.GetActiveLocations()
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(locations), nil
}

// SearchLocations searches locations by query string
func (s *locationService) SearchLocations(query string) ([]*dto.LocationSummary, error) {
	if query == "" {
		return []*dto.LocationSummary{}, nil
	}

	locations, err := s.locationRepo.Search(query)
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(locations), nil
}
