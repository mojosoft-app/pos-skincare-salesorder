package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pos-mojosoft-so-service/internal/models"
	"pos-mojosoft-so-service/internal/utils"
)

type BookkeepingDetailHandler struct {
	db *gorm.DB
}

func NewBookkeepingDetailHandler(db *gorm.DB) *BookkeepingDetailHandler {
	return &BookkeepingDetailHandler{db: db}
}

// BookkeepingDetailRequest represents the request body for creating/updating a bookkeeping detail
type BookkeepingDetailRequest struct {
	BookkeepingID   *int       `json:"bookkeeping_id"`
	TypeID          *int       `json:"type_id"`
	CategoryID      *int       `json:"category_id"`
	PaymentMethodID *int       `json:"payment_method_id"`
	PostedDate      *time.Time `json:"posted_date"`
	DocNumber       *string    `json:"doc_number"`
	Income          *float64   `json:"income"`
	Expanse         *float64   `json:"expanse"`
	Description     *string    `json:"description"`
}

// GetAll retrieves all bookkeeping details with optional filters
// @Summary Get all bookkeeping details
// @Description Get list of all bookkeeping details with optional filters
// @Tags BookkeepingDetail
// @Accept json
// @Produce json
// @Param bookkeeping_id query int false "Filter by bookkeeping ID"
// @Param type_id query int false "Filter by transaction type ID"
// @Param category_id query int false "Filter by category ID"
// @Param payment_method_id query int false "Filter by payment method ID"
// @Param posted_date_from query string false "Filter by posted date from (YYYY-MM-DD)"
// @Param posted_date_to query string false "Filter by posted date to (YYYY-MM-DD)"
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/bookkeeping-detail [get]
func (h *BookkeepingDetailHandler) GetAll(c *gin.Context) {
	var details []models.BookkeepingDetail

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Build query with preload relationships
	query := tenantDB.Model(&models.BookkeepingDetail{}).
		Preload("Bookkeeping").
		Preload("Type").
		Preload("Category").
		Preload("PaymentMethod")

	// Apply filters
	if bookkeepingID := c.Query("bookkeeping_id"); bookkeepingID != "" {
		query = query.Where("bookkeeping_id = ?", bookkeepingID)
	}
	if typeID := c.Query("type_id"); typeID != "" {
		query = query.Where("type_id = ?", typeID)
	}
	if categoryID := c.Query("category_id"); categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}
	if paymentMethodID := c.Query("payment_method_id"); paymentMethodID != "" {
		query = query.Where("paymentmethod_id = ?", paymentMethodID)
	}
	if postedDateFrom := c.Query("posted_date_from"); postedDateFrom != "" {
		query = query.Where("posteddate >= ?", postedDateFrom)
	}
	if postedDateTo := c.Query("posted_date_to"); postedDateTo != "" {
		query = query.Where("posteddate <= ?", postedDateTo)
	}

	// Execute query
	if err := query.Find(&details).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve bookkeeping details", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Bookkeeping details retrieved successfully", details)
}

// GetByID retrieves a single bookkeeping detail by ID
// @Summary Get bookkeeping detail by ID
// @Description Get a single bookkeeping detail by its ID
// @Tags BookkeepingDetail
// @Accept json
// @Produce json
// @Param id path int true "Bookkeeping Detail ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/bookkeeping-detail/{id} [get]
func (h *BookkeepingDetailHandler) GetByID(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid bookkeeping detail ID", nil)
		return
	}

	var detail models.BookkeepingDetail

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query bookkeeping detail by ID with relationships
	if err := tenantDB.Preload("Bookkeeping").
		Preload("Type").
		Preload("Category").
		Preload("PaymentMethod").
		First(&detail, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Bookkeeping detail not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve bookkeeping detail", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Bookkeeping detail retrieved successfully", detail)
}

