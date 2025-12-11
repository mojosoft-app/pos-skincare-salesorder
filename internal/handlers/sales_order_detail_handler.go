package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pos-mojosoft-so-service/internal/models"
	"pos-mojosoft-so-service/internal/utils"
)

type SalesOrderDetailHandler struct {
	db *gorm.DB
}

func NewSalesOrderDetailHandler(db *gorm.DB) *SalesOrderDetailHandler {
	return &SalesOrderDetailHandler{db: db}
}

// SalesOrderDetailRequest represents the request body for creating/updating a sales order detail
type SalesOrderDetailRequest struct {
	SalesOrderID *int     `json:"sales_order_id"`
	ItemID       *int     `json:"item_id"`
	UnitID       *int     `json:"unit_id"`
	PromoterID   *int     `json:"promoter_id"`
	ItemName     *string  `json:"item_name"`
	Quantity     *int     `json:"quantity" binding:"required"`
	Price        *float64 `json:"price" binding:"required"`
	ItemTotal    *float64 `json:"item_total"`
	DiscountPct  *int     `json:"discount_pct"`
	UsedSessions *int     `json:"used_sessions"`
}

// GetAll retrieves all sales order details with optional filters
// @Summary Get all sales order details
// @Description Get list of all sales order details with optional filters
// @Tags SalesOrderDetail
// @Accept json
// @Produce json
// @Param sales_order_id query int false "Filter by sales order ID"
// @Param item_id query int false "Filter by item ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/sales-order-details [get]
func (h *SalesOrderDetailHandler) GetAll(c *gin.Context) {
	var details []models.SalesOrderDetail

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Build query
	query := tenantDB.Model(&models.SalesOrderDetail{})

	// Apply filters
	if salesOrderID := c.Query("sales_order_id"); salesOrderID != "" {
		query = query.Where("salesorder_id = ?", salesOrderID)
	}
	if itemID := c.Query("item_id"); itemID != "" {
		query = query.Where("item_id = ?", itemID)
	}

	// Execute query
	if err := query.Find(&details).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve sales order details", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Sales order details retrieved successfully", details)
}

// GetByID retrieves a single sales order detail by ID
// @Summary Get sales order detail by ID
// @Description Get a single sales order detail by its ID
// @Tags SalesOrderDetail
// @Accept json
// @Produce json
// @Param id path int true "Detail ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/sales-order-details/{id} [get]
func (h *SalesOrderDetailHandler) GetByID(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid detail ID", nil)
		return
	}

	var detail models.SalesOrderDetail

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query detail by ID
	if err := tenantDB.First(&detail, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Sales order detail not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve sales order detail", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Sales order detail retrieved successfully", detail)
}

// GetBySalesOrderID retrieves all details for a specific sales order
// @Summary Get sales order details by sales order ID
// @Description Get all details for a specific sales order
// @Tags SalesOrderDetail
// @Accept json
// @Produce json
// @Param sales_order_id path int true "Sales Order ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/sales-order-details/by-sales-order/{sales_order_id} [get]
func (h *SalesOrderDetailHandler) GetBySalesOrderID(c *gin.Context) {
	// Parse Sales Order ID from URL parameter
	salesOrderID, err := strconv.Atoi(c.Param("sales_order_id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid sales order ID", nil)
		return
	}

	var details []models.SalesOrderDetail

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query details by sales order ID
	if err := tenantDB.Where("salesorder_id = ?", salesOrderID).Find(&details).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve sales order details", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Sales order details retrieved successfully", details)
}

// Create creates a new sales order detail
// @Summary Create a new sales order detail
// @Description Create a new sales order detail
// @Tags SalesOrderDetail
// @Accept json
// @Produce json
// @Param request body SalesOrderDetailRequest true "Sales Order Detail data"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/sales-order-details [post]
func (h *SalesOrderDetailHandler) Create(c *gin.Context) {
	var req SalesOrderDetailRequest

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

	// Calculate item total if not provided
	itemTotal := req.ItemTotal
	if itemTotal == nil && req.Quantity != nil && req.Price != nil {
		calculated := float64(*req.Quantity) * *req.Price
		// Apply discount if provided
		if req.DiscountPct != nil && *req.DiscountPct > 0 {
			discount := calculated * float64(*req.DiscountPct) / 100
			calculated = calculated - discount
		}
		itemTotal = &calculated
	}

	// Create sales order detail
	detail := models.SalesOrderDetail{
		SalesOrderID: req.SalesOrderID,
		ItemID:       req.ItemID,
		UnitID:       req.UnitID,
		PromoterID:   req.PromoterID,
		ItemName:     req.ItemName,
		Quantity:     req.Quantity,
		Price:        req.Price,
		ItemTotal:    itemTotal,
		DiscountPct:  req.DiscountPct,
		UsedSessions: req.UsedSessions,
		CreatedBy:    &userIDInt64,
	}

	// Create detail
	if err := tenantDB.Create(&detail).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create sales order detail", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Sales order detail created successfully", detail)
}

// Update updates an existing sales order detail
// @Summary Update sales order detail
// @Description Update an existing sales order detail by ID
// @Tags SalesOrderDetail
// @Accept json
// @Produce json
// @Param id path int true "Detail ID"
// @Param request body SalesOrderDetailRequest true "Sales Order Detail data"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/sales-order-details/{id} [put]
func (h *SalesOrderDetailHandler) Update(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid detail ID", nil)
		return
	}

	var req SalesOrderDetailRequest

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

	// Check if detail exists
	var detail models.SalesOrderDetail
	if err := tenantDB.First(&detail, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Sales order detail not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve sales order detail", nil)
		return
	}

	// Calculate item total if not provided
	itemTotal := req.ItemTotal
	if itemTotal == nil && req.Quantity != nil && req.Price != nil {
		calculated := float64(*req.Quantity) * *req.Price
		// Apply discount if provided
		if req.DiscountPct != nil && *req.DiscountPct > 0 {
			discount := calculated * float64(*req.DiscountPct) / 100
			calculated = calculated - discount
		}
		itemTotal = &calculated
	}

	// Update fields
	detail.SalesOrderID = req.SalesOrderID
	detail.ItemID = req.ItemID
	detail.UnitID = req.UnitID
	detail.PromoterID = req.PromoterID
	detail.ItemName = req.ItemName
	detail.Quantity = req.Quantity
	detail.Price = req.Price
	detail.ItemTotal = itemTotal
	detail.DiscountPct = req.DiscountPct
	detail.UsedSessions = req.UsedSessions
	detail.UpdatedBy = &userIDInt64

	// Save updates
	if err := tenantDB.Save(&detail).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update sales order detail", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Sales order detail updated successfully", detail)
}

// Delete soft deletes a sales order detail
// @Summary Delete sales order detail
// @Description Soft delete a sales order detail by ID
// @Tags SalesOrderDetail
// @Accept json
// @Produce json
// @Param id path int true "Detail ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/sales-order-details/{id} [delete]
func (h *SalesOrderDetailHandler) Delete(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid detail ID", nil)
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

	// Check if detail exists
	var detail models.SalesOrderDetail
	if err := tenantDB.First(&detail, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Sales order detail not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve sales order detail", nil)
		return
	}

	// Set deleted_by before soft delete
	detail.DeletedBy = &userIDInt64
	if err := tenantDB.Save(&detail).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update deleted_by", err.Error())
		return
	}

	// Soft delete
	if err := tenantDB.Delete(&detail).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete sales order detail", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Sales order detail deleted successfully", nil)
}
