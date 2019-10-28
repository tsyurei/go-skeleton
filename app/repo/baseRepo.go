package repo

import "github.com/go-pg/pg"

type BaseRepo interface {
	Inject(sess *pg.DB)
}
