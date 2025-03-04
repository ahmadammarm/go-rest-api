package services

import (
	"github.com/ahmadammarm/go-rest-api/internals/models"
	"github.com/ahmadammarm/go-rest-api/internals/repositories"
)


type CategoryService struct {
    Repo *repositories.CategoryRepository
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
    return s.Repo.GetAllCategories()
}

func (s *CategoryService) CreateCategory(category models.Category) (models.Category, error) {
    return s.Repo.CreateCategory(category)
}

func (s *CategoryService) GetCategoryById(id int) (models.Category, error) {
    return s.Repo.GetCategoryById(id)
}

func (s *CategoryService) UpdateCategory(id int, category models.Category) (models.Category, error) {
    return s.Repo.UpdateCategory(id, category)
}

func (s *CategoryService) DeleteCategory(id int) error {
    return s.Repo.DeleteCategory(id)
}