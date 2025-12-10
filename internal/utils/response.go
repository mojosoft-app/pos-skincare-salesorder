package utils

import (
	"net/http"

	"pos-mojosoft-so-service/internal/models"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	response := models.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
		Errors:  nil,
	}
	c.JSON(statusCode, response)
}

func ErrorResponse(c *gin.Context, statusCode int, message string, errors interface{}) {
	response := models.APIResponse{
		Success: false,
		Message: message,
		Data:    nil,
		Errors:  errors,
	}
	c.JSON(statusCode, response)
}

func ValidationErrorResponse(c *gin.Context, errors []models.ErrorDetail) {
	ErrorResponse(c, http.StatusBadRequest, "Validation failed", errors)
}

func UnauthorizedResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, message, nil)
}

func ForbiddenResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusForbidden, message, nil)
}

func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message, nil)
}

func InternalErrorResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusInternalServerError, message, nil)
}

func PaginatedResponse(c *gin.Context, statusCode int, message string, data interface{}, meta models.PaginationMeta) {
	response := models.PaginatedResponse{
		Success:    true,
		Message:    message,
		Data:       data,
		Pagination: meta,
	}
	c.JSON(statusCode, response)
}
