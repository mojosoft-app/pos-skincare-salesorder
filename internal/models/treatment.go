package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Treatment represents the treatment table in the database
type Treatment struct {
	ID                   uuid.UUID      `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	LocationID           *int           `gorm:"column:location_id" json:"location_id"`
	CustomerID           *int           `gorm:"column:costumer_id" json:"customer_id"`
	SalesOrderID         *int           `gorm:"column:salesorder_id" json:"sales_order_id"`
	SalesOrderDetailID   *int           `gorm:"column:salesorderdetail_id" json:"sales_order_detail_id"`
	SalesOrderServiceID  *int           `gorm:"column:salesorderservice_id" json:"sales_order_service_id"`
	ServiceID            *int           `gorm:"column:service_id" json:"service_id"`
	PatientID            *int           `gorm:"column:patient_id" json:"patient_id"`
	DoctorID             *int           `gorm:"column:doctor_id" json:"doctor_id"`
	NurseID              *int           `gorm:"column:nurse_id" json:"nurse_id"`
	BeauticianID         *int           `gorm:"column:beautician_id" json:"beautician_id"`
	DocNumber            *string        `gorm:"column:docnumber" json:"doc_number"`
	DocDate              *time.Time     `gorm:"column:docdate;type:date" json:"doc_date"`
	PostedDate           *time.Time     `gorm:"column:posteddate;type:date" json:"posted_date"`
	ServiceText          *string        `gorm:"column:servicetext" json:"service_text"`
	Note                 *string        `gorm:"column:note" json:"note"`
	StatusID             *int           `gorm:"column:status_id" json:"status_id"`
	CreatedBy            *int64         `gorm:"column:created_by" json:"created_by"`
	UpdatedBy            *int64         `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy            *int64         `gorm:"column:deleted_by" json:"deleted_by"`
	DeletedAt            gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	CreatedAt            *time.Time     `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt            *time.Time     `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relationships
	Details              []TreatmentDetail `gorm:"foreignKey:TreatmentID;references:ID" json:"details,omitempty"`
}

// TableName specifies the table name for Treatment model
func (Treatment) TableName() string {
	return "alana.treatment"
}

// BeforeCreate hook to generate UUID before creating a new record
func (t *Treatment) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}
