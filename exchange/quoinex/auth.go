package quoinex

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type header struct {
	Type string `json:"typ"`
	Algo string `json:"alg"`
}

type payload struct {
	Path    string `json:"path"`
	Nonce   int64  `json:"nonce"`
	TokenID string `json:"token_id"`
}

func (q *Quoinex) getSignature(path string) string {
	tokenId := q.ApiKey
	secret := q.Secret

	header := header{
		Algo: "HS256",
		Type: "JWT",
	}

	authPayload := payload{
		Path:    path,
		Nonce:   time.Now().Unix(),
		TokenID: tokenId,
	}

	jsonHeader, err := json.Marshal(header)
	if err != nil {
		panic(err)
	}
	jsonPayload, err := json.Marshal(authPayload)
	if err != nil {
		panic(err)
	}

	// signature := JWT.encode(authPayload, userSecret, "HS256")

	encodedHeader := base64.StdEncoding.EncodeToString(jsonHeader)
	encodedPayload := base64.StdEncoding.EncodeToString(jsonPayload)

	// ----- HMAC256
	message := encodedHeader + "." + encodedPayload
	// userSecret

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	macSum := mac.Sum(nil)

	signature := base64.StdEncoding.EncodeToString(macSum)

	return encodedHeader + "." + encodedPayload + "." + signature
}

////////////////////////////////////////////
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
