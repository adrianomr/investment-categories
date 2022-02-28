package controller

import (
	"adrianorodrigues.com.br/investment-categories/framework/external/rest/dto"
)

type CategoryController interface {
	CreateCategory(category *dto.CategoryDto) (*dto.CategoryDto, error)
	FindAllCategories() (*[]dto.CategoryDto, error)
}

type CategoryControllerImpl struct {
}

func NewCategoryController() *CategoryControllerImpl {

	return &CategoryControllerImpl{}
}

func (controller *CategoryControllerImpl) CreateCategory(category *dto.CategoryDto) (*dto.CategoryDto, error) {
	return category, nil
}

func (controller *CategoryControllerImpl) FindAllCategories() (*[]dto.CategoryDto, error) {
	listCategories := []dto.CategoryDto{dto.CategoryDto{
		ID:            "",
		Name:          "",
		Grade:         0,
		CurrentAmount: 0,
		TargetAmount:  0,
		Category:      nil,
	}, dto.CategoryDto{
		ID:            "",
		Name:          "",
		Grade:         0,
		CurrentAmount: 0,
		TargetAmount:  0,
		Category:      nil,
	}}
	return &listCategories, nil
}
