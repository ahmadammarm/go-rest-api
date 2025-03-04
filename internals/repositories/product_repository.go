package repositories

import (
	"github.com/ahmadammarm/go-rest-api/internals/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
    DB *gorm.DB
}


func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {

    var products []models.Product
    err := r.DB.Find(&products).Error

    if err != nil {
        return nil, err
    }

    r.DB.Find(&products)
    return products, err
}


func (r *ProductRepository) CreateProduct(product models.Product) (models.Product, error) {

    err := r.DB.Create(&product).Error

    if err != nil {
        return product, err
    }

    return product, err
}


func (r *ProductRepository) GetProductById(id int) (models.Product, error) {

    var product models.Product
    err := r.DB.First(&product, id).Error

    if err != nil {
        return product, err
    }

    r.DB.First(&product, id)
    return product, err
}


func (r *ProductRepository) UpdateProduct(id int, product models.Product) (models.Product, error) {

    var productUpdate models.Product
    err := r.DB.First(&productUpdate, id).Error

    if err != nil {
        return product, err
    }

    r.DB.Model(&productUpdate).Updates(product)
    return product, err
}


func (r *ProductRepository) DeleteProduct(id int) error {

    var product models.Product
    err := r.DB.First(&product, id).Error

    if err != nil {
        return err
    }

    r.DB.Delete(&product, id)
    return nil
}