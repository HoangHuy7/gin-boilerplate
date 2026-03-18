package mekyra_db

import (
	"time"

	"github.com/google/uuid"
)

// =========================
// DELIVERIES
// =========================
type Mkrtb_Delivery struct {
	Id           uuid.UUID `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	OrderId      uuid.UUID `gorm:"column:order_id;type:uuid"`
	DeliveryDate time.Time `gorm:"column:delivery_date;type:date"`
	Status       string    `gorm:"column:status;type:text"` // pending / done
	Note         string    `gorm:"column:note;type:text"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (Mkrtb_Delivery) TableName() string {
	return "mkrtb_deliveries"
}
