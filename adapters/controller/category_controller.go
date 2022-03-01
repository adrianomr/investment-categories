package controller

import (
	"adrianorodrigues.com.br/investment-categories/domain"
	"adrianorodrigues.com.br/investment-categories/framework/entrypoint/rest/dto"
	"adrianorodrigues.com.br/investment-categories/usecase"
)

type CategoryController interface {
	CreateCategory(category *dto.CategoryDto) (*dto.CategoryDto, error)
	FindAllCategories() (*[]dto.CategoryDto, error)
}

type CategoryControllerImpl struct {
	createCategory          usecase.CreateCategorieUseCase
	findAllCategoriesByUser usecase.FindAllCategoriesUseCase
}

func NewCategoryController() *CategoryControllerImpl {

	return &CategoryControllerImpl{createCategory: usecase.NewCreateCategory(), findAllCategoriesByUser: usecase.NewFindAllCategoriesUseCase()}
}

func (controller *CategoryControllerImpl) CreateCategory(categoryDto *dto.CategoryDto) (*dto.CategoryDto, error) {
	category := controller.toDomain(categoryDto)
	var err error
	category, err = controller.createCategory.Execute(category)
	response := toDto(category)
	return &response, err
}

func toDto(category *domain.Category) dto.CategoryDto {
	return dto.CategoryDto{
		ID:            category.ID,
		Name:          category.Name,
		Grade:         category.Grade,
		CurrentAmount: category.CurrentAmount,
		TargetAmount:  category.TargetAmount,
	}
}

func (controller *CategoryControllerImpl) toDomain(categoryDto *dto.CategoryDto) *domain.Category {
	return &domain.Category{
		Name:          categoryDto.Name,
		Grade:         categoryDto.Grade,
		CurrentAmount: categoryDto.CurrentAmount,
		TargetAmount:  categoryDto.TargetAmount,
	}
}

func (controller *CategoryControllerImpl) FindAllCategories() (*[]dto.CategoryDto, error) {
	categories, err := controller.findAllCategoriesByUser.Execute()
	if err != nil {
		return nil, err
	}
	categoriesDto := toDtoList(categories)
	return &categoriesDto, nil
}

func toDtoList(categories *[]domain.Category) []dto.CategoryDto {
	var categoriesDto []dto.CategoryDto
	for _, category := range *categories {
		categoriesDto = append(categoriesDto, toDto(&category))
	}
	return categoriesDto
}
