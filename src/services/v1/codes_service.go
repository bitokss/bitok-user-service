package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/domains/v1"
	repositories "github.com/bitokss/bitok-user-service/repositories/postgres/v1"
)

var (
	CodesService codesServiceInterface = &codesService{}
)

type codesServiceInterface interface {
	Send(code domains.CodeRequest) (rest_response.RestResp, rest_response.RestResp)
}

type codesService struct{}

func (*codesService) Send(code domains.CodeRequest) (rest_response.RestResp, rest_response.RestResp) {
	user, err := repositories.UsersRepository.FindByPhone(code.Phone)
	if user.ID != 0 {
		return nil, rest_response.NewBadRequestError(constants.UserExistsErr, nil)
	}
	if err.Status() != http.StatusNotFound {
		return nil, err
	}
	c, err := repositories.CodesRepository.Find(code.Phone, code.Type)
	if c.Model.ID != 0 {
		if c.CreatedAt.Add(time.Minute*2).Sub(time.Now()) > 0 {
			return nil, rest_response.NewBadRequestError(fmt.Sprintln(constants.WaitingMinutesErr, 2), nil)
		}
	}
	// send code (sms service)
	// add code to database
	if err != nil {
		return nil, err
	}
	return nil, nil
}
