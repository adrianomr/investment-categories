package usecase

import (
	"adrianorodrigues.com.br/investment-categories/adapters/gateway"
	"adrianorodrigues.com.br/investment-categories/domain"
)

type CreateCategorieUseCase interface {
	createCategorie(category *domain.Category) (*domain.Category, error)
}

type CreateCategorieUseCaseImpl struct {
	gateway gateway.CategoryGateway
}

func NewService() *CreateCategorieUseCaseImpl {

	return &CreateCategorieUseCaseImpl{
		gateway: gateway.NewCategoryGateway(),
	}
}

func (useCase *CreateCategorieUseCaseImpl) createCategorie(category *domain.Category) (*domain.Category, error) {

	response, err := useCase.gateway.CreateCategory(category)

	return response, err
}
