package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pos-mojosoft-so-service/internal/models"
	"pos-mojosoft-so-service/internal/utils"
)

type BookTransactionTypeHandler struct {
	db *gorm.DB
}

func NewBookTransactionTypeHandler(db *gorm.DB) *BookTransactionTypeHandler {
	return &BookTransactionTypeHandler{db: db}
}

// BookTransactionTypeRequest represents the request body for creating/updating a book transaction type
type BookTransactionTypeRequest struct {
	Name *string `json:"name" binding:"required"`
}

// GetAll retrieves all book transaction types
// @Summary Get all book transaction types
// @Description Get list of all book transaction types with optional filters
// @Tags BookTransactionType
// @Accept json
// @Produce json
// @Param name query string false "Filter by name (partial match)"
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/book-transaction-type [get]
func (h *BookTransactionTypeHandler) GetAll(c *gin.Context) {
	var types []models.BookTransactionType

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Build query
	query := tenantDB.Model(&models.BookTransactionType{})

	// Apply filters
	if name := c.Query("name"); name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	// Execute query
	if err := query.Find(&types).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve book transaction types", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Book transaction types retrieved successfully", types)
}

// GetByID retrieves a single book transaction type by ID
// @Summary Get book transaction type by ID
// @Description Get a single book transaction type by its ID
// @Tags BookTransactionType
// @Accept json
// @Produce json
// @Param id path int true "Book Transaction Type ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/book-transaction-type/{id} [get]
func (h *BookTransactionTypeHandler) GetByID(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid book transaction type ID", nil)
		return
	}

	var transactionType models.BookTransactionType

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query book transaction type by ID
	if err := tenantDB.First(&transactionType, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Book transaction type not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve book transaction type", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Book transaction type retrieved successfully", transactionType)
}

// Create creates a new book transaction type
// @Summary Create a new book transaction type
// @Description Create a new book transaction type
// @Tags BookTransactionType
// @Accept json
// @Produce json
// @Param request body BookTransactionTypeRequest true "Book transaction type data"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/book-transaction-type [post]
func (h *BookTransactionTypeHandler) Create(c *gin.Context) {
	var req BookTransactionTypeRequest

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

	// Create book transaction type
	transactionType := models.BookTransactionType{
		Name:      req.Name,
		CreatedBy: &userIDInt64,
	}

	// Create transaction type
	if err := tenantDB.Create(&transactionType).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create book transaction type", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Book transaction type created successfully", transactionType)
}

// Update updates an existing book transaction type
// @Summary Update book transaction type
// @Description Update an existing book transaction type by ID
// @Tags BookTransactionType
// @Accept json
// @Produce json
// @Param id path int true "Book Transaction Type ID"
// @Param request body BookTransactionTypeRequest true "Book transaction type data"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/book-transaction-type/{id} [put]
func (h *BookTransactionTypeHandler) Update(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid book transaction type ID", nil)
		return
	}

	var req BookTransactionTypeRequest

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

	// Check if transaction type exists
	var transactionType models.BookTransactionType
	if err := tenantDB.First(&transactionType, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Book transaction type not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve book transaction type", nil)
		return
	}

	// Update fields
	transactionType.Name = req.Name
	transactionType.UpdatedBy = &userIDInt64

	// Save updates
	if err := tenantDB.Save(&transactionType).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update book transaction type", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Book transaction type updated successfully", transactionType)
}

// Delete soft deletes a book transaction type
// @Summary Delete book transaction type
// @Description Soft delete a book transaction type by ID
// @Tags BookTransactionType
// @Accept json
// @Produce json
// @Param id path int true "Book Transaction Type ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/book-transaction-type/{id} [delete]
func (h *BookTransactionTypeHandler) Delete(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid book transaction type ID", nil)
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

	// Check if transaction type exists
	var transactionType models.BookTransactionType
	if err := tenantDB.First(&transactionType, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Book transaction type not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve book transaction type", nil)
		return
	}

	// Set deleted_by before soft delete
	transactionType.DeletedBy = &userIDInt64
	if err := tenantDB.Save(&transactionType).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update deleted_by", err.Error())
		return
	}

	// Soft delete
	if err := tenantDB.Delete(&transactionType).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete book transaction type", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Book transaction type deleted successfully", nil)
}
