package integration_test

import (
	"adrianorodrigues.com.br/investment-categories/framework/entrypoint/rest"
	"adrianorodrigues.com.br/investment-categories/framework/entrypoint/rest/dto"
	integration "adrianorodrigues.com.br/investment-categories/integration_test"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func init() {
	integration.NewPrepareForTests().Prepare()
}

func TestGetCategoriesShouldReturnCategories(t *testing.T) {

	req, _ := http.NewRequest("GET", "/categories", nil)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjF9.8lSCknTnRANlJ0AVzCgO2yF838WYA7bLaAR7vAKnofo")
	response := rest.HttpServerSingleton().InitTest(req)

	log.Default().Printf("Reponse: %v", response)
	assert.Equal(t, 200, response.Code)

	categories := make([]dto.CategoryDto, 0, 10)
	responseBody := dto.ResponseDto{
		Data: &categories,
	}
	categoriesExpected := &[]dto.CategoryDto{*integration.CategoryRest, *integration.CategoryRest2}
	log.Default().Printf(response.Body.String())
	json.Unmarshal(response.Body.Bytes(), &responseBody)
	log.Default().Printf(responseBody.Timestamp.String())
	assert.Equal(t, categoriesExpected, responseBody.Data)
}

func TestGetCategoriesWhenNoCategoryFoundShouldReturnEmptyList(t *testing.T) {

	req, _ := http.NewRequest("GET", "/categories", nil)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjE2NzQ2Nn0.8xd5Hppgg-U4HCflOi_jXOdrZ-o9EPSIdh34keTQYDw")
	response := rest.HttpServerSingleton().InitTest(req)

	log.Default().Printf("Reponse: %v", response)
	assert.Equal(t, 200, response.Code)

	categories := make([]dto.CategoryDto, 0, 10)
	responseBody := dto.ResponseDto{
		Data: &categories,
	}
	categoriesExpected := &[]dto.CategoryDto{}
	log.Default().Printf(response.Body.String())
	json.Unmarshal(response.Body.Bytes(), &responseBody)
	log.Default().Printf(responseBody.Timestamp.String())
	assert.Equal(t, categoriesExpected, responseBody.Data)
}
