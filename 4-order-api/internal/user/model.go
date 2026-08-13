package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PhoneNumber      string `json:"phone_number" gorm:"uniqueIndex"`
	SessionId        string
	VerificationCode string
}

func NewUser(phoneNumber, sessionId, verificationCode string) *User {

	user := &User{
		PhoneNumber:      phoneNumber,
		SessionId:        sessionId,
		VerificationCode: verificationCode,
	}

	return user
}
