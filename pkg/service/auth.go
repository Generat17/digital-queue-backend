package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"server/pkg/repository"
	"server/types"
	"time"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type tokenClaimsWorkstation struct {
	jwt.StandardClaims
	UserId        int `json:"user_id"`
	WorkstationId int `json:"workstation_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateEmployee(employee types.Employee) (int, error) {
	employee.Password = generatePasswordHash(employee.Password)
	return s.repo.CreateEmployee(employee)
}

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	employee, err := s.repo.GetEmployee(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		employee.EmployeeId,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) GenerateTokenWorkstation(login, password string, workstation int) (string, error) {
	employee, err := s.repo.GetEmployee(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaimsWorkstation{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		employee.EmployeeId,
		workstation,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
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

func (s *AuthService) ParseTokenWorkstation(accessToken string) (types.ParseTokenWorkstationResponse, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaimsWorkstation{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return types.ParseTokenWorkstationResponse{UserId: 0, WorkstationId: 0}, err
	}

	claims, ok := token.Claims.(*tokenClaimsWorkstation)
	if !ok {
		return types.ParseTokenWorkstationResponse{UserId: 0, WorkstationId: 0}, errors.New("token claims are not of type *tokenClaims")
	}

	return types.ParseTokenWorkstationResponse{UserId: claims.UserId, WorkstationId: claims.WorkstationId}, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
