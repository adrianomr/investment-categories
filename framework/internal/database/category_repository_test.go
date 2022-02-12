package database_test

import (
	"testing"

	"adrianorodrigues.com.br/investment-categories/framework/internal/database"
	"adrianorodrigues.com.br/investment-categories/framework/internal/database/dto"
	"github.com/stretchr/testify/require"
)

func CategoryRepoTest(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	category := &dto.CategoryDto{
		ID:            "TEST",
		Name:          "Ações",
		Grade:         10,
		CurrentAmount: 20,
		TargetAmount:  20,
		UserId:        1,
	}

	repository := database.NewCategoryRepository(db)

	response, err := repository.Save(category)

	require.Equal(t, category, response)
	require.Nil(t, err)

	response, err = repository.Find(category.ID)

	require.Equal(t, category, response)
	require.Nil(t, err)
}
