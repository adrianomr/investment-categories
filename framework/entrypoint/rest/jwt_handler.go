package rest

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type JwtHandler interface {
	getUser(*http.Request) (int, error)
}

type JwtHandlerImpl struct{}

func NewJwtHandler() JwtHandler {
	return JwtHandlerImpl{}
}

func (JwtHandlerImpl) getUser(r *http.Request) (int, error) {
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		return 0, errors.New("userId not found")
	}
	payloadBase64 := strings.Split(authorization, ".")[1]
	payload := make([]byte, base64.StdEncoding.DecodedLen(len(payloadBase64)))
	base64.StdEncoding.Decode(payload, []byte(payloadBase64))
	var data map[string]interface{}
	json.Unmarshal(payload, &data)
	if data == nil || data["userId"] == nil {
		return 0, errors.New("userId not found")
	}
	return int(data["userId"].(float64)), nil
}
