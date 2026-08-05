package product

import (
	"go/order-api/pkg/db"

	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	Database *db.Db
}

func NewProductRepository(database *db.Db) *ProductRepository {
	return &ProductRepository{
		Database: database,
	}
}

func (repo *ProductRepository) Create(product *Product) (*Product, error) {
	result := repo.Database.DB.Create(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repo *ProductRepository) FindByID(id uint) (*Product, error) {
	var product *Product
	result := repo.Database.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repo *ProductRepository) Update(product *Product) (*Product, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}
