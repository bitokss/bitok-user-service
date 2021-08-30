package domains

import "github.com/golang-jwt/jwt"

type Jwt struct {
	UID int `json:"uid"`
	jwt.StandardClaims
}
