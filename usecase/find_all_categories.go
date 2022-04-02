package usecase

import (
	"adrianorodrigues.com.br/investment-categories/adapters/gateway"
	"adrianorodrigues.com.br/investment-categories/domain"
)

type FindAllCategoriesUseCase interface {
	Execute(int) (*domain.Wallet, error)
}

type FindAllCategoriesUseCaseImpl struct {
	gateway gateway.CategoryGateway
}

func NewFindAllCategoriesUseCase() *FindAllCategoriesUseCaseImpl {

	return &FindAllCategoriesUseCaseImpl{
		gateway: gateway.NewCategoryGateway(),
	}
}

func (useCase *FindAllCategoriesUseCaseImpl) Execute(userId int) (*domain.Wallet, error) {

	response, err := useCase.gateway.FindAllCategories(userId)
	if err != nil {
		return nil, err
	}
	var wallet = &domain.Wallet{
		TotalAmount:       0,
		InvestedAmount:    0,
		Balance:           0,
		PercentageBalance: 0,
		Categories:        response,
	}

	wallet.Calculate()

	return wallet, err
}
