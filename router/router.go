package router

import "net/http"

import "fmt"

import "encoding/json"

import "io/ioutil"

import "html/template"

const (
	staticDir = "./html/"
)

type resultMap map[string]interface{}

func serveJSON(w http.ResponseWriter, resultMap resultMap) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, getJSONStr(resultMap))
}

func getJSONStr(resultMap resultMap) (result string) {
	bytes, err := json.Marshal(resultMap)

	if err != nil {
		result = "{}"
		return
	}

	result = string(bytes)

	return
}

//fileName not include file ext
func serveHTML(w http.ResponseWriter, fileName string) {
	w.Header().Add("Content-Type", "text/html")

	bytes, err := ioutil.ReadFile(staticDir + fileName + ".html")

	if err != nil {
		errorPage(w)
		return
	}

	fmt.Fprint(w, string(bytes))
}

func errorPage(w http.ResponseWriter) {
	var (
		err error
	)

	templ := template.New("templ")

	templ, err = templ.Parse(`
		<div>Test</div>
	`)

	if err != nil {
		panic(err.Error())
	}

	templ.Execute(w, templ)
}
