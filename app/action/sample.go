package action

import (
	"net/http"
)

type SampleAction struct {

}

var Sample SampleAction

func (self *SampleAction) SayHello(ctx *Context, w http.ResponseWriter, r *http.Request) error {
	_, err := w.Write([]byte("{'hello': 'world'}"))
	return err
}