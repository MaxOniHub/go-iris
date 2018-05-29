package services

import (
	"github.com/maxoni/auth-iris/src/errors"
	"github.com/jinzhu/gorm"
	"github.com/maxoni/auth-iris/src/models"
	"github.com/maxoni/auth-iris/src/helpers"
)

type SignInService struct {
	UserService     *UserService
	ErrorHandler    errors.Error
	PasswordManager *helpers.PasswordManager
	JwtTokenService *JwtTokenService
}

func NewSignInService(db *gorm.DB) *SignInService {
	signIn := new(SignInService)
	signIn.UserService = NewUserService(db)
	signIn.ErrorHandler = &errors.ErrorHandler{}
	signIn.PasswordManager = &helpers.PasswordManager{}
	signIn.JwtTokenService = NewJwtTokenService()
	return signIn
}


func (signIn *SignInService)SignIn(data *models.LoginModel) (string, bool) {
	User, err := signIn.UserService.FindByEmail(data.Email)

	if err != nil {
		signIn.SetError(err.Error())
		return "", false
	}

	if signIn.PasswordManager.CompareHash(User.Password, data.Password) == true {
		jwtTokenService := signIn.JwtTokenService

		token, _ := jwtTokenService.CreateToken(User.Id)

		return token, true
	}
	return "", false
}

func (signIn *SignInService) SetError(error string) {
	signIn.ErrorHandler.SetError(error)
}

func (signIn SignInService) GetError() map[string]string {
	return signIn.ErrorHandler.GetError()
}
