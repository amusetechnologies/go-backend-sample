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

// theatreService implements the TheatreService interface
type theatreService struct {
	theatreRepo     interfaces.TheatreRepository
	locationRepo    interfaces.LocationRepository
	theatreTypeRepo interfaces.TheatreTypeRepository
	mapper          *mappers.TheatreMapper
	validator       *validator.Validate
}

// NewTheatreService creates a new theatre service
func NewTheatreService(
	theatreRepo interfaces.TheatreRepository,
	locationRepo interfaces.LocationRepository,
	theatreTypeRepo interfaces.TheatreTypeRepository,
) interfaces.TheatreService {
	return &theatreService{
		theatreRepo:     theatreRepo,
		locationRepo:    locationRepo,
		theatreTypeRepo: theatreTypeRepo,
		mapper:          mappers.NewTheatreMapper(),
		validator:       validator.New(),
	}
}

// CreateTheatre creates a new theatre
func (s *theatreService) CreateTheatre(theatreDTO *dto.TheatreBase) (*dto.TheatreDetails, error) {
	// Validate input
	if err := s.validator.Struct(theatreDTO); err != nil {
		return nil, errors.New(constants.ErrorValidationFailed + ": " + err.Error())
	}

	// Validate foreign key relationships
	if err := s.validateRelationships(theatreDTO.LocationID, theatreDTO.TheatreTypeID); err != nil {
		return nil, err
	}

	// Convert DTO to model
	theatre := s.mapper.ToModel(theatreDTO)

	// Create in database
	if err := s.theatreRepo.Create(theatre); err != nil {
		return nil, err
	}

	// Get created theatre with relationships
	createdTheatre, err := s.theatreRepo.GetByID(theatre.ID)
	if err != nil {
		return nil, err
	}

	return s.mapper.ToDetailsDTO(createdTheatre), nil
}

// GetTheatreByID retrieves a theatre by ID
func (s *theatreService) GetTheatreByID(id uuid.UUID) (*dto.TheatreDetails, error) {
	theatre, err := s.theatreRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.ErrorTheatreNotFound)
		}
		return nil, err
	}

	return s.mapper.ToDetailsDTO(theatre), nil
}

// GetAllTheatres retrieves all theatres with pagination
func (s *theatreService) GetAllTheatres(limit, offset int) ([]*dto.TheatreSummary, error) {
	// Apply default and max limits
	if limit <= 0 || limit > constants.MaxLimit {
		limit = constants.DefaultLimit
	}
	if offset < 0 {
		offset = constants.DefaultOffset
	}

	theatres, err := s.theatreRepo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(theatres), nil
}

// UpdateTheatre updates an existing theatre
func (s *theatreService) UpdateTheatre(id uuid.UUID, theatreDTO *dto.TheatreBase) (*dto.TheatreDetails, error) {
	// Validate input
	if err := s.validator.Struct(theatreDTO); err != nil {
		return nil, errors.New(constants.ErrorValidationFailed + ": " + err.Error())
	}

	// Get existing theatre
	theatre, err := s.theatreRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.ErrorTheatreNotFound)
		}
		return nil, err
	}

	// Validate foreign key relationships
	if err := s.validateRelationships(theatreDTO.LocationID, theatreDTO.TheatreTypeID); err != nil {
		return nil, err
	}

	// Update model with new data
	s.mapper.UpdateModel(theatre, theatreDTO)

	// Save to database
	if err := s.theatreRepo.Update(theatre); err != nil {
		return nil, err
	}

	// Get updated theatre with relationships
	updatedTheatre, err := s.theatreRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return s.mapper.ToDetailsDTO(updatedTheatre), nil
}

// DeleteTheatre soft deletes a theatre
func (s *theatreService) DeleteTheatre(id uuid.UUID) error {
	// Check if theatre exists
	_, err := s.theatreRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constants.ErrorTheatreNotFound)
		}
		return err
	}

	return s.theatreRepo.Delete(id)
}

// GetTheatresByLocationID retrieves theatres by location ID
func (s *theatreService) GetTheatresByLocationID(locationID uuid.UUID) ([]*dto.TheatreSummary, error) {
	theatres, err := s.theatreRepo.GetByLocationID(locationID)
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(theatres), nil
}

// GetTheatresByTheatreTypeID retrieves theatres by theatre type ID
func (s *theatreService) GetTheatresByTheatreTypeID(theatreTypeID uuid.UUID) ([]*dto.TheatreSummary, error) {
	theatres, err := s.theatreRepo.GetByTheatreTypeID(theatreTypeID)
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(theatres), nil
}

// GetFeaturedTheatres retrieves all featured theatres
func (s *theatreService) GetFeaturedTheatres() ([]*dto.TheatreSummary, error) {
	theatres, err := s.theatreRepo.GetFeaturedTheatres()
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(theatres), nil
}

// GetActiveTheatres retrieves all active theatres
func (s *theatreService) GetActiveTheatres() ([]*dto.TheatreSummary, error) {
	theatres, err := s.theatreRepo.GetActiveTheatres()
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(theatres), nil
}

// SearchTheatres searches theatres by query string
func (s *theatreService) SearchTheatres(query string) ([]*dto.TheatreSummary, error) {
	if query == "" {
		return []*dto.TheatreSummary{}, nil
	}

	theatres, err := s.theatreRepo.Search(query)
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(theatres), nil
}

// GetNearbyTheatres finds theatres within a radius of given coordinates
func (s *theatreService) GetNearbyTheatres(latitude, longitude, radius float64) ([]*dto.TheatreSummary, error) {
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

	theatres, err := s.theatreRepo.GetNearbyTheatres(latitude, longitude, radius)
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(theatres), nil
}

// validateRelationships validates that location and theatre type exist
func (s *theatreService) validateRelationships(locationID, theatreTypeID uuid.UUID) error {
	// Validate location exists
	_, err := s.locationRepo.GetByID(locationID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constants.ErrorLocationNotFound)
		}
		return err
	}

	// Validate theatre type exists
	_, err = s.theatreTypeRepo.GetByID(theatreTypeID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constants.ErrorTheatreTypeNotFound)
		}
		return err
	}

	return nil
}
