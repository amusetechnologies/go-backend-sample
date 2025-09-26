package controllers

import (
	"net/http"
	"theatre-management-system/src/constants"
	"theatre-management-system/src/dto"
	"theatre-management-system/src/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ShowTypeController handles HTTP requests for show types
type ShowTypeController struct {
	showTypeService interfaces.ShowTypeService
}

// NewShowTypeController creates a new show type controller
func NewShowTypeController(showTypeService interfaces.ShowTypeService) *ShowTypeController {
	return &ShowTypeController{
		showTypeService: showTypeService,
	}
}

// CreateShowType handles POST /show-types
func (ctrl *ShowTypeController) CreateShowType(c *gin.Context) {
	var showTypeDTO dto.ShowTypeBase

	if err := c.ShouldBindJSON(&showTypeDTO); err != nil {
		BadRequestResponse(c, constants.ErrorInvalidInput, err)
		return
	}

	showType, err := ctrl.showTypeService.CreateShowType(&showTypeDTO)
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

	SuccessResponse(c, http.StatusCreated, constants.MessageShowTypeCreated, showType)
}

// GetShowTypeByID handles GET /show-types/:id
func (ctrl *ShowTypeController) GetShowTypeByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		BadRequestResponse(c, constants.ErrorInvalidUUID, err)
		return
	}

	showType, err := ctrl.showTypeService.GetShowTypeByID(id)
	if err != nil {
		if err.Error() == constants.ErrorShowTypeNotFound {
			NotFoundResponse(c, constants.ErrorShowTypeNotFound)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, showType)
}

// GetAllShowTypes handles GET /show-types
func (ctrl *ShowTypeController) GetAllShowTypes(c *gin.Context) {
	params := GetPaginationParams(c)

	showTypes, err := ctrl.showTypeService.GetAllShowTypes(params.Limit, params.Offset)
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, showTypes)
}

// UpdateShowType handles PATCH /show-types/:id
func (ctrl *ShowTypeController) UpdateShowType(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		BadRequestResponse(c, constants.ErrorInvalidUUID, err)
		return
	}

	var showTypeDTO dto.ShowTypeBase
	if err := c.ShouldBindJSON(&showTypeDTO); err != nil {
		BadRequestResponse(c, constants.ErrorInvalidInput, err)
		return
	}

	showType, err := ctrl.showTypeService.UpdateShowType(id, &showTypeDTO)
	if err != nil {
		if err.Error() == constants.ErrorShowTypeNotFound {
			NotFoundResponse(c, constants.ErrorShowTypeNotFound)
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

	SuccessResponse(c, http.StatusOK, constants.MessageShowTypeUpdated, showType)
}

// DeleteShowType handles DELETE /show-types/:id
func (ctrl *ShowTypeController) DeleteShowType(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		BadRequestResponse(c, constants.ErrorInvalidUUID, err)
		return
	}

	err = ctrl.showTypeService.DeleteShowType(id)
	if err != nil {
		if err.Error() == constants.ErrorShowTypeNotFound {
			NotFoundResponse(c, constants.ErrorShowTypeNotFound)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.MessageShowTypeDeleted, nil)
}

// GetShowTypeByName handles GET /show-types/name/:name
func (ctrl *ShowTypeController) GetShowTypeByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		BadRequestResponse(c, "name parameter is required", nil)
		return
	}

	showType, err := ctrl.showTypeService.GetShowTypeByName(name)
	if err != nil {
		if err.Error() == constants.ErrorShowTypeNotFound {
			NotFoundResponse(c, constants.ErrorShowTypeNotFound)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, showType)
}

// GetActiveShowTypes handles GET /show-types/active
func (ctrl *ShowTypeController) GetActiveShowTypes(c *gin.Context) {
	showTypes, err := ctrl.showTypeService.GetActiveShowTypes()
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, showTypes)
}
