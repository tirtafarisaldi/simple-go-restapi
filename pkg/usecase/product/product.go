package product

import (
	"net/http"
	"simple-go-restapi/pkg/repository"
	"simple-go-restapi/pkg/usecase"

	"github.com/gin-gonic/gin"
)

type productUsecase struct {
	productRepo repository.Product
}

func Usecase(repository repository.Product) usecase.Product {
	return &productUsecase{
		productRepo: repository,
	}
}

func (uc *productUsecase) Index(c *gin.Context) {

	products := uc.productRepo.Index()
	c.JSON(http.StatusOK, gin.H{"products": products})
}
