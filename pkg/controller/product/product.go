package product

import (
	productHandler "simple-go-restapi/pkg/delivery/product"
	productRepository "simple-go-restapi/pkg/repository/product"
	productUsecase "simple-go-restapi/pkg/usecase/product"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Controller(r *gin.Engine, db *gorm.DB) {
	productRepo := productRepository.Repository(db)
	productUC := productUsecase.Usecase(productRepo)

	productHandler.Handler(r, productUC)
}
