package mekyra_db

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// =========================
// ORDER ITEMS
// =========================
type Mkrtb_OrderItem struct {
	Id          uuid.UUID       `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	OrderId     uuid.UUID       `gorm:"column:order_id;type:uuid"`
	ProductId   uuid.UUID       `gorm:"column:product_id;type:uuid"`
	ProductName string          `gorm:"column:product_name;type:varchar(255)"`
	Quantity    int             `gorm:"column:quantity"`
	Price       decimal.Decimal `gorm:"column:price;type:decimal(12,2)"`
	Total       decimal.Decimal `gorm:"column:total;type:decimal(12,2)"`
}

func (Mkrtb_OrderItem) TableName() string {
	return "mkrtb_order_items"
}
