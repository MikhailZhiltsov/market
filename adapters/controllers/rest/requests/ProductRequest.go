package requests

type ProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
