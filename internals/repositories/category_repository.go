package repositories

import (
	"github.com/ahmadammarm/go-rest-api/internals/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {

	var categories []models.Category
	err := r.DB.Find(&categories).Error

	if err != nil {
		return nil, err
	}

	r.DB.Find(&categories)
	return categories, err
}

func (r *CategoryRepository) CreateCategory(category models.Category) (models.Category, error) {

	err := r.DB.Create(&category).Error

	if err != nil {
		return category, err
	}

	return category, err
}

func (r *CategoryRepository) GetCategoryById(id int) (models.Category, error) {

	var category models.Category
	err := r.DB.First(&category, id).Error

	if err != nil {
		return category, err
	}

	r.DB.First(&category, id)
	return category, err
}

func (r *CategoryRepository) UpdateCategory(id int, category models.Category) (models.Category, error) {

	var categoryUpdate models.Category
	err := r.DB.First(&categoryUpdate, id).Error

	if err != nil {
		return category, err
	}

	r.DB.Model(&categoryUpdate).Updates(category)
	return category, err
}

func (r *CategoryRepository) DeleteCategory(id int) error {

	var category models.Category
	err := r.DB.First(&category, id).Error

	if err != nil {
		return err
	}

	r.DB.Delete(&category, id)
	return err
}
