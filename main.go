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

	server := http.Server{Addr: ":8080", Handler: mux}

	fmt.Println("server on :8080")

	panic(server.ListenAndServe())

}
