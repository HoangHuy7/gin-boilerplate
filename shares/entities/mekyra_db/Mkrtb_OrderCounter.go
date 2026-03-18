package mekyra_db

// =========================
// ORDER COUNTERS
// =========================
type Mkrtb_OrderCounter struct {
	Id            string `gorm:"column:id;type:text;primaryKey"` // ví dụ: 20260317
	CurrentNumber int    `gorm:"column:current_number"`
}

func (Mkrtb_OrderCounter) TableName() string {
	return "mkrtb_order_counters"
}
