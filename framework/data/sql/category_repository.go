package sql

import (
	"fmt"
	"github.com/google/uuid"

	"adrianorodrigues.com.br/investment-categories/framework/data/sql/dto"
)

type CategoryRepository interface {
	Save(category *dto.CategoryDto) (*dto.CategoryDto, error)
	Find(id string) (*dto.CategoryDto, error)
	FindAllCategoriesByUserId(userId int) (*[]dto.CategoryDto, error)
}

type CategoryRepositoryImpl struct {
	database Database
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{database: *DatabaseSingleton()}
}

func (repo *CategoryRepositoryImpl) Save(category *dto.CategoryDto) (*dto.CategoryDto, error) {
	db, err := repo.database.GetDb()
	if err != nil {
		return nil, err
	}

	if category.ID == "" {
		newUuid, err := uuid.NewUUID()
		if err != nil {
			return nil, err
		}
		category.ID = newUuid.String()
	}
	err = db.Save(category).Error

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (repo *CategoryRepositoryImpl) Find(id string) (*dto.CategoryDto, error) {
	db, err := repo.database.GetDb()
	if err != nil {
		return nil, err
	}

	var category dto.CategoryDto
	db.Preload("Category").First(&category, "id = ?", id)

	if category.ID == "" {
		return nil, fmt.Errorf("job does not exists")
	}

	return &category, nil
}

func (repo *CategoryRepositoryImpl) FindAllCategoriesByUserId(userId int) (*[]dto.CategoryDto, error) {
	db, err := repo.database.GetDb()
	if err != nil {
		return nil, err
	}

	categories := []dto.CategoryDto{}
	err = db.Preload("Category").Where("user_id = ?", userId).Find(&categories).Error
	return &categories, err
}
