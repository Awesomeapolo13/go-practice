package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PhoneNumber string `json:"phone_number" gorm:"uniqueIndex"`
	SessionId   string
}

func NewUser(phoneNumber, sessionId string) *User {

	user := &User{
		PhoneNumber: phoneNumber,
		SessionId:   sessionId,
	}

	return user
}
