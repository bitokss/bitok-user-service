package repositories

import (
	"errors"
	"fmt"

	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/domains/v1"
	"gorm.io/gorm"
)

var (
	CodesRepository codesRepositoryInterface = &codesRepository{}
)

type codesRepositoryInterface interface {
	Find(phone, codeType string) (domains.Code, rest_response.RestResp)
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
