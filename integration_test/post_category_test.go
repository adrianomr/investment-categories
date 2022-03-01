package integration_test

import (
	"adrianorodrigues.com.br/investment-categories/framework/entrypoint/rest"
	"adrianorodrigues.com.br/investment-categories/framework/entrypoint/rest/dto"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func TestPostCategories(t *testing.T) {
	req, _ := http.NewRequest("POST", "/categories", bytes.NewBufferString(`{"name":"Test","grade":10,"currentAmount":5,"TargetAmount":15}`))
	response := rest.HttpServerSingleton().InitTest(req)

	log.Default().Printf("Reponse: %v", response)
	assert.Equal(t, 200, response.Code)
	category := &dto.CategoryDto{}
	responseBody := dto.ResponseDto{
		Data: category,
	}
	categoryExpected := &dto.CategoryDto{
		Name:          "Test",
		Grade:         10,
		CurrentAmount: 5,
		TargetAmount:  15,
	}
	json.Unmarshal(response.Body.Bytes(), &responseBody)
	categoryExpected.ID = category.ID
	assert.Equal(t, categoryExpected, responseBody.Data)
}
