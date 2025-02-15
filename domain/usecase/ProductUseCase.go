package usecase

import (
	"errors"
	"market/boundary/dto"
)

type ProductUseCase interface {
	CreateProduct(productDTO *dto.ProductDTO) error
	GetProduct(name string) (*dto.ProductDTO, error)
	UpdateProduct(productDTO *dto.ProductDTO) error
	DeleteProduct(name string) error
}

type DummyProductUseCase struct{}

func NewDummyProductUseCase() ProductUseCase {
	return &DummyProductUseCase{}
}

func (d *DummyProductUseCase) CreateProduct(productDTO *dto.ProductDTO) error {
	return nil
}

func (d *DummyProductUseCase) GetProduct(name string) (*dto.ProductDTO, error) {
	if name == "example" {
		return &dto.ProductDTO{Name: "example", Price: 9.99}, nil
	}
	return nil, errors.New("product not found")
}

func (d *DummyProductUseCase) UpdateProduct(productDTO *dto.ProductDTO) error {
	return nil
}

func (d *DummyProductUseCase) DeleteProduct(name string) error {
	return nil
}
