package workerdb

import "time"

type Tb_Screen struct {
	Id           int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Parent_id    int       `gorm:"column:parent_id;type:int" json:"parent_id"`
	Screen_code  string    `gorm:"column:screen_code;type:varchar(255);not null" json:"screen_code"`
	Title        string    `gorm:"column:title;type:varchar(255)" json:"title"`
	Description  string    `gorm:"column:description;type:varchar(255)" json:"description"`
	Icon         string    `gorm:"column:icon;type:varchar(255)" json:"icon"`
	Created_date time.Time `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	Created_by   string    `gorm:"column:created_by;type:varchar(255)" json:"created_by"`
}
