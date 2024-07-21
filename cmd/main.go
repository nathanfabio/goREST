package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanfabio/goREST/controller"
	"github.com/nathanfabio/goREST/db"
	"github.com/nathanfabio/goREST/repository"
	usecase "github.com/nathanfabio/goREST/useCase"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	//useCase
	ProductUseCase:= usecase.NewProductUseCase(ProductRepository)

	//controllers
	ProductController:= controller.NewProductController(ProductUseCase)


	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductsById)
	server.DELETE("/product/:productId", ProductController.DeleteProduct)

	server.Run(":8080")
}