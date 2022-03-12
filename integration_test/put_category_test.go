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

func TestPutCategories(t *testing.T) {
	integration.NewPrepareForTests().Prepare()
	req, _ := http.NewRequest("PUT", "/categories/TEST", bytes.NewBufferString(`{"name":"Test","grade":10,"currentAmount":5,"targetAmount":15}`))
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjF9.8lSCknTnRANlJ0AVzCgO2yF838WYA7bLaAR7vAKnofo")
	response := rest.HttpServerSingleton().InitTest(req)

	log.Default().Printf("Reponse: %v", response)
	assert.Equal(t, 200, response.Code)
	category := &dto.CategoryDto{}
	responseBody := dto.ResponseDto{
		Data: category,
	}
	categoryExpected := &dto.CategoryDto{
		ID:            "TEST",
		Name:          "Test",
		Grade:         10,
		CurrentAmount: 5,
		TargetAmount:  15,
	}
	json.Unmarshal(response.Body.Bytes(), &responseBody)
	assert.Equal(t, categoryExpected, responseBody.Data)
}

func TestPutCategoriesShouldUpdateInvestments(t *testing.T) {
	integration.NewPrepareForTests().Prepare()
	const categoryWithInvestmentsJson = `{"id":"TEST","name":"Test","grade":10,"currentAmount":5,"targetAmount":15, "investments": [{"id": "INVESTMENT-TEST", "name": "INVESTMENT 1", "origin": "B3", "currentAmount": 1000}]}`
	req, _ := http.NewRequest("PUT", "/categories/TEST", bytes.NewBufferString(categoryWithInvestmentsJson))
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjF9.8lSCknTnRANlJ0AVzCgO2yF838WYA7bLaAR7vAKnofo")
	response := rest.HttpServerSingleton().InitTest(req)

	log.Default().Printf("Reponse: %v", response)
	assert.Equal(t, 200, response.Code)
	category := &dto.CategoryDto{}
	responseBody := dto.ResponseDto{
		Data: category,
	}
	investmentsExpected := []*dto.InvestmentDto{{
		ID:            "INVESTMENT-TEST",
		Name:          "INVESTMENT 1",
		Grade:         0,
		Origin:        "B3",
		CurrentAmount: 1000,
		TargetAmount:  0,
	}}

	json.Unmarshal(response.Body.Bytes(), &responseBody)
	assert.Equal(t, investmentsExpected, category.Investments)
}
