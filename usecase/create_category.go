package usecase

import (
	"adrianorodrigues.com.br/investment-categories/adapters/gateway"
	"adrianorodrigues.com.br/investment-categories/domain"
)

type CreateCategorieUseCase interface {
	Execute(category *domain.Category) (*domain.Category, error)
}

type CreateCategorieUseCaseImpl struct {
	gateway gateway.CategoryGateway
}

func NewCreateCategory() *CreateCategorieUseCaseImpl {

	return &CreateCategorieUseCaseImpl{
		gateway: gateway.NewCategoryGateway(),
	}
}

func (useCase *CreateCategorieUseCaseImpl) Execute(category *domain.Category) (*domain.Category, error) {
	categories, err := useCase.gateway.FindAllCategories(category.UserId)
	if err != nil {
		return nil, err
	}
	category.CalculateTarget(categories)
	response, err := useCase.gateway.CreateCategory(category)

	return response, err
}
