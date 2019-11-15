package main

import (
	"FileStorageServer/db"
	"FileStorageServer/routers"
	"fmt"
	"net/http"
)

func main() {

	sqlErr := db.InitDb()

	if sqlErr != nil {
		panic(sqlErr.Error())
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/fileUpload", routers.FileUpload)
	mux.HandleFunc("/", routers.Index)

	server := http.Server{Addr: ":8080", Handler: mux}

	fmt.Println("server on :8080")

	panic(server.ListenAndServe())

}
