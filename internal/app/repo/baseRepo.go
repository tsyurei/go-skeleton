package repo

import (
	"go-skeleton/internal/util"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type BaseRepo interface {
	Inject(sess *pg.DB)
	Delete(id uint) error
}

type baseRepo struct {
	sess *pg.DB
}

func (repo *baseRepo) Inject(sess *pg.DB) {
	repo.sess = sess
}

func (repo *baseRepo) Delete(id uint) error {
	return util.NewError("Delete function need to be implemented")
}

func (repo *baseRepo) createQueryWithPager(pager *util.Pager, model interface{}) *orm.Query {
	return attachPager(repo.sess.Model(model), pager)
}

func attachPager(q *orm.Query, pager *util.Pager) *orm.Query {
	return q.Offset(pager.GetOffset()).
		Limit(pager.PageSize)
}
