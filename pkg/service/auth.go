package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"music/pkg/repository"
	"time"
)

const (
	salt      = "lk1 24g120985y12qr;lkfwqnrfkj23b r21e[" + "] `"
	singInKey = ";1 2m12 opejm`12-0e ik123[r'"
	TokenTTL  = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	ArtistId int `json:"user_id"`
}

type AuthorizationService struct {
	repo repository.Authorization
}

func NewAuthorizationService(repo repository.Authorization) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (s *AuthorizationService) SingUp(username string, password string) error {
	password = generatePasswordHash(password)
	return s.repo.CreateArtist(username, password)
}

func (s *AuthorizationService) GenerateToken(username string, password string) (string, error) {
	artist, err := s.repo.GetArtist(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		artist.Id,
	})

	return token.SignedString([]byte(singInKey))
}

func (s *AuthorizationService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(singInKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("Invalid token")
	}

	return claims.ArtistId, nil
}

func generatePasswordHash(password string) string {
	hash := sha256.Sum256([]byte(password + salt))
	return hex.EncodeToString(hash[:])
}
