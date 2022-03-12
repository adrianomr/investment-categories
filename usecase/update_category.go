package usecase

import (
	"adrianorodrigues.com.br/investment-categories/adapters/gateway"
	"adrianorodrigues.com.br/investment-categories/domain"
)

type UpdateCategoryUseCase interface {
	Execute(category *domain.Category) (*domain.Category, error)
}

type UpdateCategoryUseCaseImpl struct {
	gateway gateway.CategoryGateway
}

func NewUpdateCategoryUseCase() *UpdateCategoryUseCaseImpl {

	return &UpdateCategoryUseCaseImpl{
		gateway: gateway.NewCategoryGateway(),
	}
}

func (useCase *UpdateCategoryUseCaseImpl) Execute(category *domain.Category) (*domain.Category, error) {

	response, err := useCase.gateway.UpdateCategory(category)

	return response, err
}
