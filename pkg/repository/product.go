package repository

import (
	"simple-go-restapi/pkg/models"
)

type Product interface {
	Index() []models.Product
}
