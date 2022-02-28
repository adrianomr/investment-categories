package integration_test

import (
	"adrianorodrigues.com.br/investment-categories/framework/external/rest"
	"adrianorodrigues.com.br/investment-categories/framework/external/rest/dto"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func TestGetCategories(t *testing.T) {
	req, _ := http.NewRequest("GET", "/categories", nil)
	response := rest.HttpServerSingleton().InitTest(req)

	log.Default().Printf("Reponse: %v", response)
	assert.Equal(t, 200, response.Code)

	categories := make([]dto.CategoryDto, 0, 10)
	responseBody := dto.ResponseDto{
		Data: &categories,
	}
	categoriesExpected := []dto.CategoryDto{dto.CategoryDto{}, dto.CategoryDto{}}
	log.Default().Printf(response.Body.String())
	json.Unmarshal(response.Body.Bytes(), &responseBody)
	log.Default().Printf(responseBody.Timestamp.String())
	assert.Equal(t, categoriesExpected, categories)
}
