package models

import (
	"time"

	"gorm.io/gorm"
)

// PaymentMethod represents the payment_method table in the database
type PaymentMethod struct {
	ID        int            `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	Name      *string        `gorm:"column:name" json:"name"`
	CreatedBy *int64         `gorm:"column:created_by" json:"created_by"`
	UpdatedBy *int64         `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy *int64         `gorm:"column:deleted_by" json:"deleted_by"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	CreatedAt *time.Time     `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName specifies the table name for PaymentMethod model
func (PaymentMethod) TableName() string {
	return "alana.payment_method"
}
