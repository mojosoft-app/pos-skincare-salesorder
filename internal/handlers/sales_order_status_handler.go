package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pos-mojosoft-so-service/internal/models"
	"pos-mojosoft-so-service/internal/utils"
)

type SalesOrderStatusHandler struct {
	db *gorm.DB
}

func NewSalesOrderStatusHandler(db *gorm.DB) *SalesOrderStatusHandler {
	return &SalesOrderStatusHandler{db: db}
}

// GetAll retrieves all sales order statuses
// @Summary Get all sales order statuses
// @Description Get list of all sales order statuses
// @Tags SalesOrderStatus
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/sales-order-status [get]
func (h *SalesOrderStatusHandler) GetAll(c *gin.Context) {
	var statuses []models.SalesOrderStatus

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query all statuses (excluding soft deleted)
	if err := tenantDB.Find(&statuses).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve sales order statuses", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Sales order statuses retrieved successfully", statuses)
}

// GetByID retrieves a single sales order status by ID
// @Summary Get sales order status by ID
// @Description Get a single sales order status by its ID
// @Tags SalesOrderStatus
// @Accept json
// @Produce json
// @Param id path int true "Status ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/sales-order-status/{id} [get]
func (h *SalesOrderStatusHandler) GetByID(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid status ID", nil)
		return
	}

	var status models.SalesOrderStatus

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query status by ID
	if err := tenantDB.First(&status, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Sales order status not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve sales order status", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Sales order status retrieved successfully", status)
}
