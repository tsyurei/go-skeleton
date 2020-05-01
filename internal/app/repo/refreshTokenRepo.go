package repo

import (
	"go-skeleton/internal/app/entity"
	"time"
)

type refreshTokenRepoIml struct {
	baseRepo
}

var RefreshTokenRepo *refreshTokenRepoIml

func init() {
	RefreshTokenRepo = &refreshTokenRepoIml{}
}

func (repo *refreshTokenRepoIml) Save(r *entity.RefreshToken) error {

	return repo.sess.Insert(r)
}

func (repo *refreshTokenRepoIml) GetExpiredTokens() ([]*entity.RefreshToken, error) {
	var RefreshTokens []*entity.RefreshToken
	err := repo.sess.Model(RefreshTokens).
		Where("expired_at < ?", time.Now()).
		Select()

	return RefreshTokens, err
}

func (repo *refreshTokenRepoIml) GetByToken(token string) (*entity.RefreshToken, error) {
	RefreshToken := &entity.RefreshToken{}
	err := repo.sess.Model(RefreshToken).
		Where("token = ?", token).
		First()

	if err != nil {
		return nil, err
	}

	return RefreshToken, nil
}

func (repo *refreshTokenRepoIml) Delete(id uint) error {
	RefreshToken := &entity.RefreshToken{ID: id}
	return repo.sess.Delete(RefreshToken)
}
