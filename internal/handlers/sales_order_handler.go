package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"pos-mojosoft-so-service/internal/models"
	"pos-mojosoft-so-service/internal/utils"
)

type SalesOrderHandler struct {
	db *gorm.DB
}

func NewSalesOrderHandler(db *gorm.DB) *SalesOrderHandler {
	return &SalesOrderHandler{db: db}
}

// CreateSalesOrderRequest represents the request body for creating a sales order
type CreateSalesOrderRequest struct {
	LocationID      *int                        `json:"location_id"`
	CustomerID      *int                        `json:"customer_id" binding:"required"`
	DocDate         *string                     `json:"doc_date"`
	InvNumber       *string                     `json:"inv_number"`
	Address         *string                     `json:"address"`
	DeliveryCost    *float64                    `json:"delivery_cost"`
	TotalAmount     *float64                    `json:"total_amount"`
	TotalPayment    *float64                    `json:"total_payment"`
	Outstanding     *float64                    `json:"outstanding"`
	TotalVoucher    *float64                    `json:"total_voucher"`
	VoucherNumber   *string                     `json:"voucher_number"`
	PostedDate      *string                     `json:"posted_date"`
	AdditionalCost  *float64                    `json:"additional_cost"`
	PreviousPayment *float64                    `json:"previous_payment"`
	FullyPaid       *bool                       `json:"fully_paid"`
	Note            *string                     `json:"note"`
	StatusID        *int                        `json:"status_id"`
	Details         []CreateSalesOrderDetailRequest `json:"details"`
	Services        []CreateSalesOrderServiceRequest `json:"services"`
}

type CreateSalesOrderDetailRequest struct {
	ItemID       *int     `json:"item_id"`
	UnitID       *int     `json:"unit_id"`
	PromoterID   *int     `json:"promoter_id"`
	ItemName     *string  `json:"item_name"`
	Quantity     *int     `json:"quantity"`
	Price        *float64 `json:"price"`
	ItemTotal    *float64 `json:"item_total"`
	DiscountPct  *int     `json:"discount_pct"`
	UsedSessions *int     `json:"used_sessions"`
}

type CreateSalesOrderServiceRequest struct {
	SalesOrderDetailID *int    `json:"sales_order_detail_id"`
	ServiceID          *int    `json:"service_id"`
	TreatmentID        *int    `json:"treatment_id"`
	MessageLogDetailID *string `json:"message_log_detail_id"`
	RemindedID         *int    `json:"reminded_id"`
	ServiceName        *string `json:"service_name"`
	Treated            *bool   `json:"treated"`
	Schedule           *string `json:"schedule"`
}

// GetAll retrieves all sales orders with optional filters
// @Summary Get all sales orders
// @Description Get list of all sales orders with optional pagination and filters
// @Tags SalesOrder
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param status_id query int false "Filter by status ID"
// @Param customer_id query int false "Filter by customer ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/sales-orders [get]
func (h *SalesOrderHandler) GetAll(c *gin.Context) {
	var salesOrders []models.SalesOrder

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Build query
	query := tenantDB.Model(&models.SalesOrder{})

	// Apply filters
	if statusID := c.Query("status_id"); statusID != "" {
		query = query.Where("status_id = ?", statusID)
	}
	if customerID := c.Query("customer_id"); customerID != "" {
		query = query.Where("costumer_id = ?", customerID)
	}

	// Preload relationships
	query = query.Preload("Status").Preload("Details").Preload("Services")

	// Execute query
	if err := query.Find(&salesOrders).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve sales orders", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Sales orders retrieved successfully", salesOrders)
}

// GetByID retrieves a single sales order by ID
// @Summary Get sales order by ID
// @Description Get a single sales order by its ID with details and services
// @Tags SalesOrder
// @Accept json
// @Produce json
// @Param id path string true "Sales Order ID (UUID)"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/sales-orders/{id} [get]
func (h *SalesOrderHandler) GetByID(c *gin.Context) {
	// Parse UUID from URL parameter
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid sales order ID", nil)
		return
	}

	var salesOrder models.SalesOrder

	// Get tenant DB from context
	db, exists := c.Get("tenantDB")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Database connection not found", nil)
		return
	}
	tenantDB := db.(*gorm.DB)

	// Query sales order by ID with relationships
	if err := tenantDB.Preload("Status").Preload("Details").Preload("Services").
		First(&salesOrder, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Sales order not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve sales order", nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Sales order retrieved successfully", salesOrder)
}

