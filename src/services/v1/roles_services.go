package services

import (
	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/domains/v1"
)

var (
	RolesService rolesServiceInterface = &rolesService{}
)

type rolesServiceInterface interface {
	Create(role domains.CreateRolesRequest) (rest_response.RestResp, rest_response.RestResp)
	FindAll(limit , offset int) (rest_response.RestResp, rest_response.RestResp)
	Find(rid int) (rest_response.RestResp, rest_response.RestResp)
	Delete(rid int) (rest_response.RestResp, rest_response.RestResp)
	Update(rid int, role domains.CreateRolesRequest) (rest_response.RestResp, rest_response.RestResp)
}

type rolesService struct {}

func (r *rolesService) Update(rid int, role domains.CreateRolesRequest) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (r *rolesService) Find(rid int) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (r *rolesService) Delete(rid int) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (r *rolesService) FindAll(limit, offset int) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (r *rolesService) Create(role domains.CreateRolesRequest) (rest_response.RestResp, rest_response.RestResp) {
	return nil , nil
}
