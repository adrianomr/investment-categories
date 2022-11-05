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
	"net/http/httptest"
	"testing"
)

func TestPostCategories(t *testing.T) {
	integration.NewPrepareForTests().Prepare()
	response := postCategory(t, `{"name":"Test","grade":10,"currentAmount":5,"investedAmount": 10}`)
	category := &dto.CategoryDto{}
	responseBody := dto.ResponseDto{
		Data: category,
	}
	categoryExpected := &dto.CategoryDto{
		Name:           "Test",
		Grade:          10,
		CurrentAmount:  5,
		TargetAmount:   5,
		InvestedAmount: 10,
	}
	json.Unmarshal(response.Body.Bytes(), &responseBody)
	categoryExpected.ID = category.ID
	assert.Equal(t, categoryExpected, responseBody.Data)
}

func TestPostCategoriesWhenMultipleCategoriesShouldCalculateTargetValue(t *testing.T) {
	integration.NewPrepareForTests().Prepare()
	postCategory(t, `{"name":"Test","grade":10,"currentAmount":5,"investedAmount": 10}`)
	response := postCategory(t, `{"name":"Test","grade":5,"currentAmount":5,"investedAmount": 10}`)
	category := &dto.CategoryDto{}
	responseBody := dto.ResponseDto{
		Data: category,
	}
	categoryExpected := &dto.CategoryDto{
		Name:           "Test",
		Grade:          5,
		CurrentAmount:  5,
		TargetAmount:   3.3333335,
		InvestedAmount: 10,
	}
	json.Unmarshal(response.Body.Bytes(), &responseBody)
	categoryExpected.ID = category.ID
	assert.Equal(t, categoryExpected, responseBody.Data)
}

func postCategory(t *testing.T, postJson string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("POST", "/categories", bytes.NewBufferString(postJson))
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjJ9.QYdTqJ5SfsUkimZkmwxuMmWUx1tnY7szfN2g5UT0qDg")
	response := rest.HttpServerSingleton().InitTest(req)

	log.Default().Printf("Reponse: %v", response)
	assert.Equal(t, 200, response.Code)
	return response
}
