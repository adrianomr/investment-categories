package rest

import (
	"adrianorodrigues.com.br/investment-categories/adapters/controller"
	"adrianorodrigues.com.br/investment-categories/framework/external/rest/dto"
	"net/http"
)

var categoriesSingleton = &CategoriesControllerImpl{controller: controller.NewCategoryController()}

type CategoriesController interface {
	GetCategory(w http.ResponseWriter, r *http.Request)
	PostCategory(w http.ResponseWriter, r *http.Request)
}

type CategoriesControllerImpl struct {
	controller controller.CategoryController
}

func CategoriesControllerSingleton() *CategoriesControllerImpl {
	return categoriesSingleton
}

func (rest *CategoriesControllerImpl) GetCategory(w http.ResponseWriter, r *http.Request) {
	listCategories, err := rest.controller.FindAllCategories()
	sendResponse(w, dto.BuildResponse(listCategories, err))
}

func (rest *CategoriesControllerImpl) PostCategory(w http.ResponseWriter, r *http.Request) {
	var category *dto.CategoryDto
	var err error
	readRequest(r, &category)
	category, err = rest.controller.CreateCategory(category)
	sendResponse(w, dto.BuildResponse(category, err))
}
