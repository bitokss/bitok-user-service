package services

import (
	"fmt"
	"net/http"

	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/domains/v1"
	repositories "github.com/bitokss/bitok-user-service/repositories/postgres/v1"
	"gorm.io/gorm"
)

var (
	PermissionsService permissionsServiceInterface = &permissionsService{}
)

type permissionsServiceInterface interface {
	Create(permission domains.CreatePermissionsRequest) (rest_response.RestResp, rest_response.RestResp)
	FindAll(limit, offset int) (rest_response.RestResp, rest_response.RestResp)
	Find(pid int) (rest_response.RestResp, rest_response.RestResp)
	Delete(pid int) (rest_response.RestResp, rest_response.RestResp)
	Update(pid uint, permission domains.CreatePermissionsRequest) (rest_response.RestResp, rest_response.RestResp)
}

type permissionsService struct{}

func (*permissionsService) Update(pid uint, permission domains.CreatePermissionsRequest) (rest_response.RestResp, rest_response.RestResp) {
	pd := domains.Permission{
		Title:  permission.Title,
		Symbol: permission.Symbol,
		Model: gorm.Model{
			ID: pid,
		},
	}
	p, err := repositories.PermissionsRepository.Update(pd)
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(fmt.Sprintf(constants.SuccessUpdateOperation, constants.Permission), p, http.StatusOK), nil
}

func (*permissionsService) Find(pid int) (rest_response.RestResp, rest_response.RestResp) {
	p, err := repositories.PermissionsRepository.Find(uint(pid))
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(constants.SuccessOperation, p, http.StatusOK), nil
}

func (*permissionsService) Delete(pid int) (rest_response.RestResp, rest_response.RestResp) {
	err := repositories.PermissionsRepository.Delete(uint(pid))
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(fmt.Sprintf(constants.SuccessDeleteOperation, constants.Permission), nil, http.StatusOK), nil
}

func (*permissionsService) FindAll(limit, offset int) (rest_response.RestResp, rest_response.RestResp) {
	p, c, err := repositories.PermissionsRepository.FindAll(limit, offset)
	if err != nil {
		return nil, err
	}
	fResp := struct {
		PermissionCount int64                    `json:"permissionCount"`
		Permissions     []domains.PermissionResp `json:"permissions"`
	}{
		PermissionCount: c,
		Permissions:     p,
	}
	return rest_response.NewSuccessResponse(constants.SuccessOperation, fResp, http.StatusOK), nil
}

func (*permissionsService) Create(permission domains.CreatePermissionsRequest) (rest_response.RestResp, rest_response.RestResp) {
	pd := domains.Permission{
		Title:  permission.Title,
		Symbol: permission.Symbol,
		Model:  gorm.Model{},
	}
	p, err := repositories.PermissionsRepository.Create(pd)
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(fmt.Sprintf(constants.SuccessCreateOperation, constants.Permission), p, http.StatusOK), nil
}
