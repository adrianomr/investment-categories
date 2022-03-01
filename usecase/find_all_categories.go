package usecase

import (
	"adrianorodrigues.com.br/investment-categories/adapters/gateway"
	"adrianorodrigues.com.br/investment-categories/domain"
)

type FindAllCategoriesUseCase interface {
	Execute() (*[]domain.Category, error)
}

type FindAllCategoriesUseCaseImpl struct {
	gateway gateway.CategoryGateway
}

func NewFindAllCategoriesUseCase() *FindAllCategoriesUseCaseImpl {

	return &FindAllCategoriesUseCaseImpl{
		gateway: gateway.NewCategoryGateway(),
	}
}

func (useCase *FindAllCategoriesUseCaseImpl) Execute() (*[]domain.Category, error) {

	response, err := useCase.gateway.FindAllCategories()

	return response, err
}
