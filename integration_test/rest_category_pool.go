package integration

import "adrianorodrigues.com.br/investment-categories/framework/entrypoint/rest/dto"

var CategoryRest = &dto.CategoryDto{
	ID:            "TEST",
	Name:          "Ações",
	Grade:         10,
	CurrentAmount: 20,
	TargetAmount:  20,
}

var CategoryRest2 = &dto.CategoryDto{
	ID:            "TEST2",
	Name:          "Ações",
	Grade:         10,
	CurrentAmount: 20,
	TargetAmount:  20,
}
