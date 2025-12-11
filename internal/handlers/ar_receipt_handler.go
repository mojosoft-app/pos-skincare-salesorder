package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"pos-mojosoft-so-service/internal/models"
	"pos-mojosoft-so-service/internal/utils"
)

type ARReceiptHandler struct {
	db *gorm.DB
}

func NewARReceiptHandler(db *gorm.DB) *ARReceiptHandler {
	return &ARReceiptHandler{db: db}
}

// CreateARReceiptRequest represents the request body for creating an AR receipt
type CreateARReceiptRequest struct {
	LocationID      *int                          `json:"location_id"`
	CustomerID      *int                          `json:"customer_id" binding:"required"`
	PaymentMethodID *int                          `json:"payment_method_id"`
	DocNumber       *int                          `json:"doc_number"`
	DocDate         *string                       `json:"doc_date"`
	PostedDate      *string                       `json:"posted_date"`
	TotalAmount     *float64                      `json:"total_amount"`
	Note            *string                       `json:"note"`
	StatusID        *int                          `json:"status_id"`
	Details         []CreateARReceiptDetailRequest `json:"details"`
}

type CreateARReceiptDetailRequest struct {
	SalesOrderID  *int     `json:"sales_order_id"`
	ReceiptAmount *float64 `json:"receipt_amount"`
}

// GetAll retrieves all AR receipts with optional filters
// @Summary Get all AR receipts
// @Description Get list of all AR receipts with optional pagination and filters
// @Tags ARReceipt
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param customer_id query int false "Filter by customer ID"
// @Param status_id query int false "Filter by status ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/ar-receipts [get]
func (h *ARReceiptHandler) GetAll(c *gin.Context) {
	var arReceipts []models.ARReceipt

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Build query
	query := tenantDB.Model(&models.ARReceipt{})

	// Apply filters
	if customerID := c.Query("customer_id"); customerID != "" {
		query = query.Where("customer_id = ?", customerID)
	}
	if statusID := c.Query("status_id"); statusID != "" {
		query = query.Where("status_id = ?", statusID)
	}

	// Preload relationships
	query = query.Preload("Details")

	// Execute query
	if err := query.Find(&arReceipts).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve AR receipts", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "AR receipts retrieved successfully", arReceipts)
}

// GetByID retrieves a single AR receipt by ID
// @Summary Get AR receipt by ID
// @Description Get a single AR receipt by its ID with details
// @Tags ARReceipt
// @Accept json
// @Produce json
// @Param id path string true "AR Receipt ID (UUID)"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/ar-receipts/{id} [get]
func (h *ARReceiptHandler) GetByID(c *gin.Context) {
	// Parse UUID from URL parameter
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid AR receipt ID", nil)
		return
	}

	var arReceipt models.ARReceipt

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query AR receipt by ID with relationships
	if err := tenantDB.Preload("Details").
		First(&arReceipt, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "AR receipt not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve AR receipt", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "AR receipt retrieved successfully", arReceipt)
}

// Create creates a new AR receipt
// @Summary Create a new AR receipt
// @Description Create a new AR receipt with details
// @Tags ARReceipt
// @Accept json
// @Produce json
// @Param request body CreateARReceiptRequest true "AR Receipt data"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/ar-receipts [post]
func (h *ARReceiptHandler) Create(c *gin.Context) {
	var req CreateARReceiptRequest

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

	// Create AR receipt
	arReceipt := models.ARReceipt{
		LocationID:      req.LocationID,
		CustomerID:      req.CustomerID,
		PaymentMethodID: req.PaymentMethodID,
		DocNumber:       req.DocNumber,
		TotalAmount:     req.TotalAmount,
		Note:            req.Note,
		StatusID:        req.StatusID,
		CreatedBy:       &userIDInt64,
	}

	// Begin transaction
	tx := tenantDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create AR receipt
	if err := tx.Create(&arReceipt).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create AR receipt", err.Error())
		return
	}

	// Create details if provided
	if len(req.Details) > 0 {
		for _, detailReq := range req.Details {
			detail := models.ARReceiptDetail{
				SalesOrderID:  detailReq.SalesOrderID,
				ReceiptAmount: detailReq.ReceiptAmount,
				CreatedBy:     &userIDInt64,
			}
			if err := tx.Create(&detail).Error; err != nil {
				tx.Rollback()
				utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create AR receipt detail", err.Error())
				return
			}
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to commit transaction", err.Error())
		return
	}

	// Load the created AR receipt with relationships
	tenantDB.Preload("Details").First(&arReceipt, "id = ?", arReceipt.ID)

	utils.SuccessResponse(c, http.StatusCreated, "AR receipt created successfully", arReceipt)
}

// Update updates an existing AR receipt
// @Summary Update AR receipt
// @Description Update an existing AR receipt by ID
// @Tags ARReceipt
// @Accept json
// @Produce json
// @Param id path string true "AR Receipt ID (UUID)"
// @Param request body CreateARReceiptRequest true "AR Receipt data"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/ar-receipts/{id} [put]
func (h *ARReceiptHandler) Update(c *gin.Context) {
	// Parse UUID from URL parameter
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid AR receipt ID", nil)
		return
	}

	var req CreateARReceiptRequest

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

	// Check if AR receipt exists
	var arReceipt models.ARReceipt
	if err := tenantDB.First(&arReceipt, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "AR receipt not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve AR receipt", nil)
		return
	}

	// Update fields
	arReceipt.LocationID = req.LocationID
	arReceipt.CustomerID = req.CustomerID
	arReceipt.PaymentMethodID = req.PaymentMethodID
	arReceipt.DocNumber = req.DocNumber
	arReceipt.TotalAmount = req.TotalAmount
	arReceipt.Note = req.Note
	arReceipt.StatusID = req.StatusID
	arReceipt.UpdatedBy = &userIDInt64

	// Save updates
	if err := tenantDB.Save(&arReceipt).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update AR receipt", err.Error())
		return
	}

	// Load updated AR receipt with relationships
	tenantDB.Preload("Details").First(&arReceipt, "id = ?", arReceipt.ID)

	utils.SuccessResponse(c, http.StatusOK, "AR receipt updated successfully", arReceipt)
}

// Delete soft deletes an AR receipt
// @Summary Delete AR receipt
// @Description Soft delete an AR receipt by ID
// @Tags ARReceipt
// @Accept json
// @Produce json
// @Param id path string true "AR Receipt ID (UUID)"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/ar-receipts/{id} [delete]
func (h *ARReceiptHandler) Delete(c *gin.Context) {
	// Parse UUID from URL parameter
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid AR receipt ID", nil)
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

	// Check if AR receipt exists
	var arReceipt models.ARReceipt
	if err := tenantDB.First(&arReceipt, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "AR receipt not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve AR receipt", nil)
		return
	}

	// Set deleted_by before soft delete
	arReceipt.DeletedBy = &userIDInt64
	if err := tenantDB.Save(&arReceipt).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update deleted_by", err.Error())
		return
	}

	// Soft delete
	if err := tenantDB.Delete(&arReceipt).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete AR receipt", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "AR receipt deleted successfully", nil)
}
