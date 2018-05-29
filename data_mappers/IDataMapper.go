package data_mappers

import (
	"github.com/maxoni/auth-iris/src/models"
)

type IDataMapper interface {
	FindAll() []models.User
	FindById(id string) (*models.User, error)
	GetEntity() *models.User
	SetEntity(model *models.User)
	Insert(user *models.User)
}
