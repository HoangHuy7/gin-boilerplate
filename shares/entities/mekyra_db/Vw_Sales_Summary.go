package mekyra_db

import (
	"time"

	"github.com/shopspring/decimal"
)

// =========================
// VIEW: SALES SUMMARY
// =========================
type Vw_Sales_Summary struct {
	CreatedAt time.Time       `gorm:"column:created_at"`
	SaleDate  time.Time       `gorm:"column:sale_date"`  // DATE → map vào time.Time (00:00:00)
	SaleMonth time.Time       `gorm:"column:sale_month"` // DATE
	SaleYear  int             `gorm:"column:sale_year"`  // INT
	Quantity  int             `gorm:"column:quantity"`
	Total     decimal.Decimal `gorm:"column:total"`
}

func (Vw_Sales_Summary) TableName() string {
	return "vw_sales_summary"
}
