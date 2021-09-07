package repositories

import (
	"fmt"

	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/domains/v1"
)

var (
	ProfilesRepository profilesRepositoryInterface = &profilesRepository{}
)

type profilesRepositoryInterface interface {
	Update(profile domains.Profile) (domains.ProfileResp, rest_response.RestResp)
	Find(id uint) (domains.ProfileResp, rest_response.RestResp)
}

type profilesRepository struct{}

func (*profilesRepository) Update(profile domains.Profile) (domains.ProfileResp, rest_response.RestResp) {
	if err := DB.Model(&profile).Where("user_id = ?", profile.UserID).First(&domains.Profile{}).Updates(&profile).First(&profile).Error; err != nil {

		return domains.ProfileResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	rResp := serialaizeProfile(profile)
	return rResp, nil
}

func (*profilesRepository) Find(id uint) (domains.ProfileResp, rest_response.RestResp) {
	var p domains.Profile
	if err := DB.Where("user_id = ?", id).Find(&p).Error; err != nil {
		return domains.ProfileResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	if p.Model.ID == 0 {
		return domains.ProfileResp{}, rest_response.NewNotFoundError(fmt.Sprintf(constants.NotFoundErr, constants.Profile), nil)
	}
	pResp := serialaizeProfile(p)
	return pResp, nil
}

func serialaizeProfile(profile domains.Profile) domains.ProfileResp {
	pResp := domains.ProfileResp{
		ID:      profile.Model.ID,
		Summery: profile.Summery,
		UserID:  profile.UserID,
	}
	return pResp
}
