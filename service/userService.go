package service

import (
	"FileStorageServer/db"
	"FileStorageServer/model"
	"database/sql"
	"errors"
	"log"
)

//Signup 用户注册
func Signup(userName string, userPwd string) (err error) {
	var (
		stmt *sql.Stmt

		rs sql.Result

		rowAff int64
	)
	stmt, err = db.Db.Prepare("insert into user (`user_name`,`user_pwd`) values (?,?)")

	if err != nil {
		return
	}

	defer stmt.Close()

	rs, err = stmt.Exec(userName, userPwd)

	if err != nil {
		return
	}

	rowAff, err = rs.RowsAffected()

	if err != nil {
		return
	}

	if rowAff <= 0 {
		err = errors.New("插入失败")
		return
	}

	return
}

//LoginIn loginin
func LoginIn(username string, cryptoPwd string) *model.User {
	stmt, err := db.Db.Prepare("select `id`,`user_name`,`email`,`email_validated`,`phone`,`phone_validated`,`create_time`,`last_active_time`,`status`,`profile` from user where user_name = ? AND user_pwd = ?")

	if err != nil {
		log.Print(err.Error())
		return nil
	}

	defer stmt.Close()

	rs := stmt.QueryRow(username, cryptoPwd)

	if err != nil {
		log.Print(err.Error())
		return nil
	}

	user := &model.User{}

	err = rs.Scan(&user.ID, &user.UserName, &user.Email, &user.EmailValidated,
		&user.Phone, &user.PhoneValidated, &user.CreateTime, &user.LastActiveTime, &user.Status, &user.Profile)

	if err != nil {
		log.Print(err.Error())
		return nil
	}

	return user
}
