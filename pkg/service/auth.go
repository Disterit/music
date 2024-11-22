package service

import (
	"crypto/sha256"
	"encoding/hex"
	"music/pkg/repository"
)

const (
	salt = "lk1 24g120985y12qr;lkfwqnrfkj23b r21e[" + "] `"
)

type AuthorizationService struct {
	repo repository.Authorization
}

func NewAuthorizationService(repo repository.Authorization) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (s *AuthorizationService) SingUp(username string, password string) error {
	password = generatePasswordHash(password)
	return s.repo.SingUp(username, password)
}

func generatePasswordHash(password string) string {
	hash := sha256.Sum256([]byte(password + salt))
	return hex.EncodeToString(hash[:])
}
