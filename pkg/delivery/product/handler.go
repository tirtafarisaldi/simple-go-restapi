package delivery

import (
	"simple-go-restapi/pkg/usecase"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUsecase usecase.Product
}

func Handler(r *gin.Engine, uc usecase.Product) {
	handler := &ProductHandler{
		productUsecase: uc,
	}
	r.GET("/api/products", handler.ProductIndexHandler)

}

func (h *ProductHandler) ProductIndexHandler(c *gin.Context) {
	h.productUsecase.Index(c)
}
