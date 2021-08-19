package domains

import "gorm.io/gorm"

type Level struct {
	gorm.Model
	Title string `gorm:"column:title,not null"`
	Color string `gorm:"column:color,not null"`
}