package gateway

import (
	"adrianorodrigues.com.br/investment-categories/domain"
	"adrianorodrigues.com.br/investment-categories/framework/data/sql/dto"
)

func toInvestmentDto(domain *domain.Investment) *dto.InvestmentDto {
	if domain == nil {
		return nil
	}
	return &dto.InvestmentDto{
		ID:            domain.ID,
		Name:          domain.Name,
		Grade:         domain.Grade,
		Origin:        domain.Origin,
		CurrentAmount: domain.CurrentAmount,
		TargetAmount:  domain.TargetAmount,
		Category:      toCategoryDto(domain.Category),
	}
}

func toInvestmentsDto(dtos []*domain.Investment) []*dto.InvestmentDto {
	if dtos == nil {
		return nil
	}
	investments := []*dto.InvestmentDto{}
	for _, dto := range dtos {
		investments = append(investments, toInvestmentDto(dto))
	}
	return investments
}
