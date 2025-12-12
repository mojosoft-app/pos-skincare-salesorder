package models

import (
	"time"

	"gorm.io/gorm"
)

// SummaryByTransactionTypeAndPaymentMethod represents the summary_by_transaction_type_and_payment_method table in the database
type SummaryByTransactionTypeAndPaymentMethod struct {
	ID              int            `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	BookkeepingID   *int           `gorm:"column:bookkeeping_id" json:"bookkeeping_id"`
	TypeID          *int           `gorm:"column:type_id" json:"type_id"`
	PaymentMethodID *int           `gorm:"column:paymentmethod_id" json:"payment_method_id"`
	Total           *float64       `gorm:"column:total;type:numeric" json:"total"`
	CreatedBy       *int64         `gorm:"column:created_by" json:"created_by"`
	UpdatedBy       *int64         `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy       *int64         `gorm:"column:deleted_by" json:"deleted_by"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	CreatedAt       *time.Time     `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       *time.Time     `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relationships
	Bookkeeping   *Bookkeeping         `gorm:"foreignKey:BookkeepingID;references:ID" json:"bookkeeping,omitempty"`
	Type          *BookTransactionType `gorm:"foreignKey:TypeID;references:ID" json:"type,omitempty"`
	PaymentMethod *PaymentMethod       `gorm:"foreignKey:PaymentMethodID;references:ID" json:"payment_method,omitempty"`
}

// TableName specifies the table name for SummaryByTransactionTypeAndPaymentMethod model
func (SummaryByTransactionTypeAndPaymentMethod) TableName() string {
	return "alana.summary_by_transaction_type_and_payment_method"
}
