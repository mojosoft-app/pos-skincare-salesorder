package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pos-mojosoft-so-service/internal/models"
	"pos-mojosoft-so-service/internal/utils"
)

type PaymentMethodHandler struct {
	db *gorm.DB
}

func NewPaymentMethodHandler(db *gorm.DB) *PaymentMethodHandler {
	return &PaymentMethodHandler{db: db}
}

// PaymentMethodRequest represents the request body for creating/updating a payment method
type PaymentMethodRequest struct {
	Name *string `json:"name" binding:"required"`
}

// GetAll retrieves all payment methods
// @Summary Get all payment methods
// @Description Get list of all payment methods with optional filters
// @Tags PaymentMethod
// @Accept json
// @Produce json
// @Param name query string false "Filter by name (partial match)"
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/payment-method [get]
func (h *PaymentMethodHandler) GetAll(c *gin.Context) {
	var paymentMethods []models.PaymentMethod

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Build query
	query := tenantDB.Model(&models.PaymentMethod{})

	// Apply filters
	if name := c.Query("name"); name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	// Execute query
	if err := query.Find(&paymentMethods).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve payment methods", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Payment methods retrieved successfully", paymentMethods)
}

// GetByID retrieves a single payment method by ID
// @Summary Get payment method by ID
// @Description Get a single payment method by its ID
// @Tags PaymentMethod
// @Accept json
// @Produce json
// @Param id path int true "Payment Method ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/payment-method/{id} [get]
func (h *PaymentMethodHandler) GetByID(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid payment method ID", nil)
		return
	}

	var paymentMethod models.PaymentMethod

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query payment method by ID
	if err := tenantDB.First(&paymentMethod, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Payment method not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve payment method", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Payment method retrieved successfully", paymentMethod)
}

// Create creates a new payment method
// @Summary Create a new payment method
// @Description Create a new payment method
// @Tags PaymentMethod
// @Accept json
// @Produce json
// @Param request body PaymentMethodRequest true "Payment method data"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/payment-method [post]
func (h *PaymentMethodHandler) Create(c *gin.Context) {
	var req PaymentMethodRequest

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

	// Create payment method
	paymentMethod := models.PaymentMethod{
		Name:      req.Name,
		CreatedBy: &userIDInt64,
	}

	// Create payment method
	if err := tenantDB.Create(&paymentMethod).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create payment method", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Payment method created successfully", paymentMethod)
}

// Update updates an existing payment method
// @Summary Update payment method
// @Description Update an existing payment method by ID
// @Tags PaymentMethod
// @Accept json
// @Produce json
// @Param id path int true "Payment Method ID"
// @Param request body PaymentMethodRequest true "Payment method data"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/payment-method/{id} [put]
func (h *PaymentMethodHandler) Update(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid payment method ID", nil)
		return
	}

	var req PaymentMethodRequest

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

	// Check if payment method exists
	var paymentMethod models.PaymentMethod
	if err := tenantDB.First(&paymentMethod, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Payment method not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve payment method", nil)
		return
	}

	// Update fields
	paymentMethod.Name = req.Name
	paymentMethod.UpdatedBy = &userIDInt64

	// Save updates
	if err := tenantDB.Save(&paymentMethod).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update payment method", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Payment method updated successfully", paymentMethod)
}

// Delete soft deletes a payment method
// @Summary Delete payment method
// @Description Soft delete a payment method by ID
// @Tags PaymentMethod
// @Accept json
// @Produce json
// @Param id path int true "Payment Method ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/payment-method/{id} [delete]
func (h *PaymentMethodHandler) Delete(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid payment method ID", nil)
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

	// Check if payment method exists
	var paymentMethod models.PaymentMethod
	if err := tenantDB.First(&paymentMethod, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Payment method not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve payment method", nil)
		return
	}

	// Set deleted_by before soft delete
	paymentMethod.DeletedBy = &userIDInt64
	if err := tenantDB.Save(&paymentMethod).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update deleted_by", err.Error())
		return
	}

	// Soft delete
	if err := tenantDB.Delete(&paymentMethod).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete payment method", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Payment method deleted successfully", nil)
}
