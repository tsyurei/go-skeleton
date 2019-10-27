package sample

type SampleRepoImpl struct {

}

var Repo *SampleRepoImpl

func init() {
	Repo = &SampleRepoImpl {}
}

func (repo *SampleRepoImpl) Test() string {
	return "test"
}