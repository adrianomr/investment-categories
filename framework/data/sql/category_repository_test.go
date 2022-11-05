package sql_test

import (
	"adrianorodrigues.com.br/investment-categories/framework/data/sql"
	integration "adrianorodrigues.com.br/investment-categories/integration_test"
	"log"
	"testing"

	"adrianorodrigues.com.br/investment-categories/framework/data/sql/dto"
	"github.com/stretchr/testify/require"
)

var category = &dto.CategoryDto{
	ID:            "TEST",
	Name:          "Ações",
	Grade:         10,
	CurrentAmount: 20,
	TargetAmount:  20,
	UserId:        1,
}

var category2 = &dto.CategoryDto{
	ID:            "TEST2",
	Name:          "Ações",
	Grade:         10,
	CurrentAmount: 20,
	TargetAmount:  20,
	UserId:        1,
	Category:      category,
}
var repository = setUp()

func setUp() sql.CategoryRepository {
	integration.NewPrepareForTests().Prepare()
	repository := sql.NewCategoryRepository()

	_, err := repository.Save(category)
	if err != nil {
		log.Fatal(err)
	}
	_, err = repository.Save(category2)
	if err != nil {
		log.Fatal(err)
	}

	return repository
}

func TestCategoryRepoShouldSaveCategory(t *testing.T) {
	response, err := repository.Find(category.ID)

	log.Printf("Category %v", response)
	require.Nil(t, err)
	require.Equal(t, category, response)
}

func TestCategoryRepoShouldSaveSubCategory(t *testing.T) {
	response, err := repository.Find(category2.ID)

	log.Printf("Category %v", response)
	require.Nil(t, err)
	require.Equal(t, category, response.Category)
}

func TestCategoryRepoShouldFindAllForUser1(t *testing.T) {
	response, err := repository.FindAllCategoriesByUserId(1)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Category %v", response)
	require.Equal(t, 2, len(*response))
}
