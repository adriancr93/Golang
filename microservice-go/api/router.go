package api

import (
	"microservice-go/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	productRoutes := r.Group("/api/products")
	{
		productRoutes.GET("/", handlers.GetAllProducts)
		productRoutes.GET("/:id", handlers.GetProductByID)
		productRoutes.POST("/", handlers.CreateProduct)
		productRoutes.DELETE("/:id", handlers.DeleteProduct)
	}

	return r
}