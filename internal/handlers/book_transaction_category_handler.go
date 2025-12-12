package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pos-mojosoft-so-service/internal/models"
	"pos-mojosoft-so-service/internal/utils"
)

type BookTransactionCategoryHandler struct {
	db *gorm.DB
}

func NewBookTransactionCategoryHandler(db *gorm.DB) *BookTransactionCategoryHandler {
	return &BookTransactionCategoryHandler{db: db}
}

// BookTransactionCategoryRequest represents the request body for creating/updating a book transaction category
type BookTransactionCategoryRequest struct {
	Name *string `json:"name" binding:"required"`
}

// GetAll retrieves all book transaction categories
// @Summary Get all book transaction categories
// @Description Get list of all book transaction categories with optional filters
// @Tags BookTransactionCategory
// @Accept json
// @Produce json
// @Param name query string false "Filter by name (partial match)"
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/book-transaction-category [get]
func (h *BookTransactionCategoryHandler) GetAll(c *gin.Context) {
	var categories []models.BookTransactionCategory

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Build query
	query := tenantDB.Model(&models.BookTransactionCategory{})

	// Apply filters
	if name := c.Query("name"); name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	// Execute query
	if err := query.Find(&categories).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve book transaction categories", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Book transaction categories retrieved successfully", categories)
}

// GetByID retrieves a single book transaction category by ID
// @Summary Get book transaction category by ID
// @Description Get a single book transaction category by its ID
// @Tags BookTransactionCategory
// @Accept json
// @Produce json
// @Param id path int true "Book Transaction Category ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/book-transaction-category/{id} [get]
func (h *BookTransactionCategoryHandler) GetByID(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid book transaction category ID", nil)
		return
	}

	var category models.BookTransactionCategory

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query book transaction category by ID
	if err := tenantDB.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Book transaction category not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve book transaction category", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Book transaction category retrieved successfully", category)
}

// Create creates a new book transaction category
// @Summary Create a new book transaction category
// @Description Create a new book transaction category
// @Tags BookTransactionCategory
// @Accept json
// @Produce json
// @Param request body BookTransactionCategoryRequest true "Book transaction category data"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/book-transaction-category [post]
func (h *BookTransactionCategoryHandler) Create(c *gin.Context) {
	var req BookTransactionCategoryRequest

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

	// Create book transaction category
	category := models.BookTransactionCategory{
		Name:      req.Name,
		CreatedBy: &userIDInt64,
	}

	// Create category
	if err := tenantDB.Create(&category).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create book transaction category", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Book transaction category created successfully", category)
}

// Update updates an existing book transaction category
// @Summary Update book transaction category
// @Description Update an existing book transaction category by ID
// @Tags BookTransactionCategory
// @Accept json
// @Produce json
// @Param id path int true "Book Transaction Category ID"
// @Param request body BookTransactionCategoryRequest true "Book transaction category data"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/book-transaction-category/{id} [put]
func (h *BookTransactionCategoryHandler) Update(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid book transaction category ID", nil)
		return
	}

	var req BookTransactionCategoryRequest

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

	// Check if category exists
	var category models.BookTransactionCategory
	if err := tenantDB.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Book transaction category not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve book transaction category", nil)
		return
	}

	// Update fields
	category.Name = req.Name
	category.UpdatedBy = &userIDInt64

	// Save updates
	if err := tenantDB.Save(&category).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update book transaction category", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Book transaction category updated successfully", category)
}

// Delete soft deletes a book transaction category
// @Summary Delete book transaction category
// @Description Soft delete a book transaction category by ID
// @Tags BookTransactionCategory
// @Accept json
// @Produce json
// @Param id path int true "Book Transaction Category ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/book-transaction-category/{id} [delete]
func (h *BookTransactionCategoryHandler) Delete(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid book transaction category ID", nil)
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

	// Check if category exists
	var category models.BookTransactionCategory
	if err := tenantDB.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Book transaction category not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve book transaction category", nil)
		return
	}

	// Set deleted_by before soft delete
	category.DeletedBy = &userIDInt64
	if err := tenantDB.Save(&category).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update deleted_by", err.Error())
		return
	}

	// Soft delete
	if err := tenantDB.Delete(&category).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete book transaction category", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Book transaction category deleted successfully", nil)
}
