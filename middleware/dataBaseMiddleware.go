package middleware

import (
	"github.com/maxoni/auth-iris/src/db"
	"github.com/kataras/iris"
	"github.com/maxoni/auth-iris/src/variables"
)

func DataBaseMiddleware(ctx iris.Context) {
	repoHandler := db.NewDbHandler()

	ctx.Values().Set(variables.DB, repoHandler.GetDb())
	ctx.Next() // execute the next handler, in this case the main one.
}
