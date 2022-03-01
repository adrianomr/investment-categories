package integration

import "adrianorodrigues.com.br/investment-categories/framework/data/sql/dto"

var CategoryDatabase = &dto.CategoryDto{
	ID:            "TEST",
	Name:          "Ações",
	Grade:         10,
	CurrentAmount: 20,
	TargetAmount:  20,
	UserId:        1,
}

var CategoryDatabase2 = &dto.CategoryDto{
	ID:            "TEST2",
	Name:          "Ações",
	Grade:         10,
	CurrentAmount: 20,
	TargetAmount:  20,
	UserId:        1,
	Category:      CategoryDatabase,
}