// Create creates a new sales order
// @Summary Create a new sales order
// @Description Create a new sales order with details and services
// @Tags SalesOrder
// @Accept json
// @Produce json
// @Param request body CreateSalesOrderRequest true "Sales Order data"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/sales-orders [post]
func (h *SalesOrderHandler) Create(c *gin.Context) {
	var req CreateSalesOrderRequest

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

	// Create sales order
	salesOrder := models.SalesOrder{
		LocationID:      req.LocationID,
		CustomerID:      req.CustomerID,
		InvNumber:       req.InvNumber,
		Address:         req.Address,
		DeliveryCost:    req.DeliveryCost,
		TotalAmount:     req.TotalAmount,
		TotalPayment:    req.TotalPayment,
		Outstanding:     req.Outstanding,
		TotalVoucher:    req.TotalVoucher,
		VoucherNumber:   req.VoucherNumber,
		AdditionalCost:  req.AdditionalCost,
		PreviousPayment: req.PreviousPayment,
		FullyPaid:       req.FullyPaid,
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

	// Create sales order
	if err := tx.Create(&salesOrder).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create sales order", err.Error())
		return
	}

	// Create details if provided
	if len(req.Details) > 0 {
		for _, detailReq := range req.Details {
			detail := models.SalesOrderDetail{
				ItemID:       detailReq.ItemID,
				UnitID:       detailReq.UnitID,
				PromoterID:   detailReq.PromoterID,
				ItemName:     detailReq.ItemName,
				Quantity:     detailReq.Quantity,
				Price:        detailReq.Price,
				ItemTotal:    detailReq.ItemTotal,
				DiscountPct:  detailReq.DiscountPct,
				UsedSessions: detailReq.UsedSessions,
				CreatedBy:    &userIDInt64,
			}
			if err := tx.Create(&detail).Error; err != nil {
				tx.Rollback()
				utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create sales order detail", err.Error())
				return
			}
		}
	}

	// Create services if provided
	if len(req.Services) > 0 {
		for _, serviceReq := range req.Services {
			service := models.SalesOrderService{
				SalesOrderDetailID: serviceReq.SalesOrderDetailID,
				ServiceID:          serviceReq.ServiceID,
				TreatmentID:        serviceReq.TreatmentID,
				MessageLogDetailID: serviceReq.MessageLogDetailID,
				RemindedID:         serviceReq.RemindedID,
				ServiceName:        serviceReq.ServiceName,
				Treated:            serviceReq.Treated,
				CreatedBy:          &userIDInt64,
			}
			if err := tx.Create(&service).Error; err != nil {
				tx.Rollback()
				utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create sales order service", err.Error())
				return
			}
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to commit transaction", err.Error())
		return
	}

	// Load the created sales order with relationships
	tenantDB.Preload("Status").Preload("Details").Preload("Services").
		First(&salesOrder, "id = ?", salesOrder.ID)

	utils.SuccessResponse(c, http.StatusCreated, "Sales order created successfully", salesOrder)
}

// Update updates an existing sales order
// @Summary Update sales order
// @Description Update an existing sales order by ID
// @Tags SalesOrder
// @Accept json
// @Produce json
// @Param id path string true "Sales Order ID (UUID)"
// @Param request body CreateSalesOrderRequest true "Sales Order data"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/sales-orders/{id} [put]
func (h *SalesOrderHandler) Update(c *gin.Context) {
	// Parse UUID from URL parameter
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid sales order ID", nil)
		return
	}

	var req CreateSalesOrderRequest

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

	// Check if sales order exists
	var salesOrder models.SalesOrder
	if err := tenantDB.First(&salesOrder, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Sales order not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve sales order", nil)
		return
	}

	// Update fields
	salesOrder.LocationID = req.LocationID
	salesOrder.CustomerID = req.CustomerID
	salesOrder.InvNumber = req.InvNumber
	salesOrder.Address = req.Address
	salesOrder.DeliveryCost = req.DeliveryCost
	salesOrder.TotalAmount = req.TotalAmount
	salesOrder.TotalPayment = req.TotalPayment
	salesOrder.Outstanding = req.Outstanding
	salesOrder.TotalVoucher = req.TotalVoucher
	salesOrder.VoucherNumber = req.VoucherNumber
	salesOrder.AdditionalCost = req.AdditionalCost
	salesOrder.PreviousPayment = req.PreviousPayment
	salesOrder.FullyPaid = req.FullyPaid
	salesOrder.Note = req.Note
	salesOrder.StatusID = req.StatusID
	salesOrder.UpdatedBy = &userIDInt64

	// Save updates
	if err := tenantDB.Save(&salesOrder).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update sales order", err.Error())
		return
	}

	// Load updated sales order with relationships
	tenantDB.Preload("Status").Preload("Details").Preload("Services").
		First(&salesOrder, "id = ?", salesOrder.ID)

	utils.SuccessResponse(c, http.StatusOK, "Sales order updated successfully", salesOrder)
}

// Delete soft deletes a sales order
// @Summary Delete sales order
// @Description Soft delete a sales order by ID
// @Tags SalesOrder
// @Accept json
// @Produce json
// @Param id path string true "Sales Order ID (UUID)"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /so/api/sales-orders/{id} [delete]
func (h *SalesOrderHandler) Delete(c *gin.Context) {
	// Parse UUID from URL parameter
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid sales order ID", nil)
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

	// Check if sales order exists
	var salesOrder models.SalesOrder
	if err := tenantDB.First(&salesOrder, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Sales order not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve sales order", nil)
		return
	}

	// Set deleted_by before soft delete
	salesOrder.DeletedBy = &userIDInt64
	if err := tenantDB.Save(&salesOrder).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update deleted_by", err.Error())
		return
	}

	// Soft delete
	if err := tenantDB.Delete(&salesOrder).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete sales order", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Sales order deleted successfully", nil)
}
