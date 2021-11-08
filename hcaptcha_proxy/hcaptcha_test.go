package hcaptcha_proxy

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Ok bool
}

func handleResponse(secret, response string) interface{} {
	if response == "false" {
		return nil
	}
	return &Result{Ok: true}
}

func TestHCaptcha(t *testing.T) {
	Proxy(handleResponse, func(sUrl string) {
		result := new(Result)
		values := url.Values{
			"secret":   {""},
			"response": {"false"},
		}
		res, err := http.PostForm(sUrl, values)
		assert.NoError(t, err)
		err = json.NewDecoder(res.Body).Decode(result)
		res.Body.Close()
		assert.Error(t, err)

		values = url.Values{
			"secret":   {""},
			"response": {"true"},
		}
		res, err = http.PostForm(sUrl, values)
		assert.NoError(t, err)
		err = json.NewDecoder(res.Body).Decode(result)
		res.Body.Close()
		assert.NoError(t, err)
		assert.Equal(t, true, result.Ok)
	})
}
