package dto

import "time"

type ResponseDto struct {
	Code      int `json:"-"`
	Data      interface{}
	Timestamp time.Time
}

func BuildResponse(data interface{}, err error) ResponseDto {
	if err != nil {
		return BuildErrorResponse(err)
	}
	return ResponseDto{Code: 200, Data: data, Timestamp: time.Now()}
}

func BuildErrorResponse(data error) ResponseDto {
	return ResponseDto{Code: 500, Data: data, Timestamp: time.Now()}
}
