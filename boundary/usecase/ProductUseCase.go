package usecase

import "market/boundary/dto"

type ProductUseCase interface {
	CreateProduct(productDTO *dto.ProductDTO) error

	GetProduct(name string) (*dto.ProductDTO, error)

	UpdateProduct(productDTO *dto.ProductDTO) error

	DeleteProduct(name string) error
}
