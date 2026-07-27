package auth

type AuthRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
}

type AuthResponse struct {
	Code string `json:"code"`
}

type VerifyRequest struct {
	Code string `json:"code" validate:"required,len=4"`
}

type VerifyResponse struct {
	Token string `json:"token"`
}
