package router

import "net/http"

import "FileStorageServer/utils"

import "context"

//CtxDataStruct CtxDataStruct
type CtxDataStruct struct{}

var (
	//UserKey UserKey
	ctxKey = CtxDataStruct{}
)

//TokenInterceptor TokenInterceptor
func TokenInterceptor(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var (
			tokenData *utils.TokenData
		)

		if err := r.ParseForm(); err != nil {
			serveJSON(w, resultMap{
				"status": 0,
				"msg":    "token获取失败",
			})
			return
		}

		token := r.Form.Get("token")

		if len(token) <= 0 {
			serveJSON(w, resultMap{
				"status": 0,
				"msg":    "token不能为空",
			})
			return
		}

		data, err := utils.CheckToken(token)

		if err != nil {
			serveJSON(w, resultMap{
				"status": 0,
				"msg":    "token验证不通过",
			})
			return
		}

		var ok bool

		if tokenData, ok = data.(*utils.TokenData); !ok {
			serveJSON(w, resultMap{
				"status": 0,
				"msg":    "token验证不通过",
			})
			return
		}

		if err := tokenData.Valid(); err != nil {
			serveJSON(w, resultMap{
				"status": 0,
				"msg":    "token验证不通过",
			})
			return
		}

		ctxReq := r.WithContext(context.WithValue(context.Background(), ctxKey, tokenData))

		f(w, ctxReq)
	}
}
