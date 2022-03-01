package usecase

import (
	"adrianorodrigues.com.br/investment-categories/adapters/gateway"
	"adrianorodrigues.com.br/investment-categories/domain"
)

type FindAllCategoriesUseCase interface {
	Execute(int) (*[]domain.Category, error)
}

type FindAllCategoriesUseCaseImpl struct {
	gateway gateway.CategoryGateway
}

func NewFindAllCategoriesUseCase() *FindAllCategoriesUseCaseImpl {

	return &FindAllCategoriesUseCaseImpl{
		gateway: gateway.NewCategoryGateway(),
	}
}

func (useCase *FindAllCategoriesUseCaseImpl) Execute(userId int) (*[]domain.Category, error) {

	response, err := useCase.gateway.FindAllCategories(userId)

	return response, err
}
