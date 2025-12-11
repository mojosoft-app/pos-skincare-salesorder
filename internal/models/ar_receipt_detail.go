package models

import (
	"time"

	"gorm.io/gorm"
)

// ARReceiptDetail represents the ar_receipt_detail table in the database
type ARReceiptDetail struct {
	ID            int            `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	ARReceiptID   *int           `gorm:"column:arreceipt_id" json:"ar_receipt_id"`
	SalesOrderID  *int           `gorm:"column:salesorder_id" json:"sales_order_id"`
	ReceiptAmount *float64       `gorm:"column:receiptamount;type:numeric" json:"receipt_amount"`
	CreatedBy     *int64         `gorm:"column:created_by" json:"created_by"`
	UpdatedBy     *int64         `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy     *int64         `gorm:"column:deleted_by" json:"deleted_by"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	CreatedAt     *time.Time     `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     *time.Time     `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName specifies the table name for ARReceiptDetail model
func (ARReceiptDetail) TableName() string {
	return "alana.ar_receipt_detail"
}
