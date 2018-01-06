package quoinex

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func (q *Quoinex) getSignature() string {
	return ""
}

func (e *Quoinex) getAuthValues() url.Values {
	nonce := strconv.FormatInt(time.Now().Unix(), 10)

	message := nonce + e.CustomerID + e.ApiKey

	mac := hmac.New(sha256.New, []byte(e.Secret))
	mac.Write([]byte(message))
	macSum := mac.Sum(nil)
	sig := strings.ToUpper(hex.EncodeToString(macSum))

	form := url.Values{
		"key":       {e.ApiKey},
		"signature": {sig},
		"nonce":     {nonce},
	}
	return form
}

func (e *Quoinex) getSignedRequest(values url.Values) *http.Request {
	panic("Unimplemented")
}
