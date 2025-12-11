package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pos-mojosoft-so-service/internal/models"
	"pos-mojosoft-so-service/internal/utils"
)

type RemindedHandler struct {
	db *gorm.DB
}

func NewRemindedHandler(db *gorm.DB) *RemindedHandler {
	return &RemindedHandler{db: db}
}

// GetAll retrieves all reminded records
// @Summary Get all reminded records
// @Description Get list of all reminded records
// @Tags Reminded
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/reminded [get]
func (h *RemindedHandler) GetAll(c *gin.Context) {
	var reminded []models.Reminded

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query all reminded records (excluding soft deleted)
	if err := tenantDB.Find(&reminded).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve reminded records", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Reminded records retrieved successfully", reminded)
}

// GetByID retrieves a single reminded record by ID
// @Summary Get reminded record by ID
// @Description Get a single reminded record by its ID
// @Tags Reminded
// @Accept json
// @Produce json
// @Param id path int true "Reminded ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/reminded/{id} [get]
func (h *RemindedHandler) GetByID(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid reminded ID", nil)
		return
	}

	var reminded models.Reminded

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query reminded record by ID
	if err := tenantDB.First(&reminded, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Reminded record not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve reminded record", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Reminded record retrieved successfully", reminded)
}
