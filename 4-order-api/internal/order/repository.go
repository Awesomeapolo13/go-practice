package order

import "go/order-api/pkg/db"

type OrderRepository struct {
	Database *db.Db
}

func NewOrderRepository(database *db.Db) *OrderRepository {
	return &OrderRepository{
		Database: database,
	}
}

func (repo *OrderRepository) Create(order *Order) (*Order, error) {
	result := repo.Database.DB.Create(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return order, nil
}
