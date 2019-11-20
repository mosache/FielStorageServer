package dao

import (
	"FileStorageServer/model"
	"database/sql"
	"errors"
)

const (
	_userFileTableName = "user_file"
)

//InsertUserFile InsertUserFile
func InsertUserFile(tx *sql.Tx, data *model.UserFile) (err error) {
	var (
		stmt *sql.Stmt
		rs   sql.Result
		ra   int64
	)
	stmt, err = tx.Prepare(
		"INSERT INTO" + _userFileTableName +
			"(`UserID`,`FileHash`,`FileSize`,`FileName`,)" +
			"VALUES" +
			"(?,?,?,?)")
	if err != nil {
		return
	}

	rs, err = stmt.Exec(data.UserID, data.FileHash, data.FileSize, data.FileName)
	if err != nil {
		return
	}

	ra, err = rs.RowsAffected()
	if err != nil {
		return
	}
	if ra <= 0 {
		return errors.New("insert error")
	}
	return
}
