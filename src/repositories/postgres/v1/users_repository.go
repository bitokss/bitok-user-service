package repositories

import (
	"fmt"

	"github.com/alidevjimmy/go-rest-utils/crypto"
	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/domains/v1"
)

var (
	UsersRepository usersRepositoryInterface = &usersRepository{}
)

type usersRepositoryInterface interface {
	FindByPhoneAndPassword(phone, password string) (domains.UserResp, rest_response.RestResp)
	FindByID(id uint) (domains.UserResp, rest_response.RestResp)
	FindByUsername(username string) (domains.UserResp, rest_response.RestResp)
}

type usersRepository struct{}

func (*usersRepository) FindByPhoneAndPassword(phone, password string) (domains.UserResp, rest_response.RestResp) {
	var user domains.User
	password = crypto.GenerateSha256(password)
	if err := DB.Where("phone = ? AND password = ?", phone, password).Find(&user).Error; err != nil {
		fmt.Println(err)
		return domains.UserResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	if user.ID == 0 {
		return domains.UserResp{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.WrongPhoneOrPasswordErr), nil)
	}
	userResp, err := userSerialize(user)
	if err != nil {
		return domains.UserResp{}, err
	}
	return userResp, nil
}

func (*usersRepository) FindByID(id uint) (domains.UserResp, rest_response.RestResp) {
	var user domains.User
	if err := DB.Where("id = ?", id).Find(&user).Error; err != nil {
		fmt.Println(err)
		return domains.UserResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	if user.ID == 0 {
		return domains.UserResp{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.User), nil)
	}
	userResp, err := userSerialize(user)
	if err != nil {
		return domains.UserResp{}, err
	}
	return userResp, nil
}

func (*usersRepository) FindByUsername(username string) (domains.UserResp, rest_response.RestResp) {
	var user domains.User
	if err := DB.Where("username = ?", username).Find(&user).Error; err != nil {
		fmt.Println(err)
		return domains.UserResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	if user.ID == 0 {
		return domains.UserResp{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.User), nil)
	}
	userResp, err := userSerialize(user)
	if err != nil {
		return domains.UserResp{}, err
	}
	return userResp, nil
}

func userSerialize(user domains.User) (domains.UserResp, rest_response.RestResp) {
	if err := DB.Preload("Roles").Preload("Roles.Permissions").Preload("Level").Find(&user).Error; err != nil {
		return domains.UserResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	rolesResp := []domains.RoleResp{}
	for _, v := range user.Roles {
		rolesResp = append(rolesResp, serialaizeRole(v))
	}
	levelResp := domains.LevelResp{
		ID:    user.Level.ID,
		Title: user.Level.Title,
		Color: user.Level.Color,
	}
	userResp := domains.UserResp{
		ID:           user.ID,
		Phone:        user.Phone,
		Username:     user.Username,
		Email:        user.Email,
		PersonnelNum: user.PersonnelNum,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Blocked:      user.Blocked,
		Level:        levelResp,
		Roles:        rolesResp,
	}
	return userResp, nil
}
