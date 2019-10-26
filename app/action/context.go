package action

import (
	"github.com/sirupsen/logrus"

	"go-skeleton/conf"
)

type Context struct {
	Logger        logrus.FieldLogger
	RemoteAddress string
	AppConfig     *conf.Config
}

func (ctx *Context) WithLogger(logger logrus.FieldLogger) *Context {
	ret := *ctx
	ret.Logger = logger
	return &ret
}

func (ctx *Context) WithRemoteAddress(address string) *Context {
	ret := *ctx
	ret.RemoteAddress = address
	return &ret
}
