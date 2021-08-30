package services

import (
	"net/http"
	"time"

	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/domains/v1"
	repositories "github.com/bitokss/bitok-user-service/repositories/postgres/v1"
	"github.com/golang-jwt/jwt"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersServiceInterface interface {
	Create(user domains.CreateUsersRequest) (rest_response.RestResp, rest_response.RestResp)
	FindAll(limit, offset int) (rest_response.RestResp, rest_response.RestResp)
	Find(uid int) (rest_response.RestResp, rest_response.RestResp)
	Delete(uid int) (rest_response.RestResp, rest_response.RestResp)
	Update(uid int, user domains.CreateUsersRequest) (rest_response.RestResp, rest_response.RestResp)
	FindByToken(token string) (rest_response.RestResp, rest_response.RestResp)
	Login(body domains.LoginRequest) (rest_response.RestResp, rest_response.RestResp)
	FindByUsername(username string) (rest_response.RestResp, rest_response.RestResp)
}

type usersService struct{}

func (*usersService) Update(uid int, user domains.CreateUsersRequest) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (*usersService) Find(uid int) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (*usersService) Delete(uid int) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (*usersService) FindAll(limit, offset int) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (*usersService) Create(user domains.CreateUsersRequest) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (u *usersService) FindByToken(token string) (rest_response.RestResp, rest_response.RestResp) {
	claim, err := JwtService.Verify(token)
	if err != nil {
		return nil, err
	}
	user, err := repositories.UsersRepository.FindByID(uint(claim.UID))
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(constants.SuccessOperation, user, http.StatusOK), nil
}

func (u *usersService) Login(body domains.LoginRequest) (rest_response.RestResp, rest_response.RestResp) {
	user, err := repositories.UsersRepository.FindByPhoneAndPassword(body.Phone, body.Password)
	if err != nil {
		return nil, err
	}
	claim := domains.Jwt{
		UID: int(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7 * 2).Unix(),
		},
	}
	token, err := JwtService.Generate(claim)
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(constants.SuccessLogin, domains.TokenResp{Token: token}, http.StatusOK), nil
}

func (u *usersService) FindByUsername(username string) (rest_response.RestResp, rest_response.RestResp) {
	user, err := repositories.UsersRepository.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(constants.SuccessOperation, user, http.StatusOK), nil
}
