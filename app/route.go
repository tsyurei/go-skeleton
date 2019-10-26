package app

import (
	"github.com/gorilla/mux"

	"go-skeleton/app/action"
	"go-skeleton/app/action/api"
)

func (app *App) InitRouter(r *mux.Router) {
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Handle("/hello", app.Handler(api.Hello.SayHello)).Methods("GET")

	r.Handle("/sample", app.Handler(action.Sample.SayHello)).Methods("GET")
}