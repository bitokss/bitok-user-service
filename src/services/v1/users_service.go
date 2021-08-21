package services

import (
	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/domains/v1"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersServiceInterface interface {
	Create(role domains.CreateUsersRequest) (rest_response.RestResp, rest_response.RestResp)
	FindAll(limit , offset int) (rest_response.RestResp, rest_response.RestResp)
	Find(pid int) (rest_response.RestResp, rest_response.RestResp)
	Delete(pid int) (rest_response.RestResp, rest_response.RestResp)
	Update(pid int, role domains.CreateUsersRequest) (rest_response.RestResp, rest_response.RestResp)
	FindByToken(token string) (rest_response.RestResp , rest_response.RestResp)
}

type usersService struct {}

func (u *usersService) Update(pid int, role domains.CreateUsersRequest) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (u *usersService) Find(pid int) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (u *usersService) Delete(pid int) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (u *usersService) FindAll(limit, offset int) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (u *usersService) Create(role domains.CreateUsersRequest) (rest_response.RestResp, rest_response.RestResp) {
	return nil , nil
}

func (u *usersService) FindByToken(token string) (rest_response.RestResp , rest_response.RestResp) {
	return nil,nil
}
