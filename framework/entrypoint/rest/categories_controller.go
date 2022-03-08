package rest

import (
	"adrianorodrigues.com.br/investment-categories/adapters/controller"
	"adrianorodrigues.com.br/investment-categories/framework/entrypoint/rest/dto"
	"net/http"
)

var categoriesSingleton CategoriesController
var jwtHandler = NewJwtHandler()

type CategoriesController interface {
	GetCategory(w http.ResponseWriter, r *http.Request)
	PostCategory(w http.ResponseWriter, r *http.Request)
}

type CategoriesControllerImpl struct {
	controller controller.CategoryController
}

func CategoriesControllerSingleton() CategoriesController {
	if categoriesSingleton == nil {
		categoriesSingleton = &CategoriesControllerImpl{controller: controller.NewCategoryController()}
	}
	return categoriesSingleton
}

func (rest *CategoriesControllerImpl) GetCategory(w http.ResponseWriter, r *http.Request) {
	userId, err := jwtHandler.getUser(r)
	if err != nil {
		sendResponse(w, dto.BuildResponseForbidden(err))
		return
	}
	listCategories, err := rest.controller.FindAllCategories(userId)
	sendResponse(w, dto.BuildResponse(listCategories, err))
}

func (rest *CategoriesControllerImpl) PostCategory(w http.ResponseWriter, r *http.Request) {
	var category *dto.CategoryDto
	userId, err := jwtHandler.getUser(r)
	if err != nil {
		sendResponse(w, dto.BuildResponseForbidden(err))
		return
	}
	readRequest(r, &category)
	category.UserId = userId
	category, err = rest.controller.CreateCategory(category)
	sendResponse(w, dto.BuildResponse(category, err))
}
