package captcha

import (
	"testing"

	"github.com/gmvbr/captcha/proxy"
	"github.com/stretchr/testify/assert"
)

func handler(secret, response string) interface{} {
	if response == "" {
		return nil
	}
	if response == "20000000-aaaa-bbbb-cccc-000000000002" {
		result := HCaptchaResponse{}
		result.Success = true
		return result
	}
	result := HCaptchaResponse{}
	result.Success = false
	return result
}

func TestHCaptcha(t *testing.T) {

	captcha := NewHCaptcha("0x0000000000000000000000000000000000000000")

	// Invalid URL
	captcha.UpdateService("#INVALID_URL")
	response, err := captcha.Verify("20000000-aaaa-bbbb-cccc-000000000002")
	assert.Error(t, err)
	assert.Nil(t, response)

	proxy.Proxy(handler, func(url string) {
		captcha.UpdateService(url)

		// Invalid body
		response, err = captcha.Verify("")
		assert.Error(t, err)
		assert.Nil(t, response)

		// Invalid user
		response, err := captcha.Verify("20000000-aaaa-bbbb-cccc-000000000003")
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, false, response.Success)

		// Valid user
		response, err = captcha.Verify("20000000-aaaa-bbbb-cccc-000000000002")
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, true, response.Success)
	})
}
