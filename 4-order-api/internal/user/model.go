package user

import (
	"go/order-api/internal/order"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PhoneNumber      string `json:"phone_number" gorm:"uniqueIndex"`
	SessionId        string
	VerificationCode string
	Orders           []order.Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewUser(phoneNumber, sessionId, verificationCode string) *User {
	return &User{
		PhoneNumber:      phoneNumber,
		SessionId:        sessionId,
		VerificationCode: verificationCode,
	}
}
