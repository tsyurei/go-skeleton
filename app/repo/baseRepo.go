package repo

import "github.com/go-pg/pg/v9"

type BaseRepo interface {
	Inject(sess *pg.DB)
}
