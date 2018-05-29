package main

import (
	"github.com/kataras/iris"
	"github.com/maxoni/auth-iris/src/handlers"
	"github.com/kataras/iris/middleware/logger"

	"github.com/maxoni/auth-iris/src/middleware"

)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	jwtHandler := middleware.JwtMiddleware()
	//app.Use(recover.New())
	app.Use(logger.New())
	app.Use(middleware.DataBaseMiddleware)

	v1 := app.Party("/api/v1")

	users := v1.Party("/users")
	users.Use(jwtHandler.Serve)
	users.Get("/", handlers.GetUsers)
	users.Get("/{id:int}",handlers.UserSingle)
	users.Get("/me", middleware.AuthHeaderReader, handlers.UserMe)

	auth := v1.Party("/auth")
	auth.Post("/sign-in", handlers.SignIn)
	auth.Post("/sign-up", handlers.SignUp)

	app.Run(iris.Addr(":8080"))
}