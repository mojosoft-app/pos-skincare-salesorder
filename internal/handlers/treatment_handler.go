package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"pos-mojosoft-so-service/internal/models"
	"pos-mojosoft-so-service/internal/utils"
)

type TreatmentHandler struct {
	db *gorm.DB
}

func NewTreatmentHandler(db *gorm.DB) *TreatmentHandler {
	return &TreatmentHandler{db: db}
}

// CreateTreatmentRequest represents the request body for creating a treatment
type CreateTreatmentRequest struct {
	LocationID          *int                         `json:"location_id"`
	CustomerID          *int                         `json:"customer_id"`
	SalesOrderID        *int                         `json:"sales_order_id"`
	SalesOrderDetailID  *int                         `json:"sales_order_detail_id"`
	SalesOrderServiceID *int                         `json:"sales_order_service_id"`
	ServiceID           *int                         `json:"service_id"`
	PatientID           *int                         `json:"patient_id"`
	DoctorID            *int                         `json:"doctor_id"`
	NurseID             *int                         `json:"nurse_id"`
	BeauticianID        *int                         `json:"beautician_id"`
	DocNumber           *string                      `json:"doc_number"`
	DocDate             *string                      `json:"doc_date"`
	PostedDate          *string                      `json:"posted_date"`
	ServiceText         *string                      `json:"service_text"`
	Note                *string                      `json:"note"`
	StatusID            *int                         `json:"status_id"`
	Details             []CreateTreatmentDetailRequest `json:"details"`
}

type CreateTreatmentDetailRequest struct {
	ItemID   *int `json:"item_id"`
	UnitID   *int `json:"unit_id"`
	Quantity *int `json:"quantity"`
}

// GetAll retrieves all treatments with optional filters
// @Summary Get all treatments
// @Description Get list of all treatments with optional pagination and filters
// @Tags Treatment
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param status_id query int false "Filter by status ID"
// @Param patient_id query int false "Filter by patient ID"
// @Param doctor_id query int false "Filter by doctor ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/treatments [get]
func (h *TreatmentHandler) GetAll(c *gin.Context) {
	var treatments []models.Treatment

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Build query
	query := tenantDB.Model(&models.Treatment{})

	// Apply filters
	if statusID := c.Query("status_id"); statusID != "" {
		query = query.Where("status_id = ?", statusID)
	}
	if patientID := c.Query("patient_id"); patientID != "" {
		query = query.Where("patient_id = ?", patientID)
	}
	if doctorID := c.Query("doctor_id"); doctorID != "" {
		query = query.Where("doctor_id = ?", doctorID)
	}

	// Preload relationships
	query = query.Preload("Details")

	// Execute query
	if err := query.Find(&treatments).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve treatments", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Treatments retrieved successfully", treatments)
}

// GetByID retrieves a single treatment by ID
// @Summary Get treatment by ID
// @Description Get a single treatment by its ID with details
// @Tags Treatment
// @Accept json
// @Produce json
// @Param id path string true "Treatment ID (UUID)"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/treatments/{id} [get]
func (h *TreatmentHandler) GetByID(c *gin.Context) {
	// Parse UUID from URL parameter
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid treatment ID", nil)
		return
	}

	var treatment models.Treatment

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query treatment by ID with relationships
	if err := tenantDB.Preload("Details").
		First(&treatment, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Treatment not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve treatment", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Treatment retrieved successfully", treatment)
}

