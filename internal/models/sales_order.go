package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SalesOrder represents the sales_order table in the database
type SalesOrder struct {
	ID                uuid.UUID         `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	LocationID        *int              `gorm:"column:location_id" json:"location_id"`
	CustomerID        *int              `gorm:"column:costumer_id" json:"customer_id"`
	DocNumber         *int              `gorm:"column:docnumber" json:"doc_number"`
	DocDate           *time.Time        `gorm:"column:docdate;type:date" json:"doc_date"`
	InvNumber         *string           `gorm:"column:invnumber" json:"inv_number"`
	Address           *string           `gorm:"column:address" json:"address"`
	DeliveryCost      *float64          `gorm:"column:deliverycost;type:numeric" json:"delivery_cost"`
	TotalAmount       *float64          `gorm:"column:totalamounth;type:numeric" json:"total_amount"`
	TotalPayment      *float64          `gorm:"column:totalpayment;type:numeric" json:"total_payment"`
	Outstanding       *float64          `gorm:"column:outstanding;type:numeric" json:"outstanding"`
	TotalVoucher      *float64          `gorm:"column:totalvoucher;type:numeric" json:"total_voucher"`
	VoucherNumber     *string           `gorm:"column:vouchernumbeer " json:"voucher_number"`
	PostedDate        *time.Time        `gorm:"column:posteddate;type:date" json:"posted_date"`
	Migrated          *bool             `gorm:"column:migrated" json:"migrated"`
	AdditionalCost    *float64          `gorm:"column:maddittionalcost;type:numeric" json:"additional_cost"`
	PreviousPayment   *float64          `gorm:"column:mpreviouspayment;type:numeric" json:"previous_payment"`
	FullyPaid         *bool             `gorm:"column:fullypaid" json:"fully_paid"`
	Note              *string           `gorm:"column:note" json:"note"`
	StatusID          *int              `gorm:"column:status_id" json:"status_id"`
	CreatedBy         *int64            `gorm:"column:created_by" json:"created_by"`
	UpdatedBy         *int64            `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy         *int64            `gorm:"column:deleted_by" json:"deleted_by"`
	DeletedAt         gorm.DeletedAt    `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	CreatedAt         *time.Time        `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         *time.Time        `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relationships
	Status            *SalesOrderStatus `gorm:"foreignKey:StatusID;references:ID" json:"status,omitempty"`
	Details           []SalesOrderDetail `gorm:"foreignKey:SalesOrderID;references:ID" json:"details,omitempty"`
	Services          []SalesOrderService `gorm:"foreignKey:SalesOrderID;references:ID" json:"services,omitempty"`
}

// TableName specifies the table name for SalesOrder model
func (SalesOrder) TableName() string {
	return "alana.sales_order"
}

// BeforeCreate hook to generate UUID before creating a new record
func (so *SalesOrder) BeforeCreate(tx *gorm.DB) error {
	if so.ID == uuid.Nil {
		so.ID = uuid.New()
	}
	return nil
}
