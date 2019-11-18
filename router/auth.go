package router

import "net/http"

//TokenInterceptor TokenInterceptor
func TokenInterceptor(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		token := r.PostFormValue("token")

		if len(token) <= 0 {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		f(w, r)
	}
}
