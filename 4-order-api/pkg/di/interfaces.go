package di

import (
	"go/order-api/internal/product"
	"go/order-api/internal/user"
)

type IProductRepository interface {
	FindAllByIDs(ids []uint) ([]product.Product, error)
}

type IUserRepository interface {
	FindByPhone(phoneNumber string) (*user.User, error)
}
