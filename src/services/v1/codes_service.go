package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/domains/v1"
	repositories "github.com/bitokss/bitok-user-service/repositories/postgres/v1"
	"github.com/labstack/echo/v4"
)

var (
	CodesService codesServiceInterface = &codesService{}
)

type codesServiceInterface interface {
	Send(code domains.CodeRequest) (rest_response.RestResp, rest_response.RestResp)
	Verify(body domains.VerifyRequest) (rest_response.RestResp, rest_response.RestResp)
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
	verifyCode := rand.Intn(99999-10000) + 10000
	body := struct {
		To    string `json:"to"`
		Code  string `json:"code"`
		Token string `json:"token"`
	}{
		To:    code.Phone,
		Code:  strconv.Itoa(verifyCode),
		Token: os.Getenv(constants.KavenegarApiKey),
	}
	j, _ := json.Marshal(body)
	resp, er := http.Post(fmt.Sprintf(os.Getenv(constants.SmsServiceHost)+"/v1/otp/"), echo.MIMEApplicationJSON, bytes.NewReader(j))
	if er != nil {
		return nil, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	if resp.StatusCode != 200 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
		}
		fmt.Println(string(bodyBytes))
		return nil, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	c, err = repositories.CodesRepository.Create(code.Phone, code.Type, verifyCode)
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(constants.SuccessSendCodeOperation, nil, http.StatusOK), nil
}

func (*codesService) Verify(body domains.VerifyRequest) (rest_response.RestResp, rest_response.RestResp) {
	code, err := repositories.CodesRepository.FindByCode(body.Phone, body.Type, body.Code)
	if err != nil {
		return nil, err
	}
	// check code is expired or not
	// expiration time is hard coded as 5 minutes
	if code.CreatedAt.Add(time.Minute*5).Sub(time.Now()) < 0 {
		return nil, rest_response.NewBadRequestError(constants.CodeIsExpiredErr, nil)
	}
	// update code and set used as true
	if _, err := repositories.CodesRepository.Update(body.Phone, body.Type, body.Code); err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(constants.SuccessSendCodeOperation, nil, http.StatusOK), nil
}
