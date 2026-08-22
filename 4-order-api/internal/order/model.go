package order

import (
	"go/order-api/internal/product"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderedAt  time.Time `json:"ordered_at" gorm:"not null"`
	IsDelivery bool      `json:"is_delivery" gorm:"not null"`
	IsExpress  bool      `json:"is_express"  gorm:"not null"`
	UserID     uint
	Products   []product.Product `gorm:"many2many:order_products;"`
}
