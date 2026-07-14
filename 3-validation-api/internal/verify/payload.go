package verify

type SendVerificationResponse struct {
	Success bool `json:"success"`
}

type VerificationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
