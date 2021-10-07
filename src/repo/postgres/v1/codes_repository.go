package repo

import (
	"errors"
	"fmt"

	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/src/constants"
	"github.com/bitokss/bitok-user-service/src/domains/v1"
	"gorm.io/gorm"
)

var (
	CodesRepository codesRepositoryInterface = &codesRepository{}
)

type codesRepositoryInterface interface {
	Find(phone, codeType string) (domains.Code, rest_response.RestResp)
	FindByCode(phone, codeType string, code int) (domains.Code, rest_response.RestResp)
	Create(phone, codeType string, code int) (domains.Code, rest_response.RestResp)
	Update(phone, codeType string, code int) (domains.Code, rest_response.RestResp)
}

type codesRepository struct{}

func (*codesRepository) Find(phone, codeType string) (domains.Code, rest_response.RestResp) {
	code := domains.Code{
		Phone: phone,
		Type:  codeType,
	}
	if err := DB.Last(&code).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domains.Code{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.Code), nil)
		}
		return domains.Code{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	return code, nil
}
func (*codesRepository) FindByCode(phone, codeType string, c int) (domains.Code, rest_response.RestResp) {
	code := domains.Code{
		Phone: phone,
		Type:  codeType,
		Code:  c,
		Used:  false,
	}
	if err := DB.Last(&code).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domains.Code{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.Code), nil)
		}
		return domains.Code{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	return code, nil
}

func (*codesRepository) Create(phone, codeType string, code int) (domains.Code, rest_response.RestResp) {
	c := domains.Code{
		Phone: phone,
		Type:  codeType,
		Code:  code,
	}
	if err := DB.Create(&c).Error; err != nil {
		return domains.Code{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	return c, nil
}

func (*codesRepository) Update(phone, codeType string, c int) (domains.Code, rest_response.RestResp) {
	code := domains.Code{
		Phone: phone,
		Type:  codeType,
		Code:  c,
		Used:  true,
	}
	if err := DB.Model(&domains.Code{}).Where("phone = ? AND code = ? AND used = ? type = ?", phone, code, false, codeType).First(&domains.Code{}).Updates(&code).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domains.Code{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.Permission), nil)
		}
		return domains.Code{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	return code, nil
}
