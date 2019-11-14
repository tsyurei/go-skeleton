package action

import (
	"fmt"
	"time"
	"github.com/sirupsen/logrus"

	"go-skeleton/config"
)

type Context struct {
	Logger        logrus.FieldLogger
	RemoteAddress string
	AppConfig     *conf.Config
}

func init() {
	fmt.Println(time.Now())
}

func (ctx *Context) WithLogger(logger logrus.FieldLogger) *Context {
	ctx.Logger = logger
	return ctx
}

func (ctx *Context) WithRemoteAddress(address string) *Context {
	ctx.RemoteAddress = address
	return ctx
}
