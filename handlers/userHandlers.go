package handlers

import (
	"github.com/kataras/iris"
	"github.com/maxoni/auth-iris/src/data_mappers"
	"github.com/maxoni/auth-iris/src/services"
	"github.com/maxoni/auth-iris/src/helpers"
)

func GetUsers(ctx iris.Context) {
	connection := helpers.Context{}.GetConnection()

	userDataMapper := data_mappers.NewUserDataMapper(connection)

	userDataMapper.SetLimit(ctx.URLParams())
	userDataMapper.FindAll()

	ctx.JSON(userDataMapper.FindAll())
}

func UserSingle(ctx iris.Context) {
	connection :=  helpers.Context{}.GetConnection()
	id := ctx.Params().Get("id")

	userDataMapper := data_mappers.NewUserDataMapper(connection)

	User, err := userDataMapper.FindById(id)

	if err != nil {
		ctx.StatusCode(404)
		ctx.JSON(err.NotFound())
	} else {
		ctx.JSON(User)
	}
}

func UserMe(ctx iris.Context) {
	context := helpers.Context{Ctx:ctx}
	token := context.GetToken()
	connection := context.GetConnection()

	userId,_ := services.NewJwtTokenService().ParseToken(token)
	userDataMapper := data_mappers.NewUserDataMapper(connection)

	User, err := userDataMapper.FindById(userId.(string))

	if err != nil {
		ctx.StatusCode(404)
		ctx.JSON(err.NotFound())
	} else {
		ctx.JSON(User)
	}
}

