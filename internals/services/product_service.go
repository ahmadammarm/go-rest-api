package services

import (
	"github.com/ahmadammarm/go-rest-api/internals/models"
	"github.com/ahmadammarm/go-rest-api/internals/repositories"
)

type ProductService struct {
    Repo *repositories.ProductRepository
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
    return s.Repo.GetAllProducts()
}

func (s *ProductService) CreateProduct(product models.Product) (models.Product, error) {
    return s.Repo.CreateProduct(product)
}

func (s *ProductService) GetProductById(id int) (models.Product, error) {
    return s.Repo.GetProductById(id)
}

func (s *ProductService) UpdateProduct(id int, product models.Product) (models.Product, error) {
    return s.Repo.UpdateProduct(id, product)
}

func (s *ProductService) DeleteProduct(id int) error {
    return s.Repo.DeleteProduct(id)
}