package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanfabio/goREST/controller"
	usecase "github.com/nathanfabio/goREST/useCase"
)

func main() {

	server := gin.Default()

	ProductUseCase:= usecase.NewProductUseCase()
	//controllers
	ProductController:= controller.NewProductController(ProductUseCase)


	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8080")
}