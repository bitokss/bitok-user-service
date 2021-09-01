package repositories

import (
	"errors"
	"fmt"

	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/domains/v1"
	"gorm.io/gorm"
)

var (
	RolesRepository rolesRepositoryInterface = &rolesRepository{}
)

type rolesRepositoryInterface interface {
	Create(role domains.Role) (domains.RoleResp, rest_response.RestResp)
	Update(role domains.Role) (domains.RoleResp, rest_response.RestResp)
	FindAll(limit, offset int) ([]domains.RoleResp, int64, rest_response.RestResp)
	Find(id uint) (domains.RoleResp, rest_response.RestResp)
	Delete(id uint) rest_response.RestResp
}

type rolesRepository struct{}

func (*rolesRepository) Create(role domains.Role) (domains.RoleResp, rest_response.RestResp) {
	if err := DB.Preload("Permissions").Create(&role).First(&role).Error; err != nil {
		return domains.RoleResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	rResp := serialaizeRole(role)
	return rResp, nil
}

func (*rolesRepository) Update(role domains.Role) (domains.RoleResp, rest_response.RestResp) {
	// clear all related permissions in order to overwrite them
	if err := DB.Model(&role).Association("Permissions").Replace(&role.Permissions); err != nil {
		if errors.Is(err , gorm.ErrRecordNotFound) {
			return domains.RoleResp{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.Role), nil)
		}
	}
	// update role and its permissions
	if err := DB.Model(&role).Preload("Permissions").Where("id = ?", role.Model.ID).First(&domains.Role{}).Updates(&role).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domains.RoleResp{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.Role), nil)
		}
		return domains.RoleResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	rResp := serialaizeRole(role)
	return rResp, nil
}

func (*rolesRepository) FindAll(limit, offset int) ([]domains.RoleResp, int64, rest_response.RestResp) {
	var roles []domains.Role
	var c int64
	if err := DB.Limit(limit).Offset(offset).Preload("Permissions").Find(&roles).Error; err != nil {
		return []domains.RoleResp{}, 0, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	if err := DB.Model(&domains.Role{}).Count(&c).Error; err != nil {
		return []domains.RoleResp{}, 0, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	rResps := []domains.RoleResp{}
	for _, v := range roles {
		rResps = append(rResps, serialaizeRole(v))
	}
	return rResps, c, nil
}

func (*rolesRepository) Find(id uint) (domains.RoleResp, rest_response.RestResp) {
	var r domains.Role
	if err := DB.Where("id = ?", id).Preload("Permissions").Find(&r).Error; err != nil {
		return domains.RoleResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	if r.Model.ID == 0 {
		return domains.RoleResp{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.Role), nil)
	}
	rResp := serialaizeRole(r)
	return rResp, nil
}
func (*rolesRepository) Delete(id uint) rest_response.RestResp {
	var role = domains.Role{
		Model: gorm.Model{
			ID: id,
		},
	}
	if err := DB.First(&role).Delete(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.Role), nil)
		}
		return rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	return nil
}

func serialaizeRole(role domains.Role) domains.RoleResp {
	permissions := []domains.PermissionResp{}
	for _, v := range role.Permissions {
		permissions = append(permissions, serialaizePermission(v))
	}
	rResp := domains.RoleResp{
		ID:          role.Model.ID,
		Title:       role.Title,
		Permissions: permissions,
	}
	return rResp
}
