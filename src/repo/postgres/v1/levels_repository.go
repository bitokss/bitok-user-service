package repo

import (
	"errors"
	"fmt"

	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/src/constants"
	"github.com/bitokss/bitok-user-service/src/domains/v1"
	"gorm.io/gorm"
)

var (
	LevelsRepository levelsRepositoryInterface = &levelsRepository{}
)

type levelsRepositoryInterface interface {
	Create(level domains.Level) (domains.LevelResp, rest_response.RestResp)
	Update(level domains.Level) (domains.LevelResp, rest_response.RestResp)
	FindAll(limit, offset int) ([]domains.LevelResp, int64, rest_response.RestResp)
	Find(id uint) (domains.LevelResp, rest_response.RestResp)
	Delete(id uint) rest_response.RestResp
}

type levelsRepository struct{}

func (*levelsRepository) Create(level domains.Level) (domains.LevelResp, rest_response.RestResp) {
	if err := DB.Create(&level).Error; err != nil {
		return domains.LevelResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	pResp := serialaizeLevel(level)
	return pResp, nil
}
func (*levelsRepository) Update(level domains.Level) (domains.LevelResp, rest_response.RestResp) {
	if err := DB.Model(&level).Where("id = ?", level.Model.ID).First(&domains.Level{}).Updates(level).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domains.LevelResp{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.Level), nil)
		}
		return domains.LevelResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	pResp := serialaizeLevel(level)
	return pResp, nil
}
func (*levelsRepository) FindAll(limit, offset int) ([]domains.LevelResp, int64, rest_response.RestResp) {
	var levels []domains.Level
	var c int64
	if err := DB.Limit(limit).Offset(offset).Find(&levels).Error; err != nil {
		return []domains.LevelResp{}, 0, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	if err := DB.Model(&domains.Level{}).Count(&c).Error; err != nil {
		return []domains.LevelResp{}, 0, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	pResps := []domains.LevelResp{}
	for _, v := range levels {
		pResps = append(pResps, domains.LevelResp{
			ID:    v.Model.ID,
			Title: v.Title,
			Color: v.Color,
		})
	}
	return pResps, c, nil
}
func (*levelsRepository) Find(id uint) (domains.LevelResp, rest_response.RestResp) {
	var p domains.Level
	if err := DB.Where("id = ?", id).Find(&p).Error; err != nil {
		return domains.LevelResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	if p.Model.ID == 0 || p == (domains.Level{}) {
		return domains.LevelResp{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.Level), nil)
	}
	pResp := serialaizeLevel(p)
	return pResp, nil
}
func (*levelsRepository) Delete(id uint) rest_response.RestResp {
	var level = domains.Level{
		Model: gorm.Model{
			ID: id,
		},
	}
	if err := DB.First(&level).Delete(&level).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.Level), nil)
		}
		return rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	return nil
}

func serialaizeLevel(level domains.Level) domains.LevelResp {
	pResp := domains.LevelResp{
		ID:    level.Model.ID,
		Title: level.Title,
		Color: level.Color,
	}
	return pResp
}
