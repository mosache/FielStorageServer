package service

import (
	"FileStorageServer/dao"
	"FileStorageServer/db"
	"FileStorageServer/model"
	"database/sql"
)

//CreateFileMeta CreateFileMeta
func CreateFileMeta(fm *model.FileMeta, uf *model.UserFile) (err error) {
	var (
		tx *sql.Tx
	)
	tx, err = db.Db.Begin()
	if err != nil {
		tx.Rollback()
		return
	}

	//插入文件元信息
	err = dao.InsertFileMeta(tx, fm)

	if err != nil {
		tx.Rollback()
		return
	}

	//插入用户文件表

	err = dao.InsertUserFile(tx, uf)

	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}
