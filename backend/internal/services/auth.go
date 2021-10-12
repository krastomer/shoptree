package services

import (
	"net/mail"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/krastomer/shoptree/backend/internal/entities"
)

type authService struct {
	repo entities.CustomerRepo
}

type jwtClaims struct {
	UserID string `json:"uid"`
	User   string `json:"user"`
	jwt.StandardClaims
}

func NewAuthService(repo entities.CustomerRepo) entities.AuthService {
	return &authService{repo: repo}
}

func (s *authService) LoginCustomer(u string, p string) (string, error) {
	if _, err := mail.ParseAddress(u); err != nil {
		return "", err
	}

	cust, err := s.repo.GetCustomer(u)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClaims{
		strconv.Itoa(cust.ID),
		cust.Email,
		jwt.StandardClaims{
			Audience:  "test",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    "testets",
		},
	})

	signedToken, err := token.SignedString([]byte("september"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
