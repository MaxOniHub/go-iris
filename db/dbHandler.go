package db

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


type DbHandler struct {
	Db *sql.DB
}

func NewDbHandler() *DbHandler  {
	dbHandler := new(DbHandler)
	dbHandler.Db = InitDB()
	return  dbHandler
}

func InitDB() *sql.DB {
	Db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/gobase")
	if err != nil {
		log.Fatal(err)
	}
	return Db
}

func (repoHandler DbHandler) GetDb() *sql.DB {
	return repoHandler.Db
}
