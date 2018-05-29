package data_mappers

import (
	"github.com/maxoni/auth-iris/src/models"
	"github.com/maxoni/auth-iris/src/errors"
)

type IDataMapper interface {
	FindAll() []models.User
	FindById(id string) (*models.User, errors.Error)
	GetEntity() *models.User
	SetEntity(model *models.User)
	Insert(user *models.User)
}
