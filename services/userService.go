package services

import (
	"github.com/maxoni/auth-iris/src/data_mappers"
	"database/sql"
	"github.com/maxoni/auth-iris/src/models"
)

type UserService struct {
	IEntityService
	DataMapper *data_mappers.UserDataMapper
	error string
}

func NewUserService(db *sql.DB) *UserService {
	uS := new(UserService)
	uS.DataMapper = data_mappers.NewUserDataMapper(db)
	return uS
}

func (uS *UserService)Save() {
	uS.DataMapper.Insert(uS.GetEntity())
}

func (uS *UserService) Validate(data *models.SignUpModel) bool {
	User := uS.GetEntity()

	success, err := User.Validate(data)
	if err != nil{
		uS.SetError(err.(string))
	}
	uS.Load(data)
	return success
}

func (uS *UserService) Load(data *models.SignUpModel) {
	User := uS.GetEntity()

	User.LastName.String = data.LastName
	User.FirstName.String = data.FirstName
	User.Email = data.Email
	User.Username = data.Username
	User.Password = data.Password

	uS.DataMapper.SetEntity(User)
}

func (uS *UserService) GetEntity() *models.User {
	return uS.DataMapper.GetEntity()
}

func (uS *UserService) SetEntity(entity *models.User) {
	uS.DataMapper.SetEntity(entity)
}

func (uS *UserService)SetError(error string) {
	uS.error = error
}

func (uS UserService)GetError() string {
	return uS.error
}