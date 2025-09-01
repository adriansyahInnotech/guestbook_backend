package utils

import (
	"guestbook_backend/repository"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type ApiKey struct {
	EnvKey              string
	GuestbookRepository *repository.GuestbookRepository
}

func NewApiKey() *ApiKey {
	return &ApiKey{
		EnvKey:              os.Getenv("SERVICE_API_KEY"),
		GuestbookRepository: repository.NewGuestbookRepository(),
	}
}

func (s *ApiKey) GenerateApiKey() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(s.EnvKey), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s *ApiKey) ValidateApiKey(prefix string, hash string) bool {

	device, err := s.GuestbookRepository.DeviceRepository.GetApiKey(prefix)
	if err != nil {
		return false
	}

	if hash != device.ApiKey {
		return false
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(s.EnvKey)); err != nil {
		return false
	}

	return true

}

func (s *ApiKey) SplitAPIKey(token string) (prefix string, key string, ok bool) {
	// Cek apakah token mengandung bcrypt signature
	idx := strings.Index(token, "$2")
	if idx != -1 {
		// format bcrypt
		prefix = token[:idx-1] // buang underscore terakhir
		key = token[idx:]
		return prefix, key, true
	}

	// fallback plain format → pisahkan prefix dan key
	parts := strings.SplitN(token, "_", 2)
	if len(parts) != 2 {
		return "", "", false
	}
	return parts[0], parts[1], true
}
