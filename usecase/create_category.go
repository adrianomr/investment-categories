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

	response, err := useCase.gateway.CreateCategory(category)

	return response, err
}
