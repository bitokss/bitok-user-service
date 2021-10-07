package services

import (
	"fmt"
	"net/http"

	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/src/constants"
	"github.com/bitokss/bitok-user-service/src/domains/v1"
	repositories "github.com/bitokss/bitok-user-service/src/repo/postgres/v1"
	"gorm.io/gorm"
)

var (
	RolesService rolesServiceInterface = &rolesService{}
)

type rolesServiceInterface interface {
	Create(role domains.CreateRolesRequest) (rest_response.RestResp, rest_response.RestResp)
	FindAll(limit, offset int) (rest_response.RestResp, rest_response.RestResp)
	Find(rid int) (rest_response.RestResp, rest_response.RestResp)
	Delete(rid int) (rest_response.RestResp, rest_response.RestResp)
	Update(rid uint, role domains.CreateRolesRequest) (rest_response.RestResp, rest_response.RestResp)
}

type rolesService struct{}

func (*rolesService) Update(id uint, role domains.CreateRolesRequest) (rest_response.RestResp, rest_response.RestResp) {
	permissions := []domains.Permission{}
	for _, v := range role.Permissions {
		permissions = append(permissions, domains.Permission{
			Model: gorm.Model{
				ID: v,
			},
		})
	}
	pd := domains.Role{
		Title: role.Title,
		Model: gorm.Model{
			ID: id,
		},
		Permissions: permissions,
	}
	p, err := repositories.RolesRepository.Update(pd)
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(fmt.Sprintf(constants.SuccessUpdateOperation, constants.Role), p, http.StatusOK), nil
}

func (*rolesService) Find(id int) (rest_response.RestResp, rest_response.RestResp) {
	p, err := repositories.RolesRepository.Find(uint(id))
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(constants.SuccessOperation, p, http.StatusOK), nil
}

func (*rolesService) Delete(id int) (rest_response.RestResp, rest_response.RestResp) {
	err := repositories.RolesRepository.Delete(uint(id))
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(fmt.Sprintf(constants.SuccessDeleteOperation, constants.Role), nil, http.StatusOK), nil
}

func (*rolesService) FindAll(limit, offset int) (rest_response.RestResp, rest_response.RestResp) {
	r, c, err := repositories.RolesRepository.FindAll(limit, offset)
	if err != nil {
		return nil, err
	}
	fResp := struct {
		RolesCount int64              `json:"rolesCount"`
		Roles      []domains.RoleResp `json:"roles"`
	}{
		RolesCount: c,
		Roles:      r,
	}
	return rest_response.NewSuccessResponse(constants.SuccessOperation, fResp, http.StatusOK), nil
}

func (*rolesService) Create(role domains.CreateRolesRequest) (rest_response.RestResp, rest_response.RestResp) {
	permissions := []domains.Permission{}
	for _, v := range role.Permissions {
		permissions = append(permissions, domains.Permission{
			Model: gorm.Model{
				ID: v,
			},
		})
	}
	pd := domains.Role{
		Title:       role.Title,
		Permissions: permissions,
	}
	p, err := repositories.RolesRepository.Create(pd)
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(fmt.Sprintf(constants.SuccessCreateOperation, constants.Role), p, http.StatusOK), nil
}
