package dto

import (
	"net/http"
	"time"
)

type ResponseDto struct {
	Code      int         `json:"-"`
	Data      interface{} `json:"data"`
	Timestamp time.Time   `json:"timestamp"`
}

func BuildResponse(data interface{}, err error) ResponseDto {
	if err != nil {
		return BuildErrorResponse(err)
	}
	return ResponseDto{Code: http.StatusOK, Data: data, Timestamp: time.Now()}
}

func BuildErrorResponse(data error) ResponseDto {
	return ResponseDto{Code: http.StatusInternalServerError, Data: data.Error(), Timestamp: time.Now()}
}

func BuildResponseForbidden(data error) ResponseDto {
	return ResponseDto{Code: http.StatusForbidden, Data: data.Error(), Timestamp: time.Now()}
}

func BuildResponseUnprocessableEntity(data error) ResponseDto {
	return ResponseDto{Code: http.StatusUnprocessableEntity, Data: data.Error(), Timestamp: time.Now()}
}
