package usecase

import "github.com/gin-gonic/gin"

type Product interface {
	Index(c *gin.Context)
}
