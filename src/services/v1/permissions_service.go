package services

import (
	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/domains/v1"
)

var (
	PermissionsService permissionsServiceInterface = &permissionsService{}
)

type permissionsServiceInterface interface {
	Create(permission domains.CreatePermissionsRequest) (rest_response.RestResp, rest_response.RestResp)
	FindAll(limit , offset int) (rest_response.RestResp, rest_response.RestResp)
	Find(pid int) (rest_response.RestResp, rest_response.RestResp)
	Delete(pid int) (rest_response.RestResp, rest_response.RestResp)
	Update(pid int, permission domains.CreatePermissionsRequest) (rest_response.RestResp, rest_response.RestResp)
}

type permissionsService struct {}

func (p *permissionsService) Update(pid int, permission domains.CreatePermissionsRequest) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (p *permissionsService) Find(pid int) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (p *permissionsService) Delete(pid int) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (p *permissionsService) FindAll(limit, offset int) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (p *permissionsService) Create(permission domains.CreatePermissionsRequest) (rest_response.RestResp, rest_response.RestResp) {
	return nil , nil
}