// Create creates a new treatment
// @Summary Create a new treatment
// @Description Create a new treatment with details
// @Tags Treatment
// @Accept json
// @Produce json
// @Param request body CreateTreatmentRequest true "Treatment data"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/treatments [post]
func (h *TreatmentHandler) Create(c *gin.Context) {
	var req CreateTreatmentRequest

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

	// Parse dates
	var docDate, postedDate *time.Time
	if req.DocDate != nil {
		parsedDate, err := time.Parse("2006-01-02", *req.DocDate)
		if err == nil {
			docDate = &parsedDate
		}
	}
	if req.PostedDate != nil {
		parsedDate, err := time.Parse("2006-01-02", *req.PostedDate)
		if err == nil {
			postedDate = &parsedDate
		}
	}

	// Create treatment
	treatment := models.Treatment{
		LocationID:          req.LocationID,
		CustomerID:          req.CustomerID,
		SalesOrderID:        req.SalesOrderID,
		SalesOrderDetailID:  req.SalesOrderDetailID,
		SalesOrderServiceID: req.SalesOrderServiceID,
		ServiceID:           req.ServiceID,
		PatientID:           req.PatientID,
		DoctorID:            req.DoctorID,
		NurseID:             req.NurseID,
		BeauticianID:        req.BeauticianID,
		DocNumber:           req.DocNumber,
		DocDate:             docDate,
		PostedDate:          postedDate,
		ServiceText:         req.ServiceText,
		Note:                req.Note,
		StatusID:            req.StatusID,
		CreatedBy:           &userIDInt64,
	}

	// Begin transaction
	tx := tenantDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create treatment
	if err := tx.Create(&treatment).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create treatment", err.Error())
		return
	}

	// Create details if provided
	if len(req.Details) > 0 {
		for _, detailReq := range req.Details {
			detail := models.TreatmentDetail{
				ItemID:    detailReq.ItemID,
				UnitID:    detailReq.UnitID,
				Quantity:  detailReq.Quantity,
				CreatedBy: &userIDInt64,
			}
			if err := tx.Create(&detail).Error; err != nil {
				tx.Rollback()
				utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create treatment detail", err.Error())
				return
			}
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to commit transaction", err.Error())
		return
	}

	// Load the created treatment with relationships
	tenantDB.Preload("Details").First(&treatment, "id = ?", treatment.ID)

	utils.SuccessResponse(c, http.StatusCreated, "Treatment created successfully", treatment)
}

// Update updates an existing treatment
// @Summary Update treatment
// @Description Update an existing treatment by ID
// @Tags Treatment
// @Accept json
// @Produce json
// @Param id path string true "Treatment ID (UUID)"
// @Param request body CreateTreatmentRequest true "Treatment data"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/treatments/{id} [put]
func (h *TreatmentHandler) Update(c *gin.Context) {
	// Parse UUID from URL parameter
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid treatment ID", nil)
		return
	}

	var req CreateTreatmentRequest

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

	// Check if treatment exists
	var treatment models.Treatment
	if err := tenantDB.First(&treatment, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Treatment not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve treatment", nil)
		return
	}

	// Parse dates
	var docDate, postedDate *time.Time
	if req.DocDate != nil {
		parsedDate, err := time.Parse("2006-01-02", *req.DocDate)
		if err == nil {
			docDate = &parsedDate
		}
	}
	if req.PostedDate != nil {
		parsedDate, err := time.Parse("2006-01-02", *req.PostedDate)
		if err == nil {
			postedDate = &parsedDate
		}
	}

	// Update fields
	treatment.LocationID = req.LocationID
	treatment.CustomerID = req.CustomerID
	treatment.SalesOrderID = req.SalesOrderID
	treatment.SalesOrderDetailID = req.SalesOrderDetailID
	treatment.SalesOrderServiceID = req.SalesOrderServiceID
	treatment.ServiceID = req.ServiceID
	treatment.PatientID = req.PatientID
	treatment.DoctorID = req.DoctorID
	treatment.NurseID = req.NurseID
	treatment.BeauticianID = req.BeauticianID
	treatment.DocNumber = req.DocNumber
	treatment.DocDate = docDate
	treatment.PostedDate = postedDate
	treatment.ServiceText = req.ServiceText
	treatment.Note = req.Note
	treatment.StatusID = req.StatusID
	treatment.UpdatedBy = &userIDInt64

	// Save updates
	if err := tenantDB.Save(&treatment).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update treatment", err.Error())
		return
	}

	// Load updated treatment with relationships
	tenantDB.Preload("Details").First(&treatment, "id = ?", treatment.ID)

	utils.SuccessResponse(c, http.StatusOK, "Treatment updated successfully", treatment)
}

// Delete soft deletes a treatment
// @Summary Delete treatment
// @Description Soft delete a treatment by ID
// @Tags Treatment
// @Accept json
// @Produce json
// @Param id path string true "Treatment ID (UUID)"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/treatments/{id} [delete]
func (h *TreatmentHandler) Delete(c *gin.Context) {
	// Parse UUID from URL parameter
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid treatment ID", nil)
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

	// Check if treatment exists
	var treatment models.Treatment
	if err := tenantDB.First(&treatment, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Treatment not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve treatment", nil)
		return
	}

	// Set deleted_by before soft delete
	treatment.DeletedBy = &userIDInt64
	if err := tenantDB.Save(&treatment).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update deleted_by", err.Error())
		return
	}

	// Soft delete
	if err := tenantDB.Delete(&treatment).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete treatment", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Treatment deleted successfully", nil)
}
