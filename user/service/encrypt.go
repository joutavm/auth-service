package service

import (
	"crypto/sha256"
	"fmt"
)

type EncryptService struct {
}

func (EncryptService EncryptService) SHA256Encoder(base string) string {
	parsed := sha256.Sum256([]byte(base))
	return fmt.Sprintf("%x", parsed)
}

func InitEncryptService() *EncryptService {
	return &EncryptService{}
}
