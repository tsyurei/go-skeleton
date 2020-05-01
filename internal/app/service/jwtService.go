package service

import (
	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
}

var JWTService *jwtService

func init() {
	JWTService = &jwtService{}
}

type JWTClaims struct {
	ID   uint        `json:"id"`
	jwt.StandardClaims
}

func (service jwtService) CreateJwtToken(claims *JWTClaims, jwtKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (service jwtService) ExtractJwtToken(token string, jwtKey []byte) (*JWTClaims, error) {
	claims := &JWTClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !tkn.Valid {
		return nil, err
	}

	return claims, nil
}
