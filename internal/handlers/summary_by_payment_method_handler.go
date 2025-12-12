package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pos-mojosoft-so-service/internal/models"
	"pos-mojosoft-so-service/internal/utils"
)

type SummaryByPaymentMethodHandler struct {
	db *gorm.DB
}

func NewSummaryByPaymentMethodHandler(db *gorm.DB) *SummaryByPaymentMethodHandler {
	return &SummaryByPaymentMethodHandler{db: db}
}

// SummaryByPaymentMethodRequest represents the request body for creating/updating a summary
type SummaryByPaymentMethodRequest struct {
	BookkeepingID   *int     `json:"bookkeeping_id"`
	PaymentMethodID *int     `json:"payment_method_id"`
	Total           *float64 `json:"total"`
}

// GetAll retrieves all summaries with optional filters
// @Summary Get all summaries by payment method
// @Description Get list of all summaries by payment method with optional filters
// @Tags SummaryByPaymentMethod
// @Accept json
// @Produce json
// @Param bookkeeping_id query int false "Filter by bookkeeping ID"
// @Param payment_method_id query int false "Filter by payment method ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/summary-by-payment-method [get]
func (h *SummaryByPaymentMethodHandler) GetAll(c *gin.Context) {
	var summaries []models.SummaryByPaymentMethod

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Build query with preload relationships
	query := tenantDB.Model(&models.SummaryByPaymentMethod{}).
		Preload("Bookkeeping").
		Preload("PaymentMethod")

	// Apply filters
	if bookkeepingID := c.Query("bookkeeping_id"); bookkeepingID != "" {
		query = query.Where("bookkeeping_id = ?", bookkeepingID)
	}
	if paymentMethodID := c.Query("payment_method_id"); paymentMethodID != "" {
		query = query.Where("paymentmethod_id = ?", paymentMethodID)
	}

	// Execute query
	if err := query.Find(&summaries).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve summaries", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Summaries retrieved successfully", summaries)
}

// GetByID retrieves a single summary by ID
// @Summary Get summary by ID
// @Description Get a single summary by payment method by its ID
// @Tags SummaryByPaymentMethod
// @Accept json
// @Produce json
// @Param id path int true "Summary ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/summary-by-payment-method/{id} [get]
func (h *SummaryByPaymentMethodHandler) GetByID(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid summary ID", nil)
		return
	}

	var summary models.SummaryByPaymentMethod

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query summary by ID with relationships
	if err := tenantDB.Preload("Bookkeeping").
		Preload("PaymentMethod").
		First(&summary, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Summary not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve summary", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Summary retrieved successfully", summary)
}

// GetByBookkeepingID retrieves all summaries for a specific bookkeeping
// @Summary Get summaries by bookkeeping ID
// @Description Get all summaries for a specific bookkeeping record
// @Tags SummaryByPaymentMethod
// @Accept json
// @Produce json
// @Param bookkeeping_id path int true "Bookkeeping ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/summary-by-payment-method/by-bookkeeping/{bookkeeping_id} [get]
func (h *SummaryByPaymentMethodHandler) GetByBookkeepingID(c *gin.Context) {
	// Parse Bookkeeping ID from URL parameter
	bookkeepingID, err := strconv.Atoi(c.Param("bookkeeping_id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid bookkeeping ID", nil)
		return
	}

	var summaries []models.SummaryByPaymentMethod

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query summaries by bookkeeping ID with relationships
	if err := tenantDB.Preload("PaymentMethod").
		Where("bookkeeping_id = ?", bookkeepingID).
		Find(&summaries).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve summaries", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Summaries retrieved successfully", summaries)
}

// Create creates a new summary
// @Summary Create a new summary by payment method
// @Description Create a new summary by payment method
// @Tags SummaryByPaymentMethod
// @Accept json
// @Produce json
// @Param request body SummaryByPaymentMethodRequest true "Summary data"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/summary-by-payment-method [post]
func (h *SummaryByPaymentMethodHandler) Create(c *gin.Context) {
	var req SummaryByPaymentMethodRequest

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

	// Create summary
	summary := models.SummaryByPaymentMethod{
		BookkeepingID:   req.BookkeepingID,
		PaymentMethodID: req.PaymentMethodID,
		Total:           req.Total,
		CreatedBy:       &userIDInt64,
	}

	// Create summary
	if err := tenantDB.Create(&summary).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create summary", err.Error())
		return
	}

	// Load relationships
	tenantDB.Preload("Bookkeeping").Preload("PaymentMethod").First(&summary, summary.ID)

	utils.SuccessResponse(c, http.StatusCreated, "Summary created successfully", summary)
}

// Update updates an existing summary
// @Summary Update summary by payment method
// @Description Update an existing summary by payment method by ID
// @Tags SummaryByPaymentMethod
// @Accept json
// @Produce json
// @Param id path int true "Summary ID"
// @Param request body SummaryByPaymentMethodRequest true "Summary data"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/summary-by-payment-method/{id} [put]
func (h *SummaryByPaymentMethodHandler) Update(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid summary ID", nil)
		return
	}

	var req SummaryByPaymentMethodRequest

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

	// Check if summary exists
	var summary models.SummaryByPaymentMethod
	if err := tenantDB.First(&summary, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Summary not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve summary", nil)
		return
	}

	// Update fields
	summary.BookkeepingID = req.BookkeepingID
	summary.PaymentMethodID = req.PaymentMethodID
	summary.Total = req.Total
	summary.UpdatedBy = &userIDInt64

	// Save updates
	if err := tenantDB.Save(&summary).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update summary", err.Error())
		return
	}

	// Load relationships
	tenantDB.Preload("Bookkeeping").Preload("PaymentMethod").First(&summary, summary.ID)

	utils.SuccessResponse(c, http.StatusOK, "Summary updated successfully", summary)
}

// Delete soft deletes a summary
// @Summary Delete summary by payment method
// @Description Soft delete a summary by payment method by ID
// @Tags SummaryByPaymentMethod
// @Accept json
// @Produce json
// @Param id path int true "Summary ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/summary-by-payment-method/{id} [delete]
func (h *SummaryByPaymentMethodHandler) Delete(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid summary ID", nil)
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

	// Check if summary exists
	var summary models.SummaryByPaymentMethod
	if err := tenantDB.First(&summary, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Summary not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve summary", nil)
		return
	}

	// Set deleted_by before soft delete
	summary.DeletedBy = &userIDInt64
	if err := tenantDB.Save(&summary).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update deleted_by", err.Error())
		return
	}

	// Soft delete
	if err := tenantDB.Delete(&summary).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete summary", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Summary deleted successfully", nil)
}
