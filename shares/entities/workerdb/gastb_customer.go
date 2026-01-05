package workerdb

import "time"

type Gastb_Customer struct {
	Id           int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name         string    `gorm:"column:name;type:varchar(255)" json:"name"`
	Phone        string    `gorm:"column:phone;type:varchar(50)" json:"phone"`
	Note         string    `gorm:"column:note;type:text" json:"note"`
	Created_date time.Time `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	Created_by   string    `gorm:"column:created_by;type:varchar(255)" json:"created_by"`
}
