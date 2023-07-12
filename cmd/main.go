package main

import (
	"simple-go-restapi/config"
	productController "simple-go-restapi/pkg/controller/product"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.ConnectDatabase()
	productController.Controller(r, db)

	r.Run()
}
