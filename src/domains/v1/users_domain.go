package domains

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone string `gorm:"column:phone,unique,not null"`
	Username string `gorm:"column:username,unique,not null"`
	Email string `gorm:"column:email,unique,not null"`
	PersonnelNum int `gorm:"column:personnel_num,not null"`
	FirstName string `gorm:"column:first_name,not null"`
	LastName string `gorm:"column:last_name,not null"`
	Password string `gorm:"column:password,not null"`
	Blocked bool `gorm:"column:blocked,default:false,not null"`
	LevelID uint `gorm:"column:level_id"`
	Level Level `gorm:"foreignKey:LevelID"`
}