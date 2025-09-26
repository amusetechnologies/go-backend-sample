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

// TheatreController handles HTTP requests for theatres
type TheatreController struct {
	theatreService interfaces.TheatreService
}

// NewTheatreController creates a new theatre controller
func NewTheatreController(theatreService interfaces.TheatreService) *TheatreController {
	return &TheatreController{
		theatreService: theatreService,
	}
}

// CreateTheatre handles POST /theatres
func (ctrl *TheatreController) CreateTheatre(c *gin.Context) {
	var theatreDTO dto.TheatreBase

	if err := c.ShouldBindJSON(&theatreDTO); err != nil {
		BadRequestResponse(c, constants.ErrorInvalidInput, err)
		return
	}

	theatre, err := ctrl.theatreService.CreateTheatre(&theatreDTO)
	if err != nil {
		if err.Error() == constants.ErrorValidationFailed {
			ValidationErrorResponse(c, err)
			return
		}
		if err.Error() == constants.ErrorLocationNotFound || err.Error() == constants.ErrorTheatreTypeNotFound {
			BadRequestResponse(c, err.Error(), err)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusCreated, constants.MessageTheatreCreated, theatre)
}

// GetTheatreByID handles GET /theatres/:id
func (ctrl *TheatreController) GetTheatreByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		BadRequestResponse(c, constants.ErrorInvalidUUID, err)
		return
	}

	theatre, err := ctrl.theatreService.GetTheatreByID(id)
	if err != nil {
		if err.Error() == constants.ErrorTheatreNotFound {
			NotFoundResponse(c, constants.ErrorTheatreNotFound)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, theatre)
}

// GetAllTheatres handles GET /theatres
func (ctrl *TheatreController) GetAllTheatres(c *gin.Context) {
	params := GetPaginationParams(c)

	theatres, err := ctrl.theatreService.GetAllTheatres(params.Limit, params.Offset)
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, theatres)
}

// UpdateTheatre handles PATCH /theatres/:id
func (ctrl *TheatreController) UpdateTheatre(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		BadRequestResponse(c, constants.ErrorInvalidUUID, err)
		return
	}

	var theatreDTO dto.TheatreBase
	if err := c.ShouldBindJSON(&theatreDTO); err != nil {
		BadRequestResponse(c, constants.ErrorInvalidInput, err)
		return
	}

	theatre, err := ctrl.theatreService.UpdateTheatre(id, &theatreDTO)
	if err != nil {
		if err.Error() == constants.ErrorTheatreNotFound {
			NotFoundResponse(c, constants.ErrorTheatreNotFound)
			return
		}
		if err.Error() == constants.ErrorValidationFailed {
			ValidationErrorResponse(c, err)
			return
		}
		if err.Error() == constants.ErrorLocationNotFound || err.Error() == constants.ErrorTheatreTypeNotFound {
			BadRequestResponse(c, err.Error(), err)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.MessageTheatreUpdated, theatre)
}

// DeleteTheatre handles DELETE /theatres/:id
func (ctrl *TheatreController) DeleteTheatre(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		BadRequestResponse(c, constants.ErrorInvalidUUID, err)
		return
	}

	err = ctrl.theatreService.DeleteTheatre(id)
	if err != nil {
		if err.Error() == constants.ErrorTheatreNotFound {
			NotFoundResponse(c, constants.ErrorTheatreNotFound)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.MessageTheatreDeleted, nil)
}

// GetTheatresByLocationID handles GET /theatres/location/:locationId
func (ctrl *TheatreController) GetTheatresByLocationID(c *gin.Context) {
	locationIDParam := c.Param("locationId")
	locationID, err := uuid.Parse(locationIDParam)
	if err != nil {
		BadRequestResponse(c, constants.ErrorInvalidUUID, err)
		return
	}

	theatres, err := ctrl.theatreService.GetTheatresByLocationID(locationID)
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, theatres)
}

// GetTheatresByTheatreTypeID handles GET /theatres/type/:typeId
func (ctrl *TheatreController) GetTheatresByTheatreTypeID(c *gin.Context) {
	typeIDParam := c.Param("typeId")
	typeID, err := uuid.Parse(typeIDParam)
	if err != nil {
		BadRequestResponse(c, constants.ErrorInvalidUUID, err)
		return
	}

	theatres, err := ctrl.theatreService.GetTheatresByTheatreTypeID(typeID)
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, theatres)
}

// GetFeaturedTheatres handles GET /theatres/featured
func (ctrl *TheatreController) GetFeaturedTheatres(c *gin.Context) {
	theatres, err := ctrl.theatreService.GetFeaturedTheatres()
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, theatres)
}

// GetActiveTheatres handles GET /theatres/active
func (ctrl *TheatreController) GetActiveTheatres(c *gin.Context) {
	theatres, err := ctrl.theatreService.GetActiveTheatres()
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, theatres)
}

// SearchTheatres handles GET /theatres/search
func (ctrl *TheatreController) SearchTheatres(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		BadRequestResponse(c, "search query is required", nil)
		return
	}

	theatres, err := ctrl.theatreService.SearchTheatres(query)
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, theatres)
}

// GetNearbyTheatres handles GET /theatres/nearby
func (ctrl *TheatreController) GetNearbyTheatres(c *gin.Context) {
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

	theatres, err := ctrl.theatreService.GetNearbyTheatres(latitude, longitude, radius)
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, theatres)
}
