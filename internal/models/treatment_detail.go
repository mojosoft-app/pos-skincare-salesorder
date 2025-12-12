package models

import (
	"time"

	"gorm.io/gorm"
)

// TreatmentDetail represents the treatment_detail table in the database
type TreatmentDetail struct {
	ID          int            `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	TreatmentID *int           `gorm:"column:treatment_id" json:"treatment_id"`
	ItemID      *int           `gorm:"column:item_id" json:"item_id"`
	UnitID      *int           `gorm:"column:unit_id" json:"unit_id"`
	Quantity    *int           `gorm:"column:quantity" json:"quantity"`
	CreatedBy   *int64         `gorm:"column:created_by" json:"created_by"`
	UpdatedBy   *int64         `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy   *int64         `gorm:"column:deleted_by" json:"deleted_by"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	CreatedAt   *time.Time     `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName specifies the table name for TreatmentDetail model
func (TreatmentDetail) TableName() string {
	return "alana.treatment_detail"
}
