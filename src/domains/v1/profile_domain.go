package domains

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID uint `gorm:"column:user_id"`
	User   User `gorm:"foreignKey:UserID"`
}
