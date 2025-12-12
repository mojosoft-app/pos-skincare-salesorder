package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pos-mojosoft-so-service/internal/models"
	"pos-mojosoft-so-service/internal/utils"
)

type SummaryByTransactionTypeHandler struct {
	db *gorm.DB
}

func NewSummaryByTransactionTypeHandler(db *gorm.DB) *SummaryByTransactionTypeHandler {
	return &SummaryByTransactionTypeHandler{db: db}
}

// SummaryByTransactionTypeRequest represents the request body for creating/updating a summary
type SummaryByTransactionTypeRequest struct {
	BookkeepingID *int     `json:"bookkeeping_id"`
	TypeID        *int     `json:"type_id"`
	Total         *float64 `json:"total"`
}

// GetAll retrieves all summaries with optional filters
// @Summary Get all summaries by transaction type
// @Description Get list of all summaries by transaction type with optional filters
// @Tags SummaryByTransactionType
// @Accept json
// @Produce json
// @Param bookkeeping_id query int false "Filter by bookkeeping ID"
// @Param type_id query int false "Filter by transaction type ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/summary-by-transaction-type [get]
func (h *SummaryByTransactionTypeHandler) GetAll(c *gin.Context) {
	var summaries []models.SummaryByTransactionType

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Build query with preload relationships
	query := tenantDB.Model(&models.SummaryByTransactionType{}).
		Preload("Bookkeeping").
		Preload("Type")

	// Apply filters
	if bookkeepingID := c.Query("bookkeeping_id"); bookkeepingID != "" {
		query = query.Where("bookkeeping_id = ?", bookkeepingID)
	}
	if typeID := c.Query("type_id"); typeID != "" {
		query = query.Where("type_id = ?", typeID)
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
// @Description Get a single summary by transaction type by its ID
// @Tags SummaryByTransactionType
// @Accept json
// @Produce json
// @Param id path int true "Summary ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/summary-by-transaction-type/{id} [get]
func (h *SummaryByTransactionTypeHandler) GetByID(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid summary ID", nil)
		return
	}

	var summary models.SummaryByTransactionType

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query summary by ID with relationships
	if err := tenantDB.Preload("Bookkeeping").
		Preload("Type").
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
// @Tags SummaryByTransactionType
// @Accept json
// @Produce json
// @Param bookkeeping_id path int true "Bookkeeping ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/summary-by-transaction-type/by-bookkeeping/{bookkeeping_id} [get]
func (h *SummaryByTransactionTypeHandler) GetByBookkeepingID(c *gin.Context) {
	// Parse Bookkeeping ID from URL parameter
	bookkeepingID, err := strconv.Atoi(c.Param("bookkeeping_id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid bookkeeping ID", nil)
		return
	}

	var summaries []models.SummaryByTransactionType

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query summaries by bookkeeping ID with relationships
	if err := tenantDB.Preload("Type").
		Where("bookkeeping_id = ?", bookkeepingID).
		Find(&summaries).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve summaries", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Summaries retrieved successfully", summaries)
}

// Create creates a new summary
// @Summary Create a new summary by transaction type
// @Description Create a new summary by transaction type
// @Tags SummaryByTransactionType
// @Accept json
// @Produce json
// @Param request body SummaryByTransactionTypeRequest true "Summary data"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/summary-by-transaction-type [post]
func (h *SummaryByTransactionTypeHandler) Create(c *gin.Context) {
	var req SummaryByTransactionTypeRequest

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
	summary := models.SummaryByTransactionType{
		BookkeepingID: req.BookkeepingID,
		TypeID:        req.TypeID,
		Total:         req.Total,
		CreatedBy:     &userIDInt64,
	}

	// Create summary
	if err := tenantDB.Create(&summary).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create summary", err.Error())
		return
	}

	// Load relationships
	tenantDB.Preload("Bookkeeping").Preload("Type").First(&summary, summary.ID)

	utils.SuccessResponse(c, http.StatusCreated, "Summary created successfully", summary)
}

// Update updates an existing summary
// @Summary Update summary by transaction type
// @Description Update an existing summary by transaction type by ID
// @Tags SummaryByTransactionType
// @Accept json
// @Produce json
// @Param id path int true "Summary ID"
// @Param request body SummaryByTransactionTypeRequest true "Summary data"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/summary-by-transaction-type/{id} [put]
func (h *SummaryByTransactionTypeHandler) Update(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid summary ID", nil)
		return
	}

	var req SummaryByTransactionTypeRequest

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
	var summary models.SummaryByTransactionType
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
	summary.TypeID = req.TypeID
	summary.Total = req.Total
	summary.UpdatedBy = &userIDInt64

	// Save updates
	if err := tenantDB.Save(&summary).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update summary", err.Error())
		return
	}

	// Load relationships
	tenantDB.Preload("Bookkeeping").Preload("Type").First(&summary, summary.ID)

	utils.SuccessResponse(c, http.StatusOK, "Summary updated successfully", summary)
}

// Delete soft deletes a summary
// @Summary Delete summary by transaction type
// @Description Soft delete a summary by transaction type by ID
// @Tags SummaryByTransactionType
// @Accept json
// @Produce json
// @Param id path int true "Summary ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/summary-by-transaction-type/{id} [delete]
func (h *SummaryByTransactionTypeHandler) Delete(c *gin.Context) {
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
	var summary models.SummaryByTransactionType
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
