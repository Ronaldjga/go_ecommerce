package usecase

import (
	"go_ecommerce/model"
	"go_ecommerce/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

// METHODS
func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	repo, err := pu.repository.GetProducts()
	if err != nil {
		panic(err)
	}
	return repo, nil
}

// Functions
func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}
