package handlers

import (
	"github.com/kataras/iris"
	"database/sql"
	"github.com/maxoni/auth-iris/src/helpers"
	"github.com/maxoni/auth-iris/src/variables"
	"github.com/maxoni/auth-iris/src/models"
	"github.com/maxoni/auth-iris/src/data_mappers"
	"github.com/maxoni/auth-iris/src/services"
	"github.com/maxoni/auth-iris/src/errors"
)

func SignIn(ctx iris.Context) {

	connection := ctx.Values().Get(variables.DB).(*sql.DB)

	loginModel := &models.LoginModel{}
	passwordManager := helpers.PasswordManager{}

	mapper := data_mappers.NewUserDataMapper(connection)

	err := ctx.ReadJSON(&loginModel)

	if err != nil {
		return
	}

	User := mapper.FindByEmail(loginModel.Email)

	if passwordManager.CompareHash(User.Password, loginModel.Password) == true {
		jwtTokenService := services.NewJwtTokenService()

		token, _ := jwtTokenService.CreateToken(User.Id)

		ctx.JSON(map[string]string{
			"access_token":token,
		})
	} else {
		ctx.StatusCode(401)
		ctx.JSON(errors.ErrorHandler{}.Unauthorized())
	}

}

func SignUp(ctx iris.Context) {
	connection := ctx.Values().Get(variables.DB).(*sql.DB)

	signUpModel := models.SignUpModel{}
	SignUpService := services.NewSignUpService(connection)

	err := ctx.ReadJSON(&signUpModel)

	if err != nil {
		return
	}

	SignUpService.SignUp(&signUpModel)
}
