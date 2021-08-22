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
	userResp , err := userSerialize(user)
	if err != nil {
		return domains.UserResp{} , err
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
	userResp , err := userSerialize(user)
	if err != nil {
		return domains.UserResp{} , err
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
	userResp , err := userSerialize(user)
	if err != nil {
		return domains.UserResp{} , err
	}
	return userResp, nil
}

func userSerialize(user domains.User) (domains.UserResp , rest_response.RestResp) {
	var level domains.Level
	//var permissions []domains.Permission
	if err := DB.Where("id = ?" , user.LevelID).Find(&level).Error; err != nil {
		fmt.Println(err)
	}
	if err := DB.Preload("Roles").Find(&user).Error; err != nil {
		fmt.Println(err)
	}
	if err := DB.Preload("Permissions").Find(&user.Roles).Error; err != nil {
		fmt.Println(err)
	}
	rolesResp := []domains.RoleResp{}
	for _ , v := range user.Roles {
		permissionsResp := []domains.PermissionResp{}
		for _ , v1 := range v.Permissions{
			p := domains.PermissionResp{
				Title: v1.Title,
				Symbol: v1.Symbol,
			}
			permissionsResp = append(permissionsResp , p)
		}
		r := domains.RoleResp{
			Title: v.Title,
			Permissions: permissionsResp,
		}
		rolesResp = append(rolesResp , r)
	}

	levelResp := domains.LevelResp{
		Title: level.Title,
		Color: level.Color,
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
	return userResp , nil
}
