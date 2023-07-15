package product

import (
	"encoding/json"
	"net/http"
	"simple-go-restapi/pkg/repository"
	"simple-go-restapi/pkg/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productUsecase struct {
	productRepo repository.Product
}

func Usecase(repository repository.Product) usecase.Product {
	return &productUsecase{
		productRepo: repository,
	}
}

func (uc *productUsecase) GetProducts(c *gin.Context) {

	products := uc.productRepo.GetProducts()
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (uc *productUsecase) GetProductByID(c *gin.Context) {
	id := c.Param("id")

	product, err := uc.productRepo.GetProductByID(id)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func (uc *productUsecase) CreateProduct(c *gin.Context) {
	product, err := uc.productRepo.CreateProduct()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func (uc *productUsecase) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	product, isErr := uc.productRepo.UpdateProduct(id)
	err := c.ShouldBindJSON(&product)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if isErr {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

func (uc *productUsecase) DeleteProduct(c *gin.Context) {
	var input struct {
		Id json.Number
	}

	id, _ := input.Id.Int64()
	_, isErr := uc.productRepo.DeleteProduct(id)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if isErr {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
