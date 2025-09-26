package controllers

import (
	"net/http"
	"theatre-management-system/src/constants"

	"github.com/gin-gonic/gin"
)

// APIResponse represents a standard API response structure
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SuccessResponse sends a successful response
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse sends an error response
func ErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	response := APIResponse{
		Success: false,
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
	}

	c.JSON(statusCode, response)
}

// ValidationErrorResponse sends a validation error response
func ValidationErrorResponse(c *gin.Context, err error) {
	ErrorResponse(c, http.StatusUnprocessableEntity, constants.ErrorValidationFailed, err)
}

// NotFoundResponse sends a not found error response
func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message, nil)
}

// BadRequestResponse sends a bad request error response
func BadRequestResponse(c *gin.Context, message string, err error) {
	ErrorResponse(c, http.StatusBadRequest, message, err)
}

// InternalServerErrorResponse sends an internal server error response
func InternalServerErrorResponse(c *gin.Context, err error) {
	ErrorResponse(c, http.StatusInternalServerError, constants.ErrorInternalServerError, err)
}

// PaginationParams represents pagination parameters
type PaginationParams struct {
	Limit  int `form:"limit" json:"limit"`
	Offset int `form:"offset" json:"offset"`
}

// GetPaginationParams extracts pagination parameters from request
func GetPaginationParams(c *gin.Context) PaginationParams {
	var params PaginationParams

	// Set defaults
	params.Limit = constants.DefaultLimit
	params.Offset = constants.DefaultOffset

	// Bind query parameters
	if err := c.ShouldBindQuery(&params); err != nil {
		// Use defaults if binding fails
		params.Limit = constants.DefaultLimit
		params.Offset = constants.DefaultOffset
	}

	// Apply limits
	if params.Limit <= 0 || params.Limit > constants.MaxLimit {
		params.Limit = constants.DefaultLimit
	}

	if params.Offset < 0 {
		params.Offset = constants.DefaultOffset
	}

	return params
}
