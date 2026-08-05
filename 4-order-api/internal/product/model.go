package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// Создание продукта
// Обновление продукта
// Удаление продукта
// Получение продукта по ID
type Product struct {
	gorm.Model
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Images      pq.StringArray `json:"images"`
}

func NewProduct(name, description string, images []string) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Images:      images,
	}
}
