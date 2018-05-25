package handlers

import (
	"github.com/kataras/iris"
	"github.com/maxoni/auth-iris/src/variables"
	"github.com/maxoni/auth-iris/src/data_mappers"
	"database/sql"
	"log"
)

func GetUsers(ctx iris.Context) {

	connection := ctx.Values().Get(variables.DB).(*sql.DB)

	userDataMapper := data_mappers.NewUserDataMapper(connection)
	userDataMapper.FindAll()

	ctx.JSON(userDataMapper.FindAll())
}

func UserSingle(ctx iris.Context) {
	connection := ctx.Values().Get(variables.DB).(*sql.DB)

	id := ctx.Params().Get("id")
	log.Println(id)

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
	/*token := ctx.Values().Get(variables.Token).(string)
	tokenService := &services.JwtTokenService{}

	userId,_ := tokenService.ParseToken(token)

	log.Println(userId)*/

}


