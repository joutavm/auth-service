package request

type ProductRequest struct {
	Name  string  `json:"name,omitempty" binding:"required,min=2,max=100"`
	Price float64 `json:"price,omitempty" binding:"required,min=2,max=100"`
}

type ProductPatch struct {
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}
