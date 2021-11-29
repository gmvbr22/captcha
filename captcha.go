package captcha

import (
	"encoding/json"
	"net/http"
	"net/url"
)

/**
 * @see https://docs.hcaptcha.com/#verify-the-user-response-server-side
 */
type HCaptchaResponse struct {
	Success     bool     `json:"success"`
	ChallengeTS string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	Credit      bool     `json:"credit"`
	ErrorCodes  []string `json:"error-codes"`
	Score       float32  `json:"score"`
	ScoreReason []string `json:"score_reason"`
}

const (
	MISSING_INPUT_SECRET             = "missing-input-secret"
	INVALID_INPUT_SECRET             = "invalid-input-secret"
	MISSING_INPUT_RESPONSE           = "missing-input-response"
	INVALID_INPUT_RESPONSE           = "invalid-input-response"
	BAD_REQUEST                      = "bad-request"
	INVALID_OR_ALREADY_SEEN_RESPONSE = "invalid-or-already-seen-response"
	NOT_USING_DUMMY_PASSCODE         = "not-using-dummy-passcode"
	SITEKEY_SECRET_MISMATCH          = "sitekey-secret-mismatch"
)

const HCAPTCHA_URL = "https://hcaptcha.com/siteverify"

type HCaptcha struct {
	url    string
	secret string
}

func NewHCaptcha(secret string) *HCaptcha {
	return &HCaptcha{
		secret: secret,
		url:    HCAPTCHA_URL,
	}
}

// for testing, update service url
func (c *HCaptcha) UpdateService(url string) {
	c.url = url
}

func (c *HCaptcha) Verify(response string) (*HCaptchaResponse, error) {
	data := new(HCaptchaResponse)
	res, err := http.PostForm(c.url, url.Values{
		"secret":   {c.secret},
		"response": {response},
	})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
