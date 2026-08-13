package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"go/order-api/internal/user"
	"math/big"
)

const (
	DefaultSessIdLength = 32
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) CreateNewSessionCredentials(phone string) (string, string, error) {
	existedUser, _ := service.UserRepository.FindByPhone(phone)
	sessionId, err := generateSessionId(DefaultSessIdLength)
	if err != nil {
		return "", "", err
	}
	verificationCode, err := generateAuthCode()
	if err != nil {
		return "", "", err
	}

	if existedUser != nil {
		existedUser.SessionId = sessionId
		existedUser.VerificationCode = verificationCode
		_, err = service.UserRepository.Update(existedUser)
		if err != nil {
			return "", "", err
		}

		return sessionId, verificationCode, nil
	}

	existedUser = user.NewUser(phone, sessionId, verificationCode)
	_, err = service.UserRepository.Create(existedUser)
	if err != nil {
		return "", "", err
	}

	return sessionId, verificationCode, nil
}

func (service *AuthService) VerifyByCode(sessionId, code string) (string, error) {
	existedUser, _ := service.UserRepository.FindBySessionId(sessionId)
	if existedUser == nil {
		return "", errors.New(ErrWrongCredentials)
	}

	if code != existedUser.VerificationCode {
		return "", errors.New(ErrWrongCredentials)
	}

	return existedUser.PhoneNumber, nil
}

func generateSessionId(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

func generateAuthCode() (string, error) {
	maxInt := big.NewInt(10000)
	n, err := rand.Int(rand.Reader, maxInt)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%04d", n.Int64()), nil
}
