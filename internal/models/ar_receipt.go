package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ARReceipt represents the ar_receipt table in the database
type ARReceipt struct {
	ID              uuid.UUID         `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	LocationID      *int              `gorm:"column:lacation_id" json:"location_id"`
	CustomerID      *int              `gorm:"column:customer_id" json:"customer_id"`
	PaymentMethodID *int              `gorm:"column:paymentmethod_id" json:"payment_method_id"`
	DocNumber       *int              `gorm:"column:docnumber" json:"doc_number"`
	DocDate         *time.Time        `gorm:"column:docdate;type:date" json:"doc_date"`
	PostedDate      *time.Time        `gorm:"column:posteddate;type:date" json:"posted_date"`
	TotalAmount     *float64          `gorm:"column:totalamounth;type:numeric" json:"total_amount"`
	Note            *string           `gorm:"column:note " json:"note"`
	StatusID        *int              `gorm:"column:status_id" json:"status_id"`
	CreatedBy       *int64            `gorm:"column:created_by" json:"created_by"`
	UpdatedBy       *int64            `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy       *int64            `gorm:"column:deleted_by" json:"deleted_by"`
	DeletedAt       gorm.DeletedAt    `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	CreatedAt       *time.Time        `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       *time.Time        `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relationships
	Details         []ARReceiptDetail `gorm:"foreignKey:ARReceiptID;references:ID" json:"details,omitempty"`
}

// TableName specifies the table name for ARReceipt model
func (ARReceipt) TableName() string {
	return "alana.ar_receipt"
}

// BeforeCreate hook to generate UUID before creating a new record
func (ar *ARReceipt) BeforeCreate(tx *gorm.DB) error {
	if ar.ID == uuid.Nil {
		ar.ID = uuid.New()
	}
	return nil
}
