package services

import (
	"fmt"
	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/domains/v1"
	"github.com/golang-jwt/jwt"
	"os"
)

var (
	JwtService jwtInterface = &jwtService{}
)

type jwtService struct{}

type jwtInterface interface {
	Generate(claim domains.Jwt) (string, rest_response.RestResp)
	Verify(token string) (*domains.Jwt, rest_response.RestResp)
}

func (*jwtService) Generate(claim domains.Jwt) (string, rest_response.RestResp) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claim)
	secret := os.Getenv(constants.APP_SECRET)
	sKey := []byte(secret)
	tokenString, err := token.SignedString(sKey)
	if err != nil {
		fmt.Println(err)
		return "", rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	return tokenString, nil
}

func (*jwtService) Verify(token string) (*domains.Jwt, rest_response.RestResp) {
	claim := new(domains.Jwt)
	secret := os.Getenv(constants.APP_SECRET)
	sKey := []byte(secret)
	tkn, err := jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		return sKey, nil
	})
	if err != nil {
		if err != jwt.ErrSignatureInvalid {
			return nil, rest_response.NewUnauthorizedError(constants.UnAuthorizedErr, nil)
		}
		return nil, rest_response.NewBadRequestError(constants.BadRequestErr, nil)
	}
	if !tkn.Valid {
		return nil, rest_response.NewUnauthorizedError(constants.UnAuthorizedErr, nil)
	}
	return claim, nil
}
