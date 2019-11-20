package router

import (
	"FileStorageServer/model"
	"FileStorageServer/router/dto"
	"FileStorageServer/service"
	"FileStorageServer/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

//FileUpload FileUpload
func FileUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024 * 1024)
	if err != nil {
		serveJSON(w, resultMap{
			"status": 0,
			"msg":    err.Error(),
		})
		return
	}

	token := r.FormValue("token")

	file, fh, err := r.FormFile("file")
	if err != nil {
		serveJSON(w, resultMap{
			"status": 0,
			"msg":    err.Error(),
		})
		return
	}
	defer file.Close()

	ext := path.Ext(fh.Filename)

	saveFileName := fmt.Sprintf("%d%s", time.Now().Unix(), ext)

	wFile, err := os.OpenFile(fmt.Sprintf("./FileDir/%s", saveFileName), os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		serveJSON(w, resultMap{
			"status": 0,
			"msg":    err.Error(),
		})
		return
	}

	defer wFile.Close()

	_, err = io.Copy(wFile, file)
	if err != nil {
		serveJSON(w, resultMap{
			"status": 0,
			"msg":    err.Error(),
		})
		return
	}

	td, err := utils.CheckToken(token)

	if err != nil {
		serveJSON(w, resultMap{
			"status": 0,
			"msg":    err.Error(),
		})
		return
	}

	tokenData := td.(*utils.TokenData)

	fileMeta := &model.FileMeta{FileSha1: utils.GetSha1(file), FileSize: int(fh.Size), FileName: saveFileName}

	userFile := &model.UserFile{UserID: tokenData.UserID, FileHash: fileMeta.FileSha1, FileSize: int(fh.Size), FileName: saveFileName}

	err = service.CreateFileMeta(fileMeta, userFile)

	if err != nil {
		serveJSON(w, resultMap{
			"status": 0,
			"msg":    err.Error(),
		})
		return
	}

	serveJSON(w, resultMap{
		"status": 1,
		"msg":    "success",
		"data":   dto.FileMetaDTO{FileName: saveFileName, FileHash: fileMeta.FileSha1, FileSize: int(fh.Size)},
	})
}
