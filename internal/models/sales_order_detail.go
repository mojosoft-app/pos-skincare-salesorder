package models

import (
	"time"

	"gorm.io/gorm"
)

// SalesOrderDetail represents the sales_order_detail table in the database
type SalesOrderDetail struct {
	ID              int            `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	SalesOrderID    *int           `gorm:"column:salesorder_id" json:"sales_order_id"`
	ItemID          *int           `gorm:"column:item_id" json:"item_id"`
	UnitID          *int           `gorm:"column:unit_id" json:"unit_id"`
	PromoterID      *int           `gorm:"column:promoter_id" json:"promoter_id"`
	ItemName        *string        `gorm:"column:itemname" json:"item_name"`
	Quantity        *int           `gorm:"column:quantity" json:"quantity"`
	Price           *float64       `gorm:"column:price;type:numeric" json:"price"`
	ItemTotal       *float64       `gorm:"column:itemtotal;type:numeric" json:"item_total"`
	DiscountPct     *int           `gorm:"column:discountpct" json:"discount_pct"`
	UsedSessions    *int           `gorm:"column:usedsessions" json:"used_sessions"`
	CreatedBy       *int64         `gorm:"column:created_by" json:"created_by"`
	UpdatedBy       *int64         `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy       *int64         `gorm:"column:deleted_by" json:"deleted_by"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	CreatedAt       *time.Time     `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       *time.Time     `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName specifies the table name for SalesOrderDetail model
func (SalesOrderDetail) TableName() string {
	return "alana.sales_order_detail"
}
