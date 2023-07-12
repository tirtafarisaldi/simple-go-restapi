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

func (r *productRepo) Index() []models.Product {

	var products []models.Product

	r.DB.Find(&products)
	return products

}
