package services

import (
	"github.com/maxoni/auth-iris/src/helpers"
	"github.com/maxoni/auth-iris/src/models"
	"github.com/jinzhu/gorm"
	"github.com/maxoni/auth-iris/src/errors"
)

type SignUpService struct {
	PasswordManager *helpers.PasswordManager
	UserService     *UserService
	ErrorHandler    errors.Error
}

func NewSignUpService(db *gorm.DB) *SignUpService {
	signUp := new(SignUpService)
	signUp.ErrorHandler = &errors.ErrorHandler{}
	signUp.PasswordManager = &helpers.PasswordManager{}
	signUp.UserService = NewUserService(db)

	return signUp
}

func (signUp *SignUpService) SignUp(data *models.SignUpModel) bool {
	if isValid := signUp.UserService.Validate(data); isValid == false {
		signUp.SetError(signUp.UserService.GetError()["message"])
		return false
	}
	signUp.generatePassword()

	if res := signUp.UserService.Save(); !res {
		signUp.SetError(signUp.UserService.GetError()["message"])
		return false
	}
	return true
}

func (signUp *SignUpService) generatePassword() {
	User := signUp.UserService.GetEntity()

	User.Password, _ = signUp.PasswordManager.HashPassword(User.Password)

	signUp.UserService.SetEntity(User)
}

func (signUp *SignUpService) SetError(error string) {
	signUp.ErrorHandler.SetError(error)
}

func (signUp SignUpService) GetError() map[string]string {
	return signUp.ErrorHandler.GetError()
}
