package integration_test

import (
	"adrianorodrigues.com.br/investment-categories/framework/entrypoint/rest"
	"adrianorodrigues.com.br/investment-categories/framework/entrypoint/rest/dto"
	integration "adrianorodrigues.com.br/investment-categories/integration_test"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func init() {
	integration.NewPrepareForTests().Prepare()
}

func TestPostCategories(t *testing.T) {
	const postJson = `{"name":"Test","grade":10,"currentAmount":5,"targetAmount":15, "investedAmount": 10}`
	req, _ := http.NewRequest("POST", "/categories", bytes.NewBufferString(postJson))
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjF9.8lSCknTnRANlJ0AVzCgO2yF838WYA7bLaAR7vAKnofo")
	response := rest.HttpServerSingleton().InitTest(req)

	log.Default().Printf("Reponse: %v", response)
	assert.Equal(t, 200, response.Code)
	category := &dto.CategoryDto{}
	responseBody := dto.ResponseDto{
		Data: category,
	}
	categoryExpected := &dto.CategoryDto{
		Name:           "Test",
		Grade:          10,
		CurrentAmount:  5,
		TargetAmount:   15,
		InvestedAmount: 10,
	}
	json.Unmarshal(response.Body.Bytes(), &responseBody)
	categoryExpected.ID = category.ID
	assert.Equal(t, categoryExpected, responseBody.Data)
}
