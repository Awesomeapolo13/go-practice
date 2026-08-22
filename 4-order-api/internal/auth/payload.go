package auth

type AuthRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
}

type AuthResponse struct {
	SessionId string `json:"session_id"`
}

type VerifyRequest struct {
	SessionId string `json:"session_id" validate:"required"`
	Code      string `json:"code" validate:"required,len=4"`
}

type VerifyResponse struct {
	Token string `json:"token"`
}
