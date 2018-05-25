package services

import (
	"github.com/maxoni/auth-iris/src/helpers"
	"database/sql"
	"github.com/maxoni/auth-iris/src/models"
	"log"
)

type SignUpService struct {
	PasswordManager *helpers.PasswordManager
	UserService     *UserService
}

func NewSignUpService(db *sql.DB) *SignUpService {
	signUp := new(SignUpService)

	signUp.PasswordManager = &helpers.PasswordManager{}
	signUp.UserService = NewUserService(db)

	return signUp
}

func (signUp *SignUpService) SignUp(data *models.SignUpModel) {
	if isValid := signUp.UserService.Validate(data); isValid == false {

	}
	signUp.generatePassword()

	signUp.UserService.Save()

}

func (signUp *SignUpService) generatePassword() {
	User := signUp.UserService.GetEntity()

	log.Println(User.Password)

	User.Password, _ = signUp.PasswordManager.HashPassword(User.Password)

	signUp.UserService.SetEntity(User)
}
