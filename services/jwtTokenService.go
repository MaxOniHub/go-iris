package services

import (
	"time"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"log"
)

type JwtTokenService struct {
	secret  string
	alg     string
	expired int64
}

func NewJwtTokenService() *JwtTokenService {
	jwt := new(JwtTokenService)
	jwt.secret = "unicornsAreAwesome"
	jwt.alg = "HS256"
	jwt.expired = time.Now().Add(time.Hour * 1).Unix()

	return jwt
}

func (jwt JwtTokenService) CreateToken(uid string) (string, bool) {
	var isCreated = true

	token := jwt_lib.New(jwt_lib.GetSigningMethod(jwt.alg))

	token.Claims = jwt_lib.MapClaims{
		"Id":  uid,
		"exp": jwt.expired,
	}

	tokenString, err := token.SignedString([]byte(jwt.secret))
	if err != nil {
		isCreated = true
	}
	return tokenString, isCreated
}

func (jwt JwtTokenService) ParseToken(tokenString string) (interface{}, interface{}){
	token, err := jwt_lib.Parse(tokenString, func(token *jwt_lib.Token) (interface{}, error) {

		return []byte(jwt.secret), nil
	})

	if claims, err := token.Claims.(jwt_lib.MapClaims); err && token.Valid {
		return claims["Id"], nil
	}
	log.Println(err)
	return nil, err
}

func (jwt *JwtTokenService)GetSecret() string {
	return jwt.secret
}