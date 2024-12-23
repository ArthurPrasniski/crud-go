package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) error {
	return pu.repository.CreateProduct(product)
}

func (pu *ProductUsecase) GetProductByID(id int) (model.Product, error) {
	return pu.repository.GetProductByID(id)
}

func (pu *ProductUsecase) DeleteProductByID(id int) error {
	return pu.repository.DeleteProductByID(id)
}

func (pu *ProductUsecase) UpdateProductByID(id int, product model.Product) error {
	return pu.repository.UpdateProductByID(id, product)
}
