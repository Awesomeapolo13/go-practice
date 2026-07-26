package verify

type Varification struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

func NewVerification(email, password string) *Varification {
	return &Varification{
		Email: email,
		Hash:  password,
	}
}
