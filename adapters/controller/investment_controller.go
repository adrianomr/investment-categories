package controller

import (
	"adrianorodrigues.com.br/investment-categories/domain"
	"adrianorodrigues.com.br/investment-categories/framework/entrypoint/rest/dto"
)

func toInvestment(dto *dto.InvestmentDto) *domain.Investment {
	if dto == nil {
		return nil
	}
	return &domain.Investment{
		ID:            dto.ID,
		Name:          dto.Name,
		Grade:         dto.Grade,
		Origin:        dto.Origin,
		CurrentAmount: dto.CurrentAmount,
		TargetAmount:  dto.TargetAmount,
		Category:      toCategory(dto.Category),
	}
}

func toInvestments(dtos []*dto.InvestmentDto) []*domain.Investment {
	if dtos == nil {
		return nil
	}
	investments := []*domain.Investment{}
	for _, dto := range dtos {
		investments = append(investments, toInvestment(dto))
	}
	return investments
}
