package mappers

import (
	"theatre-management-system/src/dto"
	"theatre-management-system/src/models"
)

// TheatreMapper handles mapping between Theatre models and DTOs
type TheatreMapper struct{}

// NewTheatreMapper creates a new TheatreMapper
func NewTheatreMapper() *TheatreMapper {
	return &TheatreMapper{}
}

// ToModel converts TheatreBase DTO to Theatre model
func (m *TheatreMapper) ToModel(theatreDTO *dto.TheatreBase) *models.Theatre {
	theatre := &models.Theatre{
		Name:          theatreDTO.Name,
		Description:   theatreDTO.Description,
		Address:       theatreDTO.Address,
		Phone:         theatreDTO.Phone,
		Email:         theatreDTO.Email,
		Website:       theatreDTO.Website,
		ImageURL:      theatreDTO.ImageURL,
		LocationID:    theatreDTO.LocationID,
		TheatreTypeID: theatreDTO.TheatreTypeID,
		IsFeatured:    false, // default value
		IsActive:      true,  // default value
	}

	if theatreDTO.Capacity != nil {
		theatre.Capacity = *theatreDTO.Capacity
	}

	if theatreDTO.IsFeatured != nil {
		theatre.IsFeatured = *theatreDTO.IsFeatured
	}

	if theatreDTO.IsActive != nil {
		theatre.IsActive = *theatreDTO.IsActive
	}

	return theatre
}

// ToDetailsDTO converts Theatre model to TheatreDetails DTO
func (m *TheatreMapper) ToDetailsDTO(theatre *models.Theatre) *dto.TheatreDetails {
	theatreDTO := &dto.TheatreDetails{
		ID:            theatre.ID,
		Name:          theatre.Name,
		Description:   theatre.Description,
		Capacity:      theatre.Capacity,
		Address:       theatre.Address,
		Phone:         theatre.Phone,
		Email:         theatre.Email,
		Website:       theatre.Website,
		ImageURL:      theatre.ImageURL,
		IsFeatured:    theatre.IsFeatured,
		IsActive:      theatre.IsActive,
		CreatedAt:     theatre.CreatedAt,
		UpdatedAt:     theatre.UpdatedAt,
		LocationID:    theatre.LocationID,
		TheatreTypeID: theatre.TheatreTypeID,
	}

	// Map location if loaded
	if theatre.Location.ID != (theatre.Location.ID) {
		locationMapper := NewLocationMapper()
		theatreDTO.Location = *locationMapper.ToSummaryDTO(&theatre.Location)
	}

	// Map theatre type if loaded
	if theatre.TheatreType.ID != (theatre.TheatreType.ID) {
		theatreTypeMapper := NewTheatreTypeMapper()
		theatreDTO.TheatreType = *theatreTypeMapper.ToSummaryDTO(&theatre.TheatreType)
	}

	// Map shows if loaded
	if len(theatre.Shows) > 0 {
		showMapper := NewShowMapper()
		theatreDTO.Shows = make([]dto.ShowSummary, len(theatre.Shows))
		for i, show := range theatre.Shows {
			theatreDTO.Shows[i] = *showMapper.ToSummaryDTO(&show)
		}
	}

	return theatreDTO
}

// ToSummaryDTO converts Theatre model to TheatreSummary DTO
func (m *TheatreMapper) ToSummaryDTO(theatre *models.Theatre) *dto.TheatreSummary {
	theatreDTO := &dto.TheatreSummary{
		ID:            theatre.ID,
		Name:          theatre.Name,
		Description:   theatre.Description,
		Capacity:      theatre.Capacity,
		ImageURL:      theatre.ImageURL,
		IsFeatured:    theatre.IsFeatured,
		IsActive:      theatre.IsActive,
		LocationID:    theatre.LocationID,
		TheatreTypeID: theatre.TheatreTypeID,
	}

	// Map location if loaded
	if theatre.Location.ID != (theatre.Location.ID) {
		locationMapper := NewLocationMapper()
		theatreDTO.Location = *locationMapper.ToSummaryDTO(&theatre.Location)
	}

	// Map theatre type if loaded
	if theatre.TheatreType.ID != (theatre.TheatreType.ID) {
		theatreTypeMapper := NewTheatreTypeMapper()
		theatreDTO.TheatreType = *theatreTypeMapper.ToSummaryDTO(&theatre.TheatreType)
	}

	return theatreDTO
}

// ToSummaryDTOs converts slice of Theatre models to slice of TheatreSummary DTOs
func (m *TheatreMapper) ToSummaryDTOs(theatres []*models.Theatre) []*dto.TheatreSummary {
	dtos := make([]*dto.TheatreSummary, len(theatres))
	for i, theatre := range theatres {
		dtos[i] = m.ToSummaryDTO(theatre)
	}
	return dtos
}

// UpdateModel updates Theatre model with TheatreBase DTO data
func (m *TheatreMapper) UpdateModel(theatre *models.Theatre, theatreDTO *dto.TheatreBase) {
	theatre.Name = theatreDTO.Name
	theatre.Description = theatreDTO.Description
	theatre.Address = theatreDTO.Address
	theatre.Phone = theatreDTO.Phone
	theatre.Email = theatreDTO.Email
	theatre.Website = theatreDTO.Website
	theatre.ImageURL = theatreDTO.ImageURL
	theatre.LocationID = theatreDTO.LocationID
	theatre.TheatreTypeID = theatreDTO.TheatreTypeID

	if theatreDTO.Capacity != nil {
		theatre.Capacity = *theatreDTO.Capacity
	}

	if theatreDTO.IsFeatured != nil {
		theatre.IsFeatured = *theatreDTO.IsFeatured
	}

	if theatreDTO.IsActive != nil {
		theatre.IsActive = *theatreDTO.IsActive
	}
}
