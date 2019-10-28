package repo

import (
	"fmt"
	"go-skeleton/app/entity"

	"github.com/go-pg/pg"
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
	var users = &entity.User{}
	repo.sess.Query(users, "select id, email, hashed_password from public.user")
	fmt.Println(users)
}
