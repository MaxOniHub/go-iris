package helpers

import (
	"github.com/kataras/iris"
	"github.com/maxoni/auth-iris/src/variables"
	"github.com/jinzhu/gorm"
)

type Context struct {
	Ctx iris.Context
}

func NewContext(ctx iris.Context) *Context {
	context := new(Context)
	context.Ctx = ctx

	return context
}

func (c Context)GetConnection() *gorm.DB {
	return c.Ctx.Values().Get(variables.DB).(*gorm.DB)
}

func (c Context)GetToken() string {
	return c.Ctx.Values().Get(variables.Token).(string)
}