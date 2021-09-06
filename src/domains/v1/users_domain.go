package domains

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone        string `gorm:"column:phone;unique;not null;type:varchar(255)"`
	Username     string `gorm:"column:username;unique;not null;type:varchar(255)"`
	Email        string `gorm:"column:email;not null;type:varchar(255)"`
	PersonnelNum int    `gorm:"column:personnel_num;not null"`
	FirstName    string `gorm:"column:first_name;not null;type:varchar(255)"`
	LastName     string `gorm:"column:last_name;not null;type:varchar(255)"`
	Password     string `gorm:"column:password;not null;type:varchar(255)"`
	Blocked      bool   `gorm:"column:blocked;default:false;not null"`
	LevelID      uint   `gorm:"column:level_id"`
	Level        Level  `gorm:"foreignKey:LevelID"`
	Roles        []Role `gorm:"many2many:user_roles"`
}

type CreateUsersRequest struct {
	Phone        string `json:"phone" validate:"required"`
	Username     string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	PersonnelNum int    `json:"personnel_num" validate:"required"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Password     string `json:"password" validate:"required"`
	LevelID      uint   `json:"level_id" validate:"required"`
	Roles        []uint `json:"roles"`
}

type RegisterRequest struct {
	Phone        string `json:"phone" validate:"required"`
	Username     string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	PersonnelNum int    `json:"personnel_num" validate:"required"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Password     string `json:"password" validate:"required"`
}

type UpdateUsersRequest struct {
	Username     string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	PersonnelNum int    `json:"personnel_num" validate:"required"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Password     string `json:"password" validate:"required"`
	LevelID      uint   `json:"level_id" validate:"required"`
	Roles        []uint `json:"roles"`
}

type UserResp struct {
	ID           uint       `json:"id"`
	Phone        string     `json:"phone"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	PersonnelNum int        `json:"personnel_num"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	Blocked      bool       `json:"blocked"`
	Level        LevelResp  `json:"level"`
	Roles        []RoleResp `json:"roles"`
}

type LoginRequest struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type TokenResp struct {
	Token string `json:"token"`
}

type ResetPasswordRequest struct {
	Phone string `json:"phone"`
	NewPassword string `json:"new_password"`
}