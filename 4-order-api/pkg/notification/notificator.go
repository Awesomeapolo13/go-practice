package notification

import "fmt"

type Notificator struct {
}

func NewNotificator() *Notificator {
	return &Notificator{}
}

func (n *Notificator) SendCodeBySMS(phone, code string) {
	// Сказано пока не отправлять кода
	fmt.Println("Code " + code + " has been sent on number " + phone)
}
