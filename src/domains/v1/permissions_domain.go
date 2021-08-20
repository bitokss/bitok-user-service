package domains

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Title  string `gorm:"column:title;not null;type:varchar(255)"`
	Symbol string `gorm:"column:symbol;not null;type:varchar(255)"`
}
