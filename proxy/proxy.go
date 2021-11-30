package proxy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

type Handler = func(url string)
type Response = func(secret, response string) interface{}

func Proxy(response Response, handler Handler) {
	server := httptest.NewServer(http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[Captcha] [Proxy] handle request (response=%s)\n", r.FormValue("response"))
		r.ParseForm()
		entity := response(
			r.FormValue("secret"),
			r.FormValue("response"),
		)
		if entity != nil {
			json.NewEncoder(w).Encode(entity)
		}
	}))
	handler(server.URL)
	server.Close()
}
