package order

import (
	"go/order-api/internal/product"
	"go/order-api/internal/user"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderDate  time.Time         `json:"order_date" gorm:"not null"`
	IsDelivery bool              `json:"is_delivery" gorm:"not null"`
	IsExpress  bool              `json:"is_express"  gorm:"not null"`
	User       user.User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Products   []product.Product `gorm:"many2many:order_products;"`
}

func NewOrder(orderDate time.Time, isDelivery, isExpress bool, user user.User, products []product.Product) *Order {
	return &Order{
		OrderDate:  orderDate,
		IsDelivery: isDelivery,
		IsExpress:  isExpress,
		User:       user,
		Products:   products,
	}
}
