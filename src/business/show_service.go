package business

import (
	"errors"
	"theatre-management-system/src/constants"
	"theatre-management-system/src/dto"
	"theatre-management-system/src/interfaces"
	"theatre-management-system/src/mappers"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// showService implements the ShowService interface
type showService struct {
	showRepo     interfaces.ShowRepository
	theatreRepo  interfaces.TheatreRepository
	showTypeRepo interfaces.ShowTypeRepository
	mapper       *mappers.ShowMapper
	validator    *validator.Validate
}

// NewShowService creates a new show service
func NewShowService(
	showRepo interfaces.ShowRepository,
	theatreRepo interfaces.TheatreRepository,
	showTypeRepo interfaces.ShowTypeRepository,
) interfaces.ShowService {
	return &showService{
		showRepo:     showRepo,
		theatreRepo:  theatreRepo,
		showTypeRepo: showTypeRepo,
		mapper:       mappers.NewShowMapper(),
		validator:    validator.New(),
	}
}

// CreateShow creates a new show
func (s *showService) CreateShow(showDTO *dto.ShowBase) (*dto.ShowDetails, error) {
	// Validate input
	if err := s.validator.Struct(showDTO); err != nil {
		return nil, errors.New(constants.ErrorValidationFailed + ": " + err.Error())
	}

	// Validate business rules
	if err := s.validateShowDates(showDTO.StartDate, showDTO.EndDate); err != nil {
		return nil, err
	}

	// Validate foreign key relationships
	if err := s.validateRelationships(showDTO.TheatreID, showDTO.ShowTypeID); err != nil {
		return nil, err
	}

	// Convert DTO to model
	show := s.mapper.ToModel(showDTO)

	// Create in database
	if err := s.showRepo.Create(show); err != nil {
		return nil, err
	}

	// Get created show with relationships
	createdShow, err := s.showRepo.GetByID(show.ID)
	if err != nil {
		return nil, err
	}

	return s.mapper.ToDetailsDTO(createdShow), nil
}

// GetShowByID retrieves a show by ID
func (s *showService) GetShowByID(id uuid.UUID) (*dto.ShowDetails, error) {
	show, err := s.showRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.ErrorShowNotFound)
		}
		return nil, err
	}

	return s.mapper.ToDetailsDTO(show), nil
}

// GetAllShows retrieves all shows with pagination
func (s *showService) GetAllShows(limit, offset int) ([]*dto.ShowSummary, error) {
	// Apply default and max limits
	if limit <= 0 || limit > constants.MaxLimit {
		limit = constants.DefaultLimit
	}
	if offset < 0 {
		offset = constants.DefaultOffset
	}

	shows, err := s.showRepo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(shows), nil
}

// UpdateShow updates an existing show
func (s *showService) UpdateShow(id uuid.UUID, showDTO *dto.ShowBase) (*dto.ShowDetails, error) {
	// Validate input
	if err := s.validator.Struct(showDTO); err != nil {
		return nil, errors.New(constants.ErrorValidationFailed + ": " + err.Error())
	}

	// Get existing show
	show, err := s.showRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.ErrorShowNotFound)
		}
		return nil, err
	}

	// Validate business rules
	if err := s.validateShowDates(showDTO.StartDate, showDTO.EndDate); err != nil {
		return nil, err
	}

	// Validate foreign key relationships
	if err := s.validateRelationships(showDTO.TheatreID, showDTO.ShowTypeID); err != nil {
		return nil, err
	}

	// Update model with new data
	s.mapper.UpdateModel(show, showDTO)

	// Save to database
	if err := s.showRepo.Update(show); err != nil {
		return nil, err
	}

	// Get updated show with relationships
	updatedShow, err := s.showRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return s.mapper.ToDetailsDTO(updatedShow), nil
}

// DeleteShow soft deletes a show
func (s *showService) DeleteShow(id uuid.UUID) error {
	// Check if show exists
	_, err := s.showRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constants.ErrorShowNotFound)
		}
		return err
	}

	return s.showRepo.Delete(id)
}

// GetShowsByTheatreID retrieves shows by theatre ID
func (s *showService) GetShowsByTheatreID(theatreID uuid.UUID) ([]*dto.ShowSummary, error) {
	shows, err := s.showRepo.GetByTheatreID(theatreID)
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(shows), nil
}

// GetShowsByShowTypeID retrieves shows by show type ID
func (s *showService) GetShowsByShowTypeID(showTypeID uuid.UUID) ([]*dto.ShowSummary, error) {
	shows, err := s.showRepo.GetByShowTypeID(showTypeID)
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(shows), nil
}

// GetFeaturedShows retrieves all featured shows
func (s *showService) GetFeaturedShows() ([]*dto.ShowSummary, error) {
	shows, err := s.showRepo.GetFeaturedShows()
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(shows), nil
}

// GetActiveShows retrieves all active shows
func (s *showService) GetActiveShows() ([]*dto.ShowSummary, error) {
	shows, err := s.showRepo.GetActiveShows()
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(shows), nil
}

// GetCurrentShows retrieves shows that are currently running
func (s *showService) GetCurrentShows() ([]*dto.ShowSummary, error) {
	shows, err := s.showRepo.GetCurrentShows()
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(shows), nil
}

// GetUpcomingShows retrieves shows that will start in the future
func (s *showService) GetUpcomingShows() ([]*dto.ShowSummary, error) {
	shows, err := s.showRepo.GetUpcomingShows()
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(shows), nil
}

// SearchShows searches shows by query string
func (s *showService) SearchShows(query string) ([]*dto.ShowSummary, error) {
	if query == "" {
		return []*dto.ShowSummary{}, nil
	}

	shows, err := s.showRepo.Search(query)
	if err != nil {
		return nil, err
	}

	return s.mapper.ToSummaryDTOs(shows), nil
}

// validateShowDates validates that start and end dates make sense
func (s *showService) validateShowDates(startDate, endDate *time.Time) error {
	if startDate != nil && endDate != nil {
		if endDate.Before(*startDate) {
			return errors.New("end date cannot be before start date")
		}
	}
	return nil
}

// validateRelationships validates that theatre and show type exist
func (s *showService) validateRelationships(theatreID, showTypeID uuid.UUID) error {
	// Validate theatre exists
	_, err := s.theatreRepo.GetByID(theatreID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constants.ErrorTheatreNotFound)
		}
		return err
	}

	// Validate show type exists
	_, err = s.showTypeRepo.GetByID(showTypeID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constants.ErrorShowTypeNotFound)
		}
		return err
	}

	return nil
}
