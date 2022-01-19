package main

import (
	. "derekshop.com/product/apis"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)

	router.POST("/product", AddProductApi)

	router.GET("/products", GetProductsApi)

	router.GET("/product/:id", GetProductApi)

	router.PUT("/product/:id", ModProductApi)

	router.DELETE("/product/:id", DelProductApi)

	return router
}
