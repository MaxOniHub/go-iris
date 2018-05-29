package handlers

import (
	"github.com/kataras/iris"
	"github.com/maxoni/auth-iris/src/data_mappers"
	"github.com/maxoni/auth-iris/src/helpers"
	"github.com/maxoni/auth-iris/src/services"
)

func GetUsers(ctx iris.Context) {
	connection := helpers.NewContext(ctx).GetConnection()

	userDataMapper := data_mappers.NewUserDataMapper(connection)

	userDataMapper.SetLimit(ctx.URLParams())
	userDataMapper.FindAll()

	ctx.JSON(userDataMapper.FindAll())
}

func UserSingle(ctx iris.Context) {
	connection := helpers.Context{Ctx:ctx}.GetConnection()
	id := ctx.Params().Get("id")

	userService := services.NewUserService(connection)

	User, err := userService.FindById(id)

	if err != nil {
		ctx.StatusCode(404)
		ctx.JSON(userService.GetError())
	} else {
		ctx.JSON(User)
	}
}

func UserMe(ctx iris.Context) {
	context := helpers.NewContext(ctx)

	userId, _ := services.NewJwtTokenService().ParseToken(context.GetToken())
	userService := services.NewUserService(context.GetConnection())

	User, err := userService.FindById(userId.(string))

	if err != nil {
		ctx.StatusCode(404)
		ctx.JSON(userService.GetError())
	} else {
		ctx.JSON(User)
	}
}

