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

// showTypeService implements the ShowTypeService interface
type showTypeService struct {
	showTypeRepo interfaces.ShowTypeRepository
	mapper       *mappers.ShowTypeMapper
	validator    *validator.Validate
}

// NewShowTypeService creates a new show type service
func NewShowTypeService(showTypeRepo interfaces.ShowTypeRepository) interfaces.ShowTypeService {
	return &showTypeService{
		showTypeRepo: showTypeRepo,
		mapper:       mappers.NewShowTypeMapper(),
		validator:    validator.New(),
	}
}

// CreateShowType creates a new show type
func (s *showTypeService) CreateShowType(showTypeDTO *dto.ShowTypeBase) (*dto.ShowTypeDetails, error) {
	// Validate input
	if err := s.validator.Struct(showTypeDTO); err != nil {
		return nil, errors.New(constants.ErrorValidationFailed + ": " + err.Error())
	}

	// Check if show type with same name already exists
	existing, err := s.showTypeRepo.GetByName(showTypeDTO.Name)
	if err == nil && existing != nil {
		return nil, errors.New(constants.ErrorDuplicateEntry + ": show type name already exists")
	}

	// Convert DTO to model
	showType := s.mapper.ToModel(showTypeDTO)

	// Create in database
	if err := s.showTypeRepo.Create(showType); err != nil {
		return nil, err
	}

	// Return created show type
	return s.mapper.ToDetailsDTO(showType), nil
}

// GetShowTypeByID retrieves a show type by ID
func (s *showTypeService) GetShowTypeByID(id uuid.UUID) (*dto.ShowTypeDetails, error) {
	showType, err := s.showTypeRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.ErrorShowTypeNotFound)
		}
		return nil, err
	}

	return s.mapper.ToDetailsDTO(showType), nil
}

// GetAllShowTypes retrieves all show types with pagination
func (s *showTypeService) GetAllShowTypes(limit, offset int) ([]*dto.ShowTypeSummary, error) {
	// Apply default and max limits
	if limit <= 0 || limit > constants.MaxLimit {
		limit = constants.DefaultLimit
	}
	if offset < 0 {
		offset = constants.DefaultOffset
	}

	showTypes, err := s.showTypeRepo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(showTypes), nil
}

// UpdateShowType updates an existing show type
func (s *showTypeService) UpdateShowType(id uuid.UUID, showTypeDTO *dto.ShowTypeBase) (*dto.ShowTypeDetails, error) {
	// Validate input
	if err := s.validator.Struct(showTypeDTO); err != nil {
		return nil, errors.New(constants.ErrorValidationFailed + ": " + err.Error())
	}

	// Get existing show type
	showType, err := s.showTypeRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.ErrorShowTypeNotFound)
		}
		return nil, err
	}

	// Check if name conflicts with another show type
	if showType.Name != showTypeDTO.Name {
		existing, err := s.showTypeRepo.GetByName(showTypeDTO.Name)
		if err == nil && existing != nil && existing.ID != id {
			return nil, errors.New(constants.ErrorDuplicateEntry + ": show type name already exists")
		}
	}

	// Update model with new data
	s.mapper.UpdateModel(showType, showTypeDTO)

	// Save to database
	if err := s.showTypeRepo.Update(showType); err != nil {
		return nil, err
	}

	return s.mapper.ToDetailsDTO(showType), nil
}

// DeleteShowType soft deletes a show type
func (s *showTypeService) DeleteShowType(id uuid.UUID) error {
	// Check if show type exists
	_, err := s.showTypeRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constants.ErrorShowTypeNotFound)
		}
		return err
	}

	return s.showTypeRepo.Delete(id)
}

// GetShowTypeByName retrieves a show type by name
func (s *showTypeService) GetShowTypeByName(name string) (*dto.ShowTypeDetails, error) {
	if name == "" {
		return nil, errors.New(constants.ErrorInvalidInput + ": name cannot be empty")
	}

	showType, err := s.showTypeRepo.GetByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.ErrorShowTypeNotFound)
		}
		return nil, err
	}

	return s.mapper.ToDetailsDTO(showType), nil
}

// GetActiveShowTypes retrieves all active show types
func (s *showTypeService) GetActiveShowTypes() ([]*dto.ShowTypeSummary, error) {
	showTypes, err := s.showTypeRepo.GetActiveTypes()
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(showTypes), nil
}
