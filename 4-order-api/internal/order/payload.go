package order

import "time"

type CreateOrderRequest struct {
	OrderDate  time.Time `json:"order_date" validate:"required"`
	IsDelivery bool      `json:"is_delivery" validate:"required"`
	IsExpress  bool      `json:"is_express" validate:"required"`
	Products   []uint    `json:"products" validate:"required"`
}
