package rest

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
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
	payload, err := base64.RawStdEncoding.DecodeString(payloadBase64)
	if err != nil {
		log.Print(err)
		return 0, errors.New("invalid payload")
	}
	var data map[string]interface{}
	json.Unmarshal(payload, &data)
	if data == nil || data["userId"] == nil {
		return 0, errors.New("userId not found")
	}
	return int(data["userId"].(float64)), nil
}
