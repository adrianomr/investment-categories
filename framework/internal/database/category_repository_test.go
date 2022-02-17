package database_test

import (
	"github.com/jinzhu/gorm"
	"log"
	"testing"

	"adrianorodrigues.com.br/investment-categories/framework/internal/database"
	"adrianorodrigues.com.br/investment-categories/framework/internal/database/dto"
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

func setUp(db *gorm.DB) *database.CategoryRepositoryImpl {

	repository := database.NewCategoryRepository(db)

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
	db := database.NewDbTest()
	defer db.Close()
	repository := setUp(db)

	response, err := repository.Find(category.ID)

	log.Printf("Category %v", response)
	require.Nil(t, err)
	require.Equal(t, category, response)
}

func TestCategoryRepoShouldSaveSubCategory(t *testing.T) {

	db := database.NewDbTest()
	defer db.Close()
	repository := setUp(db)

	response, err := repository.Find(category2.ID)

	log.Printf("Category %v", response)
	require.Nil(t, err)
	require.Equal(t, category, response.Category)
}
