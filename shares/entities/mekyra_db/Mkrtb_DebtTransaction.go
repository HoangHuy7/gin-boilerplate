package mekyra_db

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// =========================
// DEBT TRANSACTIONS
// =========================
type Mkrtb_DebtTransaction struct {
	Id        uuid.UUID       `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	CustomerId uuid.UUID      `gorm:"column:customer_id;type:uuid"`
	OrderId   *uuid.UUID      `gorm:"column:order_id;type:uuid"`
	Amount    decimal.Decimal `gorm:"column:amount;type:decimal(12,2)"`
	Type      string          `gorm:"column:type;type:text"` // borrow / pay
	Note      string          `gorm:"column:note;type:text"`
	CreatedAt time.Time       `gorm:"column:created_at;autoCreateTime"`
}

func (Mkrtb_DebtTransaction) TableName() string {
	return "mkrtb_debt_transactions"
}
