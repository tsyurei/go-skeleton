package repo

import "github.com/go-pg/pg"

type SampleRepoImpl struct {
	sess *pg.DB
}

var SampleRepo *SampleRepoImpl

func init() {
	SampleRepo = &SampleRepoImpl{}
}

func (repo *SampleRepoImpl) Inject(sess *pg.DB) {
	SampleRepo.sess = sess
}

func (repo *SampleRepoImpl) Test() string {
	return "test"
}
