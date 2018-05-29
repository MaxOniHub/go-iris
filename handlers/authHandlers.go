package handlers

import (
	"github.com/kataras/iris"
	"github.com/maxoni/auth-iris/src/helpers"
	"github.com/maxoni/auth-iris/src/models"
	"github.com/maxoni/auth-iris/src/services"
)

func SignIn(ctx iris.Context) {
	connection := helpers.NewContext(ctx).GetConnection()
	loginModel := &models.LoginModel{}

	err := ctx.ReadJSON(&loginModel)

	if err != nil {
		return
	}

	signInService := services.NewSignInService(connection)
	token, success := signInService.SignIn(loginModel)

	if success {
		ctx.JSON(map[string]string{
			"access_token": token,
		})
	} else {
		ctx.StatusCode(401)
		ctx.JSON(signInService.GetError())
	}
}

func SignUp(ctx iris.Context) {
	connection := helpers.NewContext(ctx).GetConnection()

	signUpModel := &models.SignUpModel{}
	SignUpService := services.NewSignUpService(connection)

	err := ctx.ReadJSON(&signUpModel)

	if err != nil {
		return
	}

	if res := SignUpService.SignUp(signUpModel); res {
		ctx.JSON(map[string]string{
			"message": "Success",
		})
	} else {
		ctx.StatusCode(400)
		ctx.JSON(SignUpService.GetError())
	}
}
