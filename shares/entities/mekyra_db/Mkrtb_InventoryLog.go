package mekyra_db

import (
	"time"

	"github.com/google/uuid"
)

// =========================
// INVENTORY LOGS
// =========================
type Mkrtb_InventoryLog struct {
	Id        uuid.UUID `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	ProductId uuid.UUID `gorm:"column:product_id;type:uuid"`
	Type      string    `gorm:"column:type;type:text"`
	Quantity  int       `gorm:"column:quantity"`
	Note      string    `gorm:"column:note;type:text"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (Mkrtb_InventoryLog) TableName() string {
	return "mkrtb_inventory_logs"
}
