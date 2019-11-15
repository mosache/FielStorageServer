package routers

import (
	"FileStorageServer/meta"
	"FileStorageServer/utils"
	"fmt"
	"io"
	"net/http"
	"os"
)

//FileUpload FileUpload
func FileUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		serveJSON(w, resultMap{
			"status": 0,
			"msg":    err.Error(),
		})
		return
	}

	file, fh, err := r.FormFile("file")
	if err != nil {
		serveJSON(w, resultMap{
			"status": 0,
			"msg":    err.Error(),
		})
		return
	}
	defer file.Close()

	wFile, err := os.OpenFile(fmt.Sprintf("./FileDir/%s", fh.Filename), os.O_CREATE|os.O_RDWR, 0666)
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

	fileMeta := &meta.FileMeta{FileSha1: utils.GetSha1(file), FileSize: fh.Size}

	serveJSON(w, resultMap{
		"status": 1,
		"msg":    "success",
		"data":   fileMeta,
	})
}
