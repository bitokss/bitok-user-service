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
	FindByPhoneAndPassword(phone , password string) (*domains.UserResp , rest_response.RestResp)
}

type usersRepository struct {}

func (u *usersRepository) FindByPhoneAndPassword(phone, password string) (*domains.UserResp, rest_response.RestResp) {
	user := new(domains.User)
	password = crypto.GenerateSha256(password)
	if err := DB.Where("phone = ? AND password = ?" , phone , password).Find(&user).Error; err != nil {
		fmt.Println(err)
		return nil , rest_response.NewInternalServerError(constants.InternalServerErr,nil)
	}
	if user == nil || user.ID == 0{
		return nil , rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr , constants.User),nil)
	}
	userResp := domains.UserResp{
		ID : user.ID,
		Phone : user.Phone,
		Username : user.Username,
		Email : user.Email,
		PersonnelNum : user.PersonnelNum,
		FirstName : user.FirstName,
		LastName : user.LastName,
		Blocked : user.Blocked,
		LevelID : user.LevelID,
		Roles : user.Roles,
	}
	return &userResp , nil
}

