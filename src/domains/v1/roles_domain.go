package domains

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Title       string       `gorm:"column:title;not null;type:varchar(255)"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
}

type CreateRolesRequest struct {
	Title       string `json:"title" validate:"required"`
	Permissions []uint `json:"permissions"`
}

type RoleResp struct {
	ID          uint             `json:"id"`
	Title       string           `json:"title"`
	Permissions []PermissionResp `json:"permissions"`
}
