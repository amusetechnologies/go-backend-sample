package controllers

import (
	"net/http"
	"strconv"
	"theatre-management-system/src/constants"
	"theatre-management-system/src/dto"
	"theatre-management-system/src/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// LocationController handles HTTP requests for locations
type LocationController struct {
	locationService interfaces.LocationService
}

// NewLocationController creates a new location controller
func NewLocationController(locationService interfaces.LocationService) *LocationController {
	return &LocationController{
		locationService: locationService,
	}
}

// CreateLocation handles POST /locations
func (ctrl *LocationController) CreateLocation(c *gin.Context) {
	var locationDTO dto.LocationBase

	if err := c.ShouldBindJSON(&locationDTO); err != nil {
		BadRequestResponse(c, constants.ErrorInvalidInput, err)
		return
	}

	location, err := ctrl.locationService.CreateLocation(&locationDTO)
	if err != nil {
		if err.Error() == constants.ErrorValidationFailed {
			ValidationErrorResponse(c, err)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusCreated, constants.MessageLocationCreated, location)
}

// GetLocationByID handles GET /locations/:id
func (ctrl *LocationController) GetLocationByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		BadRequestResponse(c, constants.ErrorInvalidUUID, err)
		return
	}

	location, err := ctrl.locationService.GetLocationByID(id)
	if err != nil {
		if err.Error() == constants.ErrorLocationNotFound {
			NotFoundResponse(c, constants.ErrorLocationNotFound)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, location)
}

// GetAllLocations handles GET /locations
func (ctrl *LocationController) GetAllLocations(c *gin.Context) {
	params := GetPaginationParams(c)

	locations, err := ctrl.locationService.GetAllLocations(params.Limit, params.Offset)
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, locations)
}

// UpdateLocation handles PATCH /locations/:id
func (ctrl *LocationController) UpdateLocation(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		BadRequestResponse(c, constants.ErrorInvalidUUID, err)
		return
	}

	var locationDTO dto.LocationBase
	if err := c.ShouldBindJSON(&locationDTO); err != nil {
		BadRequestResponse(c, constants.ErrorInvalidInput, err)
		return
	}

	location, err := ctrl.locationService.UpdateLocation(id, &locationDTO)
	if err != nil {
		if err.Error() == constants.ErrorLocationNotFound {
			NotFoundResponse(c, constants.ErrorLocationNotFound)
			return
		}
		if err.Error() == constants.ErrorValidationFailed {
			ValidationErrorResponse(c, err)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.MessageLocationUpdated, location)
}

// DeleteLocation handles DELETE /locations/:id
func (ctrl *LocationController) DeleteLocation(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		BadRequestResponse(c, constants.ErrorInvalidUUID, err)
		return
	}

	err = ctrl.locationService.DeleteLocation(id)
	if err != nil {
		if err.Error() == constants.ErrorLocationNotFound {
			NotFoundResponse(c, constants.ErrorLocationNotFound)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.MessageLocationDeleted, nil)
}

// GetLocationsByCoordinates handles GET /locations/nearby
func (ctrl *LocationController) GetLocationsByCoordinates(c *gin.Context) {
	latStr := c.Query("latitude")
	lonStr := c.Query("longitude")
	radiusStr := c.Query("radius")

	if latStr == "" || lonStr == "" {
		BadRequestResponse(c, "latitude and longitude are required", nil)
		return
	}

	latitude, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		BadRequestResponse(c, "invalid latitude", err)
		return
	}

	longitude, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		BadRequestResponse(c, "invalid longitude", err)
		return
	}

	radius := constants.DefaultRadius
	if radiusStr != "" {
		radius, err = strconv.ParseFloat(radiusStr, 64)
		if err != nil {
			BadRequestResponse(c, "invalid radius", err)
			return
		}
	}

	locations, err := ctrl.locationService.GetLocationsByCoordinates(latitude, longitude, radius)
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, locations)
}

// GetActiveLocations handles GET /locations/active
func (ctrl *LocationController) GetActiveLocations(c *gin.Context) {
	locations, err := ctrl.locationService.GetActiveLocations()
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, locations)
}

// SearchLocations handles GET /locations/search
func (ctrl *LocationController) SearchLocations(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		BadRequestResponse(c, "search query is required", nil)
		return
	}

	locations, err := ctrl.locationService.SearchLocations(query)
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, locations)
}
