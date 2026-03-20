package mekyra_db

import (
	"time"
)

// =========================
// LOGS
// =========================
type Mkrtb_Logs struct {
	Id        int64     `gorm:"column:id;type:bigserial;primaryKey;autoIncrement"`
	Action    string    `gorm:"column:action;type:varchar(100)"`
	OldData   []byte    `gorm:"column:old_data;type:jsonb"`
	NewData   []byte    `gorm:"column:new_data;type:jsonb"`
	CreatedAt time.Time `gorm:"column:created_at;default:now()"`
	CreatedBy string    `gorm:"column:created_by;type:varchar(100)"`
	IpAddress string    `gorm:"column:ip_address;type:varchar(45)"`
	UserAgent string    `gorm:"column:user_agent;type:text"`
	RequestId string    `gorm:"column:request_id;type:varchar(100)"`
}

func (Mkrtb_Logs) TableName() string {
	return "mkrtb_logs"
}
