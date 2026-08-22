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
		Preload("Products").
		Where("id = ? AND user_id = ?", orderId, userId).
		First(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return order, nil
}

func (repo *OrderRepository) GetUserOrders(userId uint) []Order {
	var orders []Order
	repo.Database.DB.
		Table("orders").
		Preload("User").
		Preload("Products").
		Where("user_id = ?", userId).
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Scan(&orders)

	return orders
}
