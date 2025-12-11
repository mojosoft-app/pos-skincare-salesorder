package models

import (
	"time"

	"gorm.io/gorm"
)

// SalesOrderService represents the sales_order_service table in the database
type SalesOrderService struct {
	ID                  int            `gorm:"primaryKey;column:id ;autoIncrement" json:"id"`
	SalesOrderID        *int           `gorm:"column:salesorder_id" json:"sales_order_id"`
	SalesOrderDetailID  *int           `gorm:"column:salesorderdetail_id" json:"sales_order_detail_id"`
	ServiceID           *int           `gorm:"column:service_id" json:"service_id"`
	TreatmentID         *int           `gorm:"column:treatment_id" json:"treatment_id"`
	MessageLogDetailID  *string        `gorm:"column:messagelogdetail_id" json:"message_log_detail_id"`
	RemindedID          *int           `gorm:"column:reminded_id" json:"reminded_id"`
	ServiceName         *string        `gorm:"column:servicename" json:"service_name"`
	Treated             *bool          `gorm:"column:treated" json:"treated"`
	Schedule            *time.Time     `gorm:"column:schedule;type:date" json:"schedule"`
	CreatedBy           *int64         `gorm:"column:created_by" json:"created_by"`
	UpdatedBy           *int64         `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy           *int64         `gorm:"column:deleted_by" json:"deleted_by"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	CreatedAt           *time.Time     `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt           *time.Time     `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName specifies the table name for SalesOrderService model
func (SalesOrderService) TableName() string {
	return "alana.sales_order_service"
}
