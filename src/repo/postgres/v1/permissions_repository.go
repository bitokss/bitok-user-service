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
	PermissionsRepository permissionsRepositoryInterface = &permissionsRepository{}
)

type permissionsRepositoryInterface interface {
	Create(permission domains.Permission) (domains.PermissionResp, rest_response.RestResp)
	Update(permission domains.Permission) (domains.PermissionResp, rest_response.RestResp)
	FindAll(limit, offset int) ([]domains.PermissionResp, int64, rest_response.RestResp)
	Find(id uint) (domains.PermissionResp, rest_response.RestResp)
	Delete(id uint) rest_response.RestResp
}

type permissionsRepository struct{}

func (*permissionsRepository) Create(permission domains.Permission) (domains.PermissionResp, rest_response.RestResp) {
	if err := DB.Create(&permission).Error; err != nil {
		return domains.PermissionResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	pResp := serialaizePermission(permission)
	return pResp, nil
}
func (*permissionsRepository) Update(permission domains.Permission) (domains.PermissionResp, rest_response.RestResp) {
	if err := DB.Model(&permission).Where("id = ?", permission.Model.ID).First(&domains.Permission{}).Updates(permission).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domains.PermissionResp{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.Permission), nil)
		}
		return domains.PermissionResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	pResp := serialaizePermission(permission)
	return pResp, nil
}
func (*permissionsRepository) FindAll(limit, offset int) ([]domains.PermissionResp, int64, rest_response.RestResp) {
	var permissions []domains.Permission
	var c int64
	if err := DB.Limit(limit).Offset(offset).Find(&permissions).Error; err != nil {
		return []domains.PermissionResp{}, 0, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	if err := DB.Model(&domains.Permission{}).Count(&c).Error; err != nil {
		return []domains.PermissionResp{}, 0, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	pResps := []domains.PermissionResp{}
	for _, v := range permissions {
		pResps = append(pResps, domains.PermissionResp{
			ID:     v.Model.ID,
			Title:  v.Title,
			Symbol: v.Symbol,
		})
	}
	return pResps, c, nil
}
func (*permissionsRepository) Find(id uint) (domains.PermissionResp, rest_response.RestResp) {
	var p domains.Permission
	if err := DB.Where("id = ?", id).Find(&p).Error; err != nil {
		return domains.PermissionResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	if p.Model.ID == 0 || p == (domains.Permission{}) {
		return domains.PermissionResp{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.Permission), nil)
	}
	pResp := serialaizePermission(p)
	return pResp, nil
}
func (*permissionsRepository) Delete(id uint) rest_response.RestResp {
	var permission = domains.Permission{
		Model: gorm.Model{
			ID: id,
		},
	}
	if err := DB.First(&permission).Delete(&permission).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.Permission), nil)
		}
		return rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	return nil
}

func serialaizePermission(permission domains.Permission) domains.PermissionResp {
	pResp := domains.PermissionResp{
		ID:     permission.Model.ID,
		Title:  permission.Title,
		Symbol: permission.Symbol,
	}
	return pResp
}