// GetByBookkeepingID retrieves all details for a specific bookkeeping
// @Summary Get bookkeeping details by bookkeeping ID
// @Description Get all details for a specific bookkeeping record
// @Tags BookkeepingDetail
// @Accept json
// @Produce json
// @Param bookkeeping_id path int true "Bookkeeping ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/bookkeeping-detail/by-bookkeeping/{bookkeeping_id} [get]
func (h *BookkeepingDetailHandler) GetByBookkeepingID(c *gin.Context) {
	// Parse Bookkeeping ID from URL parameter
	bookkeepingID, err := strconv.Atoi(c.Param("bookkeeping_id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid bookkeeping ID", nil)
		return
	}

	var details []models.BookkeepingDetail

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query details by bookkeeping ID with relationships
	if err := tenantDB.Preload("Type").
		Preload("Category").
		Preload("PaymentMethod").
		Where("bookkeeping_id = ?", bookkeepingID).
		Find(&details).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve bookkeeping details", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Bookkeeping details retrieved successfully", details)
}

// Create creates a new bookkeeping detail
// @Summary Create a new bookkeeping detail
// @Description Create a new bookkeeping detail
// @Tags BookkeepingDetail
// @Accept json
// @Produce json
// @Param request body BookkeepingDetailRequest true "Bookkeeping detail data"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/bookkeeping-detail [post]
func (h *BookkeepingDetailHandler) Create(c *gin.Context) {
	var req BookkeepingDetailRequest

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

	// Create bookkeeping detail
	detail := models.BookkeepingDetail{
		BookkeepingID:   req.BookkeepingID,
		TypeID:          req.TypeID,
		CategoryID:      req.CategoryID,
		PaymentMethodID: req.PaymentMethodID,
		PostedDate:      req.PostedDate,
		DocNumber:       req.DocNumber,
		Income:          req.Income,
		Expanse:         req.Expanse,
		Description:     req.Description,
		CreatedBy:       &userIDInt64,
	}

	// Create detail
	if err := tenantDB.Create(&detail).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create bookkeeping detail", err.Error())
		return
	}

	// Load relationships
	tenantDB.Preload("Bookkeeping").Preload("Type").Preload("Category").Preload("PaymentMethod").First(&detail, detail.ID)

	utils.SuccessResponse(c, http.StatusCreated, "Bookkeeping detail created successfully", detail)
}

// Update updates an existing bookkeeping detail
// @Summary Update bookkeeping detail
// @Description Update an existing bookkeeping detail by ID
// @Tags BookkeepingDetail
// @Accept json
// @Produce json
// @Param id path int true "Bookkeeping Detail ID"
// @Param request body BookkeepingDetailRequest true "Bookkeeping detail data"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/bookkeeping-detail/{id} [put]
func (h *BookkeepingDetailHandler) Update(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid bookkeeping detail ID", nil)
		return
	}

	var req BookkeepingDetailRequest

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
	var detail models.BookkeepingDetail
	if err := tenantDB.First(&detail, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Bookkeeping detail not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve bookkeeping detail", nil)
		return
	}

	// Update fields
	detail.BookkeepingID = req.BookkeepingID
	detail.TypeID = req.TypeID
	detail.CategoryID = req.CategoryID
	detail.PaymentMethodID = req.PaymentMethodID
	detail.PostedDate = req.PostedDate
	detail.DocNumber = req.DocNumber
	detail.Income = req.Income
	detail.Expanse = req.Expanse
	detail.Description = req.Description
	detail.UpdatedBy = &userIDInt64

	// Save updates
	if err := tenantDB.Save(&detail).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update bookkeeping detail", err.Error())
		return
	}

	// Load relationships
	tenantDB.Preload("Bookkeeping").Preload("Type").Preload("Category").Preload("PaymentMethod").First(&detail, detail.ID)

	utils.SuccessResponse(c, http.StatusOK, "Bookkeeping detail updated successfully", detail)
}

// Delete soft deletes a bookkeeping detail
// @Summary Delete bookkeeping detail
// @Description Soft delete a bookkeeping detail by ID
// @Tags BookkeepingDetail
// @Accept json
// @Produce json
// @Param id path int true "Bookkeeping Detail ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/bookkeeping-detail/{id} [delete]
func (h *BookkeepingDetailHandler) Delete(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid bookkeeping detail ID", nil)
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
	var detail models.BookkeepingDetail
	if err := tenantDB.First(&detail, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Bookkeeping detail not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve bookkeeping detail", nil)
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
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete bookkeeping detail", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Bookkeeping detail deleted successfully", nil)
}
