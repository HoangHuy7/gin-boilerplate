package view

type Vw_UserMenu struct {
	UserID     string `gorm:"column:user_id;primaryKey" json:"user_id"`
	ID         int    `gorm:"column:id;primaryKey" json:"id"`
	ParentID   int    `gorm:"column:parent_id" json:"parent_id"`
	ScreenCode string `gorm:"column:screen_code" json:"screen_code"`
	Title      string `gorm:"column:title" json:"title"`
	Icon       string `gorm:"column:icon" json:"icon"`
	Level      int    `gorm:"column:level" json:"level"`
}
