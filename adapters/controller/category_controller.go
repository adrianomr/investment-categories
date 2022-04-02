package controller

import (
	"adrianorodrigues.com.br/investment-categories/domain"
	"adrianorodrigues.com.br/investment-categories/framework/entrypoint/rest/dto"
	"adrianorodrigues.com.br/investment-categories/usecase"
)

type CategoryController interface {
	CreateCategory(category *dto.CategoryDto) (*dto.CategoryDto, error)
	FindAllCategories(userId int) (*dto.WalletDto, error)
	UpdateCategory(category *dto.CategoryDto) (*dto.CategoryDto, error)
}

type CategoryControllerImpl struct {
	findAllCategoriesByUser usecase.FindAllCategoriesUseCase
	createCategory          usecase.CreateCategorieUseCase
	updateCategory          usecase.CreateCategorieUseCase
}

func NewCategoryController() *CategoryControllerImpl {

	return &CategoryControllerImpl{createCategory: usecase.NewCreateCategory(), findAllCategoriesByUser: usecase.NewFindAllCategoriesUseCase(), updateCategory: usecase.NewUpdateCategoryUseCase()}
}

func (controller *CategoryControllerImpl) CreateCategory(categoryDto *dto.CategoryDto) (*dto.CategoryDto, error) {
	category := toCategory(categoryDto)
	var err error
	category, err = controller.createCategory.Execute(category)
	response := toDto(category)
	return response, err
}

func (controller *CategoryControllerImpl) UpdateCategory(categoryDto *dto.CategoryDto) (*dto.CategoryDto, error) {
	category := toCategory(categoryDto)
	var err error
	category, err = controller.updateCategory.Execute(category)
	response := toDto(category)
	return response, err
}

func toDto(category *domain.Category) *dto.CategoryDto {
	if category == nil {
		return nil
	}
	return &dto.CategoryDto{
		ID:             category.ID,
		Name:           category.Name,
		Grade:          category.Grade,
		CurrentAmount:  category.CurrentAmount,
		TargetAmount:   category.TargetAmount,
		InvestedAmount: category.InvestedAmount,
		Investments:    toInvestmentDtos(category.Investments),
	}
}

func toInvestmentDtos(investments []*domain.Investment) []*dto.InvestmentDto {
	if investments == nil {
		return nil
	}
	investmentDtos := []*dto.InvestmentDto{}
	for _, investment := range investments {
		investmentDtos = append(investmentDtos, toInvestmentDto(investment))
	}
	return investmentDtos
}

func toInvestmentDto(investment *domain.Investment) *dto.InvestmentDto {
	return &dto.InvestmentDto{
		ID:            investment.ID,
		Name:          investment.Name,
		Grade:         investment.Grade,
		Origin:        investment.Origin,
		CurrentAmount: investment.CurrentAmount,
		TargetAmount:  investment.TargetAmount,
		Category:      toDto(investment.Category),
	}
}

func toCategory(categoryDto *dto.CategoryDto) *domain.Category {
	if categoryDto == nil {
		return nil
	}
	return &domain.Category{
		ID:             categoryDto.ID,
		Name:           categoryDto.Name,
		Grade:          categoryDto.Grade,
		CurrentAmount:  categoryDto.CurrentAmount,
		TargetAmount:   categoryDto.TargetAmount,
		InvestedAmount: categoryDto.InvestedAmount,
		UserId:         categoryDto.UserId,
		Category:       toCategory(categoryDto.Category),
		Investments:    toInvestments(categoryDto.Investments),
	}
}

func (controller *CategoryControllerImpl) FindAllCategories(userId int) (*dto.WalletDto, error) {
	wallet, err := controller.findAllCategoriesByUser.Execute(userId)
	if err != nil {
		return nil, err
	}
	categoriesDto := toWalletDto(wallet)
	return categoriesDto, nil
}

func toWalletDto(wallet *domain.Wallet) *dto.WalletDto {
	return &dto.WalletDto{
		TotalAmount:       wallet.TotalAmount,
		InvestedAmount:    wallet.InvestedAmount,
		Balance:           wallet.Balance,
		PercentageBalance: wallet.PercentageBalance,
		Categories:        toDtoList(wallet.Categories),
	}
}

func toDtoList(categories *[]domain.Category) *[]dto.CategoryDto {
	categoriesDto := []dto.CategoryDto{}
	for _, category := range *categories {
		categoriesDto = append(categoriesDto, *toDto(&category))
	}
	return &categoriesDto
}
