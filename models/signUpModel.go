package models

type SignUpModel struct {
	Email     string `json:"email"`
	Username  string `json:"username" validate:"nonzero"`
	FirstName string `json:"first_name" validate:"nonzero"`
	LastName  string `json:"last_name" validate:"nonzero"`
	Password  string `json:"password"`
}
