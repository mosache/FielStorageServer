package dao

import (
	"FileStorageServer/model"
	"database/sql"
	"errors"
)

const (
	_fileMetaTableName = "file_meta"
)

//InsertFileMeta InsertFileMeta
func InsertFileMeta(db *sql.Tx, data *model.FileMeta) (err error) {
	var (
		stmt         *sql.Stmt
		rs           sql.Result
		rowsAffected int64
	)
	stmt, err = db.Prepare(
		"INSERT INTO " + _fileMetaTableName +
			"`FileSha1`,`FileSize`,`FileName`" +
			"VALUES (?,?,?)")
	if err != nil {
		return
	}
	rs, err = stmt.Exec(data.FileSha1, data.FileSize, data.FileName)
	if err != nil {
		return
	}

	rowsAffected, err = rs.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected <= 0 {
		return errors.New("insert error")
	}

	return
}
