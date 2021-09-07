package repositories

import (
	"errors"
	"fmt"
	"strings"

	"github.com/alidevjimmy/go-rest-utils/crypto"
	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/domains/v1"
	"gorm.io/gorm"
)

var (
	UsersRepository usersRepositoryInterface = &usersRepository{}
)

type usersRepositoryInterface interface {
	FindByPhoneAndPassword(phone, password string) (domains.UserResp, rest_response.RestResp)
	FindByID(id uint) (domains.UserResp, rest_response.RestResp)
	FindByUsername(username string) (domains.UserResp, rest_response.RestResp)
	Create(user domains.User) (domains.UserResp, rest_response.RestResp)
	Update(user domains.User) (domains.UserResp, rest_response.RestResp)
	FindAll(limit, offset int) ([]domains.UserResp, int64, rest_response.RestResp)
	Find(id uint) (domains.UserResp, rest_response.RestResp)
	Delete(id uint) rest_response.RestResp
	FindByPhone(phone string) (domains.UserResp, rest_response.RestResp)
	UpdatePassword(phone, newPassword string) (domains.UserResp, rest_response.RestResp)
}

type usersRepository struct{}

func (*usersRepository) FindByPhoneAndPassword(phone, password string) (domains.UserResp, rest_response.RestResp) {
	var user domains.User
	password = crypto.GenerateSha256(password)
	if err := DB.Preload("Roles").Preload("Roles.Permissions").Preload("Level").Where("phone = ? AND password = ?", phone, password).Find(&user).Error; err != nil {
		fmt.Println(err)
		return domains.UserResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	if user.ID == 0 {
		return domains.UserResp{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.WrongPhoneOrPasswordErr), nil)
	}
	userResp := serializeUser(user)
	return userResp, nil
}

func (*usersRepository) FindByID(id uint) (domains.UserResp, rest_response.RestResp) {
	var user domains.User
	if err := DB.Preload("Roles").Preload("Roles.Permissions").Preload("Level").Where("id = ?", id).Find(&user).Error; err != nil {
		fmt.Println(err)
		return domains.UserResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	if user.ID == 0 {
		return domains.UserResp{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.User), nil)
	}
	userResp := serializeUser(user)
	return userResp, nil
}

func (*usersRepository) FindByUsername(username string) (domains.UserResp, rest_response.RestResp) {
	var user domains.User
	if err := DB.Preload("Roles").Preload("Roles.Permissions").Preload("Level").Where("username = ?", username).Find(&user).Error; err != nil {
		fmt.Println(err)
		return domains.UserResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	if user.ID == 0 {
		return domains.UserResp{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.User), nil)
	}
	userResp := serializeUser(user)
	return userResp, nil
}

func (*usersRepository) Create(user domains.User) (domains.UserResp, rest_response.RestResp) {
	user.Password = crypto.GenerateSha256(user.Password)
	if err := DB.Preload("Roles").Preload("Roles.Permissions").Preload("Level").Create(&user).First(&user).Error; err != nil {
		if strings.Contains(err.Error(), "phone") {
			return domains.UserResp{}, rest_response.NewBadRequestError(fmt.Sprintf(constants.UserWithSpecificVariableFieldExists, constants.Phone), nil)
		}
		if strings.Contains(err.Error(), "username") {
			return domains.UserResp{}, rest_response.NewBadRequestError(fmt.Sprintf(constants.UserWithSpecificVariableFieldExists, constants.Username), nil)
		}
		return domains.UserResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	var profile domains.Profile
	profile.UserID = user.Model.ID
	if err := DB.Create(&profile).Error; err != nil {
		fmt.Println(err)
		return domains.UserResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	rResp := serializeUser(user)
	return rResp, nil
}

func (*usersRepository) Update(user domains.User) (domains.UserResp, rest_response.RestResp) {
	// replace all related Roles in order to overwrite them
	if err := DB.Model(&user).Association("Roles").Replace(&user.Roles); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domains.UserResp{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.User), nil)
		}
	}
	// update user and its roles
	user.Password = crypto.GenerateSha256(user.Password)
	if err := DB.Model(&user).Preload("Roles").Preload("Roles.Permissions").Preload("Level").Where("id = ?", user.Model.ID).First(&domains.User{}).Updates(&user).First(&user).Error; err != nil {
		if strings.Contains(err.Error(), "username") {
			return domains.UserResp{}, rest_response.NewBadRequestError(fmt.Sprintf(constants.UserWithSpecificVariableFieldExists, constants.Username), nil)
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domains.UserResp{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.User), nil)
		}
		return domains.UserResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	rResp := serializeUser(user)
	return rResp, nil
}

func (*usersRepository) FindAll(limit, offset int) ([]domains.UserResp, int64, rest_response.RestResp) {
	var users []domains.User
	var c int64
	if err := DB.Limit(limit).Offset(offset).Preload("Roles").Preload("Roles.Permissions").Preload("Level").Find(&users).Error; err != nil {
		return []domains.UserResp{}, 0, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	if err := DB.Model(&domains.User{}).Count(&c).Error; err != nil {
		return []domains.UserResp{}, 0, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	rResps := []domains.UserResp{}
	for _, v := range users {
		rResps = append(rResps, serializeUser(v))
	}
	return rResps, c, nil
}

func (*usersRepository) Find(id uint) (domains.UserResp, rest_response.RestResp) {
	var r domains.User
	if err := DB.Where("id = ?", id).Preload("Roles").Preload("Roles.Permissions").Preload("Level").Find(&r).Error; err != nil {
		return domains.UserResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	if r.Model.ID == 0 {
		return domains.UserResp{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.User), nil)
	}
	rResp := serializeUser(r)
	return rResp, nil
}

func (*usersRepository) Delete(id uint) rest_response.RestResp {
	var user = domains.User{
		Model: gorm.Model{
			ID: id,
		},
	}
	if err := DB.First(&user).Delete(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.User), nil)
		}
		return rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	return nil
}

func (*usersRepository) UpdatePassword(phone, newPassword string) (domains.UserResp, rest_response.RestResp) {
	var user domains.User
	newPassword = crypto.GenerateSha256(newPassword)
	if err := DB.Model(&domains.User{}).Preload("Roles").Preload("Roles.Permissions").Preload("Level").Where("phone = ?", phone).First(&user).Update("password", newPassword).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domains.UserResp{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.User), nil)
		}
		return domains.UserResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	rResp := serializeUser(user)
	return rResp, nil
}

func serializeUser(user domains.User) domains.UserResp {
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
	return userResp
}

func (*usersRepository) FindByPhone(phone string) (domains.UserResp, rest_response.RestResp) {
	var user domains.User
	if err := DB.Preload("Roles").Preload("Roles.Permissions").Preload("Level").Where("phone = ?", phone).Find(&user).Error; err != nil {
		fmt.Println(err)
		return domains.UserResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	if user.ID == 0 {
		return domains.UserResp{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.User), nil)
	}
	userResp := serializeUser(user)
	return userResp, nil
}
