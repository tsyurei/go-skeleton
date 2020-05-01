package service

import (
	conf "go-skeleton/config"
	"go-skeleton/internal/app/entity"
	"go-skeleton/internal/app/repo"
	"go-skeleton/internal/util"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

type authService struct {
}

var AuthService *authService

func init() {
	AuthService = &authService{}
}

func (service authService) CreateNewAuthToken(user *entity.User) (string, string, error) {
	JWTKey := conf.AppConfig.JWTKey
	RefreshJWTKey := conf.AppConfig.RefreshJWTKey
	refreshTokenExpiredTime := time.Now().Add(time.Duration(conf.AppConfig.RefreshAuthTokenExpireMinute) * time.Minute)

	claims := &JWTClaims{
		ID:   user.ID,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "go-skeleton",
			ExpiresAt: time.Now().Add(time.Duration(conf.AppConfig.AuthTokenExpireMinute) * time.Minute).Unix(),
		},
	}

	refreshClaims := &JWTClaims{
		ID:   user.ID,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "go-skeleton",
			ExpiresAt: refreshTokenExpiredTime.Unix(),
		},
	}

	JWTToken, err := JWTService.CreateJwtToken(claims, JWTKey)
	if err != nil {
		return "", "", util.WrapError(err, "Create JWT Token failed")
	}

	refreshToken, err := JWTService.CreateJwtToken(refreshClaims, RefreshJWTKey)
	if err != nil {
		return "", "", util.WrapError(err, "Create Refresh Token failed")
	}

	err = repo.RefreshTokenRepo.Save(&entity.RefreshToken{
		Token:     refreshToken,
		ExpiredAt: refreshTokenExpiredTime,
	})

	if err != nil {
		return "", "", util.WrapError(err, "Refresh token is not successfully saved to database")
	}

	return JWTToken, refreshToken, nil
}

func (service authService) ExtractJwtRequest(r *http.Request) (*JWTClaims, error) {
	JWTKey := conf.AppConfig.JWTKey

	reqToken := r.Header.Get("Authorization")
	tokens := strings.Split(reqToken, " ")
	if len(tokens) < 2 {
		return nil, util.NewUnauthorizedError().WithContext("Authorization key is not found in Header")
	}

	token := tokens[1]

	return JWTService.ExtractJwtToken(token, JWTKey)
}

func (service authService) RefreshToken(r *http.Request) (string, error) {
	JWTKey := conf.AppConfig.JWTKey
	RefreshJWTKey := conf.AppConfig.RefreshJWTKey

	reqToken := r.Header.Get("Authorization")
	tokens := strings.Split(reqToken, " ")
	if len(tokens) < 2 {
		return "", util.NewBadRequestError().WithContext("Authorization key is not found in Header")
	}

	token := tokens[1]
	oldRefreshToken, err := repo.RefreshTokenRepo.GetByToken(token)
	if err != nil {
		return "", util.NewUnauthorizedError(err).WithContext("Invalid Token")
	}

	claims, err := JWTService.ExtractJwtToken(token, RefreshJWTKey)
	if err != nil {
		go func(oldTokenId uint) {
			if er := repo.RefreshTokenRepo.Delete(oldTokenId); er != nil {
				logrus.New().Errorf("Failed to delete refresh token with id {%v}", oldTokenId)
			}
		}(oldRefreshToken.ID)
		return "", util.NewUnauthorizedError(err).WithContext("Refresh token expired")
	}

	newClaims := &JWTClaims{
		ID:   claims.ID,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "go-skeleton",
			ExpiresAt: time.Now().Add(time.Duration(conf.AppConfig.AuthTokenExpireMinute) * time.Minute).Unix(),
		},
	}

	JWTToken, err := JWTService.CreateJwtToken(newClaims, JWTKey)
	if err != nil {
		return "", err
	}

	return JWTToken, nil
}
