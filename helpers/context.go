package helpers

import (
	"github.com/kataras/iris"
	"github.com/maxoni/auth-iris/src/variables"
	"database/sql"
)

type Context struct {
	Ctx iris.Context
}

func (c Context)GetConnection() *sql.DB {
	return c.Ctx.Values().Get(variables.DB).(*sql.DB)
}

func (c Context)GetToken() string {
	return c.Ctx.Values().Get(variables.Token).(string)
}