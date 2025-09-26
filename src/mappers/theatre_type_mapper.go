package mappers

import (
	"theatre-management-system/src/dto"
	"theatre-management-system/src/models"
)

// TheatreTypeMapper handles mapping between TheatreType models and DTOs
type TheatreTypeMapper struct{}

// NewTheatreTypeMapper creates a new TheatreTypeMapper
func NewTheatreTypeMapper() *TheatreTypeMapper {
	return &TheatreTypeMapper{}
}

// ToModel converts TheatreTypeBase DTO to TheatreType model
func (m *TheatreTypeMapper) ToModel(theatreTypeDTO *dto.TheatreTypeBase) *models.TheatreType {
	theatreType := &models.TheatreType{
		Name:        theatreTypeDTO.Name,
		Description: theatreTypeDTO.Description,
		IsActive:    true, // default value
	}

	if theatreTypeDTO.IsActive != nil {
		theatreType.IsActive = *theatreTypeDTO.IsActive
	}

	return theatreType
}

// ToDetailsDTO converts TheatreType model to TheatreTypeDetails DTO
func (m *TheatreTypeMapper) ToDetailsDTO(theatreType *models.TheatreType) *dto.TheatreTypeDetails {
	theatreTypeDTO := &dto.TheatreTypeDetails{
		ID:          theatreType.ID,
		Name:        theatreType.Name,
		Description: theatreType.Description,
		IsActive:    theatreType.IsActive,
		CreatedAt:   theatreType.CreatedAt,
		UpdatedAt:   theatreType.UpdatedAt,
	}

	// Map theatres if loaded
	if len(theatreType.Theatres) > 0 {
		theatreMapper := NewTheatreMapper()
		theatreTypeDTO.Theatres = make([]dto.TheatreSummary, len(theatreType.Theatres))
		for i, theatre := range theatreType.Theatres {
			theatreTypeDTO.Theatres[i] = *theatreMapper.ToSummaryDTO(&theatre)
		}
	}

	return theatreTypeDTO
}

// ToSummaryDTO converts TheatreType model to TheatreTypeSummary DTO
func (m *TheatreTypeMapper) ToSummaryDTO(theatreType *models.TheatreType) *dto.TheatreTypeSummary {
	return &dto.TheatreTypeSummary{
		ID:          theatreType.ID,
		Name:        theatreType.Name,
		Description: theatreType.Description,
		IsActive:    theatreType.IsActive,
	}
}

// ToSummaryDTOs converts slice of TheatreType models to slice of TheatreTypeSummary DTOs
func (m *TheatreTypeMapper) ToSummaryDTOs(theatreTypes []*models.TheatreType) []*dto.TheatreTypeSummary {
	dtos := make([]*dto.TheatreTypeSummary, len(theatreTypes))
	for i, theatreType := range theatreTypes {
		dtos[i] = m.ToSummaryDTO(theatreType)
	}
	return dtos
}

// UpdateModel updates TheatreType model with TheatreTypeBase DTO data
func (m *TheatreTypeMapper) UpdateModel(theatreType *models.TheatreType, theatreTypeDTO *dto.TheatreTypeBase) {
	theatreType.Name = theatreTypeDTO.Name
	theatreType.Description = theatreTypeDTO.Description

	if theatreTypeDTO.IsActive != nil {
		theatreType.IsActive = *theatreTypeDTO.IsActive
	}
}
