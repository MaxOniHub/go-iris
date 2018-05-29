package services

import (
	"github.com/maxoni/auth-iris/src/data_mappers"
	"github.com/maxoni/auth-iris/src/models"
	"github.com/jinzhu/gorm"
	"github.com/maxoni/auth-iris/src/errors"
)

type UserService struct {
	IEntityService
	DataMapper   *data_mappers.UserDataMapper
	ErrorHandler errors.Error
}

func NewUserService(db *gorm.DB) *UserService {
	uS := new(UserService)
	uS.DataMapper = data_mappers.NewUserDataMapper(db)
	uS.ErrorHandler = &errors.ErrorHandler{}

	return uS
}

func (uS *UserService) Save() bool {
	entity := uS.GetEntity()
	success, err := uS.DataMapper.Insert(entity)

	if err != nil {
		uS.SetError(err.Error())
	}
	return success
}

func (uS *UserService) FindById(id string) (*models.User, error) {
	User, err := uS.DataMapper.FindById(id)

	if err != nil {
		uS.ErrorHandler.SetError(err.Error())
	}
	return User, err
}

func (uS *UserService) FindByEmail(email string) (*models.User, error) {
	User, err := uS.DataMapper.FindByEmail(email)

	if err != nil {
		uS.ErrorHandler.SetError(err.Error())
	}
	return User, err
}

func (uS *UserService) Validate(data *models.SignUpModel) bool {
	User := uS.GetEntity()

	success, err := User.Validate(data)
	if err != nil{
		uS.SetError(err.Error())
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
	uS.ErrorHandler.SetError(error)
}

func (uS UserService)GetError() map[string]string {
	return uS.ErrorHandler.GetError()
}