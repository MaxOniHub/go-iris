package data_mappers

import (
	"github.com/maxoni/auth-iris/src/models"
	"strconv"
	"github.com/jinzhu/gorm"
)

type UserDataMapper struct {
	IDataMapper
	Connection   *gorm.DB
	User         *models.User
	Limit        *Limit
}

func NewUserDataMapper(db *gorm.DB) *UserDataMapper {
	mapper := new(UserDataMapper)
	mapper.Connection = db
	mapper.User = &models.User{}
	mapper.Limit = NewLimit(0, 20)
	mapper.Connection.SingularTable(true)
	return mapper
}

func (m *UserDataMapper) SetLimit(params map[string]string) {

	limit, err := strconv.ParseInt(params["limit"], 10, 64)
	offset, err := strconv.ParseInt(params["offset"], 10, 64)
	if err == nil {
		m.Limit.SetLimit(limit)
		m.Limit.SetOffset(offset)
	}
}

func (m UserDataMapper) FindAll() []models.User {

	limit := strconv.FormatInt(m.Limit.GetLimit(), 10)
	offset := strconv.FormatInt(m.Limit.GetOffset(), 10)
	var users []models.User

	m.Connection.Offset(offset).Limit(limit).Find(&users)

	return users
}

func (m UserDataMapper) FindByEmail(email string) (*models.User, error) {

	U := m.User
	if res:=m.Connection.First(&U, "email = ?", email); res.Error !=nil {
		return U, res.Error
	}
	return U, nil
}

func (m UserDataMapper) FindById(id string) (*models.User, error) {

	U := m.User
	if res := m.Connection.First(&U, id); res.Error != nil {
		return U, res.Error
	}

	return U, nil
}

func (m *UserDataMapper) Insert(user *models.User) (bool, error) {

	m.Connection.NewRecord(user)

	if result := m.Connection.Create(&user); result.Error != nil {
		return false, result.Error
	}

	return true, nil

}

func (m UserDataMapper) GetEntity() *models.User {
	return m.User
}

func (m UserDataMapper) SetEntity(entity *models.User) {
	m.User = entity
}
