package verify

type SendVerificationRequest struct {
	Email string `json:"email"`
}

type SendVerificationResponse struct {
	Success bool `json:"success"`
}

type VerificationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
