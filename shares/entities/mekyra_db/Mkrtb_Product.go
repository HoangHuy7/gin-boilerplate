package mekyra_db

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// =========================
// PRODUCTS
// =========================
type Mkrtb_Product struct {
	Id            uuid.UUID       `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	Name          string          `gorm:"column:name;type:text;not null"`
	Category      string          `gorm:"column:category;type:text"`
	Unit          string          `gorm:"column:unit;type:text"`
	Price         decimal.Decimal `gorm:"column:price;type:decimal(12,2);not null"`
	CostPrice     decimal.Decimal `gorm:"column:cost_price;type:decimal(12,2)"`
	StockQuantity int             `gorm:"column:stock_quantity;default:0"`
	Barcode       string          `gorm:"column:barcode;type:text;unique"`
	CreatedAt     time.Time       `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time       `gorm:"column:updated_at;autoUpdateTime"`
	ImgUrl        string          `gorm:"column:img_url;type:varchar(1024)"`
}

func (Mkrtb_Product) TableName() string {
	return "mkrtb_products"
}
