package middleware

import (
	"strings"
	"github.com/kataras/iris"
	"github.com/maxoni/auth-iris/src/variables"
)


func AuthHeaderReader(ctx iris.Context) {
	ctx.Values().Set(variables.Token, getValue(	ctx.GetHeader("Authorization"), "Bearer"))
	ctx.Next() // execute the next handler, in this case the main one.
}


func getValue(header string, splitBy string) string {
	value := strings.Split(header, splitBy)
	return strings.TrimSpace(value[1])
}
