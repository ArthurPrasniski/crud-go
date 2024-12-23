package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUseCase := usecase.NewProductUsecase(ProductRepository)
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/products", ProductController.GetProduct)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:id", ProductController.GetProductByID)
	server.DELETE("/product/:id", ProductController.DeleteProductByID)
	server.PUT("/product/:id", ProductController.UpdateProductByID)

	server.Run(":8080")
}
