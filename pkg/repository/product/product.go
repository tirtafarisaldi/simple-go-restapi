package product

import (
	"simple-go-restapi/pkg/models"
	"simple-go-restapi/pkg/repository"

	"gorm.io/gorm"
)

type productRepo struct {
	DB *gorm.DB
}

func Repository(db *gorm.DB) repository.Product {
	return &productRepo{
		DB: db,
	}
}

func (r *productRepo) GetProducts() []models.Product {

	var products []models.Product

	r.DB.Find(&products)
	return products

}

func (r *productRepo) GetProductByID(id string) (models.Product, error) {
	var product models.Product
	err := r.DB.First(&product, id).Error

	return product, err
}

func (r *productRepo) CreateProduct() (models.Product, error) {
	var product models.Product
	err := r.DB.Create(&product).Error

	return product, err

}

func (r *productRepo) UpdateProduct(id string) (models.Product, bool) {
	var product models.Product
	isErr := r.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0

	return product, isErr
}

func (r *productRepo) DeleteProduct(id any) (models.Product, bool) {
	var product models.Product
	isErr := r.DB.Delete(&product, id).RowsAffected == 0

	return product, isErr

}
