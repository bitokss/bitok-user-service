package services

import (
	"fmt"
	"net/http"

	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/src/constants"
	"github.com/bitokss/bitok-user-service/src/domains/v1"
	repositories "github.com/bitokss/bitok-user-service/src/repo/postgres/v1"
	"gorm.io/gorm"
)

var (
	LevelsService levelsServiceInterface = &levelsService{}
)

type levelsServiceInterface interface {
	Create(level domains.CreateLevelsRequest) (rest_response.RestResp, rest_response.RestResp)
	FindAll(limit, offset int) (rest_response.RestResp, rest_response.RestResp)
	Find(id int) (rest_response.RestResp, rest_response.RestResp)
	Delete(id int) (rest_response.RestResp, rest_response.RestResp)
	Update(id uint, level domains.CreateLevelsRequest) (rest_response.RestResp, rest_response.RestResp)
}

type levelsService struct{}

func (*levelsService) Update(id uint, level domains.CreateLevelsRequest) (rest_response.RestResp, rest_response.RestResp) {
	pd := domains.Level{
		Title: level.Title,
		Color: level.Color,
		Model: gorm.Model{
			ID: id,
		},
	}
	p, err := repositories.LevelsRepository.Update(pd)
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(fmt.Sprintf(constants.SuccessUpdateOperation, constants.Level), p, http.StatusOK), nil
}

func (*levelsService) Find(id int) (rest_response.RestResp, rest_response.RestResp) {
	p, err := repositories.LevelsRepository.Find(uint(id))
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(constants.SuccessOperation, p, http.StatusOK), nil
}

func (*levelsService) Delete(id int) (rest_response.RestResp, rest_response.RestResp) {
	err := repositories.LevelsRepository.Delete(uint(id))
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(fmt.Sprintf(constants.SuccessDeleteOperation, constants.Level), nil, http.StatusOK), nil
}

func (*levelsService) FindAll(limit, offset int) (rest_response.RestResp, rest_response.RestResp) {
	p, c, err := repositories.LevelsRepository.FindAll(limit, offset)
	if err != nil {
		return nil, err
	}
	fResp := struct {
		LevelCount int64               `json:"levelCount"`
		Levels     []domains.LevelResp `json:"levels"`
	}{
		LevelCount: c,
		Levels:     p,
	}
	return rest_response.NewSuccessResponse(constants.SuccessOperation, fResp, http.StatusOK), nil
}

func (*levelsService) Create(level domains.CreateLevelsRequest) (rest_response.RestResp, rest_response.RestResp) {
	pd := domains.Level{
		Title: level.Title,
		Color: level.Color,
		Model: gorm.Model{},
	}
	p, err := repositories.LevelsRepository.Create(pd)
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(fmt.Sprintf(constants.SuccessCreateOperation, constants.Level), p, http.StatusOK), nil
}
