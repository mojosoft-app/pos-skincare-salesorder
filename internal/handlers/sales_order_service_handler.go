package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pos-mojosoft-so-service/internal/models"
	"pos-mojosoft-so-service/internal/utils"
)

type SalesOrderServiceHandler struct {
	db *gorm.DB
}

func NewSalesOrderServiceHandler(db *gorm.DB) *SalesOrderServiceHandler {
	return &SalesOrderServiceHandler{db: db}
}

// SalesOrderServiceRequest represents the request body for creating/updating a sales order service
type SalesOrderServiceRequest struct {
	SalesOrderID       *int    `json:"sales_order_id"`
	SalesOrderDetailID *int    `json:"sales_order_detail_id"`
	ServiceID          *int    `json:"service_id"`
	TreatmentID        *int    `json:"treatment_id"`
	MessageLogDetailID *string `json:"message_log_detail_id"`
	RemindedID         *int    `json:"reminded_id"`
	ServiceName        *string `json:"service_name"`
	Treated            *bool   `json:"treated"`
	Schedule           *string `json:"schedule"`
}

// GetAll retrieves all sales order services with optional filters
// @Summary Get all sales order services
// @Description Get list of all sales order services with optional filters
// @Tags SalesOrderService
// @Accept json
// @Produce json
// @Param sales_order_id query int false "Filter by sales order ID"
// @Param treated query bool false "Filter by treated status"
// @Param service_id query int false "Filter by service ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/sales-order-services [get]
func (h *SalesOrderServiceHandler) GetAll(c *gin.Context) {
	var services []models.SalesOrderService

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Build query
	query := tenantDB.Model(&models.SalesOrderService{})

	// Apply filters
	if salesOrderID := c.Query("sales_order_id"); salesOrderID != "" {
		query = query.Where("salesorder_id = ?", salesOrderID)
	}
	if treated := c.Query("treated"); treated != "" {
		query = query.Where("treated = ?", treated)
	}
	if serviceID := c.Query("service_id"); serviceID != "" {
		query = query.Where("service_id = ?", serviceID)
	}

	// Execute query
	if err := query.Find(&services).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve sales order services", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Sales order services retrieved successfully", services)
}

// GetByID retrieves a single sales order service by ID
// @Summary Get sales order service by ID
// @Description Get a single sales order service by its ID
// @Tags SalesOrderService
// @Accept json
// @Produce json
// @Param id path int true "Service ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/sales-order-services/{id} [get]
func (h *SalesOrderServiceHandler) GetByID(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid service ID", nil)
		return
	}

	var service models.SalesOrderService

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query service by ID
	if err := tenantDB.First(&service, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Sales order service not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve sales order service", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Sales order service retrieved successfully", service)
}

// Create creates a new sales order service
// @Summary Create a new sales order service
// @Description Create a new sales order service
// @Tags SalesOrderService
// @Accept json
// @Produce json
// @Param request body SalesOrderServiceRequest true "Sales Order Service data"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/sales-order-services [post]
func (h *SalesOrderServiceHandler) Create(c *gin.Context) {
	var req SalesOrderServiceRequest

	// Bind and validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Get user ID from context (set by auth middleware)
	userID, _ := c.Get("user_id")
	userIDInt64 := int64(userID.(uint))

	// Create sales order service
	service := models.SalesOrderService{
		SalesOrderID:       req.SalesOrderID,
		SalesOrderDetailID: req.SalesOrderDetailID,
		ServiceID:          req.ServiceID,
		TreatmentID:        req.TreatmentID,
		MessageLogDetailID: req.MessageLogDetailID,
		RemindedID:         req.RemindedID,
		ServiceName:        req.ServiceName,
		Treated:            req.Treated,
		CreatedBy:          &userIDInt64,
	}

	// Create service
	if err := tenantDB.Create(&service).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create sales order service", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Sales order service created successfully", service)
}

// Update updates an existing sales order service
// @Summary Update sales order service
// @Description Update an existing sales order service by ID
// @Tags SalesOrderService
// @Accept json
// @Produce json
// @Param id path int true "Service ID"
// @Param request body SalesOrderServiceRequest true "Sales Order Service data"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/sales-order-services/{id} [put]
func (h *SalesOrderServiceHandler) Update(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid service ID", nil)
		return
	}

	var req SalesOrderServiceRequest

	// Bind and validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Get user ID from context
	userID, _ := c.Get("user_id")
	userIDInt64 := int64(userID.(uint))

	// Check if service exists
	var service models.SalesOrderService
	if err := tenantDB.First(&service, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Sales order service not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve sales order service", nil)
		return
	}

	// Update fields
	service.SalesOrderID = req.SalesOrderID
	service.SalesOrderDetailID = req.SalesOrderDetailID
	service.ServiceID = req.ServiceID
	service.TreatmentID = req.TreatmentID
	service.MessageLogDetailID = req.MessageLogDetailID
	service.RemindedID = req.RemindedID
	service.ServiceName = req.ServiceName
	service.Treated = req.Treated
	service.UpdatedBy = &userIDInt64

	// Save updates
	if err := tenantDB.Save(&service).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update sales order service", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Sales order service updated successfully", service)
}

// Delete soft deletes a sales order service
// @Summary Delete sales order service
// @Description Soft delete a sales order service by ID
// @Tags SalesOrderService
// @Accept json
// @Produce json
// @Param id path int true "Service ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/sales-order-services/{id} [delete]
func (h *SalesOrderServiceHandler) Delete(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid service ID", nil)
		return
	}

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Get user ID from context
	userID, _ := c.Get("user_id")
	userIDInt64 := int64(userID.(uint))

	// Check if service exists
	var service models.SalesOrderService
	if err := tenantDB.First(&service, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Sales order service not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve sales order service", nil)
		return
	}

	// Set deleted_by before soft delete
	service.DeletedBy = &userIDInt64
	if err := tenantDB.Save(&service).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update deleted_by", err.Error())
		return
	}

	// Soft delete
	if err := tenantDB.Delete(&service).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete sales order service", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Sales order service deleted successfully", nil)
}

// MarkAsTreated marks a service as treated
// @Summary Mark service as treated
// @Description Mark a sales order service as treated
// @Tags SalesOrderService
// @Accept json
// @Produce json
// @Param id path int true "Service ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/sales-order-services/{id}/mark-treated [patch]
func (h *SalesOrderServiceHandler) MarkAsTreated(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid service ID", nil)
		return
	}

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Get user ID from context
	userID, _ := c.Get("user_id")
	userIDInt64 := int64(userID.(uint))

	// Check if service exists
	var service models.SalesOrderService
	if err := tenantDB.First(&service, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Sales order service not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve sales order service", nil)
		return
	}

	// Mark as treated
	treated := true
	service.Treated = &treated
	service.UpdatedBy = &userIDInt64

	// Save updates
	if err := tenantDB.Save(&service).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to mark service as treated", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Service marked as treated successfully", service)
}
