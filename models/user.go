package models

import (
	"github.com/maxoni/auth-iris/src/helpers"
	"gopkg.in/validator.v2"
)

type User struct {
	IModel                              `json:"-"`
	Id        string                    `json:"id"`
	Username  string                    `json:"username"`
	FirstName helpers.JsonSqlNullString `json:"first_name"`
	LastName  helpers.JsonSqlNullString `json:"last_name"`
	Email     string                    `json:"email"`
	Password  string                    `json:"-"`
}

func (u User) Validate(data interface{}) (bool, interface{}) {
	if errs := validator.Validate(data); errs != nil {
		return false, errs.Error()
	}
	return true, nil
}
