package domains

import "gorm.io/gorm"

type (
	Profile struct {
		gorm.Model
		Summery string `gorm:"column:summery;type:text"`
		UserID  uint   `gorm:"column:user_id;index"`
		User    User   `gorm:"foreignKey:UserID"`
	}

	ProfileRequest struct {
		Summery string `json:"summery" validate:"required"`
	}

	ProfileResp struct {
		ID      uint   `json:"id"`
		Summery string `json:"summery"`
		UserID  uint   `json:"user_id"`
	}
)
