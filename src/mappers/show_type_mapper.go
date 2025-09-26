package mappers

import (
	"theatre-management-system/src/dto"
	"theatre-management-system/src/models"
)

// ShowTypeMapper handles mapping between ShowType models and DTOs
type ShowTypeMapper struct{}

// NewShowTypeMapper creates a new ShowTypeMapper
func NewShowTypeMapper() *ShowTypeMapper {
	return &ShowTypeMapper{}
}

// ToModel converts ShowTypeBase DTO to ShowType model
func (m *ShowTypeMapper) ToModel(showTypeDTO *dto.ShowTypeBase) *models.ShowType {
	showType := &models.ShowType{
		Name:        showTypeDTO.Name,
		Description: showTypeDTO.Description,
		IsActive:    true, // default value
	}

	if showTypeDTO.IsActive != nil {
		showType.IsActive = *showTypeDTO.IsActive
	}

	return showType
}

// ToDetailsDTO converts ShowType model to ShowTypeDetails DTO
func (m *ShowTypeMapper) ToDetailsDTO(showType *models.ShowType) *dto.ShowTypeDetails {
	showTypeDTO := &dto.ShowTypeDetails{
		ID:          showType.ID,
		Name:        showType.Name,
		Description: showType.Description,
		IsActive:    showType.IsActive,
		CreatedAt:   showType.CreatedAt,
		UpdatedAt:   showType.UpdatedAt,
	}

	// Map shows if loaded
	if len(showType.Shows) > 0 {
		showMapper := NewShowMapper()
		showTypeDTO.Shows = make([]dto.ShowSummary, len(showType.Shows))
		for i, show := range showType.Shows {
			showTypeDTO.Shows[i] = *showMapper.ToSummaryDTO(&show)
		}
	}

	return showTypeDTO
}

// ToSummaryDTO converts ShowType model to ShowTypeSummary DTO
func (m *ShowTypeMapper) ToSummaryDTO(showType *models.ShowType) *dto.ShowTypeSummary {
	return &dto.ShowTypeSummary{
		ID:          showType.ID,
		Name:        showType.Name,
		Description: showType.Description,
		IsActive:    showType.IsActive,
	}
}

// ToSummaryDTOs converts slice of ShowType models to slice of ShowTypeSummary DTOs
func (m *ShowTypeMapper) ToSummaryDTOs(showTypes []*models.ShowType) []*dto.ShowTypeSummary {
	dtos := make([]*dto.ShowTypeSummary, len(showTypes))
	for i, showType := range showTypes {
		dtos[i] = m.ToSummaryDTO(showType)
	}
	return dtos
}

// UpdateModel updates ShowType model with ShowTypeBase DTO data
func (m *ShowTypeMapper) UpdateModel(showType *models.ShowType, showTypeDTO *dto.ShowTypeBase) {
	showType.Name = showTypeDTO.Name
	showType.Description = showTypeDTO.Description

	if showTypeDTO.IsActive != nil {
		showType.IsActive = *showTypeDTO.IsActive
	}
}
