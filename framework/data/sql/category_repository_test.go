package sql_test

import (
	"github.com/jinzhu/gorm"
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

func setUp(db *gorm.DB) gorm.CategoryRepository {

	repository := gorm.NewCategoryRepository(db)

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
	db := gorm.NewDbTest()
	defer db.Close()
	repository := setUp(db)

	response, err := repository.Find(category.ID)

	log.Printf("Category %v", response)
	require.Nil(t, err)
	require.Equal(t, category, response)
}

func TestCategoryRepoShouldSaveSubCategory(t *testing.T) {

	db := gorm.NewDbTest()
	defer db.Close()
	repository := setUp(db)

	response, err := repository.Find(category2.ID)

	log.Printf("Category %v", response)
	require.Nil(t, err)
	require.Equal(t, category, response.Category)
}

func TestCategoryRepoShouldFindAllForUser1(t *testing.T) {

	db := gorm.NewDbTest()
	defer db.Close()
	repository := setUp(db)

	response, err := repository.FindAllCategoriesByUserId(1)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Category %v", response)
	require.Equal(t, 2, len(*response))
}
