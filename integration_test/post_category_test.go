package integration_test

import (
	"adrianorodrigues.com.br/investment-categories/framework/external/rest"
	"adrianorodrigues.com.br/investment-categories/framework/external/rest/dto"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func TestPostCategories(t *testing.T) {
	req, _ := http.NewRequest("POST", "/categories", bytes.NewBufferString(`{"id":"","name":"","grade":0,"currentAmount":0,"TargetAmount":0,"category":null}`))
	response := rest.HttpServerSingleton().InitTest(req)

	log.Default().Printf("Reponse: %v", response)
	assert.Equal(t, 200, response.Code)

	responseBody := dto.ResponseDto{
		Data: &dto.CategoryDto{},
	}
	categoryExpected := &dto.CategoryDto{}
	log.Default().Printf(response.Body.String())
	json.Unmarshal(response.Body.Bytes(), &responseBody)
	log.Default().Printf(responseBody.Timestamp.String())
	assert.Equal(t, categoryExpected, responseBody.Data)
}
