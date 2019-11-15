package routers

import "net/http"

import "fmt"

import "encoding/json"

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
