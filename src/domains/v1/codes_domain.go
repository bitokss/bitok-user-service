package domains

import "gorm.io/gorm"

type Code struct {
	gorm.Model
	Phone string `gorm:"column:phone;not null;type:varchar(255)"`
	Code  int    `gorm:"column:code;not null"`
	Used  bool   `gorm:"column:used;default:false;not null"`
}
