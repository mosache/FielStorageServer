package router

import (
	"FileStorageServer/model"
	"FileStorageServer/service"
	"fmt"
	"io/ioutil"
	"net/http"
)

//SignUp SignUp
func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		bytes, err := ioutil.ReadFile("./html/signup.html")
		if err != nil {
			serveJSON(w, resultMap{"err": err.Error()})
		}

		fmt.Fprint(w, string(bytes))
		return
	}

	err := r.ParseForm()

	if err != nil {
		serveJSON(w, resultMap{"err": err.Error()})
		return
	}

	userName := r.FormValue("username")
	pwd := r.FormValue("password")

	if len(pwd) < 6 {
		serveJSON(w, resultMap{"err": "密码过于简单"})
		return
	}

	err = service.Signup(userName, pwd)

	if err != nil {
		serveJSON(w, resultMap{"err": err.Error()})
		return
	}

	serveJSON(w, resultMap{
		"success": "1",
	})
}

//LoginIn LoginIn
func LoginIn(w http.ResponseWriter, r *http.Request) {
	var (
		user *model.User
	)

	if err := r.ParseForm(); err != nil {
		serveJSON(w, resultMap{"err": err.Error()})
		return
	}

	username := r.FormValue("username")
	pwd := r.FormValue("password")

	if user = service.LoginIn(username, pwd); user == nil {
		serveJSON(w, resultMap{"err": "登录失败"})
		return
	}

	serveJSON(w, resultMap{
		"status": 1,
		"data":   user,
	})

}
