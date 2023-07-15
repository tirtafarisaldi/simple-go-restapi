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
	r.GET("/api/products", handler.GetProductsHandler)
	r.GET("/api/products/:id", handler.GetProductByIDHandler)
	r.POST("/api/product", handler.CreateProductHandler)
	r.PUT("/api/product/:id", handler.UpdateProductHandler)
	r.DELETE("/api/product", handler.DeleteProductHandler)

}

func (h *ProductHandler) GetProductsHandler(c *gin.Context) {
	h.productUsecase.GetProducts(c)
}

func (h *ProductHandler) GetProductByIDHandler(c *gin.Context) {
	h.productUsecase.GetProductByID(c)
}

func (h *ProductHandler) CreateProductHandler(c *gin.Context) {
	h.productUsecase.CreateProduct(c)
}

func (h *ProductHandler) UpdateProductHandler(c *gin.Context) {
	h.productUsecase.CreateProduct(c)
}

func (h *ProductHandler) DeleteProductHandler(c *gin.Context) {
	h.productUsecase.CreateProduct(c)
}
