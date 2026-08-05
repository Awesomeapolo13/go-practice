package product

type CreateProductRequest struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
}

type UpdateProductRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
}
