package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathanfabio/goREST/model"
	usecase "github.com/nathanfabio/goREST/useCase"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products:= []model.Product{
		{
			ID: 1,
			Name: "Pêra",
			Price: 30.20,
		},
	}

	ctx.JSON(http.StatusOK, products)
}