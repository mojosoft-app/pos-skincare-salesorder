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

type BookkeepingHandler struct {
	db *gorm.DB
}

func NewBookkeepingHandler(db *gorm.DB) *BookkeepingHandler {
	return &BookkeepingHandler{db: db}
}

// BookkeepingRequest represents the request body for creating/updating a bookkeeping record
type BookkeepingRequest struct {
	LocationID *string    `json:"location_id"`
	BookDate   *time.Time `json:"book_date"`
	Opening    *float64   `json:"opening"`
	Income     *float64   `json:"income"`
	Expanse    *float64   `json:"expanse"`
	Balance    *float64   `json:"balance"`
	Note       *string    `json:"note"`
	StatusID   *int       `json:"status_id"`
}

// GetAll retrieves all bookkeeping records with optional filters
// @Summary Get all bookkeeping records
// @Description Get list of all bookkeeping records with optional filters
// @Tags Bookkeeping
// @Accept json
// @Produce json
// @Param location_id query string false "Filter by location ID"
// @Param status_id query int false "Filter by status ID"
// @Param book_date_from query string false "Filter by book date from (YYYY-MM-DD)"
// @Param book_date_to query string false "Filter by book date to (YYYY-MM-DD)"
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/bookkeeping [get]
func (h *BookkeepingHandler) GetAll(c *gin.Context) {
	var bookkeepings []models.Bookkeeping

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Build query with preload relationships
	query := tenantDB.Model(&models.Bookkeeping{}).
		Preload("Status").
		Preload("Details")

	// Apply filters
	if locationID := c.Query("location_id"); locationID != "" {
		query = query.Where("location_id = ?", locationID)
	}
	if statusID := c.Query("status_id"); statusID != "" {
		query = query.Where("status_id = ?", statusID)
	}
	if bookDateFrom := c.Query("book_date_from"); bookDateFrom != "" {
		query = query.Where("bookdate >= ?", bookDateFrom)
	}
	if bookDateTo := c.Query("book_date_to"); bookDateTo != "" {
		query = query.Where("bookdate <= ?", bookDateTo)
	}

	// Execute query
	if err := query.Find(&bookkeepings).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve bookkeeping records", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Bookkeeping records retrieved successfully", bookkeepings)
}

// GetByID retrieves a single bookkeeping record by ID
// @Summary Get bookkeeping by ID
// @Description Get a single bookkeeping record by its ID with all relationships
// @Tags Bookkeeping
// @Accept json
// @Produce json
// @Param id path int true "Bookkeeping ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/bookkeeping/{id} [get]
func (h *BookkeepingHandler) GetByID(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid bookkeeping ID", nil)
		return
	}

	var bookkeeping models.Bookkeeping

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query bookkeeping by ID with relationships
	if err := tenantDB.Preload("Status").
		Preload("Details").
		Preload("SummaryByTransactionType").
		Preload("SummaryByPaymentMethod").
		Preload("SummaryByTransactionTypeAndPaymentMethod").
		First(&bookkeeping, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Bookkeeping record not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve bookkeeping record", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Bookkeeping record retrieved successfully", bookkeeping)
}

// GetByLocationID retrieves all bookkeeping records for a specific location
// @Summary Get bookkeeping by location ID
// @Description Get all bookkeeping records for a specific location
// @Tags Bookkeeping
// @Accept json
// @Produce json
// @Param location_id path string true "Location ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/bookkeeping/by-location/{location_id} [get]
func (h *BookkeepingHandler) GetByLocationID(c *gin.Context) {
	locationID := c.Param("location_id")

	var bookkeepings []models.Bookkeeping

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query bookkeeping by location ID with relationships
	if err := tenantDB.Preload("Status").
		Preload("Details").
		Where("location_id = ?", locationID).
		Find(&bookkeepings).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve bookkeeping records", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Bookkeeping records retrieved successfully", bookkeepings)
}

// Create creates a new bookkeeping record
// @Summary Create a new bookkeeping record
// @Description Create a new bookkeeping record
// @Tags Bookkeeping
// @Accept json
// @Produce json
// @Param request body BookkeepingRequest true "Bookkeeping data"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/bookkeeping [post]
func (h *BookkeepingHandler) Create(c *gin.Context) {
	var req BookkeepingRequest

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

	// Create bookkeeping record
	bookkeeping := models.Bookkeeping{
		LocationID: req.LocationID,
		BookDate:   req.BookDate,
		Opening:    req.Opening,
		Income:     req.Income,
		Expanse:    req.Expanse,
		Balance:    req.Balance,
		Note:       req.Note,
		StatusID:   req.StatusID,
		CreatedBy:  &userIDInt64,
	}

	// Create bookkeeping
	if err := tenantDB.Create(&bookkeeping).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create bookkeeping record", err.Error())
		return
	}

	// Load relationships
	tenantDB.Preload("Status").First(&bookkeeping, bookkeeping.ID)

	utils.SuccessResponse(c, http.StatusCreated, "Bookkeeping record created successfully", bookkeeping)
}

// Update updates an existing bookkeeping record
// @Summary Update bookkeeping record
// @Description Update an existing bookkeeping record by ID
// @Tags Bookkeeping
// @Accept json
// @Produce json
// @Param id path int true "Bookkeeping ID"
// @Param request body BookkeepingRequest true "Bookkeeping data"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/bookkeeping/{id} [put]
func (h *BookkeepingHandler) Update(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid bookkeeping ID", nil)
		return
	}

	var req BookkeepingRequest

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

	// Check if bookkeeping exists
	var bookkeeping models.Bookkeeping
	if err := tenantDB.First(&bookkeeping, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Bookkeeping record not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve bookkeeping record", nil)
		return
	}

	// Update fields
	bookkeeping.LocationID = req.LocationID
	bookkeeping.BookDate = req.BookDate
	bookkeeping.Opening = req.Opening
	bookkeeping.Income = req.Income
	bookkeeping.Expanse = req.Expanse
	bookkeeping.Balance = req.Balance
	bookkeeping.Note = req.Note
	bookkeeping.StatusID = req.StatusID
	bookkeeping.UpdatedBy = &userIDInt64

	// Save updates
	if err := tenantDB.Save(&bookkeeping).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update bookkeeping record", err.Error())
		return
	}

	// Load relationships
	tenantDB.Preload("Status").First(&bookkeeping, bookkeeping.ID)

	utils.SuccessResponse(c, http.StatusOK, "Bookkeeping record updated successfully", bookkeeping)
}

// Delete soft deletes a bookkeeping record
// @Summary Delete bookkeeping record
// @Description Soft delete a bookkeeping record by ID
// @Tags Bookkeeping
// @Accept json
// @Produce json
// @Param id path int true "Bookkeeping ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/bookkeeping/{id} [delete]
func (h *BookkeepingHandler) Delete(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid bookkeeping ID", nil)
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

	// Check if bookkeeping exists
	var bookkeeping models.Bookkeeping
	if err := tenantDB.First(&bookkeeping, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Bookkeeping record not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve bookkeeping record", nil)
		return
	}

	// Set deleted_by before soft delete
	bookkeeping.DeletedBy = &userIDInt64
	if err := tenantDB.Save(&bookkeeping).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update deleted_by", err.Error())
		return
	}

	// Soft delete
	if err := tenantDB.Delete(&bookkeeping).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete bookkeeping record", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Bookkeeping record deleted successfully", nil)
}
