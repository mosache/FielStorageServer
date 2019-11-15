package routers

import "net/http"

import "io/ioutil"

import "os"

import "fmt"

//Index index
func Index(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile("./html/index.html", os.O_RDONLY, 0666)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprint(w, string(bytes))

}
