package action

import (
	"fmt"
	"net/http"
	"go-skeleton/app/repo/sample"
)

type SampleAction struct {

}

var Sample SampleAction

func (self *SampleAction) SayHello(ctx *Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Println(sample.Repo.Test())
	_, err := w.Write([]byte("{'hello': 'world'}"))
	return err
}