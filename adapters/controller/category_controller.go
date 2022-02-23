package controller

import "adrianorodrigues.com.br/investment-categories/domain"

type CategoryController interface {
	CreateCategory(category *domain.Category) (*domain.Category, error)
}

type CategoryControllerImpl struct {
}

func NewCategoryController() *CategoryControllerImpl {

	return &CategoryControllerImpl{}
}

func (controller *CategoryControllerImpl) CreateCategory(category *domain.Category) (*domain.Category, error) {
	return category, nil
}
