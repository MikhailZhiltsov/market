package dto

type ProductDTO struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewProductDTO(name string, price float64) *ProductDTO {
	return &ProductDTO{
		Name:  name,
		Price: price,
	}
}
