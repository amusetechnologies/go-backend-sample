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

// theatreTypeService implements the TheatreTypeService interface
type theatreTypeService struct {
	theatreTypeRepo interfaces.TheatreTypeRepository
	mapper          *mappers.TheatreTypeMapper
	validator       *validator.Validate
}

// NewTheatreTypeService creates a new theatre type service
func NewTheatreTypeService(theatreTypeRepo interfaces.TheatreTypeRepository) interfaces.TheatreTypeService {
	return &theatreTypeService{
		theatreTypeRepo: theatreTypeRepo,
		mapper:          mappers.NewTheatreTypeMapper(),
		validator:       validator.New(),
	}
}

// CreateTheatreType creates a new theatre type
func (s *theatreTypeService) CreateTheatreType(theatreTypeDTO *dto.TheatreTypeBase) (*dto.TheatreTypeDetails, error) {
	// Validate input
	if err := s.validator.Struct(theatreTypeDTO); err != nil {
		return nil, errors.New(constants.ErrorValidationFailed + ": " + err.Error())
	}

	// Check if theatre type with same name already exists
	existing, err := s.theatreTypeRepo.GetByName(theatreTypeDTO.Name)
	if err == nil && existing != nil {
		return nil, errors.New(constants.ErrorDuplicateEntry + ": theatre type name already exists")
	}

	// Convert DTO to model
	theatreType := s.mapper.ToModel(theatreTypeDTO)

	// Create in database
	if err := s.theatreTypeRepo.Create(theatreType); err != nil {
		return nil, err
	}

	// Return created theatre type
	return s.mapper.ToDetailsDTO(theatreType), nil
}

// GetTheatreTypeByID retrieves a theatre type by ID
func (s *theatreTypeService) GetTheatreTypeByID(id uuid.UUID) (*dto.TheatreTypeDetails, error) {
	theatreType, err := s.theatreTypeRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.ErrorTheatreTypeNotFound)
		}
		return nil, err
	}

	return s.mapper.ToDetailsDTO(theatreType), nil
}

// GetAllTheatreTypes retrieves all theatre types with pagination
func (s *theatreTypeService) GetAllTheatreTypes(limit, offset int) ([]*dto.TheatreTypeSummary, error) {
	// Apply default and max limits
	if limit <= 0 || limit > constants.MaxLimit {
		limit = constants.DefaultLimit
	}
	if offset < 0 {
		offset = constants.DefaultOffset
	}

	theatreTypes, err := s.theatreTypeRepo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(theatreTypes), nil
}

// UpdateTheatreType updates an existing theatre type
func (s *theatreTypeService) UpdateTheatreType(id uuid.UUID, theatreTypeDTO *dto.TheatreTypeBase) (*dto.TheatreTypeDetails, error) {
	// Validate input
	if err := s.validator.Struct(theatreTypeDTO); err != nil {
		return nil, errors.New(constants.ErrorValidationFailed + ": " + err.Error())
	}

	// Get existing theatre type
	theatreType, err := s.theatreTypeRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.ErrorTheatreTypeNotFound)
		}
		return nil, err
	}

	// Check if name conflicts with another theatre type
	if theatreType.Name != theatreTypeDTO.Name {
		existing, err := s.theatreTypeRepo.GetByName(theatreTypeDTO.Name)
		if err == nil && existing != nil && existing.ID != id {
			return nil, errors.New(constants.ErrorDuplicateEntry + ": theatre type name already exists")
		}
	}

	// Update model with new data
	s.mapper.UpdateModel(theatreType, theatreTypeDTO)

	// Save to database
	if err := s.theatreTypeRepo.Update(theatreType); err != nil {
		return nil, err
	}

	return s.mapper.ToDetailsDTO(theatreType), nil
}

// DeleteTheatreType soft deletes a theatre type
func (s *theatreTypeService) DeleteTheatreType(id uuid.UUID) error {
	// Check if theatre type exists
	_, err := s.theatreTypeRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constants.ErrorTheatreTypeNotFound)
		}
		return err
	}

	return s.theatreTypeRepo.Delete(id)
}

// GetTheatreTypeByName retrieves a theatre type by name
func (s *theatreTypeService) GetTheatreTypeByName(name string) (*dto.TheatreTypeDetails, error) {
	if name == "" {
		return nil, errors.New(constants.ErrorInvalidInput + ": name cannot be empty")
	}

	theatreType, err := s.theatreTypeRepo.GetByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.ErrorTheatreTypeNotFound)
		}
		return nil, err
	}

	return s.mapper.ToDetailsDTO(theatreType), nil
}

// GetActiveTheatreTypes retrieves all active theatre types
func (s *theatreTypeService) GetActiveTheatreTypes() ([]*dto.TheatreTypeSummary, error) {
	theatreTypes, err := s.theatreTypeRepo.GetActiveTypes()
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(theatreTypes), nil
}
