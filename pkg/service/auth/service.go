package auth

import (
	"RESTAPIService2/pkg/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	salt       = "hfwf543htrrj64gred"
	signingKey = "fgwrf2#trg%lGr#jkg6e6ry#JDifgej"
	tokenTTl   = 12 * time.Hour
)

type AuthorizationService interface {
	CreateUser(user User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ImplAuthorizationService struct {
	repo repository.AuthorizationRepository
}

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

func NewAuthorizationService(repo repository.AuthorizationRepository) *ImplAuthorizationService {
	return &ImplAuthorizationService{repo: repo}
}

func (s *ImplAuthorizationService) CreateUser(user User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.Create(user)
}

func (s *ImplAuthorizationService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.Get(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *ImplAuthorizationService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
