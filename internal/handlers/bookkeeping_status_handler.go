package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pos-mojosoft-so-service/internal/models"
	"pos-mojosoft-so-service/internal/utils"
)

type BookkeepingStatusHandler struct {
	db *gorm.DB
}

func NewBookkeepingStatusHandler(db *gorm.DB) *BookkeepingStatusHandler {
	return &BookkeepingStatusHandler{db: db}
}

// GetAll retrieves all bookkeeping statuses
// @Summary Get all bookkeeping statuses
// @Description Get list of all bookkeeping statuses
// @Tags BookkeepingStatus
// @Accept json
// @Produce json
// @Param name query string false "Filter by name (partial match)"
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/bookkeeping-status [get]
func (h *BookkeepingStatusHandler) GetAll(c *gin.Context) {
	var statuses []models.BookkeepingStatus

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Build query
	query := tenantDB.Model(&models.BookkeepingStatus{})

	// Apply filters
	if name := c.Query("name"); name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	// Execute query
	if err := query.Find(&statuses).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve bookkeeping statuses", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Bookkeeping statuses retrieved successfully", statuses)
}

// GetByID retrieves a single bookkeeping status by ID
// @Summary Get bookkeeping status by ID
// @Description Get a single bookkeeping status by its ID
// @Tags BookkeepingStatus
// @Accept json
// @Produce json
// @Param id path int true "Bookkeeping Status ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/bookkeeping-status/{id} [get]
func (h *BookkeepingStatusHandler) GetByID(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid bookkeeping status ID", nil)
		return
	}

	var status models.BookkeepingStatus

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query bookkeeping status by ID
	if err := tenantDB.First(&status, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Bookkeeping status not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve bookkeeping status", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Bookkeeping status retrieved successfully", status)
}
