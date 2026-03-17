package mekyra_db

import (
	"time"

	"github.com/google/uuid"
)

// =========================
// CUSTOMERS
// =========================
type Mkrtb_Customer struct {
	Id        uuid.UUID `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"column:name;type:text;not null"`
	Phone     string    `gorm:"column:phone;type:text"`
	Address   string    `gorm:"column:address;type:text"`
	TotalDebt string    `gorm:"column:total_debt;type:decimal(12,2);default:0"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (Mkrtb_Customer) TableName() string {
	return "mkrtb_customers"
}
