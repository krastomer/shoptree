package services

import (
	"net/mail"
	"strconv"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt"
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/errors"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repo entities.AuthRepo
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
		return "", errors.ErrEmailInvalid
	}

	if err := checkPasswordValid(p); err != nil {
		return "", errors.ErrPasswordInvalid
	}

	cust, err := s.repo.GetCustomerByEmail(u)
	if err != nil {
		return "", errors.ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(cust.Password), []byte(p)); err != nil {
		return "", errors.ErrPasswordInvalid
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClaims{
		strconv.Itoa(cust.ID),
		cust.Email,
		jwt.StandardClaims{
			Audience:  "shoptree-customer",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    "shoptree",
		},
	})

	signedToken, err := token.SignedString([]byte("september"))
	if err != nil {
		return "", errors.ErrInternalServerError
	}

	return signedToken, nil
}

func checkPasswordValid(p string) error {
	if len(p) < 8 {
		return errors.ErrPasswordInvalid
	}
	number := false
	upper := false
	for _, c := range p {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsLetter(c) || c == ' ':
		default:
		}
	}
	if !(number && upper) {
		return errors.ErrPasswordInvalid
	}
	return nil
}
