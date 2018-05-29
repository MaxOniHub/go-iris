package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type DbHandler struct {
	Db *gorm.DB
}

func NewDbHandler() *DbHandler  {
	dbHandler := new(DbHandler)
	dbHandler.Db = InitDB()
	return dbHandler
}

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/gobase?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Println(err)
	}
	return db
}

func (repoHandler DbHandler) GetDb() *gorm.DB {
	return repoHandler.Db
}
