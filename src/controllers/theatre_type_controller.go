package controllers

import (
	"net/http"
	"theatre-management-system/src/constants"
	"theatre-management-system/src/dto"
	"theatre-management-system/src/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TheatreTypeController handles HTTP requests for theatre types
type TheatreTypeController struct {
	theatreTypeService interfaces.TheatreTypeService
}

// NewTheatreTypeController creates a new theatre type controller
func NewTheatreTypeController(theatreTypeService interfaces.TheatreTypeService) *TheatreTypeController {
	return &TheatreTypeController{
		theatreTypeService: theatreTypeService,
	}
}

// CreateTheatreType handles POST /theatre-types
func (ctrl *TheatreTypeController) CreateTheatreType(c *gin.Context) {
	var theatreTypeDTO dto.TheatreTypeBase

	if err := c.ShouldBindJSON(&theatreTypeDTO); err != nil {
		BadRequestResponse(c, constants.ErrorInvalidInput, err)
		return
	}

	theatreType, err := ctrl.theatreTypeService.CreateTheatreType(&theatreTypeDTO)
	if err != nil {
		if err.Error() == constants.ErrorValidationFailed {
			ValidationErrorResponse(c, err)
			return
		}
		if err.Error() == constants.ErrorDuplicateEntry {
			BadRequestResponse(c, constants.ErrorDuplicateEntry, err)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusCreated, constants.MessageTheatreTypeCreated, theatreType)
}

// GetTheatreTypeByID handles GET /theatre-types/:id
func (ctrl *TheatreTypeController) GetTheatreTypeByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		BadRequestResponse(c, constants.ErrorInvalidUUID, err)
		return
	}

	theatreType, err := ctrl.theatreTypeService.GetTheatreTypeByID(id)
	if err != nil {
		if err.Error() == constants.ErrorTheatreTypeNotFound {
			NotFoundResponse(c, constants.ErrorTheatreTypeNotFound)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, theatreType)
}

// GetAllTheatreTypes handles GET /theatre-types
func (ctrl *TheatreTypeController) GetAllTheatreTypes(c *gin.Context) {
	params := GetPaginationParams(c)

	theatreTypes, err := ctrl.theatreTypeService.GetAllTheatreTypes(params.Limit, params.Offset)
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, theatreTypes)
}

// UpdateTheatreType handles PATCH /theatre-types/:id
func (ctrl *TheatreTypeController) UpdateTheatreType(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		BadRequestResponse(c, constants.ErrorInvalidUUID, err)
		return
	}

	var theatreTypeDTO dto.TheatreTypeBase
	if err := c.ShouldBindJSON(&theatreTypeDTO); err != nil {
		BadRequestResponse(c, constants.ErrorInvalidInput, err)
		return
	}

	theatreType, err := ctrl.theatreTypeService.UpdateTheatreType(id, &theatreTypeDTO)
	if err != nil {
		if err.Error() == constants.ErrorTheatreTypeNotFound {
			NotFoundResponse(c, constants.ErrorTheatreTypeNotFound)
			return
		}
		if err.Error() == constants.ErrorValidationFailed {
			ValidationErrorResponse(c, err)
			return
		}
		if err.Error() == constants.ErrorDuplicateEntry {
			BadRequestResponse(c, constants.ErrorDuplicateEntry, err)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.MessageTheatreTypeUpdated, theatreType)
}

// DeleteTheatreType handles DELETE /theatre-types/:id
func (ctrl *TheatreTypeController) DeleteTheatreType(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		BadRequestResponse(c, constants.ErrorInvalidUUID, err)
		return
	}

	err = ctrl.theatreTypeService.DeleteTheatreType(id)
	if err != nil {
		if err.Error() == constants.ErrorTheatreTypeNotFound {
			NotFoundResponse(c, constants.ErrorTheatreTypeNotFound)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.MessageTheatreTypeDeleted, nil)
}

// GetTheatreTypeByName handles GET /theatre-types/name/:name
func (ctrl *TheatreTypeController) GetTheatreTypeByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		BadRequestResponse(c, "name parameter is required", nil)
		return
	}

	theatreType, err := ctrl.theatreTypeService.GetTheatreTypeByName(name)
	if err != nil {
		if err.Error() == constants.ErrorTheatreTypeNotFound {
			NotFoundResponse(c, constants.ErrorTheatreTypeNotFound)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, theatreType)
}

// GetActiveTheatreTypes handles GET /theatre-types/active
func (ctrl *TheatreTypeController) GetActiveTheatreTypes(c *gin.Context) {
	theatreTypes, err := ctrl.theatreTypeService.GetActiveTheatreTypes()
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, theatreTypes)
}
