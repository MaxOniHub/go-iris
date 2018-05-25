package data_mappers

import (
	"github.com/maxoni/auth-iris/src/models"
	"github.com/maxoni/auth-iris/src/errors"
)

type IDataMapper interface {
	FindAll() []models.IModel
	FindById(id string) (models.IModel, errors.Error)
	GetEntity() models.IModel
	SetEntity(model models.IModel)
}
