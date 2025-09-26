package controllers

import (
	"net/http"
	"theatre-management-system/src/constants"
	"theatre-management-system/src/dto"
	"theatre-management-system/src/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ShowController handles HTTP requests for shows
type ShowController struct {
	showService interfaces.ShowService
}

// NewShowController creates a new show controller
func NewShowController(showService interfaces.ShowService) *ShowController {
	return &ShowController{
		showService: showService,
	}
}

// CreateShow handles POST /shows
func (ctrl *ShowController) CreateShow(c *gin.Context) {
	var showDTO dto.ShowBase

	if err := c.ShouldBindJSON(&showDTO); err != nil {
		BadRequestResponse(c, constants.ErrorInvalidInput, err)
		return
	}

	show, err := ctrl.showService.CreateShow(&showDTO)
	if err != nil {
		if err.Error() == constants.ErrorValidationFailed {
			ValidationErrorResponse(c, err)
			return
		}
		if err.Error() == constants.ErrorTheatreNotFound || err.Error() == constants.ErrorShowTypeNotFound {
			BadRequestResponse(c, err.Error(), err)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusCreated, constants.MessageShowCreated, show)
}

// GetShowByID handles GET /shows/:id
func (ctrl *ShowController) GetShowByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		BadRequestResponse(c, constants.ErrorInvalidUUID, err)
		return
	}

	show, err := ctrl.showService.GetShowByID(id)
	if err != nil {
		if err.Error() == constants.ErrorShowNotFound {
			NotFoundResponse(c, constants.ErrorShowNotFound)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, show)
}

// GetAllShows handles GET /shows
func (ctrl *ShowController) GetAllShows(c *gin.Context) {
	params := GetPaginationParams(c)

	shows, err := ctrl.showService.GetAllShows(params.Limit, params.Offset)
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, shows)
}

// UpdateShow handles PATCH /shows/:id
func (ctrl *ShowController) UpdateShow(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		BadRequestResponse(c, constants.ErrorInvalidUUID, err)
		return
	}

	var showDTO dto.ShowBase
	if err := c.ShouldBindJSON(&showDTO); err != nil {
		BadRequestResponse(c, constants.ErrorInvalidInput, err)
		return
	}

	show, err := ctrl.showService.UpdateShow(id, &showDTO)
	if err != nil {
		if err.Error() == constants.ErrorShowNotFound {
			NotFoundResponse(c, constants.ErrorShowNotFound)
			return
		}
		if err.Error() == constants.ErrorValidationFailed {
			ValidationErrorResponse(c, err)
			return
		}
		if err.Error() == constants.ErrorTheatreNotFound || err.Error() == constants.ErrorShowTypeNotFound {
			BadRequestResponse(c, err.Error(), err)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.MessageShowUpdated, show)
}

// DeleteShow handles DELETE /shows/:id
func (ctrl *ShowController) DeleteShow(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		BadRequestResponse(c, constants.ErrorInvalidUUID, err)
		return
	}

	err = ctrl.showService.DeleteShow(id)
	if err != nil {
		if err.Error() == constants.ErrorShowNotFound {
			NotFoundResponse(c, constants.ErrorShowNotFound)
			return
		}
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.MessageShowDeleted, nil)
}

// GetShowsByTheatreID handles GET /shows/theatre/:theatreId
func (ctrl *ShowController) GetShowsByTheatreID(c *gin.Context) {
	theatreIDParam := c.Param("theatreId")
	theatreID, err := uuid.Parse(theatreIDParam)
	if err != nil {
		BadRequestResponse(c, constants.ErrorInvalidUUID, err)
		return
	}

	shows, err := ctrl.showService.GetShowsByTheatreID(theatreID)
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, shows)
}

// GetShowsByShowTypeID handles GET /shows/type/:typeId
func (ctrl *ShowController) GetShowsByShowTypeID(c *gin.Context) {
	typeIDParam := c.Param("typeId")
	typeID, err := uuid.Parse(typeIDParam)
	if err != nil {
		BadRequestResponse(c, constants.ErrorInvalidUUID, err)
		return
	}

	shows, err := ctrl.showService.GetShowsByShowTypeID(typeID)
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, shows)
}

// GetFeaturedShows handles GET /shows/featured
func (ctrl *ShowController) GetFeaturedShows(c *gin.Context) {
	shows, err := ctrl.showService.GetFeaturedShows()
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, shows)
}

// GetActiveShows handles GET /shows/active
func (ctrl *ShowController) GetActiveShows(c *gin.Context) {
	shows, err := ctrl.showService.GetActiveShows()
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, shows)
}

// GetCurrentShows handles GET /shows/current
func (ctrl *ShowController) GetCurrentShows(c *gin.Context) {
	shows, err := ctrl.showService.GetCurrentShows()
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, shows)
}

// GetUpcomingShows handles GET /shows/upcoming
func (ctrl *ShowController) GetUpcomingShows(c *gin.Context) {
	shows, err := ctrl.showService.GetUpcomingShows()
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, shows)
}

// SearchShows handles GET /shows/search
func (ctrl *ShowController) SearchShows(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		BadRequestResponse(c, "search query is required", nil)
		return
	}

	shows, err := ctrl.showService.SearchShows(query)
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	SuccessResponse(c, http.StatusOK, constants.StatusOK, shows)
}
