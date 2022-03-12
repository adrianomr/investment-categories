package gateway

import (
	"adrianorodrigues.com.br/investment-categories/domain"
	"adrianorodrigues.com.br/investment-categories/framework/data/sql/dto"
)
import "adrianorodrigues.com.br/investment-categories/framework/data/sql"

type CategoryGateway interface {
	CreateCategory(category *domain.Category) (*domain.Category, error)
	FindAllCategories(int) (*[]domain.Category, error)
	UpdateCategory(category *domain.Category) (*domain.Category, error)
}

type CategoryGatewayImpl struct {
	repository sql.CategoryRepository
}

func NewCategoryGateway() *CategoryGatewayImpl {

	return &CategoryGatewayImpl{repository: sql.NewCategoryRepository()}
}

func (gateway *CategoryGatewayImpl) CreateCategory(category *domain.Category) (*domain.Category, error) {
	categoryDto := toCategoryDto(category)
	response, err := gateway.repository.Save(categoryDto)
	return toDomain(response), err
}

func (gateway *CategoryGatewayImpl) UpdateCategory(category *domain.Category) (*domain.Category, error) {
	categoryDto := toCategoryDto(category)
	response, err := gateway.repository.Save(categoryDto)
	return toDomain(response), err
}

func (gateway *CategoryGatewayImpl) FindAllCategories(userId int) (*[]domain.Category, error) {
	categoriesDtoList, err := gateway.repository.FindAllCategoriesByUserId(userId)
	if err != nil {
		return nil, err
	}
	categoriesList := toDomainList(*categoriesDtoList)
	return categoriesList, nil
}

func toDomainList(response []dto.CategoryDto) *[]domain.Category {
	categories := []domain.Category{}
	for _, categoryDto := range response {
		categories = append(categories, *toDomain(&categoryDto))
	}
	return &categories
}

func toDomain(response *dto.CategoryDto) *domain.Category {
	if response == nil {
		return nil
	}
	return &domain.Category{
		ID:            response.ID,
		Name:          response.Name,
		Grade:         response.Grade,
		CurrentAmount: response.CurrentAmount,
		TargetAmount:  response.TargetAmount,
		Category:      toDomain(response.Category),
		Investments:   toInvestments(response.Investments),
	}
}

func toInvestments(investmentsDto []*dto.InvestmentDto) []*domain.Investment {
	if investmentsDto == nil {
		return nil
	}
	investments := []*domain.Investment{}
	for _, dto := range investmentsDto {
		investments = append(investments, toInvestment(dto))
	}
	return investments
}

func toInvestment(investmentDto *dto.InvestmentDto) *domain.Investment {
	return &domain.Investment{
		ID:            investmentDto.ID,
		Name:          investmentDto.Name,
		Grade:         investmentDto.Grade,
		Origin:        investmentDto.Origin,
		CurrentAmount: investmentDto.CurrentAmount,
		TargetAmount:  investmentDto.TargetAmount,
		Category:      toDomain(investmentDto.Category),
	}
}

func toCategoryDto(category *domain.Category) *dto.CategoryDto {
	if category == nil {
		return nil
	}
	return &dto.CategoryDto{
		ID:            category.ID,
		Name:          category.Name,
		Grade:         category.Grade,
		CurrentAmount: category.CurrentAmount,
		TargetAmount:  category.TargetAmount,
		UserId:        category.UserId,
		Category:      toCategoryDto(category.Category),
		Investments:   toInvestmentsDto(category.Investments),
	}
}
