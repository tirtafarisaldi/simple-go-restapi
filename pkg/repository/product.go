package repository

import (
	"simple-go-restapi/pkg/models"
)

type Product interface {
	GetProducts() []models.Product
	GetProductByID(id string) (models.Product, error)
	CreateProduct() (models.Product, error)
	UpdateProduct(id string) (models.Product, bool)
	DeleteProduct(id any) (models.Product, bool)
}
