package usecase

import "github.com/nathanfabio/goREST/model"

type ProductUseCase struct {
}

func NewProductUseCase() ProductUseCase {
	return ProductUseCase{}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return []model.Product{}, nil
}