package main

import (
	"FileStorageServer/db"
	"FileStorageServer/router"
	"fmt"
	"net/http"
)

func main() {

	sqlErr := db.InitDb()

	if sqlErr != nil {
		panic(sqlErr.Error())
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/fileUpload", router.FileUpload)
	mux.HandleFunc("/", router.TokenInterceptor(router.Index))

	mux.HandleFunc("/user/signup", router.SignUp)
	mux.HandleFunc("/user/loginin", router.LoginIn)

	//分块上传
	mux.HandleFunc("/file/mpupload/init", router.TokenInterceptor(router.MutipartUpLoadInitalization))
	mux.HandleFunc("/file/mpupload/part", router.TokenInterceptor(router.UploadPart))
	mux.HandleFunc("/file/mpupload/complete", router.TokenInterceptor(router.CompleteUpload))

	server := http.Server{Addr: ":8080", Handler: mux}

	fmt.Println("server on :8080")

	panic(server.ListenAndServe())

}
