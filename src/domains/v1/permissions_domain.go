package domains

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Title  string `gorm:"column:title;not null;type:varchar(255)"`
	Symbol string `gorm:"column:symbol;not null;type:varchar(255)"`
}

type CreatePermissionsRequest struct {
	Title string `json:"title" validate:"required"`
	Symbol string `json:"symbol" validate:"required"`
}