package order

import (
	"go/order-api/pkg/db"
)

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

func (repo *OrderRepository) FindUserOrderById(orderId uint, userId uint) (*Order, error) {
	var order *Order
	result := repo.Database.DB.
		Preload("User").
		Where("id = ? AND user_id = ?", orderId, userId).
		First(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return order, nil
}
