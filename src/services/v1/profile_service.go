package services

import (
	"fmt"
	"net/http"

	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/src/constants"
	"github.com/bitokss/bitok-user-service/src/domains/v1"
	repositories "github.com/bitokss/bitok-user-service/src/repo/postgres/v1"
)

var (
	ProfilesService profilesServiceInterface = &profilesService{}
)

type profilesServiceInterface interface {
	Update(id uint, profile domains.ProfileRequest) (rest_response.RestResp, rest_response.RestResp)
	Find(id uint) (rest_response.RestResp, rest_response.RestResp)
}

type profilesService struct{}

func (*profilesService) Update(uid uint, profile domains.ProfileRequest) (rest_response.RestResp, rest_response.RestResp) {
	p := domains.Profile{
		Summery: profile.Summery,
		UserID:  uid,
	}
	pr, err := repositories.ProfilesRepository.Update(p)
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(fmt.Sprintf(constants.SuccessUpdateOperation, constants.Profile), pr, http.StatusOK), nil
}

func (*profilesService) Find(id uint) (rest_response.RestResp, rest_response.RestResp) {
	pr, err := repositories.ProfilesRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(constants.SuccessOperation, pr, http.StatusOK), nil
}
