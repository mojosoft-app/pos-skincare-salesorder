package models

import (
	"time"

	"gorm.io/gorm"
)

// Bookkeeping represents the bookkeeping table in the database
type Bookkeeping struct {
	ID         int            `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	LocationID *string        `gorm:"column: location_id" json:"location_id"`
	BookDate   *time.Time     `gorm:"column:bookdate;type:date" json:"book_date"`
	Opening    *float64       `gorm:"column:opening;type:numeric" json:"opening"`
	Income     *float64       `gorm:"column:income;type:numeric" json:"income"`
	Expanse    *float64       `gorm:"column:expanse;type:numeric" json:"expanse"`
	Balance    *float64       `gorm:"column:balance;type:numeric" json:"balance"`
	Note       *string        `gorm:"column:note" json:"note"`
	StatusID   *int           `gorm:"column:status_id" json:"status_id"`
	CreatedBy  *int64         `gorm:"column:created_by" json:"created_by"`
	UpdatedBy  *int64         `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy  *int64         `gorm:"column:deleted_by" json:"deleted_by"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	CreatedAt  *time.Time     `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  *time.Time     `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relationships
	Status                                      *BookkeepingStatus                             `gorm:"foreignKey:StatusID;references:ID" json:"status,omitempty"`
	Details                                     []BookkeepingDetail                            `gorm:"foreignKey:BookkeepingID;references:ID" json:"details,omitempty"`
	SummaryByTransactionType                    []SummaryByTransactionType                     `gorm:"foreignKey:BookkeepingID;references:ID" json:"summary_by_transaction_type,omitempty"`
	SummaryByPaymentMethod                      []SummaryByPaymentMethod                       `gorm:"foreignKey:BookkeepingID;references:ID" json:"summary_by_payment_method,omitempty"`
	SummaryByTransactionTypeAndPaymentMethod    []SummaryByTransactionTypeAndPaymentMethod     `gorm:"foreignKey:BookkeepingID;references:ID" json:"summary_by_transaction_type_and_payment_method,omitempty"`
}

// TableName specifies the table name for Bookkeeping model
func (Bookkeeping) TableName() string {
	return "alana.bookkeeping"
}
