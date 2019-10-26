package app

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"go-skeleton/app/action"
	"go-skeleton/conf"
	"go-skeleton/util"

	"github.com/sirupsen/logrus"
)

type App struct {
	Config *conf.Config
}

func (a *App) NewContext() *action.Context {
	return &action.Context{
		Logger: logrus.New(),
	}
}

func New() (app *App, err error) {
	app = &App{}
	app.Config, err = conf.InitConfig()
	if err != nil {
		return nil, err
	}

	return app, err
}

type statusCodeRecorder struct {
	http.ResponseWriter
	http.Hijacker
	StatusCode int
}

func (r *statusCodeRecorder) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (app *App) Handler(f func(*action.Context, http.ResponseWriter, *http.Request) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 100*1024*1024)

		beginTime := time.Now()

		hijacker, _ := w.(http.Hijacker)

		w = &statusCodeRecorder{
			ResponseWriter: w,
			Hijacker:       hijacker,
		}

		ctx := app.NewContext().WithRemoteAddress(app.IPAddressForRequest(r))
		ctx.AppConfig = app.Config
		ctx = ctx.WithLogger(ctx.Logger.WithField("request_id", util.GenerateRandomString(32)))

		defer func() {
			statusCode := w.(*statusCodeRecorder).StatusCode
			if statusCode == 0 {
				statusCode = 200
			}
			duration := time.Since(beginTime)

			logger := ctx.Logger.WithFields(logrus.Fields{
				"duration":   duration,
				"statusCode": statusCode,
				"remote":     ctx.RemoteAddress,
			})

			logger.Info(r.Method + " " + r.URL.RequestURI())
		}()

		defer func() {
			if r := recover(); r != nil {
				ctx.Logger.Error(fmt.Errorf("%v: %s", r, debug.Stack()))
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}

		}()

		w.Header().Set("Content-Type", "application/json")

		if err := f(ctx, w, r); err != nil {
			ctx.Logger.Error(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	})
}

func (app *App) IPAddressForRequest(r *http.Request) string {
	addr := r.RemoteAddr

	return strings.Split(strings.TrimSpace(addr), ":")[0]
}
