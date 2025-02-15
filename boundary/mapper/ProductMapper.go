package mapper

import (
	"market/adapters/controllers/rest/requests"
	"market/boundary/dto"
)

func MapProductRequestToDTO(req *requests.ProductRequest) *dto.ProductDTO {
	return &dto.ProductDTO{
		Name:  req.Name,
		Price: req.Price,
	}
}

func MapDTOToProductRequest(dto *dto.ProductDTO) *requests.ProductRequest {
	return &requests.ProductRequest{
		Name:  dto.Name,
		Price: dto.Price,
	}
}
