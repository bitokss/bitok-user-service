package domains

import "gorm.io/gorm"

type Level struct {
	gorm.Model
	Title string `gorm:"column:title;not null;type:varchar(255)"`
	Color string `gorm:"column:color;not null;type:varchar(255)"`
}

type CreateLevelsRequest struct {
	Title string `json:"title" validate:"required"`
	Color string `json:"color" validate:"required"`
}