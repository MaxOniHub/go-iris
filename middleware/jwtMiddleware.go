package middleware

import (
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/maxoni/auth-iris/src/services"
)

func JwtMiddleware() *jwtmiddleware.Middleware {
	jwtService := services.NewJwtTokenService()
	secret := jwtService.GetSecret()

	return jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}

