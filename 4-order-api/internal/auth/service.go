package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"go/order-api/internal/user"
)

const (
	DefaultAuthCode     = "0000"
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

func (service *AuthService) CreateNewSessionId(phone string) (string, error) {
	existedUser, _ := service.UserRepository.FindByPhone(phone)
	sessionId, err := generateSessionId(DefaultSessIdLength)
	if err != nil {
		return "", err
	}
	if existedUser != nil {
		existedUser.SessionId = sessionId
		_, err = service.UserRepository.Update(existedUser)
		if err != nil {
			return "", err
		}

		return sessionId, nil
	}

	existedUser = user.NewUser(phone, sessionId)
	_, err = service.UserRepository.Create(existedUser)
	if err != nil {
		return "", err
	}

	return sessionId, nil
}

func (service *AuthService) VerifyByCode(sessionId, code string) (string, error) {
	existedUser, _ := service.UserRepository.FindBySessionId(sessionId)
	if existedUser == nil {
		return "", errors.New(ErrWrongCredentials)
	}

	if code != DefaultAuthCode {
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
