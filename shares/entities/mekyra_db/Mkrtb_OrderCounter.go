package mekyra_db

// =========================
// ORDER COUNTERS
// =========================
type Mkrtb_OrderCounter struct {
	Id            string `gorm:"column:id;primaryKey"`
	CurrentNumber int    `gorm:"column:current_number"`
}

func (Mkrtb_OrderCounter) TableName() string {
	return "mkrtb_order_counters"
}
