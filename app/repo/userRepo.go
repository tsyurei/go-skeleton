package repo

import (
	"go-skeleton/app/entity"

	"github.com/go-pg/pg/v9"
)

type UserRepoImpl struct {
	sess *pg.DB
}

var UserRepo *UserRepoImpl

func init() {
	UserRepo = &UserRepoImpl{}
}

func (repo *UserRepoImpl) Inject(sess *pg.DB) {
	repo.sess = sess
}

func (repo *UserRepoImpl) GetAll() {
	var user = entity.User{}
	repo.sess.Query(&user, "select id, email, hashed_password from public.user")
}
