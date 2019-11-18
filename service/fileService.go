package service

import (
	"FileStorageServer/db"
	"FileStorageServer/model"
	"errors"
)

//CreateFileMeta CreateFileMeta
func CreateFileMeta(fm *model.FileMetaModel) (err error) {
	stmt, err := db.Db.Prepare("insert into file_meta (FileSha1,FileSize) values (?,?)")
	if err != nil {
		return
	}
	defer stmt.Close()

	re, err := stmt.Exec(fm.FileSha1, fm.FileSize)

	if err != nil {
		return
	}

	if rows, mErr := re.RowsAffected(); mErr == nil {
		if rows <= 0 {
			err = errors.New("insert into db affected zero row")
			return
		}
	}

	return
}
