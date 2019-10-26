package api

import (
	"net/http"

	"go-skeleton/app/action"
)

type HelloAction struct {

}

var Hello HelloAction

func (self *HelloAction) SayHello(ctx *action.Context, w http.ResponseWriter, r *http.Request) error {
	_, err := w.Write([]byte("{'hello': 'world'}"))
	return err
}