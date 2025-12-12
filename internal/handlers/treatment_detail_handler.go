package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pos-mojosoft-so-service/internal/models"
	"pos-mojosoft-so-service/internal/utils"
)

type TreatmentDetailHandler struct {
	db *gorm.DB
}

func NewTreatmentDetailHandler(db *gorm.DB) *TreatmentDetailHandler {
	return &TreatmentDetailHandler{db: db}
}

// TreatmentDetailRequest represents the request body for creating/updating a treatment detail
type TreatmentDetailRequest struct {
	TreatmentID *int `json:"treatment_id"`
	ItemID      *int `json:"item_id"`
	UnitID      *int `json:"unit_id"`
	Quantity    *int `json:"quantity" binding:"required"`
}

// GetAll retrieves all treatment details with optional filters
// @Summary Get all treatment details
// @Description Get list of all treatment details with optional filters
// @Tags TreatmentDetail
// @Accept json
// @Produce json
// @Param treatment_id query int false "Filter by treatment ID"
// @Param item_id query int false "Filter by item ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/treatment-details [get]
func (h *TreatmentDetailHandler) GetAll(c *gin.Context) {
	var details []models.TreatmentDetail

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Build query
	query := tenantDB.Model(&models.TreatmentDetail{})

	// Apply filters
	if treatmentID := c.Query("treatment_id"); treatmentID != "" {
		query = query.Where("treatment_id = ?", treatmentID)
	}
	if itemID := c.Query("item_id"); itemID != "" {
		query = query.Where("item_id = ?", itemID)
	}

	// Execute query
	if err := query.Find(&details).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve treatment details", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Treatment details retrieved successfully", details)
}

// GetByID retrieves a single treatment detail by ID
// @Summary Get treatment detail by ID
// @Description Get a single treatment detail by its ID
// @Tags TreatmentDetail
// @Accept json
// @Produce json
// @Param id path int true "Detail ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/treatment-details/{id} [get]
func (h *TreatmentDetailHandler) GetByID(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid detail ID", nil)
		return
	}

	var detail models.TreatmentDetail

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
			utils.ErrorResponse(c, http.StatusNotFound, "Treatment detail not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve treatment detail", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Treatment detail retrieved successfully", detail)
}

// GetByTreatmentID retrieves all details for a specific treatment
// @Summary Get treatment details by treatment ID
// @Description Get all details for a specific treatment
// @Tags TreatmentDetail
// @Accept json
// @Produce json
// @Param treatment_id path int true "Treatment ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/treatment-details/by-treatment/{treatment_id} [get]
func (h *TreatmentDetailHandler) GetByTreatmentID(c *gin.Context) {
	// Parse Treatment ID from URL parameter
	treatmentID, err := strconv.Atoi(c.Param("treatment_id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid treatment ID", nil)
		return
	}

	var details []models.TreatmentDetail

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query details by treatment ID
	if err := tenantDB.Where("treatment_id = ?", treatmentID).Find(&details).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve treatment details", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Treatment details retrieved successfully", details)
}

// Create creates a new treatment detail
// @Summary Create a new treatment detail
// @Description Create a new treatment detail
// @Tags TreatmentDetail
// @Accept json
// @Produce json
// @Param request body TreatmentDetailRequest true "Treatment Detail data"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/treatment-details [post]
func (h *TreatmentDetailHandler) Create(c *gin.Context) {
	var req TreatmentDetailRequest

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

	// Create treatment detail
	detail := models.TreatmentDetail{
		TreatmentID: req.TreatmentID,
		ItemID:      req.ItemID,
		UnitID:      req.UnitID,
		Quantity:    req.Quantity,
		CreatedBy:   &userIDInt64,
	}

	// Create detail
	if err := tenantDB.Create(&detail).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create treatment detail", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Treatment detail created successfully", detail)
}

// Update updates an existing treatment detail
// @Summary Update treatment detail
// @Description Update an existing treatment detail by ID
// @Tags TreatmentDetail
// @Accept json
// @Produce json
// @Param id path int true "Detail ID"
// @Param request body TreatmentDetailRequest true "Treatment Detail data"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/treatment-details/{id} [put]
func (h *TreatmentDetailHandler) Update(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid detail ID", nil)
		return
	}

	var req TreatmentDetailRequest

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
	var detail models.TreatmentDetail
	if err := tenantDB.First(&detail, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Treatment detail not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve treatment detail", nil)
		return
	}

	// Update fields
	detail.TreatmentID = req.TreatmentID
	detail.ItemID = req.ItemID
	detail.UnitID = req.UnitID
	detail.Quantity = req.Quantity
	detail.UpdatedBy = &userIDInt64

	// Save updates
	if err := tenantDB.Save(&detail).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update treatment detail", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Treatment detail updated successfully", detail)
}

// Delete soft deletes a treatment detail
// @Summary Delete treatment detail
// @Description Soft delete a treatment detail by ID
// @Tags TreatmentDetail
// @Accept json
// @Produce json
// @Param id path int true "Detail ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/treatment-details/{id} [delete]
func (h *TreatmentDetailHandler) Delete(c *gin.Context) {
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
	var detail models.TreatmentDetail
	if err := tenantDB.First(&detail, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Treatment detail not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve treatment detail", nil)
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
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete treatment detail", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Treatment detail deleted successfully", nil)
}
