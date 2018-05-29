package data_mappers

import (
	"log"
	"database/sql"
	"github.com/maxoni/auth-iris/src/errors"
	"github.com/maxoni/auth-iris/src/models"
	"strconv"
)

type UserDataMapper struct {
	IDataMapper
	Connection   *sql.DB
	User         *models.User
	Limit        *Limit
	ErrorHandler errors.Error
}

func NewUserDataMapper(db *sql.DB) *UserDataMapper {
	mapper := new(UserDataMapper)
	mapper.Connection = db
	mapper.User = &models.User{}
	mapper.Limit = NewLimit(0, 20)
	mapper.ErrorHandler = &errors.ErrorHandler{}
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

	// Execute the query
	results, err := m.Connection.Query("SELECT id, username, email FROM User LIMIT " + limit + " OFFSET " + offset)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var users []models.User

	for results.Next() {
		var u models.User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&u.Id, &u.Username, &u.Email)
		if err != nil {
			log.Printf(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		users = append(users, u)
	}
	return users
}

func (m UserDataMapper) FindByEmail(email string) *models.User {

	U := m.User

	err := m.Connection.QueryRow("SELECT * FROM User where email='" + email + "'").Scan(&U.Id, &U.Username, &U.Email, &U.FirstName, &U.LastName, &U.Password)

	if err != nil {
		log.Println(err)
	}
	return U
}

func (m UserDataMapper) FindById(id string) (*models.User, errors.Error) {

	U := m.User

	err := m.Connection.QueryRow("SELECT * FROM User where id=?", id).Scan(&U.Id, &U.Username, &U.Email, &U.FirstName, &U.LastName, &U.Password)

	log.Println(err)
	if err != nil {
		return U, m.ErrorHandler
	}

	return U, nil
}

func (m *UserDataMapper) Insert(user *models.User) {

	u := user

	stmt, err := m.Connection.Prepare("INSERT INTO User (username, email, first_name, last_name, password) values(?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(u.Username, u.Email, u.FirstName.String, u.LastName.String, u.Password)

	if err != nil {
		panic(err)
	}
	_ = res
}

func (m UserDataMapper) GetEntity() *models.User {
	return m.User
}

func (m UserDataMapper) SetEntity(entity *models.User) {
	m.User = entity
}
