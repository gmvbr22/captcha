package hcaptcha_proxy

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

type Exec = func(url string)

type Response = func(secret, response string) interface{}

func handler(h Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		entity := h(
			r.FormValue("secret"),
			r.FormValue("response"),
		)
		if entity != nil {
			json.NewEncoder(w).Encode(entity)
		}
	}
}

func Proxy(response Response, exec Exec) {
	handler := handler(response)
	srv := httptest.NewServer(http.HandlerFunc(handler))
	exec(srv.URL)
	srv.Close()
}
