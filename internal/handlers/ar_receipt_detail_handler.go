package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pos-mojosoft-so-service/internal/models"
	"pos-mojosoft-so-service/internal/utils"
)

type ARReceiptDetailHandler struct {
	db *gorm.DB
}

func NewARReceiptDetailHandler(db *gorm.DB) *ARReceiptDetailHandler {
	return &ARReceiptDetailHandler{db: db}
}

// CreateARReceiptDetailRequest represents the request body for creating an AR receipt detail
type CreateARReceiptDetailRequest struct {
	ARReceiptID   *int     `json:"ar_receipt_id" binding:"required"`
	SalesOrderID  *int     `json:"sales_order_id"`
	ReceiptAmount *float64 `json:"receipt_amount"`
}

// GetAll retrieves all AR receipt details with optional filters
// @Summary Get all AR receipt details
// @Description Get list of all AR receipt details with optional filters
// @Tags ARReceiptDetail
// @Accept json
// @Produce json
// @Param ar_receipt_id query int false "Filter by AR receipt ID"
// @Param sales_order_id query int false "Filter by sales order ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/ar-receipt-details [get]
func (h *ARReceiptDetailHandler) GetAll(c *gin.Context) {
	var details []models.ARReceiptDetail

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Build query
	query := tenantDB.Model(&models.ARReceiptDetail{})

	// Apply filters
	if arReceiptID := c.Query("ar_receipt_id"); arReceiptID != "" {
		query = query.Where("arreceipt_id = ?", arReceiptID)
	}
	if salesOrderID := c.Query("sales_order_id"); salesOrderID != "" {
		query = query.Where("salesorder_id = ?", salesOrderID)
	}

	// Execute query
	if err := query.Find(&details).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve AR receipt details", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "AR receipt details retrieved successfully", details)
}

// GetByID retrieves a single AR receipt detail by ID
// @Summary Get AR receipt detail by ID
// @Description Get a single AR receipt detail by its ID
// @Tags ARReceiptDetail
// @Accept json
// @Produce json
// @Param id path int true "AR Receipt Detail ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/ar-receipt-details/{id} [get]
func (h *ARReceiptDetailHandler) GetByID(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid AR receipt detail ID", nil)
		return
	}

	var detail models.ARReceiptDetail

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query AR receipt detail by ID
	if err := tenantDB.First(&detail, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "AR receipt detail not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve AR receipt detail", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "AR receipt detail retrieved successfully", detail)
}

// GetByARReceiptID retrieves all details for a specific AR receipt
// @Summary Get AR receipt details by AR receipt ID
// @Description Get all details for a specific AR receipt
// @Tags ARReceiptDetail
// @Accept json
// @Produce json
// @Param ar_receipt_id path int true "AR Receipt ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/ar-receipt-details/by-ar-receipt/{ar_receipt_id} [get]
func (h *ARReceiptDetailHandler) GetByARReceiptID(c *gin.Context) {
	// Parse AR receipt ID from URL parameter
	arReceiptID, err := strconv.Atoi(c.Param("ar_receipt_id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid AR receipt ID", nil)
		return
	}

	var details []models.ARReceiptDetail

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query details by AR receipt ID
	if err := tenantDB.Where("arreceipt_id = ?", arReceiptID).Find(&details).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve AR receipt details", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "AR receipt details retrieved successfully", details)
}

// Create creates a new AR receipt detail
// @Summary Create a new AR receipt detail
// @Description Create a new AR receipt detail
// @Tags ARReceiptDetail
// @Accept json
// @Produce json
// @Param request body CreateARReceiptDetailRequest true "AR Receipt Detail data"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/ar-receipt-details [post]
func (h *ARReceiptDetailHandler) Create(c *gin.Context) {
	var req CreateARReceiptDetailRequest

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

	// Create AR receipt detail
	detail := models.ARReceiptDetail{
		ARReceiptID:   req.ARReceiptID,
		SalesOrderID:  req.SalesOrderID,
		ReceiptAmount: req.ReceiptAmount,
		CreatedBy:     &userIDInt64,
	}

	// Create the detail
	if err := tenantDB.Create(&detail).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create AR receipt detail", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "AR receipt detail created successfully", detail)
}

// Update updates an existing AR receipt detail
// @Summary Update AR receipt detail
// @Description Update an existing AR receipt detail by ID
// @Tags ARReceiptDetail
// @Accept json
// @Produce json
// @Param id path int true "AR Receipt Detail ID"
// @Param request body CreateARReceiptDetailRequest true "AR Receipt Detail data"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/ar-receipt-details/{id} [put]
func (h *ARReceiptDetailHandler) Update(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid AR receipt detail ID", nil)
		return
	}

	var req CreateARReceiptDetailRequest

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

	// Check if AR receipt detail exists
	var detail models.ARReceiptDetail
	if err := tenantDB.First(&detail, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "AR receipt detail not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve AR receipt detail", nil)
		return
	}

	// Update fields
	detail.ARReceiptID = req.ARReceiptID
	detail.SalesOrderID = req.SalesOrderID
	detail.ReceiptAmount = req.ReceiptAmount
	detail.UpdatedBy = &userIDInt64

	// Save updates
	if err := tenantDB.Save(&detail).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update AR receipt detail", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "AR receipt detail updated successfully", detail)
}

// Delete soft deletes an AR receipt detail
// @Summary Delete AR receipt detail
// @Description Soft delete an AR receipt detail by ID
// @Tags ARReceiptDetail
// @Accept json
// @Produce json
// @Param id path int true "AR Receipt Detail ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/ar-receipt-details/{id} [delete]
func (h *ARReceiptDetailHandler) Delete(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid AR receipt detail ID", nil)
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

	// Check if AR receipt detail exists
	var detail models.ARReceiptDetail
	if err := tenantDB.First(&detail, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "AR receipt detail not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve AR receipt detail", nil)
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
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete AR receipt detail", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "AR receipt detail deleted successfully", nil)
}
