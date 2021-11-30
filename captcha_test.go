package captcha

import (
	"testing"

	"github.com/gmvbr/captcha/proxy"
	"github.com/stretchr/testify/assert"
)

const (
	SECRET     = "0x0000000000000000000000000000000000000000"
	RESPONSE_1 = "20000000-aaaa-bbbb-cccc-000000000002"
	RESPONSE_2 = "20000000-aaaa-bbbb-cccc-000000000003"
)

func handler(secret, response string) interface{} {
	if response == "" {
		return nil
	}
	result := HCaptchaResponse{}
	if response == RESPONSE_1 {
		result.Success = true
		return result
	}
	result.Success = false
	return result
}

func TestHCaptcha(t *testing.T) {

	captcha := NewHCaptcha(SECRET)

	// Invalid URL
	captcha.UpdateService("")
	response, err := captcha.Verify(RESPONSE_1)
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
		assert.False(t, response.Success)

		// Valid user
		response, err = captcha.Verify(RESPONSE_1)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.True(t, response.Success)
	})
}
