package mappers

import (
	"theatre-management-system/src/dto"
	"theatre-management-system/src/models"
)

// ShowMapper handles mapping between Show models and DTOs
type ShowMapper struct{}

// NewShowMapper creates a new ShowMapper
func NewShowMapper() *ShowMapper {
	return &ShowMapper{}
}

// ToModel converts ShowBase DTO to Show model
func (m *ShowMapper) ToModel(showDTO *dto.ShowBase) *models.Show {
	show := &models.Show{
		Title:       showDTO.Title,
		Description: showDTO.Description,
		Director:    showDTO.Director,
		Cast:        showDTO.Cast,
		StartDate:   showDTO.StartDate,
		EndDate:     showDTO.EndDate,
		ImageURL:    showDTO.ImageURL,
		TrailerURL:  showDTO.TrailerURL,
		TheatreID:   showDTO.TheatreID,
		ShowTypeID:  showDTO.ShowTypeID,
		IsFeatured:  false, // default value
		IsActive:    true,  // default value
	}

	if showDTO.Duration != nil {
		show.Duration = *showDTO.Duration
	}

	if showDTO.Price != nil {
		show.Price = *showDTO.Price
	}

	if showDTO.IsFeatured != nil {
		show.IsFeatured = *showDTO.IsFeatured
	}

	if showDTO.IsActive != nil {
		show.IsActive = *showDTO.IsActive
	}

	return show
}

// ToDetailsDTO converts Show model to ShowDetails DTO
func (m *ShowMapper) ToDetailsDTO(show *models.Show) *dto.ShowDetails {
	showDTO := &dto.ShowDetails{
		ID:          show.ID,
		Title:       show.Title,
		Description: show.Description,
		Director:    show.Director,
		Cast:        show.Cast,
		Duration:    show.Duration,
		StartDate:   show.StartDate,
		EndDate:     show.EndDate,
		Price:       show.Price,
		ImageURL:    show.ImageURL,
		TrailerURL:  show.TrailerURL,
		IsFeatured:  show.IsFeatured,
		IsActive:    show.IsActive,
		CreatedAt:   show.CreatedAt,
		UpdatedAt:   show.UpdatedAt,
		TheatreID:   show.TheatreID,
		ShowTypeID:  show.ShowTypeID,
	}

	// Map theatre if loaded
	if show.Theatre.ID != (show.Theatre.ID) {
		theatreMapper := NewTheatreMapper()
		showDTO.Theatre = *theatreMapper.ToSummaryDTO(&show.Theatre)
	}

	// Map show type if loaded
	if show.ShowType.ID != (show.ShowType.ID) {
		showTypeMapper := NewShowTypeMapper()
		showDTO.ShowType = *showTypeMapper.ToSummaryDTO(&show.ShowType)
	}

	return showDTO
}

// ToSummaryDTO converts Show model to ShowSummary DTO
func (m *ShowMapper) ToSummaryDTO(show *models.Show) *dto.ShowSummary {
	showDTO := &dto.ShowSummary{
		ID:          show.ID,
		Title:       show.Title,
		Description: show.Description,
		Director:    show.Director,
		Duration:    show.Duration,
		StartDate:   show.StartDate,
		EndDate:     show.EndDate,
		Price:       show.Price,
		ImageURL:    show.ImageURL,
		IsFeatured:  show.IsFeatured,
		IsActive:    show.IsActive,
		TheatreID:   show.TheatreID,
		ShowTypeID:  show.ShowTypeID,
	}

	// Map theatre if loaded
	if show.Theatre.ID != (show.Theatre.ID) {
		theatreMapper := NewTheatreMapper()
		showDTO.Theatre = *theatreMapper.ToSummaryDTO(&show.Theatre)
	}

	// Map show type if loaded
	if show.ShowType.ID != (show.ShowType.ID) {
		showTypeMapper := NewShowTypeMapper()
		showDTO.ShowType = *showTypeMapper.ToSummaryDTO(&show.ShowType)
	}

	return showDTO
}

// ToSummaryDTOs converts slice of Show models to slice of ShowSummary DTOs
func (m *ShowMapper) ToSummaryDTOs(shows []*models.Show) []*dto.ShowSummary {
	dtos := make([]*dto.ShowSummary, len(shows))
	for i, show := range shows {
		dtos[i] = m.ToSummaryDTO(show)
	}
	return dtos
}

// UpdateModel updates Show model with ShowBase DTO data
func (m *ShowMapper) UpdateModel(show *models.Show, showDTO *dto.ShowBase) {
	show.Title = showDTO.Title
	show.Description = showDTO.Description
	show.Director = showDTO.Director
	show.Cast = showDTO.Cast
	show.StartDate = showDTO.StartDate
	show.EndDate = showDTO.EndDate
	show.ImageURL = showDTO.ImageURL
	show.TrailerURL = showDTO.TrailerURL
	show.TheatreID = showDTO.TheatreID
	show.ShowTypeID = showDTO.ShowTypeID

	if showDTO.Duration != nil {
		show.Duration = *showDTO.Duration
	}

	if showDTO.Price != nil {
		show.Price = *showDTO.Price
	}

	if showDTO.IsFeatured != nil {
		show.IsFeatured = *showDTO.IsFeatured
	}

	if showDTO.IsActive != nil {
		show.IsActive = *showDTO.IsActive
	}
}
