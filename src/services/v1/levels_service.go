package services

import (
	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/domains/v1"
)

var (
	LevelsService levelsServiceInterface = &levelsService{}
)

type levelsServiceInterface interface {
	Create(level domains.CreateLevelsRequest) (rest_response.RestResp, rest_response.RestResp)
	FindAll(limit, offset int) (rest_response.RestResp, rest_response.RestResp)
	Find(lid int) (rest_response.RestResp, rest_response.RestResp)
	Delete(lid int) (rest_response.RestResp, rest_response.RestResp)
	Update(lid int, level domains.CreateLevelsRequest) (rest_response.RestResp, rest_response.RestResp)
}

type levelsService struct{}

func (l *levelsService) Update(lid int, level domains.CreateLevelsRequest) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (l *levelsService) Find(lid int) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (l *levelsService) Delete(lid int) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (l *levelsService) FindAll(limit, offset int) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}

func (l *levelsService) Create(level domains.CreateLevelsRequest) (rest_response.RestResp, rest_response.RestResp) {
	return nil, nil
}
