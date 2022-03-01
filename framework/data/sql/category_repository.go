package sql

import (
	"fmt"

	"adrianorodrigues.com.br/investment-categories/framework/data/sql/dto"
	"github.com/jinzhu/gorm"
)

type CategoryRepository interface {
	Save(category *dto.CategoryDto) (*dto.CategoryDto, error)
	Find(id string) (*dto.CategoryDto, error)
	FindAllCategoriesByUserId(userId int) (*[]dto.CategoryDto, error)
}

type CategoryRepositoryImpl struct {
	Db *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{Db: DatabaseSingleton().GetDb()}
}

func (repo *CategoryRepositoryImpl) Save(category *dto.CategoryDto) (*dto.CategoryDto, error) {
	err := repo.Db.Create(category).Error

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (repo *CategoryRepositoryImpl) Find(id string) (*dto.CategoryDto, error) {
	var category dto.CategoryDto
	repo.Db.Preload("Category").First(&category, "id = ?", id)

	if category.ID == "" {
		return nil, fmt.Errorf("job does not exists")
	}

	return &category, nil
}

func (repo *CategoryRepositoryImpl) FindAllCategoriesByUserId(userId int) (*[]dto.CategoryDto, error) {
	var categories []dto.CategoryDto
	err := repo.Db.Preload("Category").Where("user_id = ?", userId).Find(&categories).Error
	return &categories, err
}
