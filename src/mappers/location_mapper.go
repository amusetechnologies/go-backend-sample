package mappers

import (
	"theatre-management-system/src/dto"
	"theatre-management-system/src/models"
)

// LocationMapper handles mapping between Location models and DTOs
type LocationMapper struct{}

// NewLocationMapper creates a new LocationMapper
func NewLocationMapper() *LocationMapper {
	return &LocationMapper{}
}

// ToModel converts LocationBase DTO to Location model
func (m *LocationMapper) ToModel(locationDTO *dto.LocationBase) *models.Location {
	location := &models.Location{
		Name:        locationDTO.Name,
		City:        locationDTO.City,
		State:       locationDTO.State,
		Country:     locationDTO.Country,
		Latitude:    locationDTO.Latitude,
		Longitude:   locationDTO.Longitude,
		PostalCode:  locationDTO.PostalCode,
		Address:     locationDTO.Address,
		Description: locationDTO.Description,
		IsActive:    true, // default value
	}

	if locationDTO.IsActive != nil {
		location.IsActive = *locationDTO.IsActive
	}

	return location
}

// ToDetailsDTO converts Location model to LocationDetails DTO
func (m *LocationMapper) ToDetailsDTO(location *models.Location) *dto.LocationDetails {
	locationDTO := &dto.LocationDetails{
		ID:          location.ID,
		Name:        location.Name,
		City:        location.City,
		State:       location.State,
		Country:     location.Country,
		Latitude:    location.Latitude,
		Longitude:   location.Longitude,
		PostalCode:  location.PostalCode,
		Address:     location.Address,
		Description: location.Description,
		IsActive:    location.IsActive,
		CreatedAt:   location.CreatedAt,
		UpdatedAt:   location.UpdatedAt,
	}

	// Map theatres if loaded
	if len(location.Theatres) > 0 {
		theatreMapper := NewTheatreMapper()
		locationDTO.Theatres = make([]dto.TheatreSummary, len(location.Theatres))
		for i, theatre := range location.Theatres {
			locationDTO.Theatres[i] = *theatreMapper.ToSummaryDTO(&theatre)
		}
	}

	return locationDTO
}

// ToSummaryDTO converts Location model to LocationSummary DTO
func (m *LocationMapper) ToSummaryDTO(location *models.Location) *dto.LocationSummary {
	return &dto.LocationSummary{
		ID:        location.ID,
		Name:      location.Name,
		City:      location.City,
		State:     location.State,
		Country:   location.Country,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		IsActive:  location.IsActive,
	}
}

// ToSummaryDTOs converts slice of Location models to slice of LocationSummary DTOs
func (m *LocationMapper) ToSummaryDTOs(locations []*models.Location) []*dto.LocationSummary {
	dtos := make([]*dto.LocationSummary, len(locations))
	for i, location := range locations {
		dtos[i] = m.ToSummaryDTO(location)
	}
	return dtos
}

// UpdateModel updates Location model with LocationBase DTO data
func (m *LocationMapper) UpdateModel(location *models.Location, locationDTO *dto.LocationBase) {
	location.Name = locationDTO.Name
	location.City = locationDTO.City
	location.State = locationDTO.State
	location.Country = locationDTO.Country
	location.Latitude = locationDTO.Latitude
	location.Longitude = locationDTO.Longitude
	location.PostalCode = locationDTO.PostalCode
	location.Address = locationDTO.Address
	location.Description = locationDTO.Description

	if locationDTO.IsActive != nil {
		location.IsActive = *locationDTO.IsActive
	}
}
