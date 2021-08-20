package services

import (
	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/domains/v1"
)

type jwtService struct{}

type jwtInterface interface {
	Generate(claim domains.Jwt) (string, rest_response.RestResp)
	Verify(token string) (domains.Jwt, rest_response.RestResp)
}

var (
	JwtService jwtInterface = &jwtService{}
)

func (*jwtService) Generate(claim domains.Jwt) (string, rest_response.RestResp) {
	return "", nil
}

func (*jwtService) Verify(token string) (domains.Jwt, rest_response.RestResp) {
	return domains.Jwt{}, nil
}
