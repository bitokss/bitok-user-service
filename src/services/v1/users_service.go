package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/domains/v1"
	repositories "github.com/bitokss/bitok-user-service/repositories/postgres/v1"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersServiceInterface interface {
	Create(user domains.CreateUsersRequest) (rest_response.RestResp, rest_response.RestResp)
	Register(user domains.RegisterRequest) (rest_response.RestResp, rest_response.RestResp)
	ResetPassword(body domains.ResetPasswordRequest) (rest_response.RestResp, rest_response.RestResp)
	FindAll(limit, offset int) (rest_response.RestResp, rest_response.RestResp)
	Find(id int) (rest_response.RestResp, rest_response.RestResp)
	Delete(id int) (rest_response.RestResp, rest_response.RestResp)
	Update(id uint, user domains.UpdateUsersRequest) (rest_response.RestResp, rest_response.RestResp)
	FindByToken(token string) (rest_response.RestResp, rest_response.RestResp)
	Login(body domains.LoginRequest) (rest_response.RestResp, rest_response.RestResp)
	FindByUsername(username string) (rest_response.RestResp, rest_response.RestResp)
}

type usersService struct{}

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

func (*usersService) Update(id uint, user domains.UpdateUsersRequest) (rest_response.RestResp, rest_response.RestResp) {
	u := userReqToDomain(user)
	u.Model.ID = id
	p, err := repositories.UsersRepository.Update(u)
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(fmt.Sprintf(constants.SuccessUpdateOperation, constants.User), p, http.StatusOK), nil
}

func (*usersService) Find(id int) (rest_response.RestResp, rest_response.RestResp) {
	p, err := repositories.UsersRepository.Find(uint(id))
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(constants.SuccessOperation, p, http.StatusOK), nil
}

func (*usersService) Delete(id int) (rest_response.RestResp, rest_response.RestResp) {
	err := repositories.UsersRepository.Delete(uint(id))
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(fmt.Sprintf(constants.SuccessDeleteOperation, constants.User), nil, http.StatusOK), nil
}

func (*usersService) FindAll(limit, offset int) (rest_response.RestResp, rest_response.RestResp) {
	r, c, err := repositories.UsersRepository.FindAll(limit, offset)
	if err != nil {
		return nil, err
	}
	fResp := struct {
		UsersCount int64              `json:"usersCount"`
		Users      []domains.UserResp `json:"users"`
	}{
		UsersCount: c,
		Users:      r,
	}
	return rest_response.NewSuccessResponse(constants.SuccessOperation, fResp, http.StatusOK), nil
}

func (*usersService) Create(user domains.CreateUsersRequest) (rest_response.RestResp, rest_response.RestResp) {
	u := userReqToDomain(user)
	p, err := repositories.UsersRepository.Create(u)
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(fmt.Sprintf(constants.SuccessCreateOperation, constants.User), p, http.StatusOK), nil
}
func (*usersService) Register(user domains.RegisterRequest) (rest_response.RestResp, rest_response.RestResp) {
	code, err := repositories.CodesRepository.Find(user.Phone, "REGISTER")
	if err != nil {
		return nil, err
	}
	if !code.Used {
		return nil, rest_response.NewBadRequestError(constants.BadRequestErr, nil)
	}
	u := userReqToDomain(user)
	p, err := repositories.UsersRepository.Create(u)
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(constants.SuccessRegisterOperation, p, http.StatusOK), nil
}

func (*usersService) ResetPassword(body domains.ResetPasswordRequest) (rest_response.RestResp, rest_response.RestResp) {
	code, err := repositories.CodesRepository.Find(body.Phone, "FORGET_PASSWORD")
	if err != nil {
		return nil, err
	}
	if !code.Used {
		return nil, rest_response.NewBadRequestError(constants.BadRequestErr, nil)
	}
	p, err := repositories.UsersRepository.UpdatePassword(body.Phone , body.NewPassword)
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(constants.SuccessRegisterOperation, p, http.StatusOK), nil
}

func userReqToDomain(u interface{}) domains.User {
	switch user := u.(type) {
	case domains.CreateUsersRequest:
		roles := []domains.Role{}
		for _, v := range user.Roles {
			roles = append(roles, domains.Role{
				Model: gorm.Model{
					ID: v,
				},
			})
		}
		level := domains.Level{
			Model: gorm.Model{
				ID: user.LevelID,
			},
		}
		return domains.User{
			Phone:        user.Phone,
			Username:     user.Username,
			Email:        user.Email,
			PersonnelNum: user.PersonnelNum,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Password:     user.Password,
			Level:        level,
			Roles:        roles,
		}
	case domains.UpdateUsersRequest:
		roles := []domains.Role{}
		for _, v := range user.Roles {
			roles = append(roles, domains.Role{
				Model: gorm.Model{
					ID: v,
				},
			})
		}
		level := domains.Level{
			Model: gorm.Model{
				ID: user.LevelID,
			},
		}
		return domains.User{
			Username:     user.Username,
			Email:        user.Email,
			PersonnelNum: user.PersonnelNum,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Level:        level,
			Password:     user.Password,
			Roles:        roles,
		}
	case domains.RegisterRequest:
		level := domains.Level{
			Model: gorm.Model{
				ID: 1,
			},
		}
		return domains.User{
			Phone:        user.Phone,
			Username:     user.Username,
			Email:        user.Email,
			PersonnelNum: user.PersonnelNum,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Level:        level,
			Password:     user.Password,
		}
	}
	return domains.User{}
}
