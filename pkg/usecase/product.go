package usecase

import "github.com/gin-gonic/gin"

type Product interface {
	GetProducts(c *gin.Context)
	GetProductByID(c *gin.Context)
	CreateProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}
