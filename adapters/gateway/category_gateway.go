package gateway

import "adrianorodrigues.com.br/investment-categories/domain"

type CategoryGateway interface {
	CreateCategory(category *domain.Category) (*domain.Category, error)
}

type CategoryGatewayImpl struct {
	
}

func NewCategoryGateway() *CategoryGatewayImpl{

	return &CategoryGatewayImpl{}
}

func (gateway *CategoryGatewayImpl) CreateCategory(category *domain.Category) (*domain.Category, error){
	return category, nil
}