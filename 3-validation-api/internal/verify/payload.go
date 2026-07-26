package verify

type SendVerificationRequest struct {
	Email string `json:"email" validate:"required"`
}

type SendVerificationResponse struct {
	Success bool `json:"success"`
}

type VerificationRequest struct {
	Hash string `json:"hash" validate:"required"`
}

type VerificationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
