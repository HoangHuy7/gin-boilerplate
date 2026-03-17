package mekyra_db

import (
	"time"

	"github.com/google/uuid"
)

// =========================
// ORDERS
// =========================
type Mkrtb_Order struct {
	Id          uuid.UUID `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	Code        string    `gorm:"column:code;type:text"`
	CustomerId  uuid.UUID `gorm:"column:customer_id;type:uuid"`
	TotalAmount string    `gorm:"column:total_amount;type:decimal(12,2)"`
	PaidAmount  string    `gorm:"column:paid_amount;type:decimal(12,2)"`
	DebtAmount  string    `gorm:"column:debt_amount;type:decimal(12,2)"`
	Status      string    `gorm:"column:status;type:text"`
	Note        string    `gorm:"column:note;type:text"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (Mkrtb_Order) TableName() string {
	return "mkrtb_orders"
}
