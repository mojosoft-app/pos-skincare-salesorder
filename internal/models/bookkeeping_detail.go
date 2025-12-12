package models

import (
	"time"

	"gorm.io/gorm"
)

// BookkeepingDetail represents the bookeeping_detail table in the database
// Note: Table name has a typo in the database (bookeeping instead of bookkeeping)
type BookkeepingDetail struct {
	ID              int            `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	BookkeepingID   *int           `gorm:"column:bookkeeping_id" json:"bookkeeping_id"`
	TypeID          *int           `gorm:"column:type_id" json:"type_id"`
	CategoryID      *int           `gorm:"column:category_id" json:"category_id"`
	PaymentMethodID *int           `gorm:"column:paymentmethod_id" json:"payment_method_id"`
	PostedDate      *time.Time     `gorm:"column:posteddate;type:date" json:"posted_date"`
	DocNumber       *string        `gorm:"column:docnumber" json:"doc_number"`
	Income          *float64       `gorm:"column:income;type:numeric" json:"income"`
	Expanse         *float64       `gorm:"column:expanse;type:numeric" json:"expanse"`
	Description     *string        `gorm:"column:description" json:"description"`
	CreatedBy       *int64         `gorm:"column:created_by" json:"created_by"`
	UpdatedBy       *int64         `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy       *int64         `gorm:"column:deleted_by" json:"deleted_by"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	CreatedAt       *time.Time     `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       *time.Time     `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relationships
	Bookkeeping         *Bookkeeping              `gorm:"foreignKey:BookkeepingID;references:ID" json:"bookkeeping,omitempty"`
	Type                *BookTransactionType      `gorm:"foreignKey:TypeID;references:ID" json:"type,omitempty"`
	Category            *BookTransactionCategory  `gorm:"foreignKey:CategoryID;references:ID" json:"category,omitempty"`
	PaymentMethod       *PaymentMethod            `gorm:"foreignKey:PaymentMethodID;references:ID" json:"payment_method,omitempty"`
}

// TableName specifies the table name for BookkeepingDetail model
// Note: The actual table name in database has a typo (bookeeping_detail)
func (BookkeepingDetail) TableName() string {
	return "alana.bookeeping_detail"
}
