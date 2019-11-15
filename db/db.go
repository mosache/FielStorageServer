package db

import (
	"database/sql"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	//Db db
	Db *sql.DB
)

//InitDb InitDb
func InitDb() (err error) {
	Db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3308)/FileStorage?parseTime=true&charset=utf8")

	if err != nil {
		return
	}

	err = Db.Ping()

	if err != nil {
		return
	}
	return
}
